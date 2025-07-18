import api from './api'

export const dashboardService = {
  async fetchDashboardData(shopId) {
    if (!shopId) throw new Error('shopId is required')
    const response = await api.get(`/seller/shops/${shopId}/analytics/dashboard`)
    return response.data
  }
}
