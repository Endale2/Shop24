<!-- src/pages/products/AddProductPage.vue -->
<template>
  <div class="max-w-3xl mx-auto py-12 px-4 sm:px-6 lg:px-8 font-sans">
    <h1 class="text-3xl font-extrabold mb-8">Add New Product</h1>

    <form @submit.prevent="submitForm" class="space-y-6 bg-white p-6 rounded-xl shadow">
      <!-- Name -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Name</label>
        <input
          v-model="form.name"
          type="text"
          required
          class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
      </div>

      <!-- Description -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Description</label>
        <textarea
          v-model="form.description"
          required
          class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
      </div>

      <!-- Images -->
      <div>
        <label class="block text-sm font-medium text-gray-700">
          Images (comma‑separated URLs)
        </label>
        <input
          v-model="form.imagesInput"
          type="text"
          placeholder="https://…, https://…"
          class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
      </div>

      <!-- Category -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Category</label>
        <input
          v-model="form.category"
          type="text"
          required
          class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
      </div>

      <!-- Base Price -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Base Price</label>
        <input
          v-model.number="form.price"
          type="number"
          min="0"
          step="0.01"
          required
          class="mt-1 block w-32 border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
      </div>

      <!-- Variants Section -->
      <div>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Variants</h2>
          <button
            type="button"
            @click="addVariant()"
            class="inline-flex items-center px-3 py-1 bg-green-600 text-white rounded hover:bg-green-700"
          >
            + Add Variant
          </button>
        </div>

        <div v-for="(variant, vIdx) in form.variants" :key="vIdx" class="mt-4 p-4 border rounded-lg space-y-4">
          <div class="flex justify-between items-center">
            <h3 class="font-medium">Variant {{ vIdx + 1 }}</h3>
            <button
              type="button"
              @click="removeVariant(vIdx)"
              class="text-red-600 hover:underline text-sm"
            >Remove</button>
          </div>

          <!-- Variant Name -->
          <div>
            <label class="block text-sm font-medium text-gray-700">Name (e.g. Color)</label>
            <input
              v-model="variant.name"
              type="text"
              required
              class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
            />
          </div>

          <!-- Options + Prices -->
          <div v-for="(opt, oIdx) in variant.options" :key="oIdx" class="flex space-x-2 items-end">
            <div class="flex-1">
              <label class="block text-sm font-medium text-gray-700">Option</label>
              <input
                v-model="opt.value"
                type="text"
                required
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
              />
            </div>
            <div class="w-32">
              <label class="block text-sm font-medium text-gray-700">Price</label>
              <input
                v-model.number="opt.price"
                type="number"
                min="0"
                step="0.01"
                required
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
              />
            </div>
            <button
              type="button"
              @click="removeOption(vIdx, oIdx)"
              class="text-red-600 hover:underline text-sm"
            >✕</button>
          </div>
          <button
            type="button"
            @click="addOption(vIdx)"
            class="mt-2 inline-flex items-center px-3 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200"
          >
            + Add Option
          </button>
        </div>
      </div>

      <!-- Submit -->
      <div class="pt-4 border-t flex justify-end">
        <button
          type="submit"
          :disabled="submitting"
          class="inline-flex items-center px-6 py-2 bg-green-600 text-white rounded shadow hover:bg-green-700 disabled:opacity-50"
        >
          {{ submitting ? 'Saving...' : 'Save Product' }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { productService } from '@/services/product'

// Refs & state
const router = useRouter()
const shopStore = useShopStore()
const form = ref({
  name: '',
  description: '',
  imagesInput: '',
  category: '',
  price: 0,
  variants: []  // each { name: '', options: [ { value:'', price:0 } ] }
})
const submitting = ref(false)

// Helpers to manage variants/options
function addVariant() {
  form.value.variants.push({
    name: '',
    options: [ { value: '', price: 0 } ]
  })
}
function removeVariant(idx) {
  form.value.variants.splice(idx, 1)
}
function addOption(variantIdx) {
  form.value.variants[variantIdx].options.push({ value: '', price: 0 })
}
function removeOption(variantIdx, optIdx) {
  form.value.variants[variantIdx].options.splice(optIdx, 1)
}

// Form submission
async function submitForm() {
  submitting.value = true
  try {
    // Prepare payload
    const payload = {
      name: form.value.name,
      description: form.value.description,
      images: form.value.imagesInput
        .split(',')
        .map(u => u.trim())
        .filter(u => u),
      category: form.value.category,
      price: form.value.price,
      createdBy: shopStore.activeShop.userId, // or however you pass user
      variants: form.value.variants.map(v => ({
        name: v.name,
        options: v.options.map(o => o.value),
        prices: v.options.map(o => o.price)
      }))
    }

    // Call service
    await productService.create(shopStore.activeShop.id, payload)

    // Redirect back to list
    router.push({ name: 'Products' })
  } catch (e) {
    console.error('Error creating product', e)
    alert('Failed to create product.')
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
/* Add any tweaks if needed */
</style>
