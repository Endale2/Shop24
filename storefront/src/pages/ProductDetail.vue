<template>
  <div v-if="product" class="max-w-7xl mx-auto p-4 sm:p-6 lg:p-8">
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <div class="flex flex-row lg:flex-col gap-4 overflow-x-auto lg:overflow-x-visible">
        <div
          v-for="(img, idx) in galleryImages"
          :key="idx"
          @click="selectThumbnail(img)"
          :class="[
            'w-24 h-24 object-contain rounded-md border cursor-pointer flex-shrink-0',
            currentImage === img ? 'border-black ring-2 ring-black' : 'border-gray-200'
          ]"
          :style="{ background: '#f9fafb' }"
        >
          <img :src="img" :alt="`Thumbnail ${idx + 1}`" class="w-full h-full object-contain" />
        </div>
      </div>

      <div class="lg:col-span-1 flex items-start justify-center">
        <div class="w-full">
          <img
            v-if="currentImage"
            :src="currentImage"
            class="w-full h-auto max-h-[600px] object-contain rounded-lg bg-gray-50 border border-gray-200"
            alt="Product image"
          />
          <div v-else class="w-full h-96 bg-gray-200 rounded-lg flex items-center justify-center">
            <span class="text-gray-400">No Image</span>
          </div>
        </div>
      </div>

      <div class="lg:col-span-1 flex flex-col gap-4">
        <h1 class="text-3xl font-bold tracking-wide uppercase text-gray-900">{{ product?.name }}</h1>
        <p class="text-sm text-gray-500">{{ product.category }}</p>
        <p class="text-gray-600 leading-relaxed">{{ product.description }}</p>
        
        <!-- Discount Badge -->
        <div v-if="product.discounts && product.discounts.length > 0" class="my-2">
          <span class="inline-block px-3 py-1 bg-yellow-100 text-yellow-800 text-sm font-semibold rounded-full">
            {{ product.discounts[0].type === 'percentage' ? `${product.discounts[0].value}% OFF` : `$${product.discounts[0].value} OFF` }}
          </span>
        </div>

        <!-- Price Display -->
        <div class="space-y-2">
          <p class="text-3xl font-bold text-gray-900">
            <template v-if="selectedVariant">
              <span v-if="selectedVariant.discount_amount > 0" class="text-lg text-gray-400 line-through mr-2">
                ${{ (selectedVariant.price * quantity).toFixed(2) }}
              </span>
              ${{ currentPrice.toFixed(2) }}
            </template>
            <template v-else-if="product.price != null && (!product.variants || product.variants.length === 0)">
              <span v-if="product.discounts && product.discounts.length > 0" class="text-lg text-gray-400 line-through mr-2">
                ${{ (product.price * quantity).toFixed(2) }}
              </span>
              ${{ currentPrice.toFixed(2) }}
            </template>
            <template v-else-if="product.starting_price != null">
              <span class="text-xl italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
            </template>
            <template v-else>
              N/A
            </template>
          </p>
          
          <!-- Savings info -->
          <div v-if="hasDiscount && savings > 0" class="text-sm text-green-600">
            Save ${{ savings.toFixed(2) }}
          </div>
        </div>

        <div v-if="product.variants && product.variants.length > 0" class="space-y-3">
          <p class="font-semibold text-sm uppercase text-gray-700">Options:</p>
          <div class="flex flex-wrap gap-2">
            <button
              v-for="v in product.variants"
              :key="v.id"
              @click="selectVariant(v)"
              class="px-5 py-2 border-2 rounded-lg font-medium transition-colors focus:outline-none text-xs uppercase"
              :class="selectedVariant?.id === v.id ? 'bg-black text-white border-black' : 'bg-white text-gray-800 border-gray-300 hover:bg-gray-100'"
            >
              {{ v.options.map(opt => opt.value).join(' / ') }}
            </button>
          </div>
        </div>

        <div class="mt-2 text-sm font-medium text-gray-900">
          <template v-if="selectedVariant">
            <span v-if="selectedVariant.stock === 0" class="text-red-600 font-bold">Out of stock</span>
            <span v-else>In Stock: <span class="text-green-600">{{ selectedVariant.stock }}</span></span>
          </template>
          <template v-else-if="product.total_stock !== undefined">
            <span v-if="product.total_stock === 0" class="text-red-600 font-bold">Out of stock</span>
            <span v-else>Total Stock: <span class="text-green-600">{{ product.total_stock }}</span></span>
          </template>
          <template v-else-if="product.stock !== undefined">
            <span v-if="product.stock === 0" class="text-red-600 font-bold">Out of stock</span>
            <span v-else>In Stock: <span class="text-green-600">{{ product.stock }}</span></span>
          </template>
        </div>

        <!-- Quantity Selector -->
        <div v-if="canAddToCart" class="space-y-3">
          <div class="flex items-center gap-4">
            <label class="font-semibold text-sm uppercase text-gray-700">Quantity:</label>
            <div class="flex items-center border border-gray-300 rounded-lg">
              <button
                @click="decreaseQuantity"
                :disabled="quantity <= 1"
                class="px-3 py-2 text-gray-600 hover:text-black disabled:text-gray-300 disabled:cursor-not-allowed"
              >
                -
              </button>
              <input
                v-model.number="quantity"
                type="number"
                min="1"
                :max="maxQuantity"
                class="w-16 text-center border-none focus:outline-none"
                @input="validateQuantity"
                @blur="validateQuantity"
              />
              <button
                @click="increaseQuantity"
                :disabled="quantity >= maxQuantity"
                class="px-3 py-2 text-gray-600 hover:text-black disabled:text-gray-300 disabled:cursor-not-allowed"
              >
                +
              </button>
            </div>
            <span class="text-sm text-gray-500">of {{ maxQuantity }} available</span>
          </div>
        </div>

        <!-- Add to Cart Section -->
        <div class="space-y-3">
          <button
            v-if="!authStore.user"
            @click="goToLogin"
            class="w-full py-3 bg-black text-white rounded text-base font-normal transition-transform transform hover:scale-105"
          >
            Login to Add to Cart
          </button>
          <button
            v-else
            @click="addToCart"
            :disabled="!canAddToCart || cartStore.loading"
            class="w-full py-3 bg-black text-white rounded text-base font-normal transition-transform transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
          >
            <span v-if="cartStore.loading" class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
            {{ cartStore.loading ? 'Adding...' : 'Add to Cart' }}
          </button>
          
          <div v-if="cartStore.error" class="text-red-600 text-sm text-center">
            {{ cartStore.error }}
          </div>
          
          <div v-if="addToCartSuccess" class="text-green-600 text-sm text-center">
            Added to cart successfully!
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="text-center py-20 text-gray-500 text-xl">Loading product…</div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchProductDetail } from '@/services/product'
import { useCartStore } from '@/stores/cart'
import { useAuthStore } from '@/stores/auth'
import type { Product } from '@/services/product'

