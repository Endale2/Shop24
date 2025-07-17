import api from './api'

export const orderService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/orders`)
    if (!Array.isArray(res.data)) return [];
    return res.data.map(o => ({
      id: o.ID ?? o.id,
      orderNumber: o.order_number ?? o.OrderNumber,
      shopId: o.shop_id ?? o.ShopID,
      customerId: o.customer_id ?? o.CustomerID,
      items: (o.Items ?? o.items ?? []).map(item => ({
        productId: item.product_id ?? item.ProductID,
        variantId: item.variant_id ?? item.VariantID,
        variantName: item.variant_name ?? item.VariantName,
        name: item.Name ?? item.name,
        quantity: item.Quantity ?? item.quantity,
        unitPrice: Number(item.unit_price ?? item.UnitPrice) || 0,
        totalPrice: Number(item.total_price ?? item.TotalPrice) || 0,
        image: item.image ?? item.Image,
      })),
      subtotal: Number(o.Subtotal ?? o.subtotal) || 0,
      discountTotal: Number(o.discount_total ?? o.DiscountTotal) || 0,
      total: Number(o.Total ?? o.total) || 0,
      status: o.Status ?? o.status,
      couponCode: o.coupon_code ?? o.CouponCode,
      createdAt: o.CreatedAt ?? o.created_at ?? o.createdAt,
      updatedAt: o.UpdatedAt ?? o.updated_at ?? o.updatedAt,
    }))
  },

  async fetchById(shopId, orderId) {
    const res = await api.get(`/seller/shops/${shopId}/orders/${orderId}`)
    const o = res.data
    return {
      id: o.ID ?? o.id,
      orderNumber: o.order_number ?? o.OrderNumber,
      shopId: o.shop_id ?? o.ShopID,
      customerId: o.customer_id ?? o.CustomerID,
      items: (o.Items ?? o.items ?? []).map(item => ({
        productId: item.product_id ?? item.ProductID,
        variantId: item.variant_id ?? item.VariantID,
        variantName: item.variant_name ?? item.VariantName,
        name: item.Name ?? item.name,
        quantity: item.Quantity ?? item.quantity,
        unitPrice: Number(item.unit_price ?? item.UnitPrice) || 0,
        totalPrice: Number(item.total_price ?? item.TotalPrice) || 0,
        image: item.image ?? item.Image,
      })),
      subtotal: Number(o.Subtotal ?? o.subtotal) || 0,
      discountTotal: Number(o.discount_total ?? o.DiscountTotal) || 0,
      total: Number(o.Total ?? o.total) || 0,
      status: o.Status ?? o.status,
      couponCode: o.coupon_code ?? o.CouponCode,
      createdAt: o.CreatedAt ?? o.created_at ?? o.createdAt,
      updatedAt: o.UpdatedAt ?? o.updated_at ?? o.updatedAt,
    }
  },

  async fetchByIdWithCustomer(shopId, orderId) {
    const res = await api.get(`/seller/shops/${shopId}/orders/${orderId}/details`)
    const data = res.data
    
    return {
      order: {
        id: data.order.ID ?? data.order.id,
        orderNumber: data.order.order_number ?? data.order.OrderNumber,
        shopId: data.order.shop_id ?? data.order.ShopID,
        customerId: data.order.customer_id ?? data.order.CustomerID,
        items: (data.order.Items ?? data.order.items ?? []).map(item => ({
          productId: item.product_id ?? item.ProductID,
          variantId: item.variant_id ?? item.VariantID,
          variantName: item.variant_name ?? item.VariantName,
          name: item.Name ?? item.name,
          quantity: item.Quantity ?? item.quantity,
          unitPrice: Number(item.unit_price ?? item.UnitPrice) || 0,
          totalPrice: Number(item.total_price ?? item.TotalPrice) || 0,
          image: item.image ?? item.Image,
        })),
        subtotal: Number(data.order.Subtotal ?? data.order.subtotal) || 0,
        discountTotal: Number(data.order.discount_total ?? data.order.DiscountTotal) || 0,
        total: Number(data.order.Total ?? data.order.total) || 0,
        status: data.order.Status ?? data.order.status,
        couponCode: data.order.coupon_code ?? data.order.CouponCode,
        createdAt: data.order.CreatedAt ?? data.order.created_at ?? data.order.createdAt,
        updatedAt: data.order.UpdatedAt ?? data.order.updated_at ?? data.order.updatedAt,
      },
      customer: data.customer ? {
        id: data.customer.ID ?? data.customer.id,
        firstName: data.customer.firstName ?? data.customer.FirstName,
        lastName: data.customer.lastName ?? data.customer.LastName,
        profile_image: data.customer.profile_image ?? data.customer.ProfileImage,
        email: data.customer.email ?? data.customer.Email,
        phone: data.customer.phone ?? data.customer.Phone,
        address: data.customer.address ?? data.customer.Address,
        city: data.customer.city ?? data.customer.City,
        state: data.customer.state ?? data.customer.State,
        postalCode: data.customer.postalCode ?? data.customer.PostalCode,
        country: data.customer.country ?? data.customer.Country,
        createdAt: data.customer.createdAt ?? data.customer.CreatedAt,
        updatedAt: data.customer.updatedAt ?? data.customer.UpdatedAt,
      } : null,
    }
  },

  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/orders`, data)
    const created = res.data
    const id = created.id ?? created._id ?? created.ID ?? created.InsertedID
    if (!id) {
      throw new Error('Failed to retrieve new order ID from response')
    }
    return this.fetchById(shopId, id)
  },

  // Helper method to create order from cart items
  async createFromCartItems(shopId, items, couponCode = null) {
    const orderData = {
      items: items.map(item => ({
        product_id: item.productId,
        variant_id: item.variantId || '',
        quantity: item.quantity
      }))
    }
    
    if (couponCode) {
      orderData.coupon_code = couponCode
    }
    
    return this.create(shopId, orderData)
  },

  async update(shopId, orderId, data) {
    await api.patch(`/seller/shops/${shopId}/orders/${orderId}`, data)
    return this.fetchById(shopId, orderId)
  },

  async delete(shopId, orderId) {
    await api.delete(`/seller/shops/${shopId}/orders/${orderId}`)
  }
} 