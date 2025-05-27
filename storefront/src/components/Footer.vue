<template>
  <footer class="bg-slate-900 text-slate-300 pt-12 pb-8">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-8 mb-8">
        <div>
          <h3 class="text-lg font-semibold text-white mb-4">{{ shop.current?.name || 'Your Shop Name' }}</h3>
          <p class="text-sm text-slate-400 leading-relaxed">
            {{ shop.current?.short_description || 'Your go-to place for amazing products. Discover quality and style in every item.' }}
          </p>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-4">Quick Links</h3>
          <nav class="space-y-2">
            <router-link to="/products" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">All Products</router-link>
            <router-link to="/categories" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">Categories</router-link>
            <router-link to="/about-us" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">About Us</router-link> <router-link to="/contact" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">Contact</router-link> </nav>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-4">Customer Service</h3>
          <nav class="space-y-2">
            <router-link to="/faq" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">FAQ</router-link>
            <router-link to="/shipping" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">Shipping & Returns</router-link>
            <router-link to="/privacy" class="text-sm text-slate-400 hover:text-indigo-400 transition-colors duration-200 block">Privacy Policy</router-link>
          </nav>
        </div>

        <div>
          <h3 class="text-lg font-semibold text-white mb-4">Stay Connected</h3>
          <p class="text-sm text-slate-400 mb-3">Subscribe to our newsletter for updates.</p>
          <form @submit.prevent="handleNewsletterSubscription" class="flex mb-4">
            <input type="email" v-model="newsletterEmail" placeholder="Your email" class="w-full px-3 py-2 text-sm text-slate-800 bg-slate-100 border border-slate-300 rounded-l-md focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            <button type="submit" class="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 text-sm font-medium rounded-r-md transition-colors duration-200">
              Go
            </button>
          </form>
          <div class="flex space-x-4">
            <a href="#" class="text-slate-400 hover:text-indigo-400 transition-colors duration-200" aria-label="Facebook">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.988C18.343 21.128 22 16.991 22 12z" clip-rule="evenodd" />
              </svg>
            </a>
            </div>
        </div>
      </div>

      <div class="border-t border-slate-700 pt-8 mt-8 text-center text-sm text-slate-500">
        <p>&copy; {{ year }} {{ shop.current?.name || 'Your Shop Name' }}. All rights reserved.</p>
        <p class="mt-1">Powered by <a href="#" class="hover:text-indigo-400">Your Platform</a></p> </div>
    </div>
  </footer>
</template>

<script setup>
import { useShopStore } from '@/stores/shop'
import { ref } from 'vue'

const shop = useShopStore() // Access the whole store instance
const year = new Date().getFullYear()
const newsletterEmail = ref('')

const handleNewsletterSubscription = () => {
  if (newsletterEmail.value) {
    alert(`Subscribed with: ${newsletterEmail.value}`) // Replace with actual API call
    newsletterEmail.value = ''
  }
}
</script>