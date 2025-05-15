<template>
  <div>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Dashboard</h1>
      </div>

    <div class="mb-6">
       <h2 class="text-xl font-semibold text-gray-800 dark:text-white">Dashboard Overview</h2>
       <p class="mt-1 text-sm text-gray-600 dark:text-gray-300">Get a quick look at your key metrics.</p>
     </div>


    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-blue-500 dark:bg-blue-600 text-white p-3 rounded-md mr-4">
             <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10m0-10l-8 4m8-4V4"></path></svg>
           </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Products</p>
             <div v-if="loading" class="mt-1 text-lg font-semibold text-gray-900 dark:text-white">
               Loading...
             </div>
             <div v-else-if="error" class="mt-1 text-lg font-semibold text-red-500">
               Error
             </div>
             <div v-else class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">
                {{ productCount }}
             </div>
          </div>
        </div>
         <p v-if="error" class="mt-2 text-xs text-red-500">{{ error }}</p>
      </div>

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-green-500 dark:bg-green-600 text-white p-3 rounded-md mr-4">
              <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.292M12 14c0 1.657-1.594 3-3.5 3S5 15.657 5 14s1.594-3 3.5-3 3.5 1.343 3.5 3zm0 0h6c0 1.657 1.594 3 3.5 3s3.5-1.343 3.5-3s-1.594-3-3.5-3-3.5 1.343-3.5 3z"></path></svg>
           </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Customers</p>
             <div class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">
                ...
             </div>
          </div>
        </div>
      </div>

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-yellow-500 dark:bg-yellow-600 text-white p-3 rounded-md mr-4">
              <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M12 16h.01"></path></svg>
           </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Orders</p>
             <div class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">
                ...
             </div>
          </div>
        </div>
      </div>

       <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-purple-500 dark:bg-purple-600 text-white p-3 rounded-md mr-4">
               <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2-1.343-2-3-2zM21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
           </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Revenue</p>
             <div class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">
                ...
             </div>
          </div>
        </div>
      </div>

      </div>

    <div class="mt-8 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
       <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Recent Activity (Placeholder)</h3>
       <p class="text-gray-600 dark:text-gray-300">This area could contain tables, charts, or lists of recent actions.</p>
        </div>


  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DashboardPage',
  data() {
    return {
      productCount: 0,
      loading: true,
      error: null,
    };
  },
  async created() {
    // Simulate a slight delay to show loading state better if needed
    // await new Promise(resolve => setTimeout(resolve, 500));

    try {
      // Ensure this endpoint is correct and returns { count: number }
      const res = await axios.get('/admin/products/count');
      this.productCount = res.data.count;
    } catch (err) {
      console.error('Failed to fetch product count:', err);
      this.error = 'Failed to load count'; // More concise error for the card
       // Optionally set a user-friendly error message here or handle differently
    } finally {
      this.loading = false;
    }
  },
};
</script>

<style scoped>
/* No specific styles needed if using Tailwind for everything */
/* Ensure .container styles are removed from here if they existed */
/* .container {
    max-width: ...;
} */
</style>