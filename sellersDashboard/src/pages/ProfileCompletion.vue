<template>
  <div class="min-h-screen flex flex-col bg-amber-50 font-sans antialiased">
    <AppNavbar />

    <div class="flex-1 flex items-center justify-center p-4 sm:p-6 lg:p-8 animate-fade-in-up">
      <div class="bg-white p-8 rounded-lg shadow-2xl w-full max-w-md border border-green-200">
        <h2 class="text-2xl font-bold mb-6 text-center text-green-800">Complete Your Profile</h2>

        <form @submit.prevent="submit" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">First Name</label>
            <input
              v-model="firstName"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-400 transition"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Last Name</label>
            <input
              v-model="lastName"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-400 transition"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Phone (optional)</label>
            <input
              v-model="phone"
              type="tel"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-400 transition"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Business Name (optional)</label>
            <input
              v-model="businessName"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-400 transition"
            />
          </div>

          <div class="text-center mt-6">
            <button
              type="submit"
              :disabled="loading"
              class="bg-green-600 text-white px-6 py-2 rounded-md hover:bg-green-700 disabled:opacity-50 transition flex items-center justify-center mx-auto"
            >
              <span v-if="loading" class="inline-block animate-spin mr-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
                </svg>
              </span>
              {{ loading ? 'Saving…' : 'Save Profile' }}
            </button>
          </div>

          <p v-if="error" class="text-red-500 text-sm text-center mt-2">{{ error }}</p>
        </form>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import AppNavbar from '@/components/AppNavbar.vue' // Import AppNavbar
import AppFooter from '@/components/AppFooter.vue' // Import AppFooter

const auth   = useAuthStore()
const router = useRouter()

// 1) guard: if somehow not logged in, kick to Login
if (!auth.isAuthenticated) {
  router.replace({ name: 'Login' })
}

// form state
const firstName    = ref('')
const lastName     = ref('')
const phone        = ref('')
const businessName = ref('')
const loading      = ref(false)
const error        = ref('')

// 2) pre-fill from store if already present
onBeforeMount(() => {
  const u = auth.user || {}
  firstName.value    = u.firstName    || ''
  lastName.value     = u.lastName     || ''
  phone.value        = u.phone        || ''
  businessName.value = u.businessName || ''
})

async function submit() {
  loading.value = true
  error.value   = ''

  try {
    // 3) send exactly the camelCase fields your service expects
    await auth.updateProfile({
      firstName:    firstName.value.trim(),
      lastName:     lastName.value.trim(),
      phone:        phone.value.trim(),
      businessName: businessName.value.trim(),
    })

    // 4) refresh the store so auth.user is fully up to date
    await auth.fetchMe()

    // 5) push to shops
    router.replace({ name: 'ShopSelection' })
  } catch (err) {
    console.error(err)
    error.value = err.response?.data?.error
                || err.message
                || 'Failed to save profile'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>

</style>