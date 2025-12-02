import api from '../api';
// const decodeToken = (token) => {
//     try {
//         const payload = JSON.parse(atob(token.split('.')[1]));
//         return {
//             id: payload.user_id,
//             username: payload.username || '',
//             name: payload.name || '',
//             roles: payload.roles || '',
//             asalSekolah: payload.asal_sekolah || '',
//             sekolah_tenant_id: payload.sekolah_tenant_id || ''
//             // exp: payload.exp ? new Date(payload.exp * 1000) : null,
//             // raw: payload // simpan semua jika perlu
//         };
//     } catch (e) {
//         console.error('Gagal decode token:', e);
//         return null;
//     }
// };

const state = {
    user: null, // { user_id, email, ... }
    isAuthenticated: false,
    isCheckingAuth: true // untuk menampilkan loading saat cek session
};

const mutations = {
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
        localStorage.clear(); // opsional – pilih kunci spesifik
    }
};

const actions = {
    async checkAuth({ commit }) {
        commit('SET_CHECKING', true);
        try {
            const { data } = await api.get('/as/auth/web/me');
            commit('SET_USER', data);
        } catch (err) {
            // Jika 401 atau error jaringan → anggap tidak login
            commit('CLEAR_USER');
            commit('RESET');
            commit('sekolahService/resetState', null, { root: true });
            commit('semesterService/resetState', null, { root: true });
        } finally {
            commit('SET_CHECKING', false);
        }
    },

    // eslint-disable-next-line no-unused-vars
    async login({ commit, dispatch }, credentials) {
        try {
            const { data } = await api.post('/as/auth/web/login', credentials);
            await dispatch('checkAuth');
            return data;
        } catch (error) {
            console.log(error);
            throw error;
        }
    },
    async refreshToken() {
        try {
            const { data } = await api.post('/as/auth/web/refresh');
            await actions.checkAuth();
            return data;
        } catch (err) {
            console.warn('Refresh token invalid/expired. Logging out...');
            await actions.logout(); // fallback logout
            throw err.response?.data || { message: 'Refresh failed' };
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

    /*—— BOOTSTRAP di App.vue created() ——*/
    // async bootstrap({ commit }) {
    //     try {
    //         const data = await actions.me();
    //         if (data.status) {
    //             commit('SET_USER', data.user);
    //             commit('SET_USER_ROLE', data.user.role);
    //             commit('SET_SEKOLAH', data.sekolahTenant);
    //         }
    //         return true;
    //     } catch (_) {
    //         commit('RESET'); // clear user if not logged in
    //         return false;
    //     }
    // },
    async logout({ commit }) {
        try {
            await api.post('/as/auth/web/logout');
        } finally {
            commit('CLEAR_USER');
            commit('RESET');
            commit('sekolahService/resetState', null, { root: true });
            commit('semesterService/resetState', null, { root: true });
        }
    },
    // eslint-disable-next-line no-unused-vars
    async registerAdmin({ commit }, payload) {
        const response = await api.post('/as/auth:register', payload);
        return response.data;
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
