<template>
  <div
    class="bg-white rounded-2xl shadow-lg overflow-hidden group border border-gray-100 transition-all duration-300 hover:shadow-xl hover:scale-[1.02] flex flex-col h-full relative will-change-transform"
    @mouseenter="handleMouseEnter"
    @mouseleave="handleMouseLeave"
  >
    <div v-if="product.isNew" class="absolute top-4 left-4 z-10">
      <span class="bg-green-600 text-white text-xs font-bold px-3 py-1.5 rounded-full shadow-lg transform -rotate-3 origin-bottom-left">NEW</span>
    </div>

    <div class="relative w-full aspect-square overflow-hidden bg-gray-100 flex items-center justify-center p-2"> <img
        :src="currentImage"
        :alt="product.name"
        class="max-w-full max-h-full object-contain transition-transform duration-500 ease-in-out group-hover:scale-105" onerror="this.onerror=null;this.src='https://placehold.co/400x400/F0F9FF/1F2937?text=No+Image';"
      />
      <div class="absolute inset-0 bg-black opacity-0 group-hover:opacity-10 transition-opacity duration-300 pointer-events-none"></div>
    </div>

    <div class="p-4 sm:p-6 flex-1 flex flex-col justify-between">
      <div class="flex-grow">
        <p v-if="product.category?.name" class="text-xs text-gray-500 mb-1 uppercase tracking-widest font-medium">{{ product.category.name }}</p>

        <router-link :to="`/products/${product.slug}`" class="block">
          <h3 class="text-xl font-bold text-gray-900 mb-2 truncate group-hover:text-green-700 transition-colors duration-200" :title="product.name">
            {{ product.name }}
          </h3>
        </router-link>

        <p v-if="product.description" class="text-gray-600 text-sm line-clamp-2 mb-4">{{ product.description }}</p>
      </div>

      <div class="mt-auto pt-4 flex items-end justify-between border-t border-gray-100">
        <p class="text-green-600 text-2xl font-extrabold leading-none">
          ${{ product.display_price.toFixed(2) }}
          <span v-if="product.price && product.price > product.display_price" class="text-sm text-gray-500 line-through ml-2 font-normal">
            ${{ product.price.toFixed(2) }}
          </span>
        </p>

        <button
          v-if="!hasVariants"
          @click.prevent.stop="$emit('add-to-cart', product)"
          class="px-5 py-2.5 bg-green-600 text-white font-semibold rounded-lg shadow-lg hover:bg-green-700 transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2 focus:ring-offset-white opacity-0 group-hover:opacity-100 group-hover:translate-y-0 transform translate-y-3"
          aria-label="Add to cart"
        >
          Add to Cart
        </button>
        <router-link
          v-else
          :to="`/products/${product.slug}`"
          class="px-5 py-2.5 bg-gray-200 text-gray-700 font-semibold rounded-lg shadow-md hover:bg-gray-300 transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 focus:ring-offset-white flex items-center justify-center text-center opacity-0 group-hover:opacity-100 group-hover:translate-y-0 transform translate-y-3"
          aria-label="View product details"
        >
          View Product
        </router-link>
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
const hoverImageIndex = ref(0);
const hoverTimer = ref(null);
const imageCycleInterval = ref(null);

const LONG_HOVER_DELAY = 400;
const CYCLE_INTERVAL = 1800;

// NEW: Computed property to determine if the product has variants
const hasVariants = computed(() => {
  // Assuming a product has variants if `product.variants` array exists and has more than 1 item,
  // or if a property like `product.has_options` is true.
  // Adjust this logic based on your actual product data structure.
  return (props.product.variants && props.product.variants.length > 1) || props.product.has_options;
});


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

  // Add distinct general images (e.g., product gallery images)
  if (props.product.images && props.product.images.length > 0) {
    props.product.images.forEach(img => {
      if (img && img !== mainImg) {
        images.add(img);
      }
    });
  }
  return Array.from(images);
});

const currentImage = computed(() => {
  if (!isHovered.value || alternativeImages.value.length === 0) {
    return props.product.main_image || 'https://placehold.co/400x400/F0F9FF/1F2937?text=No+Image';
  }

  const index = hoverImageIndex.value % alternativeImages.value.length;
  return alternativeImages.value[index];
});

const handleMouseEnter = () => {
  isHovered.value = true;
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
  hoverImageIndex.value = 0;

  if (alternativeImages.value.length > 0) {
    hoverTimer.value = setTimeout(() => {
      if (isHovered.value) {
        imageCycleInterval.value = setInterval(() => {
          hoverImageIndex.value = (hoverImageIndex.value + 1) % alternativeImages.value.length;
        }, CYCLE_INTERVAL);
      }
    }, LONG_HOVER_DELAY);
  }
};

const handleMouseLeave = () => {
  isHovered.value = false;
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
  hoverImageIndex.value = 0;
};

onUnmounted(() => {
  clearTimeout(hoverTimer.value);
  clearInterval(imageCycleInterval.value);
});

const navigateToProduct = () => {
  router.push(`/products/${props.product.slug}`);
};
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>