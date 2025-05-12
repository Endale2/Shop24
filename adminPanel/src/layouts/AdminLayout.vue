<template>
    <div class="flex h-screen">
      <Sidebar />
      <div class="flex-1 flex flex-col">
        <Navbar :user="user" />
        <section class="bg-gray-100 p-4 border-b">
          <!-- Display additional user info -->
          <div class="text-right text-sm text-gray-600">
            <span>ID: {{ user.id }}</span>
            <span class="mx-2">|</span>
            <span>Role: {{ user.role }}</span>
          </div>
        </section>
        <main class="p-6 flex-grow overflow-auto">
          <router-view />
        </main>
      </div>
    </div>
  </template>
  
  <script>
  import Sidebar from '../components/Sidebar.vue';
  import Navbar from '../components/Navbar.vue';
  import auth from '../services/auth';
  
  export default {
    components: { Sidebar, Navbar },
    data() { return { user: null }; },
    async created() { this.user = await auth.getUser(); }
  };
  </script>