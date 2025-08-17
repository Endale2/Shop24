import api from './api'

export const componentService = {
  // =============== Component Templates API ===============

  /**
   * Get all available component templates
   * @param {Object} params - Query parameters (page, limit, category, search)
   * @returns {Promise} List of component templates
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
   * Create custom component template
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
   * Update component template
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
   * Delete component template
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
  async getCategories() {
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
   * Save layout components (drag & drop persistence)
   * @param {string} shopId - Shop ID
   * @param {string} layoutId - Layout ID
   * @param {Array} components - Components array with positions
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
   * Get available page types for layouts
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
   * @returns {Promise} Preview data (HTML/CSS)
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

  // =============== Component Utilities ===============

  /**
   * Get default component configuration
   * @param {string} type - Component type
   * @returns {Object} Default configuration
   */
  getDefaultComponentConfig(type) {
    const defaults = {
      hero: {
        settings: {
          height: '400px',
          backgroundType: 'image',
          overlayOpacity: 0.3,
          textAlignment: 'center',
          showButton: true
        },
        content: {
          title: 'Welcome to Our Store',
          subtitle: 'Discover amazing products',
          buttonText: 'Shop Now',
          buttonLink: '/products',
          backgroundImage: ''
        },
        styleConfig: {
          className: 'hero-banner',
          style: {
            minHeight: '400px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }
        }
      },
      'product-grid': {
        settings: {
          columns: 3,
          showRating: true,
          showBadges: true,
          showPrice: true,
          imageRatio: 'square',
          hoverEffect: 'lift'
        },
        content: {
          title: 'Featured Products',
          limit: 8
        }
      },
      testimonials: {
        settings: {
          layout: 'carousel',
          showRating: true,
          showAvatar: true,
          autoplay: true,
          interval: 5000
        },
        content: {
          testimonials: [
            {
              name: 'John Doe',
              rating: 5,
              text: 'Amazing products and great service!',
              avatar: ''
            }
          ]
        }
      },
      newsletter: {
        settings: {
          style: 'inline',
          showIcon: true,
          placeholder: 'Enter your email',
          buttonText: 'Subscribe'
        },
        content: {
          title: 'Stay Updated',
          description: 'Get the latest news and exclusive offers'
        }
      }
    }

    return defaults[type] || {
      settings: {},
      content: {},
      styleConfig: {}
    }
  },

  /**
   * Validate component configuration
   * @param {Object} component - Component data
   * @returns {Object} Validation result
   */
  validateComponent(component) {
    const errors = []

    if (!component.name || component.name.trim() === '') {
      errors.push('Component name is required')
    }

    if (!component.type || component.type.trim() === '') {
      errors.push('Component type is required')
    }

    if (!component.category || component.category.trim() === '') {
      errors.push('Component category is required')
    }

    return {
      isValid: errors.length === 0,
      errors
    }
  }
}

export default componentService
