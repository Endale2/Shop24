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

// Helper function to safely map shop data
function mapShopData(s) {
  return {
    id:          s._id ?? s.id ?? s.ID,
    name:        s.Name ?? s.name ?? 'Unnamed Shop',
    slug:        s.Slug ?? s.slug ?? '',
    image:       s.Image ?? s.image ?? null,
    description: s.Description ?? s.description ?? '',
    createdAt:   s.CreatedAt ?? s.createdAt ?? new Date().toISOString(),
    updatedAt:   s.UpdatedAt ?? s.updatedAt ?? new Date().toISOString(),
    
    // Business Information
    category:    s.category ?? s.Category ?? '',
    email:       s.email ?? s.Email ?? '',
    phone:       s.phone ?? s.Phone ?? '',
    address:     s.address ?? s.Address ?? '',
    currency:    s.currency ?? s.Currency ?? 'USD',
    
    // Business Status
    status:      s.status ?? s.Status ?? 'active',
    isVerified:  s.isVerified ?? s.IsVerified ?? false,
    
    // Analytics & Metrics
    customer_count: s.customerCount ?? s.customer_count ?? 0,
    product_count:  s.productCount ?? s.product_count ?? 0,
    revenue:        s.revenue ?? s.Revenue ?? 0,
    rating:         s.rating ?? s.Rating ?? 0,
    review_count:   s.reviewCount ?? s.review_count ?? 0,
  }
}

export const shopService = {
  /**
   * Fetch all shops for the current seller.
   * Returns complete shop data with all fields
   */
  async fetchShops() {
    try {
      const res = await api.get('/seller/shops')
      return res.data.map(mapShopData)
    } catch (error) {
      console.error('Error fetching shops:', error)
      throw error
    }
  },

  /**
   * Create a new shop with all the required fields.
   * Payload includes all business information fields
   */
  async createShop({ name, description, category, logo, email, phone, address, currency }) {
    try {
      const payload = {
        name,
        slug: slugify(name),
        description,
        image: logo, // Map logo to image field
        category,
        email,
        phone,
        address,
        currency,
      }
      
      const res = await api.post('/seller/shops', payload)
      
      // grab new ID
      const newId = res.data.insertedId ?? res.data.InsertedID?.$oid
      
      return {
        id:          newId,
        name,
        slug:        payload.slug,
        image:       logo,
        description,
        category,
        email,
        phone,
        address,
        currency,
        status:      'active',
        isVerified:  false,
        ownerId:     null,
        createdAt:   new Date().toISOString(),
        updatedAt:   new Date().toISOString(),
        customer_count: 0,
        product_count: 0,
        revenue: 0,
        rating: 0,
        review_count: 0,
      }
    } catch (error) {
      console.error('Error creating shop:', error)
      throw error
    }
  },

  /**
   * Fetch a single shop by ID or slug (protected).
   */
  async fetchById(shopId) {
    try {
      const res = await api.get(`/seller/shops/${shopId}`)
      return mapShopData(res.data)
    } catch (error) {
      console.error('Error fetching shop by ID:', error)
      throw error
    }
  },

  /**
   * Update shop information
   */
  async updateShop(shopId, updates) {
    try {
      const res = await api.patch(`/seller/shops/${shopId}`, updates)
      return res.data
    } catch (error) {
      console.error('Error updating shop:', error)
      throw error
    }
  },

  /**
   * Delete a shop
   */
  async deleteShop(shopId) {
    try {
      const res = await api.delete(`/seller/shops/${shopId}`)
      return res.data
    } catch (error) {
      console.error('Error deleting shop:', error)
      throw error
    }
  },
}
