import api from './api';
import { getCurrentShopSlug } from './shop'

export async function getWishlist(shopSlug?: string) {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return { data: { product_ids: [] } };
  return api.get(`/shops/${slug}/wishlist`);
}

export async function addToWishlist(shopSlug: string, productId: string) {
  return api.post(`/shops/${shopSlug}/wishlist`, { product_id: productId });
}

export async function removeFromWishlist(shopSlug: string, productId: string) {
  return api.delete(`/shops/${shopSlug}/wishlist/${productId}`);
} 