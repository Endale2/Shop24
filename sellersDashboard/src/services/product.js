// services/product.js
import api from './api'

export const productService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id: p.ID ?? p.id,
      shopId: p.shop_id ?? p.ShopID,
      name: p.Name ?? p.name,
      description: p.Description ?? p.description,
      images: p.Images ?? p.images ?? [],
      main_image: p.main_image ?? p.MainImage, // Ensure main_image is mapped
      category: p.Category ?? p.category,
      price: Number(p.Price ?? p.price) || 0,
      variants: (p.Variants ?? p.variants ?? []).map(v => ({
        variantID: v.VariantID ?? v.id, // Assuming VariantID or id for variant
        options: v.Options ?? v.options,
        price: Number(v.Price ?? v.price) || 0,
        stock: v.Stock ?? v.stock,
        image: v.Image ?? v.image, // Ensure variant image is mapped
        createdAt: v.CreatedAt ?? v.createdAt,
        updatedAt: v.UpdatedAt ?? v.updatedAt,
      })),
      createdAt: p.CreatedAt ?? p.createdAt,
      updatedAt: p.UpdatedAt ?? p.updatedAt,
    }))
  },
  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id: p.ID ?? p.id,
      shopId: p.shop_id ?? p.ShopID,
      name: p.Name ?? p.name,
      description: p.Description ?? p.description,
      images: p.Images ?? p.images ?? [],
      main_image: p.main_image ?? p.MainImage, // Ensure main_image is mapped
      category: p.Category ?? p.category,
      price: Number(p.Price ?? p.price) || 0,
      variants: (p.Variants ?? p.variants ?? []).map(v => ({
        variantID: v.VariantID ?? v.id, // Assuming VariantID or id for variant
        options: v.Options ?? v.options,
        price: Number(v.Price ?? v.price) || 0,
        stock: v.Stock ?? v.stock,
        image: v.Image ?? v.image, // Ensure variant image is mapped
        createdAt: v.CreatedAt ?? v.createdAt,
        updatedAt: v.UpdatedAt ?? v.updatedAt,
      })),
      createdAt: p.CreatedAt ?? p.createdAt,
      updatedAt: p.UpdatedAt ?? p.updatedAt,
    }
  },
  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    return this.fetchById(shopId, res.data.id ?? res.data.insertedId)
  },
  async update(shopId, productId, data) {
    // This is typically a PUT request for full replacement
    await api.put(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },
  async patch(shopId, productId, data) {
    // This is for partial updates (sending only changed fields)
    // If your backend specifically expects a PATCH verb for updates, use this.
    await api.patch(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId) // Fetch updated product after patch
  },
  async delete(shopId, productId) {
    await api.delete(`/seller/shops/${shopId}/products/${productId}`)
  }
}