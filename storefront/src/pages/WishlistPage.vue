<template>
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
            :to="`/shops/${route.params.shopSlug}/login`" 
            class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors"
          >
            Sign In
          </router-link>
          <router-link 
            :to="`/shops/${route.params.shopSlug}/register`" 
            class="inline-block bg-white text-black border border-black py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-50 transition-colors"
          >
            Create Account
          </router-link>
        </div>
      </div>

      <!-- Logged In State -->
      <div v-else>
        <!-- Wishlist Items -->
        <div v-if="wishlist.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="item in wishlist" :key="item.id" class="bg-white border border-gray-200 rounded-none p-4">
            <div class="flex items-start space-x-4">
              <img :src="item.image" :alt="item.name" class="w-20 h-20 object-cover rounded-lg" />
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
            :to="`/shops/${route.params.shopSlug}/products`" 
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
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const route = useRoute();
const authStore = useAuthStore();

const wishlist = ref([
  { id: 1, name: 'Random Product 1', price: 19.99, image: 'https://via.placeholder.com/80' },
  { id: 2, name: 'Random Product 2', price: 29.99, image: 'https://via.placeholder.com/80' },
  { id: 3, name: 'Random Product 3', price: 39.99, image: 'https://via.placeholder.com/80' },
]);

function removeFromWishlist(id: number) {
  wishlist.value = wishlist.value.filter(item => item.id !== id);
}
</script>

<style scoped>
.wishlist-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 