
<template>
  <div class="min-h-full bg-gray-50">
    <div class="p-4 sm:p-6 lg:p-8">
      
      <!-- Header Section -->
      <div class="mb-6 sm:mb-8">
        <div class="flex items-center mb-3">
          <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center mr-3 shadow-sm">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-gray-900 tracking-tight">
              Dashboard
            </h1>
            <p class="text-sm text-gray-600 mt-1">
              Welcome back! Here's a snapshot of your shop's performance.
            </p>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-16 text-gray-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-8 w-8 text-green-500 mb-3 border-3 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-sm font-medium">Loading dashboard data...</p>
        </div>
      </div>

      <div v-else class="space-y-6">
        
        <!-- Stat Cards Grid -->
        <div class="stats-grid">
          <!-- Total Products Card -->
          <div class="stat-card">
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <p class="stat-label">Total Products</p>
                <p class="stat-value">{{ stats.products }}</p>
                <div class="stat-meta">
                  <CubeIcon class="w-3 h-3" />
                  {{ stats.lowStockProducts }} low stock
                </div>
              </div>
              <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center group-hover:bg-primary-200 transition-colors duration-200">
                <CubeIcon class="w-5 h-5 text-primary" />
              </div>
            </div>
          </div>

          <!-- Total Orders Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 group">
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <div class="flex items-center mb-1">
                  <p class="text-xs font-medium text-green-600 uppercase tracking-wide">Total Orders</p>
                  <span class="ml-1 text-xs text-gray-400">(paid, shipped, delivered)</span>
                </div>
                <p class="text-2xl font-bold text-green-700 mb-1">{{ stats.orders }}</p>
                <div class="flex items-center text-xs text-green-600 font-medium">
                  <TrendingUpIcon class="w-3 h-3 mr-1" />
                  +{{ stats.newOrdersToday }} today
                </div>
              </div>
              <div class="w-10 h-10 bg-green-50 rounded-lg flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                <ShoppingBagIcon class="w-5 h-5 text-green-600" />
              </div>
            </div>
          </div>

          <!-- Pending Orders Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 group">
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <p class="text-xs font-medium text-gray-500 uppercase tracking-wide mb-1">Pending Orders</p>
                <p class="text-2xl font-bold text-amber-600 mb-1">{{ pendingOrdersCount }}</p>
                <div class="flex items-center text-xs text-amber-600 font-medium">
                  <TrendingUpIcon class="w-3 h-3 mr-1" />
                  Awaiting action
                </div>
              </div>
              <div class="w-10 h-10 bg-amber-50 rounded-lg flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                <ShoppingBagIcon class="w-5 h-5 text-amber-600" />
              </div>
            </div>
          </div>

          <!-- Total Customers Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 group">
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <p class="text-xs font-medium text-gray-500 uppercase tracking-wide mb-1">Total Customers</p>
                <p class="text-2xl font-bold text-gray-900 mb-1">{{ stats.customers }}</p>
                <div class="flex items-center text-xs text-green-600 font-medium">
                  <UserAddIcon class="w-3 h-3 mr-1" />
                  +{{ stats.newCustomersToday }} today
                </div>
              </div>
              <div class="w-10 h-10 bg-green-50 rounded-lg flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                <UsersIcon class="w-5 h-5 text-green-600" />
              </div>
            </div>
          </div>

          <!-- Total Revenue Card -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 group">
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <div class="flex items-center mb-1">
                  <p class="text-xs font-medium text-green-600 uppercase tracking-wide">Total Revenue</p>
                  <span class="ml-1 text-xs text-gray-400">(all paid, shipped, delivered)</span>
                </div>
                <p class="text-2xl font-bold text-green-700 mb-1">${{ formatCurrency(stats.totalRevenue) }}</p>
                <div class="flex items-center text-xs text-green-600 font-medium">
                  <TrendingUpIcon class="w-3 h-3 mr-1" />
                  +${{ formatCurrency(stats.revenueToday) }} today
                </div>
              </div>
              <div class="w-10 h-10 bg-green-50 rounded-lg flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                <CurrencyDollarIcon class="w-5 h-5 text-green-600" />
              </div>
            </div>
          </div>
        </div>

        <!-- Main Content Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          
          <!-- Sales Performance Chart -->
          <div class="lg:col-span-2 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Sales Performance</h2>
              <p class="text-xs text-gray-500 mt-0.5">Last 7 days revenue trend</p>
            </div>
            <div class="p-4">
              <div class="h-48 flex items-end justify-center">
                <div class="w-full">
                  <div class="grid grid-cols-7 gap-2 h-full">
                    <div 
                      v-for="(day, index) in weeklySales" 
                      :key="index"
                      class="text-center flex flex-col items-center justify-end relative group"
                    >
                      <div class="absolute -top-6 left-1/2 transform -translate-x-1/2 bg-gray-900 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity duration-200 whitespace-nowrap z-10">
                        ${{ formatCurrency(day.amount) }}
                      </div>
                      <div 
                        class="bg-gradient-to-t from-green-500 to-green-600 w-full rounded-t-sm transition-all duration-200 hover:from-green-600 hover:to-green-700 cursor-pointer"
                        :style="{ height: `${Math.max(day.amount / maxWeeklySales * 120, 2)}px` }"
                      ></div>
                      <div class="text-xs text-gray-500 font-medium uppercase tracking-wider mt-2">{{ day.day }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Quick Stats Panel -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <h2 class="text-sm font-semibold text-gray-900">Quick Stats</h2>
            </div>
            <div class="p-4 space-y-3">
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-md hover:bg-gray-100 transition-colors duration-200">
                <div>
                  <p class="text-xs text-gray-500 font-medium">Avg Order Value</p>
                  <p class="text-lg font-bold text-gray-900">${{ formatCurrency(stats.averageOrderValue) }}</p>
                </div>
                <div class="w-8 h-8 bg-blue-100 rounded-md flex items-center justify-center">
                  <ShoppingCartIcon class="w-4 h-4 text-blue-600" />
                </div>
              </div>
              
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-md hover:bg-gray-100 transition-colors duration-200">
                <div>
                  <p class="text-xs text-gray-500 font-medium">Conversion Rate</p>
                  <p class="text-lg font-bold text-gray-900">{{ stats.conversionRate }}%</p>
                </div>
                <div class="w-8 h-8 bg-green-100 rounded-md flex items-center justify-center">
                  <ChartBarIcon class="w-4 h-4 text-green-600" />
                </div>
              </div>
              
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-md hover:bg-gray-100 transition-colors duration-200">
                <div class="flex-1 min-w-0">
                  <p class="text-xs text-gray-500 font-medium">Top Seller</p>
                  <p class="text-sm font-semibold text-gray-900 truncate">{{ stats.topProduct?.name || 'N/A' }}</p>
                </div>
                <div class="w-8 h-8 bg-purple-100 rounded-md flex items-center justify-center ml-2">
                  <StarIcon class="w-4 h-4 text-purple-600" />
                </div>
              </div>
              
              <div class="flex items-center justify-between p-3 bg-gray-50 rounded-md hover:bg-red-50 transition-colors duration-200 cursor-pointer" @click="goToDiscounts">
                <div>
                  <p class="text-xs text-gray-500 font-medium">Active Discounts</p>
                  <p class="text-lg font-bold text-gray-900">{{ stats.activeDiscounts }}</p>
                </div>
                <div class="w-8 h-8 bg-red-100 rounded-md flex items-center justify-center">
                  <TagIcon class="w-4 h-4 text-red-600" />
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Activity Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          
          <!-- Recent Orders -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-semibold text-gray-900">Recent Orders</h2>
                <router-link to="/dashboard/orders" class="text-xs text-green-600 hover:text-green-700 font-medium transition-colors duration-200">
                  View all →
                </router-link>
              </div>
            </div>
            <div v-if="recentOrders.length === 0" class="p-6">
              <div class="text-center text-gray-400">
                <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
                  <ShoppingBagIcon class="w-6 h-6 text-gray-300" />
                </div>
                <p class="text-sm font-medium text-gray-600 mb-1">No orders yet</p>
                <p class="text-xs">New orders will appear here once placed.</p>
              </div>
            </div>
            <div v-else class="divide-y divide-gray-100">
              <div v-for="order in recentOrders.slice(0, 5)" :key="order.id" class="flex items-center justify-between p-3 transition-colors duration-200 hover:bg-gray-50 cursor-pointer group" @click="goToOrder(order.id)">
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 bg-green-50 rounded-md flex-shrink-0 flex items-center justify-center group-hover:scale-105 transition-transform duration-200">
                    <ShoppingBagIcon class="w-4 h-4 text-green-600" />
                  </div>
                  <div>
                    <p class="text-sm font-medium text-gray-900">#{{ order.orderNumber }}</p>
                    <p class="text-xs text-gray-500">{{ formatDateTime(order.createdAt) }}</p>
                  </div>
                </div>
                <div class="text-right">
                  <p class="text-sm font-medium text-gray-900">${{ formatCurrency(order.total) }}</p>
                  <span :class="getStatusClass(order.status)" class="inline-flex px-2 py-0.5 rounded-full text-xs font-medium leading-none mt-1">
                    {{ formatStatus(order.status) }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent Products -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
              <div class="flex items-center justify-between">
                <h2 class="text-sm font-semibold text-gray-900">Recent Products</h2>
                <router-link to="/dashboard/products" class="text-xs text-green-600 hover:text-green-700 font-medium transition-colors duration-200">
                  View all →
                </router-link>
              </div>
            </div>
            <div v-if="recentProducts.length === 0" class="p-6">
              <div class="text-center text-gray-400">
                <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
                  <CubeIcon class="w-6 h-6 text-gray-300" />
                </div>
                <p class="text-sm font-medium text-gray-600 mb-1">No products yet</p>
                <p class="text-xs">Add your first product to get started.</p>
              </div>
            </div>
            <div v-else class="divide-y divide-gray-100">
              <div v-for="product in recentProducts" :key="product.id" class="flex items-center space-x-3 p-3 transition-colors duration-200 hover:bg-gray-50 cursor-pointer group" @click="goToProduct(product.id)">
                <img :src="getProductImage(product)" :alt="product.name" class="w-12 h-12 rounded-md object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-300 transition-all duration-200" @error="$event.target.src='/placeholder-product.jpg'" />
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 truncate">{{ product.name }}</p>
                  <p class="text-xs text-gray-500">${{ formatCurrency(product.price) }}</p>
                  <span v-if="getProductStock(product) === 0" class="inline-block mt-1 px-1.5 py-0.5 bg-red-100 text-red-700 text-xs rounded-full font-medium">Out of Stock</span>
                </div>
                <div class="text-right flex-shrink-0">
                  <p class="text-xs font-medium text-gray-600">Stock: {{ getProductStock(product) }}</p>
                  <p class="text-xs text-gray-400">{{ formatDate(product.createdAt) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Alerts Section -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
            <h2 class="text-sm font-semibold text-gray-900">Alerts & Notifications</h2>
          </div>
          <div class="p-4">
            <div v-if="alerts.length === 0" class="text-center py-6 text-gray-400">
              <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-3">
                <CheckCircleIcon class="w-6 h-6 text-green-500" />
              </div>
              <p class="text-sm font-medium text-gray-600 mb-1">All good!</p>
              <p class="text-xs">No new alerts to show right now.</p>
            </div>
            <div v-else class="space-y-3">
              <div v-for="alert in alerts" :key="alert.id" :class="getAlertClass(alert.type)" class="flex items-start p-3 rounded-md border">
                <component :is="getAlertIcon(alert.type)" class="w-4 h-4 mr-2 flex-shrink-0 mt-0.5" />
                <div class="flex-1">
                  <p class="text-sm font-medium">{{ alert.title }}</p>
                  <p class="text-xs opacity-90 mt-0.5">{{ alert.message }}</p>
                </div>
                <button @click="dismissAlert(alert.id)" class="ml-3 p-1 rounded-full hover:bg-black/10 transition-colors duration-200">
                   <XIcon class="w-3 h-3" />
                </button>
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
import { dashboardService } from '@/services/dashboard'
import { format } from 'date-fns'
import {
  CubeIcon,
  ShoppingBagIcon,
  UsersIcon,
  CurrencyDollarIcon,
  TrendingUpIcon,
  UserAddIcon,
  ShoppingCartIcon,
  ChartBarIcon,
  StarIcon,
  TagIcon,
  CheckCircleIcon,
  ExclamationIcon,
  XCircleIcon,
  XIcon
} from '@heroicons/vue/outline'

const router = useRouter()
const shopStore = useShopStore()
const activeShop = computed(() => shopStore.activeShop)
const shopId = computed(() => activeShop.value?.id)

// State
const loading = ref(true)
const stats = ref({
  products: 0,
  orders: 0,
  customers: 0,
  totalRevenue: 0,
  revenueToday: 0,
  newOrdersToday: 0,
  newCustomersToday: 0,
  lowStockProducts: 0,
  averageOrderValue: 0,
  conversionRate: 0,
  topProduct: null,
  activeDiscounts: 0,
  pendingOrders: 0
})
const recentOrders = ref([])
const recentProducts = ref([])
const weeklySales = ref([])
const alerts = ref([])

// Computed
const maxWeeklySales = computed(() => {
  if (weeklySales.value.length === 0) return 1
  return Math.max(...weeklySales.value.map(day => day.amount))
})

const pendingOrdersCount = computed(() => stats.value.pendingOrders || 0)


// --- Navigation helpers ---
function goToOrder(orderId) {
  router.push({ name: 'OrderDetail', params: { orderId } })
}
function goToProduct(productId) {
  router.push({ name: 'ProductDetail', params: { productId } })
}
function goToAnalytics() {
  router.push({ name: 'Analytics' })
}

function goToDiscounts() {
  router.push({ name: 'Discounts' })
}

function getProductImage(product) {
  if (product.mainImage) return product.mainImage
  if (product.main_image) return product.main_image
  if (Array.isArray(product.images) && product.images.length > 0) return product.images[0]
  return '/placeholder-product.jpg'
}

function getProductStock(product) {
  if (product.total_stock !== undefined) return product.total_stock
  if (product.stock !== undefined) return product.stock
  if (product.variants && product.variants.length > 0) {
    return product.variants.reduce((sum, v) => sum + (v.stock || 0), 0)
  }
  return 0
}

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

function formatDateTime(dateString) {
  if (!dateString) return 'N/A'
  try {
    return format(new Date(dateString), 'MMM d, HH:mm')
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
    'pending': 'bg-amber-100 text-amber-700',
    'paid': 'bg-sky-100 text-sky-700',
    'shipped': 'bg-violet-100 text-violet-700',
    'delivered': 'bg-emerald-100 text-emerald-700',
    'cancelled': 'bg-rose-100 text-rose-700'
  }
  return classMap[status] || 'bg-gray-100 text-gray-700'
}


function getAlertClass(type) {
  const classMap = {
    'warning': 'bg-amber-50 text-amber-800 border-amber-200',
    'error': 'bg-rose-50 text-rose-800 border-rose-200',
    'info': 'bg-sky-50 text-sky-800 border-sky-200',
    'success': 'bg-emerald-50 text-emerald-800 border-emerald-200'
  }
  return classMap[type] || classMap.info
}

function getAlertIcon(type) {
  const iconMap = {
    'warning': ExclamationIcon,
    'error': XCircleIcon,
    'info': ExclamationIcon,
    'success': CheckCircleIcon
  }
  return iconMap[type] || ExclamationIcon
}

function dismissAlert(alertId) {
  alerts.value = alerts.value.filter(alert => alert.id !== alertId)
}

function getPaidRevenue() {
  return stats.value.totalRevenue || 0
}

// Generate weekly sales data from recentOrders
function generateWeeklySales(orders) {
  const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
  const today = new Date()
  const weekData = []
  for (let i = 6; i >= 0; i--) {
    const date = new Date(today)
    date.setDate(date.getDate() - i)
    date.setHours(0, 0, 0, 0)
    const dayOrders = orders.filter(order => {
      const orderDate = new Date(order.createdAt)
      orderDate.setHours(0, 0, 0, 0)
      return orderDate.getTime() === date.getTime()
    })
    const daySales = dayOrders.reduce((sum, order) => sum + (order.total || 0), 0)
    weekData.push({
      day: days[date.getDay()],
      amount: daySales
    })
  }
  return weekData
}

function generateAlerts(products, orders, discounts) {
  const alerts = []
  // Low stock alerts
  const lowStockProducts = products.filter(p => (p.stock || 0) < 5)
  if (lowStockProducts.length > 0) {
    alerts.push({
      id: 'low-stock',
      type: 'warning',
      title: 'Low Stock Alert',
      message: `${lowStockProducts.length} product(s) have low stock (less than 5 items)`
    })
  }
  // No orders today
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  const ordersToday = orders.filter(order => {
    const orderDate = new Date(order.createdAt)
    orderDate.setHours(0, 0, 0, 0)
    return orderDate.getTime() === today.getTime()
  })
  if (ordersToday.length === 0) {
    alerts.push({
      id: 'no-orders',
      type: 'info',
      title: 'No Orders Today',
      message: 'Consider running a promotion to boost sales'
    })
  }
  // Expiring discounts
  // Not available in dashboard data, skip or fetch separately if needed
  return alerts
}

async function loadDashboardData() {
  if (!shopId.value) {
    loading.value = false
    return
  }
  try {
    const data = await dashboardService.fetchDashboardData(shopId.value)
    // Map backend data to UI state
    stats.value = {
      products: data.total_products,
      orders: data.total_orders,
      customers: data.total_customers,
      totalRevenue: data.total_revenue,
      revenueToday: data.revenue_today,
      newOrdersToday: data.new_orders_today,
      newCustomersToday: data.new_customers_today,
      lowStockProducts: data.low_stock_products,
      averageOrderValue: data.average_order_value,
      conversionRate: data.total_customers > 0 ? Math.round((data.total_orders / data.total_customers) * 100) : 0,
      topProduct: data.top_product,
      activeDiscounts: data.active_discounts,
      pendingOrders: data.pending_orders
    }
    recentOrders.value = (data.recent_orders || []).map(order => ({
      id: order.id || order._id || '',
      orderNumber: order.order_number || order.orderNumber || '',
      customerId: order.customer_id || order.customerId || '',
      total: order.total,
      status: order.status,
      createdAt: order.created_at || order.createdAt
    }))
    recentProducts.value = (data.recent_products || []).map(product => ({
      id: product.id || product._id || '',
      name: product.name,
      category: product.category,
      mainImage: product.mainImage || product.main_image,
      stock: product.stock,
      price: product.price, // Use backend price
      createdAt: product.created_at || product.createdAt
    }))
    // Use backend-provided weekly_sales for the graph
    weeklySales.value = (data.weekly_sales || []).map(day => ({
      day: day.day,
      amount: day.amount
    }))
    alerts.value = generateAlerts(data.recent_products || [], data.recent_orders || [], [])
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
}

watch(() => shopStore.activeShop, (newShop) => {
  if (newShop?.id) {
    loadDashboardData()
  }
})

onMounted(() => {
  if (shopId.value) {
    loadDashboardData()
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
