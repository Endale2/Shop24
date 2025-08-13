<template>
  <div class="flex h-screen bg-gray-50 overflow-hidden font-sans antialiased">
    <!-- Mobile backdrop overlay -->
    <transition name="fade">
      <div
        v-if="showSidebar"
        class="fixed inset-0 bg-black/20 backdrop-blur-sm z-40 md:hidden"
        @click="showSidebar = false"
      />
    </transition>

    <!-- Sidebar -->
    <aside
      role="dialog"
      aria-modal="true"
      :class="[
        'fixed inset-y-0 left-0 w-64 bg-white border-r border-gray-200 shadow-sm z-50 transform transition-transform duration-300 ease-in-out',
        showSidebar ? 'translate-x-0' : '-translate-x-full',
        'md:static md:translate-x-0 md:shadow-none'
      ]"
    >
      <Sidebar @close-mobile-sidebar="closeMobileSidebar" />
    </aside>

    <!-- Main content area -->
    <div class="flex-1 flex flex-col overflow-hidden bg-gray-50">
      <!-- Top navigation bar -->
      <Navbar @toggle-sidebar="showSidebar = !showSidebar" />

      <!-- Main content -->
      <main class="flex-1 overflow-auto">
        <div class="min-h-full">
          <router-view />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Navbar  from '@/components/Navbar.vue'

const showSidebar = ref(false)

// Function to close mobile sidebar
const closeMobileSidebar = () => {
  showSidebar.value = false
}
</script>

<style scoped>
/* Fade transition for backdrop */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* Custom scrollbar styling for main content */
main::-webkit-scrollbar {
  width: 6px;
}

main::-webkit-scrollbar-track {
  background: transparent;
}

main::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;
}

main::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Ensure proper height calculations */
.h-screen {
  height: 100vh;
  height: 100dvh; /* Dynamic viewport height for mobile */
}

/* Smooth transitions for all interactive elements */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Focus styles for accessibility */
*:focus {
  outline: 2px solid #10b981;
  outline-offset: 2px;
}

/* Remove focus outline for mouse users */
*:focus:not(:focus-visible) {
  outline: none;
}
</style>