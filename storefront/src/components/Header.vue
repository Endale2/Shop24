<template>
  <header class="bg-white border-b border-gray-200 relative">
    <div class="flex items-center justify-between max-w-7xl mx-auto px-4 sm:px-6 py-4">
      <router-link :to="`/shops/${shopSlug}`" class="flex items-center space-x-4">
        <img
          v-if="shop?.image"
          :src="shop.image"
          alt="Shop Logo"
          class="h-8 w-8 object-contain"
        />
        <span class="text-2xl font-light tracking-tight text-gray-900 uppercase">{{ shop?.name || 'Store' }}</span>
      </router-link>

      <div class="hidden md:flex items-center space-x-6">
        <div class="relative w-96 max-w-lg">
          <span class="absolute inset-y-0 left-0 flex items-center pl-3">
            <MagnifyingGlassIcon class="w-5 h-5 text-gray-400" />
          </span>
          <input
            type="text"
            placeholder="Search products..."
            class="w-full pl-10 pr-4 py-2 border border-gray-200 rounded-lg bg-gray-50 text-sm focus:outline-none focus:ring-2 focus:ring-black focus:border-black transition-all duration-300 ease-in-out"
          />
        </div>

        <nav class="flex items-center space-x-8 text-base font-light">
          <router-link
            :to="`/shops/${shopSlug}/products`"
            class="text-gray-600 hover:text-black transition"
            active-class="font-semibold text-black"
          >
            Products
          </router-link>
          <router-link
            :to="`/shops/${shopSlug}/collections`"
            class="text-gray-600 hover:text-black transition"
            active-class="font-semibold text-black"
          >
            Collections
          </router-link>
        </nav>

        <div class="flex items-center space-x-5">
          <router-link :to="`/shops/${shopSlug}/wishlist`" class="text-gray-600 hover:text-black transition" title="Wishlist">
            <HeartIcon class="w-6 h-6" />
          </router-link>
          <router-link :to="`/shops/${shopSlug}/cart`" class="text-gray-600 hover:text-black transition" title="Cart">
            <ShoppingBagIcon class="w-6 h-6" />
          </router-link>
          <router-link :to="`/shops/${shopSlug}/account`" class="text-gray-600 hover:text-black transition" title="Account">
            <UserIcon class="w-6 h-6" />
          </router-link>
        </div>
      </div>

      <div class="md:hidden">
        <button @click="isMobileSearchVisible = true" class="p-2 text-gray-600 hover:text-black focus:outline-none">
          <MagnifyingGlassIcon class="w-6 h-6" />
        </button>
      </div>
    </div>

    <div v-if="isMobileSearchVisible" class="absolute top-0 left-0 w-full h-full bg-white z-10 flex items-center px-4">
      <div class="relative w-full">
        <span class="absolute inset-y-0 left-0 flex items-center pl-3">
          <MagnifyingGlassIcon class="w-5 h-5 text-gray-400" />
        </span>
        <input
          type="text"
          placeholder="Search for products..."
          class="w-full pl-10 pr-10 py-2 border-b-2 border-gray-300 bg-transparent text-sm focus:outline-none focus:border-black"
          autofocus
        />
        <button @click="isMobileSearchVisible = false" class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-black">
          <XMarkIcon class="w-6 h-6" />
        </button>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import {
  MagnifyingGlassIcon,
  HeartIcon,
  ShoppingBagIcon,
  UserIcon,
  XMarkIcon
} from '@heroicons/vue/24/outline';

interface Props {
  shop: { name: string; image?: string } | null;
}
defineProps<Props>();

const route = useRoute();
const shopSlug = route.params.shopSlug as string;

const isMobileSearchVisible = ref(false);
</script>