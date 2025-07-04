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

interface Product {
  id: string;
  name: string;
  description: string;
  main_image: string;
  images: string[];
  price?: number;
  starting_price?: number;
  stock: number;
  variants?: Variant[];
}

// --- REAL DATA FETCHING ---
async function fetchProductDetail(handle: string): Promise<Product | null> {
  // NOTE: 'sophiaboutique' is hardcoded based on your example URL.
  // In a real app, you might get this from a store or route params.
  const url = `http://localhost:8080/shops/sophiaboutique/products/${handle}`;

  try {
    const response = await fetch(url);

    if (!response.ok) {
      // Handle HTTP errors like 404 Not Found
      console.error(`Error fetching product: ${response.status} ${response.statusText}`);
      return null;
    }

    const data: Product = await response.json();
    return data;
  } catch (error) {
    // Handle network errors (e.g., server is down)
    console.error("A network error occurred:", error);
    return null;
  }
}

const route = useRoute()
const product = ref<Product | null>(null)
const selectedVariant = ref<Variant | null>(null)
const currentImage = ref<string>('')

onMounted(async () => {
  const handle = route.params.handle as string;
  if (handle) {
    product.value = await fetchProductDetail(handle);
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