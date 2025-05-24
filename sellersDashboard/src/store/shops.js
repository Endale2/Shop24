// src/stores/shops.js
import { defineStore } from 'pinia'
import { shopService } from '@/services/shop'

export const useShopStore = defineStore('shops', {
  state: () => ({
    list: [],
    active: null,
  }),
  getters: {
    allShops:    (state) => state.list,
    activeShop:  (state) => state.active,
  },
  actions: {
    /**
     * Seller-only: fetch the full list of shops for the logged-in user.
     */
    async fetchShops() {
      this.list = await shopService.fetchShops()
      return this.list
    },

    /**
     * Seller-only: create a new shop and refresh the list.
     */
    async createShop(payload) {
      const newShop = await shopService.createShop(payload)
      await this.fetchShops()
      return newShop
    },

    /**
     * Public: fetch a single shop by ID (no auth required).
     */
    async fetchPublicShop(shopId) {
      this.active = await shopService.fetchById(shopId)
      return this.active
    },

    /**
     * Set the active shop in state (and persist).
     */
    setActiveShop(shop) {
      this.active = shop
    },

    /**
     * Clear the active shop (e.g. on logout).
     */
    clearActiveShop() {
      this.active = null
    },
  },
  persist: {
    // persist both `list` (for seller dashboard) and `active` (for storefront)
    paths: ['list', 'active'],
  },
})
