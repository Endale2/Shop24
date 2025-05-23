import api from './api';

export const productService = {
  async fetchByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`);
    return res.data.map(p => ({
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      userId:      p.user_id ?? p.UserID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images,
      category:    p.Category ?? p.category,
      price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
      variants:    p.Variants ?? p.variants,
      createdAt:   p.CreatedAt ?? p.createdAt,
    }));
  },

  // fetch a single product
  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`);
    const p = res.data;
    return {
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      userId:      p.user_id ?? p.UserID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images,
      category:    p.Category ?? p.category,
      price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
      variants:    p.Variants ?? p.variants,
      createdAt:   p.CreatedAt ?? p.createdAt,
      updatedAt:   p.UpdatedAt ?? p.updatedAt,
    };
  }
};
