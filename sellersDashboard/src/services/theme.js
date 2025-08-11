import api from './api'

export const themeService = {
  // =============== Shop Customization API ===============
  
  /**
   * Get current customization data for a shop
   * @param {string} shopId - Shop ID
   * @returns {Promise} Customization data
   */
  async getCustomization(shopId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/customization`)
      return response.data
    } catch (error) {
      console.error('Error fetching customization:', error)
      throw error
    }
  },

  /**
   * Update shop customization
   * @param {string} shopId - Shop ID
   * @param {Object} customizationData - Customization data
   * @returns {Promise} Update result
   */
  async updateCustomization(shopId, customizationData) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization`, customizationData)
      return response.data
    } catch (error) {
      console.error('Error updating customization:', error)
      throw error
    }
  },

  /**
   * Save theme colors
   * @param {string} shopId - Shop ID
   * @param {Object} colors - Colors object
   * @returns {Promise} Update result
   */
  async saveThemeColors(shopId, colors) {
    return this.updateCustomization(shopId, { colors })
  },

  /**
   * Save typography settings
   * @param {string} shopId - Shop ID
   * @param {Object} fonts - Fonts object
   * @returns {Promise} Update result
   */
  async saveTypography(shopId, fonts) {
    return this.updateCustomization(shopId, { fonts })
  },

  /**
   * Save layout settings
   * @param {string} shopId - Shop ID
   * @param {Object} layout - Layout configuration
   * @returns {Promise} Update result
   */
  async saveLayout(shopId, layout) {
    return this.updateCustomization(shopId, { layout })
  },

  // =============== Theme Management API ===============

  /**
   * Get all marketplace themes
   * @returns {Promise} List of public themes
   */
  async getMarketplaceThemes() {
    try {
      const response = await api.get('/seller/themes/marketplace')
      return response.data.themes || []
    } catch (error) {
      console.error('Error fetching marketplace themes:', error)
      throw error
    }
  },

  /**
   * Get seller's custom themes
   * @returns {Promise} List of custom themes
   */
  async getMyThemes() {
    try {
      const response = await api.get('/seller/themes/my-themes')
      return response.data.themes || []
    } catch (error) {
      console.error('Error fetching custom themes:', error)
      throw error
    }
  },

  /**
   * Get specific theme by ID
   * @param {string} themeId - Theme ID
   * @returns {Promise} Theme data
   */
  async getTheme(themeId) {
    try {
      const response = await api.get(`/seller/themes/${themeId}`)
      return response.data.theme
    } catch (error) {
      console.error('Error fetching theme:', error)
      throw error
    }
  },

  /**
   * Search themes
   * @param {string} query - Search query
   * @param {boolean} publicOnly - Search only public themes
   * @returns {Promise} Search results
   */
  async searchThemes(query, publicOnly = true) {
    try {
      const response = await api.get('/seller/themes/search', {
        params: { q: query, public: publicOnly }
      })
      return response.data.themes || []
    } catch (error) {
      console.error('Error searching themes:', error)
      throw error
    }
  },

  /**
   * Create custom theme
   * @param {Object} themeData - Theme data
   * @returns {Promise} Creation result
   */
  async createTheme(themeData) {
    try {
      const response = await api.post('/seller/themes', themeData)
      return response.data
    } catch (error) {
      console.error('Error creating theme:', error)
      throw error
    }
  },

  /**
   * Update theme
   * @param {string} themeId - Theme ID
   * @param {Object} updateData - Update data
   * @returns {Promise} Update result
   */
  async updateTheme(themeId, updateData) {
    try {
      const response = await api.patch(`/seller/themes/${themeId}`, updateData)
      return response.data
    } catch (error) {
      console.error('Error updating theme:', error)
      throw error
    }
  },

  /**
   * Delete theme
   * @param {string} themeId - Theme ID
   * @returns {Promise} Deletion result
   */
  async deleteTheme(themeId) {
    try {
      const response = await api.delete(`/seller/themes/${themeId}`)
      return response.data
    } catch (error) {
      console.error('Error deleting theme:', error)
      throw error
    }
  },

  // =============== Shop Theme Configuration API ===============

  /**
   * Get shop's active theme
   * @param {string} shopId - Shop ID
   * @returns {Promise} Active theme data
   */
  async getShopActiveTheme(shopId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/themes/active`)
      return response.data
    } catch (error) {
      console.error('Error fetching active theme:', error)
      throw error
    }
  },

  /**
   * Get all shop theme configurations
   * @param {string} shopId - Shop ID
   * @returns {Promise} List of shop themes
   */
  async getShopThemes(shopId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/themes`)
      return response.data.shop_themes || []
    } catch (error) {
      console.error('Error fetching shop themes:', error)
      throw error
    }
  },

  /**
   * Create shop theme configuration
   * @param {string} shopId - Shop ID
   * @param {Object} themeConfig - Theme configuration
   * @returns {Promise} Creation result
   */
  async createShopTheme(shopId, themeConfig) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/themes`, themeConfig)
      return response.data
    } catch (error) {
      console.error('Error creating shop theme:', error)
      throw error
    }
  },

  /**
   * Update shop theme configuration
   * @param {string} shopId - Shop ID
   * @param {string} shopThemeId - Shop theme ID
   * @param {Object} updateData - Update data
   * @returns {Promise} Update result
   */
  async updateShopTheme(shopId, shopThemeId, updateData) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/themes/${shopThemeId}`, updateData)
      return response.data
    } catch (error) {
      console.error('Error updating shop theme:', error)
      throw error
    }
  },

  /**
   * Publish/activate a theme for shop
   * @param {string} shopId - Shop ID
   * @param {string} shopThemeId - Shop theme ID
   * @returns {Promise} Publish result
   */
  async publishShopTheme(shopId, shopThemeId) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/themes/${shopThemeId}/publish`)
      return response.data
    } catch (error) {
      console.error('Error publishing theme:', error)
      throw error
    }
  },

  /**
   * Delete shop theme configuration
   * @param {string} shopId - Shop ID
   * @param {string} shopThemeId - Shop theme ID
   * @returns {Promise} Deletion result
   */
  async deleteShopTheme(shopId, shopThemeId) {
    try {
      const response = await api.delete(`/seller/shops/${shopId}/themes/${shopThemeId}`)
      return response.data
    } catch (error) {
      console.error('Error deleting shop theme:', error)
      throw error
    }
  },

  // =============== Theme Presets API ===============

  /**
   * Get theme presets
   * @param {string} themeId - Theme ID
   * @returns {Promise} List of presets
   */
  async getThemePresets(themeId) {
    try {
      const response = await api.get(`/seller/themes/${themeId}/presets`)
      return response.data.presets || []
    } catch (error) {
      console.error('Error fetching theme presets:', error)
      throw error
    }
  },

  /**
   * Apply theme preset to shop theme
   * @param {string} shopId - Shop ID
   * @param {string} shopThemeId - Shop theme ID
   * @param {string} presetId - Preset ID
   * @returns {Promise} Apply result
   */
  async applyThemePreset(shopId, shopThemeId, presetId) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/themes/${shopThemeId}/presets/${presetId}/apply`)
      return response.data
    } catch (error) {
      console.error('Error applying theme preset:', error)
      throw error
    }
  },

  // =============== Utility Functions ===============

  /**
   * Get default theme configuration
   * @returns {Object} Default theme config
   */
  getDefaultThemeConfig() {
    return {
      colors: {
        primary: '#10B981',
        secondary: '#F59E0B',
        background: '#FFFFFF',
        heading: '#1F2937',
        bodyText: '#6B7280'
      },
      fonts: {
        heading: 'Inter',
        body: 'Inter'
      },
      layout: {
        headerStyle: 'classic',
        containerWidth: 'boxed',
        sidebarPosition: 'none',
        gridColumns: '3'
      }
    }
  },

  /**
   * Validate theme configuration
   * @param {Object} config - Theme configuration
   * @returns {Object} Validation result
   */
  validateThemeConfig(config) {
    const errors = {}

    // Validate colors
    if (config.colors) {
      for (const [key, value] of Object.entries(config.colors)) {
        if (!/^#[0-9A-F]{6}$/i.test(value)) {
          errors[`colors.${key}`] = 'Invalid color format'
        }
      }
    }

    // Validate fonts
    if (config.fonts) {
      for (const [key, value] of Object.entries(config.fonts)) {
        if (!value || value.trim() === '') {
          errors[`fonts.${key}`] = 'Font cannot be empty'
        }
      }
    }

    return {
      isValid: Object.keys(errors).length === 0,
      errors
    }
  }
}

export default themeService
