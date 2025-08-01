// src/services/discount.js
import api from './api';

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
      category:           d.category, // Only 'product' is supported
      type:               d.type,
      value:              d.value,
      couponCode:         d.coupon_code,
      startAt:            d.start_at,
      endAt:              d.end_at,
      active:             d.active,
      createdAt:          d.created_at,
      updatedAt:          d.updated_at,
      // Product-level targets only - map hydrated products with full data
      appliesToProducts:  (d.applies_to_products ?? []).map(p => ({
        id: p.id,
        name: p.name,
        description: p.description,
        mainImage: p.main_image,
        price: p.price,
        stock: p.stock,
      })),
      appliesToVariants:  d.applies_to_variants ?? [],
      // Customer eligibility
      eligibilityType:    d.eligibility_type ?? 'all',
      allowedCustomers:   d.allowed_customers ?? [],
      allowedSegments:    d.allowed_segments ?? [],
      // Usage limits
      usageLimit:         d.usage_limit,
      perCustomerLimit:   d.per_customer_limit,
      currentUsage:       d.current_usage ?? 0,
      usageTracking:      d.usage_tracking ?? [],
      // Buy X Get Y fields (if supported)
      buyProductIds:      d.buy_product_ids ?? [],
      buyQuantity:        d.buy_quantity,
      getProductIds:      d.get_product_ids ?? [],
      getQuantity:        d.get_quantity,
    };
  },

  /**
   * Convert datetime-local format to RFC3339 for backend
   * @param {string} datetimeLocal - datetime-local format string
   * @returns {string} RFC3339 format string
   */
  _convertToRFC3339(datetimeLocal) {
    if (!datetimeLocal) return undefined;
    return new Date(datetimeLocal).toISOString();
  },

  /**
   * Fetch all discounts for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts`);
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
    return this._mapDiscount(res.data);
  },

  /**
   * Fetch paginated products under a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Object} params
   * @param {number} params.page - Page number (default: 1)
   * @param {number} params.limit - Items per page (default: 10)
   * @returns {Promise<Object>}
   */
  async fetchProductsPaginated(shopId, discountId, params = {}) {
    const { page = 1, limit = 10 } = params;
    const res = await api.get(`/seller/shops/${shopId}/discounts/${discountId}/products`, {
      params: { page, limit }
    });
    return {
      products: res.data.products.map(p => ({
        id: p.id,
        name: p.name,
        description: p.description,
        mainImage: p.main_image,
        price: p.price,
        stock: p.stock,
      })),
      total: res.data.total,
      page: res.data.page,
      limit: res.data.limit,
      pages: res.data.pages,
    };
  },

  /**
   * Create a new discount.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.name
   * @param {string} [data.description]
   * @param {string} data.type - 'percentage' or 'fixed'
   * @param {number} data.value
   * @param {Array<string>} [data.appliesToProducts] - Product IDs
   * @param {Array<string>} [data.appliesToVariants] - Variant IDs
   * @param {string} [data.couponCode]
   * @param {string} [data.startAt] datetime-local format
   * @param {string} [data.endAt] datetime-local format
   * @param {boolean} [data.active]
   * @param {string} [data.eligibilityType] - 'all', 'specific', 'segment'
   * @param {Array<string>} [data.allowedCustomers] - Customer IDs
   * @param {Array<string>} [data.allowedSegments] - Segment IDs
   * @param {number} [data.usageLimit]
   * @param {number} [data.perCustomerLimit]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const payload = {
      name:                 data.name,
      description:          data.description,
      category:             'product', // Only product-level discounts supported
      type:                 data.type,
      value:                data.value,
      appliesToProducts:    data.appliesToProducts,
      appliesToVariants:    data.appliesToVariants,
      couponCode:           data.couponCode,
      startAt:              this._convertToRFC3339(data.startAt),
      endAt:                this._convertToRFC3339(data.endAt),
      active:               data.active,
      eligibilityType:      data.eligibilityType,
      allowedCustomers:     data.allowedCustomers,
      allowedSegments:      data.allowedSegments,
      usageLimit:           data.usageLimit,
      perCustomerLimit:     data.perCustomerLimit,
    };
    
    // Clean up undefined fields before sending
    Object.keys(payload).forEach(key => payload[key] === undefined && delete payload[key]);

    const res = await api.post(`/seller/shops/${shopId}/discounts`, payload);
    const newId = res.data.id ?? res.data.insertedId;
    return this.fetchById(shopId, newId);
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
    if (data.type !== undefined) payload.type = data.type;
    if (data.value !== undefined) payload.value = data.value;
    if (data.appliesToProducts !== undefined) payload.appliesToProducts = data.appliesToProducts;
    if (data.appliesToVariants !== undefined) payload.appliesToVariants = data.appliesToVariants;
    if (data.couponCode !== undefined) payload.couponCode = data.couponCode;
    if (data.startAt !== undefined) payload.startAt = this._convertToRFC3339(data.startAt);
    if (data.endAt !== undefined) payload.endAt = this._convertToRFC3339(data.endAt);
    if (data.active !== undefined) payload.active = data.active;
    if (data.eligibilityType !== undefined) payload.eligibilityType = data.eligibilityType;
    if (data.allowedCustomers !== undefined) payload.allowedCustomers = data.allowedCustomers;
    if (data.allowedSegments !== undefined) payload.allowedSegments = data.allowedSegments;
    if (data.usageLimit !== undefined) payload.usageLimit = data.usageLimit;
    if (data.perCustomerLimit !== undefined) payload.perCustomerLimit = data.perCustomerLimit;

    console.log('API call - update discount:', {
      url: `/seller/shops/${shopId}/discounts/${discountId}`,
      payload
    });

    await api.patch(`/seller/shops/${shopId}/discounts/${discountId}`, payload);
    console.log('API response - update discount: success');
    return this.fetchById(shopId, discountId);
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
   * Add mixed customers and segments to a discount's allowed list.
   * @param {string} shopId
   * @param {string} discountId
   * @param {Array<string>} customerIds
   * @param {Array<string>} segmentIds
   * @returns {Promise<Object>}
   */
  async addMixedEligibility(shopId, discountId, customerIds = [], segmentIds = []) {
    console.log('API call - addMixedEligibility:', {
      url: `/seller/shops/${shopId}/discounts/${discountId}/mixed-eligibility`,
      payload: { customerIds, segmentIds }
    });
    
    const res = await api.post(`/seller/shops/${shopId}/discounts/${discountId}/mixed-eligibility`, {
      customerIds: customerIds,
      segmentIds: segmentIds
    });
    console.log('API response - addMixedEligibility:', res.data);
    return res.data;
  },

  /**
   * Clear all allowed customers and segments (set to allow everyone).
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<Object>}
   */
  async clearEligibility(shopId, discountId) {
    const res = await api.post(`/seller/shops/${shopId}/discounts/${discountId}/clear-eligibility`)
    return res.data
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
   * Get all discounts eligible for a customer.
   * @param {string} shopId
   * @param {string} customerId
   * @param {Array<string>} customerSegmentIds
   * @returns {Promise<Array<Object>>}
   */
  async getEligibleForCustomer(shopId, customerId, customerSegmentIds = []) {
    const res = await api.get(`/seller/shops/${shopId}/discounts/eligible`, {
      params: {
        customerId: customerId,
        customerSegmentIds: customerSegmentIds.join(',')
      }
    });
    return res.data.map(this._mapDiscount);
  },

  /**
   * Get all eligible customers for a discount.
   * @param {string} shopId
   * @param {string} discountId
   * @returns {Promise<Object>}
   */
  async getEligibleCustomers(shopId, discountId) {
    const res = await api.get(`/seller/shops/${shopId}/discounts/${discountId}/eligible-customers`);
    return res.data;
  },
};