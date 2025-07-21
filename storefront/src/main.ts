import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import './index.css'
import { useAuthStore } from './stores/auth'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)

app.mount('#app')

// Restore user on every full-page load (like sellersDashboard)
const authStore = useAuthStore()
authStore.fetchProfile().catch(() => {})

// Optionally, keep periodic refresh if desired
setInterval(async () => {
  if (authStore.user) {
    try {
      await authStore.refreshToken()
    } catch {
      authStore.user = null
    }
  }
}, 10 * 60 * 1000)
