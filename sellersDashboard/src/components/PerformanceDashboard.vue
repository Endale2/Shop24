<template>
  <div class="performance-dashboard">
    <!-- Performance Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <ClockIcon class="h-8 w-8 text-blue-500" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Avg API Time</p>
            <p class="text-2xl font-semibold text-gray-900">
              {{ report.summary.averageApiDuration }}ms
            </p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <DatabaseIcon class="h-8 w-8 text-green-500" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Cache Hit Rate</p>
            <p class="text-2xl font-semibold text-gray-900">
              {{ report.summary.cacheHitRate }}%
            </p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <ExclamationTriangleIcon 
              :class="report.summary.errorRate > 10 ? 'text-red-500' : 'text-yellow-500'"
              class="h-8 w-8"
            />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Error Rate</p>
            <p class="text-2xl font-semibold text-gray-900">
              {{ report.summary.errorRate }}%
            </p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <CpuChipIcon class="h-8 w-8 text-purple-500" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">Memory Usage</p>
            <p class="text-2xl font-semibold text-gray-900">
              {{ report.summary.memoryUsage }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Alerts -->
    <div v-if="alerts.length > 0" class="mb-6">
      <h3 class="text-lg font-medium text-gray-900 mb-3">Performance Alerts</h3>
      <div class="space-y-2">
        <div
          v-for="alert in alerts"
          :key="alert.id"
          class="bg-yellow-50 border border-yellow-200 rounded-md p-4"
        >
          <div class="flex">
            <div class="flex-shrink-0">
              <ExclamationTriangleIcon class="h-5 w-5 text-yellow-400" />
            </div>
            <div class="ml-3 flex-1">
              <h4 class="text-sm font-medium text-yellow-800">
                {{ getAlertTitle(alert.type) }}
              </h4>
              <p class="text-sm text-yellow-700 mt-1">
                {{ getAlertMessage(alert) }}
              </p>
              <p class="text-xs text-yellow-600 mt-1">
                {{ formatTime(alert.timestamp) }}
              </p>
            </div>
            <div class="ml-3 flex-shrink-0">
              <button
                @click="dismissAlert(alert.id)"
                class="text-yellow-400 hover:text-yellow-600"
              >
                <XMarkIcon class="h-5 w-5" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Detailed Metrics -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- API Performance -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">API Performance</h3>
          <dl class="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div>
              <dt class="text-sm font-medium text-gray-500">Total Calls</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.apiCalls.total }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Success Rate</dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ getSuccessRate() }}%
              </dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Slow Calls</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.apiCalls.slowCalls }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Total Duration</dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ Math.round(metrics.apiCalls.totalDuration) }}ms
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <!-- Cache Performance -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Cache Performance</h3>
          <dl class="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div>
              <dt class="text-sm font-medium text-gray-500">Cache Hits</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.cache.hits }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Cache Misses</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.cache.misses }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Cache Size</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.cache.size }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Evictions</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.cache.evictions }}</dd>
            </div>
          </dl>
        </div>
      </div>

      <!-- State Management -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">State Management</h3>
          <dl class="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div>
              <dt class="text-sm font-medium text-gray-500">State Updates</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.state.updates }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Sync Events</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.state.syncEvents }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Conflicts</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.state.conflicts }}</dd>
            </div>
            <div>
              <dt class="text-sm font-medium text-gray-500">Recoveries</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.state.recoveries }}</dd>
            </div>
          </dl>
        </div>
      </div>

      <!-- Error Analysis -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Error Analysis</h3>
          <dl class="grid grid-cols-1 gap-4">
            <div>
              <dt class="text-sm font-medium text-gray-500">Total Errors</dt>
              <dd class="mt-1 text-sm text-gray-900">{{ metrics.errors.total }}</dd>
            </div>
            <div v-if="Object.keys(metrics.errors.byType).length > 0">
              <dt class="text-sm font-medium text-gray-500">By Type</dt>
              <dd class="mt-1 text-sm text-gray-900">
                <div v-for="(count, type) in metrics.errors.byType" :key="type" class="flex justify-between">
                  <span class="capitalize">{{ type }}:</span>
                  <span>{{ count }}</span>
                </div>
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="mt-6 bg-white rounded-lg shadow">
      <div class="px-4 py-5 sm:p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Recent Activity</h3>
        <div class="flow-root">
          <ul class="-mb-8">
            <li v-for="(entry, index) in recentHistory" :key="index" class="relative pb-8">
              <div v-if="index !== recentHistory.length - 1" class="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200"></div>
              <div class="relative flex space-x-3">
                <div>
                  <span :class="getActivityIconClass(entry.type)" class="h-8 w-8 rounded-full flex items-center justify-center ring-8 ring-white">
                    <component :is="getActivityIcon(entry.type)" class="h-4 w-4" />
                  </span>
                </div>
                <div class="min-w-0 flex-1 pt-1.5 flex justify-between space-x-4">
                  <div>
                    <p class="text-sm text-gray-500">
                      {{ getActivityDescription(entry) }}
                    </p>
                  </div>
                  <div class="text-right text-sm whitespace-nowrap text-gray-500">
                    {{ formatTime(entry.timestamp) }}
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Actions -->
    <div class="mt-6 flex justify-end space-x-3">
      <button
        @click="resetMetrics"
        class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Reset Metrics
      </button>
      <button
        @click="exportReport"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Export Report
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import {
  ClockIcon,
  DatabaseIcon,
  ExclamationTriangleIcon,
  CpuChipIcon,
  XMarkIcon,
  CloudIcon,
  CogIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/24/outline'
import { usePerformanceMonitor } from '@/utils/performanceMonitor'

// Use global performance monitor
const { metrics, alerts, report, dismissAlert, reset } = usePerformanceMonitor({ global: true })

// Computed properties
const recentHistory = computed(() => report.value.history.slice(-10).reverse())

// Methods
function getSuccessRate() {
  const total = metrics.value.apiCalls.total
  if (total === 0) return 100
  return Math.round((metrics.value.apiCalls.successful / total) * 100)
}

function getAlertTitle(type) {
  const titles = {
    'slow-api-call': 'Slow API Call',
    'high-error-rate': 'High Error Rate',
    'low-cache-hit-rate': 'Low Cache Hit Rate',
    'high-memory-usage': 'High Memory Usage'
  }
  return titles[type] || 'Performance Alert'
}

function getAlertMessage(alert) {
  const { type, data } = alert
  switch (type) {
    case 'slow-api-call':
      return `${data.operation} took ${Math.round(data.duration)}ms (threshold: ${data.threshold}ms)`
    case 'high-error-rate':
      return `Error rate is ${Math.round(data.errorRate * 100)}% (threshold: ${Math.round(data.threshold * 100)}%)`
    case 'low-cache-hit-rate':
      return `Cache hit rate is ${Math.round(data.hitRate * 100)}% (threshold: ${Math.round(data.threshold * 100)}%)`
    case 'high-memory-usage':
      return `Memory usage is ${Math.round(data.used / 1024 / 1024)}MB (threshold: ${Math.round(data.threshold / 1024 / 1024)}MB)`
    default:
      return 'Performance threshold exceeded'
  }
}

function getActivityIcon(type) {
  const icons = {
    'api-call': CloudIcon,
    'cache-event': DatabaseIcon,
    'state-update': CogIcon,
    'sync-event': CogIcon,
    'error': ExclamationCircleIcon,
    'ux-metric': ClockIcon
  }
  return icons[type] || CogIcon
}

function getActivityIconClass(type) {
  const classes = {
    'api-call': 'bg-blue-500',
    'cache-event': 'bg-green-500',
    'state-update': 'bg-purple-500',
    'sync-event': 'bg-indigo-500',
    'error': 'bg-red-500',
    'ux-metric': 'bg-yellow-500'
  }
  return classes[type] || 'bg-gray-500'
}

function getActivityDescription(entry) {
  switch (entry.type) {
    case 'api-call':
      return `API call ${entry.operation} ${entry.success ? 'succeeded' : 'failed'} in ${Math.round(entry.duration)}ms`
    case 'cache-event':
      return `Cache ${entry.event} for ${entry.key}`
    case 'state-update':
      return `State updated: ${entry.operation}`
    case 'sync-event':
      return `Sync event: ${entry.event}`
    case 'error':
      return `Error in ${entry.operation}: ${entry.message}`
    case 'ux-metric':
      return `UX metric ${entry.metric}: ${entry.value}ms`
    default:
      return 'Unknown activity'
  }
}

function formatTime(timestamp) {
  const now = Date.now()
  const diff = now - timestamp
  
  if (diff < 60000) return 'Just now'
  if (diff < 3600000) return `${Math.floor(diff / 60000)}m ago`
  if (diff < 86400000) return `${Math.floor(diff / 3600000)}h ago`
  return new Date(timestamp).toLocaleDateString()
}

function resetMetrics() {
  if (confirm('Are you sure you want to reset all performance metrics?')) {
    reset()
  }
}

function exportReport() {
  const reportData = JSON.stringify(report.value, null, 2)
  const blob = new Blob([reportData], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `performance-report-${new Date().toISOString().split('T')[0]}.json`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}
</script>

<style scoped>
.performance-dashboard {
  max-width: 7xl;
  margin: 0 auto;
  padding: 1rem;
}
</style>
