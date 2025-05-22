<template>
  <div class="p-6 space-y-6">
    <h1 class="text-3xl font-bold text-gray-800">Dashboard</h1>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow p-4 flex items-center space-x-4">
        <div class="p-3 bg-green-100 rounded-full">
          <svg class="w-6 h-6 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h18v18H3V3z" />
          </svg>
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Products</p>
          <p class="text-2xl font-semibold text-gray-800">{{ stats.products }}</p>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4 flex items-center space-x-4">
        <div class="p-3 bg-blue-100 rounded-full">
          <svg class="w-6 h-6 text-blue-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M3 14h18" />
          </svg>
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Orders</p>
          <p class="text-2xl font-semibold text-gray-800">{{ stats.orders }}</p>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4 flex items-center space-x-4">
        <div class="p-3 bg-yellow-100 rounded-full">
          <svg class="w-6 h-6 text-yellow-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Customers</p>
          <p class="text-2xl font-semibold text-gray-800">{{ stats.customers }}</p>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4 flex items-center space-x-4">
        <div class="p-3 bg-purple-100 rounded-full">
          <svg class="w-6 h-6 text-purple-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c1.657 0 3-1.343 3-3S13.657 2 12 2 9 3.343 9 5s1.343 3 3 3zM4 21v-2a4 4 0 014-4h8a4 4 0 014 4v2" />
          </svg>
        </div>
        <div>
          <p class="text-sm text-gray-500">Total Sales</p>
          <p class="text-2xl font-semibold text-gray-800">$ {{ stats.sales.toFixed(2) }}</p>
        </div>
      </div>
    </div>

    <!-- Sales Chart Placeholder -->
    <div class="bg-white rounded-lg shadow p-6">
      <h2 class="text-xl font-semibold mb-4 text-gray-700">Sales Over Time</h2>
      <!-- Replace with real chart component later -->
      <div class="h-64 flex items-center justify-center text-gray-400">
        Chart Component Here
      </div>
    </div>

    <!-- Recent Orders -->
    <div class="bg-white rounded-lg shadow p-6">
      <h2 class="text-xl font-semibold mb-4 text-gray-700">Recent Orders</h2>
      <ul class="space-y-2">
        <li v-for="order in recentOrders" :key="order.id" class="flex justify-between items-center p-2 hover:bg-gray-50 rounded">
          <span>Order #{{ order.id }}</span>
          <span class="text-sm text-gray-500">{{ formatDate(order.date) }}</span>
          <span class="font-medium">$ {{ order.total.toFixed(2) }}</span>
        </li>
      </ul>
      <p v-if="!recentOrders.length" class="text-gray-500 text-center">No recent orders.</p>
    </div>
  </div>
</template>

<script>
import { onMounted, reactive } from 'vue';
// import { dashboardService } from '@/services/dashboard'; // implement later

export default {
  name: 'DashboardPage',
  setup() {
    const stats = reactive({ products: 0, orders: 0, customers: 0, sales: 0 });
    const recentOrders = reactive([]);

    onMounted(async () => {
      // TODO: fetch real data
      // const data = await dashboardService.fetchStats();
      // Object.assign(stats, data.stats);
      // recentOrders.push(...data.recentOrders);

      // Mock data for now
      stats.products = 42;
      stats.orders = 17;
      stats.customers = 23;
      stats.sales = 1250.75;
      recentOrders.push(
        { id: '1001', date: '2025-05-20T10:24:00Z', total: 120.00 },
        { id: '1002', date: '2025-05-21T14:12:00Z', total: 85.50 }
      );
    });

    const formatDate = dt => new Date(dt).toLocaleString();
    return { stats, recentOrders, formatDate };
  }
};
</script>

<style scoped>
/* Add any scoped styles here */
</style>
