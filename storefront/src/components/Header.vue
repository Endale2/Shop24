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
      <router-link :to="`/shops/${shopSlug}`">Home</router-link> |
      <router-link :to="`/shops/${shopSlug}/products`">Products</router-link> |
      <router-link :to="`/shops/${shopSlug}/collections`">Collections</router-link> |
      <router-link :to="`/shops/${shopSlug}/auth`">Login/Register</router-link>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import { useRoute } from 'vue-router'

interface Props {
  shop: { name: string; image?: string } | null
}
defineProps<Props>()

const route = useRoute()
const shopSlug = route.params.shopSlug as string
</script>
