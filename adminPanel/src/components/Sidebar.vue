<template>
  <aside class="w-48 bg-gray-800 dark:bg-gray-800 text-white shadow-lg flex flex-col">

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
            >
            Dashboard
          </router-link>
        </li>
        <li class="mb-2 last:mb-0">
          <router-link
            to="/products"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             >
            Products
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/stores"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             >
            Stores
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/staffs"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             >
            Staffs
          </router-link>
        </li>
         <li class="mb-2 last:mb-0">
          <router-link
            to="/sellers"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             >
            Sellers
          </router-link>
        </li>
        <li class="mb-2 last:mb-0">
          <router-link
            to="/customers"
             class="flex items-center py-2 px-3 rounded-md transition duration-150 ease-in-out
                   hover:bg-gray-700 dark:hover:bg-gray-700"
            active-class="bg-gray-700 dark:bg-gray-700 font-bold"
             >
            Customers
          </router-link>
        </li>
         </ul>
    </nav>

    </aside>
</template>

<script>
export default {
  name: 'SidebarStatic', // Explicitly name for clarity
  // Removed props: ['isOpen']
  props: {
    userRole: String, // Keep if needed for styling
  },
  // Removed emits: ['close-sidebar']
};
</script>

<style scoped>
/* Add any specific sidebar styles here if needed, but Tailwind handles most */
/* Removed all responsive/transform related styles */


/* Add custom scrollbar styles if using the custom-scrollbar class */
.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #333; /* Darker track for dark sidebar */
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #555; /* Darker thumb */
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #777;
}

/* No specific dark mode needed for scrollbar in this dark sidebar */
</style>