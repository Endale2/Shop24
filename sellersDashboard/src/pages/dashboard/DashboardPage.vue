<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
        Dashboard Overview
      </h1>
      <p class="text-lg text-gray-600 mt-2">
        Monitor your shop's performance and key metrics
      </p>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <div class="animate-spin h-10 w-10 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
        <p>Loading dashboard data...</p>
      </div>
    </div>

    <!-- Dashboard Content -->
    <div v-else class="space-y-8">
      <!-- Summary Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- Total Products -->
        <div class="bg-white rounded-xl shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow duration-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">Total Products</p>
              <p class="text-3xl font-bold text-gray-900">{{ stats.products }}</p>
              <p class="text-sm text-blue-600 mt-1">
                <CubeIcon class="w-4 h-4 inline mr-1" />
                {{ stats.lowStockProducts }} low stock
              </p>
            </div>
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
              <CubeIcon class="w-6 h-6 text-green-600" />
            </div>
          </div>
        </div>

        <!-- Total Orders -->
        <div class="bg-white rounded-xl shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow duration-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">Total Orders <span class="text-xs text-gray-400" title="Only paid, shipped, or delivered orders are counted.">(paid/shipped/delivered)</span></p>
              <p class="text-3xl font-bold text-gray-900">{{ stats.orders }}</p>
              <p class="text-sm text-green-600 mt-1">
                <TrendingUpIcon class="w-4 h-4 inline mr-1" />
                +{{ stats.newOrdersToday }} today
              </p>
            </div>
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
              <ShoppingBagIcon class="w-6 h-6 text-blue-600" />
            </div>
          </div>
        </div>
        <!-- Pending Orders -->
        <div class="bg-white rounded-xl shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow duration-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">Pending Orders</p>
              <p class="text-3xl font-bold text-yellow-600">{{ pendingOrdersCount }}</p>
              <p class="text-sm text-yellow-600 mt-1">
                <TrendingUpIcon class="w-4 h-4 inline mr-1" />
                Awaiting payment or fulfillment
              </p>
            </div>
            <div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
              <ShoppingBagIcon class="w-6 h-6 text-yellow-600" />
            </div>
          </div>
        </div>

        <!-- Total Customers -->
        <div class="bg-white rounded-xl shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow duration-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">Total Customers</p>
              <p class="text-3xl font-bold text-gray-900">{{ stats.customers }}</p>
              <p class="text-sm text-green-600 mt-1">
                <UserAddIcon class="w-4 h-4 inline mr-1" />
                +{{ stats.newCustomersToday }} today
              </p>
            </div>
            <div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
              <UsersIcon class="w-6 h-6 text-yellow-600" />
            </div>
          </div>
        </div>

        <!-- Total Sales -->
        <div class="bg-white rounded-xl shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow duration-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">Total Sales</p>
              <p class="text-3xl font-bold text-gray-900">${{ formatCurrency(getPaidRevenue()) }}</p>
              <p class="text-sm text-green-600 mt-1">
                <TrendingUpIcon class="w-4 h-4 inline mr-1" />
                +${{ formatCurrency(stats.salesToday) }} today
              </p>
            </div>
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
              <CurrencyDollarIcon class="w-6 h-6 text-purple-600" />
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Metrics -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Sales Chart -->
        <div class="lg:col-span-2 bg-white rounded-xl shadow-md border border-gray-200">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Sales Performance</h2>
            <p class="text-sm text-gray-600 mt-1">Last 7 days revenue trend</p>
          </div>
          <div class="p-6">
            <div class="h-64 flex items-center justify-center">
              <div class="text-center">
                <div class="text-4xl font-bold text-green-600 mb-2">${{ formatCurrency(stats.sales) }}</div>
                <div class="text-lg text-gray-600 mb-4">Total Revenue</div>
                <div class="grid grid-cols-7 gap-2">
                  <div 
                    v-for="(day, index) in weeklySales" 
                    :key="index"
                    class="text-center"
                  >
                    <div class="text-xs text-gray-500 mb-1">{{ day.day }}</div>
                    <div 
                      class="bg-green-200 rounded-t"
                      :style="{ height: `${Math.max(day.amount / maxWeeklySales * 100, 4)}px` }"
                    ></div>
                    <div class="text-xs text-gray-600 mt-1">${{ formatCurrency(day.amount) }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Quick Stats -->
        <div class="bg-white rounded-xl shadow-md border border-gray-200">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Quick Stats</h2>
          </div>
          <div class="p-6 space-y-6">
            <!-- Average Order Value -->
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Avg Order Value</p>
                <p class="text-2xl font-bold text-gray-900">${{ formatCurrency(stats.averageOrderValue) }}</p>
              </div>
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                <ShoppingCartIcon class="w-5 h-5 text-blue-600" />
              </div>
            </div>

            <!-- Conversion Rate -->
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Conversion Rate</p>
                <p class="text-2xl font-bold text-gray-900">{{ stats.conversionRate }}%</p>
              </div>
              <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                <ChartBarIcon class="w-5 h-5 text-green-600" />
              </div>
            </div>

            <!-- Top Selling Product -->
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Top Seller</p>
                <p class="text-lg font-semibold text-gray-900 truncate max-w-32">{{ stats.topProduct?.name || 'N/A' }}</p>
              </div>
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
                <StarIcon class="w-5 h-5 text-purple-600" />
              </div>
            </div>

            <!-- Active Discounts -->
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Active Discounts</p>
                <p class="text-2xl font-bold text-gray-900">{{ stats.activeDiscounts }}</p>
              </div>
              <div class="w-10 h-10 bg-orange-100 rounded-lg flex items-center justify-center">
                <TagIcon class="w-5 h-5 text-orange-600" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Activity -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- Recent Orders -->
        <div class="bg-white rounded-xl shadow-md border border-gray-200">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">Recent Orders</h2>
              <router-link 
                to="/orders" 
                class="text-sm text-green-600 hover:text-green-700 font-medium"
              >
                View all orders
              </router-link>
            </div>
          </div>
          <div class="p-6">
            <div v-if="recentOrders.length === 0" class="text-center py-8 text-gray-500">
              <ShoppingBagIcon class="w-12 h-12 mx-auto mb-4 text-gray-300" />
              <p>No orders yet</p>
              <p class="text-sm">Orders will appear here once customers start shopping</p>
            </div>
            <div v-else class="space-y-4">
              <div 
                v-for="order in recentOrders.slice(0, 5)" 
                :key="order.id"
                class="flex items-center justify-between p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors duration-200 cursor-pointer"
                @click="goToOrder(order.id)"
              >
                <div class="flex items-center space-x-3">
                  <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                    <ShoppingBagIcon class="w-5 h-5 text-green-600" />
                  </div>
                  <div>
                    <p class="font-medium text-gray-900">#{{ order.orderNumber }}</p>
                    <p class="text-sm text-gray-600">{{ formatDateTime(order.createdAt) }}</p>
                  </div>
                </div>
                <div class="text-right">
                  <p class="font-semibold text-gray-900">${{ formatCurrency(order.total) }}</p>
                  <span :class="getStatusClass(order.status)" class="inline-flex px-2 py-1 rounded-full text-xs font-medium">
                    {{ formatStatus(order.status) }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Products -->
        <div class="bg-white rounded-xl shadow-md border border-gray-200">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">Recent Products</h2>
              <router-link 
                to="/products" 
                class="text-sm text-green-600 hover:text-green-700 font-medium"
              >
                View all products
              </router-link>
            </div>
          </div>
          <div class="p-6">
            <div v-if="recentProducts.length === 0" class="text-center py-8 text-gray-500">
              <CubeIcon class="w-12 h-12 mx-auto mb-4 text-gray-300" />
              <p>No products yet</p>
              <p class="text-sm">Add your first product to get started</p>
            </div>
            <div v-else class="space-y-4">
              <div 
                v-for="product in recentProducts.slice(0, 5)" 
                :key="product.id"
                class="flex items-center space-x-3 p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors duration-200 cursor-pointer group"
                @click="goToProduct(product.id)"
              >
                <img 
                  :src="getProductImage(product)" 
                  :alt="product.name"
                  class="w-12 h-12 min-w-[3rem] min-h-[3rem] max-w-[3rem] max-h-[3rem] rounded-lg object-cover border border-gray-200 group-hover:ring-2 group-hover:ring-green-400 transition"
                  @error="$event.target.src='/placeholder-product.jpg'"
                />
                <div class="flex-1 min-w-0">
                  <p class="font-medium text-gray-900 truncate">{{ product.name }}</p>
                  <p class="text-sm text-gray-600">${{ formatCurrency(product.price) }}</p>
                  <span v-if="product.stock === 0" class="inline-block mt-1 px-2 py-0.5 bg-red-100 text-red-700 text-xs rounded-full font-semibold">Out of Stock</span>
                </div>
                <div class="text-right flex flex-col items-end min-w-[70px]">
                  <p class="text-sm text-gray-600">Stock: {{ product.stock }}</p>
                  <p class="text-xs text-gray-500">{{ formatDate(product.createdAt) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerts & Notifications -->
      <div class="bg-white rounded-xl shadow-md border border-gray-200">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">Alerts & Notifications</h2>
        </div>
        <div class="p-6">
          <div v-if="alerts.length === 0" class="text-center py-8 text-gray-500">
            <CheckCircleIcon class="w-12 h-12 mx-auto mb-4 text-green-300" />
            <p>All good! No alerts at the moment</p>
          </div>
          <div v-else class="space-y-4">
            <div 
              v-for="alert in alerts" 
              :key="alert.id"
              :class="getAlertClass(alert.type)"
              class="flex items-center p-4 rounded-lg"
            >
              <component :is="getAlertIcon(alert.type)" class="w-5 h-5 mr-3" />
              <div class="flex-1">
                <p class="font-medium">{{ alert.title }}</p>
                <p class="text-sm opacity-90">{{ alert.message }}</p>
              </div>
              <button 
                @click="dismissAlert(alert.id)"
                class="text-sm opacity-70 hover:opacity-100"
              >
                Dismiss
              </button>
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
import { orderService } from '@/services/order'
import { productService } from '@/services/product'
import { customerService } from '@/services/customer'
import { discountService } from '@/services/discount'
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
  sales: 0,
  salesToday: 0,
  newOrdersToday: 0,
  newCustomersToday: 0,
  lowStockProducts: 0,
  averageOrderValue: 0,
  conversionRate: 0,
  topProduct: null,
  activeDiscounts: 0
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

// Add computed for pending orders
const pendingOrdersCount = computed(() => recentOrders.value.filter(order => (order.status || '').toLowerCase() === 'pending').length)

// Helper: filter orders by status considered as 'sold'
function filterSoldOrders(orders) {
  return orders.filter(order => ['paid', 'shipped', 'delivered'].includes((order.status || '').toLowerCase()))
}

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
  if (product.main_image) return product.main_image
  if (Array.isArray(product.images) && product.images.length > 0) return product.images[0]
  return '/placeholder-product.jpg'
}

// Methods
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
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800'
  }
  return classMap[status] || 'bg-gray-100 text-gray-800'
}

function getAlertClass(type) {
  const classMap = {
    'warning': 'bg-yellow-50 text-yellow-800 border border-yellow-200',
    'error': 'bg-red-50 text-red-800 border border-red-200',
    'info': 'bg-blue-50 text-blue-800 border border-blue-200',
    'success': 'bg-green-50 text-green-800 border border-green-200'
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

// --- Revenue calculation: only count paid, shipped, delivered ---
function getPaidRevenue() {
  return filterSoldOrders(recentOrders.value)
    .reduce((sum, order) => sum + (order.total || 0), 0)
}

// Generate weekly sales data
function generateWeeklySales(orders) {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const today = new Date()
  const weekData = []

  for (let i = 6; i >= 0; i--) {
    const date = new Date(today)
    date.setDate(date.getDate() - i)
    date.setHours(0, 0, 0, 0)

    const dayOrders = filterSoldOrders(orders).filter(order => {
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

// Generate alerts based on data
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
  const now = new Date()
  const expiringDiscounts = discounts.filter(discount => {
    if (!discount.active || !discount.endAt) return false
    const endDate = new Date(discount.endAt)
    const daysUntilExpiry = (endDate - now) / (1000 * 60 * 60 * 24)
    return daysUntilExpiry <= 7 && daysUntilExpiry > 0
  })

  if (expiringDiscounts.length > 0) {
    alerts.push({
      id: 'expiring-discounts',
      type: 'warning',
      title: 'Expiring Discounts',
      message: `${expiringDiscounts.length} discount(s) will expire within 7 days`
    })
  }

  return alerts
}

// Load dashboard data
async function loadDashboardData() {
  if (!shopId.value) {
    console.warn('No active shop selected')
    loading.value = false
    return
  }

  try {
    // Load data in parallel
    const [orders, products, customers, discounts] = await Promise.all([
      orderService.fetchAllByShop(shopId.value),
      productService.fetchAllByShop(shopId.value),
      customerService.fetchAll(shopId.value),
      discountService.fetchAllByShop(shopId.value)
    ])

    // Process orders
    recentOrders.value = orders.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
    
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    
    const ordersToday = orders.filter(order => {
      const orderDate = new Date(order.createdAt)
      orderDate.setHours(0, 0, 0, 0)
      return orderDate.getTime() === today.getTime()
    })

    // Process products
    recentProducts.value = products.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
    const lowStockProducts = products.filter(p => (p.stock || 0) < 10)

    // Process customers
    const customersToday = customers.filter(customer => {
      const customerDate = new Date(customer.createdAt)
      customerDate.setHours(0, 0, 0, 0)
      return customerDate.getTime() === today.getTime()
    })

    // Process discounts
    const activeDiscounts = discounts.filter(d => d.active)

    // Calculate stats
    const soldOrders = filterSoldOrders(orders)
    const totalSales = soldOrders.reduce((sum, order) => sum + (order.total || 0), 0)
    const salesToday = soldOrders.filter(order => {
      const orderDate = new Date(order.createdAt)
      orderDate.setHours(0, 0, 0, 0)
      return orderDate.getTime() === today.getTime()
    }).reduce((sum, order) => sum + (order.total || 0), 0)
    const averageOrderValue = soldOrders.length > 0 ? totalSales / soldOrders.length : 0

    // Find top product (by stock for now, could be by sales)
    const topProduct = products.length > 0 ? products.reduce((top, current) => 
      (current.stock || 0) > (top.stock || 0) ? current : top
    ) : null

    stats.value = {
      products: products.length,
      orders: soldOrders.length,
      customers: customers.length,
      sales: totalSales,
      salesToday: salesToday,
      newOrdersToday: ordersToday.length, // Optionally filter by sold status if needed
      newCustomersToday: customersToday.length,
      lowStockProducts: lowStockProducts.length,
      averageOrderValue: averageOrderValue,
      conversionRate: customers.length > 0 ? Math.round((soldOrders.length / customers.length) * 100) : 0,
      topProduct: topProduct,
      activeDiscounts: activeDiscounts.length
    }

    // Generate weekly sales data
    weeklySales.value = generateWeeklySales(orders)

    // Generate alerts
    alerts.value = generateAlerts(products, orders, discounts)

  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
}

// Watch for shop changes
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
  .min-w-[3rem] { min-width: 3rem !important; }
  .max-w-[3rem] { max-width: 3rem !important; }
  .min-h-[3rem] { min-height: 3rem !important; }
  .max-h-[3rem] { max-height: 3rem !important; }
}
</style>
