<template>
  <footer 
    class="border-t transition-colors"
    :style="footerStyle"
    :class="footerClasses"
  >
    <div :class="containerClasses">
      
      <!-- Main Footer Content -->
      <div class="py-12">
        <div :class="gridClasses">
          
          <!-- Shop Information Column -->
          <div class="space-y-4">
            <div class="flex items-center space-x-3">
              <img
                v-if="shop?.image"
                :src="shop.image"
                :alt="shop.name + ' Logo'"
                class="h-8 w-8 object-contain"
              />
              <h3 
                class="text-lg font-semibold tracking-tight uppercase"
                :style="headingStyle"
              >
                {{ shop?.name || 'Store' }}
              </h3>
            </div>
            
            <p 
              class="text-sm leading-relaxed max-w-sm"
              :style="textStyle"
            >
              {{ shop?.description || 'Discover our amazing products and exceptional quality.' }}
            </p>
            
            <!-- Contact Information -->
            <div class="space-y-2">
              <div v-if="shop?.email" class="flex items-center space-x-2 text-sm" :style="textStyle">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
                <span>{{ shop.email }}</span>
              </div>
              
              <div v-if="shop?.phone" class="flex items-center space-x-2 text-sm" :style="textStyle">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                </svg>
                <span>{{ shop.phone }}</span>
              </div>
              
              <div v-if="shop?.address" class="flex items-start space-x-2 text-sm" :style="textStyle">
                <svg class="w-4 h-4 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                <span>{{ shop.address }}</span>
              </div>
            </div>
          </div>

          <!-- Quick Links Column -->
          <div class="space-y-4">
            <h4 
              class="font-semibold text-sm uppercase tracking-wide"
              :style="headingStyle"
            >
              Quick Links
            </h4>
            <nav class="space-y-2">
              <router-link
                v-for="link in quickLinks"
                :key="link.path"
                :to="{ path: link.path }"
                class="block text-sm hover:opacity-70 transition-opacity"
                :style="linkStyle"
              >
                {{ link.label }}
              </router-link>
            </nav>
          </div>

          <!-- Customer Service Column -->
          <div class="space-y-4">
            <h4 
              class="font-semibold text-sm uppercase tracking-wide"
              :style="headingStyle"
            >
              Customer Service
            </h4>
            <nav class="space-y-2">
              <router-link
                v-for="link in serviceLinks"
                :key="link.path"
                :to="{ path: link.path }"
                class="block text-sm hover:opacity-70 transition-opacity"
                :style="linkStyle"
              >
                {{ link.label }}
              </router-link>
            </nav>
          </div>

          <!-- Newsletter Signup Column (if enabled) -->
          <div v-if="footerConfig?.showNewsletter" class="space-y-4">
            <h4 
              class="font-semibold text-sm uppercase tracking-wide"
              :style="headingStyle"
            >
              Stay Updated
            </h4>
            <p class="text-sm" :style="textStyle">
              Subscribe to get special offers and updates.
            </p>
            
            <form @submit.prevent="subscribeToNewsletter" class="space-y-3">
              <div class="relative">
                <input
                  type="email"
                  v-model="newsletterEmail"
                  placeholder="Enter your email"
                  class="w-full px-3 py-2 text-sm border rounded-lg focus:outline-none focus:ring-2 transition-all"
                  :style="inputStyle"
                  required
                />
              </div>
              <button
                type="submit"
                class="w-full px-4 py-2 text-sm font-medium rounded-lg transition-all duration-200 hover:opacity-90"
                :style="buttonStyle"
                :disabled="isSubscribing"
              >
                {{ isSubscribing ? 'Subscribing...' : 'Subscribe' }}
              </button>
            </form>
          </div>
        </div>
      </div>

      <!-- Social Media & Payment Methods (if enabled) -->
      <div 
        v-if="footerConfig?.showSocial || showPaymentMethods"
        class="py-6 border-t"
        :style="{ borderTopColor: theme?.colors?.border || '#E5E7EB' }"
      >
        <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0">
          
          <!-- Social Media Links -->
          <div v-if="footerConfig?.showSocial" class="flex items-center space-x-4">
            <span class="text-sm font-medium" :style="textStyle">Follow us:</span>
            <div class="flex items-center space-x-3">
              <a
                v-for="social in socialLinks"
                :key="social.name"
                :href="social.url"
                target="_blank"
                rel="noopener noreferrer"
                class="p-2 rounded-full transition-all duration-200 hover:scale-110"
                :style="socialButtonStyle"
                :aria-label="social.name"
              >
                <component :is="social.icon" class="w-4 h-4" />
              </a>
            </div>
          </div>

          <!-- Payment Methods -->
          <div v-if="showPaymentMethods" class="flex items-center space-x-2">
            <span class="text-sm font-medium" :style="textStyle">We accept:</span>
            <div class="flex items-center space-x-2">
              <div
                v-for="payment in paymentMethods"
                :key="payment.name"
                class="px-2 py-1 bg-white border rounded text-xs font-medium"
                :style="paymentBadgeStyle"
              >
                {{ payment.name }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Copyright -->
      <div 
        class="py-6 border-t text-center"
        :style="{ borderTopColor: theme?.colors?.border || '#E5E7EB' }"
      >
        <p class="text-sm" :style="textStyle">
          © {{ currentYear }} {{ shop?.name || 'Store' }}. All rights reserved.
          <span v-if="theme?.version" class="opacity-60 ml-2">
            • Theme v{{ theme.version }}
          </span>
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { ShopInfo, DynamicTheme, FooterConfig } from '@/services/dynamic-theme'

// =============== Props ===============

interface Props {
  shop?: ShopInfo | null
  theme?: DynamicTheme | null
  footerConfig?: FooterConfig | null
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const newsletterEmail = ref('')
const isSubscribing = ref(false)
const currentYear = new Date().getFullYear()
// const subscriptionSuccess = ref(false) // Unused for now

// =============== Footer Configuration ===============

const quickLinks = [
  { label: 'Home', path: '/' },
  { label: 'Products', path: '/products' },
  { label: 'Collections', path: '/collections' },
  { label: 'About Us', path: '/about' },
]

const serviceLinks = [
  { label: 'Contact Us', path: '/contact' },
  { label: 'Shipping Info', path: '/shipping' },
  { label: 'Returns', path: '/returns' },
  { label: 'Size Guide', path: '/size-guide' },
  { label: 'FAQ', path: '/faq' },
]

const socialLinks = [
  { name: 'Facebook', url: '#', icon: 'FacebookIcon' },
  { name: 'Instagram', url: '#', icon: 'InstagramIcon' },
  { name: 'Twitter', url: '#', icon: 'TwitterIcon' },
]

const paymentMethods = [
  { name: 'Visa' },
  { name: 'Mastercard' },
  { name: 'PayPal' },
  { name: 'Apple Pay' },
]

const showPaymentMethods = true

// =============== Computed Styles ===============

const footerStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#F9FAFB', borderTopColor: '#E5E7EB' }
  
  return {
    backgroundColor: props.theme.colors?.background || '#F9FAFB',
    borderTopColor: props.theme.colors?.border || '#E5E7EB'
  }
})

