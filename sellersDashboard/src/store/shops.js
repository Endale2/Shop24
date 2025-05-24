// src/stores/shops.js
import { defineStore } from 'pinia'
import { shopService } from '@/services/shop'

export const useShopStore = defineStore('shops', {
  state: () => ({
    list: [],
    active: null,
  }),
  getters: {
    allShops: (state) => state.list,
    activeShop: (state) => state.active,
  },
  actions: {
    async fetchShops() {
      this.list = await shopService.fetchShops()
      return this.list
    },
    async createShop(payload) {
      const shop = await shopService.createShop(payload)
      await this.fetchShops()
      return shop
    },
    setActiveShop(shop) {
      this.active = shop
    }
  },
  persist: {
    // persist both `list` and `active` to localStorage
    paths: ['list', 'active']
  }
})
