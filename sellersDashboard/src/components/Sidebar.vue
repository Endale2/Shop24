<template>
  <aside class="w-64 bg-white border-r h-screen hidden md:flex flex-col">
    <!-- Logo / Header -->
    <div class="px-4 py-6 flex items-center space-x-2 border-b">
      <ShopIcon class="h-6 w-6 text-green-600" />
      <span class="text-xl font-semibold">My Store</span>
    </div>

    <!-- Main Navigation -->
    <nav class="flex-1 px-2 py-4 space-y-1 overflow-y-auto">
      <SidebarLink to="/app" icon="HomeIcon" label="Home" />
      <SidebarLink to="/app/orders" icon="ClipboardListIcon" label="Orders" />
      <SidebarLink to="/app/products" icon="TagIcon" label="Products" />
      <SidebarLink to="/app/customers" icon="UsersIcon" label="Customers" />
      <SidebarLink to="/app/discounts" icon="GiftIcon" label="Discounts" />
      <SidebarLink to="/app/analytics" icon="ChartBarIcon" label="Analytics" />
      <SidebarLink to="/app/settings" icon="CogIcon" label="Settings" />
    </nav>

    <!-- Sales Channels -->
    <div class="px-4 py-4 border-t">
      <p class="text-xs font-semibold text-gray-500 uppercase mb-2">Sales channels</p>
      <SidebarLink to="/app/online-store" icon="GlobeAltIcon" label="Online Store" :small="true" />
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

export default {
  name: 'Sidebar',
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
