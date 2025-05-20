<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-6 rounded shadow w-full max-w-sm">
      <h2 class="text-xl font-semibold mb-4 text-center">Login</h2>

      <form @submit.prevent="handleLogin" class="space-y-3">
        <div>
          <label class="block text-sm mb-1">Email</label>
          <input
            v-model="email"
            type="email"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <div>
          <label class="block text-sm mb-1">Password</label>
          <input
            v-model="password"
            type="password"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-blue-500 text-white py-1 rounded text-sm hover:bg-blue-600 disabled:opacity-50"
        >
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>

        <p v-if="error" class="text-red-500 text-xs text-center mt-2">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'LoginPage',
  data() {
    return {
      email: '',
      password: '',
      loading: false,
      error: null
    };
  },
  methods: {
    ...mapActions('auth', ['login']),
    async handleLogin() {
      this.loading = true;
      this.error = null;
      try {
        await this.login({ email: this.email, password: this.password });
        this.$router.push({ name: 'ShopSelection' }); // Fixed name: 'ShopSelection'
      } catch (err) {
        this.error = 'Login failed. Please check your credentials.';
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
