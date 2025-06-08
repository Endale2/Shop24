import api from './api'

export const productService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images ?? [],
      main_image:  p.main_image ?? p.main_image ,
      category:    p.Category ?? p.category,
      price:       Number(p.Price ?? p.price) || 0,
      variants:    p.Variants ?? p.variants ?? [],
      createdAt:   p.CreatedAt ?? p.createdAt,
      updatedAt:   p.UpdatedAt ?? p.updatedAt,
    }))
  },
  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images ?? [],
      main_image:  p.main_image ?? p.main_image ,
      category:    p.Category ?? p.category,
      price:       Number(p.Price ?? p.price) || 0,
      variants:    p.Variants ?? p.variants ?? [],
      createdAt:   p.CreatedAt ?? p.createdAt,
      updatedAt:   p.UpdatedAt ?? p.updatedAt,
    }
  },
  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    return this.fetchById(shopId, res.data.id ?? res.data.insertedId)
  },
  async update(shopId, productId, data) {
    await api.put(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },
  async delete(shopId, productId) {
    await api.delete(`/seller/shops/${shopId}/products/${productId}`)
  }
}
