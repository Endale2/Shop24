<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-6">Search Results for "{{ searchQuery }}"</h1>
    <Loader v-if="isLoading" text="Searching products..." />
    <div v-else>
      <div v-if="products.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <ProductCard v-for="p in products" :key="p.id" :product="p" />
      </div>
      <div v-else class="text-center py-12 text-gray-500">
        No products found for your search.
      </div>
      <div v-if="totalPages > 1" class="flex justify-center mt-8">
        <button
          class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
          :disabled="currentPage === 1"
          @click="goToPage(currentPage - 1)"
        >Prev</button>
        <button
          v-for="page in totalPages"
          :key="page"
          class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
          :class="{ 'bg-black text-white': page === currentPage }"
          @click="goToPage(page)"
        >{{ page }}</button>
        <button
          class="px-4 py-2 mx-1 rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-100"
          :disabled="currentPage === totalPages"
          @click="goToPage(currentPage + 1)"
        >Next</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Loader from '@/components/Loader.vue';
import ProductCard from '@/components/ProductCard.vue';
import { fetchProductsSearchPaginated } from '@/services/product';
import { getCurrentShopSlug } from '@/services/shop';

const route = useRoute();
const router = useRouter();
const searchQuery = ref(route.query.q ? String(route.query.q) : '');
const currentPage = ref(Number(route.query.page) || 1);
const pageSize = 20;
const products = ref([]);
const totalProducts = ref(0);
const isLoading = ref(false);

async function loadResults(page = 1) {
  isLoading.value = true;
  try {
    const shopSlug = getCurrentShopSlug();
    if (!shopSlug || !searchQuery.value) return;
    const { products: fetched, total, page: backendPage } = await fetchProductsSearchPaginated(shopSlug, searchQuery.value, page, pageSize);
    products.value = fetched;
    totalProducts.value = total;
    currentPage.value = backendPage;
  } catch (error) {
    console.error('Failed to search products:', error);
  } finally {
    isLoading.value = false;
  }
}

watch(() => [route.query.q, route.query.page], ([q, p]) => {
  searchQuery.value = q ? String(q) : '';
  currentPage.value = Number(p) || 1;
  loadResults(currentPage.value);
});

onMounted(() => {
  loadResults(currentPage.value);
});

function goToPage(page: number) {
  if (page < 1 || page > totalPages.value) return;
  router.push({ path: route.path, query: { ...route.query, page } });
}

const totalPages = computed(() => Math.ceil(totalProducts.value / pageSize));
</script>
