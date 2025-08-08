<template>
  <Breadcrumbs :items="[
    { back: true, label: 'Back' },
    { label: 'Home', to: `/` },
    { label: 'Wishlist' }
  ]" />
  <div class="wishlist-container">
    <div class="max-w-6xl mx-auto">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Wishlist</h1>
        <p class="text-gray-600">Save your favorite products for later</p>
      </div>

      <!-- Session Loading State -->
      <div v-if="authStore.sessionLoading" class="flex items-center justify-center min-h-[60vh]">
        <div class="flex flex-col items-center">
          <svg class="animate-spin h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <p class="text-gray-500">Checking authentication…</p>
        </div>
      </div>

      <!-- Not Logged In State -->
      <LoginPrompt
        v-else-if="!authStore.user"
        title="Sign in to view your wishlist"
        message="Create an account or sign in to save and track your favorite products."
      >
        <template #icon>
          <svg class="w-24 h-24 text-gray-300 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
          </svg>
        </template>
      </LoginPrompt>

      <!-- Logged In State -->
      <div v-else>
        <!-- Loading/Error State -->
        <div v-if="wishlistStore.loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-black"></div>
          <p class="mt-4 text-gray-600">Loading your wishlist...</p>
        </div>
        <div v-else-if="wishlistStore.error && wishlistStore.error !== 'mongo: no documents in result'" class="text-center py-12">
          <svg class="w-16 h-16 text-red-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">Unable to load wishlist</h3>
          <p class="text-gray-600 mb-4">{{ wishlistStore.error }}</p>
          <button @click="wishlistStore.fetchWishlist" class="bg-black text-white py-2 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors">
            Try Again
          </button>
        </div>
        <!-- Wishlist Items -->
        <div v-else-if="wishlistStore.products.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="item in wishlistStore.products" :key="item.id" class="bg-white border border-gray-200 rounded-none p-4">
            <div class="flex items-start space-x-4">
              <img :src="item.main_image || item.image || '/placeholder-product.jpg'" :alt="item.name" class="w-20 h-20 object-cover rounded-lg" />
              <div class="flex-1 min-w-0">
                <h3 class="text-sm font-medium text-gray-900 truncate">{{ item.name }}</h3>
                <p class="text-sm font-semibold text-gray-900 mt-1">${{ item.price }}</p>
                <div class="mt-3 space-y-2">
                  <button @click="addToCart(item)" class="w-full bg-black text-white py-2 px-3 text-sm font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors">
                    Add to Cart
                  </button>
                  <button @click="removeFromWishlist(item.id)" class="w-full bg-white text-red-600 border border-red-600 py-2 px-3 text-sm font-semibold uppercase tracking-wide hover:bg-red-50 transition-colors">
                    Remove
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Empty Wishlist State -->
        <div v-else class="text-center py-16 flex flex-col items-center justify-center">
          <svg class="w-24 h-24 text-gray-300 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/>
          </svg>
          <h3 class="text-2xl font-bold text-gray-900 mb-2">Your wishlist is empty</h3>
          <p class="text-gray-500 mb-6 max-w-md">You haven't added any products to your wishlist yet. Start exploring and save your favorite items for later!</p>
          <button @click="goToProducts" class="inline-block bg-black text-white py-3 px-8 rounded-md font-bold uppercase tracking-wide hover:bg-gray-800 transition-colors text-lg shadow-md">
            Browse Products
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getCurrentShopSlug } from '@/services/shop';
import { useAuthStore } from '../stores/auth';
import { useWishlistStore } from '../stores/wishlist';
import { useCartStore } from '../stores/cart';
import LoginPrompt from '@/components/LoginPrompt.vue';
import Breadcrumbs from '@/components/Breadcrumbs.vue';

const router = useRouter();
const authStore = useAuthStore();
const wishlistStore = useWishlistStore();
const cartStore = useCartStore();
const shopSlug = getCurrentShopSlug() as string;

function goToProducts() {
  router.push({ path: `/products` });
}

function removeFromWishlist(productId: string) {
  wishlistStore.removeProduct(productId);
}

async function addToCart(product: any) {
  await cartStore.addToCart(product, '', 1);
}

// Only fetch wishlist when user is authenticated and session loading is complete
function fetchWishlistIfAuthenticated() {
  if (!shopSlug) return;
  if (!authStore.sessionLoading && authStore.user) {
    wishlistStore.setShopSlug(shopSlug);
    wishlistStore.fetchWishlist();
  }
}

onMounted(() => {
  fetchWishlistIfAuthenticated();
});

// Watch for authentication state changes
watch(
  () => [authStore.sessionLoading, authStore.user],
  ([sessionLoading, user]) => {
    if (!sessionLoading && user) {
      fetchWishlistIfAuthenticated();
    }
  }
);

watch(() => getCurrentShopSlug(), (newSlug) => {
  if (newSlug) {
    fetchWishlistIfAuthenticated();
  }
});
</script>

<style scoped>
.wishlist-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 