<template>
  <div class="bg-white text-gray-800 antialiased">
    <header class="w-full bg-white/80 backdrop-blur-md sticky top-0 z-50 border-b border-gray-200">
      <div class="max-w-7xl mx-auto flex justify-between items-center p-4">
        <h1 class="text-2xl font-bold text-gray-900">Seller's Dash</h1>
        <nav class="flex items-center space-x-4">
          <router-link to="/login" class="text-gray-600 font-medium hover:text-green-600 transition-colors">
            Login
          </router-link>
          <router-link to="/register" class="px-5 py-2 bg-green-600 text-white font-semibold rounded-lg shadow-sm hover:bg-green-700 transition-colors">
            Get Started
          </router-link>
        </nav>
      </div>
    </header>

    <main class="relative">
      <section class="text-center py-24 sm:py-32 px-4 bg-gray-50">
        <div class="max-w-4xl mx-auto">
          <h2 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold text-gray-900 leading-tight mb-6">
            Your All-In-One E-commerce Platform
          </h2>
          <p class="text-lg sm:text-xl text-gray-600 max-w-2xl mx-auto mb-10">
            Seller's Dash provides the tools you need to build, manage, and grow your online business. Focus on your brand; we'll handle the rest.
          </p>
          <div class="flex justify-center space-x-4">
            <router-link to="/register" class="px-8 py-3 bg-green-600 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-green-700 transition-transform transform hover:scale-105">
              Start for Free
            </router-link>
            <router-link to="/features" class="px-8 py-3 bg-white text-gray-700 text-lg font-semibold rounded-lg border border-gray-300 shadow-sm hover:bg-gray-100 transition-colors">
              Learn More
            </router-link>
          </div>
        </div>
      </section>

      <section id="features" class="py-20 sm:py-24 px-4">
        <div class="max-w-7xl mx-auto">
          <div class="text-center mb-16">
            <h3 class="text-3xl sm:text-4xl font-bold text-gray-900">
              Everything you need to succeed
            </h3>
            <p class="mt-4 text-lg text-gray-600">
              From storefront to checkout, we've got you covered.
            </p>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-8 text-center">
            <div class="p-8 bg-white rounded-xl border border-gray-200">
              <div class="flex items-center justify-center h-12 w-12 rounded-lg bg-green-100 text-green-600 mx-auto mb-5">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
                </svg>
              </div>
              <h4 class="text-xl font-semibold text-gray-900 mb-2">Beautiful Storefronts</h4>
              <p class="text-gray-600">
                Create a stunning online store with customizable themes that reflect your brand.
              </p>
            </div>
            <div class="p-8 bg-white rounded-xl border border-gray-200">
              <div class="flex items-center justify-center h-12 w-12 rounded-lg bg-green-100 text-green-600 mx-auto mb-5">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                </svg>
              </div>
              <h4 class="text-xl font-semibold text-gray-900 mb-2">Integrated Payments</h4>
              <p class="text-gray-600">
                Securely accept credit cards and other popular payment methods from day one.
              </p>
            </div>
            <div class="p-8 bg-white rounded-xl border border-gray-200">
              <div class="flex items-center justify-center h-12 w-12 rounded-lg bg-green-100 text-green-600 mx-auto mb-5">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <h4 class="text-xl font-semibold text-gray-900 mb-2">Powerful Analytics</h4>
              <p class="text-gray-600">
                Gain insights into your sales, customers, and traffic to make smarter business decisions.
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="bg-gray-800 text-gray-400">
      <div class="max-w-7xl mx-auto py-12 px-4 text-center">
        <p>&copy; {{ new Date().getFullYear() }} Seller's Dash. All rights reserved.</p>
        <div class="mt-4 flex justify-center space-x-6">
          <a href="#" class="hover:underline">Privacy Policy</a>
          <a href="#" class="hover:underline">Terms of Service</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
// No script logic needed for this static landing page.
</script>

<style scoped>
/* All styles are handled by Tailwind CSS classes in the template. */
/* The default sans-serif font from Tailwind is used for a clean, professional look. */
</style>