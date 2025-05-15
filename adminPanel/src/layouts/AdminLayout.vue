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
  name: 'AdminLayoutDesktop',
  components: {
    Sidebar,
    Navbar,
  },
  data() {
    return {
      user: null,
      userError: null,
      loadingUser: true,
      // No responsive state data needed for desktop-only
    };
  },
  // No responsive computed properties or methods
  // No window resize event listeners

  async created() {
    // Fetch user data when the layout is created
    try {
      this.user = await auth.getUser();
      if (!this.user) {
        console.warn('AdminLayoutDesktop: User not authenticated or not found, redirecting...');
        // Redirect to the login route if authentication fails
        this.$router.push('/login');
      }
    } catch (error) {
      console.error('AdminLayoutDesktop: Error fetching user:', error);
      this.userError = 'Authentication failed.'; // Set error message
       this.$router.push('/login'); // Redirect to login on fetch error
    } finally {
       this.loadingUser = false;
    }
  },
   // No beforeUnmount/beforeDestroy needed without listeners
};
</script>

<style scoped>
/*
  Scoped styles here are for layout transitions or specific component overrides.
  Tailwind classes should handle most of the styling and layout.
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

/* Custom Scrollbar (Optional, ensure styles are consistent) */
.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #555;
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