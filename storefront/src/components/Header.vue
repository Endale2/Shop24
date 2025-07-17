<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
    <div class="flex items-center justify-between max-w-7xl mx-auto px-4 sm:px-6 py-4">
      <router-link to="/" class="flex items-center space-x-4">
        <img
          v-if="shop?.image"
          :src="shop.image"
          alt="Shop Logo"
          class="h-8 w-8 object-contain"
        />
        <span class="text-2xl font-light tracking-tight text-gray-900 uppercase">{{ shop?.name || 'Store' }}</span>
      </router-link>

      <!-- Desktop Nav -->
      <div class="hidden md:flex items-center space-x-6">
        <div class="relative w-72 max-w-lg">
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
            to="/products"
            class="text-gray-600 hover:text-black transition flex items-center"
            :class="{ 'font-semibold text-black': isActive('/products') }"
          >
            <span>Products</span>
          </router-link>
          <router-link
            to="/collections"
            class="text-gray-600 hover:text-black transition flex items-center"
            :class="{ 'font-semibold text-black': isActive('/collections') }"
          >
            <span>Collections</span>
          </router-link>
        </nav>

        <div class="flex items-center space-x-5">
          <!-- Wishlist -->
          <router-link to="/wishlist" class="relative group" title="Wishlist">
            <component :is="isActive('/wishlist') ? HeartIconSolid : HeartIcon" class="w-6 h-6 transition-colors" :class="isActive('/wishlist') ? 'text-red-500' : 'text-gray-600 group-hover:text-red-500'" />
            <span 
              v-if="authStore.user && wishlistCount > 0"
              class="absolute -top-2 -right-2 bg-red-600 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-medium"
            >
              {{ wishlistCount > 99 ? '99+' : wishlistCount }}
            </span>
          </router-link>
          <!-- Cart -->
          <router-link to="/cart" class="relative group" title="Cart">
            <component :is="isActive('/cart') ? ShoppingBagIconSolid : ShoppingBagIcon" class="w-6 h-6 transition-colors" :class="isActive('/cart') ? 'text-black' : 'text-gray-600 group-hover:text-black'" />
            <span 
              v-if="cartItemCount > 0"
              class="absolute -top-2 -right-2 bg-black text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-medium"
            >
              {{ cartItemCount > 99 ? '99+' : cartItemCount }}
            </span>
          </router-link>
          <!-- Account/Login -->
          <div v-if="user" class="relative group">
            <button @click="toggleDropdown" class="focus:outline-none">
              <img v-if="user.profilePic" :src="user.profilePic" alt="Profile" class="w-8 h-8 rounded-full object-cover border border-gray-200" />
              <span v-else class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center text-lg font-bold text-gray-700">
                {{ getAvatarText() }}
              </span>
            </button>
            <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-lg shadow-lg z-50">
              <router-link to="/account" class="block px-4 py-2 text-gray-700 hover:bg-gray-50">My Account</router-link>
              <router-link to="/orders" class="block px-4 py-2 text-gray-700 hover:bg-gray-50">My Orders</router-link>
              <router-link to="/wishlist" class="block px-4 py-2 text-gray-700 hover:bg-gray-50">Wishlist</router-link>
              <button @click="logout" class="block w-full text-left px-4 py-2 text-red-600 hover:bg-red-50">Logout</button>
            </div>
          </div>
          <router-link v-else to="/login" class="group" title="Login">
            <UserIcon class="w-6 h-6 text-gray-600 group-hover:text-black transition-colors" />
          </router-link>
        </div>
      </div>

      <!-- Mobile Nav -->
      <div class="md:hidden flex items-center space-x-4">
        <router-link to="/cart" class="relative group" title="Cart">
          <component :is="isActive('/cart') ? ShoppingBagIconSolid : ShoppingBagIcon" class="w-6 h-6 transition-colors" :class="isActive('/cart') ? 'text-black' : 'text-gray-600 group-hover:text-black'" />
          <span 
            v-if="cartItemCount > 0"
            class="absolute -top-2 -right-2 bg-black text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-medium"
          >
            {{ cartItemCount > 99 ? '99+' : cartItemCount }}
          </span>
        </router-link>
        <button @click="isMobileMenuOpen = !isMobileMenuOpen" class="p-2 text-gray-600 hover:text-black focus:outline-none">
          <span v-if="user">
            <img v-if="user.profilePic" :src="user.profilePic" alt="Profile" class="w-8 h-8 rounded-full object-cover border border-gray-200" />
            <span v-else class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center text-lg font-bold text-gray-700">
              {{ getAvatarText() }}
            </span>
          </span>
          <span v-else>
            <UserIcon class="w-6 h-6" />
          </span>
        </button>
      </div>
    </div>

    <!-- Mobile Dropdown/Drawer -->
    <transition name="fade">
      <div v-if="isMobileMenuOpen" class="fixed inset-0 z-50 bg-black bg-opacity-30 flex justify-end">
        <div class="w-64 bg-white h-full shadow-lg p-6 flex flex-col space-y-4">
          <button @click="isMobileMenuOpen = false" class="self-end mb-4 text-gray-400 hover:text-black">
            <XMarkIcon class="w-6 h-6" />
          </button>
          <div v-if="user" class="mb-4 flex flex-col items-center">
            <img v-if="user.profilePic" :src="user.profilePic" alt="Profile" class="w-16 h-16 rounded-full object-cover border border-gray-200 mb-2" />
            <span v-else class="w-16 h-16 rounded-full bg-gray-200 flex items-center justify-center text-2xl font-bold text-gray-700 mb-2">
              {{ getAvatarText() }}
            </span>
            <span class="text-lg font-semibold text-gray-900">{{ user.firstName || user.username || user.email }}</span>
          </div>
          <router-link to="/account" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">My Account</router-link>
          <router-link to="/orders" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">My Orders</router-link>
          <router-link to="/wishlist" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Wishlist</router-link>
          <button v-if="user" @click="logout" class="block w-full text-left px-4 py-2 text-red-600 hover:bg-red-50 rounded">Logout</button>
          <router-link v-else to="/login" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Login</router-link>
        </div>
      </div>
    </transition>

    <!-- Mobile Search Overlay -->
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
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { useCartStore } from '../stores/cart';
import { useWishlistStore } from '../stores/wishlist';
import { getCurrentShopSlug } from '../services/shop';
import {
  MagnifyingGlassIcon,
  HeartIcon,
  HeartIcon as HeartIconSolid,
  ShoppingBagIcon,
  ShoppingBagIcon as ShoppingBagIconSolid,
  UserIcon,
  UserIcon as UserIconSolid,
  XMarkIcon
} from '@heroicons/vue/24/outline';

