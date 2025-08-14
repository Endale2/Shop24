import { defineStore } from 'pinia'
import { shopService } from '@/services/shop'

export const useShopStore = defineStore('shops', {
  state: () => ({
    list:   [],
    active: null,
    loading: false,
    error: null,
    initialized: false,
    lastFetchTime: null,
    shopStateReady: false,
  }),
  getters: {
    allShops:   (state) => state.list,
    activeShop: (state) => state.active,
    isLoading:  (state) => state.loading,
    hasError:   (state) => !!state.error,
    errorMessage: (state) => state.error,
    isInitialized: (state) => state.initialized,
    isShopStateReady: (state) => state.shopStateReady && !!state.active,
    
    // Filter shops by status
    activeShops: (state) => state.list.filter(shop => shop.status === 'active'),
    inactiveShops: (state) => state.list.filter(shop => shop.status !== 'active'),
    
    // Get shops by category
    shopsByCategory: (state) => (category) => {
      return state.list.filter(shop => shop.category === category)
    },
    
    // Get total revenue across all shops
    totalRevenue: (state) => {
      return state.list.reduce((total, shop) => total + (shop.revenue || 0), 0)
    },

    // Validate if active shop is still valid
    isActiveShopValid: (state) => {
      if (!state.active) return false
      return state.list.some(shop => shop.id === state.active.id)
    },

    // Get shop by ID with validation
    getShopById: (state) => (shopId) => {
      return state.list.find(shop => shop.id === shopId) || null
    },
    
    // Get total customers across all shops
    totalCustomers: (state) => {
      return state.list.reduce((total, shop) => total + (shop.customer_count || 0), 0)
    },
    
    // Get total products across all shops
    totalProducts: (state) => {
      return state.list.reduce((total, shop) => total + (shop.product_count || 0), 0)
    },
  },
  actions: {
    // Seller-only: list all shops
    async fetchShops(forceRefresh = false) {
      // Avoid unnecessary fetches if recently fetched (unless forced)
      if (!forceRefresh && this.initialized && this.lastFetchTime &&
          Date.now() - this.lastFetchTime < 5 * 60 * 1000) { // 5 minutes cache
        return this.list
      }

      this.loading = true
      this.error = null

      try {
        console.log('[ShopStore] Fetching shops...')
        this.list = await shopService.fetchShops()
        this.lastFetchTime = Date.now()
        this.initialized = true

        // Validate current active shop
        if (this.active && !this.isActiveShopValid) {
          console.warn('[ShopStore] Active shop is no longer valid, clearing...')
          this.active = null
        }

        // If no active shop is set and we have shops, set the first one as active
        if (!this.active && this.list.length > 0) {
          console.log('[ShopStore] Setting first shop as active:', this.list[0].name)
          this.active = this.list[0]
        }

        // Update shop state ready flag
        this.shopStateReady = !!this.active

        console.log('[ShopStore] Shops loaded successfully:', {
          count: this.list.length,
          activeShop: this.active?.name,
          shopStateReady: this.shopStateReady
        })

        return this.list
      } catch (error) {
        this.error = error.message || 'Failed to fetch shops'
        this.initialized = false
        this.shopStateReady = false
        console.error('[ShopStore] Failed to fetch shops:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    // Seller-only: create a shop
    async createShop(payload) {
      this.loading = true
      this.error = null
      
      try {
        const newShop = await shopService.createShop(payload)
        await this.fetchShops() // Refresh the list
        return newShop
      } catch (error) {
        this.error = error.message || 'Failed to create shop'
        throw error
      } finally {
        this.loading = false
      }
    },

    // Update a shop
    async updateShop(shopId, updates) {
      this.loading = true
      this.error = null
      
      try {
        const updatedShop = await shopService.updateShop(shopId, updates)
        // Update the shop in the list
        const index = this.list.findIndex(shop => shop.id === shopId)
        if (index !== -1) {
          this.list[index] = { ...this.list[index], ...updatedShop }
        }
        // Update active shop if it's the one being updated
        if (this.active && this.active.id === shopId) {
          this.active = { ...this.active, ...updatedShop }
        }
        return updatedShop
      } catch (error) {
        this.error = error.message || 'Failed to update shop'
        throw error
      } finally {
        this.loading = false
      }
    },

    // Delete a shop
    async deleteShop(shopId) {
      this.loading = true
      this.error = null
      
      try {
        await shopService.deleteShop(shopId)
        // Remove from list
        this.list = this.list.filter(shop => shop.id !== shopId)
        // Clear active shop if it's the one being deleted
        if (this.active && this.active.id === shopId) {
          this.active = null
        }
      } catch (error) {
        this.error = error.message || 'Failed to delete shop'
        throw error
      } finally {
        this.loading = false
      }
    },

    // Set active (for dashboard)
    setActiveShop(shop) {
      if (!shop) {
        console.warn('[ShopStore] Attempting to set null shop as active')
        this.active = null
        this.shopStateReady = false
        return
      }

      // Validate shop exists in the list
      const validShop = this.list.find(s => s.id === shop.id)
      if (!validShop) {
        console.error('[ShopStore] Attempting to set invalid shop as active:', shop)
        return
      }

      console.log('[ShopStore] Setting active shop:', shop.name)
      this.active = validShop
      this.shopStateReady = true
      this.error = null // Clear any previous errors
    },

    // Ensure active shop is valid and available
    async ensureActiveShop() {
      // If no active shop, try to set one
      if (!this.active) {
        if (this.list.length > 0) {
          this.setActiveShop(this.list[0])
          return this.active
        } else {
          // Try to fetch shops if list is empty
          try {
            await this.fetchShops()
            return this.active
          } catch (error) {
            console.error('[ShopStore] Failed to ensure active shop:', error)
            return null
          }
        }
      }

      // Validate current active shop
      if (!this.isActiveShopValid) {
        console.warn('[ShopStore] Current active shop is invalid, selecting new one')
        if (this.list.length > 0) {
          this.setActiveShop(this.list[0])
        } else {
          this.active = null
          this.shopStateReady = false
        }
      }

      return this.active
    },

    // Clear on logout
    clearActiveShop() {
      this.active = null
    },
    
    // Clear error
    clearError() {
      this.error = null
    },
    
    // Reset store
    $reset() {
      this.list = []
      this.active = null
      this.loading = false
      this.error = null
      this.initialized = false
      this.lastFetchTime = null
      this.shopStateReady = false
    }
  },
  persist: {
    paths: ['list', 'active', 'initialized', 'lastFetchTime', 'shopStateReady'],
    storage: localStorage,
    serializer: {
      serialize: (value) => {
        try {
          return JSON.stringify(value)
        } catch (error) {
          console.error('[ShopStore] Serialization error:', error)
          return '{}'
        }
      },
      deserialize: (value) => {
        try {
          return JSON.parse(value)
        } catch (error) {
          console.error('[ShopStore] Deserialization error:', error)
          return {}
        }
      }
    },
    afterRestore: (context) => {
      // Validate restored state
      const store = context.store

      // Check if active shop is still valid
      if (store.active && store.list.length > 0) {
        const isValid = store.list.some(shop => shop.id === store.active.id)
        if (!isValid) {
          console.warn('[ShopStore] Restored active shop is invalid, clearing...')
          store.active = null
          store.shopStateReady = false
        }
      }

      // Check if data is stale (older than 1 hour)
      if (store.lastFetchTime && Date.now() - store.lastFetchTime > 60 * 60 * 1000) {
        console.log('[ShopStore] Restored data is stale, will refresh on next fetch')
        store.initialized = false
      }
    }
  }
})
