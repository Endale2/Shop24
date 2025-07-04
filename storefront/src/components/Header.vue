<template>
  <header class="flex items-center justify-between p-4 shadow">
    <div class="flex items-center space-x-2">
      <img
        v-if="shop?.image"
        :src="shop.image"
        alt="logo"
        class="h-8 w-8 rounded-full object-cover"
      />
      <span class="text-xl font-bold">{{ shop?.name || 'Store' }}</span>
    </div>
    <nav>
      <router-link to="/">Home</router-link> |
      <router-link to="/products">Products</router-link> |
      <router-link to="/collections">Collections</router-link> |
      <router-link to="/auth">Login/Register</router-link>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface Props {
  shop: { name: string; image?: string } | null
}
defineProps<Props>()
</script>
