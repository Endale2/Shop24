<template>
  <div v-if="isLoading" class="flex min-h-[calc(100vh-200px)] items-center justify-center">
    <p class="text-xl text-gray-500 dark:text-gray-400">Loading product details...</p>
    </div>

  <div v-else-if="error" class="flex min-h-[calc(100vh-200px)] flex-col items-center justify-center text-center">
    <p class="text-xl text-red-500 dark:text-red-400">Error loading product.</p>
    <p class="text-gray-600 dark:text-gray-300 mt-2">{{ error }}</p>
    <router-link
      to="/"
      class="mt-6 px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700 transition-colors"
    >
      Go to Homepage
    </router-link>
  </div>

  <div v-else-if="product" class="container mx-auto px-4 py-8">
    <div class="md:flex md:items-start md:space-x-8 lg:space-x-12">
      <div class="md:w-1/2 mb-8 md:mb-0">
        <div v-if="product.images && product.images.length > 0" class="bg-gray-100 dark:bg-gray-700 rounded-lg shadow-lg overflow-hidden">
          <img
            :src="activeImageUrl || product.images[0]"
            :alt="product.name || 'Product image'"
            class="w-full h-auto object-cover aspect-square max-h-[500px] md:max-h-[600px] transition-all duration-300"
          >
          <div v-if="product.images.length > 1" class="grid grid-cols-4 sm:grid-cols-5 gap-2 p-2 bg-gray-200 dark:bg-gray-600">
            <button
              v-for="(image, index) in product.images"
              :key="index"
              @click="setActiveImage(image)"
              class="rounded overflow-hidden border-2 transition-colors"
              :class="image === activeImageUrl ? 'border-indigo-500' : 'border-transparent hover:border-indigo-300 dark:hover:border-indigo-400'"
            >
              <img :src="image" :alt="`${product.name} thumbnail ${index + 1}`" class="w-full h-full object-cover aspect-square">
            </button>
          </div>
        </div>
        <div v-else class="bg-gray-100 dark:bg-gray-700 rounded-lg shadow-md flex items-center justify-center h-96 aspect-square">
          <p class="text-gray-500 dark:text-gray-400">No Image Available</p>
        </div>
      </div>

      <div class="md:w-1/2">
        <h1 class="text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-2">{{ product.name || 'Product Name Unavailable' }}</h1>
        <p v-if="product.category" class="text-gray-600 dark:text-gray-400 text-md mb-4">{{ product.category }}</p>

        <p class="text-3xl font-semibold text-indigo-600 dark:text-indigo-400 mb-6">
          ${{ product.price?.toFixed(2) || 'Price Unavailable' }}
        </p>
        
        <div class="prose prose-sm sm:prose dark:prose-invert max-w-none mb-6 text-gray-700 dark:text-gray-300">
            <p>{{ product.description || 'No description available.' }}</p>
        </div>

        <div class="mb-6">
            <label for="quantity" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Quantity:</label>
            <input type="number" id="quantity" name="quantity" v-model.number="quantity" min="1" class="w-20 p-2 border border-gray-300 dark:border-gray-600 rounded-md dark:bg-gray-700 dark:text-white focus:ring-indigo-500 focus:border-indigo-500">
        </div>

        <button 
          class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-3 px-6 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800"
          @click="addToCart"
        >
          Add to Cart
        </button>

        <div class="mt-8 pt-6 border-t border-gray-200 dark:border-gray-700">
            <h3 class="text-md font-semibold text-gray-700 dark:text-gray-200 mb-3">Additional Information</h3>
            <ul class="text-sm text-gray-600 dark:text-gray-400 space-y-1">
                <li><span class="font-medium">SKU:</span> {{ product.sku || 'N/A' }}</li>
                <li><span class="font-medium">Availability:</span> {{ product.stock > 0 ? 'In Stock' : 'Out of Stock' }}</li>
                <li v-if="product.brand"><span class="font-medium">Brand:</span> {{ product.brand }}</li>
                <li v-if="product.weight"><span class="font-medium">Weight:</span> {{ product.weight }}</li>
            </ul>
        </div>
      </div>
    </div>

    </div>

  <div v-else class="flex min-h-[calc(100vh-200px)] items-center justify-center">
    <p class="text-xl text-gray-500 dark:text-gray-400">Product not found.</p>
     <router-link
      to="/products"
      class="ml-4 px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-white rounded hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
    >
      Back to Products
    </router-link>
  </div>
</template>

<script setup>
import { ref, onBeforeMount, watch, defineProps } from 'vue'
import { RouterLink } from 'vue-router' // useRoute is not needed if all params come via props
import { fetchProductBySlug } from '@/services/product'

// Define props passed from the router (due to props: true in route config)
const props = defineProps({
  shopSlug: {
    type: String,
    required: true
  },
  productSlug: {
    type: String,
    required: true
  }
})

const product = ref(null)
const isLoading = ref(true)
const error = ref(null)
const quantity = ref(1) // For the quantity input
const activeImageUrl = ref(null) // For the image gallery

async function loadProductData(currentShopSlug, currentProductSlug) {
  isLoading.value = true
  error.value = null
  product.value = null
  activeImageUrl.value = null; // Reset active image

  try {
    // Use the props directly
    const fetchedProduct = await fetchProductBySlug(currentShopSlug, currentProductSlug)
    if (fetchedProduct) {
      product.value = fetchedProduct
      if (fetchedProduct.images && fetchedProduct.images.length > 0) {
        activeImageUrl.value = fetchedProduct.images[0];
      }
    } else {
      console.warn(`Product not found for shopSlug: ${currentShopSlug}, productSlug: ${currentProductSlug}`)
      // Error ref will be null, and product.value will be null, leading to "Product not found" message
    }
  } catch (e) {
    console.error('Failed to fetch product:', e)
    error.value = e.response?.data?.message || e.message || 'An unexpected error occurred while loading the product.'
  } finally {
    isLoading.value = false
  }
}

// Set active image for gallery
const setActiveImage = (imageUrl) => {
  activeImageUrl.value = imageUrl;
}

// Initial data load
onBeforeMount(() => {
  loadProductData(props.shopSlug, props.productSlug)
})

// Watch for changes in props (e.g., if navigating between product pages directly)
watch(
  () => [props.shopSlug, props.productSlug],
  ([newShopSlug, newProductSlug], [oldShopSlug, oldProductSlug]) => {
    // Check if either slug has actually changed
    if (newShopSlug !== oldShopSlug || newProductSlug !== oldProductSlug) {
      loadProductData(newShopSlug, newProductSlug)
    }
  },
  { immediate: false } // `immediate: false` because `onBeforeMount` handles initial load
)

// Placeholder for addToCart functionality
const addToCart = () => {
  if (product.value) {
    // In a real app, you'd dispatch an action to a Pinia cart store
    console.log(`Adding ${quantity.value} of ${product.value.name} (ID: ${product.value.id}) to cart.`)
    alert(`${quantity.value} x ${product.value.name} added to cart! (This is a placeholder)`)
    // Example: cartStore.addItem({ product: product.value, quantity: quantity.value })
  }
}
</script>

<style scoped>
/* Add any specific styles for this page here */
.prose :where(p):where(:not(:where([class~="not-prose"] *))) {
  margin-top: 0.75em; /* Adjust spacing if needed */
  margin-bottom: 0.75em;
}
/* Ensure aspect ratio for images if not handled by Tailwind's aspect-square */
/* img { display: block; } */
</style>