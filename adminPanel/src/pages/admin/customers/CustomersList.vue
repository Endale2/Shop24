<template>
    <div>
      <h2 class="text-2xl mb-4">Customers</h2>
      <ul>
        <li v-for="c in customers" :key="c.id">
          <router-link :to="`/customers/${c.id}`">{{ c.name || c.id }}</router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    name: 'CustomersList',
    data() { return { customers: [] }; },
    async created() {
      const res = await axios.get('/admin/customers', { withCredentials: true });
      this.customers = res.data;
    }
  };
  </script>