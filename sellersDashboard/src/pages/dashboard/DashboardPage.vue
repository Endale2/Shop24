
<template>
  <div class="bg-slate-50 font-sans min-h-screen">
    <div class="p-4 sm:p-6 lg:p-8 max-w-7xl mx-auto">
      
      <div class="mb-10">
        <h1 class="text-3xl sm:text-4xl font-black text-slate-800 tracking-tight">
          Dashboard
        </h1>
        <p class="text-lg text-slate-500 mt-1">
          Welcome back! Here's a snapshot of your shop's performance. 🚀
        </p>
      </div>

      <div v-if="loading" class="text-center py-24 text-slate-500">
        <div class="flex flex-col items-center justify-center">
          <div class="animate-spin h-12 w-12 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
          <p class="text-lg font-medium">Summoning the latest data...</p>
        </div>
      </div>

      <div v-else class="space-y-8">
        
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-6">
          <div class="bg-white rounded-2xl shadow-sm p-6 transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-slate-500">Total Products</p>
                <p class="text-3xl font-bold text-slate-800">{{ stats.products }}</p>
                <p class="text-sm text-sky-600 mt-2 flex items-center font-semibold">
                  <CubeIcon class="w-4 h-4 inline mr-1.5" />
                  {{ stats.lowStockProducts }} low stock
                </p>
              </div>
              <div class="w-12 h-12 bg-sky-50 rounded-xl flex items-center justify-center">
                <CubeIcon class="w-6 h-6 text-sky-500" />
              </div>
            </div>
          </div>

          <!-- Total Orders (with clarification) -->
          <div class="bg-white rounded-2xl shadow-sm p-6 transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-emerald-600 flex items-center">
                  Total Orders
                  <span class="ml-2 text-xs text-slate-400 font-normal">(paid, shipped, delivered)</span>
                </p>
                <p class="text-3xl font-black text-emerald-700">{{ stats.orders }}</p>
                <p class="text-sm text-emerald-600 mt-2 flex items-center font-semibold">
                  <TrendingUpIcon class="w-4 h-4 inline mr-1.5" />
                  +{{ stats.newOrdersToday }} today
                </p>
              </div>
              <div class="w-12 h-12 bg-emerald-50 rounded-xl flex items-center justify-center">
                <ShoppingBagIcon class="w-6 h-6 text-emerald-500" />
              </div>
            </div>
          </div>

          <div class="bg-white rounded-2xl shadow-sm p-6 transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-slate-500">Pending Orders</p>
                <p class="text-3xl font-bold text-amber-500">{{ pendingOrdersCount }}</p>
                 <p class="text-sm text-amber-600 mt-2 flex items-center font-semibold">
                  <TrendingUpIcon class="w-4 h-4 inline mr-1.5" />
                   Awaiting action
                </p>
              </div>
              <div class="w-12 h-12 bg-amber-50 rounded-xl flex items-center justify-center">
                <ShoppingBagIcon class="w-6 h-6 text-amber-500" />
              </div>
            </div>
          </div>

          <div class="bg-white rounded-2xl shadow-sm p-6 transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-slate-500">Total Customers</p>
                <p class="text-3xl font-bold text-slate-800">{{ stats.customers }}</p>
                <p class="text-sm text-emerald-600 mt-2 flex items-center font-semibold">
                  <UserAddIcon class="w-4 h-4 inline mr-1.5" />
                  +{{ stats.newCustomersToday }} today
                </p>
              </div>
              <div class="w-12 h-12 bg-emerald-50 rounded-xl flex items-center justify-center">
                <UsersIcon class="w-6 h-6 text-emerald-500" />
              </div>
            </div>
          </div>

          <!-- Total Sales (Green) -->
          <div class="bg-white rounded-2xl shadow-sm p-6 transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-emerald-600 flex items-center">
                  Total Revenue
                  <span class="ml-2 text-xs text-slate-400 font-normal">(all paid, shipped, delivered)</span>
                </p>
                <p class="text-3xl font-black text-emerald-700">${{ formatCurrency(stats.totalRevenue) }}</p>
                <p class="text-sm text-emerald-600 mt-2 flex items-center font-semibold">
                  <TrendingUpIcon class="w-4 h-4 inline mr-1.5" />
                  +${{ formatCurrency(stats.revenueToday) }} today
                </p>
              </div>
              <div class="w-12 h-12 bg-emerald-50 rounded-xl flex items-center justify-center">
                <CurrencyDollarIcon class="w-6 h-6 text-emerald-500" />
              </div>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          
          <div class="lg:col-span-2 bg-white rounded-2xl shadow-sm">
            <div class="p-6 border-b border-slate-100">
              <h2 class="text-xl font-bold text-slate-800">Sales Performance</h2>
              <p class="text-sm text-slate-500 mt-1">Last 7 days revenue trend</p>
            </div>
            <div class="p-6">
              <div class="h-64 flex items-end justify-center">
                <div class="w-full">
                  <div class="grid grid-cols-7 gap-4">
                    <div 
                      v-for="(day, index) in weeklySales" 
                      :key="index"
                      class="text-center flex flex-col items-center justify-end"
                    >
                      <div class="text-sm font-semibold text-slate-700 mb-1">${{ formatCurrency(day.amount) }}</div>
                      <div 
                        class="bg-gradient-to-t from-green-400 to-green-500 w-full rounded-t-lg transition-all duration-300"
                        :style="{ height: `${Math.max(day.amount / maxWeeklySales * 150, 4)}px` }"
                      ></div>
                      <div class="text-xs text-slate-400 font-medium uppercase tracking-wider mt-2">{{ day.day }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-2xl shadow-sm">
            <div class="p-6 border-b border-slate-100">
              <h2 class="text-xl font-bold text-slate-800">Quick Stats</h2>
            </div>
            <div class="p-6 divide-y divide-slate-100">
              <div class="flex items-center justify-between py-3">
                <div>
                  <p class="text-sm text-slate-500">Avg Order Value</p>
                  <p class="text-2xl font-bold text-slate-800">${{ formatCurrency(stats.averageOrderValue) }}</p>
                </div>
                <div class="w-10 h-10 bg-sky-50 rounded-lg flex items-center justify-center">
                  <ShoppingCartIcon class="w-5 h-5 text-sky-500" />
                </div>
              </div>
              <div class="flex items-center justify-between py-3">
                <div>
                  <p class="text-sm text-slate-500">Conversion Rate</p>
                  <p class="text-2xl font-bold text-slate-800">{{ stats.conversionRate }}%</p>
                </div>
                <div class="w-10 h-10 bg-emerald-50 rounded-lg flex items-center justify-center">
                  <ChartBarIcon class="w-5 h-5 text-emerald-500" />
                </div>
              </div>
              <div class="flex items-center justify-between py-3">
                <div>
                  <p class="text-sm text-slate-500">Top Seller</p>
                  <p class="text-lg font-semibold text-slate-800 truncate max-w-40">{{ stats.topProduct?.name || 'N/A' }}</p>
                </div>
                <div class="w-10 h-10 bg-violet-50 rounded-lg flex items-center justify-center">
                  <StarIcon class="w-5 h-5 text-violet-500" />
                </div>
              </div>
              <div class="flex items-center justify-between pt-3">
                <div>
                  <p class="text-sm text-slate-500">Active Discounts</p>
                  <p class="text-2xl font-bold text-slate-800">{{ stats.activeDiscounts }}</p>
                </div>
                <div class="w-10 h-10 bg-rose-50 rounded-lg flex items-center justify-center">
                  <TagIcon class="w-5 h-5 text-rose-500" />
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
          
          <div class="bg-white rounded-2xl shadow-sm">
            <div class="p-6 border-b border-slate-100">
              <div class="flex items-center justify-between">
                <h2 class="text-xl font-bold text-slate-800">Recent Orders</h2>
                <router-link to="/dashboard/orders" class="text-sm text-green-600 hover:text-green-700 font-semibold">
                  View all
                </router-link>
              </div>
            </div>
            <div v-if="recentOrders.length === 0" class="p-6">
              <div class="text-center py-8 text-slate-400">
                <ShoppingBagIcon class="w-12 h-12 mx-auto mb-4 text-slate-300" />
                <p class="font-semibold text-slate-600">No orders yet</p>
                <p class="text-sm">New orders will appear here once placed.</p>
              </div>
            </div>
            <div v-else class="divide-y divide-slate-100">
              <div v-for="order in recentOrders.slice(0, 5)" :key="order.id" class="flex items-center justify-between p-4 transition-colors duration-200 hover:bg-slate-50/75 cursor-pointer" @click="goToOrder(order.id)">
                <div class="flex items-center space-x-4">
                  <div class="w-10 h-10 bg-emerald-50 rounded-lg flex-shrink-0 flex items-center justify-center">
                    <ShoppingBagIcon class="w-5 h-5 text-emerald-500" />
                  </div>
                  <div>
                    <p class="font-semibold text-slate-800">#{{ order.orderNumber }}</p>
                    <p class="text-sm text-slate-500">{{ formatDateTime(order.createdAt) }}</p>
                  </div>
                </div>
                <div class="text-right">
                  <p class="font-semibold text-slate-800">${{ formatCurrency(order.total) }}</p>
                  <span :class="getStatusClass(order.status)" class="inline-flex px-2.5 py-1 rounded-full text-xs font-semibold leading-none">
                    {{ formatStatus(order.status) }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-2xl shadow-sm">
            <div class="p-6 border-b border-slate-100">
              <div class="flex items-center justify-between">
                <h2 class="text-xl font-bold text-slate-800">Recent Products</h2>
                <router-link to="/dashboard/products" class="text-sm text-green-600 hover:text-green-700 font-semibold">
                  View all
                </router-link>
              </div>
            </div>
             <div v-if="recentProducts.length === 0" class="p-6">
               <div class="text-center py-8 text-slate-400">
                 <CubeIcon class="w-12 h-12 mx-auto mb-4 text-slate-300" />
                 <p class="font-semibold text-slate-600">No products yet</p>
                 <p class="text-sm">Add your first product to get started.</p>
               </div>
             </div>
            <div v-else class="divide-y divide-slate-100">
              <div v-for="product in recentProducts" :key="product.id" class="flex items-center space-x-4 p-3 transition-colors duration-200 hover:bg-slate-50/75 cursor-pointer group" @click="goToProduct(product.id)">
                <img :src="getProductImage(product)" :alt="product.name" class="w-14 h-14 rounded-lg object-cover border border-slate-100 group-hover:ring-2 group-hover:ring-green-300 transition" @error="$event.target.src='/placeholder-product.jpg'" />
                <div class="flex-1 min-w-0">
                  <p class="font-semibold text-slate-800 truncate">{{ product.name }}</p>
                  <p class="text-sm text-slate-500">${{ formatCurrency(product.price) }}</p>
                  <span v-if="getProductStock(product) === 0" class="inline-block mt-1 px-2 py-0.5 bg-rose-100 text-rose-700 text-xs rounded-full font-semibold">Out of Stock</span>
                </div>
                <div class="text-right flex-shrink-0">
                  <p class="text-sm font-medium text-slate-600">Stock: {{ getProductStock(product) }}</p>
                  <p class="text-xs text-slate-400">{{ formatDate(product.createdAt) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-2xl shadow-sm">
          <div class="p-6 border-b border-slate-100">
            <h2 class="text-xl font-bold text-slate-800">Alerts & Notifications</h2>
          </div>
          <div class="p-6">
            <div v-if="alerts.length === 0" class="text-center py-8 text-slate-400">
              <CheckCircleIcon class="w-12 h-12 mx-auto mb-4 text-emerald-400" />
              <p class="font-semibold text-slate-600">All good!</p>
              <p class="text-sm">No new alerts to show right now.</p>
            </div>
            <div v-else class="space-y-4">
              <div v-for="alert in alerts" :key="alert.id" :class="getAlertClass(alert.type)" class="flex items-start p-4 rounded-lg">
                <component :is="getAlertIcon(alert.type)" class="w-6 h-6 mr-3 flex-shrink-0 mt-0.5" />
                <div class="flex-1">
                  <p class="font-semibold">{{ alert.title }}</p>
                  <p class="text-sm opacity-90">{{ alert.message }}</p>
                </div>
                <button @click="dismissAlert(alert.id)" class="ml-4 p-1 rounded-full hover:bg-black/10 transition-colors">
                   <XIcon class="w-4 h-4" />
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
  XCircleIcon
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
  return classMap[status] || 'bg-slate-100 text-slate-700'
}


function getAlertClass(type) {
  const classMap = {
    'warning': 'bg-amber-50 text-amber-800',
    'error': 'bg-rose-50 text-rose-800',
    'info': 'bg-sky-50 text-sky-800',
    'success': 'bg-emerald-50 text-emerald-800'
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
/* Custom styles if needed */
@media (max-width: 640px) {
  .min-w-0 { min-width: 0 !important; }
  .min-w-3rem { min-width: 3rem !important; }
  .max-w-3rem { max-width: 3rem !important; }
  .min-h-3rem { min-height: 3rem !important; }
  .max-h-3rem { max-height: 3rem !important; }
}
</style>
