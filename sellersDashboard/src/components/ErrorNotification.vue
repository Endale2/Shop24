<template>
  <div class="error-notification-container">
    <!-- Error Notifications -->
    <TransitionGroup
      name="notification"
      tag="div"
      class="fixed top-4 right-4 z-50 space-y-2"
    >
      <div
        v-for="notification in notifications"
        :key="notification.id"
        :class="getNotificationClasses(notification)"
        class="max-w-md w-full bg-white shadow-lg rounded-lg pointer-events-auto ring-1 ring-black ring-opacity-5 overflow-hidden"
      >
        <div class="p-4">
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <component
                :is="getNotificationIcon(notification.type)"
                :class="getIconClasses(notification.severity)"
                class="h-6 w-6"
              />
            </div>
            <div class="ml-3 w-0 flex-1 pt-0.5">
              <p class="text-sm font-medium text-gray-900">
                {{ notification.title }}
              </p>
              <p class="mt-1 text-sm text-gray-500">
                {{ notification.message }}
              </p>
              
              <!-- Error Actions -->
              <div v-if="notification.actions?.length" class="mt-3 flex space-x-2">
                <button
                  v-for="action in notification.actions"
                  :key="action.action"
                  @click="handleAction(notification, action)"
                  :class="getActionClasses(action)"
                  class="text-xs font-medium rounded-md px-2 py-1 focus:outline-none focus:ring-2 focus:ring-offset-2"
                >
                  {{ action.label }}
                </button>
              </div>
            </div>
            <div class="ml-4 flex-shrink-0 flex">
              <button
                @click="dismissNotification(notification.id)"
                class="bg-white rounded-md inline-flex text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                <span class="sr-only">Close</span>
                <XIcon class="h-5 w-5" />
              </button>
            </div>
          </div>
        </div>
        
        <!-- Progress bar for auto-dismiss -->
        <div
          v-if="notification.autoDismiss"
          class="h-1 bg-gray-200"
        >
          <div
            class="h-full bg-red-500 transition-all duration-1000 ease-linear"
            :style="{ width: getProgressWidth(notification) + '%' }"
          ></div>
        </div>
      </div>
    </TransitionGroup>

    <!-- Global Error Overlay for Critical Errors -->
    <div
      v-if="hasCriticalError"
      class="fixed inset-0 z-50 overflow-y-auto"
      aria-labelledby="modal-title"
      role="dialog"
      aria-modal="true"
    >
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
        
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
        
        <div class="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full sm:p-6">
          <div class="sm:flex sm:items-start">
            <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
              <ExclamationIcon class="h-6 w-6 text-red-600" />
            </div>
            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
              <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                Critical Error
              </h3>
              <div class="mt-2">
                <p class="text-sm text-gray-500">
                  {{ criticalError?.userMessage || 'A critical error has occurred. Please refresh the page or contact support.' }}
                </p>
              </div>
            </div>
          </div>
          <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
            <button
              @click="handleCriticalErrorAction('refresh')"
              type="button"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Refresh Page
            </button>
            <button
              @click="handleCriticalErrorAction('dismiss')"
              type="button"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:w-auto sm:text-sm"
            >
              Dismiss
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { 
  XIcon, 
  ExclamationIcon, 
  ExclamationCircleIcon,
  InformationCircleIcon,
  CheckCircleIcon 
} from '@heroicons/vue/outline'
import { globalErrorHandler, ERROR_SEVERITY } from '@/utils/errorHandler'

const router = useRouter()

// Reactive state from global error handler with safety check
const notifications = computed(() => {
  try {
    return globalErrorHandler?.globalErrorState?.notifications || []
  } catch (error) {
    console.warn('[ErrorNotification] Failed to access notifications:', error)
    return []
  }
})
const criticalError = computed(() => {
  try {
    return notifications.value?.find(n => n.severity === ERROR_SEVERITY.CRITICAL) || null
  } catch (error) {
    console.warn('[ErrorNotification] Failed to find critical error:', error)
    return null
  }
})
const hasCriticalError = computed(() => !!criticalError.value)

