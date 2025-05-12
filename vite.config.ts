import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));

// https://vite.dev/config/
export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        party: resolve(__dirname, 'frontend/party/index.html'),
        admin: resolve(__dirname, 'frontend/admin/index.html'),
      },
    },
  },
  plugins: [svelte(), tailwindcss()],
});
