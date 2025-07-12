import api from './api';

// TypeScript interfaces for cart with discount details
export interface ItemDiscountDetail {
  product_id: string;
  variant_id?: string;
  discount_id: string;
  name: string;
  type: 'fixed' | 'percentage';
  value: number;
  category: 'product' | 'order' | 'shipping' | 'buy_x_get_y';
  amount: number;
}

export interface OrderDiscountDetail {
  discount_id: string;
  name: string;
  type: 'fixed' | 'percentage';
  value: number;
  category: 'product' | 'order' | 'shipping' | 'buy_x_get_y';
  amount: number;
}

export interface CartItem {
  id: string;
  product_id: string;
  variant_id?: string;
  product_name: string;
  variant_options?: Record<string, string>;
  image?: string;
  unit_price: number;
  quantity: number;
  line_total: number;
  discount_amount: number;
  final_line_total: number;
  applied_discount_ids: string[];
}

export interface CartWithDiscountDetails {
  id: string;
  customer_id?: string;
  session_id?: string;
  shop_id: string;
  items: CartItem[];
  subtotal: number;
  total_discounts: number;
  shipping_cost: number;
  tax_amount: number;
  grand_total: number;
  applied_discount_ids: string[];
  currency: string;
  last_updated: string;
  created_at: string;
  item_discount_details: ItemDiscountDetail[];
  order_discount_details: OrderDiscountDetail[];
}

export async function getCart(shopSlug: string) {
  return api.get<CartWithDiscountDetails>(`/shops/${shopSlug}/cart`);
}

export async function addToCart(shopSlug: string, productId: string, variantId: string, quantity: number) {
  const response = await api.post<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/items`, { product_id: productId, variant_id: variantId, quantity });
  return response;
}

export async function updateCartItem(shopSlug: string, productId: string, variantId: string, quantity: number) {
  return api.put<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/items`, { product_id: productId, variant_id: variantId, quantity });
}

export async function removeCartItem(shopSlug: string, productId: string, variantId: string) {
  return api.delete<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/items`, { data: { product_id: productId, variant_id: variantId } });
}

export async function clearCart(shopSlug: string) {
  return api.post<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/clear`);
}

export async function applyDiscount(shopSlug: string, code: string) {
  return api.post<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/discounts`, { code });
}

export async function removeDiscount(shopSlug: string, discountId: string) {
  return api.delete<CartWithDiscountDetails>(`/shops/${shopSlug}/cart/discounts/${discountId}`);
} 