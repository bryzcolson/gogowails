import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    port: 5173,
  },
  // Ensure Vite can work with Wails runtime
  optimizeDeps: {
    exclude: ['@wailsapp/runtime']
  }
})
