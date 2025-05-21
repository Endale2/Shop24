<template>
  <router-link
    :to="to"
    class="group flex items-center px-3 py-2 rounded-md font-medium transition duration-150 ease-in-out"
    :class="{
      'bg-green-50 text-green-700': isActive, /* Light green background, dark green text for active */
      'text-gray-600 hover:bg-gray-100 hover:text-gray-900': !isActive, /* Subtle hover for inactive */
      'pl-6 text-sm py-1': small /* Smaller text and padding for 'small' links */
    }"
  >
    <component
      :is="icon"
      class="mr-3 flex-shrink-0 transition-colors duration-150 ease-in-out"
      :class="{
        'h-6 w-6': !small, /* Default icon size */
        'h-5 w-5': small, /* Smaller icon size for 'small' links */
        'text-green-500 group-hover:text-green-600': isActive, /* Green icon for active and hover */
        'text-gray-400 group-hover:text-gray-600': !isActive /* Gray icon for inactive, slightly darker on hover */
      }"
    />
    <span>{{ label }}</span>
  </router-link>
</template>

<script>
import { defineComponent, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

export default defineComponent({
  name: 'SidebarLink',
  props: {
    to: { type: String, required: true },
    icon: { type: Object, required: true }, // 'Object' type for component
    label: { type: String, required: true },
    small: { type: Boolean, default: false }
  },
  setup(props) {
    const route = useRoute();
    const isActive = ref(false);

    watch(
      () => route.path,
      (newPath) => {
        isActive.value = newPath.startsWith(props.to);
      },
      { immediate: true }
    );

    return {
      isActive
    };
  }
});
</script>