// Define interfaces to match your data structure for better type safety
interface ProductOption {
  name: string;
  value: string;
}

interface Variant {
  id: string;
  options: ProductOption[];
  price: number;
  stock: number;
  image?: string;
}

const route = useRoute()
const router = useRouter()
const shopSlug = route.params.shopSlug as string
const handle = route.params.handle as string

const product = ref<Product | null>(null)
const selectedVariant = ref<Variant | null>(null)
const currentImage = ref<string>('')
const quantity = ref(1)
const addToCartSuccess = ref(false)

const cartStore = useCartStore()
const authStore = useAuthStore()

onMounted(async () => {
  if (handle) {
    product.value = await fetchProductDetail(shopSlug, handle);
    if (product.value) {
      currentImage.value = product.value.main_image;
    }
  }
  
  // Set shop slug in cart store
  cartStore.setShopSlug(shopSlug);
})

const galleryImages = computed<string[]>(() => {
  if (!product.value) return [];
  const imgs = new Set<string>();
  if (product.value.main_image) imgs.add(product.value.main_image);
  if (Array.isArray(product.value.images)) {
    product.value.images.forEach(img => img && imgs.add(img));
  }
  if (Array.isArray(product.value.variants)) {
    for (const v of product.value.variants) {
      if (v.image) imgs.add(v.image);
    }
  }
  return Array.from(imgs);
})

