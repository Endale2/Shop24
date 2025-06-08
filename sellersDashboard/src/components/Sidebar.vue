<template>
  <aside class="flex h-full w-64 flex-col overflow-y-auto border-r border-slate-200 bg-white">
    <div class="group relative h-20 w-full md:h-32">
      <div
        class="absolute inset-0 h-full w-full bg-slate-100 transition-all duration-300"
        :class="{
          'bg-cover bg-center': activeShop?.image,
          'group-hover:scale-105': activeShop?.image,
        }"
        :style="activeShop?.image ? `background-image: url(${activeShop.image})` : ''"
      >
        <div v-if="!activeShop?.image" class="flex h-full w-full items-center justify-center">
          <ShopIcon class="h-10 w-10 text-slate-400" />
        </div>
      </div>
    </div>

    <nav class="flex-1 space-y-1 p-2">
      <SidebarLink to="/dashboard" :icon="HomeIcon" label="Home" />
      <SidebarLink to="/dashboard/orders" :icon="ClipboardListIcon" label="Orders" />
      <SidebarLink to="/dashboard/products" :icon="TagIcon" label="Products" />
      <SidebarLink to="/dashboard/collections" :icon="CollectionIcon" label="Collections" />
      <SidebarLink to="/dashboard/customers" :icon="UsersIcon" label="Customers" />
      <SidebarLink to="/dashboard/discounts" :icon="GiftIcon" label="Discounts" />
      <SidebarLink to="/dashboard/analytics" :icon="ChartBarIcon" label="Analytics" />
      <SidebarLink to="/dashboard/settings" :icon="CogIcon" label="Settings" />

      <div v-if="collections.length" class="pt-4">
        <p class="mb-2 px-3 text-xs font-semibold uppercase text-slate-500">
          Your Collections
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
    </nav>

    <div class="border-t border-slate-200 p-2">
      <p class="mb-2 px-3 text-xs font-semibold uppercase text-slate-500">
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

const shopStore = useShopStore()
const activeShop = computed(() => shopStore.active)
const collections = computed(() => activeShop.value?.collections || [])
</script>