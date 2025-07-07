import { defineStore } from 'pinia';
import * as cartApi from '@/services/cart';

export const useCartStore = defineStore('cart', {
  state: () => ({
    cart: null as any,
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
  },
}); 