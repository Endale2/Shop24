// src/services/customer.js
import api from './api'

export const customerService = {
  /**
   * Fetch all customers for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAll(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/customers`)
    // normalize list entries
    return res.data.map(c => ({
      id:         c.id,
      firstName:  c.firstName,
      lastName:   c.lastName,
      username:   c.username,
      email:      c.email,
      phone:      c.phone,
      address:    c.address,
      city:       c.city,
      state:      c.state,
      postalCode: c.postalCode,
      country:    c.country,
      createdAt:  c.createdAt,
      updatedAt:  c.updatedAt,
      linkId:     c.linkId ?? null
    }))
  },

  /**
   * Fetch a single customer by ID.
   * The endpoint may return { customer: {...}, linkId } or the raw customer object.
   * @param {string} shopId
   * @param {string} customerId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, customerId) {
    const res = await api.get(`/seller/shops/${shopId}/customers/${customerId}`)
    // unwrap payload
    const payload = res.data.customer ?? res.data
    return {
      id:          payload.id,
      firstName:   payload.firstName,
      lastName:    payload.lastName,
      username:    payload.username,
      email:       payload.email,
      phone:       payload.phone,
      address:     payload.address,
      city:        payload.city,
      state:       payload.state,
      postalCode:  payload.postalCode,
      country:     payload.country,
      createdAt:   payload.createdAt,
      updatedAt:   payload.updatedAt,
      linkId:      res.data.linkId ?? null
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
   * Delete a customer.
   * @param {string} shopId
   * @param {string} customerId
   * @returns {Promise<void>}
   */
  async delete(shopId, customerId) {
    await api.delete(`/seller/shops/${shopId}/customers/${customerId}`)
  }
}
