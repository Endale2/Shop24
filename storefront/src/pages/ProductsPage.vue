<template>
  <div>
    <Breadcrumbs :items="[
      { back: true, label: 'Back' },
      { label: 'Home', to: `/${shopSlug}/` },
      { label: selectedCollection === null ? 'Products' : pageTitle }
    ]" />

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
    
    <Loader v-if="isLoading" text="Loading products..." />
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <ProductCard v-for="p in filteredProducts" :key="p.id" :product="p" />
    </div>

    <div v-if="!isLoading && filteredProducts.length === 0" class="text-center py-10 text-gray-500">
      No products found in this collection.
    </div>

    <div v-if="totalPages > 1" class="flex justify-center mt-8">
      <button
        class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
        :disabled="currentPage === 1"
        @click="goToPage(currentPage - 1)"
      >
        Prev
      </button>
      <button
        v-for="page in totalPages"
        :key="page"
        class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
        :class="{ 'bg-black text-white': page === currentPage }"
        @click="goToPage(page)"
      >
        {{ page }}
      </button>
      <button
        class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
        :disabled="currentPage === totalPages"
        @click="goToPage(currentPage + 1)"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import ProductCard from '@/components/ProductCard.vue';
import { fetchProductsPaginated } from '@/services/product';
import { fetchCollections } from '@/services/collections';
import type { Product } from '@/services/product';
import type { Collection } from '@/services/collections';
import Loader from '@/components/Loader.vue';
import Breadcrumbs from '@/components/Breadcrumbs.vue';

const route = useRoute();
const shopSlug = route.params.shopSlug as string;
const products = ref<Product[]>([]);
const collections = ref<Collection[]>([]);
const selectedCollection = ref<string | null>(null);
const collectionProducts = ref<Record<string, string[]>>({});
const isLoading = ref(true);
const currentPage = ref(1);
const pageSize = 20;
const totalProducts = ref(0);

async function loadProducts(page = 1) {
  isLoading.value = true;
  try {
    if (!shopSlug) return;
    const { products: fetchedProducts, total, page: backendPage } = await fetchProductsPaginated(shopSlug, page, pageSize);
    products.value = fetchedProducts;
    totalProducts.value = total;
    currentPage.value = backendPage;
  } catch (error) {
    console.error('Failed to load products:', error);
  } finally {
    isLoading.value = false;
  }
}

onMounted(async () => {
  if (!shopSlug) return;
  // Load collections first
  try {
    collections.value = await fetchCollections(shopSlug);
  } catch (error) {
    console.error('Failed to load collections:', error);
  }
  // Then load products
  loadProducts(currentPage.value);
});

function selectCollection(id: string | null) {
  selectedCollection.value = id;
}

function goToPage(page: number) {
  if (page < 1 || page > totalPages.value) return;
  currentPage.value = page;
  loadProducts(page);
}

const totalPages = computed(() => Math.ceil(totalProducts.value / pageSize));

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