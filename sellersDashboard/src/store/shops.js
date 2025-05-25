import { defineStore } from 'pinia'
import { shopService } from '@/services/shop'

export const useShopStore = defineStore('shops', {
  state: () => ({
    list:   [],
    active: null,
  }),
  getters: {
    allShops:   (state) => state.list,
    activeShop: (state) => state.active,
  },
  actions: {
    // Seller-only: list all shops
    async fetchShops() {
      this.list = await shopService.fetchShops()
      return this.list
    },

    // Seller-only: create a shop
    async createShop(payload) {
      const newShop = await shopService.createShop(payload)
      await this.fetchShops()
      return newShop
    },

    // Set active (for dashboard)
    setActiveShop(shop) {
      this.active = shop
    },

    // Clear on logout
    clearActiveShop() {
      this.active = null
    }
  },
  persist: {
    paths: ['list', 'active']
  }
})
