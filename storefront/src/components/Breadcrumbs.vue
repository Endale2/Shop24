<template>
  <nav class="breadcrumb-nav mb-6 w-full" :style="navStyle" aria-label="Breadcrumb">
    <ol class="flex items-center text-sm flex-wrap gap-y-1">
      <!-- Back Button -->
      <li v-if="showBack" class="flex items-center mr-3">
        <button
          @click="$router.back()"
          class="breadcrumb-back-btn inline-flex items-center gap-1 px-3 py-1 rounded-full transition-all duration-200 hover:scale-105"
          :style="backButtonStyle"
          aria-label="Go back"
        >
          <svg
            class="w-4 h-4"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
          </svg>
          <span class="uppercase tracking-wide text-xs font-semibold">Back</span>
        </button>
      </li>

      <!-- Breadcrumb Items -->
      <li v-for="(item, idx) in normalizedItems" :key="idx" class="flex items-center">
        <template v-if="idx > 0">
          <svg
            class="w-4 h-4 mx-2"
            :style="dividerStyle"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </template>
        <template v-if="item.to && idx !== normalizedItems.length - 1">
          <router-link
            :to="item.to"
            class="breadcrumb-link transition-colors duration-200 hover:text-opacity-80 font-medium"
            :style="linkStyle"
          >
            {{ item.label }}
          </router-link>
        </template>
        <template v-else>
          <span
            class="breadcrumb-active font-semibold"
            :style="activeStyle"
            :aria-current="item.label === lastItemLabel ? 'page' : undefined"
          >
            {{ item.label }}
          </span>
        </template>
      </li>
    </ol>
  </nav>
</template>

<script setup lang="ts">
import { defineProps, computed } from 'vue'
import type { DynamicTheme } from '@/services/dynamic-theme'

interface BreadcrumbItem {
  label: string
  to?: string
  back?: boolean
}

interface Props {
  items: BreadcrumbItem[]
  theme?: DynamicTheme | null
}

const props = defineProps<Props>()

// Excludes the 'back' item from the main list if present.
const showBack = computed(() => props.items.length > 0 && !!props.items[0].back)
const normalizedItems = computed(() => {
  const startIdx = showBack.value ? 1 : 0
  return props.items.slice(startIdx)
})

// Get the label of the last item for aria-current
const lastItemLabel = computed(() => {
  if (normalizedItems.value.length > 0) {
    return normalizedItems.value[normalizedItems.value.length - 1].label
  }
  return null
})

// --- Dynamic Styles from Theme Service ---

const navStyle = computed(() => ({
  fontFamily: props.theme?.fonts?.body ? `'${props.theme.fonts.body}', sans-serif` : undefined
}))

const linkStyle = computed(() => ({
  color: props.theme?.colors?.bodyText || '#6B7280'
}))

const dividerStyle = computed(() => ({
  color: props.theme?.colors?.border || '#E5E7EB'
}))

const activeStyle = computed(() => ({
  color: props.theme?.colors?.heading || '#1F2937'
}))

const backButtonStyle = computed(() => ({
  backgroundColor: props.theme?.colors?.background || 'white',
  color: props.theme?.colors?.bodyText || '#4B5563',
  border: `1px solid ${props.theme?.colors?.border || '#D1D5DB'}`,
  borderRadius: getBorderRadiusValue(props.theme?.layout?.borderStyle || 'rounded')
}))

function getBorderRadiusValue(borderStyle: string): string {
  switch (borderStyle) {
    case 'sharp': return '0px'
    case 'rounded': return '9999px' // Use a large value for pill shape
    case 'pill': return '9999px'
    default: return '9999px'
  }
}
</script>

