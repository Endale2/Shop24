// store/navigation.js
import { defineStore } from 'pinia'

export const useNavigationStore = defineStore('navigation', {
  state: () => ({
    activeNavItem: null,
    isNavigating: false
  }),
  
  getters: {
    getActiveNavItem: (state) => state.activeNavItem,
    getIsNavigating: (state) => state.isNavigating
  },
  
  actions: {
    setActiveNavItem(path) {
      this.activeNavItem = path
      this.isNavigating = true
    },
    
    clearNavigating() {
      this.isNavigating = false
    },
    
    // Initialize active nav item from current route
    initializeFromRoute(routePath) {
      if (!this.activeNavItem) {
        this.activeNavItem = routePath
      }
      this.isNavigating = false
    }
  }
})
