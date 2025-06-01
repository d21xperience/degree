import { useSekolahService } from '@/composables/useSekolahService';
import router from '@/router';
import { computed } from 'vue';
import { useStore } from 'vuex';
const isEmail = (input) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(input);
};

const isValidUsername = (input) => {
    const usernameRegex = /^[a-zA-Z0-9_-]{3,20}$/;
    return usernameRegex.test(input);
};

export function useAuth() {
    const { fetchSekolah } = useSekolahService();
    const store = useStore();
    const currentUser = computed(() => store.getters['authService/getUserProfile']);
    const onLogin = async ({ values }) => {
        const { username, password } = values;
        // cek apakah username merupakan username atau email
        let loginIdentifier = 'username';

        if (isEmail(username)) {
            loginIdentifier = 'email';
        } else if (!isValidUsername(username)) {
            alert('Username atau email tidak valid.');
            return;
        }
        let cek = {};
        try {
            const response = await store.dispatch('authService/login', {
                [loginIdentifier]: username,
                password
            });
            if (response.status) {
                await store.dispatch('sekolahService/fetchTabeltenant', response?.user.sekolahTenantId);
                await fetchSekolah();

                // Ambil tahun ajaran
                await store.dispatch('sekolahService/fetchTahunAjaran');
                await store.dispatch('sekolahService/fetchSemester');

                const namaSekolah = response?.sekolahTenant.namaSekolah.toLowerCase().replace(/\s+/g, '');

                await router.push({ name: 'dashboard', params: { sekolah: namaSekolah } });
            }
        } catch (error) {
            console.error('Login gagal:', error);
            alert('Login gagal. Silakan periksa kembali informasi Anda.');
            store.dispatch('authService/logout');
            return; // pastikan keluar supaya finally tidak berjalan seolah login sukses
        }
    };

    const onLogout = async () => {
        await store.dispatch('authService/logout');
        await store.dispatch('sekolahService/resetState');
        router.push({ name: 'landing' });
    };
    return {
        onLogin,
        onLogout,
        currentUser
    };
}
