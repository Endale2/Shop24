// src/stores/collection.js
import { defineStore } from 'pinia'
import { fetchCollections, fetchCollectionDetail } from '@/services/collection'

export const useCollectionStore = defineStore('collection', {
  state: () => ({
    list: [],      // list of { id, title, handle, image }
    detail: null,  // full collection object with products
    loading: false
  }),

  actions: {
    async loadAll(shopSlug) {
      this.loading = true
      try {
        this.list = await fetchCollections(shopSlug)
      } finally {
        this.loading = false
      }
    },

    async loadDetail(shopSlug, handle) {
      this.loading = true
      try {
        this.detail = await fetchCollectionDetail(shopSlug, handle)
      } finally {
        this.loading = false
      }
    }
  }
})
