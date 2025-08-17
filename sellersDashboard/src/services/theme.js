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
   * Save theme colors (optimized endpoint)
   * @param {string} shopId - Shop ID
   * @param {Object} colors - Colors object
   * @returns {Promise} Update result
   */
  async saveThemeColors(shopId, colors) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/colors`, colors)
      return response.data
    } catch (error) {
      console.error('Error saving colors:', error)
      throw error
    }
  },

  /**
   * Save typography settings (optimized endpoint)
   * @param {string} shopId - Shop ID
   * @param {Object} fonts - Fonts object
   * @returns {Promise} Update result
   */
  async saveTypography(shopId, fonts) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/fonts`, fonts)
      return response.data
    } catch (error) {
      console.error('Error saving typography:', error)
      throw error
    }
  },

  /**
   * Save layout settings (optimized endpoint)
   * @param {string} shopId - Shop ID
   * @param {Object} layout - Layout configuration
   * @returns {Promise} Update result
   */
  async saveLayout(shopId, layout) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/layout`, layout)
      return response.data
    } catch (error) {
      console.error('Error saving layout:', error)
      throw error
    }
  },

  /**
   * Reset all customization to defaults
   * @param {string} shopId - Shop ID
   * @returns {Promise} Reset result
   */
  async resetToDefaults(shopId) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/customization/reset`)
      return response.data
    } catch (error) {
      console.error('Error resetting customization:', error)
      throw error
    }
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

  // =============== Advanced Styling API ===============

  /**
   * Update theme gradients
   * @param {string} shopId - Shop ID
   * @param {Object} gradients - Gradients configuration
   * @returns {Promise} Update result
   */
  async updateGradients(shopId, gradients) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/gradients`, { gradients })
      return response.data
    } catch (error) {
      console.error('Error updating gradients:', error)
      throw error
    }
  },

  /**
   * Update theme shadows
   * @param {string} shopId - Shop ID
   * @param {Object} shadows - Shadows configuration
   * @returns {Promise} Update result
   */
  async updateShadows(shopId, shadows) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/shadows`, { shadows })
      return response.data
    } catch (error) {
      console.error('Error updating shadows:', error)
      throw error
    }
  },

  /**
   * Update theme animations
   * @param {string} shopId - Shop ID
   * @param {Object} animations - Animations configuration
   * @returns {Promise} Update result
   */
  async updateAnimations(shopId, animations) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/animations`, { animations })
      return response.data
    } catch (error) {
      console.error('Error updating animations:', error)
      throw error
    }
  },

  /**
   * Update mobile configuration
   * @param {string} shopId - Shop ID
   * @param {Object} mobileConfig - Mobile configuration
   * @returns {Promise} Update result
   */
  async updateMobileConfig(shopId, mobileConfig) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/mobile`, { mobileConfig })
      return response.data
    } catch (error) {
      console.error('Error updating mobile config:', error)
      throw error
    }
  },

  /**
   * Update SEO configuration
   * @param {string} shopId - Shop ID
   * @param {Object} seoConfig - SEO configuration
   * @returns {Promise} Update result
   */
  async updateSEOConfig(shopId, seoConfig) {
    try {
      const response = await api.patch(`/seller/shops/${shopId}/customization/seo`, { seoConfig })
      return response.data
    } catch (error) {
      console.error('Error updating SEO config:', error)
      throw error
    }
  },

  // =============== CSS Management API ===============

  /**
   * Validate custom CSS
   * @param {string} shopId - Shop ID
   * @param {string} css - CSS content
   * @returns {Promise} Validation result
   */
  async validateCSS(shopId, css) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/customization/css/validate`, { css })
      return response.data
    } catch (error) {
      console.error('Error validating CSS:', error)
      throw error
    }
  },

  /**
   * Compile custom CSS
   * @param {string} shopId - Shop ID
   * @param {string} css - CSS content
   * @returns {Promise} Compilation result
   */
  async compileCSS(shopId, css) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/customization/css/compile`, { css })
      return response.data
    } catch (error) {
      console.error('Error compiling CSS:', error)
      throw error
    }
  },

  /**
   * Get CSS variables for theme
   * @param {string} shopId - Shop ID
   * @returns {Promise} CSS variables
   */
  async getCSSVariables(shopId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/customization/css/variables`)
      return response.data
    } catch (error) {
      console.error('Error fetching CSS variables:', error)
      throw error
    }
  },

  // =============== Component Library API ===============

  /**
   * Get available components
   * @param {Object} params - Query parameters
   * @returns {Promise} List of components
   */
  async getComponents(params = {}) {
    try {
      const response = await api.get('/seller/components', { params })
      return response.data
    } catch (error) {
      console.error('Error fetching components:', error)
      throw error
    }
  },

  /**
   * Get component by ID
   * @param {string} componentId - Component ID
   * @returns {Promise} Component data
   */
  async getComponent(componentId) {
    try {
      const response = await api.get(`/seller/components/${componentId}`)
      return response.data
    } catch (error) {
      console.error('Error fetching component:', error)
      throw error
    }
  },

  /**
   * Create custom component
   * @param {Object} componentData - Component data
   * @returns {Promise} Creation result
   */
  async createComponent(componentData) {
    try {
      const response = await api.post('/seller/components', componentData)
      return response.data
    } catch (error) {
      console.error('Error creating component:', error)
      throw error
    }
  },

  /**
   * Update component
   * @param {string} componentId - Component ID
   * @param {Object} updateData - Update data
   * @returns {Promise} Update result
   */
  async updateComponent(componentId, updateData) {
    try {
      const response = await api.put(`/seller/components/${componentId}`, updateData)
      return response.data
    } catch (error) {
      console.error('Error updating component:', error)
      throw error
    }
  },

  /**
   * Delete component
   * @param {string} componentId - Component ID
   * @returns {Promise} Deletion result
   */
  async deleteComponent(componentId) {
    try {
      const response = await api.delete(`/seller/components/${componentId}`)
      return response.data
    } catch (error) {
      console.error('Error deleting component:', error)
      throw error
    }
  },

  /**
   * Get component categories
   * @returns {Promise} List of categories
   */
  async getComponentCategories() {
    try {
      const response = await api.get('/seller/components/categories')
      return response.data
    } catch (error) {
      console.error('Error fetching component categories:', error)
      throw error
    }
  },

  /**
   * Search components
   * @param {string} query - Search query
   * @returns {Promise} Search results
   */
  async searchComponents(query) {
    try {
      const response = await api.get('/seller/components/search', { params: { q: query } })
      return response.data
    } catch (error) {
      console.error('Error searching components:', error)
      throw error
    }
  },

  /**
   * Get components by category
   * @param {string} category - Category name
   * @returns {Promise} Components in category
   */
  async getComponentsByCategory(category) {
    try {
      const response = await api.get(`/seller/components/category/${category}`)
      return response.data
    } catch (error) {
      console.error('Error fetching components by category:', error)
      throw error
    }
  },

  // =============== Layout Management API ===============

  /**
   * Get shop layouts
   * @param {string} shopId - Shop ID
   * @returns {Promise} List of layouts
   */
  async getLayouts(shopId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/layouts`)
      return response.data
    } catch (error) {
      console.error('Error fetching layouts:', error)
      throw error
    }
  },

  /**
   * Get specific layout
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @returns {Promise} Layout data
   */
  async getLayout(shopId, layoutId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/layouts/${layoutId}`)
      return response.data
    } catch (error) {
      console.error('Error fetching layout:', error)
      throw error
    }
  },

  /**
   * Create layout
   * @param {string} shopId - Shop ID
   * @param {Object} layoutData - Layout data
   * @returns {Promise} Creation result
   */
  async createLayout(shopId, layoutData) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/layouts`, layoutData)
      return response.data
    } catch (error) {
      console.error('Error creating layout:', error)
      throw error
    }
  },

  /**
   * Update layout
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @param {Object} updateData - Update data
   * @returns {Promise} Update result
   */
  async updateLayout(shopId, layoutId, updateData) {
    try {
      const response = await api.put(`/seller/shops/${shopId}/layouts/${layoutId}`, updateData)
      return response.data
    } catch (error) {
      console.error('Error updating layout:', error)
      throw error
    }
  },

  /**
   * Delete layout
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @returns {Promise} Deletion result
   */
  async deleteLayout(shopId, layoutId) {
    try {
      const response = await api.delete(`/seller/shops/${shopId}/layouts/${layoutId}`)
      return response.data
    } catch (error) {
      console.error('Error deleting layout:', error)
      throw error
    }
  },

  /**
   * Save layout components
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @param {Array} components - Components array
   * @returns {Promise} Save result
   */
  async saveLayoutComponents(shopId, layoutId, components) {
    try {
      const response = await api.post(`/seller/shops/${shopId}/layouts/${layoutId}/components`, { components })
      return response.data
    } catch (error) {
      console.error('Error saving layout components:', error)
      throw error
    }
  },

  /**
   * Get page types
   * @returns {Promise} List of page types
   */
  async getPageTypes() {
    try {
      const response = await api.get('/seller/layouts/page-types')
      return response.data
    } catch (error) {
      console.error('Error fetching page types:', error)
      throw error
    }
  },

  /**
   * Preview layout
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @returns {Promise} Preview data
   */
  async previewLayout(shopId, layoutId) {
    try {
      const response = await api.get(`/seller/shops/${shopId}/layouts/${layoutId}/preview`)
      return response.data
    } catch (error) {
      console.error('Error generating layout preview:', error)
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
        bodyText: '#6B7280',
        border: '#E5E7EB',
        accent: '#3B82F6',
        success: '#10B981',
        warning: '#F59E0B',
        error: '#EF4444'
      },
      fonts: {
        heading: 'Inter',
        body: 'Inter',
        headingWeight: '600',
        bodyWeight: '400',
        headingSize: '2rem',
        bodySize: '1rem'
      },
      layout: {
        headerStyle: 'classic',
        containerWidth: 'boxed',
        sidebarPosition: 'none',
        gridColumns: '3',
        spacing: 'medium',
        borderRadius: '0.375rem'
      },
      gradients: {
        primary: 'linear-gradient(135deg, #10B981 0%, #059669 100%)',
        secondary: 'linear-gradient(135deg, #F59E0B 0%, #D97706 100%)',
        hero: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
        card: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)'
      },
      shadows: {
        sm: '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
        md: '0 4px 6px -1px rgba(0, 0, 0, 0.1)',
        lg: '0 10px 15px -3px rgba(0, 0, 0, 0.1)',
        xl: '0 20px 25px -5px rgba(0, 0, 0, 0.1)',
        card: '0 4px 6px -1px rgba(0, 0, 0, 0.1)',
        button: '0 2px 4px -1px rgba(0, 0, 0, 0.1)'
      },
      animations: {
        fadeIn: 'fadeIn 0.3s ease-in-out',
        slideUp: 'slideUp 0.3s ease-out',
        bounce: 'bounce 0.5s ease-in-out',
        pulse: 'pulse 2s infinite',
        hover: 'transform 0.2s ease-in-out'
      },
      spacing: {
        xs: '0.25rem',
        sm: '0.5rem',
        md: '1rem',
        lg: '1.5rem',
        xl: '3rem',
        xxl: '6rem'
      },
      mobile: {
        navigation: {
          style: 'drawer',
          position: 'left',
          showLabels: true
        },
        layout: {
          stackOnMobile: true,
          mobileColumns: 1,
          hideOnMobile: []
        },
        performance: {
          lazyLoading: true,
          imageOptimization: true,
          minifyCSS: true
        }
      },
      seo: {
        meta: {
          title: '',
          description: '',
          keywords: '',
          author: ''
        },
        openGraph: {
          title: '',
          description: '',
          image: '',
          type: 'website'
        },
        twitter: {
          card: 'summary_large_image',
          title: '',
          description: '',
          image: ''
        },
        structuredData: {
          enabled: true,
          type: 'Store',
          logo: '',
          contactPoint: {}
        }
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
