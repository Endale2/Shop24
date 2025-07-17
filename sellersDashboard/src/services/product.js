import api from './api'

export const productService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id: p.id || p._id || p.ID,
      shopId: p.shop_id || p.shopId || p.ShopID,
      name: p.name || p.Name,
      description: p.description || p.Description,
      images: p.images || p.Images || [],
      main_image: p.main_image || p.MainImage || '',
      category: p.category || p.Category,
      price: Number(p.price || p.Price) || 0,
      stock: p.stock || p.Stock || 0,
      starting_price: p.starting_price,
      total_stock: p.total_stock,
      variants: (p.variants || p.Variants || []).map(v => ({
        id: v.id || v.variant_id || v.VariantID,
        options: v.options || v.Options || [],
        price: Number(v.price || v.Price) || 0,
        stock: v.stock || v.Stock || 0,
        image: v.image || v.Image || '',
        createdAt: v.createdAt || v.CreatedAt,
        updatedAt: v.updatedAt || v.UpdatedAt,
      })),
      createdAt: p.createdAt || p.CreatedAt,
      updatedAt: p.updatedAt || p.UpdatedAt,
      meta_title: p.meta_title || '',
      meta_description: p.meta_description || '',
    }))
  },

  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id: p.id || p._id || p.ID,
      shopId: p.shop_id || p.shopId || p.ShopID,
      name: p.name || p.Name,
      description: p.description || p.Description,
      images: p.images || p.Images || [],
      main_image: p.main_image || p.MainImage || '',
      category: p.category || p.Category,
      price: Number(p.price || p.Price) || 0,
      stock: p.stock || p.Stock || 0,
      starting_price: p.starting_price,
      total_stock: p.total_stock,
      variants: (p.variants || p.Variants || []).map(v => ({
        id: v.id || v.variant_id || v.VariantID,
        options: v.options || v.Options || [],
        price: Number(v.price || v.Price) || 0,
        stock: v.stock || v.Stock || 0,
        image: v.image || v.Image || '',
        createdAt: v.createdAt || v.CreatedAt,
        updatedAt: v.updatedAt || v.UpdatedAt,
      })),
      createdAt: p.createdAt || p.CreatedAt,
      updatedAt: p.updatedAt || p.UpdatedAt,
      meta_title: p.meta_title || '',
      meta_description: p.meta_description || '',
    }
  },

  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    const created = res.data
    const id = created.id || created._id || created.ID || created.InsertedID
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