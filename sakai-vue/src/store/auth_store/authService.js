import api from '../api';

const state = {
    user: null,
    isAuthenticated: false,
    isCheckingAuth: true, // untuk menampilkan loading saat cek session
    expiresAt: JSON.parse(localStorage.getItem('expiresAt')) || null,
    expiresIn: JSON.parse(localStorage.getItem('expiresIn')) || null,
    sIdleWarning: false,
    idleCountdown: 0
};

const mutations = {
    SET_TOKEN_EXPIRY(state, { expiresIn, expiresAt }) {
        state.expiresIn = expiresIn; // simpan durasi (detik)
        localStorage.setItem('expiresIn', JSON.stringify(expiresIn));
        state.expiresAt = expiresAt; // simpan timestamp (ms)
        localStorage.setItem('expiresAt', JSON.stringify(expiresAt));
    },
    SET_USER(state, user) {
        state.user = user;
        // localStorage.setItem('user', JSON.stringify(user)); // Simpan user ke localStorage
        state.isAuthenticated = !!user;
    },
    CLEAR_USER(state) {
        // state.user = null;
        // state.isAuthenticated = false;
        // console.log('[Vuex] CLEAR_USER called');
        state.user = null;
        state.isAuthenticated = false;
        // console.log('[Vuex] After clear:', { user: state.user, auth: state.isAuthenticated });
    },
    SET_CHECKING(state, value) {
        state.isCheckingAuth = value;
    },
    // 🔴 Mutations untuk idle
    SET_IDLE_WARNING(state, { isWarning, countdown }) {
        state.isIdleWarning = isWarning;
        state.idleCountdown = countdown;
    },
    SET_USER_ROLE(state, value) {
        state.userRole = value;
        localStorage.setItem('userRole', JSON.stringify(value));
    },
    SET_USER_PROFILE(state, userProfile) {
        state.userProfile = userProfile;
        localStorage.setItem('userProfile', JSON.stringify(userProfile));
    },
    CLEAR_AUTH(state) {
        state.user = null;
        state.sekolah = null;
        state.userRole = null;
        state.userProfile = null;
        localStorage.clear();
    },
    SET_SEKOLAH(state, sekolah) {
        state.sekolah = sekolah;
        localStorage.setItem('sekolah', JSON.stringify(sekolah));
    },

    RESET(state) {
        state.user = null;
        state.sekolah = null;
        state.expiresIn = null; // simpan durasi (detik)
        state.expiresAt = null;
        localStorage.clear(); // opsional – pilih kunci spesifik
    }
};

