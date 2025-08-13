// main.js
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import App from './App.vue'
import router from './router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'
import { useNavigationStore } from '@/store/navigation'

const app = createApp(App)
const pinia = createPinia()

// Configure persistence with error handling
pinia.use(piniaPersist)

app.use(pinia)
app.use(router)

// Enhanced initialization with proper error handling and recovery
async function initializeApp() {
  const auth = useAuthStore()
  const shops = useShopStore()
  const navigation = useNavigationStore()

  try {
    // First, check if we have a persisted user
    if (auth.user) {
      console.log('[App] Found persisted user, verifying session...')

      // Verify the session is still valid
      const isValid = await auth.verifySession()

      if (isValid) {
        console.log('[App] Session verified, loading shops...')
        // Only load shops if auth is valid
        try {
          await shops.fetchShops()
        } catch (shopError) {
          console.warn('[App] Failed to load shops, but auth is valid:', shopError)
          // Don't fail the entire initialization if shops fail to load
        }
      } else {
        console.warn('[App] Session verification failed, user will need to re-authenticate')
      }
    } else {
      console.log('[App] No persisted user found')
    }
  } catch (error) {
    console.error('[App] Initialization error:', error)

    // Handle persistence errors gracefully
    if (error.message?.includes('localStorage') || error.message?.includes('persist')) {
      auth.handlePersistenceError(error)
    }

    // Don't block app initialization for auth errors
    // The router guards will handle redirecting unauthenticated users
  }
}

// Mount app and initialize
app.mount('#app')

// Initialize after mounting to ensure stores are ready
initializeApp().catch(error => {
  console.error('[App] Failed to initialize app:', error)
})
