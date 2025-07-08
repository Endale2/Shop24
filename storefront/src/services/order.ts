import api from './api'

export interface OrderItem {
  product_id: string
  variant_id?: string
  quantity: number
}

export interface PlaceOrderRequest {
  items: OrderItem[]
}

export interface OrderItemResponse {
  product_id: string
  variant_id: string
  name: string
  quantity: number
  unit_price: number
  total_price: number
  image: string
}

export interface OrderResponse {
  id: string
  shop_id: string
  customer_id: string
  items: OrderItemResponse[]
  subtotal: number
  discount_total: number
  total: number
  status: string
  created_at: string
  updated_at: string
}

export const orderService = {
  async placeOrder(orderData: PlaceOrderRequest): Promise<OrderResponse> {
    const response = await api.post('/orders', orderData)
    return response.data
  },

  async getOrders(): Promise<OrderResponse[]> {
    const response = await api.get('/orders')
    return response.data
  },

  async getOrderDetail(orderId: string): Promise<OrderResponse> {
    const response = await api.get(`/orders/${orderId}`)
    return response.data
  },

  // Helper method to create order from cart items
  async createFromCartItems(items: OrderItem[]): Promise<OrderResponse> {
    const orderData: PlaceOrderRequest = {
      items: items.map(item => ({
        product_id: item.product_id,
        variant_id: item.variant_id || '',
        quantity: item.quantity
      }))
    }
    
    return this.placeOrder(orderData)
  }
}

export async function placeOrder(shopSlug: string, items: Array<{ product_id: string, variant_id: string, quantity: number }>) {
  return api.post(`/shops/${shopSlug}/orders`, { items });
}

export async function getOrderDetail(shopSlug: string, orderId: string) {
  return api.get(`/shops/${shopSlug}/orders/${orderId}`);
}

export async function getCustomerOrders(shopSlug: string) {
  return api.get(`/shops/${shopSlug}/orders`);
} 