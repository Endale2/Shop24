<template>
  <div>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Dashboard</h1>
    </div>
    <div class="mb-6">
      <h2 class="text-xl font-semibold text-gray-800 dark:text-white">Dashboard Overview</h2>
      <p class="mt-1 text-sm text-gray-600 dark:text-gray-300">Get a quick look at your key metrics.</p>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      <!-- Total Products -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-blue-500 dark:bg-blue-600 text-white p-3 rounded-md mr-4">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10m0-10l-8 4m8-4V4"></path></svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Products</p>
            <div v-if="loadingProducts" class="mt-1 text-lg font-semibold text-gray-900 dark:text-white">Loading...</div>
            <div v-else-if="errorProducts" class="mt-1 text-lg font-semibold text-red-500">Error</div>
            <div v-else class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">{{ productCount }}</div>
          </div>
        </div>
        <p v-if="errorProducts" class="mt-2 text-xs text-red-500">{{ errorProducts }}</p>
      </div>
      <!-- Total Orders -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-yellow-500 dark:bg-yellow-600 text-white p-3 rounded-md mr-4">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M12 16h.01"></path></svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Orders</p>
            <div v-if="loadingOrders" class="mt-1 text-lg font-semibold text-gray-900 dark:text-white">Loading...</div>
            <div v-else-if="errorOrders" class="mt-1 text-lg font-semibold text-red-500">Error</div>
            <div v-else class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">{{ orderCount }}</div>
          </div>
        </div>
        <p v-if="errorOrders" class="mt-2 text-xs text-red-500">{{ errorOrders }}</p>
      </div>
      <!-- Total Revenue -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-purple-500 dark:bg-purple-600 text-white p-3 rounded-md mr-4">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2-1.343-2-3-2zM21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Total Revenue</p>
            <div v-if="loadingRevenue" class="mt-1 text-lg font-semibold text-gray-900 dark:text-white">Loading...</div>
            <div v-else-if="errorRevenue" class="mt-1 text-lg font-semibold text-red-500">Error</div>
            <div v-else class="mt-1 text-2xl font-semibold text-gray-900 dark:text-white">${{ totalRevenue.toLocaleString() }}</div>
          </div>
        </div>
        <p v-if="errorRevenue" class="mt-2 text-xs text-red-500">{{ errorRevenue }}</p>
      </div>
      <!-- Top Products -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0 bg-green-500 dark:bg-green-600 text-white p-3 rounded-md mr-4">
            <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.292M12 14c0 1.657-1.594 3-3.5 3S5 15.657 5 14s1.594-3 3.5-3 3.5 1.343 3.5 3zm0 0h6c0 1.657 1.594 3 3.5 3s3.5-1.343 3.5-3s-1.594-3-3.5-3-3.5 1.343-3.5 3z"></path></svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Top Products</p>
            <div v-if="loadingTopProducts" class="mt-1 text-lg font-semibold text-gray-900 dark:text-white">Loading...</div>
            <div v-else-if="errorTopProducts" class="mt-1 text-lg font-semibold text-red-500">Error</div>
            <ul v-else>
              <li v-for="p in topProducts" :key="p.Key" class="text-gray-900 dark:text-white">{{ p.Key }}: {{ p.Value }} sales</li>
            </ul>
          </div>
        </div>
        <p v-if="errorTopProducts" class="mt-2 text-xs text-red-500">{{ errorTopProducts }}</p>
      </div>
    </div>
    <!-- Top Customers Section -->
    <div class="mt-8 bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 border-b border-gray-200 dark:border-gray-700 pb-2">Top Customers</h3>
      <div v-if="loadingTopCustomers" class="text-gray-900 dark:text-white">Loading...</div>
      <div v-else-if="errorTopCustomers" class="text-red-500">Error</div>
      <ul v-else>
        <li v-for="c in topCustomers" :key="c.Key" class="text-gray-900 dark:text-white">{{ c.Key }}: {{ c.Value }} orders</li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DashboardPage',
  data() {
    return {
      productCount: 0,
      orderCount: 0,
      totalRevenue: 0,
      topProducts: [],
      topCustomers: [],
      loadingProducts: true,
      loadingOrders: true,
      loadingRevenue: true,
      loadingTopProducts: true,
      loadingTopCustomers: true,
      errorProducts: null,
      errorOrders: null,
      errorRevenue: null,
      errorTopProducts: null,
      errorTopCustomers: null,
    };
  },
  async created() {
    // Product count
    try {
      const res = await axios.get('/admin/products/count', { withCredentials: true });
      this.productCount = res.data.count;
    } catch (err) {
      this.errorProducts = 'Failed to load count';
    } finally {
      this.loadingProducts = false;
    }
    // Order count
    try {
      const res = await axios.get('/admin/analytics/order-count', { withCredentials: true });
      this.orderCount = res.data.order_count;
    } catch (err) {
      this.errorOrders = 'Failed to load orders';
    } finally {
      this.loadingOrders = false;
    }
    // Revenue
    try {
      const res = await axios.get('/admin/analytics/total-revenue', { withCredentials: true });
      this.totalRevenue = res.data.total_revenue;
    } catch (err) {
      this.errorRevenue = 'Failed to load revenue';
    } finally {
      this.loadingRevenue = false;
    }
    // Top products
    try {
      const res = await axios.get('/admin/analytics/top-products', { withCredentials: true });
      this.topProducts = res.data;
    } catch (err) {
      this.errorTopProducts = 'Failed to load top products';
    } finally {
      this.loadingTopProducts = false;
    }
    // Top customers
    try {
      const res = await axios.get('/admin/analytics/top-customers', { withCredentials: true });
      this.topCustomers = res.data;
    } catch (err) {
      this.errorTopCustomers = 'Failed to load top customers';
    } finally {
      this.loadingTopCustomers = false;
    }
  },
};
</script>

