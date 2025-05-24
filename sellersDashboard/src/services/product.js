// src/services/product.js
import api from './api'

export const productService = {
  // Seller-only (dashboard) fetch:
  async fetchByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return mapProducts(res.data)
  },

  // Public storefront fetch:
  async fetchPublicByShop(shopId) {
    // This hits your public route: GET /shops/:shopId/products
    const res = await api.get(`/shops/${shopId}/products`)
    return mapProducts(res.data)
  },

  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    return mapProduct(res.data)
  },

  // … other methods unchanged …

}

// helper to map an array
function mapProducts(arr) {
  return arr.map(p => mapProduct(p))
}

// helper to normalize a single product
function mapProduct(p) {
  return {
    id:          p.ID ?? p.id,
    shopId:      p.shop_id ?? p.ShopID,
    userId:      p.user_id ?? p.UserID,
    name:        p.Name ?? p.name,
    description: p.Description ?? p.description,
    images:      p.Images ?? p.images ?? [],
    category:    p.Category ?? p.category ?? null,
    price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
    variants:    p.Variants ?? p.variants ?? [],
    createdAt:   p.CreatedAt ?? p.createdAt,
    updatedAt:   p.UpdatedAt ?? p.updatedAt ?? null,
  }
}
