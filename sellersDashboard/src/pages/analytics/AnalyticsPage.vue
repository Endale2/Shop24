<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
          <div class="flex items-center mb-3 md:mb-0">
            <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
            <div>
              <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
                Analytics Dashboard
              </h1>
              <p class="text-sm text-gray-600 mt-1">
                Deep insights into your shop's performance
              </p>
            </div>
          </div>
          <div class="inline-flex bg-white p-1 rounded-lg shadow-sm border border-gray-200">
            <button
              v-for="period in periods"
              :key="period.id"
              @click="changePeriod(period.id)"
              :class="[
                'px-3 py-2 text-xs font-medium rounded-md transition-all duration-200',
                selectedPeriod === period.id
                  ? 'bg-green-600 text-white shadow-sm'
                  : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
              ]"
            >
              {{ period.label }}
            </button>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading analytics data...</p>
        </div>
      </div>

      <!-- Analytics Content -->
      <div v-else class="space-y-6">
        <!-- Stats Cards -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
          <div
            v-for="stat in analyticsStats"
            :key="stat.title"
            class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all duration-200 border-l-4 group"
            :class="stat.borderColor"
          >
            <div class="flex items-center space-x-3">
              <div class="p-2 rounded-lg" :class="stat.iconBg">
                <component :is="stat.icon" class="h-5 w-5" :class="stat.iconColor" />
              </div>
              <div class="flex-1">
                <p class="text-xs font-medium text-gray-500 uppercase tracking-wide mb-1">{{ stat.title }}</p>
                <p class="text-2xl font-bold text-gray-900 mb-1">{{ stat.value }}</p>
                <p
                  class="text-xs font-medium"
                  :class="stat.change >= 0 ? 'text-green-600' : 'text-red-600'"
                >
                  <component :is="stat.change >= 0 ? TrendingUpIcon : TrendingDownIcon" class="w-3 h-3 inline mr-1" />
                  {{ stat.change >= 0 ? '+' : '' }}{{ stat.change }}% vs last period
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Charts Section -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- Revenue Over Time -->
          <div class="lg:col-span-2 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-semibold text-gray-900">Revenue Over Time</h2>
                <div class="text-xs text-gray-500">
                  {{ getPeriodLabel(selectedPeriod) }}
                </div>
              </div>
            </div>
            <div class="p-4">
              <div class="h-64">
                <Line
                  v-if="revenueChartData && Array.isArray(revenueChartData.labels) && revenueChartData.labels.length"
                  :data="revenueChartData"
                  :options="lineChartOptions"
                />
                <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
                  <p class="text-sm text-gray-500">No revenue data available</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Sales by Category -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Sales by Category</h2>
            </div>
            <div class="p-4">
              <div class="h-64">
                <Doughnut
                  v-if="categoryChartData"
                  :data="categoryChartData"
                  :options="doughnutChartOptions"
                />
                <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
                  <p class="text-sm text-gray-500">No category data available</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Additional Charts -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- Orders Over Time -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Orders Over Time</h2>
            </div>
            <div class="p-4">
              <div class="h-48">
                <Line
                  v-if="ordersChartData && Array.isArray(ordersChartData.labels) && ordersChartData.labels.length"
                  :data="ordersChartData"
                  :options="lineChartOptions"
                />
                <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
                  <p class="text-sm text-gray-500">No order data available</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Customer Growth -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Customer Growth</h2>
            </div>
            <div class="p-4">
              <div class="h-48">
                <Line
                  v-if="customersChartData"
                  :data="customersChartData"
                  :options="lineChartOptions"
                />
                <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
                  <p class="text-sm text-gray-500">No customer data available</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Top Products & Recent Orders -->
        <div class="grid grid-cols-1 lg:grid-cols-5 gap-6">
          <!-- Top Selling Products -->
          <div class="lg:col-span-2 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-semibold text-gray-900 flex items-center">Top Selling Products
                  <span class="ml-2 text-xs text-gray-400" title="Only paid, shipped, or delivered orders are counted as sold.">(sold = paid/shipped/delivered)</span>
                </h2>
                <router-link 
                  to="/dashboard/products" 
                  class="text-xs text-green-600 hover:text-green-700 font-medium transition-colors duration-200"
                >
                  View all →
                </router-link>
              </div>
            </div>
            <div class="p-4">
              <div v-if="Array.isArray(topProductsList) && topProductsList.length === 0" class="text-center py-6 text-gray-400">
                <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
                  <CubeIcon class="w-6 h-6 text-gray-300" />
                </div>
                <p class="text-sm font-medium text-gray-600 mb-1">No products yet</p>
                <p class="text-xs">Add products to see analytics</p>
              </div>
              <div v-else class="space-y-3">
                <div
                  v-for="(product, index) in (Array.isArray(topProductsList) ? topProductsList : [])"
                  :key="product.id"
                  class="flex items-center space-x-3 p-3 bg-gray-50 rounded-md hover:bg-gray-100 transition-colors duration-200 cursor-pointer group"
                  @click="goToProduct(product.id)"
                >
                  <div class="flex-shrink-0">
                    <div class="w-10 h-10 bg-gray-200 rounded-md overflow-hidden">
                      <img 
                        v-if="product.mainImage"
                        :src="product.mainImage"
                        :alt="product.name"
                        class="w-full h-full object-cover"
                      />
                      <div v-else class="w-full h-full bg-gray-200 flex items-center justify-center text-gray-400">
                        <span class="text-xs">No Image</span>
                      </div>
                    </div>
                  </div>
                  <div class="flex-1 min-w-0">
                    <p class="text-sm font-medium text-gray-900 truncate">{{ product.name }}</p>
                    <p class="text-xs text-gray-500">{{ product.category }}</p>
                  </div>
                  <div class="text-right">
                    <p class="text-sm font-semibold text-gray-900">${{ formatCurrency(product.totalRevenue) }}</p>
                    <p class="text-xs text-gray-500">{{ product.totalSold }} sold</p>
                  </div>
                  <div class="flex-shrink-0">
                    <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                      <span class="text-xs font-bold text-green-600">#{{ index + 1 }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent Orders Table -->
          <div class="lg:col-span-3 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-semibold text-gray-900">Recent Orders</h2>
                <router-link 
                  to="/orders" 
                  class="text-xs text-green-600 hover:text-green-700 font-medium transition-colors duration-200"
                >
                  View all →
                </router-link>
              </div>
            </div>
            <div class="p-4">
              <div class="overflow-x-auto">
                <table class="w-full text-sm">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="p-3 text-left font-medium text-gray-600 text-xs">Order ID</th>
                      <th class="p-3 text-left font-medium text-gray-600 text-xs">Customer</th>
                      <th class="p-3 text-left font-medium text-gray-600 text-xs">Amount</th>
                      <th class="p-3 text-left font-medium text-gray-600 text-xs">Status</th>
                      <th class="p-3 text-left font-medium text-gray-600 text-xs">Date</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100">
                    <tr
                      v-for="order in (Array.isArray(recentOrders) ? recentOrders : [])"
                      :key="order.id"
                      class="hover:bg-gray-50 cursor-pointer transition-colors duration-200"
                      @click="goToOrder(order.id)"
                    >
                      <td class="p-3 font-mono text-gray-500 text-xs">#{{ order.order_number }}</td>
                      <td class="p-3 text-gray-800 text-xs">{{ order.customer_id }}</td>
                      <td class="p-3 font-semibold text-gray-900 text-xs">${{ formatCurrency(order.total) }}</td>
                      <td class="p-3">
                        <span
                          class="px-2 py-0.5 text-xs font-medium rounded-full"
                          :class="getStatusClass(order.status)"
                        >
                          {{ formatStatus(order.status) }}
                        </span>
                      </td>
                      <td class="p-3 text-gray-600 text-xs">{{ formatDate(order.created_at) }}</td>
                    </tr>
                  </tbody>
                </table>
                <div v-if="Array.isArray(recentOrders) && recentOrders.length === 0" class="text-center py-6 text-gray-400">
                  <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
                    <ShoppingBagIcon class="w-6 h-6 text-gray-300" />
                  </div>
                  <p class="text-sm font-medium text-gray-600 mb-1">No orders yet</p>
                  <p class="text-xs">Orders will appear here once customers start shopping</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Performance Metrics -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- Customer Insights -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Customer Insights</h2>
            </div>
            <div class="p-4">
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">New Customers</span>
                  <span class="font-semibold text-green-600 text-sm">+{{ customerInsights?.newCustomers ?? 0 }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Returning Customers</span>
                  <span class="font-semibold text-blue-600 text-sm">{{ customerInsights?.returningCustomers ?? 0 }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Avg Order Value</span>
                  <span class="font-semibold text-gray-900 text-sm">${{ formatCurrency(customerInsights?.avgOrderValue ?? 0) }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Customer Lifetime Value</span>
                  <span class="font-semibold text-purple-600 text-sm">${{ formatCurrency(customerInsights?.lifetimeValue ?? 0) }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Inventory Status -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Inventory Status</h2>
            </div>
            <div class="p-4">
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Total Products</span>
                  <span class="font-semibold text-gray-900 text-sm">{{ inventory.total_products }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">In Stock</span>
                  <span class="font-semibold text-green-600 text-sm">{{ inventory.in_stock }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Low Stock</span>
                  <span class="font-semibold text-yellow-600 text-sm">{{ inventory.low_stock }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-600">Out of Stock</span>
                  <span class="font-semibold text-red-600 text-sm">{{ inventory.out_of_stock }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'
import { analyticsService } from '@/services/analytics'
import { format } from 'date-fns'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  ArcElement,
} from 'chart.js'
import { Line, Doughnut } from 'vue-chartjs'
import {
  CurrencyDollarIcon,
  ShoppingBagIcon,
  UsersIcon,
  ChartBarIcon,
  TrendingUpIcon,
  TrendingDownIcon,
  CubeIcon
} from '@heroicons/vue/outline'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, ArcElement)

const router = useRouter()
const shopStore = useShopStore()
const activeShop = computed(() => shopStore.activeShop)
const shopId = computed(() => activeShop.value?.id)

// State
const loading = ref(true)
const selectedPeriod = ref('last_30_days')
const summary = ref(null)
const revenueOverTime = ref({ labels: [], values: [] })
const ordersOverTime = ref({ labels: [], values: [] })
const topProducts = ref([])
const inventoryStatus = ref(null)

// Add new state for categorySales, customersOverTime, and recentOrders
const categorySales = ref({ labels: [], values: [] })
const customersOverTime = ref({ labels: [], values: [] })
const recentOrders = ref([])

const periods = [
  { id: 'last_7_days', label: '7 Days' },
  { id: 'last_30_days', label: '30 Days' },
  { id: 'last_90_days', label: '90 Days' },
  { id: 'last_12_months', label: '12 Months' },
]

const lineChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { 
    legend: { display: false },
    tooltip: {
      callbacks: {
        label: function(context) {
          return `$${context.parsed.y.toLocaleString()}`
        }
      }
    }
  },
  scales: { 
    y: { 
      beginAtZero: true,
      ticks: { 
        callback: function(value) {
          return '$' + value.toLocaleString()
        }
      } 
    } 
  }
}

const doughnutChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { 
    legend: { 
      position: 'bottom',
      labels: {
        padding: 20,
        usePointStyle: true
      }
    } 
  }
}

const analyticsStats = computed(() => {
  if (!summary.value) return []
  // For change %, you may want to fetch previous period summary as well (not implemented here)
  return [
    {
      title: 'Total Revenue',
      value: `$${formatCurrency(summary.value.total_revenue)}`,
      change: 0, // Placeholder
      icon: CurrencyDollarIcon,
      iconBg: 'bg-green-100',
      iconColor: 'text-green-600',
      borderColor: 'border-green-500'
    },
    {
      title: 'Total Orders',
      value: summary.value.total_orders?.toLocaleString() ?? '0',
      change: 0,
      icon: ShoppingBagIcon,
      iconBg: 'bg-blue-100',
      iconColor: 'text-blue-600',
      borderColor: 'border-blue-500'
    },
    {
      title: 'New Customers',
      value: summary.value.new_customers?.toLocaleString() ?? '0',
      change: 0,
      icon: UsersIcon,
      iconBg: 'bg-indigo-100',
      iconColor: 'text-indigo-600',
      borderColor: 'border-indigo-500'
    },
    {
      title: 'Avg Order Value',
      value: `$${formatCurrency(summary.value.avg_order_value)}`,
      change: 0,
      icon: ChartBarIcon,
      iconBg: 'bg-amber-100',
      iconColor: 'text-amber-600',
      borderColor: 'border-amber-500'
    }
  ]
})

const revenueChartData = computed(() => {
  if (!Array.isArray(revenueOverTime.value?.labels) || revenueOverTime.value.labels.length === 0) return null
  return {
    labels: revenueOverTime.value.labels.map(date => format(new Date(date), 'MMM d')),
    datasets: [{
      label: 'Revenue',
      data: Array.isArray(revenueOverTime.value.values) ? revenueOverTime.value.values : [],
      borderColor: 'rgb(16, 185, 129)',
      backgroundColor: 'rgba(16, 185, 129, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

const ordersChartData = computed(() => {
  if (!Array.isArray(ordersOverTime.value?.labels) || ordersOverTime.value.labels.length === 0) return null
  return {
    labels: ordersOverTime.value.labels.map(date => format(new Date(date), 'MMM d')),
    datasets: [{
      label: 'Orders',
      data: Array.isArray(ordersOverTime.value.values) ? ordersOverTime.value.values : [],
      borderColor: 'rgb(59, 130, 246)',
      backgroundColor: 'rgba(59, 130, 246, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

const topProductsList = computed(() => Array.isArray(topProducts.value) ? topProducts.value : [])

const inventory = computed(() => inventoryStatus.value && typeof inventoryStatus.value === 'object' ? inventoryStatus.value : {
  total_products: 0,
  in_stock: 0,
  low_stock: 0,
  out_of_stock: 0
})

// Defensive computed for customerInsights
const customerInsights = computed(() => {
  if (!summary.value) {
    return {
      newCustomers: 0,
      returningCustomers: 0,
      avgOrderValue: 0,
      lifetimeValue: 0
    }
  }
  // If you have real backend data for these, use it. Otherwise, fallback to summary values or 0.
  return {
    newCustomers: summary.value.new_customers ?? 0,
    returningCustomers: (summary.value.total_customers ?? 0) - (summary.value.new_customers ?? 0),
    avgOrderValue: summary.value.avg_order_value ?? 0,
    lifetimeValue: (summary.value.avg_order_value ?? 0) * 2.5 // or use backend value if available
  }
})

// Update chart computed for category sales
const categoryChartData = computed(() => {
  if (!Array.isArray(categorySales.value?.labels) || categorySales.value.labels.length === 0) return null
  return {
    labels: categorySales.value.labels,
    datasets: [{
      data: Array.isArray(categorySales.value.values) ? categorySales.value.values : [],
      backgroundColor: ['#10B981', '#3B82F6', '#8B5CF6', '#F59E0B', '#EF4444', '#06B6D4']
    }]
  }
})

// Update chart computed for customer growth
const customersChartData = computed(() => {
  if (!Array.isArray(customersOverTime.value?.labels) || customersOverTime.value.labels.length === 0) return null
  return {
    labels: customersOverTime.value.labels.map(date => format(new Date(date), 'MMM d')),
    datasets: [{
      label: 'Customers',
      data: Array.isArray(customersOverTime.value.values) ? customersOverTime.value.values : [],
      borderColor: 'rgb(139, 92, 246)',
      backgroundColor: 'rgba(139, 92, 246, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

function formatCurrency(amount) {
  return (amount || 0).toFixed(2)
}

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  try {
    return format(new Date(dateString), 'MMM d, yyyy')
  } catch (e) {
    return 'Invalid Date'
  }
}

function formatStatus(status) {
  const statusMap = {
    'pending': 'Pending',
    'paid': 'Paid',
    'shipped': 'Shipped',
    'delivered': 'Delivered',
    'cancelled': 'Cancelled'
  }
  return statusMap[status] || status
}

function getStatusClass(status) {
  const classMap = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800'
  }
  return classMap[status] || 'bg-gray-100 text-gray-800'
}

function getPeriodLabel(period) {
  const labels = {
    'last_7_days': 'Last 7 Days',
    'last_30_days': 'Last 30 Days',
    'last_90_days': 'Last 90 Days',
    'last_12_months': 'Last 12 Months'
  }
  return labels[period] || period
}

function getDaysForPeriod(period) {
  const daysMap = {
    'last_7_days': 7,
    'last_30_days': 30,
    'last_90_days': 90,
    'last_12_months': 365
  }
  return daysMap[period] || 30
}

function calculatePercentageChange(current, previous) {
  if (previous === 0) return current > 0 ? 100 : 0
  return Math.round(((current - previous) / previous) * 100)
}

function getChartData(type) {
  const days = getDaysForPeriod(selectedPeriod.value)
  const labels = []
  const values = []
  
  for (let i = days - 1; i >= 0; i--) {
    const date = new Date()
    date.setDate(date.getDate() - i)
    date.setHours(0, 0, 0, 0)
    
    const nextDate = new Date(date)
    nextDate.setDate(nextDate.getDate() + 1)
    
    let value = 0
    
    if (type === 'revenue') {
      value = orders.value
        .filter(order => {
          const orderDate = new Date(order.createdAt)
          return orderDate >= date && orderDate < nextDate
        })
        .reduce((sum, order) => sum + (order.total || 0), 0)
    } else if (type === 'orders') {
      value = orders.value.filter(order => {
        const orderDate = new Date(order.createdAt)
        return orderDate >= date && orderDate < nextDate
      }).length
    } else if (type === 'customers') {
      value = customers.value.filter(customer => {
        const customerDate = new Date(customer.createdAt)
        return customerDate >= date && customerDate < nextDate
      }).length
    }
    
    labels.push(format(date, 'MMM d'))
    values.push(value)
  }
  
  return { labels, values }
}

function getCategoryData() {
  const categoryMap = {}
  
  products.value.forEach(product => {
    const category = product.category || 'Uncategorized'
    if (!categoryMap[category]) {
      categoryMap[category] = 0
    }
    categoryMap[category] += calculateProductRevenue(product.id)
  })
  
  const sortedCategories = Object.entries(categoryMap)
    .sort(([,a], [,b]) => b - a)
    .slice(0, 6)
  
  return {
    labels: sortedCategories.map(([category]) => category),
    values: sortedCategories.map(([, revenue]) => revenue)
  }
}

function calculateProductRevenue(productId) {
  return orders.value
    .flatMap(order => order.items || [])
    .filter(item => item.productId === productId)
    .reduce((sum, item) => sum + (item.totalPrice || 0), 0)
}

function calculateProductSold(productId) {
  return orders.value
    .flatMap(order => order.items || [])
    .filter(item => item.productId === productId)
    .reduce((sum, item) => sum + (item.quantity || 0), 0)
}

function getCustomerName(customerId) {
  const customer = customers.value.find(c => c.id === customerId)
  return customer ? `${customer.firstName} ${customer.lastName}` : 'Unknown Customer'
}

function changePeriod(period) {
  selectedPeriod.value = period
  loadAnalyticsData()
}

function goToOrder(orderId) {
  router.push({ name: 'OrderDetail', params: { orderId } })
}

function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}

// Load analytics data
async function loadAnalyticsData() {
  if (!shopId.value) {
    loading.value = false
    return
  }
  loading.value = true
  try {
    const [summaryData, revenueData, ordersData, topProductsData, inventoryData, categorySalesData, customersOverTimeData, recentOrdersData] = await Promise.all([
      analyticsService.fetchSummary(shopId.value),
      analyticsService.fetchRevenueOverTime(shopId.value, getDaysForPeriod(selectedPeriod.value)),
      analyticsService.fetchOrdersOverTime(shopId.value, getDaysForPeriod(selectedPeriod.value)),
      analyticsService.fetchTopProducts(shopId.value),
      analyticsService.fetchInventoryStatus(shopId.value),
      analyticsService.fetchCategorySales(shopId.value, getDaysForPeriod(selectedPeriod.value)),
      analyticsService.fetchCustomersOverTime(shopId.value, getDaysForPeriod(selectedPeriod.value)),
      analyticsService.fetchRecentOrders(shopId.value, 10)
    ])
    summary.value = summaryData
    revenueOverTime.value = revenueData
    ordersOverTime.value = ordersData
    topProducts.value = topProductsData
    inventoryStatus.value = inventoryData
    categorySales.value = categorySalesData
    customersOverTime.value = customersOverTimeData
    recentOrders.value = recentOrdersData
  } catch (e) {
    console.error('Failed to load analytics data:', e)
  } finally {
    loading.value = false
  }
}

// Watch for shop changes
watch(() => shopStore.activeShop, (newShop) => {
  if (newShop?.id) {
    loadAnalyticsData()
  }
})

onMounted(async () => {
  try {
    if (!shopId.value) {
      const ensured = await shopStore.ensureActiveShop()
      if (!ensured?.id) {
        router.replace({ name: 'ShopSelection' })
        return
      }
    }
    await loadAnalyticsData()
  } catch (e) {
    console.error('[AnalyticsPage] Initialization failed:', e)
  }
})
</script>

<style scoped>
/* Custom styles for enhanced design */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

/* Smooth transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Custom scrollbar for better UX */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>
