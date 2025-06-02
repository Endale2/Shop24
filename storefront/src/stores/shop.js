import { defineStore } from 'pinia'
import api from '@/services/api'

export const useShopStore = defineStore('shop', {
  state: () => ({
    current: null,
    products: []
  }),
  actions: {
    async fetchShop(slug) {
      const { data: s } = await api.get(`/shops/${slug}`)
      this.current = {
        id:          s.id,
        slug:        s.slug,
        name:        s.name,
        description: s.description,
        banner:      s.banner,
        createdAt:   s.createdAt,
        updatedAt:   s.updatedAt
      }
    },

    async fetchProducts(shopSlug) {
      const { data } = await api.get(`/shops/${shopSlug}/products`)
      this.products = data.map(p => ({
        id:            p.id,
        shop_id:       p.shop_id,
        user_id:       p.user_id,
        discount_id:   p.discount_id || [],
        name:          p.name,
        slug:          p.slug,
        description:   p.description,
        main_image:    p.main_image || '',
        images:        Array.isArray(p.images) ? p.images : [],
        category:      p.category,
        price:         typeof p.price === 'number' ? p.price : 0,
        display_price: typeof p.display_price === 'number' ? p.display_price : p.price,
        createdBy:     p.createdBy,
        variants:      Array.isArray(p.variants) ? p.variants.map(v => ({
                         id:        v.id,
                         options:   v.options || {},
                         price:     typeof v.price === 'number' ? v.price : 0,
                         stock:     v.stock || 0,
                         image:     v.image || '',
                         createdAt: v.createdAt,
                         updatedAt: v.updatedAt
                       })) : [],
        createdAt:     p.createdAt,
        updatedAt:     p.updatedAt
      }))
    }
  },
  persist: true
})
