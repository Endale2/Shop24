// src/stores/cart.js
import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: []  // Each item: { productId, name, variantOptions, price, quantity, image }
  }),
  getters: {
    totalItems: (state) => state.items.reduce((sum, i) => sum + i.quantity, 0),
    totalPrice: (state) => state.items.reduce((sum, i) => sum + i.price * i.quantity, 0)
  },
  actions: {
    /**
     * Add an item to the cart. If same product+variant exists, increment quantity.
     */
    addItem(item) {
      const existing = this.items.find(i =>
        i.productId === item.productId &&
        JSON.stringify(i.variantOptions) === JSON.stringify(item.variantOptions)
      )

      if (existing) {
        existing.quantity += item.quantity
      } else {
        this.items.push({ ...item })
      }
    },

    /**
     * Remove an item completely or decrement quantity.
     */
    removeItem(productId, variantOptions, qty = null) {
      const index = this.items.findIndex(i =>
        i.productId === productId &&
        JSON.stringify(i.variantOptions) === JSON.stringify(variantOptions)
      )
      if (index === -1) return

      if (qty === null || this.items[index].quantity <= qty) {
        this.items.splice(index, 1)
      } else {
        this.items[index].quantity -= qty
      }
    },

    /**
     * Clear the entire cart.
     */
    clearCart() {
      this.items = []
    }
  },
  persist: true
})
