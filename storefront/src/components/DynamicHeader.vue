<template>
  <header 
    class="sticky top-0 z-50 shadow-sm transition-all duration-300"
    :style="headerStyle"
    :class="headerClasses"
  >
    <div :class="containerClasses">
      <div class="flex items-center justify-between" :style="{ height: headerHeight }">
        
        <!-- Logo and Shop Name -->
        <router-link 
          :to="{ path: '/' }" 
          class="flex items-center space-x-3 hover:opacity-80 transition-opacity"
          aria-label="Home"
        >
          <img
            v-if="shop?.image"
            :src="shop.image"
            :alt="shop.name + ' Logo'"
            class="object-contain transition-transform hover:scale-105"
            :class="logoClasses"
          />
          <span 
            class="font-light tracking-tight uppercase transition-colors"
            :style="shopNameStyle"
            :class="shopNameClasses"
          >
            {{ shop?.name || 'Store' }}
          </span>
        </router-link>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex flex-1 items-center justify-end space-x-6">
          
          <!-- Search Bar -->
          <div class="w-full max-w-sm relative">
            <form @submit.prevent="handleSearchSubmit" class="relative">
              <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </span>
              <input
                type="text"
                placeholder="Search products..."
                class="w-full pl-10 pr-10 py-2 border rounded-lg bg-opacity-80 text-sm focus:outline-none focus:ring-2 transition-all duration-300"
                :style="searchInputStyle"
                v-model="searchInput"
                @input="onSearchInput"
                @focus="showSearchResults = true"
              />
              <button
                v-if="searchInput"
                type="button"
                @click="clearSearch"
                class="absolute inset-y-0 right-0 flex items-center pr-3"
              >
                <svg class="w-4 h-4 text-gray-400 hover:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </form>
          </div>

          <!-- Dynamic Navigation Items -->
          <nav class="flex items-center space-x-6 text-base font-light">
            <router-link
              v-for="item in navigation?.items || defaultNavItems"
              :key="item.path"
              :to="{ path: item.path }"
              class="transition-colors flex items-center gap-2"
              :style="navLinkStyle"
              :class="navLinkClasses(item.path)"
              :aria-label="item.label"
            >
              <svg v-if="navigation?.showIcons && item.icon" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="item.icon === 'home'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
                <path v-else-if="item.icon === 'shopping-bag'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l-1 6a2 2 0 01-2 2H8a2 2 0 01-2-2L5 9z"></path>
                <path v-else-if="item.icon === 'collection'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
              </svg>
              <span>{{ item.label }}</span>
            </router-link>
          </nav>

          <!-- User Actions -->
          <div class="flex items-center space-x-4">
            
            <!-- Wishlist -->
            <router-link
              :to="{ path: '/wishlist' }"
              class="p-2 rounded-full transition-all duration-200 hover:scale-110"
              :style="iconButtonStyle"
              aria-label="Wishlist"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
              </svg>
            </router-link>

            <!-- Cart -->
            <router-link
              :to="{ path: '/cart' }"
              class="p-2 rounded-full transition-all duration-200 hover:scale-110 relative"
              :style="iconButtonStyle"
              aria-label="Shopping Cart"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13h10m-10 0L5.5 18M17 13l2.5 5"></path>
              </svg>
              <!-- Cart badge (dynamic based on actual cart state) -->
              <span 
                v-if="cartItemCount > 0"
                class="absolute -top-1 -right-1 rounded-full text-xs font-bold min-w-[18px] h-[18px] flex items-center justify-center"
                :style="badgeStyle"
              >
                {{ cartItemCount }}
              </span>
            </router-link>

            <!-- Account -->
            <router-link
              :to="{ path: '/account' }"
              class="p-2 rounded-full transition-all duration-200 hover:scale-110"
              :style="iconButtonStyle"
              aria-label="Account"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
            </router-link>
          </div>
        </div>

        <!-- Mobile Menu Button -->
        <button
          @click="toggleMobileMenu"
          class="md:hidden p-2 rounded-lg transition-colors"
          :style="mobileMenuButtonStyle"
          aria-label="Toggle menu"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path v-if="!showMobileMenu" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
            <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div 
      v-if="showMobileMenu"
      class="md:hidden border-t transition-all duration-300"
      :style="mobileMenuStyle"
    >
      <div :class="containerClasses">
        <div class="py-4 space-y-3">
          
          <!-- Mobile Search -->
          <div class="relative">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </span>
            <input
              type="text"
              placeholder="Search products..."
              class="w-full pl-10 pr-4 py-2 border rounded-lg text-sm focus:outline-none focus:ring-2 transition-all"
              :style="searchInputStyle"
              v-model="searchInput"
              @input="onSearchInput"
            />
          </div>

          <!-- Mobile Navigation -->
          <nav class="space-y-2">
            <router-link
              v-for="item in navigation?.items || defaultNavItems"
              :key="item.path"
              :to="{ path: item.path }"
              class="block py-2 px-3 rounded-lg transition-colors"
              :style="mobileNavLinkStyle"
              @click="showMobileMenu = false"
            >
              {{ item.label }}
            </router-link>
          </nav>

          <!-- Mobile User Actions -->
          <div class="pt-3 border-t space-y-2" :style="{ borderColor: theme?.colors?.border || '#E5E7EB' }">
            <router-link :to="{ path: '/wishlist' }" class="block py-2 px-3 rounded-lg transition-colors" :style="mobileNavLinkStyle">Wishlist</router-link>
            <router-link :to="{ path: '/cart' }" class="block py-2 px-3 rounded-lg transition-colors" :style="mobileNavLinkStyle">Cart</router-link>
            <router-link :to="{ path: '/account' }" class="block py-2 px-3 rounded-lg transition-colors" :style="mobileNavLinkStyle">Account</router-link>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import type { ShopInfo, DynamicTheme, NavigationConfig, ComponentConfig } from '@/services/dynamic-theme'

