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
    // const { fetchSekolah } = useSekolah();
    const store = useStore();
    const currentUser = computed(() => store.getters['authService/getUserProfile']);
    const user = store.getters['authService/userRole'];

    const onLogin = async ({ values }) => {
        const { username, password, rememberMe } = values;
        // cek apakah username merupakan username atau email
        let loginIdentifier = 'username';

        if (isEmail(username)) {
            loginIdentifier = 'email';
        } else if (!isValidUsername(username)) {
            alert('Username atau email tidak boleh kosong.');
            return;
        }
        try {
            const response = await store.dispatch('authService/login', {
                [loginIdentifier]: username,
                password,
                remember_me: rememberMe
            });

            // console.log(response);
            if (response.status) {
                if (response.userRole != 'superadmin') {
                    // await store.dispatch('sekolahService/fetchTabeltenant', response?.user.sekolahTenantId);
                    // await fetchSekolah();
                    // // Ambil tahun ajaran
                    // await store.dispatch('semesterService/fetchTahunAjaran');
                    // await store.dispatch('semesterService/fetchSemester');
                    const namaSekolah = response?.sekolahTenant.namaSekolah.toLowerCase().replace(/\s+/g, '');
                    await router.push({ name: 'dashboard', params: { sekolah: namaSekolah } });
                } else {
                    await router.push({ name: 'suDashboard' });
                }
            }
        } catch (error) {
            // throw new Error('Gagal login, periksa kembali user dan password anda');
            console.error('Login gagal:', error);
            alert('Login gagal. Silakan periksa kembali informasi Anda.');
            // store.dispatch('authService/logout');
            // return; // pastikan keluar supaya finally tidak berjalan seolah login sukses
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
            console.log(response);
            if (response.status) {
                return false;
            }
            return true;
        } catch (e) {
            console.log(e);
            return true;
        }
    };
    const onRegisterAdmin = async (dataReg) => {
        try {
            const response = await store.dispatch('authService/registerAdmin', dataReg);
            console.log('onRegisterAdmin =>', response);
            // if (response.status) {
            //     return response;
            // }
            return false;
        } catch (e) {
            console.log(e);
            return true;
        }
    };
    return {
        onLogin,
        onLogout,
        currentUser,
        user,
        cekSekolahByNPSN,
        onRegisterAdmin
    };
}
