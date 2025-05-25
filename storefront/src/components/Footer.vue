<template>
  <footer class="bg-gray-900 text-gray-300 py-10 md:py-12">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-8">

      <div>
        <h3 class="text-white text-lg font-bold mb-4">
          {{ shop.current?.name || 'My Storefront' }}
        </h3>
        <p class="text-sm leading-relaxed mb-4">
          Discover the latest trends and unique products. We're committed to bringing you quality and style directly to your doorstep.
        </p>
        <router-link to="/about" class="text-indigo-400 hover:text-indigo-300 text-sm transition duration-300 ease-in-out">
          Learn More &rarr;
        </router-link>
      </div>

      <div>
        <h3 class="text-white text-lg font-bold mb-4">Quick Links</h3>
        <ul class="space-y-2">
          <li><router-link to="/products" class="hover:text-white transition duration-300 ease-in-out">Products</router-link></li>
          <li><router-link to="/categories" class="hover:text-white transition duration-300 ease-in-out">Categories</router-link></li>
          <li><router-link to="/faq" class="hover:text-white transition duration-300 ease-in-out">FAQ</router-link></li>
          <li><router-link to="/returns" class="hover:text-white transition duration-300 ease-in-out">Returns</router-link></li>
          <li><router-link to="/privacy" class="hover:text-white transition duration-300 ease-in-out">Privacy Policy</router-link></li>
        </ul>
      </div>

      <div>
        <h3 class="text-white text-lg font-bold mb-4">Contact Us</h3>
        <address class="not-italic text-sm space-y-2 mb-4">
          <p><a href="mailto:info@yourstore.com" class="hover:text-white transition duration-300 ease-in-out">info@yourstore.com</a></p>
          <p>123 Commerce St, Suite 456<br>
             Retail City, ST 98765</p>
          <p><a href="tel:+1-555-123-4567" class="hover:text-white transition duration-300 ease-in-out">+1-555-123-4567</a></p>
        </address>
        <div class="flex space-x-4">
          <a href="#" class="text-gray-400 hover:text-white transition duration-300 ease-in-out" aria-label="Facebook">
            <i class="fab fa-facebook-f text-xl"></i> </a>
          <a href="#" class="text-gray-400 hover:text-white transition duration-300 ease-in-out" aria-label="Twitter">
            <i class="fab fa-twitter text-xl"></i> </a>
          <a href="#" class="text-gray-400 hover:text-white transition duration-300 ease-in-out" aria-label="Instagram">
            <i class="fab fa-instagram text-xl"></i> </a>
        </div>
      </div>

      <div class="md:col-span-full lg:col-span-1"> <h3 class="text-white text-lg font-bold mb-4">Stay Connected</h3>
        <p class="text-sm mb-4">Subscribe to our newsletter for exclusive offers and updates!</p>
        <form @submit.prevent="subscribeNewsletter">
          <input
            type="email"
            placeholder="Enter your email"
            class="w-full p-3 mb-3 bg-gray-700 border border-gray-600 rounded-md text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          />
          <button
            type="submit"
            class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 rounded-md transition duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-900"
          >
            Subscribe
          </button>
        </form>
      </div>

    </div>

    <div class="border-t border-gray-700 mt-10 pt-6 text-center text-gray-500 text-sm">
      <p>&copy; {{ currentYear }} {{ shop.current?.name || 'My Storefront' }}. All rights reserved.</p>
      <div class="mt-4 flex justify-center space-x-4">
        <img src="https://via.placeholder.com/40x25?text=Visa" alt="Visa" class="h-6 opacity-75">
        <img src="https://via.placeholder.com/40x25?text=Mastercard" alt="Mastercard" class="h-6 opacity-75">
        <img src="https://via.placeholder.com/40x25?text=Amex" alt="American Express" class="h-6 opacity-75">
        <img src="https://via.placeholder.com/40x25?text=PayPal" alt="PayPal" class="h-6 opacity-75">
      </div>
    </div>
  </footer>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import { RouterLink } from 'vue-router' // Import RouterLink for proper usage

const shop = useShopStore()
const currentYear = new Date().getFullYear()

// Placeholder for newsletter subscription logic
const subscribeNewsletter = () => {
  alert('Newsletter subscribed! (This is a placeholder action)')
  // In a real application, you'd send the email to your backend
}
</script>

<style scoped>
/* You might include custom CSS for actual icon sprites or specific SVG sizing */
</style>