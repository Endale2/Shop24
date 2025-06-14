<template>
  <div class="min-h-screen bg-gradient-to-br from-green-50 to-white text-gray-800 font-sans">
    <!-- Navigation (Optional, can be added if more pages exist) -->
    <header class="w-full bg-white bg-opacity-90 shadow-sm fixed top-0 left-0 z-40 p-4">
      <div class="max-w-7xl mx-auto flex justify-between items-center">
        <h1 class="text-2xl font-extrabold text-green-700">Seller’s Dash</h1>
        <nav>
          <button
            @click="openAuthModal('login')"
            class="px-5 py-2.5 bg-green-600 text-white font-medium rounded-lg shadow-md
                   hover:bg-green-700 transition duration-200 ease-in-out transform hover:scale-105"
          >
            Login
          </button>
        </nav>
      </div>
    </header>

    <!-- Hero Section -->
    <section class="relative flex items-center justify-center h-screen py-16 px-4 text-center bg-green-700 text-white overflow-hidden">
      <!-- Background elements for visual appeal -->
      <div class="absolute inset-0 z-0 opacity-10">
        <svg class="w-full h-full" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid slice">
          <defs>
            <pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse">
              <path d="M 10 0 L 0 0 L 0 10" fill="none" stroke="white" stroke-width="0.1"></path>
            </pattern>
          </defs>
          <rect width="100%" height="100%" fill="url(#grid)"></rect>
        </svg>
      </div>

      <div class="relative z-10 max-w-4xl mx-auto">
        <h2 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold leading-tight mb-6 animate-fade-in-up">
          Launch Your E-commerce Empire, Effortlessly.
        </h2>
        <p class="text-lg sm:text-xl mb-10 opacity-90 animate-fade-in delay-200">
          Seller’s Dash empowers you to create, manage, and scale your online shops with intuitive tools and powerful features.
        </p>
        <div class="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-6 animate-fade-in delay-400">
          <button
            @click="openAuthModal('register')"
            class="px-8 py-4 bg-white text-green-700 text-xl font-bold rounded-xl shadow-lg
                   hover:bg-gray-100 transition duration-300 ease-in-out transform hover:scale-105"
          >
            Get Started Free
          </button>
          <button
            @click="openAuthModal('login')"
            class="px-8 py-4 border-2 border-white text-white text-xl font-bold rounded-xl shadow-lg
                   hover:bg-white hover:text-green-700 transition duration-300 ease-in-out transform hover:scale-105"
          >
            Existing User? Login
          </button>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-16 px-4 bg-white">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-3xl sm:text-4xl font-extrabold text-gray-900 mb-12 text-center">
          What You Can Do
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-10">
          <div class="flex flex-col items-center text-center p-6 bg-gray-50 rounded-xl shadow-md transition-transform duration-300 hover:scale-105 hover:shadow-lg">
            <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-6 shadow-sm">
              <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 24 24"><path d="M10 12a2 2 0 100-4 2 2 0 000 4z"/><path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd"/></svg>
            </div>
            <h3 class="text-xl font-bold text-gray-900 mb-3">Create & Customize Shops</h3>
            <p class="text-gray-600">
              Build beautiful online stores tailored to your brand, no coding required.
            </p>
          </div>
          <div class="flex flex-col items-center text-center p-6 bg-gray-50 rounded-xl shadow-md transition-transform duration-300 hover:scale-105 hover:shadow-lg">
            <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-6 shadow-sm">
              <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 24 24"><path d="M7 8H6a2 2 0 00-2 2v3a2 2 0 002 2h1v4a1 1 0 001 1h3a1 1 0 001-1v-4h1a2 2 0 002-2V9a2 2 0 00-2-2h-1V3a1 1 0 00-1-1h-3a1 1 0 00-1 1v4H7zM9 4h2v3H9V4zm2 10H9v4h2v-4zM7 9h10V9H7z"/></svg>
            </div>
            <h3 class="text-xl font-bold text-gray-900 mb-3">Manage Products & Inventory</h3>
            <p class="text-gray-600">
              Effortlessly add products, track stock, and organize your catalog.
            </p>
          </div>
          <div class="flex flex-col items-center text-center p-6 bg-gray-50 rounded-xl shadow-md transition-transform duration-300 hover:scale-105 hover:shadow-lg">
            <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mb-6 shadow-sm">
              <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-6h2v6zm0-8h-2V7h2v2z"/></svg>
            </div>
            <h3 class="text-xl font-bold text-gray-900 mb-3">Powerful Discount Engine</h3>
            <p class="text-gray-600">
              Create flexible discounts, coupons, and promotions to boost your sales.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="py-16 px-4 bg-green-600 text-white text-center">
      <div class="max-w-3xl mx-auto">
        <h2 class="text-3xl sm:text-4xl font-extrabold mb-6">Ready to Get Started?</h2>
        <p class="text-lg opacity-90 mb-10">
          Join thousands of successful sellers transforming their businesses with Seller’s Dash.
        </p>
        <button
          @click="openAuthModal('register')"
          class="px-8 py-4 bg-white text-green-700 text-xl font-bold rounded-xl shadow-lg
                 hover:bg-gray-100 transition duration-300 ease-in-out transform hover:scale-105"
        >
          Start Your Free Shop Today
        </button>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-8 px-4 bg-gray-800 text-gray-300 text-center text-sm">
      <p>&copy; 2025 Seller’s Dash. All rights reserved.</p>
      <p class="mt-2">
        <a href="#" class="hover:underline text-gray-400">Privacy Policy</a> |
        <a href="#" class="hover:underline text-gray-400">Terms of Service</a>
      </p>
    </footer>

    <!-- Auth Modal -->
    <div v-if="showAuthModal" class="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center z-50 p-4 animate-fade-in">
      <div class="max-w-md w-full bg-white shadow-2xl rounded-xl overflow-hidden transform scale-95 opacity-0 animate-scale-in">
        <!-- Header -->
        <div class="bg-green-600 px-6 py-4">
          <h2 class="text-2xl font-bold text-white text-center">Seller’s Dash</h2>
          <p class="text-green-200 text-center mt-1">Create & manage your e-commerce shops</p>
          <button @click="closeAuthModal" class="absolute top-4 right-4 text-white hover:text-green-100 transition">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </button>
        </div>

        <!-- Tabs -->
        <div class="flex border-b border-gray-200">
          <button
            @click="modalMode = 'login'"
            :class="tabClass('login')"
            class="flex-1 py-4 text-center font-semibold text-lg transition-colors duration-200"
          >
            Login
          </button>
          <button
            @click="modalMode = 'register'"
            :class="tabClass('register')"
            class="flex-1 py-4 text-center font-semibold text-lg transition-colors duration-200"
          >
            Register
          </button>
        </div>

        <!-- Forms -->
        <div class="px-6 py-8">
          <!-- OAuth button -->
          <button
            @click="loginWithGoogle"
            class="w-full mb-6 flex items-center justify-center px-4 py-3 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors duration-200 shadow-md transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
          >
            <svg class="h-5 w-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
              <path d="M21.805 10.023h-9.728v3.953h5.565c-.238 1.495-1.347 4.378-5.565 4.378-3.345 0-6.07-2.756-6.07-6.153s2.725-6.153 6.07-6.153c1.905 0 3.177.811 3.909 1.503l2.664-2.58C17.646 2.96 15.392 1.75 12.077 1.75 6.013 1.75 1.248 6.562 1.248 12.145c0 5.583 4.765 10.395 10.829 10.395 6.244 0 10.396-4.38 10.396-10.55 0-.713-.077-1.253-.668-1.917z"/>
            </svg>
            Continue with Google
          </button>

          <!-- Divider -->
          <div class="flex items-center mb-6">
            <div class="flex-grow h-px bg-gray-200"></div>
            <span class="px-3 text-gray-500 font-medium">or</span>
            <div class="flex-grow h-px bg-gray-200"></div>
          </div>

          <!-- Login Form -->
          <form v-if="modalMode === 'login'" @submit.prevent="handleLogin" class="space-y-5">
            <div>
              <label for="login-email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
              <input
                id="login-email"
                v-model="email"
                type="email"
                required
                class="mt-1 w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm
                       focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                placeholder="your.email@example.com"
              />
            </div>
            <div>
              <label for="login-password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
              <input
                id="login-password"
                v-model="password"
                type="password"
                required
                class="mt-1 w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm
                       focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                placeholder="••••••••"
              />
            </div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full px-4 py-3 bg-green-600 text-white font-semibold rounded-lg shadow-md
                     hover:bg-green-700 transition duration-200 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed
                     flex items-center justify-center"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ loading ? 'Logging in...' : 'Login' }}
            </button>
            <p v-if="error" class="text-red-500 text-xs text-center mt-3">{{ error }}</p>
          </form>

          <!-- Register Form -->
          <form v-else @submit.prevent="handleRegister" class="space-y-5">
            <div>
              <label for="register-username" class="block text-sm font-medium text-gray-700 mb-1">Username</label>
              <input
                id="register-username"
                v-model="username"
                type="text"
                required
                class="mt-1 w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm
                       focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                placeholder="your_username"
              />
            </div>
            <div>
              <label for="register-email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
              <input
                id="register-email"
                v-model="email"
                type="email"
                required
                class="mt-1 w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm
                       focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                placeholder="your.email@example.com"
              />
            </div>
            <div>
              <label for="register-password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
              <input
                id="register-password"
                v-model="password"
                type="password"
                required
                class="mt-1 w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm
                       focus:ring-green-500 focus:border-green-500 transition duration-150 ease-in-out"
                placeholder="••••••••"
              />
            </div>
            <button
              type="submit"
              :disabled="loading"
              class="w-full px-4 py-3 bg-green-600 text-white font-semibold rounded-lg shadow-md
                     hover:bg-green-700 transition duration-200 ease-in-out disabled:opacity-60 disabled:cursor-not-allowed
                     flex items-center justify-center"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ loading ? 'Creating Account...' : 'Create Account' }}
            </button>
            <p v-if="error" class="text-red-500 text-xs text-center mt-3">{{ error }}</p>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'

