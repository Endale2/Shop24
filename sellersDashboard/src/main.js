// src/main.js
import { createApp }    from 'vue'
import { createPinia }  from 'pinia'
import piniaPersist     from 'pinia-plugin-persistedstate'
import App              from './App.vue'
import router           from './router'
import { useShopStore } from '@/store/shops'
import { useAuthStore } from '@/store/auth'

const app   = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)

// Detect if we’re in “storefront” subdomain mode
const { hostname, port } = window.location
const isStorefront = port === '5174' && hostname.endsWith('.localhost')

if (!isStorefront) {
  // Normal dashboard app: fetch auth + shops
  const auth = useAuthStore()
  const shops = useShopStore()

  // Try to restore user & shops
  auth.fetchMe().catch(() => {})          // no redirect here
  shops.fetchShops().catch(() => {})      // ignore 401

} else {
  // Storefront: skip auth & shop list entirely
  console.log('🛍️ Running in storefront mode for', hostname)
}

app.mount('#app')
