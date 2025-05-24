<template>
  <header v-if="shop">
    <h1>{{ shop.name }}</h1>
    <p>{{ shop.description }}</p>
  </header>
</template>

<script setup>
defineProps({ shop: Object })
</script>
