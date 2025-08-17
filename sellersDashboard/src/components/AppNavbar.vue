<template>
  <header class="w-full sticky top-0 z-50 bg-white/90 backdrop-blur-xl border-b border-gray-200/50 shadow-sm transition-all duration-300 ease-in-out">
    <!-- Main Navigation -->
    <div class="max-w-7xl mx-auto flex items-center justify-between p-4 lg:px-6">
      <!-- Logo Section -->
      <router-link to="/" class="flex items-center space-x-3 group outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 rounded-lg">
        <div class="relative">
          <div class="w-12 h-12 bg-gradient-to-br from-emerald-500 via-green-500 to-emerald-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-500/30 group-hover:shadow-emerald-500/50 transition-all duration-300 group-hover:scale-110 group-hover:rotate-3">
            <svg class="w-7 h-7 text-white group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
          </div>
        </div>
        <div class="flex flex-col">
          <span class="text-2xl font-bold text-gray-800 tracking-tight group-hover:text-gray-900 transition-colors duration-200">
            Shop<span class="text-amber-600 group-hover:text-amber-700">24</span>
          </span>
          <span class="text-xs text-gray-500 font-medium -mt-1 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
            Power your business
          </span>
        </div>
      </router-link>

      <!-- Right Side Actions -->
      <div class="flex items-center space-x-3">
        <!-- Notifications (for authenticated users) -->
        <div v-if="user && user.id" class="relative group">
          <button class="relative p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2">
            <svg class="h-5 w-5 text-gray-600 group-hover:text-gray-800 transition-colors duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM10.5 3.75a6 6 0 00-6 6v3.75a6 6 0 01-6 6h.75a6 6 0 006 6h3.75a6 6 0 006-6v-3.75a6 6 0 00-6-6h-.75z" />
            </svg>
            <!-- Notification Badge -->
            <span v-if="notificationCount > 0" class="absolute -top-1 -right-1 bg-gradient-to-r from-red-500 to-pink-500 text-white text-xs font-bold w-4 h-4 rounded-full flex items-center justify-center ring-2 ring-white animate-pulse">
              {{ notificationCount > 9 ? '9+' : notificationCount }}
            </span>
          </button>
          
          <!-- Notifications Dropdown -->
          <div class="absolute right-0 mt-2 w-72 bg-white rounded-xl shadow-xl py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-300 transform scale-95 group-hover:scale-100 origin-top-right border border-gray-100">
            <div class="px-4 py-2 border-b border-gray-100">
              <h3 class="text-sm font-semibold text-gray-800">Notifications</h3>
              <p class="text-xs text-gray-500">{{ notificationCount }} new notifications</p>
            </div>
            <div class="max-h-56 overflow-y-auto">
              <div v-for="notification in notifications" :key="notification.id" class="px-4 py-2 hover:bg-gray-50 transition-colors duration-150 cursor-pointer">
                <div class="flex items-start space-x-3">
                  <div class="w-2 h-2 bg-green-500 rounded-full mt-2 flex-shrink-0"></div>
                  <div class="flex-1">
                    <p class="text-sm font-medium text-gray-800">{{ notification.title }}</p>
                    <p class="text-xs text-gray-500 mt-1">{{ notification.message }}</p>
                    <p class="text-xs text-gray-400 mt-1">{{ formatTime(notification.time) }}</p>
                  </div>
                </div>
              </div>
            </div>
            <div class="px-4 py-2 border-t border-gray-100">
              <button class="w-full text-center text-xs text-green-600 hover:text-green-700 font-medium">
                View all notifications
              </button>
            </div>
          </div>
        </div>

        <!-- User Menu -->
        <template v-if="user && user.id">
          <div class="relative group">
            <button class="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-100 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2">
              <div class="relative">
                <img
                  v-if="user.profile_image"
                  :src="user.profile_image"
                  alt="Profile"
                  class="h-8 w-8 rounded-full object-cover ring-2 ring-green-400 group-hover:ring-green-600 transition-all duration-200 shadow-md"
                />
                <div
                  v-else
                  class="h-8 w-8 bg-gradient-to-br from-green-500 to-green-700 rounded-full flex items-center justify-center text-white font-bold text-sm ring-2 ring-green-400 group-hover:ring-green-600 transition-all duration-200 shadow-md"
                >
                  {{ userInitials }}
                </div>
                <!-- Online Status Indicator -->
                <div class="absolute -bottom-1 -right-1 w-3 h-3 bg-green-500 rounded-full border-2 border-white"></div>
              </div>
              <div class="hidden md:block text-left">
                <p class="text-sm font-medium text-gray-800">{{ user.name || user.email?.split('@')[0] }}</p>
                <p class="text-xs text-gray-500">Seller</p>
              </div>
              <svg class="h-4 w-4 text-gray-400 group-hover:text-gray-600 transition-colors duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
            
            <!-- User Dropdown Menu -->
            <div class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-xl py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-300 transform scale-95 group-hover:scale-100 origin-top-right border border-gray-100">
              <!-- User Info Header -->
              <div class="px-4 py-3 bg-gradient-to-r from-gray-50 to-gray-100 border-b border-gray-200">
                <p class="text-sm font-medium text-gray-800">{{ user.email }}</p>
                <p class="text-xs text-gray-500 mt-1">Seller Account</p>
              </div>
              
              <!-- Menu Items -->
              <div class="py-2">
                <router-link
                  to="/profile"
                  class="flex items-center px-4 py-2 text-gray-700 hover:bg-green-50 hover:text-green-800 transition-colors duration-150 group"
                >
                  <svg class="h-4 w-4 mr-3 text-gray-400 group-hover:text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                  <span class="text-sm">Profile Settings</span>
                </router-link>
                
                <router-link
                  to="/settings"
                  class="flex items-center px-4 py-2 text-gray-700 hover:bg-green-50 hover:text-green-800 transition-colors duration-150 group"
                >
                  <svg class="h-4 w-4 mr-3 text-gray-400 group-hover:text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  <span class="text-sm">Account Settings</span>
                </router-link>
                
                <div class="border-t border-gray-100 my-2"></div>
                
                <button
                  @click="handleLogout"
                  class="w-full flex items-center px-4 py-2 text-red-600 hover:bg-red-50 hover:text-red-700 transition-colors duration-150 group"
                >
                  <svg class="h-4 w-4 mr-3 group-hover:scale-110 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                  </svg>
                  <span class="text-sm">Sign Out</span>
                </button>
              </div>
            </div>
          </div>
        </template>
        
        <!-- Auth Buttons (for unauthenticated users) -->
        <template v-else>
          <router-link 
            to="/auth" 
            class="px-5 py-2 bg-gradient-to-r from-green-600 to-green-700 text-white font-medium rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-0.5 active:translate-y-0 active:shadow-md focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 text-sm"
          >
            Sign In
          </router-link>
        </template>
      </div>
    </div>
  </header>
