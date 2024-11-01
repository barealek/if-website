import { defineConfig } from 'astro/config';

import tailwind from '@astrojs/tailwind';

import svelte from '@astrojs/svelte';

// https://astro.build/config
export default defineConfig({
  integrations: [tailwind(), svelte()],
  build: {
    assets: "_if",
  },
  vite: {
    server: {
      proxy: {
        '/api': {
          target: 'http://localhost',

          changeOrigin: true,
        }
      }
    }
  }
});
