<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 px-4 py-12">
        <div class="max-w-md w-full bg-white rounded-xl shadow-md p-8 text-center ring-1 ring-gray-200">
            <!-- Icon Sukses -->
            <div class="mb-6">
                <div class="mx-auto flex items-center justify-center h-16 w-16 rounded-full bg-green-100">
                    <svg class="h-8 w-8 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                </div>
            </div>

            <!-- Judul & Pesan -->
            <h1 class="text-2xl font-bold text-gray-900 mb-2">Registrasi Berhasil!</h1>
            <p class="text-gray-600 mb-6">
                Akun Anda telah dibuat. Anda akan dialihkan ke halaman login dalam
                <span class="font-semibold text-emerald-600">{{ countdown }} detik</span>.
            </p>

            <!-- Tombol Login Manual -->
            <div class="mb-4">
                <button
                    type="button"
                    class="inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-emerald-600 hover:bg-emerald-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-500 transition-colors w-full"
                    @click="goToLogin"
                >
                    <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13a1 1 0 102 0V9.414l1.293 1.293a1 1 0 001.414-1.414z" clip-rule="evenodd" />
                    </svg>
                    Login Sekarang
                </button>
            </div>

            <!-- Petunjuk Tambahan -->
            <p class="text-sm text-gray-500">Atau tunggu otomatis dialihkan.</p>
        </div>
    </div>
</template>

<script>
import { onMounted, onUnmounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

export default {
    name: 'RegisterSuccess',
    setup() {
        const router = useRouter();
        const countdown = ref(5);
        const redirectTimeout = ref(null);

        // Jalankan countdown saat komponen dimount
        onMounted(() => {
            redirectTimeout.value = setTimeout(() => {
                countdown.value = 4;
                const timer = setInterval(() => {
                    countdown.value--;
                    if (countdown.value <= 0) {
                        clearInterval(timer);
                        goToLogin();
                    }
                }, 1000);
                redirectTimeout.value = timer;
            }, 500); // Delay 0.5s agar animasi lebih natural
        });

        // Cleanup timer saat komponen di-unmount
        onUnmounted(() => {
            if (redirectTimeout.value) {
                clearTimeout(redirectTimeout.value);
                clearInterval(redirectTimeout.value);
            }
        });

        // Watcher opsional (jika ingin logika tambahan saat countdown berubah)
        watch(countdown, (newVal) => {
            if (newVal === 0) {
                goToLogin();
            }
        });

        // Fungsi redirect ke login
        const goToLogin = () => {
            console.log('➡️ Redirecting to login with query...');
            console.log('Target:', { path: '/login', query: { from: 'register-success' } });

            router
                .push({
                    path: '/login',
                    query: { from: 'register-success' }
                })
                .catch((err) => {
                    console.error('Router error:', err); // misal: navigation duplicated
                });
        };

        return {
            countdown,
            goToLogin
        };
    }
};
</script>

<style scoped>
/* Optional: tambahkan transisi lembut */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}
</style>
