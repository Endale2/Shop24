<template>
  <router-link
    :to="`/collections/${collection.handle}`"
    class="collection-card group relative overflow-hidden transition-all duration-300 hover:scale-105"
    :class="cardClasses"
    :style="cardStyle"
  >
    <!-- Collection Image -->
    <div
      class="relative overflow-hidden"
      :class="imageContainerClasses"
    >
      <img
        :src="collection.image || collection.featured_image || defaultCollectionImage"
        :alt="collection.name"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
        loading="lazy"
      />

      <!-- Overlay -->
      <div
        class="absolute inset-0 bg-gradient-to-t from-black/60 via-black/20 to-transparent"
      ></div>

      <!-- Product Count Badge -->
      <div
        v-if="collection.product_count > 0"
        class="absolute top-3 right-3 px-2 py-1 text-xs font-bold rounded-full"
        :style="badgeStyle"
      >
        {{ collection.product_count }} {{ collection.product_count === 1 ? 'item' : 'items' }}
      </div>
    </div>

    <!-- Collection Info Overlay -->
    <div class="absolute inset-0 flex items-end p-6 z-10">
      <div class="space-y-2 w-full">
        <h3
          class="font-bold tracking-tight uppercase leading-tight"
          :style="collectionNameStyle"
          :class="collectionNameClasses"
        >
          {{ collection.name }}
        </h3>

        <p
          v-if="collection.description"
          class="text-sm line-clamp-2 opacity-90"
          :style="collectionDescriptionStyle"
        >
          {{ collection.description }}
        </p>

        <!-- Call to Action -->
        <div class="pt-2">
          <span
            class="inline-flex items-center text-sm font-medium transition-all duration-200 group-hover:translate-x-1"
            :style="ctaStyle"
          >
            Explore Collection
            <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
            </svg>
          </span>
        </div>
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { DynamicTheme } from '@/services/dynamic-theme'

// =============== Props ===============

interface Collection {
  id: string
  name: string
  handle: string
  description?: string
  image?: string
  featured_image?: string
  product_count: number
}

interface Props {
  collection: Collection
  theme?: DynamicTheme | null
}

const props = defineProps<Props>()

// =============== Constants ===============

const defaultCollectionImage = 'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=800&h=600&fit=crop'

// =============== Computed Styles ===============

// Card styling
const cardStyle = computed(() => {
  if (!props.theme) return {}

  return {
    borderRadius: getBorderRadiusValue(props.theme.layout?.borderStyle || 'rounded'),
    boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)'
  }
})

const cardClasses = computed(() => {
  const classes = ['block']

  // Add border radius class
  const borderStyle = props.theme?.layout?.borderStyle || 'rounded'
  if (borderStyle === 'sharp') {
    classes.push('rounded-none')
  } else if (borderStyle === 'pill') {
    classes.push('rounded-3xl')
  } else {
    classes.push('rounded-lg')
  }

  return classes
})

// Image container classes
const imageContainerClasses = computed(() => {
  return ['aspect-[4/3]'] // Default aspect ratio for collections
})

// Collection name styling
const collectionNameStyle = computed(() => {
  if (!props.theme) return { color: 'white' }

  return {
    color: 'white', // Always white for better contrast over images
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined,
    textShadow: '0 2px 4px rgba(0,0,0,0.8)'
  }
})

const collectionNameClasses = computed(() => {
  switch (props.theme?.layout?.headerStyle) {
    case 'compact': return 'text-lg md:text-xl'
    case 'large': return 'text-2xl md:text-3xl'
    default: return 'text-xl md:text-2xl'
  }
})

// Collection description styling
const collectionDescriptionStyle = computed(() => {
  return {
    color: 'white',
    fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined,
    textShadow: '0 1px 2px rgba(0,0,0,0.8)'
  }
})

// Badge styling
const badgeStyle = computed(() => {
  return {
    backgroundColor: props.theme?.colors?.primary || '#10B981',
    color: props.theme?.colors?.background || 'white'
  }
})

// Call-to-action styling
const ctaStyle = computed(() => {
  return {
    color: props.theme?.colors?.primary || '#10B981',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    padding: '0.5rem 1rem',
    borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded'),
    backdropFilter: 'blur(10px)'
  }
})

// =============== Helper Functions ===============

function getBorderRadiusValue(borderStyle: string): string {
  switch (borderStyle) {
    case 'sharp': return '0px'
    case 'rounded': return '8px'
    case 'pill': return '50px'
    default: return '8px'
  }
}
</script>

<style scoped>
/* Collection card base styles */
.collection-card {
  display: block;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.collection-card:hover {
  transform: scale(1.05);
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

/* Aspect ratio utilities */
.aspect-\[4\/3\] {
  aspect-ratio: 4 / 3;
}

/* Text truncation */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Hover effects */
.group:hover .group-hover\:scale-110 {
  transform: scale(1.1);
}

.group:hover .group-hover\:translate-x-1 {
  transform: translateX(0.25rem);
}

/* Focus states */
.collection-card:focus {
  outline: 2px solid var(--color-primary, #10B981);
  outline-offset: 2px;
}

/* Image transitions */
img {
  transition: transform 0.5s ease;
}

/* Backdrop blur support */
.backdrop-blur {
  backdrop-filter: blur(10px);
}

/* Mobile optimizations */
@media (max-width: 640px) {
  .collection-card {
    margin: 0.5rem 0;
  }

  .collection-info {
    padding: 1rem;
  }
}

/* Dark theme support */
@media (prefers-color-scheme: dark) {
  .collection-card {
    border-color: rgba(255, 255, 255, 0.1);
  }
}
</style>
