<template>
  <div class="min-h-screen bg-gray-50 flex flex-col justify-center items-center px-6 py-4">
    <div class="w-full max-w-md">
      <div class="text-center mb-6">
        <h1 class="text-3xl font-extrabold text-gray-900">
          Create a Seller Account
        </h1>
        <p class="mt-2 text-sm text-gray-600">
          Already have an account?
          <router-link to="/login" class="font-medium text-green-600 hover:text-green-500">
            Sign in
          </router-link>
        </p>
      </div>

      <div class="bg-white p-8 rounded-2xl shadow-lg">
        <form @submit.prevent="handleRegister" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label for="firstName" class="block text-sm font-medium text-gray-700">First Name</label>
              <input
                id="firstName"
                v-model="firstName"
                required
                class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
              />
            </div>
            <div>
              <label for="lastName" class="block text-sm font-medium text-gray-700">Last Name</label>
              <input
                id="lastName"
                v-model="lastName"
                required
                class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
              />
            </div>
          </div>

          <div>
            <label for="businessName" class="block text-sm font-medium text-gray-700">Business Name</label>
            <input
              id="businessName"
              v-model="businessName"
              class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
            />
          </div>

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email Address</label>
            <input
              id="email"
              type="email"
              v-model="email"
              required
              class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
            />
          </div>

            <div>
            <label for="phone" class="block text-sm font-medium text-gray-700">Phone Number (Optional)</label>
            <input
              id="phone"
              type="tel"
              v-model="phone"
              class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
            />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
              id="password"
              type="password"
              v-model="password"
              required
              class="mt-1 w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-green-500 focus:border-green-500 transition"
            />
          </div>

          <div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center items-center bg-green-600 text-white py-2.5 px-4 rounded-lg font-semibold hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition disabled:opacity-60"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
              </svg>
              <span>{{ loading ? 'Creating Account...' : 'Create Account' }}</span>
            </button>
          </div>
          <p v-if="error" class="text-sm text-red-600 text-center">{{ error }}</p>
        </form>

        <div class="mt-6 relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <div class="mt-6">
          <button
            @click="loginWithGoogle"
            class="w-full flex justify-center items-center space-x-3 border border-gray-300 rounded-lg py-2.5 px-4 hover:bg-gray-50 transition"
          >
            <img src="https://www.svgrepo.com/show/355037/google.svg" alt="Google" class="h-5 w-5" />
            <span class="text-sm font-medium text-gray-700">Continue with Google</span>
          </button>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/store/auth';

const router = useRouter();
const auth = useAuthStore();

const firstName = ref('');
const lastName = ref('');
const businessName = ref('');
const email = ref('');
const phone = ref('');
const password = ref('');
const loading = ref(false);
const error = ref(null);

async function handleRegister() {
  loading.value = true;
  error.value = null;

  const registrationData = {
    firstName: firstName.value,
    lastName: lastName.value,
    businessName: businessName.value,
    email: email.value,
    phone: phone.value,
    password: password.value,
  };

  try {
    await auth.register(registrationData);
    router.push({ name: 'ShopSelection' });
  } catch (err) {
    error.value = err.response?.data?.error || 'An error occurred during registration.';
    console.error('Registration error:', err);
  } finally {
    loading.value = false;
  }
}

// Kick off OAuth login by hitting your backend endpoint
function loginWithGoogle() {
  // We redirect to the login page which is already equipped to handle the OAuth callback.
  // This avoids duplicating the callback logic.
  const redirect = encodeURIComponent(
    window.location.origin + '/login?oauth=google'
  );
  window.location.href =
    `http://localhost:8080/auth/seller/oauth/google?redirect_uri=${redirect}`;
}
</script>