interface Props {
  shop: { name: string; image?: string } | null;
}
defineProps<Props>();

const router = useRouter();
const shopSlug = getCurrentShopSlug();
const isMobileSearchVisible = ref(false);
const isMobileMenuOpen = ref(false);
const dropdownOpen = ref(false);

const authStore = useAuthStore();
const cartStore = useCartStore();
const wishlistStore = useWishlistStore();

const user = computed(() => authStore.user);

const cartItemCount = computed(() => {
  if (!cartStore.cart || !cartStore.cart.items) return 0;
  return cartStore.cart.items.reduce((total: number, item: any) => total + item.quantity, 0);
});

const wishlistCount = computed(() => {
  if (!authStore.user) return 0;
  return wishlistStore.productIds.length;
});

function isActive(path: string) {
  return router.currentRoute.value.path.includes(path);
}

function getAvatarText() {
  if (!user.value) return '';
  if (user.value.firstName) return user.value.firstName[0].toUpperCase();
  if (user.value.email) return user.value.email[0].toUpperCase();
  return user.value.username ? user.value.username[0].toUpperCase() : '';
}

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value;
}

async function logout() {
  await (authStore as any).logout();
  router.push(`/shops/${shopSlug}/login`);
}
</script>

<style scoped>
header {
  transition: box-shadow 0.2s;
}
@media (max-width: 768px) {
  .max-w-7xl {
    padding-left: 0.5rem;
    padding-right: 0.5rem;
  }
}
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>