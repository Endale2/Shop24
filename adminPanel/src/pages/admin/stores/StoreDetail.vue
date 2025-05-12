<template>
    <div>
      <h2 class="text-2xl mb-4">Store Detail</h2>
      <div v-if="store">
        <p><strong>Name:</strong> {{ store.name }}</p>
        <p><strong>Address:</strong> {{ store.address }}</p>
      </div>
      <div v-else>Loading...</div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    name: 'StoreDetail',
    props: ['id'],
    data() { return { store: null }; },
    async created() {
      const res = await axios.get(`/admin/stores/${this.$route.params.id}`, { withCredentials: true });
      this.store = res.data;
    }
  };
  </script>