// import { fileURLToPath, URL } from 'node:url';

// import { PrimeVueResolver } from '@primevue/auto-import-resolver';
// import vue from '@vitejs/plugin-vue';
// import Components from 'unplugin-vue-components/vite';
// import { defineConfig } from 'vite';

// // https://vitejs.dev/config/
// export default defineConfig({
//     optimizeDeps: {
//         noDiscovery: true
//     },
//     plugins: [
//         vue(),
//         Components({
//             resolvers: [PrimeVueResolver()]
//         })
//     ],
//     resolve: {
//         alias: {
//             '@': fileURLToPath(new URL('./src', import.meta.url))
//         }
//     }
//     // server: {
//     //     https: true // Aktifkan HTTPS dengan sertifikat self-signed
//     // }
// });
import { fileURLToPath, URL } from 'node:url';

import { PrimeVueResolver } from '@primevue/auto-import-resolver';
import vue from '@vitejs/plugin-vue';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
    optimizeDeps: {
        noDiscovery: true
    },
    plugins: [
        vue(),
        Components({
            resolvers: [PrimeVueResolver()]
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
        sourcemap: false, // Nonaktifkan sourcemap untuk production
        // Optimasi untuk production
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
