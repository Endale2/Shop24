<template>
  <nav class="flex items-center space-x-2 text-sm text-gray-500 mb-6">
    <button @click="$router.back()" class="text-gray-700 hover:text-black flex items-center gap-1 font-medium text-xs mr-2">
      <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
      Back
    </button>
    <span class="font-semibold text-black">Home</span>
    <span>/</span>
    <span class="text-gray-900 font-medium">Shop Selection</span>
  </nav>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full bg-white rounded-lg shadow-md p-8">
      <h1 class="text-2xl font-bold text-center mb-6">Welcome to DRPS Storefront</h1>
      <p class="text-gray-600 text-center mb-8">
        Please select a shop to continue:
      </p>
      
      <div class="space-y-4">
        <router-link
          to="/"
          class="block w-full p-4 border border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-colors"
        >
          <h2 class="font-semibold text-lg">Sophia Boutique</h2>
          <p class="text-gray-600 text-sm">Fashion and accessories</p>
        </router-link>
        
        <!-- Add more shops here as needed -->
        <div class="text-center text-gray-500 text-sm mt-6">
          <p>More shops coming soon...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// This page allows users to select which shop they want to visit
</script> 