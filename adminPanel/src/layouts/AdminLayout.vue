<template>
  <div class="flex flex-col h-screen bg-gray-100 dark:bg-gray-900">
    <!-- Navbar on top -->
    <header class="flex-shrink-0">
      <Navbar :user="user" />
    </header>

    <!-- Content: Sidebar + Main -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar on the left -->
      <Sidebar :user-role="user ? user.role : ''" />

      <!-- Main content area -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 md:p-8 custom-scrollbar">
        <div v-if="loadingUser" class="flex items-center justify-center h-full">
          <p class="text-gray-500 dark:text-gray-400">Loading user data...</p>
        </div>
        <div v-else-if="userError" class="flex items-center justify-center h-full">
          <p class="text-red-500 dark:text-red-400">Error loading user data. Please try again.</p>
        </div>
        <router-view v-else v-slot="{ Component }">
          <transition name="fade-transform" mode="out-in">
            <component :is="Component" :key="$route.path" />
          </transition>
        </router-view>
      </main>
    </div>

    <!-- Optional Footer -->
    <!-- <footer class="flex-shrink-0 p-4 bg-white dark:bg-gray-800 shadow-inner">
         Footer here
    </footer> -->
  </div>
</template>

<script>
import Sidebar from '../components/Sidebar.vue';
import Navbar from '../components/Navbar.vue';
import auth from '../services/auth';

export default {
  name: 'AdminLayout',
  components: { Sidebar, Navbar },
  data() {
    return {
      user: null,
      loadingUser: true,
      userError: null,
    };
  },
  async created() {
    try {
      this.user = await auth.getUser();
      if (!this.user) this.$router.push({ name: 'Login' });
    } catch (error) {
      this.userError = 'Authentication failed.';
      this.$router.push({ name: 'Login' });
    } finally {
      this.loadingUser = false;
    }
  }
};
</script>

<style scoped>
/* Fade & Transform transitions */
.fade-transform-enter-active,
.fade-transform-leave-active {
  transition: opacity 0.3s ease-in-out, transform 0.3s ease-in-out;
}
.fade-transform-enter-from { opacity: 0; transform: translateY(10px); }
.fade-transform-leave-to { opacity: 0; transform: translateY(-10px); }

/* Custom scrollbar */
.custom-scrollbar::-webkit-scrollbar { width: 8px; }
.custom-scrollbar::-webkit-scrollbar-track { background: #f1f1f1; border-radius: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #888; border-radius: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: #555; }
.dark .custom-scrollbar::-webkit-scrollbar-track { background: #333; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #666; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover { background: #888; }
</style>