// Notification styling
function getNotificationClasses(notification) {
  const baseClasses = 'transform transition-all duration-300 ease-in-out'
  const severityClasses = {
    [ERROR_SEVERITY.LOW]: 'border-l-4 border-yellow-400',
    [ERROR_SEVERITY.MEDIUM]: 'border-l-4 border-orange-400',
    [ERROR_SEVERITY.HIGH]: 'border-l-4 border-red-400',
    [ERROR_SEVERITY.CRITICAL]: 'border-l-4 border-red-600'
  }
  
  return `${baseClasses} ${severityClasses[notification.severity] || ''}`
}

function getNotificationIcon(type) {
  const icons = {
    error: ExclamationCircleIcon,
    warning: ExclamationIcon,
    info: InformationCircleIcon,
    success: CheckCircleIcon
  }
  return icons[type] || ExclamationCircleIcon
}

function getIconClasses(severity) {
  const classes = {
    [ERROR_SEVERITY.LOW]: 'text-yellow-400',
    [ERROR_SEVERITY.MEDIUM]: 'text-orange-400',
    [ERROR_SEVERITY.HIGH]: 'text-red-400',
    [ERROR_SEVERITY.CRITICAL]: 'text-red-600'
  }
  return classes[severity] || 'text-red-400'
}

function getActionClasses(action) {
  if (action.primary) {
    return 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500'
  }
  return 'bg-gray-100 text-gray-700 hover:bg-gray-200 focus:ring-gray-500'
}

function getProgressWidth(notification) {
  if (!notification.autoDismiss) return 0
  const elapsed = Date.now() - notification.timestamp
  const duration = notification.autoDismissDuration || 5000
  return Math.max(0, 100 - (elapsed / duration) * 100)
}

// Action handlers
function handleAction(notification, action) {
  switch (action.action) {
    case 'retry':
      handleRetry(notification)
      break
    case 'login':
      router.push({ name: 'Landing' })
      break
    case 'refresh':
      window.location.reload()
      break
    case 'dismiss':
      dismissNotification(notification.id)
      break
    default:
      console.warn('Unknown action:', action.action)
  }
}

function handleRetry(notification) {
  // Emit retry event or call retry function
  console.log('Retrying operation for notification:', notification.id)
  dismissNotification(notification.id)
  
  // You can emit events here for parent components to handle
  // this.$emit('retry', notification)
}

function handleCriticalErrorAction(action) {
  if (action === 'refresh') {
    window.location.reload()
  } else if (action === 'dismiss') {
    if (criticalError.value) {
      dismissNotification(criticalError.value.id)
    }
  }
}

function dismissNotification(notificationId) {
  try {
    if (globalErrorHandler && typeof globalErrorHandler.dismissNotification === 'function') {
      globalErrorHandler.dismissNotification(notificationId)
    } else {
      console.warn('[ErrorNotification] globalErrorHandler not available for dismissing notification')
    }
  } catch (error) {
    console.error('[ErrorNotification] Failed to dismiss notification:', error)
  }
}

// Auto-dismiss logic
let dismissTimers = new Map()

function setupAutoDismiss(notification) {
  if (notification.severity === ERROR_SEVERITY.LOW) {
    const timer = setTimeout(() => {
      dismissNotification(notification.id)
      dismissTimers.delete(notification.id)
    }, 5000)
    
    dismissTimers.set(notification.id, timer)
  }
}

// Watch for new notifications
let lastNotificationCount = 0
function checkForNewNotifications() {
  if (notifications.value.length > lastNotificationCount) {
    const newNotifications = notifications.value.slice(lastNotificationCount)
    newNotifications.forEach(setupAutoDismiss)
  }
  lastNotificationCount = notifications.value.length
}

// Lifecycle
onMounted(() => {
  // Set up periodic check for new notifications
  const interval = setInterval(checkForNewNotifications, 100)
  
  onUnmounted(() => {
    clearInterval(interval)
    // Clear all timers
    dismissTimers.forEach(timer => clearTimeout(timer))
    dismissTimers.clear()
  })
})
</script>

<style scoped>
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.notification-move {
  transition: transform 0.3s ease;
}
</style>
