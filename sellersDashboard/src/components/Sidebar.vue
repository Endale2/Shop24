<template>
  <aside
    class="flex h-full w-64 flex-col bg-white shadow-lg border-r border-slate-200"
  >
    <div
      class="flex h-16 items-center justify-center border-b border-slate-200 p-2"
    >
      <img
        v-if="activeShop?.image"
        :src="activeShop.image"
        alt="Shop Logo"
        class="h-10 w-10 rounded-md object-cover"
      />
      <ShopIcon v-else class="h-10 w-10 text-slate-400" />
    </div>

    <nav class="flex-1 overflow-y-auto py-3 space-y-0.5">
      <SidebarLink
        v-for="item in mainLinks"
        :key="item.to"
        :to="item.to"
        :icon="item.icon"
        :label="item.label"
        compact
        large-icon
      />
      </nav>

    <div
      v-if="collections.length"
      class="mt-3 border-t border-slate-200 pt-3 px-3 space-y-0.5"
    >
      <p class="text-xs font-semibold uppercase text-slate-500 mb-1">
        Collections
      </p>
      <SidebarLink
        v-for="col in collections"
        :key="col.id"
        :to="`/dashboard/collections/${col.id}`"
        :icon="CollectionIcon"
        :label="col.title"
        compact
      />
      </div>

    <div class="mt-auto border-t border-slate-200 p-3">
      <p class="text-xs font-semibold uppercase text-slate-500 mb-1">
        Sales Channels
      </p>
      <SidebarLink
        to="/dashboard/online-store"
        :icon="GlobeAltIcon"
        label="Online Store"
        compact
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
  CollectionIcon,
} from '@heroicons/vue/outline'

const shopStore = useShopStore()
const activeShop = computed(() => shopStore.active)
const collections = computed(() => activeShop.value?.collections || [])

// Primary navigation items
const mainLinks = [
  { to: '/dashboard', icon: HomeIcon, label: 'Home' },
  { to: '/dashboard/orders', icon: ClipboardListIcon, label: 'Orders' },
  { to: '/dashboard/products', icon: TagIcon, label: 'Products' },
  { to: '/dashboard/collections', icon: CollectionIcon, label: 'Collections' },
  { to: '/dashboard/customers', icon: UsersIcon, label: 'Customers' },
  { to: '/dashboard/discounts', icon: GiftIcon, label: 'Discounts' },
  { to: '/dashboard/analytics', icon: ChartBarIcon, label: 'Analytics' },
  { to: '/dashboard/settings', icon: CogIcon, label: 'Settings' },
]
</script>