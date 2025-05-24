<!-- src/layouts/AppLayout.vue -->
<template>
  <div class="flex h-screen bg-gray-100 overflow-hidden">
    <!-- Backdrop when sidebar is open on mobile -->
    <transition name="fade">
      <div
        v-if="showSidebar"
        class="fixed inset-0 bg-black bg-opacity-40 z-20 md:hidden"
        @click="showSidebar = false"
      />
    </transition>

    <!-- Sidebar -->
    <Sidebar
      :class="[
        'fixed z-30 inset-y-0 left-0 w-64 bg-white border-r border-gray-200 transform transition-transform duration-200 ease-in-out',
        showSidebar ? 'translate-x-0' : '-translate-x-full',
        'md:translate-x-0 md:static md:flex'
      ]"
    />

    <!-- Main content -->
    <div class="flex-1 flex flex-col">
      <Navbar @toggle-sidebar="showSidebar = !showSidebar" />
      <main class="flex-1 overflow-auto p-6 bg-gray-50">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Navbar from '@/components/Navbar.vue'

const showSidebar = ref(false)
</script>
