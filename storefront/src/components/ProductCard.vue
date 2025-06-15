<template>
  <div
    class="bg-white rounded-xl shadow-xl overflow-hidden group border border-gray-100 transition-all duration-300 hover:shadow-2xl hover:scale-103 flex flex-col h-full relative"
    @mouseenter="handleMouseEnter"
    @mouseleave="handleMouseLeave"
  >
    <!-- New Badge (Sale badge removed) -->
    <div v-if="product.isNew" class="absolute top-3 left-3 z-10 flex space-x-1">
      <span class="bg-green-600 text-white text-xs font-bold px-2.5 py-1 rounded-full shadow-md">NEW</span>
    </div>

    <!-- Product Image -->
    <div class="relative w-full aspect-square overflow-hidden bg-gray-50">
      <img
        :src="currentImage"
        :alt="product.name"
        class="w-full h-full object-cover transition-transform duration-300"
        onerror="this.onerror=null;this.src='https://placehold.co/400x400/F0F9FF/1F2937?text=No+Image';"
      />
    </div>

    <!-- Product Details -->
    <div class="p-4 sm:p-6 flex-1 flex flex-col">
      <div class="flex-grow">
        <!-- Category (if available) -->
        <p v-if="product.category?.name" class="text-xs text-gray-500 mb-1 uppercase tracking-wider">{{ product.category.name }}</p>

        <!-- Product Name -->
        <router-link :to="`/products/${product.slug}`" class="block">
          <h3 class="text-xl font-semibold text-gray-900 mb-2 truncate group-hover:text-green-700 transition-colors duration-200" :title="product.name">
            {{ product.name }}
          </h3>
        </router-link>

        <!-- Product Description (truncated) -->
        <p v-if="product.description" class="text-gray-600 text-sm line-clamp-2 mb-4">{{ product.description }}</p>
      </div>

      <!-- Price and Add to Cart Button -->
      <div class="mt-auto pt-2 flex items-center justify-between">
        <p class="text-green-600 text-lg font-bold">
          ${{ product.display_price.toFixed(2) }}
          <span v-if="product.price && product.price > product.display_price" class="text-sm text-gray-500 line-through ml-2">
            ${{ product.price.toFixed(2) }}
          </span>
        </p>

        <!-- Add to Cart Button -->
        <button
          @click.prevent.stop="$emit('add-to-cart', product)"
          class="px-4 py-2 bg-green-600 text-white font-semibold rounded-lg shadow-md hover:bg-green-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 opacity-0 group-hover:opacity-100 group-hover:translate-y-0 transform translate-y-2"
        >
          Add to Cart
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, ref, computed, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps({
  product: {
    type: Object,
    required: true,
    validator: (value) => {
      return (
        typeof value.id === 'string' &&
        typeof value.name === 'string' &&
        typeof value.display_price === 'number' &&
        (typeof value.main_image === 'string' || value.main_image === null || value.main_image === undefined)
      );
    }
  }
});

const emits = defineEmits(['quick-view', 'add-to-cart']);
const router = useRouter();

const isHovered = ref(false);
const hoverImageIndex = ref(0); // Tracks which alternative image is currently displayed
const hoverTimer = ref(null); // Stores the setTimeout ID for the long hover delay
const imageCycleInterval = ref(null); // Stores the setInterval ID for cycling images

const LONG_HOVER_DELAY = 500; // milliseconds before image cycling starts
const CYCLE_INTERVAL = 1500; // milliseconds between image changes during cycling

// Computed property to gather all unique alternative images (excluding main_image)
const alternativeImages = computed(() => {
  const images = new Set();
  const mainImg = props.product.main_image;

  // Add distinct variant images
  if (props.product.variants && props.product.variants.length > 0) {
    props.product.variants.forEach(v => {
      if (v.image && v.image !== mainImg) {
        images.add(v.image);
      }
    });
  }

  // Add distinct general images
  if (props.product.images && props.product.images.length > 0) {
    props.product.images.forEach(img => {
      if (img && img !== mainImg) {
        images.add(img);
      }
    });
  }
  return Array.from(images);
});

// Computed property to determine which image to display
const currentImage = computed(() => {
  // If not hovered, or if hovered but no alternative images, show main image
  if (!isHovered.value || alternativeImages.value.length === 0) {
    return props.product.main_image || 'https://placehold.co/400x400/F0F9FF/1F2937?text=No+Image';
  }

  // If hovered and there are alternative images, cycle them
  const index = hoverImageIndex.value % alternativeImages.value.length;
  return alternativeImages.value[index];
});

const handleMouseEnter = () => {
  isHovered.value = true;
  // Clear any existing timers/intervals to prevent overlaps from previous hovers
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
  hoverImageIndex.value = 0; // Reset index for a fresh hover sequence

  if (alternativeImages.value.length > 0) {
    // Start a timer to initiate cycling after a delay
    hoverTimer.value = setTimeout(() => {
      // Start cycling images after the long hover delay
      imageCycleInterval.value = setInterval(() => {
        hoverImageIndex.value = (hoverImageIndex.value + 1) % alternativeImages.value.length;
      }, CYCLE_INTERVAL);
    }, LONG_HOVER_DELAY);
  }
};

const handleMouseLeave = () => {
  isHovered.value = false;
  // Clear all timers and intervals when mouse leaves
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
  hoverImageIndex.value = 0; // Reset index for next hover
};

// Clean up intervals and timeouts when the component is unmounted
onUnmounted(() => {
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
});

// Function to navigate to product detail page if needed (not directly used by quick-view anymore)
const navigateToProduct = () => {
  router.push(`/products/${props.product.slug}`);
};
</script>

<style scoped>
/* Ensure line-clamp utility works if not natively supported by Tailwind JIT or older versions */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
