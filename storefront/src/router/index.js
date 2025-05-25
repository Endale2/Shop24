import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/HomePage.vue'
import ProductPage from '@/pages/ProductPage.vue'

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/products/:id', name: 'Product', component: ProductPage, props: true },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

export default createRouter({
  history: createWebHistory(),
  routes
})
