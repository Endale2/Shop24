<template>
  <div class="flex">
    <Sidebar />
    <div class="flex-1">
      <Navbar />
      <main class="p-4"><router-view/></main>
    </div>
  </div>
</template>
<script>
import Sidebar from '@/components/Sidebar.vue';
import Navbar  from '@/components/Navbar.vue';
export default { components: { Sidebar, Navbar } };
</script>
