<template>
  <!-- Breadcrumb and Back Button -->
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
      Back
    </button>
    <router-link to="/" class="hover:underline">Home</router-link>
    <span>/</span>
    <span>Wishlist</span>
  </nav>
  <div class="wishlist-container">
    <div class="max-w-6xl mx-auto">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Wishlist</h1>
        <p class="text-gray-600">Save your favorite products for later</p>
      </div>

      <!-- Not Logged In State -->
      <div v-if="!authStore.user" class="text-center py-12">
        <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
        </svg>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Sign in to view your wishlist</h3>
        <p class="text-gray-600 mb-6">Create an account or sign in to save and track your favorite products</p>
        <div class="space-x-4">
          <router-link 
            :to="'/login'" 
            class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors"
          >
            Sign In
          </router-link>
          <router-link 
            :to="'/register'" 
            class="inline-block bg-white text-black border border-black py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-50 transition-colors"
          >
            Create Account
          </router-link>
        </div>
      </div>

      <!-- Logged In State -->
      <div v-else>
        <!-- Loading/Error State -->
        <div v-if="wishlistStore.loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-black"></div>
          <p class="mt-4 text-gray-600">Loading your wishlist...</p>
        </div>
        <div v-else-if="wishlistStore.error" class="text-center py-12">
          <svg class="w-16 h-16 text-red-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
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
                  <button class="w-full bg-black text-white py-2 px-3 text-sm font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors">
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
        <div v-else class="text-center py-12">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
          </svg>
          <h3 class="text-lg font-semibold text-gray-900 mb-2">Your wishlist is empty</h3>
          <p class="text-gray-600 mb-6">Start adding products to your wishlist</p>
          <router-link 
            :to="'/products'" 
            class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors"
          >
            Browse Products
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { useWishlistStore } from '../stores/wishlist';

const route = useRoute();
const authStore = useAuthStore();
const wishlistStore = useWishlistStore();

onMounted(() => {
  wishlistStore.setShopSlug(route.params.shopSlug as string);
  if (authStore.user) {
    wishlistStore.fetchWishlist();
  }
});

function removeFromWishlist(productId: string) {
  wishlistStore.removeProduct(productId);
}
</script>

<style scoped>
.wishlist-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 