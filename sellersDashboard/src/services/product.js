import api from './api'

export const productService = {
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id: p.id || p._id || p.ID,
      shopId: p.shop_id || p.shopId || p.ShopID,
      name: p.name || p.Name,
      description: p.description || p.Description,
      images: p.images || p.Images || [],
      main_image: p.main_image || p.MainImage || '',
      collection_ids: p.collection_ids || p.CollectionIDs || [],
      price: Number(p.price || p.Price) || 0,
      stock: p.stock || p.Stock || 0,
      starting_price: p.starting_price,
      total_stock: p.total_stock,
      variants: (p.variants || p.Variants || []).map(v => ({
        id: v.id || v.variant_id || v.VariantID,
        options: v.options || v.Options || [],
        price: Number(v.price || v.Price) || 0,
        stock: v.stock || v.Stock || 0,
        image: v.image || v.Image || '',
        createdAt: v.createdAt || v.CreatedAt,
        updatedAt: v.updatedAt || v.UpdatedAt,
      })),
      createdAt: p.createdAt || p.CreatedAt,
      updatedAt: p.updatedAt || p.UpdatedAt,
      meta_title: p.meta_title || '',
      meta_description: p.meta_description || '',
    }))
  },

  // New paginated version for ProductsPage.vue
  async fetchAllByShopPaginated(shopId, { page = 1, limit = 10, search = '', stockStatus = '' } = {}) {
    const params = new URLSearchParams({
      page: page.toString(),
      limit: limit.toString(),
    })
    if (search) params.append('search', search)
    if (stockStatus) params.append('stock_status', stockStatus)

    const res = await api.get(`/seller/shops/${shopId}/products?${params}`)
    if (res.data && Array.isArray(res.data.products)) {
      return {
        products: res.data.products.map(p => ({
          id: p.id || p._id || p.ID,
          shopId: p.shop_id || p.shopId || p.ShopID,
          name: p.name || p.Name,
          description: p.description || p.Description,
          images: p.images || p.Images || [],
          main_image: p.main_image || p.MainImage || '',
          collection_ids: p.collection_ids || p.CollectionIDs || [],
          price: Number(p.price || p.Price) || 0,
          stock: p.stock || p.Stock || 0,
          starting_price: p.starting_price,
          total_stock: p.total_stock,
          variants: (p.variants || p.Variants || []).map(v => ({
            id: v.id || v.variant_id || v.VariantID,
            options: v.options || v.Options || [],
            price: Number(v.price || v.Price) || 0,
            stock: v.stock || v.Stock || 0,
            image: v.image || v.Image || '',
            createdAt: v.createdAt || v.CreatedAt,
            updatedAt: v.updatedAt || v.UpdatedAt,
          })),
          createdAt: p.createdAt || p.CreatedAt,
          updatedAt: p.updatedAt || p.UpdatedAt,
          meta_title: p.meta_title || '',
          meta_description: p.meta_description || '',
        })),
        pagination: res.data.pagination,
        stats: res.data.stats,
      }
    }
    // fallback: old array format
    return {
      products: res.data.map(p => ({
        id: p.id || p._id || p.ID,
        shopId: p.shop_id || p.shopId || p.ShopID,
        name: p.name || p.Name,
        description: p.description || p.Description,
        images: p.images || p.Images || [],
        main_image: p.main_image || p.MainImage || '',
        collection_ids: p.collection_ids || p.CollectionIDs || [],
        price: Number(p.price || p.Price) || 0,
        stock: p.stock || p.Stock || 0,
        starting_price: p.starting_price,
        total_stock: p.total_stock,
        variants: (p.variants || p.Variants || []).map(v => ({
          id: v.id || v.variant_id || v.VariantID,
          options: v.options || v.Options || [],
          price: Number(v.price || v.Price) || 0,
          stock: v.stock || v.Stock || 0,
          image: v.image || v.Image || '',
          createdAt: v.createdAt || v.CreatedAt,
          updatedAt: v.updatedAt || v.UpdatedAt,
        })),
        createdAt: p.createdAt || p.CreatedAt,
        updatedAt: p.updatedAt || p.UpdatedAt,
        meta_title: p.meta_title || '',
        meta_description: p.meta_description || '',
      })),
      pagination: null,
      stats: null,
    }
  },

  /**
   * Fetch all products for a shop, with optional filter for collection_id.
   * @param {string} shopId
   * @param {Object} filter - e.g. { collection_id: null } or { collection_id: 'someId' }
   * @returns {Promise<Array<Object>>}
   */
  async fetchAllByShopWithFilter(shopId, filter = {}) {
    const params = new URLSearchParams()
    if (filter.collection_id === null) {
      params.append('collection_id', 'null')
    } else if (filter.collection_id) {
      params.append('collection_id', filter.collection_id)
    }
    const url = params.toString()
      ? `/seller/shops/${shopId}/products?${params}`
      : `/seller/shops/${shopId}/products`
    const res = await api.get(url)
    return res.data.map(p => ({
      id: p.id || p._id || p.ID,
      shopId: p.shop_id || p.shopId || p.ShopID,
      name: p.name || p.Name,
      description: p.description || p.Description,
      images: p.images || p.Images || [],
      main_image: p.main_image || p.MainImage || '',
      collection_ids: p.collection_ids || p.CollectionIDs || [],
      price: Number(p.price || p.Price) || 0,
      stock: p.stock || p.Stock || 0,
      starting_price: p.starting_price,
      total_stock: p.total_stock,
      variants: (p.variants || p.Variants || []).map(v => ({
        id: v.id || v.variant_id || v.VariantID,
        options: v.options || v.Options || [],
        price: Number(v.price || v.Price) || 0,
        stock: v.stock || v.Stock || 0,
        image: v.image || v.Image || '',
        createdAt: v.createdAt || v.CreatedAt,
        updatedAt: v.updatedAt || v.UpdatedAt,
      })),
      createdAt: p.createdAt || p.CreatedAt,
      updatedAt: p.updatedAt || p.UpdatedAt,
      meta_title: p.meta_title || '',
      meta_description: p.meta_description || '',
    }))
  },

  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id: p.id || p._id || p.ID,
      shopId: p.shop_id || p.shopId || p.ShopID,
      name: p.name || p.Name,
      description: p.description || p.Description,
      images: p.images || p.Images || [],
      main_image: p.main_image || p.MainImage || '',
      collection_ids: p.collection_ids || p.CollectionIDs || [],
      price: Number(p.price || p.Price) || 0,
      stock: p.stock || p.Stock || 0,
      starting_price: p.starting_price,
      total_stock: p.total_stock,
      variants: (p.variants || p.Variants || []).map(v => ({
        id: v.id || v.variant_id || v.VariantID,
        options: v.options || v.Options || [],
        price: Number(v.price || v.Price) || 0,
        stock: v.stock || v.Stock || 0,
        image: v.image || v.Image || '',
        createdAt: v.createdAt || v.CreatedAt,
        updatedAt: v.updatedAt || v.UpdatedAt,
      })),
      createdAt: p.createdAt || p.CreatedAt,
      updatedAt: p.updatedAt || p.UpdatedAt,
      meta_title: p.meta_title || '',
      meta_description: p.meta_description || '',
    }
  },

  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    const created = res.data
    const id = created.id || created._id || created.ID || created.InsertedID
    if (!id) {
      throw new Error('Failed to retrieve new product ID from response')
    }
    return this.fetchById(shopId, id)
  },

  async update(shopId, productId, data) {
    await api.put(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },

  async patch(shopId, productId, data) {
    await api.patch(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },

  async delete(shopId, productId) {
    await api.delete(`/seller/shops/${shopId}/products/${productId}`)
  },

  /**
   * Add products to a collection using the new API endpoint
   * @param {string} shopId
   * @param {string} collectionId
   * @param {Array<string>} productIds
   * @returns {Promise<void>}
   */
  async addToCollection(shopId, collectionId, productIds) {
    await api.post(`/seller/shops/${shopId}/collections/${collectionId}/products`, {
      product_ids: productIds
    })
  },

  /**
   * Remove products from a collection using the new API endpoint
   * @param {string} shopId
   * @param {string} collectionId
   * @param {Array<string>} productIds
   * @returns {Promise<void>}
   */
  async removeFromCollection(shopId, collectionId, productIds) {
    await api.delete(`/seller/shops/${shopId}/collections/${collectionId}/products`, {
      data: { product_ids: productIds }
    })
  }
}