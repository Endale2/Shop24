// src/services/shop.js
import api from './api';

export const shopService = {
  fetchShops: async () => {
    const res = await api.get('/seller/shops');
    return res.data.map(s => ({
      id: s._id || s.ID,
      name: s.Name,
      ownerId: s.OwnerID,
      description: s.Description,
      createdAt: s.CreatedAt,
      updatedAt: s.UpdatedAt,
    }));
  },

  createShop: async ({ name, description }) => {
    const res = await api.post('/seller/shops', { name, description });
    const newId = res.data.insertedId || res.data.InsertedID?.$oid;
    return { id: newId, name, description, ownerId: null, createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() };
  },
};