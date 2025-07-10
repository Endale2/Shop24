import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import './index.css'
import { useAuthStore } from './stores/auth'
import { nextTick } from 'vue'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)

app.mount('#app')

nextTick(() => {
  const authStore = useAuthStore()
  // On app start, verify session (refresh token and fetch profile)
  authStore.verifySession()

  // Periodic background token refresh every 10 minutes (if user is logged in)
  setInterval(async () => {
    if (authStore.user) {
      try {
        await authStore.refreshToken()
      } catch {
        authStore.user = null
      }
    }
  }, 10 * 60 * 1000)
})
