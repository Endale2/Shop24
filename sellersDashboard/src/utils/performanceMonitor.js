// utils/performanceMonitor.js
import { reactive, ref, computed } from 'vue'

/**
 * Performance monitoring utility for shop state management
 */
export class PerformanceMonitor {
  constructor(options = {}) {
    this.options = {
      enableMetrics: true,
      enableLogging: true,
      sampleRate: 1.0, // 100% sampling by default
      maxHistorySize: 1000,
      alertThresholds: {
        slowApiCall: 5000,      // 5 seconds
        highErrorRate: 0.1,     // 10%
        lowCacheHitRate: 0.5,   // 50%
        memoryUsage: 50 * 1024 * 1024 // 50MB
      },
      ...options
    }
    
    this.metrics = reactive({
      // API Performance
      apiCalls: {
        total: 0,
        successful: 0,
        failed: 0,
        totalDuration: 0,
        averageDuration: 0,
        slowCalls: 0
      },
      
      // Cache Performance
      cache: {
        hits: 0,
        misses: 0,
        hitRate: 0,
        size: 0,
        evictions: 0
      },
      
      // State Management
      state: {
        updates: 0,
        syncEvents: 0,
        conflicts: 0,
        recoveries: 0
      },
      
      // Error Tracking
      errors: {
        total: 0,
        byType: {},
        byOperation: {},
        rate: 0
      },
      
      // User Experience
      ux: {
        loadingTime: 0,
        interactionDelay: 0,
        navigationTime: 0,
        errorRecoveryTime: 0
      },
      
      // Memory Usage
      memory: {
        used: 0,
        peak: 0,
        storeSize: 0
      }
    })
    
    this.history = []
    this.activeOperations = new Map()
    this.alerts = ref([])
    
    this.init()
  }
  
  /**
   * Initialize performance monitoring
   */
  init() {
    if (this.options.enableMetrics) {
      // Start periodic memory monitoring
      this.startMemoryMonitoring()
      
      // Start performance observer if available
      this.startPerformanceObserver()
    }
  }
  
