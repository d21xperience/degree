// src/router/guards/authGuard.js
import store from '@/store';

export default async function authGuard(to, from, next) {
    console.log(store.getters['authService/isAuthenticated']);
    if (store.getters['authService/isAuthenticated']) {
        return next();
    }

    try {
        await store.dispatch('authService/refreshToken');
        return next();
    } catch {
        return next('/auth/login');
    }
}
