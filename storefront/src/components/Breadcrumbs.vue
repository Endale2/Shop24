<template>
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <template v-for="(item, idx) in items" :key="idx">
      <template v-if="idx === 0">
        <button v-if="item.back" @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
          Back
        </button>
      </template>
      <template v-else>
        <span>/</span>
      </template>
      <template v-if="item.to && idx !== items.length - 1">
        <router-link :to="item.to" class="hover:underline">{{ item.label }}</router-link>
      </template>
      <template v-else>
        <span class="text-gray-900 font-medium">{{ item.label }}</span>
      </template>
    </template>
  </nav>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface BreadcrumbItem {
  label: string
  to?: string
  back?: boolean // If true, renders a back button
}

defineProps<{ items: BreadcrumbItem[] }>()
</script> 