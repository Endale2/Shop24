// src/main.js
import { createApp } from 'vue'
import { createPinia }   from 'pinia'
import piniaPersist      from 'pinia-plugin-persistedstate'
import App               from './App.vue'
import router            from './router'
import { useShopStore }  from '@/store/shops'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)

// extract subdomain (e.g. "shop123" from "shop123.localhost")
const host = window.location.hostname       // e.g. "shop123.localhost"
const subdomain = host.split('.')[0] === 'localhost' 
  ? null 
  : host.split('.')[0]

if (subdomain) {
  // if a subdomain exists, set it as your activeShop in Pinia
  const shops = useShopStore()
  // you may want to fetch all shops first, then pick the one matching `subdomain`
  shops.fetchShops().then(list => {
    const match = list.find(s => s.id === subdomain)
    if (match) shops.setActiveShop(match)
  })
}

app.mount('#app')
