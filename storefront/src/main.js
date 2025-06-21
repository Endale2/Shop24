// src/main.js
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import App from './App.vue'
import router from './router'
import { useAuthStore } from '@/stores/auth'

const app = createApp(App)
const pinia = createPinia()

// wire up the persist plugin
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(router)

// now get your auth store
const auth = useAuthStore()

// 2) on startup, either hydrate from localStorage *or* re‐fetch from cookie:
// If user is in localStorage, this will auto‐hydrate `auth.user`.
// If you want to re‐validate with the server, you can still call fetchCurrent():
//
auth.$hydrate()       // <-- plugin hook, hydrates from storage
// OR to refresh from server each load:
// await auth.fetchCurrent()

app.mount('#app')
