import api from './api';

export async function getWishlist(shopSlug: string) {
  return api.get(`/shops/${shopSlug}/wishlist`);
}

export async function addToWishlist(shopSlug: string, productId: string) {
  return api.post(`/shops/${shopSlug}/wishlist`, { product_id: productId });
}

export async function removeFromWishlist(shopSlug: string, productId: string) {
  return api.delete(`/shops/${shopSlug}/wishlist/${productId}`);
} 