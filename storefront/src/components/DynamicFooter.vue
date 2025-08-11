<template>
  <footer class="mt-auto" :style="footerStyle">
    <div :class="containerClasses" class="px-6 sm:px-8">
      <!-- Top Section: Links and Shop Info -->
      <div class="py-16">
        <div class="grid gap-12" :style="gridStyle">
          <!-- Shop Info & Tagline -->
          <div class="space-y-4 pr-8">
            <div class="flex items-center gap-x-3">
              <img
                v-if="shop?.image"
                :src="shop.image"
                :alt="`${shop.name} Logo`"
                class="h-10 w-10 object-contain rounded-full shadow-sm"
                :style="{ backgroundColor: theme?.colors.primary + '20' }"
              />
              <h3 class="text-xl font-bold" :style="headingStyle">
                {{ shop?.name || 'Our Store' }}
              </h3>
            </div>
            <p class="text-sm leading-relaxed" :style="textStyle">
              ✨ Spreading joy with every order. Discover your new favorite things with us!
            </p>
          </div>

          <!-- Dynamic Link Columns -->
          <div
            v-for="column in linkColumns"
            :key="column.title"
            class="space-y-4"
          >
            <h4
              class="font-semibold text-sm uppercase tracking-wide"
              :style="headingStyle"
            >
              {{ column.title }}
            </h4>
            <nav class="space-y-2">
              <router-link
                v-for="link in column.links"
                :key="link.path"
                :to="{ path: link.path }"
                class="block text-sm transition-opacity hover:opacity-75"
                :style="linkStyle"
              >
                {{ link.label }}
              </router-link>
            </nav>
          </div>
        </div>
      </div>

      <!-- Bottom Section: Social Links & Copyright -->
      <div class="py-8 border-t" :style="{ borderTopColor: theme?.colors.border }">
        <div class="flex flex-col sm:flex-row items-center justify-between gap-y-6">
          <!-- Social Media Icons -->
          <div v-if="footerConfig?.showSocial && socialLinks.length" class="flex items-center gap-x-3">
            <a
              v-for="social in socialLinks"
              :key="social.name"
              :href="social.url"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-full flex items-center justify-center transition-all duration-300"
              :style="socialIconStyle"
              :aria-label="`Find us on ${social.name}`"
            >
              <component :is="social.icon" class="w-5 h-5" />
            </a>
          </div>

          <!-- Copyright Notice -->
          <p class="text-xs" :style="textStyle">
            &copy; {{ currentYear }} {{ shop?.name }}. Lovingly built with Vue.
          </p>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { ref, computed, defineComponent } from 'vue';
// Import types from your actual theme service file
import type { ShopInfo, DynamicTheme, FooterConfig } from '@/services/dynamic-theme';

// --- Mock Icons (Replace with your actual icon components) ---
const HeartIcon = defineComponent({ template: '<span>❤️</span>' });
const StarIcon = defineComponent({ template: '<span>⭐</span>' });
const SmileIcon = defineComponent({ template: '<span>😊</span>' });

// --- Component Props ---
interface Props {
  shop: ShopInfo | null;
  theme: DynamicTheme | null;
  footerConfig: FooterConfig | null;
}
const props = defineProps<Props>();

const currentYear = ref(new Date().getFullYear());

// --- Data Driven Content ---
// This data should ideally come from the footerConfig prop in a real application.
const linkColumns = computed(() => (props.footerConfig as any)?.linkColumns || [
  {
    title: 'Explore',
    links: [
      { label: 'New Arrivals', path: '/products?sort=newest' },
      { label: 'Best Sellers', path: '/collections/best-sellers' },
      { label: 'Special Offers', path: '/sale' },
    ],
  },
  {
    title: 'Our World',
    links: [
      { label: 'About Us', path: '/about' },
      { label: 'Contact', path: '/contact' },
      { label: 'FAQs', path: '/faq' },
    ],
  },
]);

const socialLinks = computed(() => (props.footerConfig as any)?.socialLinks || [
  { name: 'Love', url: '#', icon: HeartIcon },
  { name: 'Favorites', url: '#', icon: StarIcon },
  { name: 'Community', url: '#', icon: SmileIcon },
]);

// --- Computed Styles (Driven by the Theme Service) ---

const footerStyle = computed(() => ({
  backgroundColor: props.theme?.colors.background ? `${props.theme.colors.background}E6` : '#FFFFFF',
  color: props.theme?.colors.bodyText || '#4A5568',
}));

const containerClasses = computed(() => {
  switch (props.theme?.layout.containerWidth) {
    case 'full': return 'max-w-none';
    case 'wide': return 'max-w-screen-xl mx-auto';
    default: return 'max-w-6xl mx-auto';
  }
});

const gridStyle = computed(() => ({
  '--grid-columns': props.footerConfig?.columns || 3,
}));

const headingStyle = computed(() => ({
  color: props.theme?.colors.heading || '#1A202C',
  fontFamily: props.theme?.fonts.heading ? `'${props.theme.fonts.heading}', sans-serif` : 'sans-serif',
}));

const textStyle = computed(() => ({
  color: props.theme?.colors.bodyText || '#4A5568',
  fontFamily: props.theme?.fonts.body ? `'${props.theme.fonts.body}', sans-serif` : 'sans-serif',
}));

const linkStyle = computed(() => ({
  color: props.theme?.colors.bodyText || '#4A5568',
  fontFamily: props.theme?.fonts.body ? `'${props.theme.fonts.body}', sans-serif` : 'sans-serif',
}));

const socialIconStyle = computed(() => {
  const primaryColor = props.theme?.colors.primary || '#4299E1';
  const textColor = props.theme?.colors.background || '#FFFFFF';
  return {
    backgroundColor: `${primaryColor}20`,
    color: primaryColor,
    '--hover-bg': primaryColor,
    '--hover-color': textColor,
  };
});
</script>

<style scoped>
.grid {
  display: grid;
  /* Mobile-first: 1 column */
  grid-template-columns: repeat(1, minmax(0, 1fr));
}

/* Apply dynamic grid columns on larger screens */
@media (min-width: 768px) {
  .grid {
    grid-template-columns: repeat(var(--grid-columns, 3), minmax(0, 1fr));
  }
}

/* Cute hover effect for social icons */
a[aria-label^="Find us on"] {
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

a[aria-label^="Find us on"]:hover {
  background-color: var(--hover-bg) !important;
  color: var(--hover-color) !important;
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 6px 12px rgba(0,0,0,0.1);
}
</style>
