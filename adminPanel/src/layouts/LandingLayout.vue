<template>
    <div>
      <header class="p-4 bg-blue-600 text-white text-center">
        <h1 class="text-2xl">Welcome to Admin Portal</h1>
      </header>
      <main class="flex-grow p-8">
        <router-view />
      </main>
      <footer class="p-4 bg-gray-200 text-center">
        © 2025 Company
      </footer>
    </div>
  </template>