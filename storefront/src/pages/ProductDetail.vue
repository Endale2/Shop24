<template>
  <div class="min-h-screen flex flex-col bg-gray-50 text-gray-900 font-sans antialiased">
    <div class="container mx-auto px-4 py-12 lg:py-16">
      <div v-if="loading" class="flex items-center justify-center h-96 bg-white rounded-xl shadow-lg">
        <svg class="animate-spin h-16 w-16 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
          <path class="opacity-75" fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291 A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
        </svg>
      </div>

      <div v-else-if="error" class="text-center text-red-600 bg-red-50 p-6 rounded-xl shadow-md">
        <p class="text-lg font-semibold">Oops! Something went wrong.</p>
        <p class="text-base">Failed to load product details. Please try again later.</p>
      </div>

      <div v-else-if="product" class="bg-white rounded-xl shadow-xl p-6 lg:p-10 flex flex-col lg:flex-row gap-8 lg:gap-12">
        <div class="lg:w-1/2 flex flex-col">
          <div class="relative w-full rounded-lg overflow-hidden border border-gray-200 aspect-square flex items-center justify-center shadow-md">
            <img
              :src="currentImage"
              :alt="product.name"
              class="w-full h-full object-contain"
              @error="onImageError($event, 500)"
            />
          </div>
          <div v-if="gallery.length > 1" class="mt-4 grid grid-cols-4 gap-3">
            <button
              v-for="(img, i) in gallery"
              :key="i"
              @click="currentImage = img"
              :class="[
                'border-2 rounded-lg overflow-hidden transition-all duration-200 cursor-pointer',
                img === currentImage
                  ? 'border-green-500 ring-2 ring-green-300'
                  : 'border-gray-200 hover:border-green-300'
              ]"
            >
              <img
                :src="img"
                class="w-full h-20 object-cover"
                @error="onImageError($event, 80)"
              />
            </button>
          </div>
        </div>

        <div class="lg:w-1/2 flex flex-col">
          <div>
            <h1 class="text-4xl lg:text-5xl font-extrabold text-gray-900 mb-4 leading-tight">
              {{ product.name }}
            </h1>
            <p class="text-green-600 text-3xl lg:text-4xl font-bold mb-4">
              ${{ priceDisplay.toFixed(2) }}
              <span v-if="product.price > product.display_price" class="text-lg text-gray-500 line-through ml-2">
                ${{ product.price.toFixed(2) }}
              </span>
            </p>
            <p class="text-gray-700 text-base leading-relaxed mb-6">{{ product.description }}</p>

            <div v-if="Object.keys(variantOptions).length > 0" class="space-y-6">
                <div v-for="(values, key) in variantOptions" :key="key">
                <label v-if="Object.keys(variantOptions).length > 1" class="block text-base font-semibold text-gray-800 mb-3">{{ key }}:</label>
                <div class="flex flex-wrap gap-3">
                    <button
                    v-for="val in values" :key="val"
                    @click="selectOption(key, val)"
                    :class="[
                        'px-5 py-2 rounded-full border-2 text-base font-medium transition-all duration-200 shadow-sm',
                        selectedOptions[key] === val
                        ? 'bg-green-600 text-white border-green-600 shadow-md'
                        : isOptionAvailable(key, val)
                            ? 'bg-gray-100 text-gray-700 border-gray-200 hover:bg-gray-200 hover:border-green-300 cursor-pointer'
                            : 'bg-gray-50 text-gray-400 border-gray-100 cursor-not-allowed opacity-50'
                    ]"
                    :disabled="!isOptionAvailable(key, val)"
                    >{{ val }}</button>
                </div>
                </div>
            </div>
            
            <div class="flex items-center space-x-6 my-6">
              <span
                v-if="selectedVariant"
                :class="selectedVariant.stock > 0
                  ? 'text-green-600 bg-green-50 px-3 py-1 rounded-full'
                  : 'text-red-600 bg-red-50 px-3 py-1 rounded-full'"
                class="text-sm font-semibold"
              >
                {{ selectedVariant.stock > 0 ? 'In Stock' : 'Out of Stock' }}
              </span>
              <span v-if="selectedVariant?.id" class="text-sm text-gray-500">
                SKU: <span class="font-mono text-gray-600">{{ selectedVariant.id }}</span>
              </span>
            </div>

            <div class="flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0 sm:space-x-4 mt-6">
              <input
                type="number"
                v-model.number="quantity"
                min="1"
                :max="selectedVariant ? selectedVariant.stock : 1"
                class="w-full sm:w-24 p-3 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-400 focus:border-green-400 transition-all duration-200 text-center"
              />
              <button
                @click="addToCart"
                class="w-full sm:w-auto flex-grow bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-8 rounded-lg shadow-md hover:shadow-lg disabled:opacity-60 disabled:cursor-not-allowed transition-all duration-200"
                :disabled="!selectedVariant || selectedVariant.stock === 0 || quantity < 1"
              >Add to Cart</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { defineProps } from 'vue';
