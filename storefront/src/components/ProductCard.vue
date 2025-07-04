<template>
  <div class="border rounded p-4 hover:shadow">
    <img
      :src="product.main_image"
      alt=""
      class="w-full h-48 object-cover mb-2"
    />
    <h2 class="font-semibold">{{ product.name }}</h2>
    <p class="text-gray-700">
      <template v-if="product.price > 0 && (!product.variants || product.variants.length === 0)">
        ${{ product.price.toFixed(2) }}
      </template>
      <template v-else-if="product.starting_price > 0">
        <span class="italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
      </template>
      <template v-else>
        N/A
      </template>
    </p>
    <router-link :to="`/products/${product.slug}`" class="text-blue-500">
      View
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface Props {
  product: any
}
defineProps<Props>()
</script>
