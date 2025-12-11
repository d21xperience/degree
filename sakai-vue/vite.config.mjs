import { PrimeVueResolver } from '@primevue/auto-import-resolver';
import vue from '@vitejs/plugin-vue';
import { fileURLToPath, URL } from 'node:url';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';
// https://vitejs.dev/config/
export default defineConfig({
    optimizeDeps: {
        // exclude: ['primevue']
        include: [
            'primevue',
            'chart.js',
            'lodash-es' // karena PrimeVue pakai lodash-es dalam
        ]
    },
    plugins: [
        vue(),
        Components({
            resolvers: [PrimeVueResolver()],
            dts: false
        })
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    // Tambahkan konfigurasi untuk production
    server: {
        host: '0.0.0.0', // Penting untuk Docker
        port: 3000, // Untuk development mode
        strictPort: true
    },
    build: {
        outDir: 'dist',
        assetsDir: 'assets',
        sourcemap: false,
        minify: 'esbuild',
        rollupOptions: {
            output: {
                manualChunks: {
                    vendor: ['vue', 'vue-router', 'vuex'],
                    primevue: ['primevue'],
                    charts: ['chart.js']
                }
            }
        }
    },
    // Base path untuk production - relative path
    base: './'
});
