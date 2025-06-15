<template>
  <footer class="bg-gradient-to-r from-green-200 to-amber-100 border-t border-green-300 shadow-xl py-10">
    <div class="max-w-7xl mx-auto px-4 grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-8 text-gray-800">

      <!-- Column 1: Company Info -->
      <div class="col-span-1">
        <h3 class="text-xl font-bold text-green-900 mb-4">Seller’s Dash</h3>
        <p class="text-sm leading-relaxed mb-4">
          Empowering sellers with intuitive tools and insights to boost their online business presence and sales.
        </p>
        <p class="text-xs text-gray-700">
          © {{ new Date().getFullYear() }} Seller’s Dash. All rights reserved.
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

      <!-- Column 4: Social Media / Newsletter (Optional: if you want to add newsletter signup) -->
      <div class="col-span-1">
        <h3 class="text-lg font-semibold text-green-800 mb-4">Connect With Us</h3>
        <div class="flex space-x-4">
          <!-- Social Media Icons (using SVG for simplicity and consistency) -->
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33V22C17.34 21.128 22 16.991 22 12z" clip-rule="evenodd" />
            </svg>
          </a>
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path d="M19.611 12.009c0 2.906-2.357 5.263-5.264 5.263-2.906 0-5.262-2.357-5.262-5.263 0-2.907 2.356-5.263 5.262-5.263 2.907 0 5.264 2.356 5.264 5.263zm-5.264-3.565c-1.309 0-2.368 1.059-2.368 2.368 0 1.309 1.059 2.368 2.368 2.368s2.368-1.059 2.368-2.368c0-1.309-1.059-2.368-2.368-2.368z" />
              <circle cx="18.9" cy="5.8" r="1.4" />
              <path d="M22.213 14.28c-.012-.489-.038-.97-.09-1.442-.053-.473-.131-.937-.24-1.39-.108-.453-.254-.889-.438-1.308-.184-.42-.4-1.054-.432-1.472a4.428 4.428 0 0 0-.256-.81 4.544 4.544 0 0 0-.464-.817c-.201-.274-.424-.52-.676-.738a4.9 4.9 0 0 0-.796-.583c-.272-.19-.57-.34-.89-.45-.32-.11-.66-.19-1.02-.24-.36-.05-.72-.08-1.1-.09-.38-.01-.76-.01-1.15-.01h-.202a6.052 6.052 0 0 0-2.316.425c-.714.212-1.39.52-2.008.918-.62.398-1.18.88-1.68 1.44-.5.56-.94 1.18-1.31 1.87-.37.69-.67 1.45-.9 2.27-.23.82-.38 1.69-.47 2.6-.09.91-.12 1.83-.12 2.76v.22a6.052 6.052 0 0 0 .425 2.316c.212.714.52 1.39.918 2.008.398.62.88 1.18 1.44 1.68.56.5 1.18.94 1.87 1.31.69.37 1.45.67 2.27.9.82.23 1.69.38 2.6.47.91.09 1.83.12 2.76.12h.22c.93 0 1.85-.03 2.76-.12.91-.09 1.78-.24 2.6-.47.82-.23 1.58-.53 2.27-.9.69-.37 1.31-.81 1.87-1.31.56-.5 1.04-1.06 1.44-1.68.398-.62.706-1.294.918-2.008A6.052 6.052 0 0 0 22.213 14.28zM14.35 20.32c-.818.083-1.64.125-2.46.125h-.22c-.93 0-1.85-.03-2.76-.12-.91-.09-1.78-.24-2.6-.47-.82-.23-1.58-.53-2.27-.9-.69-.37-1.31-.81-1.87-1.31-.56-.5-1.04-1.06-1.44-1.68-.398-.62-.706-1.294-.918-2.008a4.428 4.428 0 0 1-.256-.81c-.032-.42-.254-1.054-.438-1.472-.184-.42-.33-.855-.438-1.308-.108-.453-.187-.917-.24-1.39-.053-.473-.078-.954-.09-1.442v-.22c0-.93.03-1.85.12-2.76.09-.91.24-1.78.47-2.6.23-.82.53-1.58.9-2.27.37-.69.81-1.31 1.31-1.87.5-.56 1.06-1.04 1.68-1.44.62-.398 1.294-.706 2.008-.918a4.428 4.428 0 0 1 2.316-.425h.202c.38 0 .76.01 1.15.01.36.01.72.04 1.1.09.36.05.70.13 1.02.24.32.11.62.26.89.45.252.218.475.464.676.738.201.274.363.543.464.817.101.274.19.58.256.81.032.42.254 1.054.438 1.472.184.42.33.855.438 1.308.108.453.187.917.24 1.39.053.473.078.954.09 1.442z" />
            </svg>
          </a>
          <a href="#" class="text-gray-700 hover:text-green-800 transition-colors duration-200">
            <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path d="M20 2H4c-1.104 0-2 .896-2 2v16c0 1.104.896 2 2 2h16c1.104 0 2-.896 2-2V4c0-1.104-.896-2-2-2zm-3.5 8.5h-2v7h-3v-7h-2v-2h7v2z" />
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
