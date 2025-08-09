<template>
  <nav class="mb-6" aria-label="Breadcrumb">
    <ol class="flex items-center text-sm text-gray-500">
      <li v-if="showBack" class="mr-3">
        <button @click="$router.back()" class="inline-flex items-center gap-1 text-gray-600 hover:text-black font-medium px-2 py-1 rounded transition-colors" aria-label="Go back">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
          <span class="uppercase tracking-wide text-xs">Back</span>
        </button>
      </li>
      <li v-for="(item, idx) in normalizedItems" :key="idx" class="flex items-center">
        <template v-if="idx > 0">
          <span class="mx-2 text-gray-300">/</span>
        </template>
        <template v-if="item.to && idx !== normalizedItems.length - 1">
          <router-link :to="item.to" class="text-gray-600 hover:text-black hover:underline transition-colors">{{ item.label }}</router-link>
        </template>
        <template v-else>
          <span class="text-gray-900 font-semibold">{{ item.label }}</span>
        </template>
      </li>
    </ol>
  </nav>
  </template>

<script setup lang="ts">
import { defineProps, computed } from 'vue'

interface BreadcrumbItem {
  label: string
  to?: string
  back?: boolean // If true, renders a back button
}

const props = defineProps<{ items: BreadcrumbItem[] }>()

const showBack = computed(() => props.items.length > 0 && !!props.items[0].back)
const normalizedItems = computed(() => {
  // Filter out the first back item for breadcrumb trail
  const startIdx = showBack.value ? 1 : 0
  return props.items.slice(startIdx)
})
</script> 