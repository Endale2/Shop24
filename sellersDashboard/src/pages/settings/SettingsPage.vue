<template>
  <div class="p-6 bg-gray-50 min-h-screen">
    <div class="max-w-3xl mx-auto bg-white shadow rounded-lg p-6 space-y-6">
      <h1 class="text-2xl font-bold text-gray-800">Shop Settings</h1>

      <form @submit.prevent="save" class="space-y-4">
        <!-- Shop Name -->
        <div>
          <label for="name" class="block text-sm font-medium text-gray-700">
            Shop Name
          </label>
          <input
            id="name"
            v-model="form.name"
            type="text"
            required
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          />
        </div>

        <!-- Description -->
        <div>
          <label for="description" class="block text-sm font-medium text-gray-700">
            Description
          </label>
          <textarea
            id="description"
            v-model="form.description"
            rows="3"
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          />
          <p class="mt-1 text-sm text-gray-500">A brief description of your shop.</p>
        </div>

        <!-- Logo Upload + Preview -->
        <div>
          <label class="block text-sm font-medium text-gray-700">Logo</label>
          <div class="mt-1 flex items-center space-x-4">
            <img
              v-if="previewLogo"
              :src="previewLogo"
              alt="Logo preview"
              class="h-16 w-16 rounded-md object-cover border"
            />
            <div v-else class="h-16 w-16 bg-gray-100 rounded-md flex items-center justify-center">
              <svg class="h-8 w-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M3 7v4m0 0v4m0-4h4m-4 0H3m11-4v12m0 0h4m-4 0H9" />
              </svg>
            </div>
            <input
              type="file"
              accept="image/*"
              @change="onFileChange"
              class="text-sm text-gray-500"
            />
          </div>
        </div>

        <!-- Theme Color -->
        <div>
          <label for="themeColor" class="block text-sm font-medium text-gray-700">
            Theme Accent Color
          </label>
          <input
            id="themeColor"
            v-model="form.themeColor"
            type="color"
            class="mt-1 h-10 w-16 p-0 border-0"
          />
        </div>

        <!-- Buttons -->
        <div class="pt-4 border-t border-gray-200 flex justify-end space-x-3">
          <button
            type="button"
            @click="resetForm"
            class="px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition"
          >
            Reset
          </button>
          <button
            type="submit"
            :disabled="saving"
            class="px-6 py-2 bg-green-600 text-white rounded hover:bg-green-700 focus:ring-2 focus:ring-green-500 transition"
          >
            {{ saving ? 'Saving...' : 'Save Settings' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useShopStore } from '@/store/shops'
import { shopService } from '@/services/shop'

// grab current shop from Pinia
const shopStore = useShopStore()
const current = ref(null)

const saving = ref(false)
const previewLogo = ref('')

// form state, seeded with random/demo values
const form = reactive({
  name: '',
  description: '',
  logoFile: null,
  themeColor: '#10B981', // default green
})

function loadCurrent() {
  current.value = shopStore.active
  if (current.value) {
    form.name = current.value.name
    form.description = current.value.description || ''
    previewLogo.value = current.value.image || ''
    form.themeColor = current.value.themeColor || form.themeColor
  }
}

function onFileChange(ev) {
  const file = ev.target.files[0]
  if (!file) return
  form.logoFile = file
  previewLogo.value = URL.createObjectURL(file)
}

function resetForm() {
  loadCurrent()
}

async function save() {
  saving.value = true
  try {
    // upload logo if changed (you’d replace this stub with your real upload logic)
    let imageUrl = current.value.image
    if (form.logoFile) {
      const uploadRes = await shopService.uploadLogo(current.value.id, form.logoFile)
      imageUrl = uploadRes.url
    }

    // update shop
    const updated = await shopService.updateShop(current.value.id, {
      name: form.name,
      description: form.description,
      image: imageUrl,
      themeColor: form.themeColor,
    })
    shopStore.setActiveShop(updated)
    alert('Settings saved!')
  } catch (err) {
    console.error(err)
    alert('Failed to save settings.')
  } finally {
    saving.value = false
  }
}

onMounted(loadCurrent)
</script>

<style scoped>
/* nothing beyond Tailwind */
</style>
