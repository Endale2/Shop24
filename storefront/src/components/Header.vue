<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
    <div class="flex items-center justify-between max-w-7xl mx-auto px-4 sm:px-6 py-4">
      <router-link :to="{ path: `/${shopSlug}/` }" class="flex items-center space-x-4" aria-label="Home">
        <img
          v-if="shop?.image"
          :src="shop.image"
          alt="Shop Logo"
          class="h-8 w-8 object-contain"
        />
        <span class="text-2xl font-light tracking-tight text-gray-900 uppercase">{{ shop?.name || 'Store' }}</span>
      </router-link>

      <!-- Desktop Nav -->
      <div class="hidden md:flex flex-1 items-center justify-end space-x-2 lg:space-x-6">
        <div class="w-full max-w-xs lg:max-w-sm">
          <div class="relative">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <MagnifyingGlassIcon class="w-5 h-5 text-gray-400" />
            </span>
            <input
              type="text"
              placeholder="Search products..."
              class="w-full pl-10 pr-4 py-2 border border-gray-200 rounded-lg bg-gray-50 text-sm focus:outline-none focus:ring-2 focus:ring-black focus:border-black transition-all duration-300 ease-in-out"
              v-model="searchInput"
              @input="onSearchInput"
            />
          </div>
        </div>

        <nav class="flex items-center space-x-4 lg:space-x-6 text-base font-light">
          <router-link
            :to="{ path: `/${shopSlug}/products` }"
            class="text-gray-600 hover:text-black transition flex items-center"
            :class="{ 'font-semibold text-black': isActive('/products') }"
            aria-label="Products"
          >
            <span>Products</span>
          </router-link>
          <router-link
            :to="{ path: `/${shopSlug}/collections` }"
            class="text-gray-600 hover:text-black transition flex items-center"
            :class="{ 'font-semibold text-black': isActive('/collections') }"
            aria-label="Collections"
          >
            <span>Collections</span>
          </router-link>
        </nav>

        <div class="flex items-center space-x-3 lg:space-x-5">
          <!-- Wishlist -->
          <router-link :to="{ path: `/${shopSlug}/wishlist` }" class="relative group" title="Wishlist" aria-label="Wishlist">
            <component :is="isActive('/wishlist') ? HeartIconSolid : HeartIcon" class="w-6 h-6 transition-colors" :class="isActive('/wishlist') ? 'text-red-500' : 'text-gray-600 group-hover:text-red-500'" />
            <span 
              v-if="authStore.user && wishlistCount > 0"
              class="absolute -top-2 -right-2 bg-red-600 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-medium"
            >
              {{ wishlistCount > 99 ? '99+' : wishlistCount }}
            </span>
          </router-link>
          <!-- Cart -->
          <router-link :to="{ path: `/${shopSlug}/cart` }" class="relative group" title="Cart" aria-label="Cart">
            <component :is="isActive('/cart') ? ShoppingBagIconSolid : ShoppingBagIcon" class="w-6 h-6 transition-colors" :class="isActive('/cart') ? 'text-black' : 'text-gray-600 group-hover:text-black'" />
            <span 
              v-if="cartItemCount > 0"
              class="absolute -top-2 -right-2 bg-black text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-medium"
            >
              {{ cartItemCount > 99 ? '99+' : cartItemCount }}
            </span>
          </router-link>
          <!-- Account/Login -->
          <div v-if="user" class="relative" v-click-outside="closeDropdown">
            <button @click="toggleDropdown" class="focus:outline-none" aria-label="Account menu">
              <img v-if="user.profilePic" :src="user.profilePic" alt="Profile" class="w-8 h-8 rounded-full object-cover border-2 border-transparent group-hover:border-gray-300 transition" />
              <span v-else class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center text-lg font-bold text-gray-700 group-hover:bg-gray-300 transition">
                {{ getAvatarText() }}
              </span>
            </button>
            <transition name="fade-down">
              <div 
                v-if="dropdownOpen" 
                class="absolute right-0 mt-2 w-56 bg-white border border-gray-100 rounded-lg shadow-xl z-50 overflow-hidden"
              >
                <div class="px-4 py-3 bg-gray-50 border-b border-gray-100">
                  <p class="text-sm font-semibold text-gray-800 truncate">{{ user.firstName || user.username }}</p>
                  <p class="text-xs text-gray-500 truncate">{{ user.email }}</p>
                </div>
                <div class="py-1">
                  <router-link :to="{ path: `/${shopSlug}/account` }" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors">My Account</router-link>
                  <router-link :to="{ path: `/${shopSlug}/orders` }" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors">My Orders</router-link>
                  <router-link :to="{ path: `/${shopSlug}/wishlist` }" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors">Wishlist</router-link>
                </div>
                <div class="border-t border-gray-100 py-1">
                  <button @click="logout" class="block w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors">Logout</button>
                </div>
              </div>
            </transition>
          </div>
          <router-link v-else :to="{ path: `/${shopSlug}/login` }" class="group" title="Login" aria-label="Login">
            <UserIcon class="w-6 h-6 text-gray-600 group-hover:text-black transition-colors" />
          </router-link>
        </div>
      </div>

      <!-- Mobile Nav -->
      <div class="md:hidden flex items-center space-x-4">
        <button @click="isMobileMenuOpen = !isMobileMenuOpen" class="p-2 text-gray-600 hover:text-black focus:outline-none" aria-label="Open menu">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16"/></svg>
        </button>
        <button @click="isMobileSearchVisible = !isMobileSearchVisible" class="p-2 text-gray-600 hover:text-black focus:outline-none" aria-label="Open search">
          <MagnifyingGlassIcon class="w-6 h-6" />
        </button>
      </div>
    </div>

    <!-- Mobile Dropdown/Drawer -->
    <transition name="fade">
      <div v-if="isMobileMenuOpen" class="fixed inset-0 z-50 flex justify-end" style="background: transparent;">
        <div class="w-64 bg-white h-full shadow-lg p-6 flex flex-col space-y-4">
          <button @click="isMobileMenuOpen = false" class="self-end mb-4 text-gray-400 hover:text-black" aria-label="Close menu">
            <XMarkIcon class="w-6 h-6" />
          </button>
          <router-link :to="{ path: `/${shopSlug}/products` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Products</router-link>
          <router-link :to="{ path: `/${shopSlug}/collections` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Collections</router-link>
          <router-link :to="{ path: `/${shopSlug}/wishlist` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Wishlist</router-link>
          <router-link :to="{ path: `/${shopSlug}/cart` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Cart</router-link>
          <router-link v-if="user" :to="{ path: `/${shopSlug}/account` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">My Account</router-link>
          <router-link v-if="user" :to="{ path: `/${shopSlug}/orders` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">My Orders</router-link>
          <button v-if="user" @click="logout" class="block w-full text-left px-4 py-2 text-red-600 hover:bg-red-50 rounded">Logout</button>
          <router-link v-else :to="{ path: `/${shopSlug}/login` }" class="block px-4 py-2 text-gray-700 hover:bg-gray-50 rounded">Login</router-link>
        </div>
      </div>
    </transition>

    <!-- Mobile Search Overlay -->
    <transition name="fade">
      <div v-if="isMobileSearchVisible" class="fixed inset-0 z-50 bg-white flex items-center px-4">
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
          <button @click="isMobileSearchVisible = false" class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-black" aria-label="Close search">
            <XMarkIcon class="w-6 h-6" />
          </button>
        </div>
      </div>
    </transition>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';
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
  XMarkIcon
} from '@heroicons/vue/24/outline';

