<template>
  <div class="container mx-auto py-8 px-4 max-w-md">
    <h2 class="text-2xl font-bold mb-6 text-center">Create Your Account</h2>
    <form @submit.prevent="handleRegister" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
          Email:
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="email"
          type="email"
          v-model="email"
          required
        />
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
          Password:
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
          id="password"
          type="password"
          v-model="password"
          required
        />
        </div>
      <div class="flex items-center justify-between">
        <button
          class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
          type="submit"
        >
          Create Account
        </button>
      </div>
    </form>
    </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'RegisterPage',
  data() {
    return {
      email: '',
      password: ''
    };
  },
  methods: {
    ...mapActions('auth', ['register']),
    async handleRegister() {
      try {
        await this.register({ email: this.email, password: this.password });
        // Assuming 'ShopSelector' is the name of the route for selecting a shop
        this.$router.push({ name: 'ShopSelection' }); // Corrected route name based on previous discussion
      } catch (err) {
        alert('Registration failed.'); // Basic error handling
        console.error('Registration error:', err); // Log error for debugging
      }
    }
  }
};
</script>


