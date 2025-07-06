<template>
  <div v-if="product" class="space-y-6">
    <div class="grid md:grid-cols-2 gap-6">
      <div>
        <img
          v-if="currentImage"
          :src="currentImage"
          class="w-full h-80 object-cover rounded"
          alt="Product image"
        />
        <div v-else class="w-full h-80 bg-gray-200 rounded flex items-center justify-center">
            <span class="text-gray-500">No Image</span>
        </div>

        <div class="flex mt-2 space-x-2 overflow-x-auto">
          <img
            v-for="(img, idx) in galleryImages"
            :key="idx"
            :src="img"
            @click="selectThumbnail(img)"
            :class="[
              'w-20 h-20 object-cover rounded cursor-pointer',
              currentImage === img ? 'ring-2 ring-blue-500' : ''
            ]"
            alt="Thumbnail"
          />
        </div>
      </div>

      <div>
        <h1 class="text-3xl font-bold">{{ product?.name }}</h1>
        <p class="text-gray-700 my-2">{{ product?.description }}</p>

        <!-- Discount Information -->
        <div v-if="product.discounts && product.discounts.length > 0" class="mb-4">
          <div class="bg-green-50 border border-green-200 rounded-lg p-3">
            <h3 class="text-sm font-medium text-green-800 mb-2">Active Discounts:</h3>
            <div class="space-y-2">
              <div v-for="discount in product.discounts" :key="discount.id" class="text-sm">
                <span class="font-medium text-green-700">{{ discount.name }}</span>
                <span class="text-green-600 ml-2">
                  {{ discount.type === 'percentage' ? `${discount.value}% off` : `$${discount.value} off` }}
                </span>
                <div v-if="discount.coupon_code" class="text-xs text-green-600 mt-1">
                  Code: {{ discount.coupon_code }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <p class="text-xl font-semibold mb-4">
          <template v-if="selectedVariant">
            ${{ selectedVariant.price.toFixed(2) }}
          </template>
          <template v-else-if="product.price != null && (!product.variants || product.variants.length === 0)">
            ${{ product.price.toFixed(2) }}
          </template>
          <template v-else-if="product.starting_price != null">
            <span class="italic text-gray-700">from</span> ${{ product.starting_price.toFixed(2) }}
          </template>
          <template v-else>
            N/A
          </template>
        </p>

        <div v-if="product.variants && product.variants.length > 0" class="space-y-2 mb-4">
          <p class="font-semibold">Options:</p>
          <div class="flex flex-wrap gap-2">
            <button
              v-for="v in product.variants"
              :key="v.id"
              @click="selectVariant(v)"
              class="px-4 py-2 border rounded hover:bg-gray-100 transition-colors"
              :class="{ 'bg-blue-500 text-white border-blue-500': selectedVariant?.id === v.id }"
            >
              {{ v.options.map(opt => opt.value).join(' / ') }}
            </button>
          </div>
        </div>

        <button
          class="w-full py-2 bg-blue-500 text-white rounded disabled:opacity-50"
          :disabled="!canAddToCart"
        >
          Add to cart
        </button>
      </div>
    </div>
  </div>

  <div v-else class="text-center py-12">Loading product…</div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { fetchProductDetail } from '@/services/product'
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
const shopSlug = route.params.shopSlug as string
const handle = route.params.handle as string

const product = ref<Product | null>(null)
const selectedVariant = ref<Variant | null>(null)
const currentImage = ref<string>('')

onMounted(async () => {
  if (handle) {
    product.value = await fetchProductDetail(shopSlug, handle);
    if (product.value) {
      currentImage.value = product.value.main_image;
    }
  }
})

const galleryImages = computed<string[]>(() => {
  if (!product.value) return [];

  const imgs = new Set<string>();
  if (product.value.main_image) imgs.add(product.value.main_image);
  if (Array.isArray(product.value.images)) {
      product.value.images.forEach(img => img && imgs.add(img));
  }

  // Safely check for variants array before iterating
  if (Array.isArray(product.value.variants)) {
    for (const v of product.value.variants) {
      if (v.image) imgs.add(v.image);
    }
  }
  return Array.from(imgs);
})

function selectThumbnail(img: string) {
  currentImage.value = img;
  selectedVariant.value = null; // Reset variant selection
}

function selectVariant(v: Variant) {
  selectedVariant.value = v;
  currentImage.value = v.image || product.value?.main_image || '';
}

const canAddToCart = computed<boolean>(() => {
  if (!product.value) return false;

  if (selectedVariant.value) {
    return selectedVariant.value.stock > 0;
  }
  
  if (!product.value.variants || product.value.variants.length === 0) {
    return product.value.stock > 0;
  }

  return false; // Disable if variants exist but none are selected
})
</script>