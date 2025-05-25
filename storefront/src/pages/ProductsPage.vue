<template>
  <div v-if="shop.isLoadingShop && firstLoad" class="flex min-h-[calc(100vh-200px)] items-center justify-center">
    <p class="text-xl text-gray-500 dark:text-gray-400">Loading shop information...</p>
    </div>

  <div v-else-if="shop.shopError && !shop.current" class="flex min-h-[calc(100vh-200px)] flex-col items-center justify-center text-center">
    <p class="text-xl text-red-500 dark:text-red-400">Error loading shop data.</p>
    <p class="text-gray-600 dark:text-gray-300 mt-2">{{ shop.shopError }}</p>
    <router-link
      to="/"
      class="mt-6 px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700 transition-colors"
    >
      Try Homepage
    </router-link>
  </div>

  <div v-else-if="shop.current">
    <section class="container mx-auto px-4 py-8">
      <div class="flex flex-col sm:flex-row justify-between items-center mb-8 pb-4 border-b border-gray-200 dark:border-gray-700">
        <h1 class="text-3xl md:text-4xl font-bold text-gray-800 dark:text-white mb-4 sm:mb-0">
          All Products
        </h1>
        <div class="flex items-center space-x-2">
          <label for="sort-by" class="text-sm text-gray-600 dark:text-gray-400">Sort by:</label>
          <select 
            id="sort-by" 
            v-model="sortBy" 
            class="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm"
          >
            <option value="default">Default</option>
            <option value="name_asc">Name (A-Z)</option>
            <option value="name_desc">Name (Z-A)</option>
            <option value="price_asc">Price (Low to High)</option>
            <option value="price_desc">Price (High to Low)</option>
            <option value="date_desc">Newest</option>
          </select>
        </div>
      </div>
      
      <div v-if="shop.isLoadingProducts" class="text-center py-20">
        <p class="text-lg text-gray-500 dark:text-gray-400">Loading products...</p>
        </div>
      <div v-else-if="shop.productsError" class="text-center py-20">
        <p class="text-lg text-red-500 dark:text-red-400">Error loading products.</p>
        <p class="text-gray-600 dark:text-gray-300 mt-1">{{ shop.productsError }}</p>
      </div>
      <ProductGrid
        v-else-if="sortedProducts && sortedProducts.length > 0"
        :products="sortedProducts" 
        :shopSlug="shop.current.slug" 
      />
      <div v-else class="text-center py-20">
         <p class="text-lg text-gray-500 dark:text-gray-400">No products found for this shop.</p>
         <p v-if="sortBy !== 'default'" class="text-sm text-gray-400 dark:text-gray-500 mt-1">Try a different sort option or check back later.</p>
      </div>
    </section>
  </div>
  
  <div v-else-if="!firstLoad" class="flex min-h-[calc(100vh-200px)] items-center justify-center">
     <p class="text-xl text-gray-500 dark:text-gray-400">Shop information is currently unavailable.</p>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useShopStore } from '@/stores/shop'
import ProductGrid from '@/components/ProductGrid.vue'
import { RouterLink } from 'vue-router'

const shop = useShopStore()
const sortBy = ref('default') // Default sort option
const firstLoad = ref(true) // To manage initial display before store might be populated

// Computed property for sorted products
const sortedProducts = computed(() => {
  if (!shop.products) return [];
  
  // Create a shallow copy to avoid mutating the store's state directly
  let productsCopy = [...shop.products];

  switch (sortBy.value) {
    case 'name_asc':
      return productsCopy.sort((a, b) => (a.name || '').localeCompare(b.name || ''));
    case 'name_desc':
      return productsCopy.sort((a, b) => (b.name || '').localeCompare(a.name || ''));
    case 'price_asc':
      // Ensure price is a number for sorting, default to Infinity for undefined/null prices to push them to the end
      return productsCopy.sort((a, b) => (parseFloat(a.price) || Infinity) - (parseFloat(b.price) || Infinity));
    case 'price_desc':
      // Default to -Infinity for undefined/null prices to push them to the end when sorting descending
      return productsCopy.sort((a, b) => (parseFloat(b.price) || -Infinity) - (parseFloat(a.price) || -Infinity));
    case 'date_desc': // Assuming products have a 'createdAt' or 'updatedAt' ISO date string
      return productsCopy.sort((a, b) => new Date(b.updatedAt || b.createdAt || 0).getTime() - new Date(a.updatedAt || a.createdAt || 0).getTime());
    case 'default':
    default:
      return productsCopy; // Return original order or server-defined order
  }
});

onMounted(() => {
  // This helps manage the "firstLoad" state. 
  // DefaultLayout is primarily responsible for triggering fetches.
  // This flag helps avoid briefly showing "Shop information unavailable" if the store
  // takes a moment to update after DefaultLayout's onBeforeMount.
  if (shop.isLoadingShop || shop.current || shop.shopError) {
    firstLoad.value = false;
  }
  // If products are not loaded and shop is available, DefaultLayout should have triggered product fetch.
  // If for some reason it didn't (e.g., direct navigation, complex state), you could add a safeguard here:
  // if (shop.current && !shop.isLoadingProducts && shop.products.length === 0 && !shop.productsError) {
  //   shop.fetchProducts(shop.current.slug);
  // }
});

// Watch for store changes to turn off firstLoad if it was true and shop data arrives
watch(() => [shop.isLoadingShop, shop.current, shop.shopError], () => {
    if (!(shop.isLoadingShop && firstLoad.value)) {
        firstLoad.value = false;
    }
});

</script>

<style scoped>
/* Add any specific styles for ProductsPage if needed */
</style>