<!-- src/pages/products/AddProductPage.vue -->
<template>
  <div class="p-4 sm:p-6 max-w-6xl mx-auto space-y-8 font-sans">
    <!-- Back Button -->
    <button @click="goBack" class="inline-flex items-center text-gray-600 hover:text-green-700 mb-6">
      <ChevronLeftIcon class="h-5 w-5 mr-1 text-gray-500"/> 
      <span class="text-sm font-medium">Back to Products</span>
    </button>

    <!-- Loading / Error -->
    <div v-if="loading" class="flex flex-col items-center text-gray-600 py-12">
      <SpinnerIcon class="animate-spin h-8 w-8 text-green-500 mb-3"/>
      <p class="mt-3 text-lg">Preparing form...</p>
    </div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded">
      <strong>Error:</strong> {{ error }}
    </div>

    <!-- Form -->
    <div v-else class="bg-white shadow rounded-xl p-8 space-y-8 animate-fade-in">
      <h2 class="text-3xl font-bold">Add New Product</h2>

      <form @submit.prevent="submitForm" class="space-y-6">
        <!-- Basic Fields -->
        <div>
          <label class="block text-sm font-medium">Name</label>
          <input v-model="form.name" type="text" required class="mt-1 w-full border rounded px-3 py-2"/>
        </div>

        <div>
          <label class="block text-sm font-medium">Description</label>
          <textarea v-model="form.description" rows="4" class="mt-1 w-full border rounded px-3 py-2"></textarea>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium">Price</label>
            <input v-model.number="form.price" type="number" min="0" step="0.01" required
                   class="mt-1 w-32 border rounded px-3 py-2"/>
          </div>
          <div>
            <label class="block text-sm font-medium">Category</label>
            <input v-model="form.category" type="text" class="mt-1 w-full border rounded px-3 py-2"/>
          </div>
        </div>

        <!-- Images -->
        <div>
          <label class="block text-sm font-medium">Main Image URL</label>
          <input v-model="form.main_image" type="url" placeholder="https://…" class="mt-1 w-full border rounded px-3 py-2"/>
        </div>

        <div>
          <label class="block text-sm font-medium">Additional Images</label>
          <div v-for="(img,i) in form.images" :key="i" class="flex items-center space-x-2 mb-2">
            <input v-model="form.images[i]" type="url" placeholder="https://…" class="flex-1 border rounded px-3 py-2"/>
            <button type="button" @click="removeImage(i)" class="text-red-600">
              <MinusCircleIcon class="h-5 w-5"/>
            </button>
          </div>
          <button type="button" @click="addImage" class="text-blue-600 inline-flex items-center">
            <PlusCircleIcon class="h-5 w-5 mr-1"/> Add Image
          </button>
        </div>

        <!-- Variants -->
        <div>
          <h3 class="text-lg font-semibold mb-4">Variants</h3>
          <div v-for="(v,vi) in form.variants" :key="vi" class="border rounded p-4 mb-4">
            <div class="flex justify-between mb-3">
              <span>Variant {{ vi+1 }}</span>
              <button type="button" @click="removeVariant(vi)" class="text-red-600">
                <XIcon class="h-5 w-5"/>
              </button>
            </div>
            <!-- Options -->
            <div v-for="(val,key) in v.options" :key="key" class="flex items-center space-x-2 mb-2">
              <input v-model="variantKeyTemp[vi][key]"
                     @blur="renameOptionKey(vi, key, variantKeyTemp[vi][key])"
                     placeholder="Option Name"
                     class="w-1/3 border rounded px-2 py-1 text-sm"/>
              <input v-model="v.options[key]" placeholder="Option Value" class="flex-1 border rounded px-2 py-1 text-sm"/>
              <button type="button" @click="removeVariantOption(vi,key)" class="text-red-600">
                <MinusCircleIcon class="h-4 w-4"/>
              </button>
            </div>
            <button type="button" @click="addVariantOption(vi)" class="inline-flex items-center text-sm text-blue-600">
              <PlusCircleIcon class="h-4 w-4 mr-1"/> Add Option
            </button>

            <!-- Price & Image -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 mt-4">
              <div>
                <label class="block text-sm font-medium">Variant Price</label>
                <input v-model.number="v.price" type="number" step="0.01" min="0" required class="mt-1 w-32 border rounded px-3 py-2"/>
              </div>
              <div>
                <label class="block text-sm font-medium">Variant Image URL</label>
                <input v-model="v.image" type="url" placeholder="https://…" class="mt-1 w-full border rounded px-3 py-2"/>
              </div>
            </div>
          </div>

          <button type="button" @click="addVariant" class="inline-flex items-center px-4 py-2 bg-green-600 text-white rounded">
            <PlusIcon class="h-5 w-5 mr-1"/> Add Variant
          </button>
        </div>

        <!-- Actions -->
        <div class="flex justify-end space-x-4 mt-8">
          <button type="button" @click="goBack" class="px-6 py-2 border rounded">Cancel</button>
          <button type="submit" :disabled="saving" class="px-6 py-2 bg-green-600 text-white rounded">
            <SpinnerIcon v-if="saving" class="animate-spin h-5 w-5 mr-2 inline-block"/>
            {{ saving ? 'Saving...' : 'Save Product' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter }       from 'vue-router'
import { useShopStore }    from '@/store/shops'
import { useAuthStore }    from '@/store/auth'
import { productService }  from '@/services/product'
import {
  ChevronLeftIcon,
  RefreshIcon   as SpinnerIcon,
  PlusCircleIcon,
  MinusCircleIcon,
  XIcon,
  PlusIcon
} from '@heroicons/vue/outline'

const router    = useRouter()
const shopStore = useShopStore()
const authStore = useAuthStore()

const loading = ref(true)
const saving  = ref(false)
const error   = ref(null)

const form = reactive({
  name:        '',
  description: '',
  main_image:  '',
  images:      [],
  category:    '',
  price:       0,
  variants:    []
})

// temp keys for option renaming
const variantKeyTemp = reactive({})

function goBack() {
  router.push({ name: 'Products' })
}

onMounted(() => {
  loading.value = false
})

// Images
function addImage()    { form.images.push('') }
function removeImage(i){ form.images.splice(i,1) }

// Variants
function addVariant() {
  form.variants.push({ options: {}, price: 0, image: '' })
  variantKeyTemp[form.variants.length-1] = {}
}
function removeVariant(i) {
  form.variants.splice(i,1)
  delete variantKeyTemp[i]
}

// Options
function addVariantOption(vi) {
  const key = `opt_${Date.now()}`
  form.variants[vi].options[key] = ''
  variantKeyTemp[vi][key] = key
}
function removeVariantOption(vi,key) {
  delete form.variants[vi].options[key]
  delete variantKeyTemp[vi][key]
}
function renameOptionKey(vi, oldKey, newKey) {
  if (!newKey || oldKey===newKey) return
  const val = form.variants[vi].options[oldKey]
  delete form.variants[vi].options[oldKey]
  form.variants[vi].options[newKey] = val
  delete variantKeyTemp[vi][oldKey]
  variantKeyTemp[vi][newKey] = newKey
}

// Submit
async function submitForm() {
  saving.value = true
  error.value  = null

  try {
    const payload = {
      name:        form.name,
      description: form.description,
      main_image:  form.main_image || undefined,
      images:      form.images.filter(i=>i),
      category:    form.category,
      price:       form.price,
      createdBy:   authStore.user.id,
      variants:    form.variants.map(v => ({
        name:    Object.keys(v.options)[0] || '',
        options: Object.values(v.options),
        prices:  [v.price]
      }))
    }

    await productService.create(shopStore.active.id, payload)
    router.push({ name: 'Products' })
  } catch (e) {
    console.error(e)
    error.value = 'Failed to create product. Please try again.'
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity:0; transform:translateY(10px) }
  to   { opacity:1; transform:none         }
}
.animate-fade-in { animation:fadeIn .4s ease-out forwards }
</style>