const maxQuantity = computed<number>(() => {
  if (!product.value) return 1;
  if (selectedVariant.value) {
    return selectedVariant.value.stock;
  }
  if (!product.value.variants || product.value.variants.length === 0) {
    return product.value.stock ?? product.value.total_stock ?? 1;
  }
  return 1;
})

function selectThumbnail(img: string) {
  currentImage.value = img;
  // Optional: Find a variant that matches this image and select it
  const matchingVariant = product.value?.variants.find(v => v.image === img);
  if (matchingVariant) {
    selectedVariant.value = matchingVariant;
    validateQuantity();
  }
}

function selectVariant(v: Variant) {
  selectedVariant.value = v;
  currentImage.value = v.image || product.value?.main_image || '';
  validateQuantity();
}

function validateQuantity() {
  if (quantity.value < 1) {
    quantity.value = 1;
  }
  if (quantity.value > maxQuantity.value) {
    quantity.value = maxQuantity.value;
  }
}

function increaseQuantity() {
  if (quantity.value < maxQuantity.value) {
    quantity.value++;
  }
}

function decreaseQuantity() {
  if (quantity.value > 1) {
    quantity.value--;
  }
}

const canAddToCart = computed<boolean>(() => {
  if (!product.value) return false;
  if (selectedVariant.value) {
    return selectedVariant.value.stock > 0;
  }
  if (!product.value.variants || product.value.variants.length === 0) {
    // Check stock for products without variants
    return (product.value.stock ?? product.value.total_stock ?? 0) > 0;
  }
  return false; // Disable if variants exist but none are selected
})

const hasDiscount = computed<boolean>(() => {
  if (!product.value) return false;
  if (selectedVariant.value && selectedVariant.value.discount_amount > 0) {
    return true;
  }
  if (product.value.discounts && product.value.discounts.length > 0) {
    return true;
  }
  return false;
})

const savings = computed(() => {
  if (!product.value) return 0;
  let price = 0;
  let discount = 0;
  let type = '';
  if (selectedVariant.value) {
    price = selectedVariant.value.price;
    discount = selectedVariant.value.discount_amount || 0;
    type = selectedVariant.value.discount_type || (product.value.discounts && product.value.discounts[0]?.type) || '';
    if (type === 'percentage') {
      return ((price * (discount / 100)) * quantity.value);
    } else {
      return (discount * quantity.value);
    }
  } else if (product.value.discounts && product.value.discounts.length > 0) {
    price = product.value.price;
    discount = product.value.discounts[0].value;
    type = product.value.discounts[0].type;
    if (type === 'percentage') {
      return ((price * (discount / 100)) * quantity.value);
    } else {
      return (discount * quantity.value);
    }
  }
  return 0;
});

const currentPrice = computed(() => {
  if (!product.value) return 0;
  let price = 0;
  let discount = 0;
  let type = '';
  if (selectedVariant.value) {
    price = selectedVariant.value.price;
    discount = selectedVariant.value.discount_amount || 0;
    type = selectedVariant.value.discount_type || (product.value.discounts && product.value.discounts[0]?.type) || '';
    if (type === 'percentage') {
      return (price * (1 - discount / 100)) * quantity.value;
    } else {
      return (price - discount) * quantity.value;
    }
  } else if (product.value.discounts && product.value.discounts.length > 0) {
    price = product.value.price;
    discount = product.value.discounts[0].value;
    type = product.value.discounts[0].type;
    if (type === 'percentage') {
      return (price * (1 - discount / 100)) * quantity.value;
    } else {
      return (price - discount) * quantity.value;
    }
  } else if (product.value.price) {
    return product.value.price * quantity.value;
  }
  return 0;
});

async function addToCart() {
  if (!product.value || !canAddToCart.value) return;
  
  addToCartSuccess.value = false;
  
  try {
    const productId = product.value.id;
    const variantId = selectedVariant.value?.id || '';
    
    await cartStore.addToCart(productId, variantId, quantity.value);
    
    addToCartSuccess.value = true;
    
    // Reset success message after 3 seconds
    setTimeout(() => {
      addToCartSuccess.value = false;
    }, 3000);
    
  } catch (error) {
    console.error('Failed to add to cart:', error);
    // Don't show success message if there was an error
    addToCartSuccess.value = false;
  }
}

function goToLogin() {
  router.push({ name: 'Login', params: { shopSlug } });
}
</script>