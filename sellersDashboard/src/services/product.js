import api from './api';

export const productService = {
  // Fetch all products for a given shop ID
  async fetchByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`);
    return res.data.map(p => ({
      id:        p.ID ?? p.id,
      shopId:    p.shop_id ?? p.ShopID,
      userId:    p.user_id ?? p.UserID,
      name:      p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:    p.Images ?? p.images,
      category:  p.Category ?? p.category,
      price:     isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
      variants:  p.Variants ?? p.variants,
      createdAt: p.CreatedAt ?? p.createdAt,
    }));
  }
};
