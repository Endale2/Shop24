<!-- src/App.vue -->
<template>
  <!-- Wrap all pages in your DefaultLayout -->
  <DefaultLayout>
    <router-view />
  </DefaultLayout>
</template>

<script setup>
import DefaultLayout from '@/layouts/DefaultLayout.vue'
</script>

<style>
/* Global styles (if any) can go here */
</style>
