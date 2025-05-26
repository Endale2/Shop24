import api from './api'

export const shopService = {
  async fetchShops() {
    const res = await api.get('/seller/shops')
    return res.data.map(s => ({
      id:          s._id ?? s.id,
      name:        s.Name ?? s.name,
      ownerId:     s.OwnerID ?? s.ownerId,
      description: s.Description ?? s.description,
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }))
  },
  async createShop({ name, description }) {
    const res = await api.post('/seller/shops', { name, description })
    const newId = res.data.insertedId ?? res.data.InsertedID?.$oid
    return { id: newId, name, description, ownerId: null, createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() }
  },
  async fetchById(shopId) {
    // if you have a protected GET /seller/shops/:id
    const res = await api.get(`/seller/shops/${shopId}`)
    const s = res.data
    return {
      id:          s._id ?? s.ID,
      name:        s.Name ?? s.name,
      ownerId:     s.OwnerID ?? s.ownerId,
      description: s.Description ?? s.description,
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }
  },
}