// stores & router
const auth = useAuthStore()
const router = useRouter()
const route = useRoute()

// authentication form state
const showAuthModal = ref(false) // Controls visibility of the entire auth modal
const modalMode = ref('login')    // 'login' or 'register' for the modal's tabs

const email = ref('')
const password = ref('')
const username = ref('')
const loading = ref(false)
const error = ref(null)

// Helper to apply active tab classes within the modal
function tabClass(tab) {
  return modalMode.value === tab
    ? 'border-green-600 text-green-600 bg-green-50 border-b-2' // Active tab styling
    : 'text-gray-600 hover:text-green-600 hover:bg-gray-50 border-b-2 border-transparent'; // Inactive tab styling
}

// Functions to control the auth modal
function openAuthModal(mode) {
  modalMode.value = mode; // Set initial mode (login or register)
  showAuthModal.value = true;
  // Reset form fields when opening the modal
  email.value = '';
  password.value = '';
  username.value = '';
  error.value = null;
}

function closeAuthModal() {
  showAuthModal.value = false;
  // Optional: Reset form fields here too if preferred
}

// Google OAuth redirect
function loginWithGoogle() {
  window.location.href =
    'http://localhost:8080/auth/seller/oauth/google?redirect_uri=' +
    encodeURIComponent(window.location.origin + '?oauth=google');
}