const vClickOutside = {
  mounted(el: any, binding: any) {
    el.__ClickOutsideHandler__ = (event: MouseEvent) => {
      if (!(el === event.target || el.contains(event.target))) {
        binding.value(event);
      }
    };
    document.body.addEventListener('click', el.__ClickOutsideHandler__);
  },
  unmounted(el: any) {
    document.body.removeEventListener('click', el.__ClickOutsideHandler__);
  },
};

interface Props {
  shop: { name: string; image?: string } | null;
}
defineProps<Props>();

const router = useRouter();
const shopSlug = getCurrentShopSlug();
const isMobileSearchVisible = ref(false);
const isMobileMenuOpen = ref(false);
const dropdownOpen = ref(false);
const searchInput = ref('');
let searchTimeout: ReturnType<typeof setTimeout> | null = null;

function onSearchInput() {
  if (searchTimeout) clearTimeout(searchTimeout);
  searchTimeout = setTimeout(() => {
    if (searchInput.value.trim()) {
      router.push({ path: `/${shopSlug}/search`, query: { q: searchInput.value.trim() } });
    }
  }, 400);
}

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

// Ensure wishlist is fetched after login and on navigation
onMounted(() => {
  if (authStore.user && shopSlug) {
    wishlistStore.setShopSlug(shopSlug);
    wishlistStore.fetchWishlist();
  }
});
watch(() => authStore.user, (user) => {
  if (user && shopSlug) {
    wishlistStore.setShopSlug(shopSlug);
    wishlistStore.fetchWishlist();
  }
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

function closeDropdown() {
  dropdownOpen.value = false;
}

async function logout() {
  await authStore.logout();
  closeDropdown();
  if (shopSlug) {
    router.push({ path: `/${shopSlug}/login` });
  }
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
.fade-down-enter-active, .fade-down-leave-active {
  transition: all 0.2s ease-out;
}
.fade-down-enter-from, .fade-down-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>