<template>
  <div v-if="product" class="max-w-4xl mx-auto py-8 space-y-8">
    <!-- Hero image (first image or placeholder) -->
    <div class="rounded-lg overflow-hidden shadow-lg h-96 bg-gray-100 flex items-center justify-center">
      <img
        v-if="product.images && product.images.length"
        :src="product.images[0]"
        alt="Product image"
        class="object-cover w-full h-full"
      />
      <svg
        v-else
        class="w-16 h-16 text-gray-400"
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path
          d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-4 4 4 4-4V5h-2v6l-4-4-4 4z"
        />
      </svg>
    </div>

    <!-- Title, Category & Price -->
    <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-4">
      <div>
        <h1 class="text-4xl font-extrabold text-gray-900">{{ product.name }}</h1>
        <p class="text-sm text-gray-500 uppercase">{{ product.category || 'Uncategorized' }}</p>
      </div>
      <div class="flex items-center space-x-4">
        <span class="text-3xl font-bold text-green-600">
          ${{ product.price.toFixed(2) }}
        </span>
        <button
          class="px-6 py-2 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition"
        >
          Add to Cart
        </button>
      </div>
    </div>

    <!-- Description -->
    <section class="prose prose-lg max-w-none">
      <h2>Description</h2>
      <p>{{ product.description }}</p>
    </section>

    <!-- Variants -->
    <section v-if="product.variants && product.variants.length" class="space-y-4">
      <h2 class="text-2xl font-semibold text-gray-800">Variants</h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div
          v-for="(v, idx) in product.variants"
          :key="idx"
          class="p-4 border rounded-lg hover:shadow-lg transition"
        >
          <p class="text-gray-700 mb-2">
            {{ Object.entries(v.options).map(([k, val]) => `${k}: ${val}`).join(', ') }}
          </p>
          <p class="font-bold text-gray-900">${{ v.price.toFixed(2) }}</p>
        </div>
      </div>
    </section>

    <!-- Timestamps -->
    <footer class="text-sm text-gray-500 space-y-1">
      <p><strong>Created:</strong> {{ formatDate(product.createdAt) }}</p>
      <p><strong>Last Updated:</strong> {{ formatDate(product.updatedAt) }}</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue'
import { useRoute } from 'vue-router'
import { fetchProductById } from '@/services/product'

const route = useRoute()
const product = ref(null)

onBeforeMount(async () => {
  const host = window.location.host.split(':')[0]
  const shopId = host.split('.')[0]
  product.value = await fetchProductById(shopId, route.params.id)
})

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleString(undefined, {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    : '—'
}
</script>

<style scoped>
/* Ensure your project includes the "prose" classes from Tailwind Typography for the description */
</style>
