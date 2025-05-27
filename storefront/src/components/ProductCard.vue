<template>
  <div class="bg-white border border-gray-200 rounded-lg overflow-hidden shadow-md hover:shadow-xl transition-all duration-300 ease-in-out flex flex-col h-full group relative">
    <div v-if="product.isNew || product.onSale" class="absolute top-3 left-3 z-10">
      <span v-if="product.isNew" class="bg-green-500 text-white text-xs font-semibold px-2.5 py-1 rounded-full">NEW</span>
      <span v-if="product.onSale" class="bg-red-500 text-white text-xs font-semibold px-2.5 py-1 rounded-full ml-1">SALE</span>
    </div>

    <div class="w-full h-48 sm:h-56 bg-gray-100 flex items-center justify-center overflow-hidden relative">
      <img
        v-if="product.images?.length"
        :src="product.images[0]"
        :alt="product.name"
        class="w-full h-full object-cover object-center transition-transform duration-500 group-hover:scale-110"
        loading="lazy"
      />
      <div v-else class="text-gray-400 text-sm flex flex-col items-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-300 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 3.75v4.5m0-4.5h4.5m-4.5 0L9 9M3.75 20.25v-4.5m0 4.5h4.5m-4.5 0L9 15M20.25 3.75h-4.5m4.5 0v4.5m0-4.5L15 9m5.25 11.25h-4.5m4.5 0v-4.5m0 4.5L15 15" />
        </svg>
        No Image
      </div>
      <div class="absolute inset-0 bg-black bg-opacity-20 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <button @click.prevent.stop="quickViewProduct(product)" class="bg-white text-slate-700 hover:bg-indigo-500 hover:text-white text-sm font-medium py-2 px-4 rounded-md shadow-sm transition-colors duration-200">
          Quick View
        </button>
      </div>
    </div>

    <div class="p-4 flex-1 flex flex-col">
      <div class="flex-grow">
        <p v-if="product.category?.name" class="text-xs text-gray-500 mb-1 uppercase tracking-wider">{{ product.category.name }}</p>
        <h3 class="text-md sm:text-lg font-semibold text-slate-800 truncate group-hover:text-indigo-600 transition-colors duration-200" :title="product.name">
          {{ product.name }}
        </h3>
        <p v-if="product.shortDescription" class="text-sm text-gray-600 mt-1 mb-2 h-10 overflow-hidden text-ellipsis">
          {{ product.shortDescription }}
        </p>
      </div>

      <div class="mt-auto pt-2 flex items-center justify-between">
        <p class="text-indigo-600 text-lg sm:text-xl font-bold">
          ${{ product.price.toFixed(2) }}
          <span v-if="product.originalPrice && product.originalPrice > product.price" class="text-sm text-gray-500 line-through ml-2">
            ${{ product.originalPrice.toFixed(2) }}
          </span>
        </p>
        <button @click.prevent.stop="addToCartInternal(product)" class="text-slate-500 hover:text-indigo-600 p-1 rounded-full hover:bg-indigo-50 transition-colors duration-200" title="Add to Cart">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
             <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 00-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 00-16.536-1.84M7.5 14.25L5.106 5.272M6 20.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm12.75 0a.75.75 0 11-1.5 0 .75.75 0 011.5 0z" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  product: { type: Object, required: true }
})

const emit = defineEmits(['quick-view', 'add-to-cart'])

const quickViewProduct = (product) => {
  emit('quick-view', product)
}

const addToCartInternal = (product) => {
  emit('add-to-cart', product)
}
</script>

<style scoped>
.text-ellipsis {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>