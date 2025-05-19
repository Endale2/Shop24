<template>
  <aside style="width: 200px; background: #eee; padding: 1rem;">
    <h3>Shop: {{ activeShop?.name || '...' }}</h3>
    <nav>
      <ul>
        <li><router-link to="/app/products">Products</router-link></li>
        <li><router-link to="/app/orders">Orders</router-link></li>
        <li><router-link to="/app/customers">Customers</router-link></li>
        <li><router-link to="/app/discounts">Discounts</router-link></li>
        <li><router-link to="/app/analytics">Analytics</router-link></li>
        <li><router-link to="/app/settings">Settings</router-link></li>
      </ul>
    </nav>

    <button @click="logout" style="margin-top: 1rem;">Logout</button>
  </aside>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'Sidebar',
  computed: {
    ...mapState('shops', ['activeShop'])
  },
  methods: {
    ...mapActions('auth', ['logout']),
    logout() {
      this.logout();
      this.$router.push({ name: 'Login' });
    }
  }
};
</script>
