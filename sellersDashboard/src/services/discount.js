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
      currentUsage: d.current_usage ?? 0,
      usageTracking: d.usage_tracking ?? [],
      eligibilityType: d.eligibility_type ?? 'all',
      allowedCustomers: d.allowed_customers ?? [],
      allowedSegments: d.allowed_segments ?? [],
      buyProductIds: d.buy_product_ids ?? [],
      buyQuantity: d.buy_quantity,
      getProductIds: d.get_product_ids ?? [],
      getQuantity: d.get_quantity,
      appliesToVariants: d.applies_to_variants ?? [],
    };
  },

  /**
   * Convert datetime-local format to RFC3339 for backend
   * @param {string} datetimeLocal - datetime-local format string
   * @returns {string} RFC3339 format string
   */
  _convertToRFC3339(datetimeLocal) {
    if (!datetimeLocal) return undefined;
    // datetime-local format: "2024-01-15T10:30"
    // RFC3339 format: "2024-01-15T10:30:00Z"
    return new Date(datetimeLocal).toISOString();
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
   * @param {string} [data.startAt] datetime-local format
   * @param {string} [data.endAt] datetime-local format
   * @param {boolean} [data.active]
   * @param {string} [data.eligibilityType]
   * @param {Array<string>} [data.allowedCustomers]
   * @param {Array<string>} [data.allowedSegments]
   * @param {number} [data.usageLimit]
   * @param {number} [data.perCustomerLimit]
   * @param {Array<string>} [data.buyProductIds]
   * @param {number} [data.buyQuantity]
   * @param {Array<string>} [data.getProductIds]
   * @param {number} [data.getQuantity]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const payload = {
      name:                 data.name,
      description:          data.description,
      category:             data.category,
      type:                 data.type,
      value:                data.value,
      appliesToProducts:    data.appliesToProducts,   // Sending IDs
      appliesToCollections: data.appliesToCollections, // Sending IDs
      appliesToVariants:    data.appliesToVariants,
      couponCode:           data.couponCode,
      startAt:              this._convertToRFC3339(data.startAt),
      endAt:                this._convertToRFC3339(data.endAt),
      active:               data.active,
      freeShipping:         data.freeShipping,
      minimumOrderSubtotal: data.minimumOrderSubtotal,
      minimumOrderForFreeShipping: data.minimumFreeShipping,
      usageLimit:           data.usageLimit,
      perCustomerLimit:     data.perCustomerLimit,
      eligibilityType:      data.eligibilityType,
      allowedCustomers:     data.allowedCustomers,
      allowedSegments:      data.allowedSegments,
      buyProductIds:        data.buyProductIds,
      buyQuantity:          data.buyQuantity,
      getProductIds:        data.getProductIds,
      getQuantity:          data.getQuantity,
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
    if (data.appliesToProducts !== undefined) payload.appliesToProducts = data.appliesToProducts; // Sending IDs
    if (data.appliesToCollections !== undefined) payload.appliesToCollections = data.appliesToCollections; // Sending IDs
    if (data.appliesToVariants !== undefined) payload.appliesToVariants = data.appliesToVariants;
    if (data.couponCode !== undefined) payload.couponCode = data.couponCode;
    if (data.startAt !== undefined) payload.startAt = this._convertToRFC3339(data.startAt);
    if (data.endAt !== undefined) payload.endAt = this._convertToRFC3339(data.endAt);
    if (data.active !== undefined) payload.active = data.active;
    if (data.freeShipping !== undefined) payload.freeShipping = data.freeShipping;
    if (data.minimumOrderSubtotal !== undefined) payload.minimumOrderSubtotal = data.minimumOrderSubtotal;
    if (data.minimumFreeShipping !== undefined) payload.minimumOrderForFreeShipping = data.minimumFreeShipping;
    if (data.usageLimit !== undefined) payload.usageLimit = data.usageLimit;
    if (data.perCustomerLimit !== undefined) payload.perCustomerLimit = data.perCustomerLimit;
    if (data.eligibilityType !== undefined) payload.eligibilityType = data.eligibilityType;
    if (data.allowedCustomers !== undefined) payload.allowedCustomers = data.allowedCustomers;
    if (data.allowedSegments !== undefined) payload.allowedSegments = data.allowedSegments;
    if (data.buyProductIds !== undefined) payload.buyProductIds = data.buyProductIds;
    if (data.buyQuantity !== undefined) payload.buyQuantity = data.buyQuantity;
    if (data.getProductIds !== undefined) payload.getProductIds = data.getProductIds;
    if (data.getQuantity !== undefined) payload.getQuantity = data.getQuantity;

    await api.patch(`/seller/shops/${shopId}/discounts/${discountId}`, payload);
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
  },

  /**
   * Add customers to a discount's allowed list.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Array<string>} customerIds
   * @returns {Promise<Object>}
   */
  async addCustomers(shopId, discountId, customerIds) {
    const res = await api.post(`/seller/shops/${shopId}/discounts/${discountId}/customers`, {
      customerIds: customerIds
    });
    return res.data;
  },

  /**
   * Add customer segments to a discount's allowed list.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Array<string>} segmentIds
   * @returns {Promise<Object>}
   */
  async addSegments(shopId, discountId, segmentIds) {
    const res = await api.post(`/seller/shops/${shopId}/discounts/${discountId}/segments`, {
      segmentIds: segmentIds
    });
    return res.data;
  },

  /**
   * Get usage statistics for a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<Object>}
   */
  async getUsageStats(shopId, discountId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts/${discountId}/usage`);
    return res.data;
  },

  /**
   * Validate if a customer can use a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @param {string} customerId
   * @param {Array<string>} customerSegmentIds
   * @returns {Promise<Object>}
   */
  async validateForCustomer(shopId, discountId, customerId, customerSegmentIds = []) {
    const res = await api.post(`/seller/shops/${shopId}/discounts/${discountId}/validate`, {
      customerId: customerId,
      customerSegmentIds: customerSegmentIds
    });
    return res.data;
  },

  /**
   * Get eligible discounts for a customer.
   * @param {string} shopId
   * @param {string} customerId
   * @param {Array<string>} customerSegmentIds
   * @returns {Promise<Array<Object>>}
   */
  async getEligibleForCustomer(shopId, customerId, customerSegmentIds = []) {
    // This would need a new endpoint, but for now we'll filter client-side
    const allDiscounts = await this.fetchAllByShop(shopId);
    const eligibleDiscounts = [];
    
    for (const discount of allDiscounts) {
      try {
        const validation = await this.validateForCustomer(shopId, discount.id, customerId, customerSegmentIds);
        if (validation.valid) {
          eligibleDiscounts.push(discount);
        }
      } catch (error) {
        // Skip invalid discounts
        continue;
      }
    }
    
    return eligibleDiscounts;
  }
};