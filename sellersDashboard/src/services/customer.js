// src/services/customer.js
import api from './api';

export const customerService = {
  // …

  async fetchById(shopId, customerId) {
    const res = await api.get(
      `/seller/shops/${shopId}/customers/${customerId}`
    );
    // The endpoint might return { customer: {...}, linkId } OR return the customer object directly.
    const payload = res.data.customer ? res.data.customer : res.data;
    const linkId  = res.data.linkId ?? null;
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
      linkId
    };
  }
};
