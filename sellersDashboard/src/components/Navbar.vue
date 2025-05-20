<template>
  <header class="bg-white shadow p-4 flex justify-between items-center">
    <h1 class="text-xl font-semibold">Dashboard</h1>
    <div class="flex items-center space-x-4">
      <span v-if="user" class="text-gray-700">{{ user.email }}</span>
      <button
        @click="handleLogout"
        class="text-sm text-red-500 hover:underline"
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
    ...mapGetters('auth', ['user'])
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
