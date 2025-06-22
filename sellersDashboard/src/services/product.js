import api from './api'

export const productService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id:           p.ID ?? p.id,
      shopId:       p.shop_id ?? p.ShopID,
      name:         p.Name ?? p.name,
      description:  p.Description ?? p.description,
      images:       p.Images ?? p.images ?? [],
      main_image:   p.main_image ?? p.MainImage ?? '',
      category:     p.Category ?? p.category,
      price:        Number(p.Price ?? p.price) || 0,
      stock:        p.Stock ?? p.stock ?? 0, // <--- This line is already correct here
      variants:     (p.Variants ?? p.variants ?? []).map(v => ({
        id:         v.VariantID ?? v.id,
        options:    v.Options ?? v.options ?? {},
        price:      Number(v.Price ?? v.price) || 0,
        stock:      v.Stock ?? v.stock ?? 0,
        image:      v.Image ?? v.image ?? '',
        createdAt:  v.CreatedAt ?? v.createdAt,
        updatedAt:  v.UpdatedAt ?? v.updatedAt,
      })),
      createdAt:    p.CreatedAt ?? p.createdAt,
      updatedAt:    p.UpdatedAt ?? p.updatedAt,
    }))
  },

  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id:           p.ID ?? p.id,
      shopId:       p.shop_id ?? p.ShopID,
      name:         p.Name ?? p.name,
      description:  p.Description ?? p.description,
      images:       p.Images ?? p.images ?? [],
      main_image:   p.main_image ?? p.MainImage ?? '',
      category:     p.Category ?? p.category,
      price:        Number(p.Price ?? p.price) || 0,
      stock:        p.Stock ?? p.stock ?? 0, // <--- ADD THIS LINE!
      variants:     (p.Variants ?? p.variants ?? []).map(v => ({
        id:         v.VariantID ?? v.id,
        options:    v.Options ?? v.options ?? {},
        price:      Number(v.Price ?? v.price) || 0,
        stock:      v.Stock ?? v.stock ?? 0,
        image:      v.Image ?? v.image ?? '',
        createdAt:  v.CreatedAt ?? v.createdAt,
        updatedAt:  v.UpdatedAt ?? v.updatedAt,
      })),
      createdAt:    p.CreatedAt ?? p.createdAt,
      updatedAt:    p.UpdatedAt ?? p.updatedAt,
    }
  },

  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    const created = res.data
    const id = created.id
               ?? created._id
               ?? created.ID
               ?? created.InsertedID
    if (!id) {
      throw new Error('Failed to retrieve new product ID from response')
    }
    return this.fetchById(shopId, id)
  },

  async update(shopId, productId, data) {
    await api.put(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },

  async patch(shopId, productId, data) {
    await api.patch(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },

  async delete(shopId, productId) {
    await api.delete(`/seller/shops/${shopId}/products/${productId}`)
  }
}