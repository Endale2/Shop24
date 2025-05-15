<template>
  <div class="flex flex-col min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-800 dark:text-gray-200">

    <header class="bg-gray-800 dark:bg-gray-900 text-white p-4 shadow-md">
      <div class="container mx-auto">
        <h1 class="text-2xl font-semibold text-center">Welcome to Admin Portal</h1>
        </div>
    </header>

    <main class="flex-grow container mx-auto p-4 sm:p-6 md:p-8 overflow-y-auto">
      <router-view />
    </main>

    <footer class="bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-300 p-4 text-center shadow-inner mt-auto">
      <div class="container mx-auto">
        © 2025 Company. All rights reserved.
      </div>
    </footer>

  </div>
</template>

<script>
// No script needed for a simple layout component unless you add logic later
export default {
  name: 'BasicLayout', // Give it a descriptive name
};
</script>

<style scoped>


</style>