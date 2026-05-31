import { defineConfig } from "vite";
import path from 'node:path'
import vue from "@vitejs/plugin-vue";
import tailwindcss from '@tailwindcss/vite'
import svgLoader from 'vite-svg-loader'
// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss(), svgLoader()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    host: "0.0.0.0",
    proxy: {
      '/api': 'http://localhost:8080',

    }
  }
})
