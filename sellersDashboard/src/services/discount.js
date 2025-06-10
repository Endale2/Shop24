// src/services/discount.js
import api from './api'

export const discountService = {
  /**
   * Fetch all discounts for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts`)
    // Assume res.data is an array of discount objects
    return res.data.map(d => ({
      id:               d.id,
      name:             d.name,
      description:      d.description,
      category:         d.category,           // e.g. "product" or "order"
      type:             d.type,               // e.g. "percentage" or "fixed_amount"
      value:            d.value,              // numeric
      shopId:           d.shop_id,
      sellerId:         d.seller_id,
      appliesToProducts: d.applies_to_products ?? [], // array of product IDs
      couponCode:       d.coupon_code,
      startAt:          d.startAt || d.start_at,    // depending on backend naming
      endAt:            d.endAt   || d.end_at,
      active:           d.active,
      createdAt:        d.created_at,
      updatedAt:        d.updated_at,
    }))
  },

  /**
   * Fetch a single discount by ID.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, discountId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts/${discountId}`)
    const d = res.data
    return {
      id:               d.id,
      name:             d.name,
      description:      d.description,
      category:         d.category,
      type:             d.type,
      value:            d.value,
      shopId:           shopId,
      sellerId:         d.seller_id,
      appliesToProducts: d.applies_to_products ?? [],
      couponCode:       d.coupon_code,
      startAt:          d.startAt || d.start_at,
      endAt:            d.endAt   || d.end_at,
      active:           d.active,
      createdAt:        d.created_at,
      updatedAt:        d.updated_at,
    }
  },

  /**
   * Create a new discount.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.name
   * @param {string} [data.description]
   * @param {string} data.category
   * @param {string} data.type
   * @param {number} data.value
   * @param {Array<string>} [data.appliesToProducts]
   * @param {string} [data.couponCode]
   * @param {string} [data.startAt] ISO string
   * @param {string} [data.endAt] ISO string
   * @param {boolean} [data.active]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const payload = {
      name:               data.name,
      description:        data.description,
      category:           data.category,
      type:               data.type,
      value:              data.value,
      applies_to_products: data.appliesToProducts,
      coupon_code:        data.couponCode,
      startAt:            data.startAt,
      endAt:              data.endAt,
      active:             data.active,
    }
    const res = await api.post(`/seller/shops/${shopId}/discounts`, payload)
    const newId = res.data.id ?? res.data.insertedId
    return this.fetchById(shopId, newId)
  },

  /**
   * Update an existing discount.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, discountId, data) {
    const payload = {}
    if (data.name               !== undefined) payload.name               = data.name
    if (data.description        !== undefined) payload.description        = data.description
    if (data.category           !== undefined) payload.category           = data.category
    if (data.type               !== undefined) payload.type               = data.type
    if (data.value              !== undefined) payload.value              = data.value
    if (data.appliesToProducts  !== undefined) payload.applies_to_products = data.appliesToProducts
    if (data.couponCode         !== undefined) payload.coupon_code        = data.couponCode
    if (data.startAt            !== undefined) payload.startAt            = data.startAt
    if (data.endAt              !== undefined) payload.endAt              = data.endAt
    if (data.active             !== undefined) payload.active             = data.active

    await api.put(`/seller/shops/${shopId}/discounts/${discountId}`, payload)
    return this.fetchById(shopId, discountId)
  },

  /**
   * Delete a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<void>}
   */
  async delete(shopId, discountId) {
    await api.delete(`/seller/shops/${shopId}/discounts/${discountId}`)
  }
}
