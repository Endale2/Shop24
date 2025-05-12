<template>
    <div>
      <h2 class="text-2xl mb-4">Stores</h2>
      <ul>
        <li v-for="s in stores" :key="s.id">
          <router-link :to="`/stores/${s.id}`">{{ s.name || s.id }}</router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    name: 'StoresList',
    data() { return { stores: [] }; },
    async created() {
      const res = await axios.get('/admin/stores', { withCredentials: true });
      this.stores = res.data;
    }
  };
  </script>