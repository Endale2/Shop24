// utils/stateSynchronizer.js
import { ref, reactive, watch, nextTick } from 'vue'

/**
 * State synchronization utility for cross-component and cross-tab consistency
 */
export class StateSynchronizer {
  constructor(options = {}) {
    this.options = {
      enableCrossTab: true,
      enableBroadcast: true,
      debounceDelay: 100,
      maxRetries: 3,
      ...options
    }
    
    this.subscribers = new Map()
    this.eventQueue = []
    this.isProcessing = false
    this.broadcastChannel = null
    this.syncState = reactive({
      lastSync: null,
      syncCount: 0,
      conflicts: 0,
      errors: 0
    })
    
    this.init()
  }
  
  /**
   * Initialize synchronizer
   */
  init() {
    if (this.options.enableCrossTab && typeof BroadcastChannel !== 'undefined') {
      this.broadcastChannel = new BroadcastChannel('shop-state-sync')
      this.broadcastChannel.addEventListener('message', this.handleBroadcastMessage.bind(this))
    }
    
    // Listen for visibility changes to sync when tab becomes active
    if (typeof document !== 'undefined') {
      document.addEventListener('visibilitychange', this.handleVisibilityChange.bind(this))
    }
    
    // Listen for online/offline events
    if (typeof window !== 'undefined') {
      window.addEventListener('online', this.handleOnline.bind(this))
      window.addEventListener('offline', this.handleOffline.bind(this))
    }
  }
  
  /**
   * Subscribe to state changes
   */
  subscribe(key, callback, options = {}) {
    const subscription = {
      key,
      callback,
      options: {
        immediate: false,
        debounce: this.options.debounceDelay,
        ...options
      },
      lastValue: null,
      debounceTimer: null
    }
    
    if (!this.subscribers.has(key)) {
      this.subscribers.set(key, new Set())
    }
    
    this.subscribers.get(key).add(subscription)
    
    // Return unsubscribe function
    return () => {
      const keySubscribers = this.subscribers.get(key)
      if (keySubscribers) {
        keySubscribers.delete(subscription)
        if (keySubscribers.size === 0) {
          this.subscribers.delete(key)
        }
      }
      
      if (subscription.debounceTimer) {
        clearTimeout(subscription.debounceTimer)
      }
    }
  }
  
  /**
   * Emit state change event
   */
  emit(key, value, metadata = {}) {
    const event = {
      key,
      value,
      metadata: {
        timestamp: Date.now(),
        source: 'local',
        version: this.generateVersion(),
        ...metadata
      }
    }
    
    this.eventQueue.push(event)
    this.processEventQueue()
    
    // Broadcast to other tabs if enabled
    if (this.options.enableCrossTab && this.broadcastChannel) {
      this.broadcastChannel.postMessage({
        type: 'state-change',
        event: {
          ...event,
          metadata: {
            ...event.metadata,
            source: 'broadcast'
          }
        }
      })
    }
  }
  
  /**
   * Process event queue
   */
  async processEventQueue() {
    if (this.isProcessing || this.eventQueue.length === 0) {
      return
    }
    
    this.isProcessing = true
    
    try {
      while (this.eventQueue.length > 0) {
        const event = this.eventQueue.shift()
        await this.processEvent(event)
      }
    } catch (error) {
      console.error('[StateSynchronizer] Error processing event queue:', error)
      this.syncState.errors += 1
    } finally {
      this.isProcessing = false
    }
  }
  
  /**
   * Process individual event
   */
  async processEvent(event) {
    const subscribers = this.subscribers.get(event.key)
    if (!subscribers || subscribers.size === 0) {
      return
    }
    
    for (const subscription of subscribers) {
      try {
        await this.notifySubscriber(subscription, event)
      } catch (error) {
        console.error('[StateSynchronizer] Error notifying subscriber:', error)
      }
    }
    
    this.syncState.lastSync = Date.now()
    this.syncState.syncCount += 1
  }
  
