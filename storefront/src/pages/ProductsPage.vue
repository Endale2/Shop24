<template>
  <div class="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-green-50 text-gray-900 antialiased font-sans py-12 px-4 sm:px-6 lg:px-8">
    <div class="container mx-auto">
      <section v-if="shop.current">
        <div class="mb-8 md:mb-10 text-center sm:text-left">
          <h1 class="text-3xl sm:text-4xl font-bold tracking-tight text-gray-900">
            {{ currentSectionTitle }}
          </h1>
          <p class="mt-2 text-md text-gray-700">
            {{ currentSectionDescription }}
          </p>
        </div>

        <!-- Collections Tab Bar -->
        <div v-if="colStore.list && colStore.list.length > 0" class="mb-8 md:mb-10">
          <nav class="flex space-x-4 lg:space-x-6 pb-4 overflow-x-auto hide-scrollbar">
            <!-- All Products Tab -->
            <button
              @click="selectCollection('all')"
              :class="[
                'flex-shrink-0 px-5 py-2 rounded-full text-base font-medium transition-all duration-200 shadow-sm',
                selectedCollectionHandle === 'all'
                  ? 'bg-green-600 text-white shadow-md'
                  : 'bg-white text-gray-700 border border-gray-200 hover:bg-gray-100 hover:border-green-300'
              ]"
            >
              <img src="https://placehold.co/24x24/E0F2F7/263238?text=All" alt="All Products Icon" class="inline-block h-6 w-6 rounded-full mr-2 object-cover" />
              All Products
            </button>

            <!-- Individual Collection Tabs -->
            <button
              v-for="collection in colStore.list"
              :key="collection.handle"
              @click="selectCollection(collection.handle)"
              :class="[
                'flex-shrink-0 px-5 py-2 rounded-full text-base font-medium transition-all duration-200 shadow-sm',
                selectedCollectionHandle === collection.handle
                  ? 'bg-green-600 text-white shadow-md'
                  : 'bg-white text-gray-700 border border-gray-200 hover:bg-gray-100 hover:border-green-300'
              ]"
            >
              <img
                :src="collection.image || 'https://placehold.co/24x24/E0F2F7/263238?text=Col'"
                :alt="`${collection.title} Icon`"
                class="inline-block h-6 w-6 rounded-full mr-2 object-cover"
                onerror="this.onerror=null;this.src='https://placehold.co/24x24/E0F2F7/263238?text=Col';"
              />
              {{ collection.title }}
            </button>
          </nav>
        </div>

        <!-- Product Grid Display -->
        <div v-if="shop.loading || colStore.loading" class="flex items-center justify-center h-64 bg-white rounded-xl shadow-lg">
          <svg class="animate-spin h-14 w-14 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
        </div>
        <div v-else-if="displayedProducts.length > 0">
          <ProductGrid :products="displayedProducts" />
        </div>
        <div v-else class="text-center py-16 md:py-20 bg-white rounded-xl shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1">
            <path stroke-linecap="round" stroke-linejoin="round" d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" />
          </svg>
          <h2 class="text-2xl font-semibold text-gray-700 mb-2">No Products Found</h2>
          <p class="text-gray-500">It looks like there are no products in this selection yet. Please try another collection.</p>
        </div>
      </section>

      <!-- Initial Shop Loading State (when shop.current is not yet available) -->
      <div v-else class="flex flex-col items-center justify-center min-h-[calc(100vh-200px)] py-20 text-center bg-white rounded-xl shadow-lg m-4">
        <svg class="animate-spin h-14 w-14 text-green-600 mb-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <p class="text-2xl font-semibold text-gray-700 mb-2">Loading your amazing shop...</p>
        <p class="text-base text-gray-500">Please wait a moment.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useShopStore } from '@/stores/shop'
import { useCollectionStore } from '@/stores/collection' // Import collection store
import ProductGrid from '@/components/ProductGrid.vue'

const shop = useShopStore()
const colStore = useCollectionStore() // Initialize collection store

const selectedCollectionHandle = ref('all'); // 'all' for all products, or a collection handle
const shopSlug = window.location.host.split('.')[0]; // Derive shopSlug

// Computed property to determine the products to display
const displayedProducts = computed(() => {
  if (!shop.current) return []; // If shop data not loaded, return empty
  if (selectedCollectionHandle.value === 'all') {
    return shop.products;
  } else {
    // If a collection is selected, return its products if available and loaded
    return colStore.detail && colStore.detail.handle === selectedCollectionHandle.value
      ? colStore.detail.products
      : [];
  }
});

// Computed properties for dynamic title and description
const currentSectionTitle = computed(() => {
  if (selectedCollectionHandle.value === 'all') {
    return 'All Products';
  }
  return colStore.detail?.title || 'Products'; // Fallback if collection detail not loaded yet
});

const currentSectionDescription = computed(() => {
  if (selectedCollectionHandle.value === 'all') {
    return 'Explore our full collection of available items.';
  }
  return colStore.detail?.description || 'Discover a curated selection of items.';
});

// Function to handle collection tab clicks
const selectCollection = (handle) => {
  selectedCollectionHandle.value = handle;
};

// On mount, load all shop products and collections
onMounted(async () => {
  if (!shop.current) {
    // Assuming useShopStore has a loadAll method if needed, otherwise data might come from context/initial load
    // If shop data isn't guaranteed to be loaded by a parent, uncomment/implement loadAll here:
    // await shop.loadAll(shopSlug);
  }
  await colStore.loadAll(shopSlug); // Load all collections for the tabs
});

// Watch for changes in selectedCollectionHandle to load specific collection products
watch(selectedCollectionHandle, async (newHandle) => {
  if (newHandle !== 'all' && shop.current) { // Only load detail if a specific collection is selected and shop is loaded
    await colStore.loadDetail(shopSlug, newHandle);
  }
}, { immediate: true }); // Run immediately to set initial products (All Products)
</script>

<style scoped>
/* Custom scrollbar hide for horizontal tabs */
.hide-scrollbar {
  -ms-overflow-style: none;  /* IE and Edge */
  scrollbar-width: none;  /* Firefox */
}
.hide-scrollbar::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Opera */
}
</style>
