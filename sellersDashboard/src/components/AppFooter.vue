<template>
  <footer class="bg-white border-t border-green-200 shadow-inner">
    <div class="max-w-7xl mx-auto px-4 py-6 flex flex-col sm:flex-row justify-between items-center text-sm text-gray-600">
      <p>© {{ new Date().getFullYear() }} Seller’s Dash. All rights reserved.</p>
      <nav class="flex space-x-4 mt-2 sm:mt-0">
        <a href="#features" class="hover:text-green-600 transition">Features</a>
        <router-link to="/privacy" class="hover:text-green-600 transition">Privacy</router-link>
        <router-link to="/terms" class="hover:text-green-600 transition">Terms</router-link>
      </nav>
    </div>
  </footer>
</template>

<script setup>
// no script needed
</script>

<style scoped>
/* nothing extra here */
</style>
