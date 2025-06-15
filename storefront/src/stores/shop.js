// src/stores/shop.js
import { defineStore }   from 'pinia'
import { getShop, getProducts } from '@/services/shop'

export const useShopStore = defineStore('shop', {
  state: () => ({
    current: null,
    products: [],
    loading: false
  }),
  actions: {
    async fetchShopAndProducts(slug) {
      this.loading = true
      const [s, products] = await Promise.all([
        getShop(slug),
        getProducts(slug)
      ])
      this.current  = s
      this.products = products
      this.loading  = false
    }
  },
  persist: true
})
