import { defineStore } from 'pinia';
import { getWishlist, addToWishlist, removeFromWishlist } from '@/services/wishlist';
import { fetchProductDetailById } from '@/services/product';

export const useWishlistStore = defineStore('wishlist', {
  state: () => ({
    productIds: [] as string[],
    products: [] as any[],
    loading: false,
    error: null as string | null,
    shopSlug: '',
  }),
  actions: {
    setShopSlug(slug: string) {
      this.shopSlug = slug;
    },
    async fetchWishlist() {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        const { data } = await getWishlist(this.shopSlug);
        this.productIds = data.product_ids || [];
        // Fetch product details for each product ID
        this.products = [];
        for (const id of this.productIds) {
          try {
            const { data: product } = await fetchProductDetailById(this.shopSlug, id);
            if (product) this.products.push(product);
          } catch (e) {
            // Ignore missing products
          }
        }
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to fetch wishlist';
      } finally {
        this.loading = false;
      }
    },
    async addProduct(productId: string) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        await addToWishlist(this.shopSlug, productId);
        await this.fetchWishlist();
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to add to wishlist';
      } finally {
        this.loading = false;
      }
    },
    async removeProduct(productId: string) {
      if (!this.shopSlug) return;
      this.loading = true;
      this.error = null;
      try {
        await removeFromWishlist(this.shopSlug, productId);
        await this.fetchWishlist();
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to remove from wishlist';
      } finally {
        this.loading = false;
      }
    },
    clearWishlistState() {
      this.productIds = [];
      this.products = [];
      this.error = null;
    },
  },
  persist: {
    paths: ['productIds', 'shopSlug'],
  },
}); 