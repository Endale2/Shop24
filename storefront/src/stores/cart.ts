import { defineStore } from 'pinia';
import * as cartApi from '@/services/cart';
import * as orderApi from '@/services/order';
import type { CartWithDiscountDetails } from '@/services/cart';
import { getCurrentShopSlug } from '@/services/shop';

export const useCartStore = defineStore('cart', {
  state: () => ({
    cart: null as CartWithDiscountDetails | null,
    loading: false,
    error: null as null | string,
    shopSlug: getCurrentShopSlug() || '', // Default to subdomain
  }),
  getters: {
    itemCount: (state) => {
      return state.cart?.items?.reduce((total, item) => total + item.quantity, 0) || 0;
    },
    totalPrice: (state) => {
      return state.cart?.grand_total || 0;
    },
    subtotal: (state) => {
      return state.cart?.subtotal || 0;
    },
    totalDiscount: (state) => {
      return state.cart?.total_discounts || 0;
    },
    hasItems: (state) => {
      return (state.cart?.items?.length || 0) > 0;
    },
    isEmpty: (state) => {
      return (state.cart?.items?.length || 0) === 0;
    }
  },
  actions: {
    setShopSlug(slug?: string) {
      this.shopSlug = slug || getCurrentShopSlug() || '';
    },
    async fetchCart() {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.getCart(slug);
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
    async addToCart(productData: any, variantId: string, quantity: number) {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      // Ensure product ID is a string, not an object
      const productId = typeof productData === 'object' && productData !== null ? productData.id : productData;
      try {
        const { data } = await cartApi.addToCart(slug, productId, variantId, quantity);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to add to cart';
      } finally {
        this.loading = false;
      }
    },
    async updateCartItem(productId: string, variantId: string, quantity: number) {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.updateCartItem(slug, productId, variantId, quantity);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to update cart item';
      } finally {
        this.loading = false;
      }
    },
    async removeCartItem(productId: string, variantId: string) {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.removeCartItem(slug, productId, variantId);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to remove cart item';
      } finally {
        this.loading = false;
      }
    },
    async applyDiscount(code: string) {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.applyDiscount(slug, code);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to apply discount';
      } finally {
        this.loading = false;
      }
    },
    async removeDiscount(discountId: string) {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.removeDiscount(slug, discountId);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to remove discount';
      } finally {
        this.loading = false;
      }
    },
    async clearCart() {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await cartApi.clearCart(slug);
        this.cart = data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to clear cart';
      } finally {
        this.loading = false;
      }
    },
    async placeOrder() {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (!slug || !this.cart?.items || this.cart.items.length === 0) {
        throw new Error('No items in cart to place order');
      }
      this.loading = true;
      this.error = null;
      try {
        const orderItems = this.cart.items.map(item => ({
          product_id: item.product_id,
          variant_id: item.variant_id || '',
          quantity: item.quantity
        }));
        const response = await orderApi.placeOrder(slug, orderItems);
        return response.data;
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to place order';
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async refreshCart() {
      const slug = this.shopSlug || getCurrentShopSlug();
      if (slug) {
        await this.fetchCart();
      }
    },
    clearCartState() {
      this.cart = null;
      this.error = null;
    },
  },
}); 