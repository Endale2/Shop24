<template>
  <aside
    class="flex h-full w-64 flex-col bg-white shadow-lg border-r border-slate-200"
  >
    <!-- Logo Section -->
    <div
      class="flex h-20 items-center justify-center border-b border-slate-200 p-2"
    >
      <img
        v-if="activeShop?.image"
        :src="activeShop.image"
        alt="Shop Logo"
        class="h-12 w-12 rounded-md object-cover"
      />
      <ShopIcon
        v-else
        class="h-12 w-12 text-slate-400"
      />
    </div>

    <!-- Main Navigation -->
    <nav class="flex-1 overflow-y-auto py-4 space-y-1">
      <SidebarLink
        v-for="item in mainLinks"
        :key="item.to"
        :to="item.to"
        :icon="item.icon"
        :label="item.label"
      />
    </nav>

    <!-- Collections Group -->
    <div
      v-if="collections.length"
      class="mt-4 border-t border-slate-200 pt-4 px-3 space-y-1"
    >
      <p class="text-xs font-semibold uppercase text-slate-500 mb-2">
        Collections
      </p>
      <SidebarLink
        v-for="col in collections"
        :key="col.id"
        :to="`/dashboard/collections/${col.id}`"
        :icon="CollectionIcon"
        :label="col.title"
        small
      />
    </div>

    <!-- Sales Channels -->
    <div class="mt-auto border-t border-slate-200 p-4">
      <p class="text-xs font-semibold uppercase text-slate-500 mb-2">
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

const shopStore   = useShopStore()
const activeShop  = computed(() => shopStore.active)
const collections = computed(() => activeShop.value?.collections || [])

// Primary navigation items
const mainLinks = [
  { to: '/dashboard',          icon: HomeIcon,          label: 'Home'       },
  { to: '/dashboard/orders',    icon: ClipboardListIcon, label: 'Orders'     },
  { to: '/dashboard/products',  icon: TagIcon,           label: 'Products'   },
  { to: '/dashboard/collections', icon: CollectionIcon,   label: 'Collections'},
  { to: '/dashboard/customers', icon: UsersIcon,         label: 'Customers'  },
  { to: '/dashboard/discounts', icon: GiftIcon,          label: 'Discounts'  },
  { to: '/dashboard/analytics', icon: ChartBarIcon,      label: 'Analytics'  },
  { to: '/dashboard/settings',  icon: CogIcon,           label: 'Settings'   },
]
</script>
