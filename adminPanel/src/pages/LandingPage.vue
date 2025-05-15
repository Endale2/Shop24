<template>
  <div class="container mx-auto p-8 bg-white dark:bg-gray-800 shadow-lg rounded-lg text-gray-800 dark:text-gray-200">
    <div class="text-center">
      <h2 class="text-3xl sm:text-4xl font-bold text-gray-900 dark:text-white mb-6">
        Welcome to the Admin Portal
      </h2>
      <p class="text-lg text-gray-600 dark:text-gray-400 mb-8">
        Access administrative features by logging in.
      </p>
      <router-link
        to="/login"
        class="inline-flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-offset-gray-800 transition-colors duration-150"
      >
        Go to Login Page
      </router-link>
    </div>
  </div>
</template>

<script>
// No script is needed for this static component
export default {
  name: 'WelcomePage', // Give it a descriptive name
};
</script>

<style scoped>

.container {
    max-width: 600px; 
}
</style>