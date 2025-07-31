<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <button @click="goBack" class="inline-flex items-center text-gray-600 hover:text-green-700 transition duration-200 ease-in-out mb-6 group">
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500 group-hover:text-green-600 transition-colors duration-200" />
      <span class="text-sm font-medium group-hover:text-green-700 transition-colors duration-200">Back to Product Details</span>
    </button>

    <div v-if="loading" class="flex flex-col items-center justify-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3" />
      <p class="mt-3 text-lg">Loading product for editing...</p>
    </div>

    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg shadow-sm" role="alert">
      <strong class="font-bold">Error:</strong>
      <span class="ml-2">{{ error }}</span>
      <p class="text-sm mt-2">Could not load product. Please try again.</p>
    </div>

    <div v-else-if="productToEdit.id" class="bg-white rounded-xl shadow-lg p-6 md:p-8 space-y-8 animate-fade-in">
      <h2 class="text-3xl font-bold text-gray-800 mb-6">Edit Product: {{ productToEdit.name }}</h2>

      <form @submit.prevent="submitChanges" class="space-y-6">
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Product Name</label>
          <input
            type="text"
            id="name"
            v-model="productToEdit.name"
            class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm"
            required
          />
        </div>

        <div>
          <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
          <textarea
            id="description"
            v-model="productToEdit.description"
            rows="4"
            class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm"
          ></textarea>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <template v-if="productToEdit.variants && productToEdit.variants.length > 0">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Starting Price</label>
              <input type="number" :value="computedPrice" disabled class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 bg-gray-100" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Highest Display Price</label>
              <input type="number" :value="computedDisplayPrice" disabled class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 bg-gray-100" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Total Stock</label>
              <input type="number" :value="computedStock" disabled class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 bg-gray-100" />
            </div>
            <div>
              <label for="main_image" class="block text-sm font-medium text-gray-700 mb-1">Main Image URL</label>
              <input
                type="url"
                id="main_image"
                v-model="productToEdit.main_image"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm"
                placeholder="URL for the main product image (optional)"
              />
            </div>
          </template>
          <template v-else>
            <div>
              <label for="price" class="block text-sm font-medium text-gray-700 mb-1">Product Price</label>
              <input id="price" type="number" v-model.number="productToEdit.price" step="0.01" min="0" class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm" />
            </div>
            <div>
              <label for="stock" class="block text-sm font-medium text-gray-700 mb-1">Stock</label>
              <input id="stock" type="number" v-model.number="productToEdit.stock" min="0" class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm" />
            </div>
            <div>
              <label for="main_image" class="block text-sm font-medium text-gray-700 mb-1">Main Image URL</label>
              <input
                type="url"
                id="main_image"
                v-model="productToEdit.main_image"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm"
                placeholder="URL for the main product image (optional)"
              />
            </div>
          </template>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Collections</label>
          <div class="space-y-3">
            <div v-for="coll in collections" :key="coll.id" class="flex items-center">
              <input
                :id="'collection-' + coll.id"
                type="checkbox"
                :value="coll.id"
                v-model="productToEdit.collection_ids"
                class="h-4 w-4 text-green-600 focus:ring-green-500 border-gray-300 rounded"
              />
              <label :for="'collection-' + coll.id" class="ml-3 text-sm font-medium text-gray-700">
                {{ coll.title }}
              </label>
            </div>
            <div v-if="collections.length === 0" class="text-sm text-gray-500 italic">
              No collections available.
            </div>
          </div>
          <div v-if="productToEdit.collection_ids && productToEdit.collection_ids.length > 0" class="mt-3 space-y-2">
            <div v-for="collectionId in productToEdit.collection_ids" :key="collectionId" class="flex items-center gap-3 p-3 bg-green-50 border border-green-200 rounded-lg shadow-sm animate-fade-in">
              <img v-if="getCollectionById(collectionId)?.image" :src="getCollectionById(collectionId).image" alt="Collection image" class="w-12 h-12 object-cover rounded-md border border-green-200" />
              <div>
                <div class="font-semibold text-green-800 text-lg">{{ getCollectionById(collectionId)?.title }}</div>
                <div class="text-gray-600 text-sm">{{ getCollectionById(collectionId)?.description || 'No description' }}</div>
              </div>
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Additional Images (URLs)</label>
          <div v-for="(image, index) in productToEdit.images" :key="index" class="flex items-center space-x-2 mb-2">
            <input
              type="url"
              v-model="productToEdit.images[index]"
              class="flex-1 border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-green-500 focus:border-green-500 sm:text-sm"
              placeholder="Image URL"
            />
            <button type="button" @click="removeImage(index)" class="p-2 text-red-600 hover:text-red-800 rounded-full hover:bg-red-50">
              <MinusCircleIcon class="h-5 w-5" />
            </button>
          </div>
          <button type="button" @click="addImage" class="mt-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
            <PlusCircleIcon class="h-5 w-5 mr-2" /> Add Image URL
          </button>
        </div>

        <div class="border-t border-gray-200 pt-6 pb-6">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <span class="text-yellow-600 font-semibold text-sm">SEO</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-800">Advanced Fields (SEO)</h3>
            <button type="button" @click="showAdvanced = !showAdvanced" class="ml-4 px-3 py-1.5 rounded bg-gray-100 text-gray-700 text-sm font-medium hover:bg-yellow-50">
              {{ showAdvanced ? 'Hide' : 'Show' }}
            </button>
          </div>
          <div v-if="showAdvanced" class="space-y-4 mt-4">
            <div>
              <label for="meta-title" class="block text-sm font-medium text-gray-700 mb-2">Meta Title</label>
              <input id="meta-title" v-model="productToEdit.meta_title" type="text" class="w-full border border-gray-300 rounded-md px-4 py-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition duration-150 shadow-sm" placeholder="SEO meta title (optional)" />
            </div>
            <div>
              <label for="meta-description" class="block text-sm font-medium text-gray-700 mb-2">Meta Description</label>
              <textarea id="meta-description" v-model="productToEdit.meta_description" rows="2" class="w-full border border-gray-300 rounded-md px-4 py-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition duration-150 shadow-sm resize-y" placeholder="SEO meta description (optional)"></textarea>
            </div>
          </div>
        </div>

        <div class="border-t border-gray-200 pt-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">Product Variants</h3>
          <div v-for="(variant, vIndex) in productToEdit.variants" :key="vIndex" class="bg-gray-50 p-4 rounded-md mb-4 border border-gray-200">
            <div class="flex justify-end">
              <button type="button" @click="removeVariant(vIndex)" class="text-red-600 hover:text-red-800">
                <XIcon class="h-5 w-5" />
              </button>
            </div>
            <div class="space-y-3">
              <label class="block text-sm font-medium text-gray-700">Options</label>
              <div v-for="(opt, oi) in variant.options" :key="oi" class="flex items-center space-x-2 mb-2">
                <input
                  v-model="productToEdit.variants[vIndex].options[oi].name"
                  class="w-1/3 border border-gray-300 rounded-md py-2 px-3 text-sm"
                  placeholder="Option Name (e.g., Color)"
                />
                <input
                  v-model="productToEdit.variants[vIndex].options[oi].value"
                  class="w-1/3 border border-gray-300 rounded-md py-2 px-3 text-sm"
                  placeholder="Option Value (e.g., Red)"
                />
                <button type="button" @click="removeVariantOption(vIndex, oi)" class="p-1 text-red-600 hover:text-red-800">
                  <MinusCircleIcon class="h-4 w-4" />
                </button>
              </div>
              <button type="button" @click="addVariantOption(vIndex)" class="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-800">
                <PlusCircleIcon class="h-4 w-4 mr-1" /> Add Option
              </button>
            </div>
            <div>
              <label :for="`variant-price-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3">Variant Price</label>
              <input
                :id="`variant-price-${vIndex}`"
                type="number"
                v-model.number="productToEdit.variants[vIndex].price"
                step="0.01"
                min="0"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 sm:text-sm"
                required
              />
            </div>

            <div>
              <label :for="`variant-stock-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3">Variant Stock</label>
              <input
                :id="`variant-stock-${vIndex}`"
                type="number"
                v-model.number="productToEdit.variants[vIndex].stock"
                min="0"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 sm:text-sm"
                required
              />
            </div>
            <div>
              <label :for="`variant-image-${vIndex}`" class="block text-sm font-medium text-gray-700 mt-3">Variant Image URL</label>
              <input
                :id="`variant-image-${vIndex}`"
                type="url"
                v-model="productToEdit.variants[vIndex].image"
                class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 sm:text-sm"
                placeholder="URL for variant image (optional)"
              />
            </div>
          </div>
          <button type="button" @click="addVariant" class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
            <PlusIcon class="h-5 w-5 mr-2" /> Add Variant
          </button>
        </div>

        <div v-if="variantOptionError" class="bg-red-50 border border-red-300 text-red-700 px-4 py-2 rounded mb-4">
          <ExclamationCircleIcon class="h-5 w-5 inline-block mr-2" />
          {{ variantOptionError }}
        </div>

        <div class="flex justify-end space-x-4 mt-8">
          <button
            type="button"
            @click="goBack"
            class="px-6 py-3 border border-gray-300 rounded-md shadow-sm text-base font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="saving || !!variantOptionError"
            class="px-6 py-3 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition duration-150 ease-in-out"
          >
            <SpinnerIcon v-if="saving" class="animate-spin h-5 w-5 mr-3 inline-block" />
            {{ saving ? 'Saving...' : 'Save Changes' }}
          </button>
        </div>
      </form>
    </div>

    <div v-else class="bg-gray-100 border border-gray-200 text-gray-700 px-6 py-8 rounded-lg text-center mt-8 shadow-sm">
      <p class="text-lg font-medium mb-2">Product not found for editing.</p>
      <button @click="goBack" class="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
        Go to Product Details
      </button>
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
  ChevronLeftIcon,
  RefreshIcon as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon,
  PlusIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/outline';

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
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-in {
  animation: fadeIn 0.5s ease-out forwards;
}
</style>