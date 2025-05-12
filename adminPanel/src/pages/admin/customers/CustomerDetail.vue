<template>
    <div>
      <h2 class="text-2xl mb-4">Customer Detail</h2>
      <div v-if="customer">
        <p><strong>Name:</strong> {{ customer.name }}</p>
        <p><strong>Phone:</strong> {{ customer.phone }}</p>
        <p><strong>Address:</strong> {{ customer.address }}</p>
      </div>
      <div v-else>Loading...</div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  export default {
    name: 'CustomerDetail',
    props: ['id'],
    data() { return { customer: null }; },
    async created() {
      const res = await axios.get(`/admin/customers/${this.$route.params.id}`, { withCredentials: true });
      this.customer = res.data;
    }
  };
  </script>