  /**
   * Notify individual subscriber
   */
  async notifySubscriber(subscription, event) {
    const { callback, options, lastValue } = subscription
    
    // Check if value actually changed
    if (options.skipDuplicates && this.deepEqual(event.value, lastValue)) {
      return
    }
    
    subscription.lastValue = event.value
    
    if (options.debounce > 0) {
      // Debounced notification
      if (subscription.debounceTimer) {
        clearTimeout(subscription.debounceTimer)
      }
      
      subscription.debounceTimer = setTimeout(() => {
        callback(event.value, event.metadata)
        subscription.debounceTimer = null
      }, options.debounce)
    } else {
      // Immediate notification
      await nextTick()
      callback(event.value, event.metadata)
    }
  }
  
  /**
   * Handle broadcast messages from other tabs
   */
  handleBroadcastMessage(messageEvent) {
    const { type, event } = messageEvent.data
    
    if (type === 'state-change') {
      // Avoid processing our own broadcasts
      if (event.metadata.source === 'broadcast') {
        this.eventQueue.push(event)
        this.processEventQueue()
      }
    }
  }
  
  /**
   * Handle tab visibility changes
   */
  handleVisibilityChange() {
    if (!document.hidden) {
      // Tab became visible, trigger sync
      this.emit('tab-visible', { timestamp: Date.now() })
    }
  }
  
  /**
   * Handle online event
   */
  handleOnline() {
    this.emit('network-online', { timestamp: Date.now() })
  }
  
  /**
   * Handle offline event
   */
  handleOffline() {
    this.emit('network-offline', { timestamp: Date.now() })
  }
  
  /**
   * Generate version for conflict resolution
   */
  generateVersion() {
    return `${Date.now()}-${Math.random().toString(36).substr(2, 9)}`
  }
  
  /**
   * Deep equality check
   */
  deepEqual(a, b) {
    if (a === b) return true
    if (a == null || b == null) return false
    if (typeof a !== typeof b) return false
    
    if (typeof a === 'object') {
      const keysA = Object.keys(a)
      const keysB = Object.keys(b)
      
      if (keysA.length !== keysB.length) return false
      
      for (const key of keysA) {
        if (!keysB.includes(key)) return false
        if (!this.deepEqual(a[key], b[key])) return false
      }
      
      return true
    }
    
    return false
  }
  
  /**
   * Get synchronization statistics
   */
  getStats() {
    return {
      ...this.syncState,
      subscriberCount: Array.from(this.subscribers.values())
        .reduce((total, set) => total + set.size, 0),
      queueLength: this.eventQueue.length,
      isProcessing: this.isProcessing
    }
  }
  
  /**
   * Clear all subscriptions and cleanup
   */
  destroy() {
    // Clear all subscriptions
    for (const [key, subscribers] of this.subscribers) {
      for (const subscription of subscribers) {
        if (subscription.debounceTimer) {
          clearTimeout(subscription.debounceTimer)
        }
      }
    }
    this.subscribers.clear()
    
    // Clear event queue
    this.eventQueue = []
    
    // Close broadcast channel
    if (this.broadcastChannel) {
      this.broadcastChannel.close()
      this.broadcastChannel = null
    }
    
    // Remove event listeners
    if (typeof document !== 'undefined') {
      document.removeEventListener('visibilitychange', this.handleVisibilityChange.bind(this))
    }
    
    if (typeof window !== 'undefined') {
      window.removeEventListener('online', this.handleOnline.bind(this))
      window.removeEventListener('offline', this.handleOffline.bind(this))
    }
  }
}

// Global synchronizer instance
export const globalSynchronizer = new StateSynchronizer()

/**
 * Composable for state synchronization
 */
export function useStateSynchronizer(options = {}) {
  const synchronizer = options.global ? globalSynchronizer : new StateSynchronizer(options)
  const subscriptions = ref([])
  
  const subscribe = (key, callback, subscriptionOptions = {}) => {
    const unsubscribe = synchronizer.subscribe(key, callback, subscriptionOptions)
    subscriptions.value.push(unsubscribe)
    return unsubscribe
  }
  
  const emit = (key, value, metadata = {}) => {
    synchronizer.emit(key, value, metadata)
  }
  
  const cleanup = () => {
    subscriptions.value.forEach(unsubscribe => unsubscribe())
    subscriptions.value = []
    
    if (!options.global) {
      synchronizer.destroy()
    }
  }
  
  return {
    subscribe,
    emit,
    cleanup,
    stats: synchronizer.syncState,
    getStats: () => synchronizer.getStats()
  }
}