// =============== Props ===============

interface Props {
  shop?: ShopInfo | null
  theme?: DynamicTheme | null
  navigation?: NavigationConfig | null
  components?: ComponentConfig | null
}

const props = defineProps<Props>()

// =============== Reactive State ===============

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()
const searchInput = ref('')
const showMobileMenu = ref(false)
const searchResults = ref([])
const showSearchResults = ref(false)

// Get cart item count from store
const cartItemCount = computed(() => cartStore.itemCount)

// =============== Default Navigation ===============

const defaultNavItems = [
  { label: 'Home', path: '/', icon: 'home' },
  { label: 'Products', path: '/products', icon: 'shopping-bag' },
  { label: 'Collections', path: '/collections', icon: 'collection' },
]

// =============== Computed Styles ===============

// Header style based on theme
const headerStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#ffffff', borderBottomColor: '#E5E7EB' }
  
  return {
    backgroundColor: props.theme.colors?.background || '#ffffff',
    borderBottomColor: props.theme.colors?.border || '#E5E7EB',
    backdropFilter: props.components?.navigation?.transparent ? 'blur(10px)' : 'none',
    backgroundColor: props.components?.navigation?.transparent 
      ? `${props.theme.colors?.background || '#ffffff'}CC` 
      : props.theme.colors?.background || '#ffffff'
  }
})

const headerHeight = computed(() => {
  if (!props.theme?.layout) return '80px'
  
  switch (props.theme.layout.headerStyle) {
    case 'compact': return '60px'
    case 'large': return '100px'
    default: return '80px'
  }
})

const headerClasses = computed(() => {
  const classes = ['border-b']
  
  if (props.navigation?.sticky !== false) {
    classes.push('sticky', 'top-0')
  }
  
  return classes
})

const containerClasses = computed(() => {
  const classes = ['mx-auto', 'px-4', 'sm:px-6']
  
  if (props.theme?.layout?.containerWidth === 'full') {
    classes.push('max-w-none')
  } else if (props.theme?.layout?.containerWidth === 'wide') {
    classes.push('max-w-7xl')
  } else {
    classes.push('max-w-6xl')
  }
  
  return classes
})

// Logo styling
const logoClasses = computed(() => {
  switch (props.theme?.layout?.headerStyle) {
    case 'compact': return 'h-8 w-8'
    case 'large': return 'h-12 w-12'
    default: return 'h-10 w-10'
  }
})

