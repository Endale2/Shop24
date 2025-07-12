import { defineStore } from 'pinia';
import * as cartApi from '@/services/cart';
import * as orderApi from '@/services/order';
import type { CartWithDiscountDetails } from '@/services/cart';

export const useCartStore = defineStore('cart', {
  state: () => ({
    cart: null as CartWithDiscountDetails | null,
    loading: false,
    error: null as null | string,
    shopSlug: '', // Set this when user selects a shop
  }),
  actions: {
    setShopSlug(slug: string) {
      this.shopSlug = slug;
    },
    async fetchCart() {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.getCart(this.shopSlug);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to fetch cart';
        // Clear cart if unauthorized (user not logged in)
        if (e.response?.status === 401) {
          this.cart = null;
        }
      } finally {
        this.loading = false;
      }
    },
    async addToCart(productId: string, variantId: string, quantity: number) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.addToCart(this.shopSlug, productId, variantId, quantity);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to add to cart';
      } finally {
        this.loading = false;
      }
    },
    async updateCartItem(productId: string, variantId: string, quantity: number) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.updateCartItem(this.shopSlug, productId, variantId, quantity);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to update cart item';
      } finally {
        this.loading = false;
      }
    },
    async removeCartItem(productId: string, variantId: string) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.removeCartItem(this.shopSlug, productId, variantId);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to remove cart item';
      } finally {
        this.loading = false;
      }
    },
    async applyDiscount(code: string) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.applyDiscount(this.shopSlug, code);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to apply discount';
      } finally {
        this.loading = false;
      }
    },
    async removeDiscount(discountId: string) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.removeDiscount(this.shopSlug, discountId);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to remove discount';
      } finally {
        this.loading = false;
      }
    },
    async clearCart() {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.clearCart(this.shopSlug);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to clear cart';
      } finally {
        this.loading = false;
      }
    },
    // Method to place order from cart - sends only minimal data to backend
    async placeOrder() {
      if (!this.shopSlug || !this.cart?.items || this.cart.items.length === 0) {
        throw new Error('No items in cart to place order');
      }
      
      this.loading = true;
      this.error = null;
      try {
        // Send only minimal data to backend - NO pricing information
        // Backend will calculate all pricing, discounts, and totals securely
        const orderItems = this.cart.items.map(item => ({
          product_id: item.product_id,
          variant_id: item.variant_id || '',
          quantity: item.quantity
          // Note: NO pricing data sent - backend calculates everything
        }));
        
        // Call the order service to place the order
        // Backend will:
        // 1. Fetch product prices from database
        // 2. Calculate all discounts server-side
        // 3. Apply the best eligible discounts
        // 4. Calculate final totals
        // 5. Return the order with server-calculated pricing
        const response = await orderApi.placeOrder(this.shopSlug, orderItems);
        return response.data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to place order';
        throw e;
      } finally {
        this.loading = false;
      }
    },
    // Method to refresh cart when user logs in
    async refreshCart() {
      if (this.shopSlug) {
        await this.fetchCart();
      }
    },
    // Method to clear cart state when user logs out
    clearCartState() {
      this.cart = null;
      this.error = null;
    },
  },
  persist: {
    paths: ['cart', 'shopSlug'],
  },
}); 