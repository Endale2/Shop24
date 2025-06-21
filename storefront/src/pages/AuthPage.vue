<template>
  <div class="max-w-md mx-auto mt-16 p-6 bg-white rounded-lg shadow-xl border border-gray-100"> <h2 class="text-3xl font-extrabold text-gray-800 mb-8 text-center">
      Welcome to <span class="text-transparent bg-clip-text bg-gradient-to-r from-green-600 to-green-800">{{ shopName }}</span>
    </h2>

    <form @submit.prevent="mode === 'login' ? onLogin() : onRegister()" class="space-y-5"> <template v-if="mode === 'register'">
        <input
          v-model="firstName"
          placeholder="First Name"
          type="text"
          :disabled="loading"
          required
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
        />
        <input
          v-model="lastName"
          placeholder="Last Name"
          type="text"
          :disabled="loading"
          required
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
        />
        <input
          v-model="phone"
          placeholder="Phone (Optional)"
          type="tel"
          :disabled="loading"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
        />
        <input
          v-model="address"
          placeholder="Address (Optional)"
          type="text"
          :disabled="loading"
          class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
        />
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <input
            v-model="city"
            placeholder="City (Optional)"
            type="text"
            :disabled="loading"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
          />
          <input
            v-model="state"
            placeholder="State (Optional)"
            type="text"
            :disabled="loading"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
          />
          <input
            v-model="postalCode"
            placeholder="Postal Code (Optional)"
            type="text"
            :disabled="loading"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
          />
          <input
            v-model="country"
            placeholder="Country (Optional)"
            type="text"
            :disabled="loading"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
          />
        </div>
      </template>

      <input
        v-model="email"
        placeholder="Email"
        type="email"
        :disabled="loading"
        required
        class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
      />
      <input
        v-model="password"
        placeholder="Password"
        type="password"
        :disabled="loading"
        required
        class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-green-400 transition duration-200 text-gray-700 placeholder-gray-400 shadow-sm"
      />

      <button
        type="submit"
        :disabled="loading"
        class="w-full p-3 text-white rounded-lg font-semibold text-lg shadow-md hover:shadow-lg transition duration-200 ease-in-out transform hover:-translate-y-0.5"
        :class="mode === 'login' ? 'bg-blue-600 hover:bg-blue-700' : 'bg-green-600 hover:bg-green-700'"
      >
        {{ loading ? (mode === 'login' ? 'Logging in…' : 'Registering…') : (mode === 'login' ? 'Login' : 'Create Account') }} </button>

      <p class="text-sm text-center text-gray-600 mt-4">
        <template v-if="mode === 'login'">
          Don't have an account? <a @click.prevent="mode = 'register'" href="#" class="text-green-600 hover:text-green-700 font-medium transition-colors duration-200">Register now</a>
        </template>
        <template v-else>
          Already have an account? <a @click.prevent="mode = 'login'" href="#" class="text-blue-600 hover:text-blue-700 font-medium transition-colors duration-200">Login here</a>
        </template>
      </p>
    </form>

    <div class="relative flex items-center justify-center my-8">
      <div class="absolute inset-x-0 h-px bg-gray-200"></div>
      <span class="relative bg-white px-4 text-sm text-gray-500">OR</span>
    </div>

    <div class="mt-6 text-center">
      <button
        @click="onOAuth"
        :disabled="loading"
        class="w-full p-3 border border-gray-300 rounded-lg flex items-center justify-center space-x-2 text-gray-700 font-medium hover:bg-gray-50 transition duration-200 ease-in-out shadow-sm"
      >
        <img src="https://www.svgrepo.com/show/354431/google.svg" alt="Google icon" class="h-5 w-5" />
        <span>Continue with Google</span>
      </button>
    </div>

    <p v-if="error" class="mt-6 text-red-500 text-center text-sm font-medium bg-red-50 p-3 rounded-lg border border-red-200">
      {{ error }}
    </p>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useShopStore } from '@/stores/shop'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const shop = useShopStore()
const router = useRouter()

const mode = ref('login')
const email = ref('')
const password = ref('')

// New fields for registration
const firstName = ref('')
const lastName = ref('')
const phone = ref('')
const address = ref('')
const city = ref('')
const state = ref('')
const postalCode = ref('')
const country = ref('')

const loading = computed(() => auth.loading)
const error = computed(() => auth.error)
const shopName = computed(() => shop.current?.name || 'Store')

// Get shopId from the current shop store
const currentShopId = computed(() => shop.current?.id); // Assuming your shop store has an 'id' property

// Ensure shop data is loaded (for shopName and shopId)
onMounted(async () => {
  if (!shop.current) {
    const slug = window.location.host.split('.')[0]
    try {
      await shop.fetchShopAndProducts(slug)
    } catch (e) {
      console.error('Failed to load shop data:', e)
    }
  }
})

// Login or register
const onLogin = async () => {
  try {
    await auth.authenticate({ email: email.value, password: password.value })
    router.push('/')
  } catch { /* Error handled by auth store and displayed in template */ }
}

const onRegister = async () => {
  try {
    const customerData = {
      firstName: firstName.value,
      lastName: lastName.value,
      email: email.value,
      password: password.value,
      phone: phone.value || undefined, // Send undefined if empty
      address: address.value || undefined,
      city: city.value || undefined,
      state: state.value || undefined,
      postalCode: postalCode.value || undefined,
      country: country.value || undefined,
      shopId: currentShopId.value // Pass the shopId here
    };

    // Filter out undefined values to send a cleaner object
    const payload = Object.fromEntries(
      Object.entries(customerData).filter(([, value]) => value !== undefined && value !== '')
    );

    await auth.register(payload);
    router.push('/');
  } catch { /* Error handled by auth store and displayed in template */ }
}

// Google OAuth
const onOAuth = () => {
  auth.oauthLogin()
}
</script>

<style scoped>
/* Tailwind handles styles */
</style>