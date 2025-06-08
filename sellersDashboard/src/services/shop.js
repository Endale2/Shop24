import api from './api'

// Simple slugify: lowercase, spaces → dash, alphanumerics only
function slugify(str) {
  return str
    .toString()
    .toLowerCase()
    .trim()
    .replace(/\s+/g, '-')       // spaces → dash
    .replace(/[^\w\-]+/g, '')   // remove non-word chars
    .replace(/\-\-+/g, '-')     // collapse multiple dashes
}

export const shopService = {
  /**
   * Fetch all shops for the current seller.
   * Returns { id, name, slug, image, description, createdAt, updatedAt }
   */
  async fetchShops() {
    const res = await api.get('/seller/shops')
    return res.data.map(s => ({
      id:          s._id ?? s.id,
      name:        s.Name ?? s.name,
      slug:        s.Slug ?? s.slug,
      image:       s.Image ?? s.image ?? null,
      description: s.Description ?? s.description ?? '',
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }))
  },

  /**
   * Create a new shop.
   * Payload now includes `slug` and optional `image`.
   */
  async createShop({ name, description, image = '' }) {
    const payload = {
      name,
      slug: slugify(name),
      description,
      image,
    }
    const res = await api.post('/seller/shops', payload)
    // grab new ID
    const newId = res.data.insertedId ?? res.data.InsertedID?.$oid
    return {
      id:          newId,
      name,
      slug:        payload.slug,
      image,
      description,
      ownerId:     null,
      createdAt:   new Date().toISOString(),
      updatedAt:   new Date().toISOString(),
    }
  },

  /**
   * Fetch a single shop by ID or slug (protected).
   */
  async fetchById(shopId) {
    const res = await api.get(`/seller/shops/${shopId}`)
    const s = res.data
    return {
      id:          s._id ?? s.ID,
      name:        s.Name ?? s.name,
      slug:        s.Slug ?? s.slug,
      image:       s.Image ?? s.image ?? null,
      description: s.Description ?? s.description ?? '',
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }
  },
}
