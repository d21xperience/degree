import { computed, ref } from 'vue';
import { useStore } from 'vuex';
const isEmail = (input) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(input);
};

const isValidUsername = (input) => {
    const usernameRegex = /^[a-zA-Z0-9_-]{3,20}$/;
    return usernameRegex.test(input);
};
const isValidSlug = (slug) => {
    if (typeof slug !== 'string') return false;
    // hanya huruf kecil, angka, dan tanda hubung; panjang 2–50 karakter
    return /^[a-z0-9](?:[a-z0-9-]{0,48}[a-z0-9])?$/.test(slug);
};
export function useAuth() {
    // const { fetchSekolah } = useSekolah();
    const store = useStore();
    const currentUser = computed(() => store.getters['authService/user']);
    const user = store.getters['authService/user'];
    const authError = ref(null);
    const isAuthLoading = ref(false);

    const login = async ({ username, password, rememberMe = false }) => {
        authError.value = null;
        isAuthLoading.value = true;

        // Validasi input
        if (!username || !username.trim()) {
            authError.value = 'Username atau email tidak boleh kosong.';
            isAuthLoading.value = false;
            return { success: false, error: authError.value };
        }
        if (!password || !password.trim()) {
            authError.value = 'Password tidak boleh kosong.';
            isAuthLoading.value = false;
            return { success: false, error: authError.value };
        }

        let loginIdentifier = 'username';
        if (isEmail(username)) {
            loginIdentifier = 'email';
        } else if (!isValidUsername(username)) {
            authError.value = 'Username minimal 3 karakter dan tanpa spasi.';
            isAuthLoading.value = false;
            return { success: false, error: authError.value };
        }

        try {
            const payload = {
                [loginIdentifier]: username.trim(),
                password: password.trim(),
                remember_me: rememberMe
            };

            const response = await store.dispatch('authService/login', payload);
            console.log(response);
            if (response?.status === true) {
                const user = store.getters['authService/user'];

                if (!user) {
                    throw new Error('Data pengguna tidak tersedia setelah login.');
                }

                let redirectRoute;

                if (user.role === 'superadmin') {
                    redirectRoute = { name: 'suDashboard' };
                } else if (user.sekolahSlug && isValidSlug(user.sekolahSlug)) {
                    redirectRoute = { name: 'dashboard', params: { sekolah: user.sekolahSlug } };
                } else {
                    // Fallback: jika sekolahSlug tidak ada/malformed
                    console.warn('SekolahSlug tidak valid atau tidak tersedia:', user.sekolahSlug);
                    redirectRoute = { name: 'dashboard' };
                }

                return { success: true, redirectRoute };
            } else {
                const errorMsg = response?.message || 'Login gagal. Silakan coba lagi.';
                authError.value = errorMsg;
                return { success: false, error: errorMsg };
            }
        } catch (error) {
            let errorMsg = 'Terjadi kesalahan. Silakan periksa koneksi Anda.';

            if (error?.response) {
                const status = error.response.status;
                if (status === 401) {
                    errorMsg = 'Username/email atau password salah.';
                } else if (status === 422) {
                    errorMsg = 'Data tidak valid.';
                } else if (error.response.data?.message) {
                    errorMsg = error.response.data.message;
                }
            } else if (error?.message) {
                errorMsg = error.message;
            }

            authError.value = errorMsg;
            return { success: false, error: errorMsg };
        } finally {
            isAuthLoading.value = false;
        }
    };

    const onLogout = async () => {
        // try {
        await store.dispatch('authService/logout');
        await store.dispatch('sekolahService/resetState');
        // } finally {
        // router.push({ name: 'landing' });
        // }
    };

    const cekSekolahByNPSN = async (npsn) => {
        try {
            const response = await store.dispatch('authService/ceknpsn', npsn);
            if (response) {
                return false;
            }
            return true;
        } catch (e) {
            // console.log(e);
            return true;
        }
    };
    const onRegisterAdmin = async (dataReg) => {
        try {
            const response = await store.dispatch('authService/registerAdmin', dataReg);
            // console.log('onRegisterAdmin =>', response);
            return response;
        } catch (e) {
            // console.log(e.response.data);
            throw e.response.data;
        }
    };
    return {
        login,
        onLogout,
        currentUser,
        user,
        cekSekolahByNPSN,
        onRegisterAdmin,
        authError,
        isAuthLoading
    };
}
