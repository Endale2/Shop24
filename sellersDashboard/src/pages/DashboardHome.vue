<template>
  <div>
    <h1>Welcome to your shop dashboard, {{ activeShop?.name || 'Shop' }}!</h1>
    <p>Use the sidebar to navigate through your products, orders, customers, and more.</p>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'DashboardHome',
  computed: {
    ...mapState('shops', ['activeShop'])
  }
};
</script>
