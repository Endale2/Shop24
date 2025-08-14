// utils/errorHandler.js
import { ref, reactive } from 'vue'

// Error types and categories
export const ERROR_TYPES = {
  NETWORK: 'network',
  VALIDATION: 'validation',
  PERMISSION: 'permission',
  NOT_FOUND: 'not_found',
  SERVER: 'server',
  TIMEOUT: 'timeout',
  RATE_LIMIT: 'rate_limit',
  SHOP_STATE: 'shop_state',
  AUTH: 'auth'
}

export const ERROR_SEVERITY = {
  LOW: 'low',
  MEDIUM: 'medium',
  HIGH: 'high',
  CRITICAL: 'critical'
}

// Global error state
const globalErrorState = reactive({
  errors: new Map(),
  notifications: [],
  retryQueue: new Map(),
  errorCount: 0,
  lastError: null
})

// Error classification rules
const ERROR_CLASSIFICATION = {
  [ERROR_TYPES.NETWORK]: {
    severity: ERROR_SEVERITY.MEDIUM,
    retryable: true,
    userMessage: 'Network connection issue. Please check your internet connection.',
    autoRetry: true
  },
  [ERROR_TYPES.VALIDATION]: {
    severity: ERROR_SEVERITY.LOW,
    retryable: false,
    userMessage: 'Please check your input and try again.',
    autoRetry: false
  },
  [ERROR_TYPES.PERMISSION]: {
    severity: ERROR_SEVERITY.HIGH,
    retryable: false,
    userMessage: 'You don\'t have permission to perform this action.',
    autoRetry: false
  },
  [ERROR_TYPES.NOT_FOUND]: {
    severity: ERROR_SEVERITY.MEDIUM,
    retryable: false,
    userMessage: 'The requested resource was not found.',
    autoRetry: false
  },
  [ERROR_TYPES.SERVER]: {
    severity: ERROR_SEVERITY.HIGH,
    retryable: true,
    userMessage: 'Server error occurred. Please try again later.',
    autoRetry: true
  },
  [ERROR_TYPES.TIMEOUT]: {
    severity: ERROR_SEVERITY.MEDIUM,
    retryable: true,
    userMessage: 'Request timed out. Please try again.',
    autoRetry: true
  },
  [ERROR_TYPES.RATE_LIMIT]: {
    severity: ERROR_SEVERITY.MEDIUM,
    retryable: true,
    userMessage: 'Too many requests. Please wait a moment and try again.',
    autoRetry: true
  },
  [ERROR_TYPES.SHOP_STATE]: {
    severity: ERROR_SEVERITY.HIGH,
    retryable: true,
    userMessage: 'Shop state error. Attempting to recover...',
    autoRetry: true
  },
  [ERROR_TYPES.AUTH]: {
    severity: ERROR_SEVERITY.CRITICAL,
    retryable: false,
    userMessage: 'Authentication error. Please log in again.',
    autoRetry: false
  }
}

/**
 * Enhanced error handler class
 */
export class ErrorHandler {
  constructor(context = 'global') {
    this.context = context
    this.errorHistory = []
    this.retryAttempts = new Map()
    this.maxRetries = 3
    this.retryDelay = 1000
  }

  /**
   * Classify error type based on error object
   */
  classifyError(error) {
    if (!error) return ERROR_TYPES.SERVER

    // Network errors
    if (!error.response && (error.code === 'NETWORK_ERROR' || !navigator.onLine)) {
      return ERROR_TYPES.NETWORK
    }

    // HTTP status-based classification
    const status = error.response?.status
    if (status === 401 || status === 403) return ERROR_TYPES.PERMISSION
    if (status === 404) return ERROR_TYPES.NOT_FOUND
    if (status === 422) return ERROR_TYPES.VALIDATION
    if (status === 429) return ERROR_TYPES.RATE_LIMIT
    if (status >= 400 && status < 500) return ERROR_TYPES.VALIDATION
    if (status >= 500) return ERROR_TYPES.SERVER

    // Timeout errors
    if (error.code === 'ECONNABORTED' || error.message?.includes('timeout')) {
      return ERROR_TYPES.TIMEOUT
    }

    // Shop state specific errors
    if (error.message?.includes('shop') || error.context === 'shop') {
      return ERROR_TYPES.SHOP_STATE
    }

    // Auth specific errors
    if (error.message?.includes('auth') || error.context === 'auth') {
      return ERROR_TYPES.AUTH
    }

    return ERROR_TYPES.SERVER
  }

