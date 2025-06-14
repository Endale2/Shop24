// src/services/discount.js
import api from './api';
// No need to import productService here if the backend hydrates applies_to_products

export const discountService = {
  /**
   * Helper to map raw API discount object to camelCase and ensure expected formats.
   * This handles the backend already hydrating `applies_to_products` with full data.
   * @param {Object} d - Raw discount object from API
   * @returns {Object} Mapped discount object
   */
  _mapDiscount(d) {
    return {
      id:                 d.id,
      name:               d.name,
      description:        d.description,
      category:           d.category,
      type:               d.type,
      value:              d.value,
      shopId:             d.shop_id,
      sellerId:           d.seller_id,
      couponCode:         d.coupon_code,
      startAt:            d.start_at, // Directly use as ISO string
      endAt:              d.end_at,   // Directly use as ISO string
      active:             d.active,
      createdAt:          d.created_at,
      updatedAt:          d.updated_at,
      // Backend provides full product objects for applies_to_products, so just map it
      appliesToProducts:  d.applies_to_products ?? [],
      // Backend provides collection IDs for applies_to_collections
      appliesToCollections: d.applies_to_collections ?? [],
      freeShipping: d.free_shipping,
      minimumOrderSubtotal: d.minimum_order_subtotal,
      minimumFreeShipping: d.minimum_free_shipping,
      usageLimit: d.usage_limit,
      perCustomerLimit: d.per_customer_limit,
      allowedCustomers: d.allowed_customers ?? [],
      buyProductIds: d.buy_product_ids ?? [],
      buyQuantity: d.buy_quantity,
      getProductIds: d.get_product_ids ?? [],
      getQuantity: d.get_quantity,
      appliesToVariants: d.applies_to_variants ?? [],
    };
  },

  /**
   * Fetch all discounts for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts`);
    // Map each raw discount object using the helper
    return res.data.map(this._mapDiscount);
  },

  /**
   * Fetch a single discount by ID.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, discountId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts/${discountId}`);
    // Map the single raw discount object using the helper
    return this._mapDiscount(res.data);
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
   * @param {Array<string>} [data.appliesToProducts] // Expects product IDs
   * @param {Array<string>} [data.appliesToCollections] // Expects collection IDs
   * @param {string} [data.couponCode]
   * @param {string} [data.startAt] ISO string
   * @param {string} [data.endAt] ISO string
   * @param {boolean} [data.active]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const payload = {
      name:                 data.name,
      description:          data.description,
      category:             data.category,
      type:                 data.type,
      value:                data.value,
      applies_to_products:  data.appliesToProducts,   // Sending IDs
      applies_to_collections: data.appliesToCollections, // Sending IDs
      coupon_code:          data.couponCode,
      start_at:             data.startAt,
      end_at:               data.endAt,
      active:               data.active,
      free_shipping:        data.freeShipping,
      minimum_order_subtotal: data.minimumOrderSubtotal,
      minimum_free_shipping: data.minimumFreeShipping,
      usage_limit:          data.usageLimit,
      per_customer_limit:   data.perCustomerLimit,
      allowed_customers:    data.allowedCustomers,
      buy_product_ids:      data.buyProductIds,
      buy_quantity:         data.buyQuantity,
      get_product_ids:      data.getProductIds,
      get_quantity:         data.getQuantity,
      applies_to_variants:  data.appliesToVariants,
    };
    // Clean up undefined fields before sending
    Object.keys(payload).forEach(key => payload[key] === undefined && delete payload[key]);

    const res = await api.post(`/seller/shops/${shopId}/discounts`, payload);
    const newId = res.data.id ?? res.data.insertedId; // Handle potential different backend responses for ID
    return this.fetchById(shopId, newId); // Fetch the newly created discount for consistent data structure
  },

  /**
   * Update an existing discount.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, discountId, data) {
    const payload = {};
    if (data.name !== undefined) payload.name = data.name;
    if (data.description !== undefined) payload.description = data.description;
    if (data.category !== undefined) payload.category = data.category;
    if (data.type !== undefined) payload.type = data.type;
    if (data.value !== undefined) payload.value = data.value;
    if (data.appliesToProducts !== undefined) payload.applies_to_products = data.appliesToProducts; // Sending IDs
    if (data.appliesToCollections !== undefined) payload.applies_to_collections = data.appliesToCollections; // Sending IDs
    if (data.couponCode !== undefined) payload.coupon_code = data.couponCode;
    if (data.startAt !== undefined) payload.start_at = data.startAt;
    if (data.endAt !== undefined) payload.end_at = data.endAt;
    if (data.active !== undefined) payload.active = data.active;
    if (data.freeShipping !== undefined) payload.free_shipping = data.freeShipping;
    if (data.minimumOrderSubtotal !== undefined) payload.minimum_order_subtotal = data.minimumOrderSubtotal;
    if (data.minimumFreeShipping !== undefined) payload.minimum_free_shipping = data.minimumFreeShipping;
    if (data.usageLimit !== undefined) payload.usage_limit = data.usageLimit;
    if (data.perCustomerLimit !== undefined) payload.per_customer_limit = data.perCustomerLimit;
    if (data.allowedCustomers !== undefined) payload.allowed_customers = data.allowedCustomers;
    if (data.buyProductIds !== undefined) payload.buy_product_ids = data.buyProductIds;
    if (data.buyQuantity !== undefined) payload.buy_quantity = data.buyQuantity;
    if (data.getProductIds !== undefined) payload.get_product_ids = data.getProductIds;
    if (data.getQuantity !== undefined) payload.get_quantity = data.getQuantity;
    if (data.appliesToVariants !== undefined) payload.applies_to_variants = data.appliesToVariants;


    await api.put(`/seller/shops/${shopId}/discounts/${discountId}`, payload);
    return this.fetchById(shopId, discountId); // Fetch updated discount for consistent data structure
  },

  /**
   * Delete a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<void>}
   */
  async delete(shopId, discountId) {
    await api.delete(`/seller/shops/${shopId}/discounts/${discountId}`);
  }
};