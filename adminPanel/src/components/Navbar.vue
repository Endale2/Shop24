<template>
  <nav class="bg-white dark:bg-gray-800 shadow-md p-4 flex justify-between items-center z-20">
    <div class="flex items-center">
      <button
        @click="$emit('toggle-sidebar')"
        class="text-gray-500 dark:text-gray-300 focus:outline-none focus:text-gray-600 dark:focus:text-gray-400 lg:hidden mr-4"
      >
        <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
      </button>

      <h1 class="text-xl font-semibold text-gray-800 dark:text-white hidden sm:block">Dashboard</h1>
       </div>

    <div class="flex items-center">
      <div v-if="user" class="text-gray-800 dark:text-gray-200 mr-4 hidden md:block">
         Hello, <span class="font-semibold">{{ user.name }}</span>
        <span class="text-sm text-gray-500 dark:text-gray-400 ml-1">({{ user.role }})</span>
      </div>


      <button
        @click="logout"
        class="px-3 py-1 bg-red-600 hover:bg-red-700 text-white rounded-md transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
      >
        Logout
      </button>
    </div>
  </nav>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    user: Object, // Expecting the user object
  },
   emits: ['toggle-sidebar'], // Declare emitted event
  methods: {
    async logout() {
      try {
        // Assuming your backend returns a successful response on logout
        await axios.post('/auth/admin/logout');
        console.log('Logged out successfully');
      } catch (error) {
        console.error('Logout failed:', error);
        // Even if the request fails, clear local session/redirect
        // depending on your auth strategy
      } finally {
         // Always redirect to login after attempting logout
         // Ensure your router is configured and available
         this.$router.replace({ name: 'Login' });
      }
    }
  }
};
</script>

<style scoped>
/* No specific styles needed if using Tailwind */
</style>