  /**
   * Handle error with classification and recovery
   */
  async handleError(error, options = {}) {
    const {
      context = this.context,
      operation = 'unknown',
      silent = false,
      customMessage = null,
      onRetry = null,
      maxRetries = this.maxRetries
    } = options

    const errorType = this.classifyError(error)
    const classification = ERROR_CLASSIFICATION[errorType]
    const errorId = `${context}_${operation}_${Date.now()}`

    const errorInfo = {
      id: errorId,
      type: errorType,
      severity: classification.severity,
      message: error.message || 'Unknown error',
      userMessage: customMessage || classification.userMessage,
      context,
      operation,
      timestamp: Date.now(),
      retryable: classification.retryable,
      autoRetry: classification.autoRetry,
      originalError: error
    }

    // Store error
    this.errorHistory.push(errorInfo)
    globalErrorState.errors.set(errorId, errorInfo)
    globalErrorState.errorCount += 1
    globalErrorState.lastError = errorInfo

    console.error(`[ErrorHandler:${context}] ${operation} failed:`, {
      type: errorType,
      severity: classification.severity,
      message: error.message,
      error
    })

    // Show user notification if not silent
    if (!silent && classification.severity !== ERROR_SEVERITY.LOW) {
      this.showErrorNotification(errorInfo)
    }

    // Attempt auto-retry if applicable
    if (classification.autoRetry && classification.retryable) {
      const retryKey = `${context}_${operation}`
      const currentRetries = this.retryAttempts.get(retryKey) || 0

      if (currentRetries < maxRetries) {
        this.retryAttempts.set(retryKey, currentRetries + 1)
        
        console.log(`[ErrorHandler] Auto-retry ${currentRetries + 1}/${maxRetries} for ${operation}`)
        
        const delay = this.retryDelay * Math.pow(2, currentRetries)
        
        setTimeout(async () => {
          try {
            if (onRetry) {
              await onRetry()
              this.retryAttempts.delete(retryKey)
              this.clearError(errorId)
            }
          } catch (retryError) {
            console.error(`[ErrorHandler] Retry failed for ${operation}:`, retryError)
            // Don't auto-retry the retry
          }
        }, delay)
      } else {
        console.error(`[ErrorHandler] Max retries exceeded for ${operation}`)
        this.retryAttempts.delete(retryKey)
      }
    }

    return errorInfo
  }

  /**
   * Show error notification to user
   */
  showErrorNotification(errorInfo) {
    const notification = {
      id: errorInfo.id,
      type: 'error',
      title: this.getErrorTitle(errorInfo.type),
      message: errorInfo.userMessage,
      severity: errorInfo.severity,
      timestamp: errorInfo.timestamp,
      actions: this.getErrorActions(errorInfo)
    }

    globalErrorState.notifications.push(notification)

    // Auto-dismiss low severity errors
    if (errorInfo.severity === ERROR_SEVERITY.LOW) {
      setTimeout(() => {
        this.dismissNotification(notification.id)
      }, 5000)
    }
  }

  /**
   * Get user-friendly error title
   */
  getErrorTitle(errorType) {
    const titles = {
      [ERROR_TYPES.NETWORK]: 'Connection Error',
      [ERROR_TYPES.VALIDATION]: 'Validation Error',
      [ERROR_TYPES.PERMISSION]: 'Permission Denied',
      [ERROR_TYPES.NOT_FOUND]: 'Not Found',
      [ERROR_TYPES.SERVER]: 'Server Error',
      [ERROR_TYPES.TIMEOUT]: 'Request Timeout',
      [ERROR_TYPES.RATE_LIMIT]: 'Rate Limited',
      [ERROR_TYPES.SHOP_STATE]: 'Shop State Error',
      [ERROR_TYPES.AUTH]: 'Authentication Error'
    }
    return titles[errorType] || 'Error'
  }

  /**
   * Get available actions for error
   */
  getErrorActions(errorInfo) {
    const actions = []

    if (errorInfo.retryable) {
      actions.push({
        label: 'Retry',
        action: 'retry',
        primary: true
      })
    }

    if (errorInfo.type === ERROR_TYPES.AUTH) {
      actions.push({
        label: 'Login',
        action: 'login',
        primary: true
      })
    }

    if (errorInfo.severity === ERROR_SEVERITY.CRITICAL) {
      actions.push({
        label: 'Refresh Page',
        action: 'refresh',
        primary: false
      })
    }

    actions.push({
      label: 'Dismiss',
      action: 'dismiss',
      primary: false
    })

    return actions
  }

  /**
   * Clear specific error
   */
  clearError(errorId) {
    globalErrorState.errors.delete(errorId)
    this.dismissNotification(errorId)
  }

  /**
   * Clear all errors for context
   */
  clearContextErrors() {
    for (const [id, error] of globalErrorState.errors) {
      if (error.context === this.context) {
        this.clearError(id)
      }
    }
  }

  /**
   * Dismiss notification
   */
  dismissNotification(notificationId) {
    const index = globalErrorState.notifications.findIndex(n => n.id === notificationId)
    if (index !== -1) {
      globalErrorState.notifications.splice(index, 1)
    }
  }

  /**
   * Get error statistics
   */
  getErrorStats() {
    const stats = {
      total: this.errorHistory.length,
      byType: {},
      bySeverity: {},
      recent: this.errorHistory.slice(-10)
    }

    this.errorHistory.forEach(error => {
      stats.byType[error.type] = (stats.byType[error.type] || 0) + 1
      stats.bySeverity[error.severity] = (stats.bySeverity[error.severity] || 0) + 1
    })

    return stats
  }
}

// Global error handler instance
export const globalErrorHandler = new ErrorHandler('global')

// Composable for error handling
export function useErrorHandler(context = 'component') {
  const errorHandler = new ErrorHandler(context)
  const errors = ref([])
  const hasErrors = ref(false)

  const handleError = async (error, options = {}) => {
    const errorInfo = await errorHandler.handleError(error, options)
    errors.value.push(errorInfo)
    hasErrors.value = errors.value.length > 0
    return errorInfo
  }

  const clearErrors = () => {
    errorHandler.clearContextErrors()
    errors.value = []
    hasErrors.value = false
  }

  const clearError = (errorId) => {
    errorHandler.clearError(errorId)
    errors.value = errors.value.filter(e => e.id !== errorId)
    hasErrors.value = errors.value.length > 0
  }

  return {
    errors,
    hasErrors,
    handleError,
    clearErrors,
    clearError,
    errorHandler,
    globalErrorState
  }
}
