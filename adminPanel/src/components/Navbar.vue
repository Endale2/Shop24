<template>
  <nav class="bg-white dark:bg-gray-800 shadow-md p-4 flex justify-between items-center z-20 flex-shrink-0">
    <!-- Brand/Logo on left -->
    <div class="flex items-center">
      <div class="flex items-center justify-center h-12 w-12 bg-gradient-to-r from-purple-600 to-indigo-600 text-xl font-bold rounded-md">
        <span class="uppercase tracking-wide text-white">AP</span>
      </div>
      <span class="ml-3 text-lg font-semibold text-gray-800 dark:text-gray-100 hidden sm:block">Admin Panel</span>
    </div>

    <!-- User & Actions -->
    <div class="flex items-center">
      <div v-if="user" class="text-gray-800 dark:text-gray-200 mr-4 hidden md:block">
         Hello, <span class="font-semibold">{{ user.name }}</span>
        <span class="text-sm text-gray-500 dark:text-gray-400 ml-1">({{ user.role }})</span>
      </div>

      <button
        @click="logout"
        class="px-3 py-1 bg-red-600 hover:bg-red-700 text-white rounded-md transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50 dark:bg-red-700 dark:hover:bg-red-600 dark:focus:ring-red-600"
      >
        Logout
      </button>
    </div>
  </nav>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Navbar',
  props: {
    user: Object,
  },
  methods: {
    async logout() {
      try {
        await axios.post('/auth/admin/logout', null, { withCredentials: true });
        console.log('Logged out successfully');
      } catch (error) {
        console.error('Logout failed:', error);
      } finally {
         this.$router.replace({ name: 'Login' });
      }
    }
  }
};
</script>

<style scoped>
/* No specific styles needed beyond Tailwind */
</style>
