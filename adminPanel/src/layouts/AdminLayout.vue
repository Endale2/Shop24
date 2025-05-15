<template>
  <div class="flex h-screen bg-gray-100 dark:bg-gray-900 overflow-hidden">

    <Sidebar :user-role="user ? user.role : ''" />

    <div class="flex-1 flex flex-col overflow-hidden ml-48">

      <Navbar :user="user" />

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
  name: 'AdminLayoutDesktop', // Explicitly name for clarity
  components: {
    Sidebar,
    Navbar,
  },
  data() {
    return {
      user: null,
      userError: null,
      loadingUser: true,
      // Removed all responsiveness data (isSidebarOpen, windowWidth)
    };
  },
  // Removed all responsiveness computed properties and methods (isLargeScreen, toggleSidebar, handleResize, etc.)
  // Removed all window resize event listeners

  async created() {
    // Fetch user data (needed to show user name in Navbar, or redirect if not authenticated)
    try {
      this.user = await auth.getUser();
      if (!this.user) {
        console.warn('AdminLayoutDesktop: User not authenticated or not found, redirecting...');
        // Ensure router is available and configured with a 'Login' route
        this.$router.push('/login');
      }
    } catch (error) {
      console.error('AdminLayoutDesktop: Error fetching user:', error);
      this.userError = 'Authentication failed.'; // Set error message
       // Redirect to login on fetch error
       this.$router.push('/login');
    } finally {
       this.loadingUser = false;
    }
  },
   // Removed beforeUnmount/beforeDestroy
};
</script>

<style scoped>
/*
  Scoped styles are primarily for transitions or specific component overrides.
  Tailwind handles the layout and appearance.
*/

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

/* Custom Scrollbar (Optional, shared styling) */
/* Make sure these styles match the ones you are using elsewhere if any */
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