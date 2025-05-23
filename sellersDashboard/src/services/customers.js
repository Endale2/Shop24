// src/services/customer.js
import api from './api';

export const customerService = {
  /**
   * Fetch all customers for a given shop
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/customers`);
    // res.data is array of { customer: {...}, linkId }
    return res.data.map(item => ({
      id:          item.customer.id,
      firstName:   item.customer.firstName,
      lastName:    item.customer.lastName,
      username:    item.customer.username,
      email:       item.customer.email,
      phone:       item.customer.phone,
      address:     item.customer.address,
      city:        item.customer.city,
      state:       item.customer.state,
      postalCode:  item.customer.postalCode,
      country:     item.customer.country,
      createdAt:   item.customer.createdAt,
      updatedAt:   item.customer.updatedAt,
      linkId:      item.linkId
    }));
  },

  /**
   * Fetch a single customer by ID
   * @param {string} shopId
   * @param {string} customerId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, customerId) {
    const res = await api.get(
      `/seller/shops/${shopId}/customers/${customerId}`
    );
    // res.data is { customer: {...}, linkId }
    const c = res.data.customer;
    return {
      id:          c.id,
      firstName:   c.firstName,
      lastName:    c.lastName,
      username:    c.username,
      email:       c.email,
      phone:       c.phone,
      address:     c.address,
      city:        c.city,
      state:       c.state,
      postalCode:  c.postalCode,
      country:     c.country,
      createdAt:   c.createdAt,
      updatedAt:   c.updatedAt,
      linkId:      res.data.linkId
    };
  }
};
