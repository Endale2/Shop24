<!-- src/components/Sidebar.vue -->
<template>
  <aside class="flex flex-col h-full bg-white border-r border-gray-200">
    <!-- Header -->
    <div class="flex items-center px-4 py-6 border-b border-gray-200">
      <ShopIcon class="h-7 w-7 text-green-600 mr-2" />
      <span class="text-2xl font-bold text-gray-800">
        {{ activeShop?.name || '—' }}
      </span>
    </div>

    <!-- Main Links -->
    <nav class="flex-1 px-2 py-4 overflow-y-auto space-y-1">
      <SidebarLink to="/dashboard"        :icon="HomeIcon"          label="Home" />
      <SidebarLink to="/dashboard/orders"  :icon="ClipboardListIcon" label="Orders" />
      <SidebarLink to="/dashboard/products":icon="TagIcon"           label="Products" />
      <SidebarLink to="/dashboard/collections":icon="CollectionIcon"  label="Collections" />
      <SidebarLink to="/dashboard/customers":icon="UsersIcon"         label="Customers" />
      <SidebarLink to="/dashboard/discounts":icon="GiftIcon"          label="Discounts" />
      <SidebarLink to="/dashboard/analytics":icon="ChartBarIcon"      label="Analytics" />
      <SidebarLink to="/dashboard/settings" :icon="CogIcon"           label="Settings" />

      <!-- Collections Section -->
      <div v-if="collections.length" class="mt-6 pt-4 border-t border-gray-200">
        <p class="text-xs font-semibold text-gray-500 uppercase px-3 mb-2">
          Collections
        </p>
        <SidebarLink
          v-for="col in collections"
          :key="col.id"
          :to="`/dashboard/collections/${col.id}`"
          :icon="CollectionIcon"
          :label="col.name"
          small
        />
      </div>
    </nav>

    <!-- Sales channels -->
    <div class="px-4 py-4 border-t border-gray-200">
      <p class="text-xs font-semibold text-gray-500 uppercase mb-2">
        Sales Channels
      </p>
      <SidebarLink
        to="/dashboard/online-store"
        :icon="GlobeAltIcon"
        label="Online Store"
        small
      />
    </div>
  </aside>
</template>

<script setup>
import { computed } from 'vue'
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
  CollectionIcon  
} from '@heroicons/vue/outline'

// pinia store
const shopStore = useShopStore()
const activeShop = computed(() => shopStore.active)
const collections = computed(() => activeShop.value?.collections || [])
</script>
