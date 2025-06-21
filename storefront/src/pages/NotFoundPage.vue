<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-gray-50 to-red-50 text-gray-900 antialiased font-sans px-4 py-16 text-center">
    <div class="max-w-2xl w-full bg-white rounded-2xl shadow-xl p-8 sm:p-12 border border-gray-100 transform transition-all duration-300 hover:scale-[1.01] hover:shadow-2xl">
      <p class="text-7xl sm:text-8xl font-extrabold text-red-500 mb-6 drop-shadow-lg animate-bounce-slow">
        404
      </p>

      <h1 class="text-4xl sm:text-5xl font-bold text-gray-800 mb-4 tracking-tight leading-tight">
        Oops! Page Not Found
      </h1>

      <p class="text-lg text-gray-600 mb-8 max-w-prose mx-auto">
        It looks like the page you're trying to reach doesn't exist or has been moved. Don't worry, we'll help you find your way back!
      </p>

      <div class="flex flex-col sm:flex-row justify-center items-center gap-4">
        <router-link
          to="/"
          class="inline-flex items-center px-8 py-3 bg-green-600 text-white text-lg font-semibold rounded-full shadow-lg hover:bg-green-700 transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl focus:outline-none focus:ring-4 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-white"
        >
          <svg class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m0 0l7 7m-17 0v10a1 1 0 001 1h3m-10 4h10a2 2 0 002-2v-6H4v6a2 2 0 002 2z" />
          </svg>
          Go to Homepage
        </router-link>

        <router-link
          to="/products"
          class="inline-flex items-center px-8 py-3 border-2 border-gray-300 text-gray-700 text-lg font-semibold rounded-full bg-white shadow-md hover:border-gray-400 hover:text-gray-800 transition-all duration-300 transform hover:-translate-y-1 hover:shadow-lg focus:outline-none focus:ring-4 focus:ring-gray-300 focus:ring-offset-2 focus:ring-offset-white"
        >
          <svg class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
          </svg>
          Browse Products
        </router-link>
      </div>

      <div class="mt-10 pt-6 border-t border-gray-200">
        <p class="text-md text-gray-500 mb-4">
          Looking for something specific? Try searching:
        </p>
        <form @submit.prevent="handleSearch" class="flex max-w-md mx-auto rounded-full overflow-hidden shadow-md focus-within:ring-2 focus-within:ring-green-500 focus-within:border-green-500 transition-all duration-200">
          <input
            type="text"
            v-model="searchQuery"
            placeholder="e.g., 'green tea' or 'latest shirts'"
            class="flex-1 px-5 py-3 text-gray-800 bg-gray-100 placeholder-gray-500 outline-none border-none"
            aria-label="Search the shop"
          />
          <button
            type="submit"
            class="bg-green-500 hover:bg-green-600 text-white px-6 py-3 font-semibold transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-white"
            aria-label="Submit search"
          >
            Search
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const searchQuery = ref('');
const router = useRouter();

const handleSearch = () => {
  if (searchQuery.value.trim()) {
    router.push({ path: '/products', query: { q: searchQuery.value.trim() } });
  }
};
</script>

<style scoped>
/* Optional: A subtle animation for the 404 text */
@keyframes bounce-slow {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px); /* Subtle bounce */
  }
}

.animate-bounce-slow {
  animation: bounce-slow 4s infinite ease-in-out;
}
</style>