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

        <div v-if="colStore.list && colStore.list.length > 0 && shop.products.length > 0" class="mb-8 md:mb-10">
          <nav class="flex space-x-3 lg:space-x-4 pb-4 overflow-x-auto hide-scrollbar">
            <button
              @click="selectCollection('all')"
              :class="[
                'flex-shrink-0 px-6 py-2.5 rounded-full text-base font-medium transition-all duration-200 ease-in-out shadow-sm border',
                selectedCollectionHandle === 'all'
                  ? 'bg-green-600 text-white border-green-600 shadow-md transform scale-105'
                  : 'bg-white text-gray-700 border-gray-200 hover:bg-green-50 hover:text-green-700 hover:border-green-300'
              ]"
            >
              <img src="https://placehold.co/24x24/E0F2F7/263238?text=All" alt="All Products Icon" class="inline-block h-5 w-5 rounded-full mr-2 object-cover" />
              All Products
            </button>

            <button
              v-for="collection in colStore.list"
              :key="collection.handle"
              @click="selectCollection(collection.handle)"
              :class="[
                'flex-shrink-0 px-6 py-2.5 rounded-full text-base font-medium transition-all duration-200 ease-in-out shadow-sm border',
                selectedCollectionHandle === collection.handle
                  ? 'bg-green-600 text-white border-green-600 shadow-md transform scale-105'
                  : 'bg-white text-gray-700 border-gray-200 hover:bg-green-50 hover:text-green-700 hover:border-green-300'
              ]"
            >
              <img
                :src="collection.image || 'https://placehold.co/24x24/E0F2F7/263238?text=Col'"
                :alt="`${collection.title} Icon`"
                class="inline-block h-5 w-5 rounded-full mr-2 object-cover"
                onerror="this.onerror=null;this.src='https://placehold.co/24x24/E0F2F7/263238?text=Col';"
              />
              {{ collection.title }}
            </button>
          </nav>
        </div>

        <div v-if="shop.loading || colStore.loading" class="flex flex-col items-center justify-center h-80 bg-white rounded-xl shadow-lg animate-pulse">
          <svg class="animate-spin h-16 w-16 text-green-500 mb-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <p class="text-xl font-semibold text-gray-600">Loading products...</p>
        </div>

        <div v-else-if="shop.products.length === 0" class="text-center py-20 bg-white rounded-xl shadow-lg border border-gray-100 flex flex-col items-center justify-center min-h-[400px]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-28 w-28 mx-auto text-yellow-500 mb-6 drop-shadow-md" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <h2 class="text-4xl font-bold text-gray-800 mb-4">We're Brewing Something Special!</h2>
          <p class="text-lg text-gray-600 max-w-prose mx-auto mb-6">
            Our shelves are currently being stocked with amazing new products.
            We're working hard to bring you a fantastic selection. Please check back soon!
          </p>
          <p class="text-sm text-gray-500">
            In the meantime, feel free to contact us if you have any questions.
          </p>
        </div>

        <div v-else-if="displayedProducts.length > 0">
          <ProductGrid :products="displayedProducts" />
        </div>

        <div v-else class="text-center py-20 bg-white rounded-xl shadow-lg border border-gray-100 flex flex-col items-center justify-center min-h-[400px]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-24 w-24 mx-auto text-green-400 mb-6 drop-shadow-md" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1">
            <path stroke-linecap="round" stroke-linejoin="round" d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" />
          </svg>
          <h2 class="text-3xl font-bold text-gray-800 mb-3">No Treasures Here Yet!</h2>
          <p class="text-lg text-gray-600 max-w-prose mx-auto mb-6">
            It looks like this selection is a bit empty right now. We're working on adding exciting new products.
            Why not explore our <router-link to="/products" class="text-green-600 hover:text-green-700 font-semibold underline">All Products</router-link> section?
          </p>
          <button
            @click="selectCollection('all')"
            class="px-8 py-3 bg-green-600 text-white font-semibold rounded-full shadow-lg hover:bg-green-700 transition duration-300 ease-in-out transform hover:scale-105"
          >
            Browse All Products
          </button>
        </div>
      </section>

      <div v-else class="flex flex-col items-center justify-center min-h-[calc(100vh-200px)] py-20 text-center bg-white rounded-xl shadow-lg m-4">
        <svg class="animate-spin h-16 w-16 text-green-600 mb-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <p class="text-2xl font-semibold text-gray-700 mb-2">Fetching your shop details...</p>
        <p class="text-base text-gray-500">Just a moment while we prepare everything for you.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useShopStore } from '@/stores/shop';
import { useCollectionStore } from '@/stores/collection';
import ProductGrid from '@/components/ProductGrid.vue';
import { useRoute } from 'vue-router';

const shop = useShopStore();
const colStore = useCollectionStore();
const route = useRoute();

const selectedCollectionHandle = ref('all');
const shopSlug = window.location.host.split('.')[0];

// Computed property to determine the products to display
const displayedProducts = computed(() => {
  if (!shop.current) return []; // If shop data not loaded, return empty
  if (selectedCollectionHandle.value === 'all') {
    // If 'all' is selected, filter shop.products by search query if present
    if (route.query.q) {
      const searchTerm = route.query.q.toLowerCase();
      return shop.products.filter(p =>
        p.title.toLowerCase().includes(searchTerm) ||
        (p.description && p.description.toLowerCase().includes(searchTerm))
      );
    }
    return shop.products;
  } else {
    // If a collection is selected, return its products if available and loaded, also filter by search query
    const collectionProducts = colStore.detail && colStore.detail.handle === selectedCollectionHandle.value
      ? colStore.detail.products
      : [];

    if (route.query.q) {
      const searchTerm = route.query.q.toLowerCase();
      return collectionProducts.filter(p =>
        p.title.toLowerCase().includes(searchTerm) ||
        (p.description && p.description.toLowerCase().includes(searchTerm))
      );
    }
    return collectionProducts;
  }
});

// Computed properties for dynamic title and description
const currentSectionTitle = computed(() => {
  if (route.query.q) {
    return `Search Results for "${route.query.q}"`;
  }
  if (selectedCollectionHandle.value === 'all') {
    return 'All Products';
  }
  return colStore.detail?.title || 'Products';
});

const currentSectionDescription = computed(() => {
  if (route.query.q) {
    const totalResults = displayedProducts.value.length;
    return totalResults > 0
      ? `Found ${totalResults} item${totalResults !== 1 ? 's' : ''} matching your search.`
      : `No items matched your search for "${route.query.q}".`;
  }
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
    await shop.fetchShopAndProducts(shopSlug);
  }
  await colStore.loadAll(shopSlug);

  if (route.params.handle) {
    selectedCollectionHandle.value = route.params.handle;
  }
});

// Watch for changes in selectedCollectionHandle to load specific collection products
watch(selectedCollectionHandle, async (newHandle) => {
  if (newHandle !== 'all' && shop.current) {
    await colStore.loadDetail(shopSlug, newHandle);
  }
}, { immediate: true });

// Watch for changes in route query for search
watch(() => route.query.q, () => {
  // No specific action needed here beyond what displayedProducts computed property handles.
  // We can optionally reset selectedCollectionHandle to 'all' if a new search is performed,
  // to ensure search covers all products unless a new collection is clicked.
  // For now, we'll let the search filter the currently selected view (all or specific collection).
});
</script>

<style scoped>
/* Custom scrollbar hide for horizontal tabs */
.hide-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.hide-scrollbar::-webkit-scrollbar {
  display: none;
}
</style>