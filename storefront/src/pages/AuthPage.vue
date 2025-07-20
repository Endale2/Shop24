<template>
  <div class="login-container">
    <div class="max-w-md mx-auto text-center">
      <h1 class="text-3xl font-bold text-gray-900 tracking-tight uppercase mb-2">Sign In</h1>
      <p class="text-gray-600">Please use the Sign In page to authenticate with your email via OTP.</p>
      <router-link to="/login" class="inline-block bg-black text-white py-3 px-6 rounded-none font-semibold uppercase tracking-wide hover:bg-gray-800 transition-colors mt-8">Go to Sign In</router-link>
    </div>
  </div>
</template>
<script setup lang="ts">
</script>
<style scoped>
.login-container {
  min-height: 80vh;
  padding: 2rem 0;
  background: #f7f7f9;
}
</style> 