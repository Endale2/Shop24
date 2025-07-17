import api from './api'

export const analyticsService = {
  async fetchSummary(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/summary`)
    return res.data
  },
  async fetchRevenueOverTime(shopId, days = 30) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/revenue-over-time`, { params: { days } })
    return res.data
  },
  async fetchOrdersOverTime(shopId, days = 30) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/orders-over-time`, { params: { days } })
    return res.data
  },
  async fetchTopProducts(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/top-products`)
    return res.data
  },
  async fetchTopCustomers(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/top-customers`)
    return res.data
  },
  async fetchInventoryStatus(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/inventory-status`)
    return res.data
  },
  async fetchDiscountPerformance(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/discount-performance`)
    return res.data
  },
  async fetchCustomersOverTime(shopId, days = 30) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/customers-over-time`, { params: { days } })
    return res.data
  },
  async fetchCategorySales(shopId, days = 30) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/category-sales`, { params: { days } })
    return res.data
  },
  async fetchRecentOrders(shopId, limit = 10) {
    const res = await api.get(`/seller/shops/${shopId}/analytics/recent-orders`, { params: { limit } })
    return res.data
  },
} 