<template>
  <aside class="w-64 bg-white border-r border-gray-200 h-screen hidden md:flex flex-col shadow-sm">
    <div class="px-4 py-6 flex items-center space-x-2 border-b border-gray-200">
      <ShopIcon class="h-7 w-7 text-green-600" />
      <span class="text-2xl font-bold text-gray-800">{{ activeShop.name }}</span>
    </div>

    <nav class="flex-1 px-2 py-4 space-y-1 overflow-y-auto">
      <SidebarLink to="/dashboard" :icon="HomeIcon" label="Home" />
      <SidebarLink to="/dashboard/orders" :icon="ClipboardListIcon" label="Orders" />
      <SidebarLink to="/dashboard/products" :icon="TagIcon" label="Products" />
      <SidebarLink to="/dashboard/customers" :icon="UsersIcon" label="Customers" />
      <SidebarLink to="/dashboard/discounts" :icon="GiftIcon" label="Discounts" />
      <SidebarLink to="/dashboard/analytics" :icon="ChartBarIcon" label="Analytics" />
      <SidebarLink to="/dashboard/settings" :icon="CogIcon" label="Settings" />
    </nav>

    <div class="px-4 py-4 border-t border-gray-200">
      <p class="text-xs font-semibold text-gray-500 uppercase mb-2">Sales channels</p>
      <SidebarLink to="/dashboard/online-store" :icon="GlobeAltIcon" label="Online Store" :small="true" />
    </div>
  </aside>
</template>

<script>
import SidebarLink from './SidebarLink.vue';
import {
  HomeIcon,
  ClipboardListIcon,
  TagIcon,
  UsersIcon,
  GiftIcon,
  ChartBarIcon,
  CogIcon,
  GlobeAltIcon,
  ShoppingBagIcon as ShopIcon
} from '@heroicons/vue/outline';
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'Sidebar',
  computed: {
    ...mapGetters('auth', ['user']),
    ...mapGetters('shops', ['activeShop'])
  },
  components: {
    SidebarLink,
    HomeIcon,
    ClipboardListIcon,
    TagIcon,
    UsersIcon,
    GiftIcon,
    ChartBarIcon,
    CogIcon,
    GlobeAltIcon,
    ShopIcon
  }
};
</script>