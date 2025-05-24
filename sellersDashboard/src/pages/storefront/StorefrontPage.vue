<template>
  <div class="storefront-page bg-gray-100 text-gray-800 font-sans flex flex-col min-h-screen">
    <header class="storefront-header bg-gradient-to-r from-blue-600 to-indigo-700 text-white py-8 sm:py-12 shadow-xl">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl sm:text-5xl font-bold font-heading mb-2 tracking-tight">
          {{ shop?.name || 'Unnamed Shop' }}
        </h1>
        <p v-if="shop?.description" class="shop-description text-lg sm:text-xl text-indigo-100 max-w-3xl mx-auto">
          {{ shop.description }}
        </p>
      </div>
    </header>

    <main class="storefront-main container mx-auto px-4 sm:px-6 lg:px-8 flex-grow py-8 sm:py-12">
      <div v-if="loading" class="loading-state flex flex-col items-center justify-center text-center p-8 sm:p-12 mt-8 bg-white rounded-xl shadow-lg">
        <div class="spinner w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full animate-spin mb-6"></div>
        <p class="text-xl text-gray-600 font-medium">Loading products...</p>
      </div>
      <div v-else-if="error" class="error-state flex flex-col items-center justify-center text-center p-8 sm:p-12 mt-8 bg-white rounded-xl shadow-lg border-l-4 border-red-500">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-red-500 mb-4" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
        </svg>
        <p class="text-xl text-red-700 font-semibold mb-2">Oops! Something went wrong.</p>
        <p class="text-gray-600">{{ error }}</p>
        <p class="text-sm text-gray-500 mt-3">Please try refreshing the page or contact support if the issue persists.</p>
      </div>
      <div v-else>
        <div v-if="products.length" class="product-grid grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 sm:gap-8">
          <div v-for="product in products" :key="product.id" class="product-card bg-white rounded-xl shadow-lg overflow-hidden flex flex-col transition-all duration-300 ease-in-out hover:shadow-2xl hover:-translate-y-1.5 group">
            <div class="product-image-wrapper w-full aspect-[4/3] bg-gray-200 relative overflow-hidden">
              <img
                :src="getProductImage(product)"
                :alt="product.name"
                class="product-image absolute inset-0 w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                @error="onImageError"
              />
            </div>
            <div class="product-info p-5 sm:p-6 flex flex-col flex-grow">
              <h3 class="product-name text-xl sm:text-2xl font-semibold font-heading text-gray-900 mb-2">
                {{ product.name }}
              </h3>
              <p v-if="product.description" class="product-description text-sm text-gray-600 mb-4 leading-relaxed flex-grow min-h-[4.5rem] line-clamp-3">
                {{ product.description }} </p>
              <p v-else class="flex-grow min-h-[4.5rem]"></p> <div class="mt-auto">
                <p class="product-price text-2xl font-bold text-blue-600 mb-4 text-right">
                  ${{ product.price?.toFixed(2) || 'N/A' }}
                </p>
                <button class="add-to-cart-button block w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-4 rounded-lg transition duration-200 ease-in-out transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
                  View Details
                </button>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="no-products-found flex flex-col items-center justify-center text-center p-8 sm:p-12 mt-8 bg-white rounded-xl shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
             <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
          </svg>
          <p class="text-xl text-gray-700 font-semibold mb-2">No Products Found</p>
          <p class="text-gray-500">It looks like there are no products in this shop yet. Please check back soon!</p>
        </div>
      </div>
    </main>

    <footer class="storefront-footer bg-gray-800 text-gray-400 py-8 text-center text-sm">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <p>&copy; {{ new Date().getFullYear() }} {{ shop?.name || 'Our Store' }}. All rights reserved.</p>
        <p class="mt-1">Powered by Vue.js & Tailwind CSS</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
// import { useRoute }      from 'vue-router' // Not used in the provided script logic for fetching
import { useShopStore }  from '@/store/shops' // Assuming this path is correct
import { productService }from '@/services/product' // Assuming this path is correct

// const route    = useRoute() // Not used in current fetching logic
const shops    = useShopStore()
const shopId   = window.location.hostname.split('.')[0] // This might be too simple for production, consider router params
const shop     = ref(null)
const products = ref([])
const loading  = ref(true)
const error    = ref(null)

const placeholderImage = 'https://via.placeholder.com/400x300.png?text=No+Image+Available' // Adjusted placeholder size

const getProductImage = (product) => {
  if (product.images && product.images.length > 0 && product.images[0]) {
    return typeof product.images[0] === 'string' ? product.images[0] : product.images[0].url || placeholderImage;
  }
  return placeholderImage;
}

const onImageError = (event) => {
  event.target.src = placeholderImage;
}

// Description truncation is now handled by Tailwind's line-clamp plugin if you install and configure it.
// If not using the plugin, you can revert to the JS function or simply let the text flow (or set a fixed height).
// For simplicity, I've used `line-clamp-3` in the template. If you prefer the JS function:
/*
const truncateDescription = (description, maxLength = 100) => {
  if (!description) return '';
  if (description.length <= maxLength) return description;
  return description.substring(0, maxLength) + '...';
}
*/

onMounted(async () => {
  loading.value = true
  error.value = null // Reset error on new load
  try {
    shop.value = await shops.fetchPublicShop(shopId)
    products.value = await productService.fetchPublicByShop(shopId)
  } catch (e) {
    console.error('Error fetching storefront data:', e)
    error.value = 'Could not load this storefront. The shop might not exist or there was a network issue.'
    if (e.response && e.response.status === 404) {
        error.value = `Sorry, we couldn't find a shop with the ID '${shopId}'.`
    }
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
/*
  With Tailwind CSS, you generally want to avoid writing custom CSS here.
  Most styles are applied directly in the template using utility classes.
  If you have very specific needs not covered by Tailwind utilities,
  you might add them here, or better, try to create a custom Tailwind component/plugin.

  For this example, all styling has been moved to Tailwind classes in the <template>.
  The `font-sans` and `font-heading` in the template rely on your tailwind.config.js `theme.extend.fontFamily`.
*/

/* Example: If you wanted a global body background if not doing it on the top div:
body {
  @apply bg-gray-100;
}
*/
</style>