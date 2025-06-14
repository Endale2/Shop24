<template>
  <div class="pt-16 pb-8 bg-gray-50 min-h-[calc(100vh-4rem)]">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
      <!-- Header & Period Selector -->
      <div class="flex flex-col md:flex-row items-start md:items-center justify-between space-y-4 md:space-y-0">
        <h1 class="text-2xl md:text-3xl font-extrabold text-gray-900">
          Analytics Dashboard
        </h1>
        <div class="inline-flex bg-white p-1 rounded-lg shadow-sm">
          <button
            v-for="period in periods"
            :key="period.id"
            @click="selectedPeriod = period.id"
            :class="[
              'px-4 py-1 text-sm font-medium rounded transition',
              selectedPeriod === period.id
                ? 'bg-green-600 text-white'
                : 'text-gray-600 hover:bg-gray-100'
            ]"
          >
            {{ period.label }}
          </button>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <template v-if="loading">
          <div v-for="n in 4" :key="n" class="h-32 bg-white rounded-lg shadow animate-pulse"></div>
        </template>
        <!-- only render real stats when analyticsData is ready -->
        <template v-else-if="analyticsData">
          <div
            v-for="stat in analyticsData.stats"
            :key="stat.title"
            class="bg-white p-4 rounded-lg shadow hover:shadow-md transition flex items-center space-x-4 border-l-4"
            :class="stat.borderColor"
          >
            <div class="p-2 rounded-full" :class="stat.iconBg">
              <component :is="stat.icon" class="h-6 w-6" :class="stat.iconColor" />
            </div>
            <div>
              <p class="text-sm text-gray-500">{{ stat.title }}</p>
              <p class="text-xl font-semibold text-gray-900">{{ stat.value }}</p>
              <p
                class="text-xs"
                :class="stat.change.startsWith('+') ? 'text-green-600' : 'text-red-600'"
              >
                {{ stat.change }} vs last period
              </p>
            </div>
          </div>
        </template>
      </div>

      <!-- Charts -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Revenue Over Time -->
        <div class="lg:col-span-2 bg-white p-6 rounded-lg shadow">
          <h2 class="text-lg font-semibold text-gray-800 mb-4">Revenue Over Time</h2>
          <div class="h-64">
            <Line
              v-if="analyticsData"
              :data="analyticsData.revenueChart.data"
              :options="chartOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg"></div>
          </div>
        </div>

        <!-- Sales by Category -->
        <div class="bg-white p-6 rounded-lg shadow">
          <h2 class="text-lg font-semibold text-gray-800 mb-4">Sales by Category</h2>
          <div class="h-64">
            <Doughnut
              v-if="analyticsData"
              :data="analyticsData.categoryChart.data"
              :options="doughnutOptions"
            />
            <div v-else class="h-full bg-gray-100 animate-pulse rounded-lg"></div>
          </div>
        </div>
      </div>

      <!-- Top Products & Recent Orders -->
      <div class="grid grid-cols-1 lg:grid-cols-5 gap-6" v-if="analyticsData">
        <!-- Top Selling Products -->
        <div class="lg:col-span-2 bg-white p-6 rounded-lg shadow space-y-4">
          <h2 class="text-lg font-semibold text-gray-800">Top Selling Products</h2>
          <ul class="space-y-3">
            <li
              v-for="prod in analyticsData.topProducts"
              :key="prod.name"
              class="flex items-center space-x-3"
            >
              <div class="h-12 w-12 bg-gray-100 rounded-lg overflow-hidden">
                <img :src="prod.image" alt="" class="h-full w-full object-cover" />
              </div>
              <div class="flex-1">
                <p class="font-medium text-gray-900">{{ prod.name }}</p>
                <p class="text-sm text-gray-500">{{ prod.category }}</p>
              </div>
              <p class="font-semibold text-gray-900">{{ prod.unitsSold.toLocaleString() }} units</p>
            </li>
          </ul>
        </div>

        <!-- Recent Orders Table -->
        <div class="lg:col-span-3 bg-white p-6 rounded-lg shadow overflow-x-auto">
          <h2 class="text-lg font-semibold text-gray-800 mb-4">Recent Orders</h2>
          <table class="w-full text-sm">
            <thead class="bg-gray-100">
              <tr>
                <th class="p-3 text-left font-medium text-gray-600">Order ID</th>
                <th class="p-3 text-left font-medium text-gray-600">Customer</th>
                <th class="p-3 text-left font-medium text-gray-600">Amount</th>
                <th class="p-3 text-left font-medium text-gray-600">Status</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr
                v-for="order in analyticsData.recentOrders"
                :key="order.id"
                class="hover:bg-gray-50"
              >
                <td class="p-3 font-mono text-gray-500">{{ order.id }}</td>
                <td class="p-3 text-gray-800">{{ order.customerName }}</td>
                <td class="p-3 text-gray-800">{{ order.amount }}</td>
                <td class="p-3">
                  <span
                    class="px-2 py-1 text-xs font-semibold rounded-full"
                    :class="order.status.class"
                  >
                    {{ order.status.label }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, defineAsyncComponent } from 'vue'
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

