<template>
  <div class="border rounded p-4 hover:shadow">
    <img
      :src="product.main_image"
      alt=""
      class="w-full h-48 object-cover mb-2"
    />
    <h2 class="font-semibold">{{ product.name }}</h2>
    
    <!-- Discount Badge -->
    <div v-if="product.discounts && product.discounts.length > 0" class="mb-2">
      <div class="inline-flex items-center px-2 py-1 bg-green-100 text-green-800 text-xs rounded-full">
        <span class="font-medium">
          {{ product.discounts[0].type === 'percentage' ? `${product.discounts[0].value}% OFF` : `$${product.discounts[0].value} OFF` }}
        </span>
      </div>
    </div>
    
    <p class="text-gray-700">
      <template v-if="product.price > 0 && (!product.variants || product.variants.length === 0)">
        <span v-if="product.display_price && product.display_price < product.price" class="line-through text-gray-500 mr-2">
          ${{ product.price.toFixed(2) }}
        </span>
        <span class="font-semibold text-green-600">
          ${{ (product.display_price || product.price).toFixed(2) }}
        </span>
      </template>
      <template v-else-if="product.starting_price > 0">
        <span class="italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
      </template>
      <template v-else>
        N/A
      </template>
    </p>
    <router-link :to="`/shops/${shopSlug}/products/${product.slug}`" class="text-blue-500">
      View
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import { useRoute } from 'vue-router'

interface Props {
  product: any
}
defineProps<Props>()

const route = useRoute()
const shopSlug = route.params.shopSlug as string
</script>
