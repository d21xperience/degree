import axios from 'axios';

// let refreshPromise = null;
const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    withCredentials: true,
    timeout: 10000,
    headers: { 'Content-Type': 'application/json' }
});

// Tambahkan request interceptor untuk logging
api.interceptors.request.use(
    (config) => {
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);
// Interceptor response: tangani 401 → logout & redirect
api.interceptors.response.use(
    (res) => res,
    async (error) => {
        const { status } = error.response || {};
        const router = window.$router;
        if (status === 401 || status === 403) {
            // ✅ Cek apakah route saat ini protected
            const currentRoute = router.currentRoute.value;
            if (currentRoute?.meta?.requiresAuth) {
                window.dispatchEvent(new CustomEvent('unauthorized'));
            }
            // 🌐 Jika di public route → biarkan UI handle (misal: sembunyikan widget "Hi, User")
        }
        return Promise.reject(error);
    }
);

export default api;
