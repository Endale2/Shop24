<template>
  <div class="p-4 sm:p-6 max-w-7xl mx-auto font-sans">
    <!-- Header & Period Selector -->
    <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0 mb-8">
      <div>
        <h1 class="text-3xl sm:text-4xl font-extrabold text-gray-900 leading-tight">
          Analytics Dashboard
        </h1>
        <p class="text-lg text-gray-600 mt-2">
          Deep insights into your shop's performance
        </p>
      </div>
      <div class="inline-flex bg-white p-1 rounded-lg shadow-sm border border-gray-200">
        <button
          v-for="period in periods"
          :key="period.id"
          @click="changePeriod(period.id)"
          :class="[
            'px-4 py-2 text-sm font-medium rounded-md transition-all duration-200',
            selectedPeriod === period.id
              ? 'bg-green-600 text-white shadow-md'
              : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
          ]"
        >
          {{ period.label }}
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 text-gray-600 text-lg">
      <div class="flex flex-col items-center justify-center">
        <div class="animate-spin h-10 w-10 text-green-500 mb-4 border-4 border-green-200 border-t-green-500 rounded-full"></div>
        <p>Loading analytics data...</p>
      </div>
    </div>

    <!-- Analytics Content -->
    <div v-else class="space-y-8">
      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <div
          v-for="stat in analyticsStats"
          :key="stat.title"
          class="bg-white p-6 rounded-xl shadow-md border border-gray-200 hover:shadow-lg transition-shadow duration-200 border-l-4"
          :class="stat.borderColor"
        >
          <div class="flex items-center space-x-4">
            <div class="p-3 rounded-lg" :class="stat.iconBg">
              <component :is="stat.icon" class="h-6 w-6" :class="stat.iconColor" />
            </div>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-600">{{ stat.title }}</p>
              <p class="text-2xl font-bold text-gray-900">{{ stat.value }}</p>
              <p
                class="text-sm font-medium"
                :class="stat.change >= 0 ? 'text-green-600' : 'text-red-600'"
              >
                <component :is="stat.change >= 0 ? TrendingUpIcon : TrendingDownIcon" class="w-4 h-4 inline mr-1" />
                {{ stat.change >= 0 ? '+' : '' }}{{ stat.change }}% vs last period
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Revenue Over Time -->
        <div class="lg:col-span-2 bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-semibold text-gray-900">Revenue Over Time</h2>
            <div class="text-sm text-gray-600">
              {{ getPeriodLabel(selectedPeriod) }}
            </div>
          </div>
          <div class="h-80">
            <Line
              v-if="revenueChartData"
              :data="revenueChartData"
              :options="lineChartOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
              <p class="text-gray-500">No revenue data available</p>
            </div>
          </div>
        </div>

        <!-- Sales by Category -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Sales by Category</h2>
          <div class="h-80">
            <Doughnut
              v-if="categoryChartData"
              :data="categoryChartData"
              :options="doughnutChartOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
              <p class="text-gray-500">No category data available</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Additional Charts -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- Orders Over Time -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Orders Over Time</h2>
          <div class="h-64">
            <Line
              v-if="ordersChartData"
              :data="ordersChartData"
              :options="lineChartOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
              <p class="text-gray-500">No order data available</p>
            </div>
          </div>
        </div>

        <!-- Customer Growth -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Customer Growth</h2>
          <div class="h-64">
            <Line
              v-if="customersChartData"
              :data="customersChartData"
              :options="lineChartOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg flex items-center justify-center">
              <p class="text-gray-500">No customer data available</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Top Products & Recent Orders -->
      <div class="grid grid-cols-1 lg:grid-cols-5 gap-8">
        <!-- Top Selling Products -->
        <div class="lg:col-span-2 bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-semibold text-gray-900">Top Selling Products</h2>
            <router-link 
              to="/products" 
              class="text-sm text-green-600 hover:text-green-700 font-medium"
            >
              View all products
            </router-link>
          </div>
          <div v-if="topProducts.length === 0" class="text-center py-8 text-gray-500">
            <CubeIcon class="w-12 h-12 mx-auto mb-4 text-gray-300" />
            <p>No products yet</p>
            <p class="text-sm">Add products to see analytics</p>
          </div>
          <div v-else class="space-y-4">
            <div
              v-for="(product, index) in topProducts"
              :key="product.id"
              class="flex items-center space-x-4 p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors duration-200 cursor-pointer"
              @click="goToProduct(product.id)"
            >
              <div class="flex-shrink-0">
                <div class="w-12 h-12 bg-gray-200 rounded-lg overflow-hidden">
                  <img 
                    :src="product.mainImage || '/placeholder-product.jpg'" 
                    :alt="product.name"
                    class="w-full h-full object-cover"
                  />
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <p class="font-medium text-gray-900 truncate">{{ product.name }}</p>
                <p class="text-sm text-gray-500">{{ product.category }}</p>
              </div>
              <div class="text-right">
                <p class="font-semibold text-gray-900">${{ formatCurrency(product.totalRevenue) }}</p>
                <p class="text-sm text-gray-500">{{ product.totalSold }} sold</p>
              </div>
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <span class="text-sm font-bold text-green-600">#{{ index + 1 }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Orders Table -->
        <div class="lg:col-span-3 bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-semibold text-gray-900">Recent Orders</h2>
            <router-link 
              to="/orders" 
              class="text-sm text-green-600 hover:text-green-700 font-medium"
            >
              View all orders
            </router-link>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead class="bg-gray-50">
                <tr>
                  <th class="p-3 text-left font-medium text-gray-600">Order ID</th>
                  <th class="p-3 text-left font-medium text-gray-600">Customer</th>
                  <th class="p-3 text-left font-medium text-gray-600">Amount</th>
                  <th class="p-3 text-left font-medium text-gray-600">Status</th>
                  <th class="p-3 text-left font-medium text-gray-600">Date</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200">
                <tr
                  v-for="order in recentOrders"
                  :key="order.id"
                  class="hover:bg-gray-50 cursor-pointer transition-colors duration-200"
                  @click="goToOrder(order.id)"
                >
                  <td class="p-3 font-mono text-gray-500">#{{ order.orderNumber }}</td>
                  <td class="p-3 text-gray-800">{{ order.customerName }}</td>
                  <td class="p-3 font-semibold text-gray-900">${{ formatCurrency(order.total) }}</td>
                  <td class="p-3">
                    <span
                      class="px-2 py-1 text-xs font-semibold rounded-full"
                      :class="getStatusClass(order.status)"
                    >
                      {{ formatStatus(order.status) }}
                    </span>
                  </td>
                  <td class="p-3 text-gray-600">{{ formatDate(order.createdAt) }}</td>
                </tr>
              </tbody>
            </table>
            <div v-if="recentOrders.length === 0" class="text-center py-8 text-gray-500">
              <ShoppingBagIcon class="w-12 h-12 mx-auto mb-4 text-gray-300" />
              <p>No orders yet</p>
              <p class="text-sm">Orders will appear here once customers start shopping</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Metrics -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Conversion Funnel -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Conversion Funnel</h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Total Visitors</span>
              <span class="font-semibold text-gray-900">{{ conversionFunnel.visitors }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-600 h-2 rounded-full" :style="{ width: `${conversionFunnel.visitorRate}%` }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Add to Cart</span>
              <span class="font-semibold text-gray-900">{{ conversionFunnel.addToCart }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-green-600 h-2 rounded-full" :style="{ width: `${conversionFunnel.addToCartRate}%` }"></div>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Purchases</span>
              <span class="font-semibold text-gray-900">{{ conversionFunnel.purchases }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-purple-600 h-2 rounded-full" :style="{ width: `${conversionFunnel.purchaseRate}%` }"></div>
            </div>
          </div>
        </div>

        <!-- Customer Insights -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Customer Insights</h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">New Customers</span>
              <span class="font-semibold text-green-600">+{{ customerInsights.newCustomers }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Returning Customers</span>
              <span class="font-semibold text-blue-600">{{ customerInsights.returningCustomers }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Avg Order Value</span>
              <span class="font-semibold text-gray-900">${{ formatCurrency(customerInsights.avgOrderValue) }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Customer Lifetime Value</span>
              <span class="font-semibold text-purple-600">${{ formatCurrency(customerInsights.lifetimeValue) }}</span>
            </div>
          </div>
        </div>

        <!-- Inventory Status -->
        <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">Inventory Status</h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Total Products</span>
              <span class="font-semibold text-gray-900">{{ inventoryStatus.totalProducts }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">In Stock</span>
              <span class="font-semibold text-green-600">{{ inventoryStatus.inStock }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Low Stock</span>
              <span class="font-semibold text-yellow-600">{{ inventoryStatus.lowStock }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Out of Stock</span>
              <span class="font-semibold text-red-600">{{ inventoryStatus.outOfStock }}</span>
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
const orders = ref([])
const products = ref([])
const customers = ref([])
const discounts = ref([])

// Period options
const periods = [
  { id: 'last_7_days', label: '7 Days' },
  { id: 'last_30_days', label: '30 Days' },
  { id: 'last_90_days', label: '90 Days' },
  { id: 'last_12_months', label: '12 Months' },
]

// Chart options
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

// Computed properties
const analyticsStats = computed(() => {
  const periodData = getPeriodData()
  const previousPeriodData = getPreviousPeriodData()
  
  return [
    {
      title: 'Total Revenue',
      value: `$${formatCurrency(periodData.totalRevenue)}`,
      change: calculatePercentageChange(periodData.totalRevenue, previousPeriodData.totalRevenue),
      icon: CurrencyDollarIcon,
      iconBg: 'bg-green-100',
      iconColor: 'text-green-600',
      borderColor: 'border-green-500'
    },
    {
      title: 'Total Orders',
      value: periodData.totalOrders.toLocaleString(),
      change: calculatePercentageChange(periodData.totalOrders, previousPeriodData.totalOrders),
      icon: ShoppingBagIcon,
      iconBg: 'bg-blue-100',
      iconColor: 'text-blue-600',
      borderColor: 'border-blue-500'
    },
    {
      title: 'New Customers',
      value: periodData.newCustomers.toLocaleString(),
      change: calculatePercentageChange(periodData.newCustomers, previousPeriodData.newCustomers),
      icon: UsersIcon,
      iconBg: 'bg-indigo-100',
      iconColor: 'text-indigo-600',
      borderColor: 'border-indigo-500'
    },
    {
      title: 'Avg Order Value',
      value: `$${formatCurrency(periodData.avgOrderValue)}`,
      change: calculatePercentageChange(periodData.avgOrderValue, previousPeriodData.avgOrderValue),
      icon: ChartBarIcon,
      iconBg: 'bg-amber-100',
      iconColor: 'text-amber-600',
      borderColor: 'border-amber-500'
    }
  ]
})

const revenueChartData = computed(() => {
  const data = getChartData('revenue')
  if (!data.labels.length) return null
  
  return {
    labels: data.labels,
    datasets: [{
      label: 'Revenue',
      data: data.values,
      borderColor: 'rgb(16, 185, 129)',
      backgroundColor: 'rgba(16, 185, 129, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

const ordersChartData = computed(() => {
  const data = getChartData('orders')
  if (!data.labels.length) return null
  
  return {
    labels: data.labels,
    datasets: [{
      label: 'Orders',
      data: data.values,
      borderColor: 'rgb(59, 130, 246)',
      backgroundColor: 'rgba(59, 130, 246, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

const customersChartData = computed(() => {
  const data = getChartData('customers')
  if (!data.labels.length) return null
  
  return {
    labels: data.labels,
    datasets: [{
      label: 'Customers',
      data: data.values,
      borderColor: 'rgb(139, 92, 246)',
      backgroundColor: 'rgba(139, 92, 246, 0.1)',
      tension: 0.3,
      fill: true
    }]
  }
})

const categoryChartData = computed(() => {
  const categoryData = getCategoryData()
  if (!categoryData.labels.length) return null
  
  return {
    labels: categoryData.labels,
    datasets: [{
      data: categoryData.values,
      backgroundColor: ['#10B981', '#3B82F6', '#8B5CF6', '#F59E0B', '#EF4444', '#06B6D4']
    }]
  }
})

const topProducts = computed(() => {
  return products.value
    .map(product => ({
      ...product,
      totalRevenue: calculateProductRevenue(product.id),
      totalSold: calculateProductSold(product.id)
    }))
    .sort((a, b) => b.totalRevenue - a.totalRevenue)
    .slice(0, 5)
})

const recentOrders = computed(() => {
  return orders.value
    .sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
    .slice(0, 10)
    .map(order => ({
      ...order,
      customerName: getCustomerName(order.customerId)
    }))
})

const conversionFunnel = computed(() => {
  const totalVisitors = customers.value.length * 3 // Estimate
  const addToCart = Math.floor(totalVisitors * 0.15) // 15% add to cart rate
  const purchases = orders.value.length
  
  return {
    visitors: totalVisitors.toLocaleString(),
    addToCart: addToCart.toLocaleString(),
    purchases: purchases.toLocaleString(),
    visitorRate: 100,
    addToCartRate: (addToCart / totalVisitors * 100).toFixed(1),
    purchaseRate: (purchases / totalVisitors * 100).toFixed(1)
  }
})

const customerInsights = computed(() => {
  const periodData = getPeriodData()
  const avgOrderValue = periodData.totalOrders > 0 ? periodData.totalRevenue / periodData.totalOrders : 0
  const lifetimeValue = avgOrderValue * 2.5 // Estimate
  
  return {
    newCustomers: periodData.newCustomers,
    returningCustomers: customers.value.length - periodData.newCustomers,
    avgOrderValue: avgOrderValue,
    lifetimeValue: lifetimeValue
  }
})

const inventoryStatus = computed(() => {
  const totalProducts = products.value.length
  const inStock = products.value.filter(p => (p.stock || 0) > 10).length
  const lowStock = products.value.filter(p => (p.stock || 0) <= 10 && (p.stock || 0) > 0).length
  const outOfStock = products.value.filter(p => (p.stock || 0) === 0).length
  
  return {
    totalProducts,
    inStock,
    lowStock,
    outOfStock
  }
})

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

function getPeriodData() {
  const now = new Date()
  const days = getDaysForPeriod(selectedPeriod.value)
  const startDate = new Date(now.getTime() - (days * 24 * 60 * 60 * 1000))
  
  const periodOrders = orders.value.filter(order => 
    new Date(order.createdAt) >= startDate
  )
  
  const periodCustomers = customers.value.filter(customer => 
    new Date(customer.createdAt) >= startDate
  )
  
  const totalRevenue = periodOrders.reduce((sum, order) => sum + (order.total || 0), 0)
  const totalOrders = periodOrders.length
  const newCustomers = periodCustomers.length
  const avgOrderValue = totalOrders > 0 ? totalRevenue / totalOrders : 0
  
  return {
    totalRevenue,
    totalOrders,
    newCustomers,
    avgOrderValue
  }
}

function getPreviousPeriodData() {
  const now = new Date()
  const days = getDaysForPeriod(selectedPeriod.value)
  const currentStartDate = new Date(now.getTime() - (days * 24 * 60 * 60 * 1000))
  const previousStartDate = new Date(currentStartDate.getTime() - (days * 24 * 60 * 60 * 1000))
  
  const periodOrders = orders.value.filter(order => {
    const orderDate = new Date(order.createdAt)
    return orderDate >= previousStartDate && orderDate < currentStartDate
  })
  
  const periodCustomers = customers.value.filter(customer => {
    const customerDate = new Date(customer.createdAt)
    return customerDate >= previousStartDate && customerDate < currentStartDate
  })
  
  const totalRevenue = periodOrders.reduce((sum, order) => sum + (order.total || 0), 0)
  const totalOrders = periodOrders.length
  const newCustomers = periodCustomers.length
  const avgOrderValue = totalOrders > 0 ? totalRevenue / totalOrders : 0
  
  return {
    totalRevenue,
    totalOrders,
    newCustomers,
    avgOrderValue
  }
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
    console.warn('No active shop selected')
    loading.value = false
    return
  }

  try {
    // Load data in parallel
    const [ordersData, productsData, customersData, discountsData] = await Promise.all([
      orderService.fetchAllByShop(shopId.value),
      productService.fetchAllByShop(shopId.value),
      customerService.fetchAll(shopId.value),
      discountService.fetchAllByShop(shopId.value)
    ])

    orders.value = ordersData
    products.value = productsData
    customers.value = customersData
    discounts.value = discountsData

  } catch (error) {
    console.error('Failed to load analytics data:', error)
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

onMounted(() => {
  if (shopId.value) {
    loadAnalyticsData()
  }
})
</script>

<style scoped>
/* Custom styles if needed */
</style>
