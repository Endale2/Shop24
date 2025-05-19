<template>
  <div class="container mx-auto py-12 text-center">
    <h1 class="text-4xl font-bold mb-4">Welcome to Seller's Dash</h1>
    <p class="text-lg text-gray-700 mb-8">Your platform to create and manage e-commerce shops.</p>

    <div class="space-x-4">
      <button
        @click="goToRegister"
        class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Get Started
      </button>
      <button
        @click="goToLogin"
        class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Login
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LandingPage',
  methods: {
    goToRegister() {
      this.$router.push({ name: 'Register' });
    },
    goToLogin() {
      this.$router.push({ name: 'Login' });
    }
  }
};
</script>


