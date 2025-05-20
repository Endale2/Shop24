<template>
  <div class="flex h-screen">
    <!-- Sidebar (hidden on small screens) -->
    <Sidebar />

    <div class="flex-1 flex flex-col">
      <!-- Top navbar -->
      <Navbar />

      <!-- Main content area -->
      <main class="p-4 bg-gray-50 flex-1 overflow-auto">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script>
import Sidebar from '@/components/Sidebar.vue';
import Navbar from '@/components/Navbar.vue';

export default {
  name: 'DashboardLayout',
  components: { Sidebar, Navbar }
};
</script>
