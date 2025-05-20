// src/services/shop.js
import api from './api'; // your axios instance

export default {
  // GET /seller/shops → normalize keys
  async fetchShops() {
    const res = await api.get('/seller/shops');
    return res.data.map(s => ({
      id:          s.ID,
      name:        s.Name,
      ownerId:     s.OwnerID,
      description: s.Description,
      createdAt:   s.CreatedAt,
      updatedAt:   s.UpdatedAt
    }));
  },

  // POST /seller/shops
  async createShop({ name, description }) {
    const res = await api.post('/seller/shops', { name, description });
    // backend returns InsertOneResult; we need to re-fetch or construct:
    const newId = res.data.InsertedID.$oid; // adjust to however your backend returns it
    return { id: newId, name, ownerId: null, description, createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() };
  }
};
