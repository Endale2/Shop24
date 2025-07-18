<template>
  <div class="flex h-screen bg-amber-50 overflow-hidden font-sans antialiased">
    <transition name="fade">
      <div
        v-if="showSidebar"
        class="fixed inset-0 bg-white/30 backdrop-blur-none z-40 md:hidden"
        @click="showSidebar = false"
      />
    </transition>

    <aside
      role="dialog"
      aria-modal="true"
      :class="[
        'fixed inset-y-0 left-0 w-64 bg-white border-r border-green-200 z-50 transform transition-transform duration-300 ease-in-out',
        showSidebar ? 'translate-x-0' : '-translate-x-full',
        'md:static md:translate-x-0'
      ]"
    >
      <Sidebar />
    </aside>

    <div class="flex-1 flex flex-col overflow-hidden">
      <Navbar @toggle-sidebar="showSidebar = !showSidebar" />

      <main class="flex-1 overflow-auto p-4 md:p-6 bg-amber-50">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Navbar  from '@/components/Navbar.vue'

const showSidebar = ref(false)
</script>

<style scoped>
/* fade transition for backdrop */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>