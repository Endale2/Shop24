<template>
  <footer class="bg-gradient-to-r from-green-200 to-amber-100 border-t border-green-300 shadow-xl py-10">
    <div class="max-w-7xl mx-auto px-4 grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-8 text-gray-800">

      <!-- Column 1: Company Info -->
      <div class="col-span-1">
        <div class="flex items-center space-x-3 mb-4">
          <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-500/30">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-green-900">
            Shop<span class="text-amber-600">24</span>
          </h3>
        </div>
        <p class="text-sm leading-relaxed mb-4">
          Empowering sellers with intuitive tools and insights to boost their online business presence and sales.
        </p>
        <p class="text-xs text-gray-700">
          © {{ new Date().getFullYear() }} Shop24. All rights reserved.
        </p>
      </div>

      <!-- Column 2: Quick Links -->
      <div class="col-span-1">
        <h3 class="text-lg font-semibold text-green-800 mb-4">Quick Links</h3>
        <nav class="flex flex-col space-y-2">
          <router-link to="/" class="text-sm hover:text-green-900 transition-colors duration-200">Home</router-link>
          <a href="#features" class="text-sm hover:text-green-900 transition-colors duration-200">Features</a>
          <router-link to="/pricing" class="text-sm hover:text-green-900 transition-colors duration-200">Pricing</router-link>
          <router-link to="/contact" class="text-sm hover:text-green-900 transition-colors duration-200">Contact Us</router-link>
        </nav>
      </div>

      <!-- Column 3: Legal & Support -->
      <div class="col-span-1">
        <h3 class="text-lg font-semibold text-green-800 mb-4">Legal & Support</h3>
        <nav class="flex flex-col space-y-2">
          <router-link to="/privacy" class="text-sm hover:text-green-900 transition-colors duration-200">Privacy Policy</router-link>
          <router-link to="/terms" class="text-sm hover:text-green-900 transition-colors duration-200">Terms of Service</router-link>
          <router-link to="/faq" class="text-sm hover:text-green-900 transition-colors duration-200">FAQ</router-link>
        </nav>
      </div>

      <!-- Column 4: Social Media -->
      <div class="col-span-1">
        <h3 class="text-lg font-semibold text-green-800 mb-4">Connect With Us</h3>
        <div class="flex space-x-4">
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33V22C17.34 21.128 22 16.991 22 12z" clip-rule="evenodd" />
            </svg>
          </a>
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path d="M8.29 20.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0022 5.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.072 4.072 0 012.8 9.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 012 18.407a11.616 11.616 0 006.29 1.84" />
            </svg>
          </a>
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path fill-rule="evenodd" d="M12.017 0C5.396 0 .029 5.367.029 11.987c0 5.079 3.158 9.417 7.618 11.174-.105-.949-.199-2.403.041-3.439.219-.937 1.406-5.957 1.406-5.957s-.359-.72-.359-1.781c0-1.663.967-2.911 2.168-2.911 1.024 0 1.518.769 1.518 1.688 0 1.029-.653 2.567-.992 3.992-.285 1.193.6 2.165 1.775 2.165 2.128 0 3.768-2.245 3.768-5.487 0-2.861-2.063-4.869-5.008-4.869-3.41 0-5.409 2.562-5.409 5.199 0 1.033.394 2.143.889 2.741.097.118.112.221.083.343-.09.375-.293 1.199-.334 1.363-.053.225-.172.271-.402.165-1.495-.69-2.433-2.878-2.433-4.646 0-3.776 2.748-7.252 7.92-7.252 4.158 0 7.392 2.967 7.392 6.923 0 4.135-2.607 7.462-6.233 7.462-1.214 0-2.357-.629-2.75-1.378l-.748 2.853c-.271 1.043-1.002 2.35-1.492 3.146C9.57 23.812 10.763 24.009 12.017 24.009c6.624 0 11.99-5.367 11.99-11.988C24.007 5.367 18.641.001.012.001z" clip-rule="evenodd" />
            </svg>
          </a>
        </div>
      </div>

    </div>
  </footer>
</template>

<script setup>
// No script needed for this component.
</script>

<style scoped>
/* No additional custom styles are necessary as Tailwind CSS handles the styling. */
</style>
