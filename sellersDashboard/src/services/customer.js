// src/services/customer.js
import api from './api'

export const customerService = {
  /**
   * Fetch all customers for a given shop, paginated and filtered by search/segment.
   * @param {string} shopId
   * @param {Object} options
   * @param {number} options.page
   * @param {number} options.limit
   * @param {string} options.search
   * @param {string} options.segmentId
   * @returns {Promise<Object>} { customers, total, page, pageSize }
   */
  async fetchAll(shopId, { page = 1, limit = 10, search = '', segmentId = '' } = {}) {
    if (!shopId) return { customers: [], total: 0, page: 1, pageSize: 10 };
    try {
      const params = []
      if (page) params.push(`page=${page}`)
      if (limit) params.push(`limit=${limit}`)
      if (search) params.push(`search=${encodeURIComponent(search)}`)
      if (segmentId) params.push(`segmentId=${segmentId}`)
      const query = params.length ? `?${params.join('&')}` : ''
      const res = await api.get(`/seller/shops/${shopId}/customers${query}`)
      if (!res.data || !Array.isArray(res.data.customers)) {
        return { customers: [], total: 0, page: 1, pageSize: 10 };
      }
      // Normalize list entries
      const customers = res.data.customers.map(item => ({
        id:           item.customer.id,
        firstName:    item.customer.firstName,
        lastName:     item.customer.lastName,
        email:        item.customer.email,
        phone:        item.customer.phone,
        address:      item.customer.address,
        city:         item.customer.city,
        state:        item.customer.state,
        postalCode:   item.customer.postalCode,
        country:      item.customer.country,
        createdAt:    item.customer.createdAt,
        updatedAt:    item.customer.updatedAt,
        linkId:       item.linkId ?? null,
        profile_image: item.customer.profile_image
      }))
      return {
        customers,
        total: res.data.total,
        page: res.data.page,
        pageSize: res.data.pageSize
      }
    } catch (err) {
      // Optionally log error
      return { customers: [], total: 0, page: 1, pageSize: 10 };
    }
  },

  /**
   * Fetch customer dashboard statistics.
   * @param {string} shopId
   * @returns {Promise<Object>} { total, segments, thisMonth, segmented }
   */
  async fetchStats(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/customers/stats`)
    return res.data
  },

  /**
   * Fetch a single customer by ID.
   * The endpoint may return { customer: {...}, segments: [...], linkId } or the raw customer object.
   * @param {string} shopId
   * @param {string} customerId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, customerId) {
    const res = await api.get(`/seller/shops/${shopId}/customers/${customerId}`)
    const c = res.data.customer
    return {
      id: c.id,
      firstName: c.firstName,
      lastName: c.lastName,
      email: c.email,
      phone: c.phone,
      address: c.address,
      city: c.city,
      state: c.state,
      postalCode: c.postalCode,
      country: c.country,
      createdAt: c.createdAt,
      updatedAt: c.updatedAt,
      linkId: res.data.linkId ?? null,
      segments: res.data.segments ?? [],
      profile_image: c.profile_image // Add profile_image
    }
  },

  /**
   * Create a new customer in a shop.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.firstName
   * @param {string} data.lastName
   * @param {string} data.email
   * @param {string} [data.phone]
   * @param {string} [data.address]
   * @param {string} [data.city]
   * @param {string} [data.state]
   * @param {string} [data.postalCode]
   * @param {string} [data.country]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/customers`, data)
    return this.fetchById(shopId, res.data.id || res.data.insertedId)
  },

  /**
   * Update an existing customer.
   * @param {string} shopId
   * @param {string} customerId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, customerId, data) {
    await api.put(`/seller/shops/${shopId}/customers/${customerId}`, data)
    return this.fetchById(shopId, customerId)
  },

  /**
   * Unlink a customer from a shop (doesn't delete customer data).
   * @param {string} shopId
   * @param {string} linkId - The link ID from the customer data
   * @returns {Promise<void>}
   */
  async delete(shopId, linkId) {
    await api.delete(`/seller/shops/${shopId}/customers/link/${linkId}`)
  },

  /**
   * Fetch order history and stats for a customer in a shop.
   * @param {string} shopId
   * @param {string} customerId
   * @returns {Promise<Object>} - { orders, order_count, total_spend, last_order }
   */
  async fetchHistory(shopId, customerId) {
    const res = await api.get(`/seller/shops/${shopId}/customers/${customerId}/history`)
    return res.data
  }
}