// Assuming your product service is correctly aliased in your build config
import { fetchProductBySlug } from '@/services/product';

const props = defineProps({ shopSlug: String, productSlug: String });

const product = ref(null);
const loading = ref(true);
const error = ref(false);
const gallery = ref([]);
const currentImage = ref('');
const quantity = ref(1);
const selectedOptions = ref({});

// --- DATA FETCHING AND INITIALIZATION ---
async function loadProduct() {
  loading.value = true;
  error.value = false;
  try {
    const data = await fetchProductBySlug(props.shopSlug, props.productSlug);
    product.value = data;
    initializeState(data);
  } catch (e) {
    console.error("Failed to load product:", e);
    error.value = true;
  } finally {
    loading.value = false;
  }
}

function initializeState(data) {
  const imageSet = new Set();
  if (data.main_image) imageSet.add(data.main_image);
  (data.images || []).forEach(img => imageSet.add(img));
  (data.variants || []).forEach(v => {
    if (v.image) imageSet.add(v.image);
  });
  gallery.value = Array.from(imageSet);
  currentImage.value = data.main_image || gallery.value[0] || '';

  if (data.variants && data.variants.length > 0) {
    selectedOptions.value = { ...data.variants[0].options };
  } else {
    selectedOptions.value = {};
  }
}

onMounted(loadProduct);

// --- COMPUTED PROPERTIES ---

const variantOptions = computed(() => {
  if (!product.value) return {}; // Guard against null product
  const optionsMap = {};
  (product.value.variants || []).forEach(variant => {
    for (const [key, value] of Object.entries(variant.options)) {
      if (!optionsMap[key]) {
        optionsMap[key] = new Set();
      }
      optionsMap[key].add(value);
    }
  });
  Object.keys(optionsMap).forEach(key => {
    optionsMap[key] = Array.from(optionsMap[key]);
  });
  return optionsMap;
});

const selectedVariant = computed(() => {
  // FIXED: Add a guard clause to prevent running on a null product.
  if (!product.value) {
    return null;
  }

  const variants = product.value.variants;
  if (!variants || variants.length === 0) {
    // This is a simple product with no variants.
    return {
      id: product.value.id,
      price: product.value.display_price,
      stock: product.value.stock ?? 1,
      options: {},
      image: product.value.main_image
    };
  }

  const match = variants.find(variant => {
    const variantKeys = Object.keys(variant.options);
    const selectedKeys = Object.keys(selectedOptions.value);
    if (variantKeys.length !== selectedKeys.length) return false;
    return variantKeys.every(key => variant.options[key] === selectedOptions.value[key]);
  });
  
  return match; // Returns the matched variant or `undefined`
});

const priceDisplay = computed(() => {
  // Use optional chaining `?.` to safely access properties
  return selectedVariant.value?.price ?? product.value?.display_price ?? 0;
});

// --- CORE LOGIC ---

function isOptionAvailable(keyToTest, valueToTest) {
  if (!product.value?.variants) return false;
  return product.value.variants.some(variant => {
    if (variant.options[keyToTest] !== valueToTest) return false;
    return Object.keys(selectedOptions.value).every(key => {
      if (key === keyToTest) return true;
      return selectedOptions.value[key] === variant.options[key];
    });
  });
}

function selectOption(key, value) {
  selectedOptions.value = { ...selectedOptions.value, [key]: value };
  const compatibleVariant = (product.value?.variants || []).find(variant =>
    Object.entries(selectedOptions.value).every(
      ([k, v]) => variant.options[k] === v
    )
  );
  if (compatibleVariant) {
    selectedOptions.value = { ...compatibleVariant.options };
  }
}

// --- WATCHERS AND UTILITY FUNCTIONS ---

watch(selectedVariant, (newVariant) => {
  if (newVariant?.image) {
    currentImage.value = newVariant.image;
  }
});

function addToCart() {
  if (!selectedVariant.value || !selectedVariant.value.id) return alert('Please select all variant options to continue.');
  if (quantity.value < 1) { quantity.value = 1; return alert('Quantity must be at least 1.'); }
  if (quantity.value > selectedVariant.value.stock) return alert('The selected quantity exceeds available stock.');

  const cartItem = {
    variantId: selectedVariant.value.id,
    productId: product.value.id,
    shopSlug: props.shopSlug,
    name: product.value.name,
    price: priceDisplay.value,
    quantity: quantity.value,
    options: selectedVariant.value.options,
    image: currentImage.value
  };
  
  console.log("Adding to cart:", cartItem);
  alert(`Added ${quantity.value} x ${product.value.name} to cart!`);
}

function onImageError(e, size = 100) { 
  e.target.src = `https://placehold.co/${size}x${size}/E0F2F7/263238?text=No+Image`;
}
</script>

<style scoped>
/* Your component's scoped styles remain unchanged */
</style>