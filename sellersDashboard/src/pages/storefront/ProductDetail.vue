<template>
  <div v-if="product">
    <h2>{{ product.name }}</h2>
    <p>{{ product.description }}</p>
    <p>Price: {{ product.price }}</p>
    <img v-for="img in product.images" :src="img" :key="img" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { productService } from '@/services/product'

const route = useRoute()
const product = ref(null)
const shopId = window.location.hostname.split('.')[0]
const productId = route.params.productId

onMounted(async () => {
  product.value = await productService.fetchPublicById(shopId, productId)
})
</script>
