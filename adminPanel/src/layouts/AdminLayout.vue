<template>
  <div class="flex h-screen bg-gray-100 dark:bg-gray-900 overflow-hidden">

    <Sidebar
      :user-role="user ? user.role : ''"
      :is-open="isSidebarOpen"
      @close-sidebar="closeSidebar"
    />

    <div
      v-if="isSidebarOpen && !isLargeScreen"
      class="fixed inset-0 bg-black opacity-50 z-20"
      @click="closeSidebar"
    ></div>

    <div
      class="flex-1 flex flex-col overflow-hidden transition-all duration-300 ease-in-out"
      :class="{ 'ml-48': isSidebarOpen && isLargeScreen }"
    >

      <Navbar
        :user="user"
        @toggle-sidebar="toggleSidebar"
      />

      <main class="flex-grow p-4 sm:p-6 md:p-8 overflow-y-auto custom-scrollbar">
        <div v-if="!user && loadingUser" class="flex items-center justify-center h-full">
          <p class="text-gray-500 dark:text-gray-400">Loading user data...</p>
          </div>
         <div v-else-if="userError" class="flex items-center justify-center h-full">
            <p class="text-red-500 dark:text-red-400">Error loading user data. Please try again.</p>
         </div>
        <router-view v-slot="{ Component }" v-else>
          <transition name="fade-transform" mode="out-in">
            <component :is="Component" :key="$route.path" />
          </transition>
        </router-view>
      </main>

    </div>

  </div>
</template>

<script>
import Sidebar from '../components/Sidebar.vue';
import Navbar from '../components/Navbar.vue';
import auth from '../services/auth'; // Your authentication service

export default {
  name: 'AdminLayout',
  components: {
    Sidebar,
    Navbar,
  },
  data() {
    return {
      user: null,
      userError: null,
      loadingUser: true,
      isSidebarOpen: true, // Default state (consider screen size on mount)
      windowWidth: typeof window !== 'undefined' ? window.innerWidth : 0,
    };
  },
  computed: {
    // Define breakpoints using matchMedia or simply check width against Tailwind defaults (lg: 1024px)
    isLargeScreen() {
       return this.windowWidth >= 1024;
    }
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    closeSidebar() {
      this.isSidebarOpen = false;
    },
    // Handle window resize
    handleResize() {
      this.windowWidth = window.innerWidth;
      // Automatically adjust sidebar state based on breakpoint
      // Only change if it crosses the breakpoint
      if (this.windowWidth < 1024 && this.isSidebarOpen && this.isLargeScreen) {
          this.isSidebarOpen = false; // Close sidebar when shrinking below lg
      } else if (this.windowWidth >= 1024 && !this.isSidebarOpen && !this.isLargeScreen) {
          this.isSidebarOpen = true; // Open sidebar when expanding to lg or above
      }
       // Note: This simple logic doesn't preserve manual toggle state across breakpoints.
       // For more complex needs, track if the state was set by user interaction vs resize.
    }
  },
  async created() {
    // Set initial sidebar state based on screen size
    if (typeof window !== 'undefined') {
       this.windowWidth = window.innerWidth;
       this.isSidebarOpen = this.windowWidth >= 1024; // Open by default on large screens
    }


    // Listen for window resize
    if (typeof window !== 'undefined') {
      window.addEventListener('resize', this.handleResize);
    }


    // Fetch user data
    try {
      this.user = await auth.getUser();
      if (!this.user) {
        console.warn('AdminLayout: User not authenticated or not found, redirecting...');
        // Use nextTick to ensure router is ready if needed, though $router is usually available in created
        this.$router.push('/login');
      }
    } catch (error) {
      console.error('AdminLayout: Error fetching user:', error);
      this.userError = 'Authentication failed.'; // Set error message
      // Optionally redirect to login on fetch error
       this.$router.push('/login');
    } finally {
       this.loadingUser = false;
    }
  },
   beforeUnmount() { // Use beforeUnmount in Vue 3
    // Remove resize listener
     if (typeof window !== 'undefined') {
      window.removeEventListener('resize', this.handleResize);
     }
   }
   // For Vue 2, use beforeDestroy() instead of beforeUnmount()
};
</script>

<style scoped>
/*
  Scoped styles for transitions, overlay, and custom scrollbars.
  Tailwind classes handle most of the layout and appearance.
*/

/* Overlay for small screens */
.fixed.inset-0 {
  /* Handled by Tailwind */
}


/* Fade and Transform Transition for Router View */
.fade-transform-enter-active,
.fade-transform-leave-active {
  transition: opacity 0.3s ease-in-out, transform 0.3s ease-in-out;
}

.fade-transform-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Custom Scrollbar (Optional, feel free to adjust or remove) */
/* Existing scrollbar styles */
.custom-scrollbar::-webkit-scrollbar {
  width: 8px; /* width of the scrollbar */
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1; /* color of the track */
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #888; /* color of the scrollbar thumb */
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #555; /* color of the scrollbar thumb on hover */
}

/* Dark mode scrollbar */
.dark .custom-scrollbar::-webkit-scrollbar-track {
  background: #333;
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: #666;
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #888;
}
</style>