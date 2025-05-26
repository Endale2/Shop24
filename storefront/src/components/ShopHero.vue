<template>
  <section class="h-72 bg-cover bg-center" :style="`background-image:url(${shop.banner})`">
    <div class="bg-black bg-opacity-50 h-full flex items-center justify-center">
      <div class="text-center text-white">
        <h2 class="text-4xl font-bold">{{ shop.name }}</h2>
        <p class="mt-2">{{ shop.description }}</p>
        <router-link to="/products" class="mt-4 inline-block bg-indigo-600 py-2 px-4 rounded">
          Shop Now
        </router-link>
      </div>
    </div>
  </section>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({ shop: Object })
</script>