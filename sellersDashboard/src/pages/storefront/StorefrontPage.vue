<template>
  <div class="storefront">
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <div v-else>
      <h1>{{ shop?.name || 'Unnamed Shop' }}</h1>
      
      <div v-if="products.length">
        <h2>Products</h2>
        <ul>
          <li v-for="product in products" :key="product.id">
            {{ product.name }} - ${{ product.price }}
          </li>
        </ul>
      </div>
      <div v-else>
        <p>No products found.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute }      from 'vue-router'
import { useShopStore }  from '@/store/shops'
import { productService }from '@/services/product'

const route    = useRoute()
const shops    = useShopStore()
const shopId   = window.location.hostname.split('.')[0]
const shop     = ref(null)
const products = ref([])
const loading  = ref(true)
const error    = ref(null)

onMounted(async () => {
  loading.value = true
  try {
    shop.value = await shops.fetchPublicShop(shopId)
    products.value = await productService.fetchPublicByShop(shopId)
  } catch (e) {
    console.error(e)
    error.value = 'Could not load this storefront.'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.storefront {
  padding: 2rem;
  font-family: sans-serif;
}
</style>
