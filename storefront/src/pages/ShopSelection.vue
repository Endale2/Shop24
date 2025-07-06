<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full bg-white rounded-lg shadow-md p-8">
      <h1 class="text-2xl font-bold text-center mb-6">Welcome to DRPS Storefront</h1>
      <p class="text-gray-600 text-center mb-8">
        Please select a shop to continue:
      </p>
      
      <div class="space-y-4">
        <router-link
          to="/shops/sophiaboutique"
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