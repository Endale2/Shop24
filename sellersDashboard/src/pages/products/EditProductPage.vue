<template>
  <div class="min-h-full bg-gradient-to-br from-gray-50 via-blue-50 to-green-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <button @click="goBack" class="inline-flex items-center text-gray-600 hover:text-green-600 transition-all duration-200 ease-in-out mr-4 group">
              <svg
                class="w-5 h-5 mr-2 transition-all duration-200 group-hover:scale-110 group-hover:-translate-x-0.5"
                fill="none"
                stroke="currentColor"
                stroke-width="2.5"
                viewBox="0 0 24 24"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <path d="M15 18l-6-6 6-6" />
              </svg>
              <span class="text-sm font-medium transition-colors duration-200">Back to Product Details</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="relative">
            <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
            <div class="absolute inset-0 animate-ping h-8 w-8 text-green-400 rounded-full"></div>
          </div>
          <p class="text-sm font-medium">Loading product for editing...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg shadow-sm mb-6 backdrop-blur-sm" role="alert">
        <div class="flex items-center">
          <ExclamationCircleIcon class="h-5 w-5 mr-2 flex-shrink-0 animate-pulse" />
          <div>
            <strong class="font-semibold">Error:</strong>
            <span class="ml-1">{{ error }}</span>
          </div>
        </div>
        <p class="text-sm mt-2">Could not load product. Please try again.</p>
      </div>

      <!-- Edit Form -->
      <div v-else-if="productToEdit.id" class="space-y-6">
        <div class="bg-white/80 backdrop-blur-sm rounded-2xl shadow-xl border border-white/50 p-6 md:p-8 space-y-6 hover:shadow-2xl transition-all duration-500">
          <div class="px-4 py-3 border-b border-gray-100 bg-gradient-to-r from-gray-50 to-blue-50 -mx-6 -mt-6 mb-6 rounded-t-2xl">
            <div class="flex items-center space-x-3">
              <div class="w-8 h-8 bg-gradient-to-br from-green-400 to-blue-500 rounded-lg flex items-center justify-center shadow-lg">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </div>
              <h2 class="text-sm font-semibold text-gray-900">Edit Product: {{ productToEdit.name }}</h2>
            </div>
          </div>

          <form @submit.prevent="submitChanges" class="space-y-6">
            <div class="group">
              <label for="name" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Product Name</label>
              <input
                type="text"
                id="name"
                v-model="productToEdit.name"
                class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300"
                required
              />
            </div>

            <div class="group">
              <label for="description" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Description</label>
              <textarea
                id="description"
                v-model="productToEdit.description"
                rows="4"
                class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300 resize-none"
              ></textarea>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <template v-if="productToEdit.variants && productToEdit.variants.length > 0">
                <div class="group">
                  <label class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Starting Price</label>
                  <input type="number" :value="computedPrice" disabled class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 bg-gradient-to-r from-gray-100 to-gray-50 text-sm" />
                </div>
                <div class="group">
                  <label class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Highest Display Price</label>
                  <input type="number" :value="computedDisplayPrice" disabled class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 bg-gradient-to-r from-gray-100 to-gray-50 text-sm" />
                </div>
                <div class="group">
                  <label class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Total Stock</label>
                  <input type="number" :value="computedStock" disabled class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 bg-gradient-to-r from-gray-100 to-gray-50 text-sm" />
                </div>
                <div class="group">
                  <label for="main_image" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Main Image URL</label>
                  <input
                    type="url"
                    id="main_image"
                    v-model="productToEdit.main_image"
                    class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300"
                    placeholder="URL for the main product image (optional)"
                  />
                </div>
              </template>
              <template v-else>
                <div class="group">
                  <label for="price" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Product Price</label>
                  <input id="price" type="number" v-model.number="productToEdit.price" step="0.01" min="0" class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300" />
                </div>
                <div class="group">
                  <label for="stock" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Stock</label>
                  <input id="stock" type="number" v-model.number="productToEdit.stock" min="0" class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300" />
                </div>
                <div class="group">
                  <label for="main_image" class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Main Image URL</label>
                  <input
                    type="url"
                    id="main_image"
                    v-model="productToEdit.main_image"
                    class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300"
                    placeholder="URL for the main product image (optional)"
                  />
                </div>
              </template>
            </div>

            <div class="group">
              <label class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Collections</label>
              <div class="space-y-3">
                <div v-for="coll in collections" :key="coll.id" class="flex items-center hover:bg-gray-50 p-2 rounded-lg transition-colors duration-200">
                  <input
                    :id="'collection-' + coll.id"
                    type="checkbox"
                    :value="coll.id"
                    v-model="productToEdit.collection_ids"
                    class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-300 rounded transition-all duration-200 hover:scale-110"
                  />
                  <label :for="'collection-' + coll.id" class="ml-3 text-sm font-medium text-gray-700 cursor-pointer hover:text-green-600 transition-colors duration-200">
                    {{ coll.title }}
                  </label>
                </div>
                <div v-if="collections.length === 0" class="text-sm text-gray-500 italic">
                  No collections available.
                </div>
              </div>
              <div v-if="productToEdit.collection_ids && productToEdit.collection_ids.length > 0" class="mt-3 space-y-2">
                <div v-for="collectionId in productToEdit.collection_ids" :key="collectionId" class="flex items-center gap-3 p-3 bg-gradient-to-r from-green-50 to-blue-50 border border-green-200 rounded-xl shadow-sm hover:shadow-md transition-all duration-200 hover:scale-105">
                  <img v-if="getCollectionById(collectionId)?.image" :src="getCollectionById(collectionId).image" alt="Collection image" class="w-10 h-10 object-cover rounded-lg border border-green-200 shadow-sm" />
                  <div>
                    <div class="font-semibold text-green-800 text-sm">{{ getCollectionById(collectionId)?.title }}</div>
                    <div class="text-gray-600 text-xs">{{ getCollectionById(collectionId)?.description || 'No description' }}</div>
                  </div>
                </div>
              </div>
            </div>

            <div class="group">
              <label class="block text-sm font-medium text-gray-700 mb-1 group-hover:text-green-600 transition-colors duration-200">Additional Images (URLs)</label>
              <div v-for="(image, index) in productToEdit.images" :key="index" class="flex items-center space-x-2 mb-2">
                <input
                  type="url"
                  v-model="productToEdit.images[index]"
                  class="flex-1 border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm transition-all duration-200 hover:shadow-md group-hover:border-green-300"
                  placeholder="Image URL"
                />
                <button type="button" @click="removeImage(index)" class="p-2 text-red-600 hover:text-red-800 rounded-full hover:bg-red-50 transition-all duration-200 hover:scale-110">
                  <MinusCircleIcon class="h-4 w-4" />
                </button>
              </div>
              <button type="button" @click="addImage" class="mt-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-xl shadow-sm text-white bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200 hover:scale-105">
                <PlusCircleIcon class="h-4 w-4 mr-1.5" /> Add Image URL
              </button>
            </div>

            <div class="border-t border-gray-200 pt-6 pb-6">
              <div class="flex items-center space-x-3">
                <div class="w-6 h-6 bg-gradient-to-br from-yellow-400 to-orange-500 rounded-full flex items-center justify-center shadow-lg">
                  <span class="text-white font-semibold text-xs">SEO</span>
                </div>
                <h3 class="text-sm font-semibold text-gray-800">Advanced Fields (SEO)</h3>
                <button type="button" @click="showAdvanced = !showAdvanced" class="ml-4 px-3 py-1.5 rounded-xl bg-gradient-to-r from-gray-100 to-gray-200 text-gray-700 text-xs font-medium hover:from-yellow-50 hover:to-orange-50 transition-all duration-200 hover:scale-105">
                  {{ showAdvanced ? 'Hide' : 'Show' }}
                </button>
              </div>
              <div v-if="showAdvanced" class="space-y-4 mt-4 animate-fade-in">
                <div class="group">
                  <label for="meta-title" class="block text-sm font-medium text-gray-700 mb-2 group-hover:text-yellow-600 transition-colors duration-200">Meta Title</label>
                  <input id="meta-title" v-model="productToEdit.meta_title" type="text" class="w-full border border-gray-300 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all duration-200 shadow-sm text-sm hover:shadow-md group-hover:border-yellow-300" placeholder="SEO meta title (optional)" />
                </div>
                <div class="group">
                  <label for="meta-description" class="block text-sm font-medium text-gray-700 mb-2 group-hover:text-yellow-600 transition-colors duration-200">Meta Description</label>
                  <textarea id="meta-description" v-model="productToEdit.meta_description" rows="2" class="w-full border border-gray-300 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all duration-200 shadow-sm resize-none text-sm hover:shadow-md group-hover:border-yellow-300" placeholder="SEO meta description (optional)"></textarea>
                </div>
              </div>
            </div>

            <div class="border-t border-gray-200 pt-6">
              <h3 class="text-sm font-semibold text-gray-800 mb-4 flex items-center">
                <div class="w-5 h-5 bg-gradient-to-br from-purple-400 to-pink-500 rounded-full flex items-center justify-center mr-2 shadow-lg">
                  <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                  </svg>
                </div>
                Product Variants
              </h3>
              <div v-for="(variant, vIndex) in productToEdit.variants" :key="vIndex" class="bg-gradient-to-r from-gray-50 to-purple-50 p-4 rounded-xl mb-4 border border-gray-200 hover:shadow-md transition-all duration-200 hover:scale-105">
                <div class="flex justify-end">
                  <button type="button" @click="removeVariant(vIndex)" class="text-red-600 hover:text-red-800 transition-all duration-200 hover:scale-110">
                    <XIcon class="h-4 w-4" />
                  </button>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-medium text-gray-700">Options</label>
                  <div v-for="(opt, oi) in variant.options" :key="oi" class="flex items-center space-x-2 mb-2">
                    <input
                      v-model="productToEdit.variants[vIndex].options[oi].name"
                      class="w-1/3 border border-gray-300 rounded-xl py-2 px-3 text-sm transition-all duration-200 hover:shadow-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="Option Name (e.g., Color)"
                    />
                    <input
                      v-model="productToEdit.variants[vIndex].options[oi].value"
                      class="w-1/3 border border-gray-300 rounded-xl py-2 px-3 text-sm transition-all duration-200 hover:shadow-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="Option Value (e.g., Red)"
                    />
                    <button type="button" @click="removeVariantOption(vIndex, oi)" class="p-1 text-red-600 hover:text-red-800 transition-all duration-200 hover:scale-110">
                      <MinusCircleIcon class="h-3 w-3" />
                    </button>
                  </div>
                  <button type="button" @click="addVariantOption(vIndex)" class="mt-2 inline-flex items-center text-sm font-medium text-purple-600 hover:text-purple-800 transition-all duration-200 hover:scale-105">
                    <PlusCircleIcon class="h-3 w-3 mr-1" /> Add Option
                  </button>
                </div>
                <div class="group">
                  <label :for="`variant-price-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3 group-hover:text-purple-600 transition-colors duration-200">Variant Price</label>
                  <input
                    :id="`variant-price-${vIndex}`"
                    type="number"
                    v-model.number="productToEdit.variants[vIndex].price"
                    step="0.01"
                    min="0"
                    class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 text-sm transition-all duration-200 hover:shadow-md group-hover:border-purple-300 focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    required
                  />
                </div>

                <div class="group">
                  <label :for="`variant-stock-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3 group-hover:text-purple-600 transition-colors duration-200">Variant Stock</label>
                  <input
                    :id="`variant-stock-${vIndex}`"
                    type="number"
                    v-model.number="productToEdit.variants[vIndex].stock"
                    min="0"
                    class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 text-sm transition-all duration-200 hover:shadow-md group-hover:border-purple-300 focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    required
                  />
                </div>
                <div class="group">
                  <label :for="`variant-image-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3 group-hover:text-purple-600 transition-colors duration-200">Variant Image URL</label>
                  <input
                    :id="`variant-image-${vIndex}`"
                    type="url"
                    v-model="productToEdit.variants[vIndex].image"
                    class="mt-1 block w-full border border-gray-300 rounded-xl shadow-sm py-2.5 px-3 text-sm transition-all duration-200 hover:shadow-md group-hover:border-purple-300 focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    placeholder="URL for variant image (optional)"
                  />
                </div>
              </div>
              <button type="button" @click="addVariant" class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-xl shadow-sm text-white bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 transition-all duration-200 hover:scale-105">
                <PlusIcon class="h-4 w-4 mr-1.5" /> Add Variant
              </button>
            </div>

            <div v-if="variantOptionError" class="bg-red-50 border border-red-300 text-red-700 px-4 py-2 rounded-xl mb-4 backdrop-blur-sm animate-pulse">
              <ExclamationCircleIcon class="h-4 w-4 inline-block mr-2" />
              {{ variantOptionError }}
            </div>

            <div class="flex justify-end space-x-3 mt-8">
              <button
                type="button"
                @click="goBack"
                class="px-4 py-2.5 border border-gray-300 rounded-xl shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-all duration-200 hover:scale-105"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="saving || !!variantOptionError"
                class="px-4 py-2.5 border border-transparent rounded-xl shadow-sm text-sm font-medium text-white bg-gradient-to-r from-green-600 to-blue-600 hover:from-green-700 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-all duration-200 hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <SpinnerIcon v-if="saving" class="animate-spin h-4 w-4 mr-1.5 inline-block" />
                {{ saving ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Not Found State -->
      <div v-else class="text-center py-16 text-gray-400">
        <div class="max-w-md mx-auto">
          <div class="w-12 h-12 bg-gradient-to-br from-gray-100 to-gray-200 rounded-full flex items-center justify-center mx-auto mb-3 shadow-lg">
            <ExclamationCircleIcon class="w-6 h-6 text-gray-300" />
          </div>
          <p class="text-sm font-medium text-gray-600 mb-1">Product not found for editing</p>
          <p class="text-xs">The product ID might be invalid or the product does not exist for the active shop.</p>
          <button @click="goBack" class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-xl shadow-sm text-white bg-gradient-to-r from-green-600 to-blue-600 hover:from-green-700 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-all duration-200 hover:scale-105">
            Go to Product Details
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useShopStore } from '@/store/shops';
import { productService } from '@/services/product';
import { collectionService } from '@/services/collection'
import {
  RefreshIcon as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon,
  PlusIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline'
import BackButton from '@/components/BackButton.vue';

const router = useRouter();
const route = useRoute();
const shopStore = useShopStore();

const productToEdit = ref({
  id: '',
  name: '',
  description: '',
  price: null,
  stock: null,
  category: '',
  images: [],
  main_image: null,
  variants: [],
  meta_title: '',
  meta_description: '',
});
const loading = ref(true);
const saving = ref(false);
const error = ref(null);
const variantOptionError = ref(null);
const showAdvanced = ref(false);
const collections = ref([])

const activeShop = computed(() => shopStore.activeShop);
const getCollectionById = (id) => collections.value.find(c => c.id === id);

const computedPrice = computed(() => {
  if (!productToEdit.value.variants || productToEdit.value.variants.length === 0) return productToEdit.value.price;
  return Math.min(...productToEdit.value.variants.map(v => v.price));
});
const computedDisplayPrice = computed(() => {
  if (!productToEdit.value.variants || productToEdit.value.variants.length === 0) return productToEdit.value.display_price;
  const displayPrices = productToEdit.value.variants.map(v => v.display_price).filter(p => p != null);
  return displayPrices.length ? Math.max(...displayPrices) : '';
});
const computedStock = computed(() => {
  if (!productToEdit.value.variants || productToEdit.value.variants.length === 0) return productToEdit.value.stock;
  return productToEdit.value.variants.reduce((sum, v) => sum + (v.stock || 0), 0);
});

function goBack() {
  const productId = route.params.productId;
  if (productId) {
    router.push({ name: 'ProductDetail', params: { productId } });
  } else {
    router.push({ name: 'Products' });
  }
}

function addImage() {
  productToEdit.value.images.push('');
}

function removeImage(index) {
  productToEdit.value.images.splice(index, 1);
}

function addVariant() {
  productToEdit.value.variants.push({ options: [{ name: '', value: '' }], price: 0, stock: 0, image: null });
}

function removeVariant(index) {
  productToEdit.value.variants.splice(index, 1);
}

function addVariantOption(vi) {
  productToEdit.value.variants[vi].options.push({ name: '', value: '' });
}

function removeVariantOption(vi, oi) {
  productToEdit.value.variants[vi].options.splice(oi, 1);
}

async function fetchProductDetails() {
  if (!activeShop.value) {
    error.value = 'No shop selected. Please select a shop to edit product details.';
    loading.value = false;
    return;
  }

  const productId = route.params.productId;
  if (!productId) {
    error.value = 'Product ID is missing for editing.';
    loading.value = false;
    return;
  }

  loading.value = true;
  error.value = null;
  try {
    const fetchedProduct = await productService.fetchById(activeShop.value.id, productId);
    // Map backend fields to frontend
    productToEdit.value = {
      id: fetchedProduct.id || fetchedProduct._id,
      name: fetchedProduct.name,
      slug: fetchedProduct.slug,
      images: fetchedProduct.images || [],
      main_image: fetchedProduct.main_image,
      category: fetchedProduct.category,
      collection_ids: fetchedProduct.collection_ids || [],
      price: (typeof fetchedProduct.price === 'number' && !isNaN(fetchedProduct.price)) ? fetchedProduct.price : 0,
      stock: (typeof fetchedProduct.stock === 'number' && !isNaN(fetchedProduct.stock)) ? fetchedProduct.stock : 0,
      description: fetchedProduct.description,
      variants: (fetchedProduct.variants || []).map(v => ({
        ...v,
        options: v.option || v.options || {},
        price: v.price,
        stock: v.stock,
        image: v.image
      })),
      meta_title: fetchedProduct.meta_title || '',
      meta_description: fetchedProduct.meta_description || '',
    };
  } catch (e) {
    console.error('Failed to load product details for editing:', e);
    error.value = 'Failed to load product details for editing. Please try again later.';
  } finally {
    loading.value = false;
  }
}

async function submitChanges() {
  if (!activeShop.value?.id || !productToEdit.value?.id) {
    alert('Missing shop or product ID for saving changes.');
    return;
  }

  saving.value = true;
  error.value = null;
  variantOptionError.value = null;

  let dataToSend = {
    name: productToEdit.value.name,
    description: productToEdit.value.description,
    collection_ids: productToEdit.value.collection_ids,
    images: productToEdit.value.images.filter(img => img),
    main_image: productToEdit.value.main_image || null,
    meta_title: productToEdit.value.meta_title,
    meta_description: productToEdit.value.meta_description,
  };

  if (productToEdit.value.variants.length > 0) {
    dataToSend.variants = productToEdit.value.variants.map(v => ({
      options: v.options.filter(o => o.name && o.value),
      price: v.price,
      stock: v.stock,
      image: v.image || null,
    }));
    // Do NOT send price/stock for variant products
  } else {
    // Only include price/stock if they are not null/undefined
    if (typeof productToEdit.value.price === 'number') {
      dataToSend.price = productToEdit.value.price;
    }
    if (typeof productToEdit.value.stock === 'number') {
      dataToSend.stock = productToEdit.value.stock;
    }
    // Do NOT include variants at all for simple products
    // Always send main_image as a string (even if empty)
    dataToSend.main_image = productToEdit.value.main_image || '';
  }

  try {
    // Debug: log the payload being sent
    console.log('PATCH payload:', JSON.stringify(dataToSend, null, 2));
    await productService.patch(activeShop.value.id, productToEdit.value.id, dataToSend);
    alert('Product updated successfully!');
    router.push({ name: 'ProductDetail', params: { productId: productToEdit.value.id } });
  } catch (e) {
    console.error('Failed to update product:', e);
    error.value = `Failed to update product: ${e.response?.data?.message || e.message}. Please try again.`;
    alert(error.value);
  } finally {
    saving.value = false;
  }
}

onMounted(async () => {
  collections.value = await collectionService.fetchAllByShop(shopStore.activeShop.id)
  fetchProductDetails();
});

watch([() => activeShop.value?.id, () => route.params.productId], () => {
  fetchProductDetails();
});
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Cool animations */
@keyframes fade-in {
  from { 
    opacity: 0; 
    transform: translateY(10px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

.animate-fade-in {
  animation: fade-in 0.5s ease-out forwards;
}

/* Gradient text animation */
@keyframes gradient-shift {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

/* Floating animation for subtle movement */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
}

/* Pulse animation for loading states */
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

/* Glass morphism effect */
.backdrop-blur-sm {
  backdrop-filter: blur(8px);
}

/* Enhanced shadows */
.shadow-xl {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

.shadow-2xl {
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}
</style>