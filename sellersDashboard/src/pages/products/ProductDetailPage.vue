<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button
      @click="goBack"
      class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group"
    >
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Products</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="mt-3 text-lg">Loading product details...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm"
      role="alert"
    >
      <strong class="font-bold">Oops!</strong>
      <span class="ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Please try again or contact support if the issue persists.</p>
    </div>

    <div v-else-if="product.id" class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8 animate-fade-in">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <div class="space-y-4">
          <div class="w-full h-80 sm:h-96 md:h-80 lg:h-96 rounded-xl overflow-hidden shadow-md border border-gray-200 relative bg-gray-50 flex items-center justify-center">
            <img
              v-if="mainImage"
              :src="mainImage"
              alt="Main product image"
              class="max-w-full max-h-full object-contain transition-transform duration-300 ease-in-out transform hover:scale-105"
            />
            <div v-else class="w-full h-full flex items-center justify-center bg-gray-100">
              <PlaceholderIcon class="w-20 h-20 text-gray-400" />
            </div>
          </div>

          <div class="flex space-x-3 overflow-x-auto pb-2">
            <img
              v-for="(img, idx) in product.images"
              :key="`prod-img-${idx}`"
              :src="img"
              alt="Product thumbnail"
              @click="setMainImage(img)"
              :class="{ 'border-green-500 ring-2 ring-green-500': img === mainImage, 'border-gray-300': img !== mainImage }"
              class="w-20 h-20 object-cover rounded-lg cursor-pointer border-2 shadow-sm transform hover:scale-105 transition-all duration-200 ease-in-out"
            />
            <img
              v-for="(variant, idx) in product.variants"
              :key="`variant-img-${idx}`"
              v-if="variant && variant.image" :src="variant.image"
              :alt="`Variant image for ${formatVariant(variant.options)}`"
              @click="setMainImage(variant.image)"
              :class="{ 'border-purple-500 ring-2 ring-purple-500': variant.image === mainImage, 'border-gray-300': variant.image !== mainImage }"
              class="w-20 h-20 object-cover rounded-lg cursor-pointer border-2 shadow-sm transform hover:scale-105 transition-all duration-200 ease-in-out"
            />
          </div>
        </div>

        <div class="lg:sticky lg:top-8 self-start space-y-6">
          <div>
            <h1 class="text-4xl sm:text-5xl font-extrabold text-gray-900 leading-tight">{{ product.name }}</h1>
            <p class="text-lg text-gray-600 mt-2">{{ product.category || 'Uncategorized' }}</p>
          </div>

          <p class="text-4xl sm:text-5xl font-bold text-green-700 animate-pulse-price">
            {{ formatPrice(product.price) }}
          </p>

          <div class="mt-4">
            <p class="text-lg text-gray-800">
              Availability:
              <span
                :class="{
                  'text-green-600 font-semibold': product.stock > 0,
                  'text-red-600 font-semibold': product.stock <= 0
                }"
              >
                {{ product.stock > 0 ? `${product.stock} In Stock` : 'Out of Stock' }}
              </span>
            </p>
          </div>
          <p class="text-gray-700 leading-relaxed text-base md:text-lg">
            {{ product.description || 'No description available.' }}
          </p>

          <div v-if="displayableVariants.length > 0" class="mt-6 border-t border-gray-200 pt-6">
            <h3 class="font-semibold text-gray-800 text-xl mb-4">Available Variants</h3>
            <ul class="space-y-3">
              <li
                v-for="(v, i) in displayableVariants"
                :key="i"
                class="flex flex-col sm:flex-row sm:justify-between items-start sm:items-center p-4 bg-green-50 rounded-lg shadow-sm border border-green-100 transition-all duration-200 ease-in-out hover:shadow-md hover:border-green-200"
              >
                <div class="flex items-center gap-2 mb-2 sm:mb-0">
                  <img
                    v-if="v && v.image" :src="v.image"
                    alt="Variant thumbnail"
                    class="w-10 h-10 object-cover rounded-md border border-gray-200"
                  />
                  <span class="text-gray-700 font-medium text-base">
                    {{ formatVariant(v.options) }}
                  </span>
                </div>
                <div class="text-right">
                  <span class="font-bold text-green-800 text-lg sm:text-xl">
                    {{ formatPrice(v.price) }}
                  </span>
                  <p class="text-sm font-medium" :class="{ 'text-green-600': v.stock > 0, 'text-red-600': v.stock <= 0 }">
                    {{ v.stock > 0 ? `${v.stock} in stock` : 'Out of Stock' }}
                  </p>
                </div>
                </li>
            </ul>
          </div>

          <div class="flex flex-col sm:flex-row space-y-3 sm:space-y-0 sm:space-x-4 mt-8">
            <button
              @click="goToEditProduct(product.id)"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out transform hover:-translate-y-0.5"
            >
              <PencilIcon class="w-5 h-5 mr-2" />
              Edit Product
            </button>
            <button
              @click="confirmDeleteProduct"
              class="flex-1 w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-red-300 text-base font-medium rounded-md shadow-sm text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition duration-150 ease-in-out"
            >
              <TrashIcon class="w-5 h-5 mr-2" />
              Delete Product
            </button>
          </div>
        </div>
      </div>

      <div class="border-t border-gray-200 pt-6 text-sm text-gray-500 text-right space-y-1">
        <p><strong>Created:</strong> {{ formatDate(product.createdAt) }}</p>
        <p><strong>Last Updated:</strong> {{ formatDate(product.updatedAt) }}</p>
      </div>
    </div>

    <div v-else class="bg-gray-100 border border-gray-200 text-gray-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <p class="text-lg font-medium mb-2">Product not found.</p>
      <p class="text-sm">The product ID might be invalid or the product does not exist for the active shop.</p>
      <button @click="goBack" class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
        Go to Products List
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'
import {
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PhotographIcon as PlaceholderIcon,
  PencilIcon,
  TrashIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const route = useRoute()
const shopStore = useShopStore()

const product = ref({ images: [], variants: [] })
const loading = ref(true)
const error = ref(null)
const mainImage = ref(null)

const activeShop = computed(() => shopStore.activeShop)

// ADDED: Computed property to determine if variants should be displayed
const displayableVariants = computed(() => {
  const variants = product.value?.variants;
  if (!variants || variants.length === 0) {
    return [];
  }
  // Hide variants section if there is only one variant AND it has no options (e.g. options: {})
  if (variants.length === 1 && Object.keys(variants[0].options).length === 0) {
    return [];
  }
  return variants;
});


function goBack() {
  router.push({ name: 'Products' })
}

function goToEditProduct(productId) {
  router.push({ name: 'EditProduct', params: { productId } })
}

async function confirmDeleteProduct() {
  if (confirm(`Are you sure you want to delete "${product.value.name}"? This action cannot be undone.`)) {
    await deleteProduct();
  }
}

async function deleteProduct() {
  if (!activeShop.value?.id || !product.value?.id) {
    alert('Missing shop or product ID for deletion.')
    return;
  }

  loading.value = true;
  error.value = null;
  try {
    await productService.delete(activeShop.value.id, product.value.id);
    alert('Product deleted successfully!');
    router.push({ name: 'Products' });
  } catch (e) {
    console.error('Failed to delete product:', e);
    error.value = `Failed to delete product: ${e.response?.data?.message || e.message}. Please try again.`;
    alert(error.value);
  } finally {
    loading.value = false;
  }
}

function formatDate(dt) {
  return dt
    ? new Date(dt).toLocaleDateString(undefined, {
        year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit'
      })
    : 'Not available'
}

function formatPrice(p) {
  return p != null ? `$${p.toFixed(2)}` : 'N/A'
}

function formatVariant(options) {
  if (!options || typeof options !== 'object') {
    return ''; // Return an empty string instead of 'N/A' for cleaner display
  }
  const entries = Object.entries(options);
  if (entries.length === 0) {
    return ''; // Also return empty if the options object is empty
  }
  return entries.map(([k, v]) => `${k}: ${v}`).join(', ')
}

function setMainImage(imageUrl) {
  mainImage.value = imageUrl
}

async function fetchProductDetails() {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to view product details.'
    loading.value = false
    return
  }

  const productId = route.params.productId
  if (!productId) {
    error.value = 'Product ID is missing.'
    loading.value = false
    return;
  }

  loading.value = true
  error.value = null
  try {
    const fetchedProduct = await productService.fetchById(activeShop.value.id, productId)
    product.value = fetchedProduct

    // Set the main image: Prioritize main_image, then product images, then first variant image
    if (product.value.main_image) {
        mainImage.value = product.value.main_image;
    } else if (product.value.images && product.value.images.length > 0) {
      mainImage.value = product.value.images[0];
    } else if (product.value.variants && product.value.variants.length > 0 && product.value.variants[0]?.image) {
      mainImage.value = product.value.variants[0].image;
    } else {
      mainImage.value = null; // No images available
    }
  } catch (e) {
    console.error('Failed to load product details:', e)
    error.value = 'Failed to load product details. Please try again later.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProductDetails()
})

watch([() => activeShop.value?.id, () => route.params.productId], () => {
  fetchProductDetails()
})
</script>

<style scoped>
/* Keyframe for a subtle fade-in animation on content load */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fadeIn 0.5s ease-out forwards;
}

/* Keyframe for a subtle pulse on the price */
@keyframes pulsePrice {
  0% { transform: scale(1); }
  50% { transform: scale(1.02); }
  100% { transform: scale(1); }
}
.animate-pulse-price {
  animation: pulsePrice 2s infinite ease-in-out;
}
</style>