// On mount, handle OAuth callback or redirect if already authenticated
onMounted(async () => {
  if (auth.isAuthenticated) {
    return router.replace({ name: 'ShopSelection' });
  }
  if (route.query.oauth === 'google') {
    loading.value = true;
    try {
      await auth.fetchMe(); // Attempt to fetch user data after OAuth redirect
      // Clear query params to prevent re-triggering on refresh
      router.replace({ query: {} });
      return router.replace({ name: 'ShopSelection' });
    } catch {
      error.value = 'Google login failed. Please try again or use email/password.';
      openAuthModal('login'); // Show login form on OAuth failure
    } finally {
      loading.value = false;
    }
  }
});

// handle login
async function handleLogin() {
  loading.value = true;
  error.value = null;
  try {
    await auth.login({ email: email.value, password: password.value });
    closeAuthModal(); // Close modal on successful login
    router.push({ name: 'ShopSelection' });
  } catch (err) {
    console.error('Login error:', err);
    error.value = err.response?.data?.message || 'Invalid email or password.';
  } finally {
    loading.value = false;
  }
}

// handle register
async function handleRegister() {
  loading.value = true;
  error.value = null;
  try {
    await auth.register({ username: username.value, email: email.value, password: password.value });
    closeAuthModal(); // Close modal on successful registration
    router.push({ name: 'ShopSelection' });
  } catch (err) {
    console.error('Registration error:', err);
    error.value = err.response?.data?.message || 'Registration failed. Please check your details.';
  } finally {
    loading.value = false;
  }
}
</script>

<style>
/* Add keyframe animations for modal */
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes scaleIn {
  from { transform: scale(0.95); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.animate-fade-in {
  animation: fadeIn 0.3s ease-out forwards;
}

.animate-scale-in {
  animation: scaleIn 0.3s ease-out forwards;
}

/* Animations for hero section */
@keyframes fadeInFromBottom {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fadeInElement {
  from { opacity: 0; }
  to { opacity: 1; }
}

.animate-fade-in-up {
  animation: fadeInFromBottom 0.8s ease-out forwards;
}

.animate-fade-in {
  animation: fadeInElement 0.8s ease-out forwards;
}

/* Delay utility for hero section elements */
.delay-200 { animation-delay: 0.2s; }
.delay-400 { animation-delay: 0.4s; }
</style>