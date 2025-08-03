import { defineStore } from 'pinia'
import { shopService } from '@/services/shop'

export const useShopStore = defineStore('shops', {
  state: () => ({
    list:   [],
    active: null,
    loading: false,
    error: null,
  }),
  getters: {
    allShops:   (state) => state.list,
    activeShop: (state) => state.active,
    isLoading:  (state) => state.loading,
    hasError:   (state) => !!state.error,
    errorMessage: (state) => state.error,
    
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
    async fetchShops() {
      this.loading = true
      this.error = null
      
      try {
        this.list = await shopService.fetchShops()
        return this.list
      } catch (error) {
        this.error = error.message || 'Failed to fetch shops'
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
      this.active = shop
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
    }
  },
  persist: {
    paths: ['list', 'active']
  }
})
