// main.js
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import App from './App.vue'
import router from './router'
import { useAuthStore } from '@/store/auth'
import { useShopStore } from '@/store/shops'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)

// re-hydrate on every full-page load
const auth = useAuthStore()
const shops = useShopStore()
auth.fetchMe().catch(() => {})
shops.fetchShops().catch(() => {})

app.mount('#app')