  /**
   * Start API call tracking
   */
  startApiCall(operation, metadata = {}) {
    if (!this.shouldSample()) return null
    
    const callId = `${operation}_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
    const startTime = performance.now()
    
    const callInfo = {
      id: callId,
      operation,
      startTime,
      metadata,
      timestamp: Date.now()
    }
    
    this.activeOperations.set(callId, callInfo)
    
    if (this.options.enableLogging) {
      console.log(`[PerformanceMonitor] API call started: ${operation}`, metadata)
    }
    
    return callId
  }
  
  /**
   * End API call tracking
   */
  endApiCall(callId, success = true, error = null) {
    if (!callId || !this.activeOperations.has(callId)) return
    
    const callInfo = this.activeOperations.get(callId)
    const endTime = performance.now()
    const duration = endTime - callInfo.startTime
    
    // Update metrics
    this.metrics.apiCalls.total += 1
    this.metrics.apiCalls.totalDuration += duration
    this.metrics.apiCalls.averageDuration = 
      this.metrics.apiCalls.totalDuration / this.metrics.apiCalls.total
    
    if (success) {
      this.metrics.apiCalls.successful += 1
    } else {
      this.metrics.apiCalls.failed += 1
      this.recordError(error, callInfo.operation)
    }
    
    if (duration > this.options.alertThresholds.slowApiCall) {
      this.metrics.apiCalls.slowCalls += 1
      this.createAlert('slow-api-call', {
        operation: callInfo.operation,
        duration,
        threshold: this.options.alertThresholds.slowApiCall
      })
    }
    
    // Add to history
    this.addToHistory({
      type: 'api-call',
      operation: callInfo.operation,
      duration,
      success,
      error: error?.message,
      timestamp: Date.now()
    })
    
    this.activeOperations.delete(callId)
    
    if (this.options.enableLogging) {
      console.log(`[PerformanceMonitor] API call ended: ${callInfo.operation}`, {
        duration: Math.round(duration),
        success
      })
    }
  }
  
  /**
   * Record cache hit/miss
   */
  recordCacheEvent(type, key, metadata = {}) {
    if (!this.shouldSample()) return
    
    if (type === 'hit') {
      this.metrics.cache.hits += 1
    } else if (type === 'miss') {
      this.metrics.cache.misses += 1
    }
    
    const total = this.metrics.cache.hits + this.metrics.cache.misses
    this.metrics.cache.hitRate = total > 0 ? this.metrics.cache.hits / total : 0
    
    // Check for low cache hit rate
    if (total > 10 && this.metrics.cache.hitRate < this.options.alertThresholds.lowCacheHitRate) {
      this.createAlert('low-cache-hit-rate', {
        hitRate: this.metrics.cache.hitRate,
        threshold: this.options.alertThresholds.lowCacheHitRate
      })
    }
    
    this.addToHistory({
      type: 'cache-event',
      event: type,
      key,
      hitRate: this.metrics.cache.hitRate,
      timestamp: Date.now(),
      ...metadata
    })
  }
  
  /**
   * Record state update
   */
  recordStateUpdate(operation, metadata = {}) {
    if (!this.shouldSample()) return
    
    this.metrics.state.updates += 1
    
    this.addToHistory({
      type: 'state-update',
      operation,
      timestamp: Date.now(),
      ...metadata
    })
  }
  
  /**
   * Record synchronization event
   */
  recordSyncEvent(type, metadata = {}) {
    if (!this.shouldSample()) return
    
    this.metrics.state.syncEvents += 1
    
    if (type === 'conflict') {
      this.metrics.state.conflicts += 1
    } else if (type === 'recovery') {
      this.metrics.state.recoveries += 1
    }
    
    this.addToHistory({
      type: 'sync-event',
      event: type,
      timestamp: Date.now(),
      ...metadata
    })
  }
  
  /**
   * Record error
   */
  recordError(error, operation = 'unknown') {
    if (!this.shouldSample()) return
    
    this.metrics.errors.total += 1
    
    const errorType = this.classifyError(error)
    this.metrics.errors.byType[errorType] = (this.metrics.errors.byType[errorType] || 0) + 1
    this.metrics.errors.byOperation[operation] = (this.metrics.errors.byOperation[operation] || 0) + 1
    
    // Calculate error rate
    const totalOperations = this.metrics.apiCalls.total + this.metrics.state.updates
    this.metrics.errors.rate = totalOperations > 0 ? this.metrics.errors.total / totalOperations : 0
    
    // Check for high error rate
    if (totalOperations > 10 && this.metrics.errors.rate > this.options.alertThresholds.highErrorRate) {
      this.createAlert('high-error-rate', {
        errorRate: this.metrics.errors.rate,
        threshold: this.options.alertThresholds.highErrorRate
      })
    }
    
    this.addToHistory({
      type: 'error',
      errorType,
      operation,
      message: error?.message,
      timestamp: Date.now()
    })
  }
  
  /**
   * Record UX metric
   */
  recordUXMetric(metric, value, metadata = {}) {
    if (!this.shouldSample()) return
    
    if (this.metrics.ux.hasOwnProperty(metric)) {
      this.metrics.ux[metric] = value
    }
    
    this.addToHistory({
      type: 'ux-metric',
      metric,
      value,
      timestamp: Date.now(),
      ...metadata
    })
  }
  
  /**
   * Start memory monitoring
   */
  startMemoryMonitoring() {
    if (typeof performance === 'undefined' || !performance.memory) return
    
    const updateMemoryMetrics = () => {
      const memory = performance.memory
      this.metrics.memory.used = memory.usedJSHeapSize
      this.metrics.memory.peak = Math.max(this.metrics.memory.peak, memory.usedJSHeapSize)
      
      // Check for high memory usage
      if (memory.usedJSHeapSize > this.options.alertThresholds.memoryUsage) {
        this.createAlert('high-memory-usage', {
          used: memory.usedJSHeapSize,
          threshold: this.options.alertThresholds.memoryUsage
        })
      }
    }
    
    // Update every 30 seconds
    setInterval(updateMemoryMetrics, 30000)
    updateMemoryMetrics() // Initial update
  }
  
  /**
   * Start performance observer
   */
  startPerformanceObserver() {
    if (typeof PerformanceObserver === 'undefined') return
    
    try {
      const observer = new PerformanceObserver((list) => {
        for (const entry of list.getEntries()) {
          if (entry.entryType === 'navigation') {
            this.recordUXMetric('navigationTime', entry.loadEventEnd - entry.fetchStart)
          }
        }
      })
      
      observer.observe({ entryTypes: ['navigation'] })
    } catch (error) {
      console.warn('[PerformanceMonitor] Failed to start performance observer:', error)
    }
  }
  
  /**
   * Classify error type
   */
  classifyError(error) {
    if (!error) return 'unknown'
    
    const status = error.response?.status
    if (status === 401 || status === 403) return 'auth'
    if (status === 404) return 'not-found'
    if (status >= 400 && status < 500) return 'client'
    if (status >= 500) return 'server'
    if (error.code === 'NETWORK_ERROR') return 'network'
    if (error.message?.includes('timeout')) return 'timeout'
    
    return 'unknown'
  }
  
  /**
   * Create performance alert
   */
  createAlert(type, data) {
    const alert = {
      id: `${type}_${Date.now()}`,
      type,
      data,
      timestamp: Date.now(),
      acknowledged: false
    }
    
    this.alerts.value.push(alert)
    
    if (this.options.enableLogging) {
      console.warn(`[PerformanceMonitor] Alert: ${type}`, data)
    }
    
    // Auto-dismiss after 5 minutes
    setTimeout(() => {
      this.dismissAlert(alert.id)
    }, 5 * 60 * 1000)
  }
  
  /**
   * Dismiss alert
   */
  dismissAlert(alertId) {
    const index = this.alerts.value.findIndex(alert => alert.id === alertId)
    if (index !== -1) {
      this.alerts.value.splice(index, 1)
    }
  }
  
  /**
   * Add entry to history
   */
  addToHistory(entry) {
    this.history.push(entry)
    
    // Maintain max history size
    if (this.history.length > this.options.maxHistorySize) {
      this.history.shift()
    }
  }
  
  /**
   * Check if should sample this event
   */
  shouldSample() {
    return Math.random() < this.options.sampleRate
  }
  
  /**
   * Get performance report
   */
  getReport() {
    return {
      metrics: { ...this.metrics },
      alerts: this.alerts.value,
      summary: {
        totalApiCalls: this.metrics.apiCalls.total,
        averageApiDuration: Math.round(this.metrics.apiCalls.averageDuration),
        cacheHitRate: Math.round(this.metrics.cache.hitRate * 100),
        errorRate: Math.round(this.metrics.errors.rate * 100),
        memoryUsage: Math.round(this.metrics.memory.used / 1024 / 1024) + 'MB'
      },
      history: this.history.slice(-100) // Last 100 entries
    }
  }
  
  /**
   * Reset all metrics
   */
  reset() {
    Object.keys(this.metrics.apiCalls).forEach(key => {
      this.metrics.apiCalls[key] = 0
    })
    Object.keys(this.metrics.cache).forEach(key => {
      this.metrics.cache[key] = 0
    })
    Object.keys(this.metrics.state).forEach(key => {
      this.metrics.state[key] = 0
    })
    Object.keys(this.metrics.errors).forEach(key => {
      if (typeof this.metrics.errors[key] === 'object') {
        this.metrics.errors[key] = {}
      } else {
        this.metrics.errors[key] = 0
      }
    })
    
    this.history = []
    this.alerts.value = []
    this.activeOperations.clear()
  }
}

// Global performance monitor instance
export const globalPerformanceMonitor = new PerformanceMonitor()

/**
 * Composable for performance monitoring
 */
export function usePerformanceMonitor(options = {}) {
  const monitor = options.global ? globalPerformanceMonitor : new PerformanceMonitor(options)
  
  const startApiCall = (operation, metadata) => monitor.startApiCall(operation, metadata)
  const endApiCall = (callId, success, error) => monitor.endApiCall(callId, success, error)
  const recordCacheEvent = (type, key, metadata) => monitor.recordCacheEvent(type, key, metadata)
  const recordStateUpdate = (operation, metadata) => monitor.recordStateUpdate(operation, metadata)
  const recordSyncEvent = (type, metadata) => monitor.recordSyncEvent(type, metadata)
  const recordError = (error, operation) => monitor.recordError(error, operation)
  const recordUXMetric = (metric, value, metadata) => monitor.recordUXMetric(metric, value, metadata)
  
  const metrics = computed(() => monitor.metrics)
  const alerts = computed(() => monitor.alerts.value)
  const report = computed(() => monitor.getReport())
  
  return {
    startApiCall,
    endApiCall,
    recordCacheEvent,
    recordStateUpdate,
    recordSyncEvent,
    recordError,
    recordUXMetric,
    metrics,
    alerts,
    report,
    dismissAlert: (id) => monitor.dismissAlert(id),
    reset: () => monitor.reset()
  }
}
