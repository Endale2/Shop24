// src/services/shop.js
import api from './api';

export default {
  // GET /seller/shops
  async fetchShops() {
    const res = await api.get('/seller/shops');
    return res.data; // [ { id, name, ownerId, … }, … ]
  },

  // POST /seller/shops
  async createShop({ name, description }) {
    const res = await api.post('/seller/shops', { name, description });
    return res.data; // InsertOneResult
  },

  // Optionally, select a shop in your Vuex store
  selectShop(shop) {
    // dispatch to shops module: e.g. store.commit('shops/setActiveShop', shop)
  },

  // GET a single shop
  async getShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}`);
    return res.data;
  },

  // PATCH update shop
  async updateShop(shopId, updates) {
    const res = await api.patch(`/seller/shops/${shopId}`, updates);
    return res.data;
  },

  // DELETE shop
  async deleteShop(shopId) {
    const res = await api.delete(`/seller/shops/${shopId}`);
    return res.data;
  },
};
