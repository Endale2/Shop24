<template>
  <div>
    <!-- Breadcrumb and Back Button -->
    <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
      <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
        Back
      </button>
      <router-link :to="`/shops/${shopSlug}`" class="hover:underline">Home</router-link>
      <span>/</span>
      <span v-if="selectedCollection === null">Products</span>
      <span v-else>{{ pageTitle }}</span>
    </nav>
    <!-- End Breadcrumb and Back Button -->

    <div v-if="collections.length > 0" class="flex space-x-2 overflow-x-auto mb-8 pb-2 border-b border-gray-200">
      <button
        class="flex-shrink-0 px-4 py-2 rounded-t border-b-2 transition-colors text-xs font-medium uppercase focus:outline-none"
        :class="selectedCollection === null ? 'border-black text-black bg-gray-50' : 'border-transparent text-gray-500 hover:text-black hover:border-gray-400'"
        @click="selectCollection(null)"
      >
        <span>All</span>
      </button>
      
      <button
        v-for="col in collections"
        :key="col.id"
        class="flex-shrink-0 flex items-center space-x-2 px-4 py-2 rounded-t border-b-2 transition-colors text-xs font-medium uppercase focus:outline-none"
        :class="selectedCollection === col.id ? 'border-black text-black bg-gray-50' : 'border-transparent text-gray-500 hover:text-black hover:border-gray-400'"
        @click="selectCollection(col.id)"
      >
        <img v-if="col.image" :src="col.image" :alt="`${col.title} collection`" class="w-6 h-6 object-contain rounded" />
        <span>{{ col.title }}</span>
      </button>
    </div>

    <h1 class="text-3xl font-bold mb-8 text-gray-900 tracking-tight uppercase">{{ pageTitle }}</h1>
    
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <ProductCard v-for="p in filteredProducts" :key="p.id" :product="p" />
    </div>

    <div v-if="isLoading" class="text-center py-10 text-gray-500">
      Loading products...
    </div>
    <div v-else-if="filteredProducts.length === 0" class="text-center py-10 text-gray-500">
      No products found in this collection.
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import ProductCard from '@/components/ProductCard.vue';
import { fetchAllProducts } from '@/services/product';
import { fetchCollections, fetchCollectionDetail } from '@/services/collections';
import type { Product } from '@/services/product';
import type { Collection } from '@/services/collections';

const route = useRoute();
const shopSlug = route.params.shopSlug as string;

const products = ref<Product[]>([]);
const collections = ref<Collection[]>([]);
const selectedCollection = ref<string | null>(null);
const collectionProducts = ref<Record<string, string[]>>({});
const isLoading = ref(true);

onMounted(async () => {
  try {
    // Fetch products and collections concurrently for better performance
    const [fetchedProducts, fetchedCollections] = await Promise.all([
      fetchAllProducts(shopSlug),
      fetchCollections(shopSlug)
    ]);
    
    products.value = fetchedProducts;
    collections.value = fetchedCollections;

    // Fetch product IDs for each collection
    for (const col of collections.value) {
      // Use fetchCollectionDetail (not fetchCollectionDetails)
      const detail = await fetchCollectionDetail(shopSlug, col.handle);
      collectionProducts.value[col.id] = (detail.products || []).map((p: any) => p.id);
    }
  } catch (error) {
    console.error("Failed to load products or collections:", error);
  } finally {
    isLoading.value = false;
  }
});

function selectCollection(id: string | null) {
  selectedCollection.value = id;
}

// Dynamic title based on the selected collection
const pageTitle = computed(() => {
  if (selectedCollection.value === null) {
    return 'All Products';
  }
  const collection = collections.value.find(c => c.id === selectedCollection.value);
  return collection ? collection.title : 'Products';
});

// Memoizes the products shown on screen
const filteredProducts = computed(() => {
  if (selectedCollection.value === null) {
    return products.value;
  }
  const productIds = collectionProducts.value[selectedCollection.value] || [];
  return products.value.filter(p => productIds.includes(p.id));
});
</script>