import api from './api';

export async function getCart(shopSlug: string) {
  return api.get(`/shops/${shopSlug}/cart`);
}

export async function addToCart(shopSlug: string, productId: string, variantId: string, quantity: number) {
  const response = await api.post(`/shops/${shopSlug}/cart/items`, { product_id: productId, variant_id: variantId, quantity });
  return response;
}

export async function updateCartItem(shopSlug: string, productId: string, variantId: string, quantity: number) {
  return api.put(`/shops/${shopSlug}/cart/items`, { product_id: productId, variant_id: variantId, quantity });
}

export async function removeCartItem(shopSlug: string, productId: string, variantId: string) {
  return api.delete(`/shops/${shopSlug}/cart/items`, { data: { product_id: productId, variant_id: variantId } });
}

export async function applyDiscount(shopSlug: string, code: string) {
  return api.post(`/shops/${shopSlug}/cart/discount`, { code });
}

export async function clearCart(shopSlug: string) {
  return api.post(`/shops/${shopSlug}/cart/clear`);
} 