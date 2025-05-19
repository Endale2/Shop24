<template>
  <div style="display: flex; min-height: 100vh;">
    <!-- Sidebar -->
    <aside style="width: 200px; background: #eee; padding: 1rem;">
      <h3>Shop: {{ activeShop?.name }}</h3>
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
    </aside>

    <!-- Main Content -->
    <main style="flex: 1; padding: 2rem;">
      <router-view />
    </main>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'DashboardLayout',
  computed: {
    ...mapState('shops', ['activeShop'])
  }
};
</script>
