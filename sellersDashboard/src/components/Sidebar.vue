<template>
  <aside
    class="flex h-full w-64 flex-col bg-white shadow-xl border-r border-gray-200"
  >
    <!-- Header Section -->
    <div
      class="flex h-20 items-center justify-between px-4 border-b border-gray-200 bg-gradient-to-r from-gray-50 to-white"
    >
      <div class="flex items-center space-x-3">
        <div class="relative">
          <img
            v-if="activeShop?.image"
            :src="activeShop.image"
            alt="Shop Logo"
            class="h-12 w-12 rounded-xl object-cover shadow-md ring-2 ring-white"
          />
          <div
            v-else
            class="h-12 w-12 bg-gradient-to-br from-green-500 to-emerald-600 rounded-xl flex items-center justify-center shadow-md ring-2 ring-white"
          >
            <ShopIcon class="h-6 w-6 text-white" />
          </div>
          <div class="absolute -bottom-1 -right-1 w-4 h-4 bg-green-500 rounded-full border-2 border-white shadow-sm"></div>
        </div>
        <div class="flex-1 min-w-0">
          <h2 class="text-sm font-bold text-gray-900 truncate">{{ activeShop?.name || 'Shop Name' }}</h2>
          <p class="text-xs text-gray-500 truncate" :title="activeShop?.description">
            {{ truncateDescription(activeShop?.description) }}
          </p>
        </div>
      </div>
    </div>

    <!-- Main Navigation -->
    <nav class="flex-1 overflow-y-auto py-4 px-3 space-y-1">
      <div class="mb-4">
        <p class="text-xs font-semibold uppercase text-gray-400 tracking-wider px-3 mb-2">
          Main Menu
        </p>
        <div class="space-y-1">
          <SidebarLink
            v-for="item in mainLinks"
            :key="item.to"
            :to="item.to"
            :icon="item.icon"
            :label="item.label"
            :badge="item.badge"
            compact
            large-icon
          />
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="mb-4">
        <p class="text-xs font-semibold uppercase text-gray-400 tracking-wider px-3 mb-2">
          Quick Actions
        </p>
        <div class="space-y-1">
          <SidebarLink
            to="/dashboard/products/add"
            :icon="PlusIcon"
            label="Add Product"
            compact
            variant="success"
          />
          <SidebarLink
            to="/dashboard/discounts/create"
            :icon="GiftIcon"
            label="Create Discount"
            compact
            variant="warning"
          />
          <SidebarLink
            to="/dashboard/media"
            :icon="PhotographIcon"
            label="Media Library"
            compact
            variant="info"
          />
        </div>
      </div>

      <!-- Collections Section -->
      <div
        v-if="collections.length"
        class="mb-4"
      >
        <p class="text-xs font-semibold uppercase text-gray-400 tracking-wider px-3 mb-2">
          Collections
        </p>
        <div class="space-y-1">
          <SidebarLink
            v-for="col in collections.slice(0, 3)"
            :key="col.id"
            :to="`/dashboard/collections/${col.id}`"
            :icon="CollectionIcon"
            :label="col.title"
            compact
            variant="info"
          />
          <SidebarLink
            v-if="collections.length > 3"
            to="/dashboard/collections"
            :icon="ViewListIcon"
            :label="`View All (${collections.length})`"
            compact
            variant="secondary"
          />
        </div>
      </div>
    </nav>

    <!-- Bottom Section -->
    <div class="border-t border-gray-200 bg-gradient-to-r from-gray-50 to-white">
      <!-- Sales Channels -->
      <div class="p-3">
        <p class="text-xs font-semibold uppercase text-gray-400 tracking-wider mb-2">
          Sales Channels
        </p>
        <div class="space-y-1">
          <SidebarLink
            to="/storefront"
            :icon="GlobeAltIcon"
            label="Online Store"
            compact
            variant="primary"
          />
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useShopStore } from '@/store/shops'
import SidebarLink from './SidebarLink.vue'

import {
  HomeIcon,
  ClipboardListIcon,
  TagIcon,
  UsersIcon,
  GiftIcon,
  ChartBarIcon,
  CogIcon,
  GlobeAltIcon,
  ShoppingBagIcon as ShopIcon,
  CollectionIcon,
  PlusIcon,
  ViewListIcon,
  PhotographIcon
} from '@heroicons/vue/outline'

const shopStore = useShopStore()
const activeShop = computed(() => shopStore.active)
const collections = computed(() => activeShop.value?.collections || [])

// Function to truncate long descriptions
function truncateDescription(description) {
  if (!description) return 'No description'
  return description.length > 25 ? description.substring(0, 25) + '...' : description
}

// Primary navigation items with badges - routes match router configuration
const mainLinks = [
  { 
    to: '/dashboard', 
    icon: HomeIcon, 
    label: 'Dashboard',
    badge: null
  },
  { 
    to: '/dashboard/orders', 
    icon: ClipboardListIcon, 
    label: 'Orders',
    badge: '12'
  },
  { 
    to: '/dashboard/products', 
    icon: TagIcon, 
    label: 'Products',
    badge: null
  },
  { 
    to: '/dashboard/collections', 
    icon: CollectionIcon, 
    label: 'Collections',
    badge: null
  },
  { 
    to: '/dashboard/customers', 
    icon: UsersIcon, 
    label: 'Customers',
    badge: '5'
  },
  { 
    to: '/dashboard/discounts', 
    icon: GiftIcon, 
    label: 'Discounts',
    badge: '3'
  },
  { 
    to: '/dashboard/analytics', 
    icon: ChartBarIcon, 
    label: 'Analytics',
    badge: null
  },
  { 
    to: '/dashboard/settings', 
    icon: CogIcon, 
    label: 'Settings',
    badge: null
  },
]
</script>

<style scoped>
/* Custom scrollbar for sidebar */
.overflow-y-auto::-webkit-scrollbar {
  width: 4px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Smooth transitions */
* {
  transition: all 0.2s ease-in-out;
}
</style>