import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    tailwindcss()
  ],
  safelist: [
    {
      pattern: /(bg|text|border|ring)-(green|amber|gray|red)-(50|100|200|300|400|500|600|700|800|900)/,
      variants: ['hover', 'focus', 'group-hover', 'sm', 'lg'],
    },
    'shadow-sm', 'shadow-md', 'shadow-lg', 'shadow-xl', 'shadow-2xl', 'shadow-inner',
    'backdrop-blur-md', 'object-cover', 'resize-y',
    'animate-fade-in-up', 'animate-blob', 'animate-scale-in',
    'animation-delay-200', 'animation-delay-400',
    'line-clamp-3',
  ],
 
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    host: 'localhost', // or '0.0.0.0' if you want LAN access
    port: 5174,
    proxy: {
      '/auth': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
