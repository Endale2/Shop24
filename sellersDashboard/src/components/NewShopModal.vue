<template>
  <Transition name="modal-fade">
    <div v-if="isVisible" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-gray-900 bg-opacity-50 backdrop-blur-sm" @click.self="$emit('close')">
      <div class="bg-white p-8 rounded-2xl shadow-2xl w-full max-w-lg border border-green-200 animate-scale-in" @click.stop>
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-3xl font-bold text-green-800">Launch a New Shop</h3>
          <button @click="$emit('close')" class="text-gray-500 hover:text-gray-800 transition-colors duration-200">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="submitForm" class="space-y-6">
          <input
            v-model="name"
            placeholder="Enter your dazzling shop name (Required)"
            required
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500 focus:ring-4 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm text-lg"
          />
          <textarea
            v-model="description"
            rows="4"
            placeholder="Share a brief description of your shop (Optional, but recommended!)"
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500 focus:ring-4 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm text-lg resize-y"
          />
          <input
            v-model="image"
            placeholder="Link to an eye-catching image for your shop (Optional)"
            class="w-full px-5 py-3 border border-gray-300 rounded-xl text-gray-800 placeholder-gray-500 focus:ring-4 focus:ring-green-400 focus:border-green-400 transition-all duration-200 shadow-sm text-lg"
          />
          <button
            type="submit"
            :disabled="loading || !String(name).trim()"
            class="w-full bg-green-700 text-white py-3.5 rounded-xl font-bold text-lg hover:bg-green-800 disabled:opacity-60 disabled:cursor-not-allowed focus:outline-none focus:ring-4 focus:ring-green-500 focus:ring-offset-2 transition-all duration-200 transform hover:scale-[1.01] flex items-center justify-center shadow-lg"
          >
            <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            {{ loading ? 'Creating Shop...' : 'Create Your Shop Now' }}
          </button>
        </form>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  isVisible: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['create-shop', 'close']);

const name = ref('');
const description = ref('');
const image = ref('');

watch(() => props.isVisible, (newVal) => {
  if (!newVal) {
    name.value = '';
    description.value = '';
    image.value = '';
  }
});

function submitForm() {
  if (!String(name.value).trim()) return;
  emit('create-shop', {
    name: name.value,
    description: description.value,
    image: image.value
  });
}
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from, .modal-fade-leave-to {
  opacity: 0;
}

@keyframes scale-in {
  from {
    opacity: 0;
    transform: scale(0.9);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
.animate-scale-in {
  animation: scale-in 0.3s ease-out forwards;
}
</style>