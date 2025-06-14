<template>
  <div class="min-h-screen bg-gray-50 py-6 px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto">
      <h1 class="text-2xl font-bold text-gray-900 mb-6">Shop Settings</h1>

      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="p-6 grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Left column: form inputs -->
          <form @submit.prevent="save" class="space-y-6">
            <!-- Shop Name -->
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700">Shop Name</label>
              <input
                id="name"
                v-model="form.name"
                type="text"
                required
                class="mt-1 block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <!-- Description -->
            <div>
              <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
              <textarea
                id="description"
                v-model="form.description"
                rows="3"
                class="mt-1 block w-full border-gray-300 rounded-lg shadow-sm focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <!-- Theme Color Picker -->
            <div class="flex items-center space-x-4">
              <div>
                <label for="themeColor" class="block text-sm font-medium text-gray-700">Accent Color</label>
                <input
                  id="themeColor"
                  v-model="form.themeColor"
                  type="color"
                  class="mt-1 h-10 w-12 p-0 border-0 cursor-pointer"
                />
              </div>
              <div class="flex-1">
                <p class="text-sm text-gray-600">Your shop’s accent color.</p>
              </div>
            </div>
          </form>

          <!-- Right column: logo upload + preview -->
          <div class="flex flex-col items-center justify-center space-y-4">
            <label class="block text-sm font-medium text-gray-700">Shop Logo</label>

            <div
              class="w-32 h-32 border-2 border-dashed rounded-lg flex items-center justify-center cursor-pointer hover:border-green-500 transition"
              @click="$refs.fileInput.click()"
            >
              <img
                v-if="previewLogo"
                :src="previewLogo"
                alt="Logo preview"
                class="max-h-full max-w-full rounded-md object-cover"
              />
              <div v-else class="text-gray-400">
                <svg class="h-10 w-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M3 7v4m0 0v4m0-4h4m-4 0H3m11-4v12m0 0h4m-4 0H9" />
                </svg>
              </div>
            </div>

            <input
              type="file"
              accept="image/*"
              class="hidden"
              ref="fileInput"
              @change="onFileChange"
            />

            <p class="text-sm text-gray-500">Click above to upload or replace your logo.</p>
          </div>
        </div>

        <!-- action buttons span full width -->
        <div class="px-6 py-4 bg-gray-50 flex flex-col sm:flex-row justify-end gap-3">
          <button
            @click="resetForm"
            class="px-4 py-2 bg-white border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-100 transition"
          >
            Reset
          </button>
          <button
            @click="save"
            :disabled="saving"
            class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:ring-2 focus:ring-green-500 transition disabled:opacity-50"
          >
            {{ saving ? 'Saving...' : 'Save Changes' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useShopStore } from '@/store/shops'
import { shopService } from '@/services/shop'

const shopStore = useShopStore()
const current = ref(null)
const saving = ref(false)
const previewLogo = ref('')

// reactive form
const form = reactive({
  name: '',
  description: '',
  logoFile: null,
  themeColor: '#10B981',
})

function loadCurrent() {
  current.value = shopStore.active
  if (!current.value) return
  form.name = current.value.name
  form.description = current.value.description || ''
  previewLogo.value = current.value.image || ''
  form.themeColor = current.value.themeColor || form.themeColor
}

function onFileChange(e) {
  const f = e.target.files[0]
  if (!f) return
  form.logoFile = f
  previewLogo.value = URL.createObjectURL(f)
}

function resetForm() {
  loadCurrent()
}

async function save() {
  if (!current.value) return
  saving.value = true
  try {
    let imageUrl = current.value.image
    if (form.logoFile) {
      // replace with real upload endpoint
      const up = await shopService.uploadLogo(current.value.id, form.logoFile)
      imageUrl = up.url
    }
    const updated = await shopService.updateShop(current.value.id, {
      name: form.name,
      description: form.description,
      image: imageUrl,
      themeColor: form.themeColor,
    })
    shopStore.setActiveShop(updated)
    // optionally a toast here
  } catch (err) {
    console.error(err)
    // optionally error toast
  } finally {
    saving.value = false
  }
}

onMounted(loadCurrent)
</script>

<style scoped>
/* nothing extra—Tailwind handles it all */
</style>
