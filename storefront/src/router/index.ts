import { createRouter, createWebHistory } from 'vue-router'
import StoreLayout from '@/layouts/StoreLayout.vue'
import HomePage from '@/pages/HomePage.vue'
import ProductsPage from '@/pages/ProductsPage.vue'
import ProductDetail from '@/pages/ProductDetail.vue'
import CollectionsPage from '@/pages/CollectionsPage.vue'
import CollectionDetail from '@/pages/CollectionDetail.vue'
import AuthPage from '@/pages/AuthPage.vue'
import ShopSelection from '@/pages/ShopSelection.vue'

const routes = [
  {
    path: '/shops/:shopSlug',
    component: StoreLayout,
    props: true,
    children: [
      { path: '', component: HomePage },
      { path: 'products', component: ProductsPage },
      { path: 'products/:handle', component: ProductDetail, props: true },
      { path: 'collections', component: CollectionsPage },
      {
        path: 'collections/:handle',
        component: CollectionDetail,
        props: true,
      },
      { path: 'auth', component: AuthPage },
    ],
  },
  // Shop selection page
  {
    path: '/',
    component: ShopSelection
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
