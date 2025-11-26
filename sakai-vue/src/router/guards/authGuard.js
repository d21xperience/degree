import store from '@/store';
export async function authGuard(to, from, next) {
    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

    if (!requiresAuth) {
        return next(); // route publik, lanjutkan
    }

    if (store.getters['authService/isAuthenticated']) {
        return next(); // sudah login
    }

    // try {
    //     await store.dispatch('authService/refreshToken');
    //     return next(); // berhasil refresh
    // } catch {
    //     console.warn('Token expired or not found. Redirecting to login.');
    return next('/auth/login');
    // }
}

export async function redirectIfAuthenticated(to, from, next) {
    // if (store.getters['authService/isAuthenticated']) {
    //     const nmSekolah = await store.getters['authService/getSekolah']?.namaSekolah;
    //     return next({ name: 'dashboard', params: { sekolah: nmSekolah.toLowerCase().replace(/\s+/g, '') } });
    // }
    // return next();
}

export function isTokenExpired(token) {
    try {
        const payloadBase64 = token.split('.')[1];
        const payload = JSON.parse(atob(payloadBase64));
        return Date.now() >= payload.exp * 1000;
    } catch {
        return true; // invalid token
    }
}
