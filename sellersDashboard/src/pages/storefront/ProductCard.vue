<template>
  <router-link :to="`/products/${product.id}`" class="card">
    <img :src="product.images[0]" alt="" />
    <h3>{{ product.name }}</h3>
    <p>{{ product.price }}</p>
  </router-link>
</template>

<script setup>
defineProps({ product: Object })
</script>
