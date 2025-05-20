<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100 px-4">
    <div class="bg-white p-6 rounded shadow-md w-full max-w-md">
      <h2 class="text-xl font-semibold mb-4 text-center">Create Your Account</h2>

      <form @submit.prevent="handleRegister" class="space-y-4">
        <!-- Username -->
        <div>
          <label class="block text-sm mb-1" for="username">Username</label>
          <input
            id="username"
            v-model="username"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <!-- Email -->
        <div>
          <label class="block text-sm mb-1" for="email">Email</label>
          <input
            id="email"
            type="email"
            v-model="email"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <!-- Password -->
        <div>
          <label class="block text-sm mb-1" for="password">Password</label>
          <input
            id="password"
            type="password"
            v-model="password"
            required
            class="w-full px-2 py-1 border rounded text-sm"
          />
        </div>

        <!-- Submit -->
        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-green-500 text-white py-1 rounded text-sm hover:bg-green-600 disabled:opacity-50"
        >
          {{ loading ? 'Creating...' : 'Create Account' }}
        </button>

        <!-- Error -->
        <p v-if="error" class="text-red-500 text-xs text-center mt-2">
          {{ error }}
        </p>
      </form>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'RegisterPage',
  data() {
    return {
      username: '',
      email: '',
      password: '',
      loading: false,
      error: null
    };
  },
  methods: {
    ...mapActions('auth', ['register']),
    async handleRegister() {
      this.loading = true;
      this.error = null;
      try {
        // your auth service expects { username, email, password }
        await this.register({ username: this.username, email: this.email, password: this.password });
        this.$router.push({ name: 'ShopSelection' });
      } catch (err) {
        this.error = err.response?.data?.error || 'Registration failed.';
        console.error('Registration error:', err);
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
