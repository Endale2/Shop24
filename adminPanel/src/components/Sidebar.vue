<template>
  <aside
    class="fixed z-30 inset-y-0 left-0 transform transition-transform duration-300 ease-in-out
           lg:static lg:inset-auto lg:translate-x-0
           w-48 bg-gray-800 dark:bg-gray-800 text-white shadow-lg
           flex flex-col"
    :class="{ '-translate-x-full': !isOpen, 'translate-x-0': isOpen }"
  >
    <div class="flex items-center justify-center h-16 bg-gray-900 dark:bg-gray-900 text-xl font-semibold text-white">
      Admin Panel
    </div>

    <nav class="flex-1 p-4 overflow-y-auto custom-scrollbar">
      <ul>
        <li class="mb-2 last:mb-0">
          <router-link
            to="/dashboard"
            class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
            @click="$emit('close-sidebar')" >
            Dashboard
          </router-link>
        </li>
        <li class="mb-2 last:mb-0">
          <router-link
            to="/products"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             @click="$emit('close-sidebar')" >
            Products
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/stores"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             @click="$emit('close-sidebar')" >
            Stores
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/staffs"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             @click="$emit('close-sidebar')" >
            Staffs
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/sellers"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             @click="$emit('close-sidebar')" >
            Sellers
          </router-link>
        </li>
        <li class="mb-2 last:mb-0">
          <router-link
            to="/customers"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             @click="$emit('close-sidebar')" >
            Customers
          </router-link>
        </li>
         </ul>
    </nav>
  </aside>
</template>

<script>
export default {
  props: {
    userRole: String,
    isOpen: {
      type: Boolean,
      required: true
    }
  },
  emits: ['close-sidebar'], // Declare emitted event
};
</script>

<style scoped>
/* Existing styles */
/* ... custom-scrollbar styles ... */
</style>