</template>

<script setup>
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/store/auth';

const router = useRouter();
const auth = useAuthStore();

// Reactive data
const notificationCount = ref(3);

// Mock notifications data
const notifications = ref([
  {
    id: 1,
    title: 'New Order Received',
    message: 'Order #12345 has been placed for $299.99',
    time: new Date(Date.now() - 5 * 60 * 1000) // 5 minutes ago
  },
  {
    id: 2,
    title: 'Low Stock Alert',
    message: 'Product "Premium Widget" is running low on stock',
    time: new Date(Date.now() - 15 * 60 * 1000) // 15 minutes ago
  },
  {
    id: 3,
    title: 'Payment Successful',
    message: 'Payment of $150.00 has been processed successfully',
    time: new Date(Date.now() - 30 * 60 * 1000) // 30 minutes ago
  }
]);

const user = computed(() => auth.user || {});

const userInitials = computed(() => {
  if (user.value.name) {
    return user.value.name
      .split(' ')
      .map(n => n[0])
      .join('')
      .toUpperCase();
  }
  const [first, last] = (user.value.email || '').split('@')[0].split('.');
  return ((first?.[0]||'') + (last?.[0]||'')).toUpperCase();
});

// Utility functions
function formatTime(date) {
  const now = new Date();
  const diff = now - date;
  const minutes = Math.floor(diff / (1000 * 60));
  const hours = Math.floor(diff / (1000 * 60 * 60));
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));

  if (minutes < 1) return 'Just now';
  if (minutes < 60) return `${minutes}m ago`;
  if (hours < 24) return `${hours}h ago`;
  return `${days}d ago`;
}

async function handleLogout() {
  await auth.logout();
  router.push({ name: 'Landing' });
}
</script>

<style scoped>
/* Custom animations */
@keyframes ping {
  75%, 100% {
    transform: scale(2);
    opacity: 0;
  }
}

.animate-ping {
  animation: ping 1s cubic-bezier(0, 0, 0.2, 1) infinite;
}

/* Smooth transitions for all interactive elements */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style>
