import { defineStore } from 'pinia'
import { getProducts } from '@/services/shop'

export const useProductStore = defineStore('product', {
  state: () => ({
    list: [],
    loading: false,
    currentShopSlug: null,   // Tracks which shop's products are loaded
  }),
  actions: {
    /**
     * Fetch all products for a given shop slug.
     * Sets `currentShopSlug` so we know when to reload.
     */
    async fetchAll(shopSlug) {
      this.loading = true
      try {
        const products = await getProducts(shopSlug)
        this.list = products
        this.currentShopSlug = shopSlug
      } catch (error) {
        console.error(`Error fetching products for shop "${shopSlug}":`, error)
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Clear the current product list (e.g., on logout or shop change).
     */
    clear() {
      this.list = []
      this.currentShopSlug = null
    }
  }
})
