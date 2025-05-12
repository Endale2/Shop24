<template>
    <nav class="bg-white shadow p-4 flex justify-between items-center">
      <div>
        <span class="font-semibold">Hello, {{ user.name }}</span>
        <span class="text-sm text-gray-500 ml-2">({{ user.role }})</span>
      </div>
      <div class="flex items-center">
        <button @click="logout" class="px-3 py-1 bg-red-500 text-white rounded">Logout</button>
      </div>
    </nav>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    props: ['user'],
    methods: {
      async logout() {
        await axios.post('/auth/admin/logout');
        this.$router.replace({ name: 'Login' });
      }
    }
  };
  </script>