// Shop name styling
const shopNameStyle = computed(() => {
  if (!props.theme) return { color: '#1F2937' }
  
  return {
    color: props.theme.colors?.heading || '#1F2937',
    fontFamily: props.theme.fonts?.heading ? `'${props.theme.fonts.heading}', sans-serif` : undefined
  }
})

const shopNameClasses = computed(() => {
  switch (props.theme?.layout?.headerStyle) {
    case 'compact': return 'text-lg'
    case 'large': return 'text-3xl'
    default: return 'text-2xl'
  }
})

// Search input styling
const searchInputStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    borderColor: props.theme.colors?.border || '#E5E7EB',
    backgroundColor: `${props.theme.colors?.background || '#ffffff'}F0`,
    color: props.theme.colors?.bodyText || '#6B7280',
    '--tw-ring-color': props.theme.colors?.primary || '#10B981'
  }
})

// Navigation link styling
const navLinkStyle = computed(() => {
  if (!props.theme) return { color: '#6B7280' }
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280'
  }
})

function navLinkClasses(path: string) {
  const isActive = route.path === path || (path !== '/' && route.path.startsWith(path))
  const classes = ['hover:opacity-80']
  
  if (isActive) {
    classes.push('font-semibold')
    if (props.theme?.colors?.primary) {
      return classes.concat([`text-[${props.theme.colors.primary}]`])
    } else {
      classes.push('text-black')
    }
  }
  
  return classes
}

// Icon button styling
const iconButtonStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    '--hover-bg': `${props.theme.colors?.primary || '#10B981'}20`
  }
})

// Cart badge styling
const badgeStyle = computed(() => {
  if (!props.theme) return { backgroundColor: '#10B981', color: 'white' }
  
  return {
    backgroundColor: props.theme.colors?.primary || '#10B981',
    color: props.theme.colors?.background || 'white'
  }
})

// Mobile menu styling
const mobileMenuButtonStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280'
  }
})

const mobileMenuStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    backgroundColor: props.theme.colors?.background || '#ffffff',
    borderTopColor: props.theme.colors?.border || '#E5E7EB'
  }
})

const mobileNavLinkStyle = computed(() => {
  if (!props.theme) return {}
  
  return {
    color: props.theme.colors?.bodyText || '#6B7280',
    '--hover-bg': `${props.theme.colors?.primary || '#10B981'}10`
  }
})

// =============== Methods ===============

function toggleMobileMenu() {
  showMobileMenu.value = !showMobileMenu.value
}

async function onSearchInput() {
  if (!searchInput.value.trim()) {
    showSearchResults.value = false
    searchResults.value = []
    return
  }

  // Debounced search implementation
  if (searchInput.value.length >= 2) {
    try {
      // TODO: Implement actual search API call
      console.log('Searching for:', searchInput.value)
      showSearchResults.value = true

      // Simulate search results
      searchResults.value = []
    } catch (error) {
      console.error('Search error:', error)
    }
  }
}

function handleSearchSubmit() {
  if (searchInput.value.trim()) {
    router.push({
      path: '/search',
      query: { q: searchInput.value.trim() }
    })
    showMobileMenu.value = false
    showSearchResults.value = false
  }
}

function clearSearch() {
  searchInput.value = ''
  showSearchResults.value = false
  searchResults.value = []
}

// Close mobile menu when clicking outside
function handleClickOutside(event: Event) {
  const target = event.target as Element
  if (!target.closest('.mobile-menu-container')) {
    showMobileMenu.value = false
  }
}

// =============== Watchers ===============

// Close mobile menu on route change
watch(route, () => {
  showMobileMenu.value = false
})

// Add click outside listener for mobile menu
watch(showMobileMenu, (isOpen) => {
  if (isOpen) {
    document.addEventListener('click', handleClickOutside)
  } else {
    document.removeEventListener('click', handleClickOutside)
  }
})
</script>

<style scoped>
/* Custom hover effects */
.hover\:opacity-80:hover {
  opacity: 0.8;
}

/* Dynamic focus rings */
input:focus {
  ring-width: 2px;
  ring-color: var(--tw-ring-color);
}

/* Mobile menu animations */
.mobile-menu-enter-active, .mobile-menu-leave-active {
  transition: all 0.3s ease;
}

.mobile-menu-enter-from, .mobile-menu-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>