const footerClasses = computed(() => {
  const classes = ['mt-auto']
  
  // Add style-specific classes based on footer config
  if (props.footerConfig?.style === 'minimal') {
    classes.push('py-8')
  } else if (props.footerConfig?.style === 'full') {
    classes.push('py-12')
  } else {
    classes.push('py-10')
  }
  
  return classes
})

const containerClasses = computed(() => {
  const classes = ['mx-auto', 'px-4', 'sm:px-6']
  
  if (props.theme?.layout?.containerWidth === 'full') {
    classes.push('max-w-none')
  } else if (props.theme?.layout?.containerWidth === 'wide') {
    classes.push('max-w-7xl')
  } else {
    classes.push('max-w-6xl')
  }
  
  return classes
})

const gridClasses = computed(() => {
  const columns = props.footerConfig?.columns || 4
  
  return [
    'grid',
    'gap-8',
    `grid-cols-1`,
    `md:grid-cols-${Math.min(columns, 4)}`,
    'lg:gap-12'
  ]
})

const headingStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }
  
  return {
    color: props.theme.colors?.heading || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const textStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    fontFamily: props.theme.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
  }
})

const linkStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280'
  }
})

const inputStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    borderColor: props.theme.colors?.border || '#E5E7EB',
    backgroundColor: props.theme.colors?.background || '#ffffff',
    color: props.theme.colors?.bodyText || '#6B7280',
    '--tw-ring-color': props.theme.colors?.primary || '#10B981'
  }
})

const buttonStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#10B981', color: 'white' }
  
  return {
    backgroundColor: props.theme.colors?.primary || '#10B981',
    color: props.theme.colors?.background || 'white',
    borderRadius: props.theme.layout?.borderStyle === 'sharp' ? '0px' : 
                 props.theme.layout?.borderStyle === 'pill' ? '50px' : '8px'
  }
})

const socialButtonStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    backgroundColor: `${props.theme.colors?.primary || '#10B981'}10`,
    '--hover-bg': `${props.theme.colors?.primary || '#10B981'}20`
  }
})

const paymentBadgeStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    borderColor: props.theme.colors?.border || '#E5E7EB',
    color: props.theme.colors?.bodyText || '#6B7280'
  }
})

// =============== Methods ===============

async function subscribeToNewsletter() {
  if (!newsletterEmail.value) return
  
  isSubscribing.value = true
  
  try {
    // TODO: Implement newsletter subscription
    await new Promise(resolve => setTimeout(resolve, 1000)) // Simulate API call
    
    console.log('Newsletter subscription:', newsletterEmail.value)
    newsletterEmail.value = ''
    
    // Show success message
    alert('Successfully subscribed to newsletter!')
    
  } catch (error) {
    console.error('Newsletter subscription error:', error)
    alert('Failed to subscribe. Please try again.')
  } finally {
    isSubscribing.value = false
  }
}
</script>

<style scoped>
/* Custom grid classes for dynamic columns */
.grid-cols-1 { grid-template-columns: repeat(1, minmax(0, 1fr)); }
.md\:grid-cols-2 { grid-template-columns: repeat(2, minmax(0, 1fr)); }
.md\:grid-cols-3 { grid-template-columns: repeat(3, minmax(0, 1fr)); }
.md\:grid-cols-4 { grid-template-columns: repeat(4, minmax(0, 1fr)); }

@media (min-width: 768px) {
  .md\:grid-cols-2 { grid-template-columns: repeat(2, minmax(0, 1fr)); }
  .md\:grid-cols-3 { grid-template-columns: repeat(3, minmax(0, 1fr)); }
  .md\:grid-cols-4 { grid-template-columns: repeat(4, minmax(0, 1fr)); }
}

/* Hover effects */
.hover\:scale-110:hover {
  transform: scale(1.1);
}

.hover\:opacity-70:hover {
  opacity: 0.7;
}

/* Focus rings */
input:focus {
  ring-width: 2px;
  ring-color: var(--tw-ring-color);
}

/* Button hover effects */
button:hover:not(:disabled) {
  opacity: 0.9;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
