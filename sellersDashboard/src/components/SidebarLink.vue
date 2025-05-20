<template>
  <router-link
    :to="to"
    class="group flex items-center px-3 py-2 rounded-md text-sm font-medium"
    :class="{
      'bg-gray-200 text-gray-900': isActive,
      'text-gray-600 hover:bg-gray-50 hover:text-gray-900': !isActive,
      'pl-6': small
    }"
  >
    <component
      :is="icon"
      class="mr-3 h-5 w-5 flex-shrink-0"
      :class="{
        'text-gray-500 group-hover:text-gray-700': !isActive,
        'text-gray-900': isActive
      }"
    />
    <span>{{ label }}</span>
  </router-link>
</template>

<script>
export default {
  name: 'SidebarLink',
  props: {
    to: { type: String, required: true },
    icon: { type: String, required: true },
    label: { type: String, required: true },
    small: { type: Boolean, default: false }
  },
  computed: {
    isActive() {
      return this.$route.path === this.to;
    }
  }
};
</script>
