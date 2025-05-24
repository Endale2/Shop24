<template>
  <div>
    <ProductList :products="products" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { productService } from '@/services/product'
import ProductList from '@/components/storefront/ProductList.vue'

const products = ref([])
const shopId = window.location.hostname.split('.')[0]

onMounted(async () => {
  products.value = await productService.fetchPublicByShop(shopId)
})
</script>