const actions = {
    async checkAuth({ commit, dispatch }) {
        commit('SET_CHECKING', true);
        try {
            const { data } = await api.get('/as/auth/web/me');
            commit('SET_USER', data);
        } catch (err) {
            // Jika 401 atau error jaringan → anggap tidak login
            // commit('CLEAR_USER');
            // commit('RESET');
            // commit('sekolahService/resetState', null, { root: true });
            // commit('semesterService/resetState', null, { root: true });
            await dispatch.reset;
        } finally {
            commit('SET_CHECKING', false);
        }
    },

    // eslint-disable-next-line no-unused-vars
    async login({ commit, dispatch }, credentials) {
        try {
            const { data } = await api.post('/as/auth/web/login', credentials);
            // Ambil expiresIn DARI RESPON LOGIN
            const expiresIn = data.expires_in || data.expiresIn || 1800; // support snake/kebab

            // Simpan ke state (bisa via mutation khusus, atau langsung)
            commit('SET_TOKEN_EXPIRY', {
                expiresIn,
                expiresAt: Date.now() + expiresIn * 1000
            });
            await dispatch('checkAuth');
            await dispatch('configureIdleDetection');
            return data;
        } catch (error) {
            console.log(error);
            throw error;
        }
    },
    async refreshToken({ commit, dispatch }) {
        try {
            const { data } = await api.post('/as/auth/web/refresh');

            const expiresIn = data.expires_in || data.expiresIn || 1800;
            commit('SET_TOKEN_EXPIRY', {
                expiresIn,
                expiresAt: Date.now() + expiresIn * 1000
            });

            await dispatch('checkAuth');
            await dispatch('configureIdleDetection'); // ✅ perbarui idle timeout

            return data;
        } catch (err) {
            await dispatch('logout');
            throw err;
        }
    },
    async me() {
        try {
            const response = await api.get('/as/auth/web/me');
            console.log(response);

            return response.data;
        } catch (err) {
            throw err.response?.data; //|| { message: 'Unauthorized' };
        }
    },
    configureIdleDetection({ state }) {
        console.log('[DEBUG] isAuthenticated:', state.isAuthenticated);
        console.log('[DEBUG] expiresAt:', state.expiresAt);
        // if (!state.isAuthenticated || !state.expiresAt) return;

        // // 🔑 Hitung idle timeout = 90% dari sisa waktu
        // const now = Date.now();
        // const timeLeft = state.expiresAt - now;
        // if (timeLeft <= 0) return;

        // const idleTimeout = Math.min(
        //     timeLeft * 0.9, // 90% dari sisa waktu
        //     25 * 60 * 1000 // maks 25 menit
        // );

        // // Kirim ke idle detection
        // window.dispatchEvent(
        //     new CustomEvent('idle-configure', {
        //         detail: { timeout: idleTimeout }
        //     })
        // );
        if (!state.isAuthenticated || !state.expiresAt) return;

        const timeLeft = state.expiresAt - Date.now();
        if (timeLeft <= 0) return;

        const timeout = Math.min(timeLeft * 0.9, 25 * 60 * 1000);

        // 🔥 Kirim event ke App.vue
        window.dispatchEvent(
            new CustomEvent('idle-configure', {
                detail: { timeout }
            })
        );
    },
    // 🔴 Action: mulai idle detection
    // eslint-disable-next-line no-unused-vars
    startIdleDetection({ dispatch }) {
        // Hanya mulai jika sudah login
        if (!state.isAuthenticated) return;

        // Gunakan event untuk komunikasi (hindari direct import)
        window.dispatchEvent(new CustomEvent('idle-start'));
    },

    // 🔴 Action: logout karena idle
    async logoutDueToIdle({ dispatch }) {
        await dispatch('logout');
        // Redirect hanya jika di protected route
        const router = window.$router;
        if (router) {
            const route = router.currentRoute.value;
            if (route.meta?.requiresAuth) {
                router.push({ name: 'login', query: { reason: 'idle' } });
            }
        }
    },
    // eslint-disable-next-line no-unused-vars
    async logout({ commit, dispatch }) {
        try {
            await api.post('/as/auth/web/logout');
        } finally {
            await dispatch('reset');
            // commit('CLEAR_USER');
            // commit('RESET');
            // commit('sekolahService/resetState', null, { root: true });
            // commit('semesterService/resetState', null, { root: true });
        }
    },
    // eslint-disable-next-line no-unused-vars
    async registerAdmin({ commit }, payload) {
        const response = await api.post('/as/auth:register', payload);
        return response.data;
    },

    async reset({ commit }) {
        commit('CLEAR_USER');
        commit('RESET');
        commit('sekolahService/resetState', null, { root: true });
        commit('semesterService/resetState', null, { root: true });
        commit('kurikulumService/resetState', null, { root: true });
        commit('siswaService/resetState', null, { root: true });
        commit('kelasService/resetState', null, { root: true });
    },

    // Fitur baru ceknpsn
    // eslint-disable-next-line no-unused-vars
    async ceknpsn({ commit }, npsn) {
        const { data } = await api.get(`/as/sekolah-tenant/npsn`, {
            params: {
                npsn: npsn
            }
        });
        return data; // Mengembalikan data sekolah
    },
    async getSekolahByID({ commit }, sekolahId) {
        try {
            const response = await api.get(`/as/sekolah`, {
                params: {
                    sekolah_id: sekolahId
                }
            });
            commit('SET_SEKOLAH', response.data);
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan id sekolah', error);
        }
    },
    //  Ambil Profil Pengguna
    async getUserProfile({ commit }, userID) {
        console.log(userID);
        try {
            const response = await api.get(`/as/users/${userID}/profile`);
            commit('SET_USER_PROFILE', response.data.userProfile);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan profile user', error);
        }
    },

    // Update Profil Pengguna
    async updateUserProfile({ commit }, updatedProfile) {
        console.log('Mengirim data ke server:', updatedProfile);
        try {
            const response = await api.put(`/as/user/${updatedProfile.userId}/profile`, {
                user_id: updatedProfile.userId, // Sesuai dengan .proto
                user_profile: {
                    // Harus dikirim dalam objek "user_profile"
                    nama: updatedProfile.nama,
                    jk: updatedProfile.jk,
                    phone: updatedProfile.phone,
                    tpt_lahir: updatedProfile.tptLahir,
                    alamat_jalan: updatedProfile.alamatJalan,
                    kota_kab: updatedProfile.kotaKab,
                    prov: updatedProfile.prov,
                    kode_pos: updatedProfile.kodePos,
                    nama_ayah: updatedProfile.namaAyah,
                    nama_ibu: updatedProfile.namaIbu
                }
            });

            if (response.data.status === 'success') {
                commit('SET_USER_PROFILE', response.data.user_profile);
                return response.data;
            } else {
                console.error('Gagal memperbarui profil:', response.data);
                return null;
            }
        } catch (error) {
            console.error('Gagal memperbarui profil pengguna:', error.response?.data || error.message);
            return null;
        }
    },
    // Upload Foto Profil Pengguna
    async uploadUserPhotoProfile({ commit }, file) {
        try {
            const formData = new FormData();
            formData.append('photo', file);

            const response = await api.post('/as/user/profile/photo', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });

            commit('SET_USER', response.data);
            return response.data;
        } catch (error) {
            console.error('Gagal mengunggah foto profil:', error);
            return null;
        }
    }
};

const getters = {
    user: (state) => state.user,
    isAuthenticated: (state) => state.isAuthenticated,
    isCheckingAuth: (state) => state.isCheckingAuth
    // isAuthenticated: (state) => !!state.user,
    // userRole: (state) => state.userRole,
    // currentUser: (state) => state.user,
    // getSekolah: (state) => state.sekolah,
    // getUserProfile(state) {
    //     const userData = { ...state.user, ...state.userProfile };
    //     return userData;
    // }
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
