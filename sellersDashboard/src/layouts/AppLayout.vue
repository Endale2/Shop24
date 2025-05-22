<template>
  <div class="flex h-screen bg-gray-100">
    <!-- Mobile hamburger + sidebar -->
    <transition name="slide">
      <aside
        v-if="showSidebar"
        class="fixed inset-0 z-30 bg-black bg-opacity-50 md:hidden"
        @click="showSidebar = false"
      ></aside>
    </transition>

    <Sidebar
      :class="{'translate-x-0': showSidebar, '-translate-x-full': !showSidebar}"
      class="fixed inset-y-0 left-0 w-64 bg-white border-r border-gray-200 transform md:translate-x-0 transition-transform duration-200 ease-in-out z-40"
    />

    <div class="flex-1 flex flex-col overflow-hidden md:pl-64">
      <Navbar @toggle-sidebar="showSidebar = !showSidebar" />
      <main class="flex-1 overflow-auto bg-gray-50 p-6">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Navbar  from '@/components/Navbar.vue'

export default {
  name: 'AppLayout',
  components: { Sidebar, Navbar },
  setup() {
    const showSidebar = ref(false)
    return { showSidebar }
  }
}
</script>
