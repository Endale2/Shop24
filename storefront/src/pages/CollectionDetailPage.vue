<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-green-50 text-gray-900 antialiased font-sans py-12 px-4 sm:px-6 lg:px-8">
    <div class="container mx-auto">
      <button @click="$router.back()" class="inline-flex items-center text-green-600 hover:text-green-800 font-semibold mb-8 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 rounded-md">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        Back to Collections
      </button>

      <!-- Loading state -->
      <div v-if="colStore.loading" class="flex items-center justify-center h-64 bg-white rounded-xl shadow-lg">
        <svg class="animate-spin h-14 w-14 text-green-600 mb-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </div>

      <!-- Collection Not Found State -->
      <div v-else-if="!colStore.detail" class="text-center text-red-600 bg-red-50 p-6 rounded-xl shadow-md">
        <p class="text-lg font-semibold">Collection not found.</p>
        <p class="text-base">The collection you are looking for does not exist or has been removed.</p>
      </div>

      <!-- Collection Details and Products -->
      <div v-else class="bg-white rounded-xl shadow-xl p-6 lg:p-10">
        <h2 class="text-4xl sm:text-5xl font-extrabold tracking-tight text-gray-900 mb-4 relative inline-block">
          {{ colStore.detail.title }}
          <span class="absolute left-1/2 transform -translate-x-1/2 bottom-0 w-24 h-1 bg-green-500 rounded-full"></span>
        </h2>
        <p class="text-lg text-gray-700 leading-relaxed mb-8">{{ colStore.detail.description || 'A detailed look into this exclusive collection.' }}</p>

        <!-- Product Display for Collection Products using ProductCard -->
        <div v-if="colStore.detail.products && colStore.detail.products.length > 0">
          <h3 class="text-2xl font-bold text-gray-800 mb-6">Products in this Collection</h3>
          <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 xl:gap-8">
            <router-link
              v-for="p in colStore.detail.products"
              :key="p.id"
              :to="`/products/${p.slug}`"
              class="block group"
            >
              <ProductCard :product="p" @add-to-cart="handleAddToCart" />
            </router-link>
          </div>
        </div>
        <div v-else class="text-center text-gray-600 p-8 rounded-lg bg-gray-50 border border-gray-200">
            <p class="text-lg font-medium">No products found in this collection yet.</p>
            <p class="text-sm">Please check back later or browse <router-link to="/products" class="text-green-600 hover:underline">all products</router-link>.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useCollectionStore } from '@/stores/collection'
import ProductCard from '@/components/ProductCard.vue' // Re-import ProductCard

const route = useRoute()
const handle = route.params.handle
const shopSlug = window.location.host.split('.')[0]
const colStore = useCollectionStore()

onMounted(() => {
  colStore.loadDetail(shopSlug, handle)
})

// Handler for the add-to-cart event from ProductCard
const handleAddToCart = (product) => {
  console.log('Collection Detail: Add to cart triggered for', product.name);
  // Implement your global cart logic here, e.g., dispatching to a Vuex/Pinia store
  // Example: useCartStore().addItem(product);
  alert(`Added "${product.name}" to cart!`); // Placeholder notification
};
</script>

<style scoped>
/* No additional custom styles needed as Tailwind handles everything. */
</style>
