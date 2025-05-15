<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
    <div class="max-w-sm w-full p-8 bg-white dark:bg-gray-800 rounded-lg shadow-xl">
      <h2 class="text-3xl font-bold text-center text-gray-900 dark:text-white mb-6">
        Admin Login
      </h2>
      <form @submit.prevent="submit" class="space-y-6">
        <div>
          <label for="email" class="sr-only">Email address</label>
          <input
            id="email"
            v-model="email"
            type="email"
            autocomplete="email"
            required
            class="appearance-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-600 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500 sm:text-sm bg-white dark:bg-gray-700"
            placeholder="Email address"
          />
        </div>

        <div>
          <label for="password" class="sr-only">Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            autocomplete="current-password"
            required
            class="appearance-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-600 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500 sm:text-sm bg-white dark:bg-gray-700"
            placeholder="Password"
          />
        </div>

        <div v-if="error" class="text-red-600 dark:text-red-400 text-sm text-center">
          {{ error }}
        </div>

        <div>
          <button
            type="submit"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="loading"
          >
            <span v-if="loading" class="absolute left-0 inset-y-0 flex items-center pl-3">
              <svg class="animate-spin h-5 w-5 text-blue-300 dark:text-blue-200" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
              </svg>
            </span>
            {{ loading ? 'Logging In...' : 'Login' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import auth from '../services/auth';

export default {
  name: 'AdminLogin', // Changed name to be more specific
  data() {
    return {
      email: '',
      password: '',
      error: null,
      loading: false, // Added loading state
    };
  },
  methods: {
    async submit() {
      this.error = null;
      this.loading = true; // Start loading
      try {
        // await auth.login({ email: this.email, password: this.password });
         // Assuming auth.login returns a promise that resolves on success
        const response = await auth.login({ email: this.email, password: this.password });
         // Optional: check response data if needed, but auth.login handles the API call

        // Redirect only after successful login API call
        console.log('Login successful, redirecting...');
        this.$router.replace({ name: 'Dashboard' }); // Use replace to prevent going back to login

      } catch (e) {
        console.error('Login failed:', e);
        // Provide more specific error message if available from backend response
        this.error = e.response?.data?.message || 'Invalid email or password.'; // Use a more common message
      } finally {
        this.loading = false; // Stop loading regardless of success or failure
      }
    },
  },
  // Optional: Redirect if already authenticated
  async created() {
    try {
      const isAuthenticated = await auth.isAuthenticated();
      if (isAuthenticated) {
        console.log('Already authenticated, redirecting to dashboard...');
        this.$router.replace({ name: 'Dashboard' });
      }
    } catch (e) {
        console.warn('Auth check failed on login page load, proceeding with login form:', e.message);
       // If isAuthenticated throws (e.g. API down), just allow login
    }
  }
};
</script>

<style scoped>
/* Add any specific styles not covered by Tailwind if needed */
/* Custom spinner animation (already added in AdminLayout, but good to have here too if used independently) */
@keyframes spin {
  to { transform: rotate(360deg); }
}
.animate-spin {
  animation: spin 1s linear infinite;
}

/* Utility to hide labels visually but keep for screen readers */
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}
</style>