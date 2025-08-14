// composables/useShopState.js
import { computed, watch, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useShopStore } from '@/store/shops'

/**
 * Composable for managing shop state across components
 * Provides consistent shop state validation, error handling, and recovery
 */
export function useShopState() {
  const router = useRouter()
  const shopStore = useShopStore()
  
  // Reactive state
  const isInitializing = ref(false)
  const initializationError = ref(null)
  
  // Computed properties
  const activeShop = computed(() => shopStore.activeShop)
  const isShopStateReady = computed(() => shopStore.isShopStateReady)
  const allShops = computed(() => shopStore.allShops)
  const isLoading = computed(() => shopStore.isLoading)
  const hasError = computed(() => shopStore.hasError)
  const errorMessage = computed(() => shopStore.errorMessage)
  
  /**
   * Ensures we have a valid shop state before proceeding
   * @param {boolean} redirectOnFailure - Whether to redirect to shop selection on failure
   * @returns {Promise<Object|null>} - The active shop or null if failed
   */
  async function ensureValidShopState(redirectOnFailure = true) {
    try {
      isInitializing.value = true
      initializationError.value = null
      
      console.log('[useShopState] Ensuring valid shop state...')
      
      // First check if we have an active shop
      if (activeShop.value?.id && shopStore.isActiveShopValid) {
        console.log('[useShopState] Active shop is valid:', activeShop.value.name)
        return activeShop.value
      }
      
      // Try to ensure we have an active shop
      const shop = await shopStore.ensureActiveShop()
      
      if (!shop) {
        console.warn('[useShopState] No shops available')
        initializationError.value = 'No shops available. Please create a shop first.'
        
        if (redirectOnFailure) {
          router.replace({ name: 'ShopSelection' })
        }
        return null
      }
      
      console.log('[useShopState] Shop state ensured:', shop.name)
      return shop
      
    } catch (error) {
      console.error('[useShopState] Failed to ensure valid shop state:', error)
      initializationError.value = error.message || 'Failed to initialize shop state'
      
      if (redirectOnFailure) {
        router.replace({ name: 'ShopSelection' })
      }
      return null
    } finally {
      isInitializing.value = false
    }
  }
  
  /**
   * Waits for shop state to be ready
   * @param {number} timeout - Timeout in milliseconds (default: 10 seconds)
   * @returns {Promise<Object|null>} - The active shop or null if timeout
   */
  async function waitForShopState(timeout = 10000) {
    return new Promise((resolve) => {
      // If already ready, resolve immediately
      if (isShopStateReady.value && activeShop.value) {
        resolve(activeShop.value)
        return
      }
      
      let timeoutId
      const unwatch = watch(
        [isShopStateReady, activeShop],
        ([ready, shop]) => {
          if (ready && shop) {
            clearTimeout(timeoutId)
            unwatch()
            resolve(shop)
          }
        },
        { immediate: true }
      )
      
      // Set timeout
      timeoutId = setTimeout(() => {
        unwatch()
        console.warn('[useShopState] Timeout waiting for shop state')
        resolve(null)
      }, timeout)
    })
  }
  
  /**
   * Refreshes shop data and ensures valid state
   * @param {boolean} forceRefresh - Force refresh even if recently fetched
   * @returns {Promise<Object|null>} - The active shop or null if failed
   */
  async function refreshShopState(forceRefresh = false) {
    try {
      console.log('[useShopState] Refreshing shop state...')
      await shopStore.fetchShops(forceRefresh)
      return await ensureValidShopState()
    } catch (error) {
      console.error('[useShopState] Failed to refresh shop state:', error)
      initializationError.value = error.message || 'Failed to refresh shop state'
      return null
    }
  }
  
  /**
   * Sets the active shop with validation
   * @param {Object} shop - The shop to set as active
   * @returns {boolean} - Success status
   */
  function setActiveShop(shop) {
    try {
      shopStore.setActiveShop(shop)
      return true
    } catch (error) {
      console.error('[useShopState] Failed to set active shop:', error)
      initializationError.value = error.message || 'Failed to set active shop'
      return false
    }
  }
  
  /**
   * Validates if a shop ID is valid
   * @param {string} shopId - The shop ID to validate
   * @returns {boolean} - Whether the shop ID is valid
   */
  function isValidShopId(shopId) {
    if (!shopId) return false
    return allShops.value.some(shop => shop.id === shopId)
  }
  
  /**
   * Gets a shop by ID
   * @param {string} shopId - The shop ID
   * @returns {Object|null} - The shop or null if not found
   */
  function getShopById(shopId) {
    return shopStore.getShopById(shopId)
  }
  
  /**
   * Clears any initialization errors
   */
  function clearError() {
    initializationError.value = null
    shopStore.clearError()
  }
  
  // Auto-recovery watcher
  watch(
    () => shopStore.hasError,
    (hasError) => {
      if (hasError) {
        console.warn('[useShopState] Shop store error detected, attempting recovery...')
        // Attempt recovery after a short delay
        setTimeout(() => {
          refreshShopState(true).catch(console.error)
        }, 2000)
      }
    }
  )
  
  return {
    // State
    activeShop,
    isShopStateReady,
    allShops,
    isLoading,
    hasError,
    errorMessage,
    isInitializing,
    initializationError,
    
    // Methods
    ensureValidShopState,
    waitForShopState,
    refreshShopState,
    setActiveShop,
    isValidShopId,
    getShopById,
    clearError
  }
}
