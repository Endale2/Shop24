<template>
  <Breadcrumbs :items="[
    { back: true, label: 'Back' },
    { label: 'Home', to: `/` },
    { label: 'Account' }
  ]" />
  
  <!-- Loading State -->
  <div v-if="authStore.sessionLoading" class="account-container flex items-center justify-center min-h-[60vh]">
    <div class="flex flex-col items-center">
      <svg class="animate-spin h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <p class="text-gray-500">Checking authentication…</p>
    </div>
  </div>
  
  <!-- Not Authenticated State -->
  <LoginPrompt
    v-else-if="!user"
    :title="'Sign in to view your account'"
    :message="'Please log in to access your account details and profile.'"
  />
  
  <!-- Authenticated State -->
  <div v-else class="account-container">
    <!-- Profile Completion Form -->
    <div v-if="!profileComplete" class="max-w-md mx-auto bg-white border border-gray-200 rounded-none p-8 mt-8">
      <h2 class="text-2xl font-bold text-gray-900 mb-4">Complete Your Profile</h2>
      <form @submit.prevent="onCompleteProfile">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">First Name</label>
            <input v-model="form.firstName" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Last Name</label>
            <input v-model="form.lastName" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Phone</label>
            <input v-model="form.phone" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Address</label>
            <input v-model="form.address" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">City</label>
            <input v-model="form.city" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">State</label>
            <input v-model="form.state" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Postal Code</label>
            <input v-model="form.postalCode" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide mb-2">Country</label>
            <input v-model="form.country" type="text" required class="w-full p-3 border border-gray-200 rounded-none focus:ring-2 focus:ring-black focus:outline-none text-sm uppercase" />
          </div>
        </div>
        <button type="submit" :disabled="loading" class="w-full bg-black text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center mt-6">
          <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ loading ? 'Saving...' : 'Save Profile' }}
        </button>
        <div v-if="error" class="bg-red-50 border border-red-200 rounded-none p-4 mt-4">
          <p class="text-red-800 text-sm">{{ error }}</p>
        </div>
      </form>
    </div>
    
    <!-- Profile Display -->
    <div v-else class="max-w-4xl mx-auto">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">My Account</h1>
        <p class="text-gray-600">Manage your profile and preferences</p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Profile Card -->
        <div class="lg:col-span-2">
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h2 class="text-xl font-semibold text-gray-900 mb-6 uppercase">Profile Information</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Email</label>
                  <p class="mt-1 text-gray-900">{{ user.email }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">First Name</label>
                  <p class="mt-1 text-gray-900">{{ user.firstName || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Last Name</label>
                  <p class="mt-1 text-gray-900">{{ user.lastName || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Phone</label>
                  <p class="mt-1 text-gray-900">{{ user.phone || 'Not provided' }}</p>
                </div>
              </div>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Address</label>
                  <p class="mt-1 text-gray-900">{{ user.address || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">City</label>
                  <p class="mt-1 text-gray-900">{{ user.city || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">State</label>
                  <p class="mt-1 text-gray-900">{{ user.state || 'Not provided' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Country</label>
                  <p class="mt-1 text-gray-900">{{ user.country || 'Not provided' }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          <!-- Quick Actions -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 uppercase">Quick Actions</h3>
            <div class="space-y-3">
              <router-link 
               :to="`/orders`" 
                class="flex items-center p-3 text-gray-700 hover:bg-gray-50 transition-colors rounded-lg"
              >
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                My Orders
              </router-link>
              <router-link 
               :to="`/wishlist`" 
                class="flex items-center p-3 text-gray-700 hover:bg-gray-50 transition-colors rounded-lg"
              >
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                </svg>
                Wishlist
              </router-link>
            </div>
          </div>

          <!-- Account Info -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4 uppercase">Account Info</h3>
            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">User ID</label>
                <p class="mt-1 text-sm text-gray-600 font-mono">{{ user.id || user._id }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 uppercase tracking-wide">Member Since</label>
                <p class="mt-1 text-sm text-gray-600">{{ formatDate(user.createdAt) }}</p>
              </div>
            </div>
          </div>

          <!-- Logout -->
          <div class="bg-white border border-gray-200 rounded-none p-6">
            <button 
              @click="onLogout" 
              class="w-full bg-red-600 text-white py-3 px-4 rounded-none font-semibold uppercase tracking-wide hover:bg-red-700 transition-colors"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, reactive, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import { updateCustomerProfile } from '../services/auth';
import LoginPrompt from '../components/LoginPrompt.vue';
import Breadcrumbs from '../components/Breadcrumbs.vue';

const router = useRouter();
const authStore = useAuthStore();
// No direct use; slug retrieved via services when needed
const user = computed(() => authStore.user);
const profileComplete = computed(() => authStore.profileComplete);
const loading = ref(false);
const error = ref('');

const form = reactive({
  firstName: '',
  lastName: '',
  phone: '',
  address: '',
  city: '',
  state: '',
  postalCode: '',
  country: '',
});

// Initialize form when user data is available
watch(user, (newUser) => {
  if (newUser) {
    form.firstName = newUser.firstName || '';
    form.lastName = newUser.lastName || '';
    form.phone = newUser.phone || '';
    form.address = newUser.address || '';
    form.city = newUser.city || '';
    form.state = newUser.state || '';
    form.postalCode = newUser.postalCode || '';
    form.country = newUser.country || '';
  }
}, { immediate: true });

async function onCompleteProfile() {
  loading.value = true;
  error.value = '';
  try {
    await updateCustomerProfile(form);
    await authStore.fetchProfile();
    // Redirect to homepage after profile completion
    router.push({ path: `/` });
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Failed to update profile';
  } finally {
    loading.value = false;
  }
}

function formatDate(dateString: string) {
  if (!dateString) return 'Unknown';
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
}

async function onLogout() {
  await authStore.logout();
  // Redirect is now handled in the auth store
}
</script>

<style scoped>
.account-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 