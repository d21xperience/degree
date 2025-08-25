import axios from 'axios';

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    withCredentials: true,
    timeout: 10000,
    headers: { 'Content-Type': 'application/json' }
});

let refreshPromise = null;

api.interceptors.response.use(
    (res) => res,
    async (err) => {
        const cfg = err.config || {};
        if (err.response?.status === 401 && !cfg._retry) {
            cfg._retry = true;

            if (!refreshPromise) {
                refreshPromise = store.dispatch('authService/refreshToken').finally(() => {
                    refreshPromise = null;
                });
            }
            try {
                await refreshPromise;
                return api(cfg); // retry original request
            } catch (_) {
                // refresh gagal → logout sudah ditangani di authService.refresh()
                return Promise.reject(err);
            }
        }

        return Promise.reject(err);
    }
);

export default api;