const CurrencyDollarIcon = defineAsyncComponent(() => import('@heroicons/vue/outline/CurrencyDollarIcon'))
const ShoppingCartIcon   = defineAsyncComponent(() => import('@heroicons/vue/outline/ShoppingCartIcon'))
const UserGroupIcon      = defineAsyncComponent(() => import('@heroicons/vue/outline/UserGroupIcon'))
const ChartBarIcon       = defineAsyncComponent(() => import('@heroicons/vue/outline/ChartBarIcon'))

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, ArcElement)

const loading = ref(true)
const analyticsData = ref(null)
const selectedPeriod = ref('last_30_days')
const periods = [
  { id: 'last_24_hours', label: '24h' },
  { id: 'last_7_days',   label: '7d' },
  { id: 'last_30_days',  label: '30d' },
  { id: 'last_12_months',label: '12m' },
]

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { legend: { display: false } },
  scales: { y: { ticks: { callback: v => `$${v/1000}k` } } }
}
const doughnutOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { legend: { position: 'bottom' } }
}

const random = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min
const fmt$ = n => `$${n.toLocaleString()}`
const fmtPct = () => `${Math.random() > 0.5 ? '+' : '-'}${random(1,9)}.${random(0,9)}%`
const names = ['Liam','Olivia','Noah','Emma','Oliver','Ava']

function genData() {
  return {
    stats: [
      { title:'Total Revenue',   value:fmt$(random(50000,150000)), change:fmtPct(), icon:CurrencyDollarIcon, iconBg:'bg-green-100', iconColor:'text-green-600', borderColor:'border-green-500' },
      { title:'Total Orders',    value:random(1000,5000).toLocaleString(), change:fmtPct(), icon:ShoppingCartIcon, iconBg:'bg-blue-100', iconColor:'text-blue-600', borderColor:'border-blue-500' },
      { title:'New Customers',   value:random(200,800).toLocaleString(),   change:fmtPct(), icon:UserGroupIcon, iconBg:'bg-indigo-100', iconColor:'text-indigo-600', borderColor:'border-indigo-500' },
      { title:'Avg Order Value', value:fmt$(random(80,200)),              change:fmtPct(), icon:ChartBarIcon, iconBg:'bg-amber-100', iconColor:'text-amber-600', borderColor:'border-amber-500' },
    ],
    revenueChart: {
      data: {
        labels:['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug'],
        datasets:[{
          label:'Revenue',
          data:Array.from({length:8},()=>random(10000,80000)),
          borderColor:'rgb(16,185,129)',
          backgroundColor:'rgba(16,185,129,0.1)',
          tension:0.3, fill:true
        }]
      }
    },
    categoryChart: {
      data: {
        labels:['Electronics','Apparel','Home Goods','Books'],
        datasets:[{
          data:Array.from({length:4},()=>random(20,100)),
          backgroundColor:['#10B981','#3B82F6','#8B5CF6','#F59E0B']
        }]
      }
    },
    topProducts: Array.from({length:4},(_,i)=>({
      name:['Wireless Headphones','Organic T-Shirt','Smart Hub','Novel'][i],
      category:['Electronics','Apparel','Home Goods','Books'][i],
      unitsSold:random(500,2000),
      image:`https://picsum.photos/seed/p${i}/100/100`
    })),
    recentOrders: Array.from({length:5},()=>({
      id:'#'+random(10000,99999),
      customerName:`${names[random(0,5)]} ${names[random(0,5)]}`,
      amount:fmt$(random(20,500)),
      status:[
        {label:'Completed',class:'bg-green-100 text-green-700'},
        {label:'Processing',class:'bg-blue-100 text-blue-700'},
        {label:'Shipped',class:'bg-amber-100 text-amber-700'}
      ][random(0,2)]
    }))
  }
}

onMounted(() => {
  setTimeout(() => {
    analyticsData.value = genData()
    loading.value = false
  }, 1200)
})
</script>

<style scoped>
/* purely Tailwind; no extra CSS needed */
</style>
