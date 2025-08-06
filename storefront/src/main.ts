import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import './index.css'
import { useAuthStore } from './stores/auth'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)
app.use(router)

app.mount('#app')

// Initialize authentication state on app startup
const authStore = useAuthStore()
authStore.verifySession().catch((error) => {
  console.error('[main.ts] Failed to verify session:', error)
})
