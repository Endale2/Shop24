<template>
  <footer class="bg-gray-900 text-gray-300 pt-16 pb-10 shadow-inner border-t border-gray-800"> <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-8 gap-y-12 mb-14"> <div>
          <h3 class="text-2xl font-extrabold text-white mb-5 tracking-wide leading-tight"> {{ shop.current?.name || 'Your Shop Name' }}
          </h3>
          <p class="text-sm text-gray-400 leading-relaxed max-w-xs"> {{ shop.current?.short_description || 'Your go-to place for amazing products. Discover quality and style in every item.' }}
          </p>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-5 tracking-tight border-b border-gray-700 pb-2">Quick Links</h3> <nav class="space-y-3"> <router-link to="/products" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">All Products</router-link>
            <router-link to="/categories" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">Categories</router-link>
            <router-link to="/about-us" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">About Us</router-link>
            <router-link to="/contact" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">Contact</router-link>
          </nav>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-5 tracking-tight border-b border-gray-700 pb-2">Customer Service</h3>
          <nav class="space-y-3">
            <router-link to="/faq" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">FAQ</router-link>
            <router-link to="/shipping" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">Shipping & Returns</router-link>
            <router-link to="/privacy" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">Privacy Policy</router-link>
            <router-link to="/terms" class="text-base text-gray-400 hover:text-green-400 transition-colors duration-200 block">Terms of Service</router-link> </nav>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-5 tracking-tight border-b border-gray-700 pb-2">Stay Connected</h3>
          <p class="text-sm text-gray-400 mb-4 leading-normal">Subscribe to our newsletter for exclusive updates, new product launches, and special offers.</p>
          <form @submit.prevent="handleNewsletterSubscription" class="flex mb-6 rounded-lg overflow-hidden border border-gray-700 focus-within:ring-2 focus-within:ring-green-500 focus-within:border-green-500 transition-all duration-200"> <input
              type="email"
              v-model="newsletterEmail"
              placeholder="Your email address"
              class="flex-1 px-4 py-3 text-sm text-gray-800 bg-gray-100 placeholder-gray-500 outline-none" aria-label="Your email for newsletter subscription"
            />
            <button
              type="submit"
              class="bg-green-600 hover:bg-green-700 text-white px-6 py-3 text-sm font-medium shadow-md hover:shadow-lg transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-gray-900" >
              Subscribe
            </button>
          </form>
          <div class="flex space-x-5 mt-6"> <a href="#" class="text-gray-400 hover:text-green-400 transition-transform transform hover:scale-110 duration-200" aria-label="Facebook">
              <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 24 24"> <path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.988C18.343 21.128 22 16.991 22 12z" clip-rule="evenodd" />
              </svg>
            </a>
            <a href="#" class="text-gray-400 hover:text-green-400 transition-transform transform hover:scale-110 duration-200" aria-label="Twitter">
              <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8.29 20.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.007-.532A8.318 8.318 0 0022 5.92a8.19 8.19 0 01-2.357.646 4.124 4.124 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.288 4.109 4.109 0 001.27 5.477A4.072 4.072 0 014 9.6c-.054.02-.108.04-.163.06 0 4.47 3.193 8.18 7.465 9.041a4.093 4.093 0 01-1.853.07 4.107 4.107 0 003.834 2.85A8.204 8.204 0 013 18.005v.5c0 2.209 1.251 4.12 3.05 5.172A11.64 11.64 0 008.29 20.25z" />
              </svg>
            </a>
            <a href="#" class="text-gray-400 hover:text-green-400 transition-transform transform hover:scale-110 duration-200" aria-label="Instagram">
              <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M12.315 2c2.43 0 2.723.01 3.693.047 3.25.12 4.79 1.48 4.908 4.908.037.97.047 1.263.047 3.693v.001c0 2.43-.01 2.723-.047 3.693-.12 3.25-1.48 4.79-4.908 4.908-.037-.97-.047-1.263-.047-3.693v-.001c0-2.43.01-2.723.047-3.693.12-3.25 1.48-4.79 4.908-4.908.97-.037 1.263-.047 3.693-.047H12.315zM12 7c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5zm0 8c-1.656 0-3-1.344-3-3s1.344-3 3-3 3 1.344 3 3-1.344 3-3 3zm5.426-11.458a1.226 1.226 0 011.226 1.226 1.226 1.226 0 01-1.226 1.226A1.226 1.226 0 0116.2 4.542a1.226 1.226 0 011.226-1.226z" clip-rule="evenodd" />
              </svg>
            </a>
          </div>
        </div>
      </div>

      <div class="border-t border-gray-700 pt-8 mt-8 text-center text-sm text-gray-500 flex flex-col sm:flex-row justify-between items-center"> <p class="mb-2 sm:mb-0">&copy; {{ year }} <span class="font-semibold">{{ shop.current?.name || 'Your Shop Name' }}</span>. All rights reserved.</p>
        <p class="mt-1 sm:mt-0">Built with <span class="text-green-400 font-semibold">Passion</span> by <a href="#" class="hover:text-green-400 transition-colors duration-200 underline-offset-2 hover:underline">Your Platform</a></p> </div>
    </div>
  </footer>
</template>

<script setup>
import { useShopStore } from '@/stores/shop';
import { ref } from 'vue';

const shop = useShopStore(); // Access the whole store instance
const year = new Date().getFullYear();
const newsletterEmail = ref('');

const handleNewsletterSubscription = () => {
  if (newsletterEmail.value) {
    // In a real application, you would send this to your backend API
    // For demonstration, a simple alert is used.
    alert(`Thank you for subscribing, ${newsletterEmail.value}!`);
    newsletterEmail.value = '';
  }
};
</script>

<style scoped>
/* No additional custom styles needed as Tailwind handles everything. */
</style>