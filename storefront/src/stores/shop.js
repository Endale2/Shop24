import { defineStore } from 'pinia'
import api from '@/services/api'

export const useShopStore = defineStore('shop', {
  state: () => ({
    current: null,
    products: []
  }),
  actions: {
    // Fetch and normalize a single shop
    async fetchShop(id) {
      const { data: s } = await api.get(`/shops/${id}`)
      this.current = {
        id:          s.ID,
        name:        s.Name,
        description: s.Description,
        createdAt:   s.CreatedAt,
        updatedAt:   s.UpdatedAt
      }
    },

    // Fetch and normalize products list (example)
    async fetchProducts(id) {
      const { data } = await api.get(`/shops/${id}/products`)
      this.products = data.map(p => ({
        id:          p.ID ?? p.id,
        name:        p.Name ?? p.name,
        description: p.Description ?? p.description,
        images:      p.Images ?? p.images ?? [],
        price:       isFinite(p.Price) ? p.Price : p.price
      }))
    }
  },
  persist: true
})
