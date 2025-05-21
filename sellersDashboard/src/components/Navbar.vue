<template>
  <header class="bg-white shadow p-4 flex justify-between items-center border-b border-gray-200">
    <div class="flex items-center">
      
      <div class="text-sm text-gray-600 ml-6 flex items-center space-x-4">
        <span v-if="user" class="flex items-center">
          <svg class="w-4 h-4 mr-1 text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd"></path></svg>
          User: {{ user.id }}
        </span>
        
      </div>
    </div>
    <div class="flex items-center space-x-6">
      <span v-if="user" class="text-gray-700 text-md font-medium">{{ user.email }}</span>
      <button
        @click="handleLogout"
        class="bg-green-600 text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-300 focus:ring-opacity-75 font-semibold py-2 px-4 rounded-full transition duration-300 ease-in-out transform hover:scale-105"
      >
        Logout
      </button>
    </div>
  </header>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'Navbar',
  computed: {
    ...mapGetters('auth', ['user']),
    ...mapGetters('shops', ['activeShop'])
  },
  methods: {
    ...mapActions('auth', ['logout']),
    async handleLogout() {
      await this.logout();
      this.$router.push({ name: 'Login' });
    }
  }
};
</script>