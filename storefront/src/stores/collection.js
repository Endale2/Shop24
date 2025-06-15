import { defineStore } from 'pinia'
import {
  fetchCollections,
  fetchCollectionDetail
} from '@/services/collection'

export const useCollectionStore = defineStore('collection', {
  state: () => ({
    list:    [],
    current: null,
    loading: false
  }),
  actions: {
    async loadList(shopSlug) {
      this.loading = true
      try {
        this.list = await fetchCollections(shopSlug)
      } finally {
        this.loading = false
      }
    },
    async loadDetail(shopSlug, collId) {
      this.loading = true
      try {
        this.current = await fetchCollectionDetail(shopSlug, collId)
      } finally {
        this.loading = false
      }
    }
  },
  persist: true
})
