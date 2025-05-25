import { defineStore } from 'pinia'
import api from '@/services/api'

export const useShopStore = defineStore('shop', {
  state: () => ({
    current: null,
    products: []
  }),
  actions: {
    // Fetch and normalize a single shop by slug
    async fetchShop(slug) {
      const { data: s } = await api.get(`/shops/${slug}`)
      this.current = {
        id:          s.id,
        slug:        s.slug,
        name:        s.name,
        description: s.description,
        createdAt:   s.createdAt,
        updatedAt:   s.updatedAt
      }
    },

    // Fetch and normalize products list by shop slug
    async fetchProducts(shopSlug) {
      const { data } = await api.get(`/shops/${shopSlug}/products`)
      this.products = data.map(p => ({
        id:            p.id,
        slug:          p.slug,
        name:          p.name,
        description:   p.description,
        images:        p.images || [],
        price:         isFinite(p.price) ? p.price : 0,
        createdAt:     p.createdAt,
        updatedAt:     p.updatedAt
      }))
    }
  },
  persist: true
})
