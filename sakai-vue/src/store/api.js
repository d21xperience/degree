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
        // console.log(`Making ${config.method?.toUpperCase()} request to: ${config.url}`);
        // return config;
        const token = localStorage.getItem('accessToken');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (res) => res,
    // async (err) => {
    //     console.error('API Error:', err.response?.status, err.config?.url);

    //     const cfg = err.config || {};
    //     if (err.response?.status === 401 && !cfg._retry) {
    //         cfg._retry = true;

    //         if (!refreshPromise) {
    //             refreshPromise = store.dispatch('authService/refreshToken').finally(() => {
    //                 refreshPromise = null;
    //             });
    //         }
    //         try {
    //             await refreshPromise;
    //             return api(cfg);
    //         } catch (_) {
    //             return Promise.reject(err);
    //         }
    //     }

    //     return Promise.reject(err);
    // }
    (error) => {
        if (error.response?.status === 401) {
            // Hapus token & redirect ke login
            localStorage.removeItem('accessToken');
            // opsional: tampilkan notifikasi
            alert('Sesi berakhir. Silakan login kembali.');
            // window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export default api;
