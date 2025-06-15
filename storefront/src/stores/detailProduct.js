// src/stores/detailProduct.js
import { defineStore } from 'pinia'
import { getProduct } from '@/services/productService'

export const useDetailProductStore = defineStore('detailProduct', {
  state: () => ({ item: null, loading: false }),
  actions: {
    async fetchOne(shopSlug, slug) {
      this.loading = true
      this.item = await getProduct(shopSlug, slug)
      this.loading = false
    }
  }
})
