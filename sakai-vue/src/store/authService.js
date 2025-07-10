import axios from 'axios';
const api = axios.create({
    baseURL: import.meta.env.VITE_API_AUTH_BASE_URL,
    withCredentials: true, // penting agar cookie dikirim
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
});

// Optional: response interceptor untuk auto-refresh
api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        if (error.response?.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;
            try {
                await api.post('/auth/web/refresh');
                return api(originalRequest); // retry
            } catch (refreshError) {
                return Promise.reject(refreshError);
            }
        }
        return Promise.reject(error);
    }
);

const state = {
    token: localStorage.getItem('token') || null,
    refreshToken: null,
    userRole: localStorage.getItem('userRole') || null,
    userProfile: JSON.parse(localStorage.getItem('userProfile')) || null, // Ambil dari localStorage
    sekolah: JSON.parse(localStorage.getItem('sekolah')) || null, // Ambil dari localStorage
    user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')) : null,
    isAuthenticated: false
};

const mutations = {
    SET_USER(state, user) {
        state.user = user;
        localStorage.setItem('user', JSON.stringify(user)); // Simpan user ke localStorage
    },
    SET_USER_ROLE(state, userRole) {
        state.userRole = userRole;
        localStorage.setItem('userRole', userRole);
    },
    SET_USER_PROFILE(state, userProfile) {
        state.userProfile = userProfile;
        localStorage.setItem('userProfile', JSON.stringify(userProfile));
    },
    SET_AUTH(state, user) {
        state.isAuthenticated = true;
        state.user = user;
    },
    // setToken(state, token) {
    //     state.token = token;
    //     localStorage.setItem('token', token);
    // },
    // setRefreshToken(state, refreshToken) {
    //     state.refreshToken = refreshToken;
    // },
    CLEAR_AUTH(state) {
        state.isAuthenticated = false;
        state.user = null;
        localStorage.removeItem('userRole'); // Hapus userRole saat logout
        localStorage.removeItem('userProfile'); // Hapus userProfile saat logout
        localStorage.removeItem('user');
        localStorage.removeItem('sekolah');
        localStorage.removeItem('tabelTenant');
        localStorage.removeItem('BCNETWORK');
        localStorage.removeItem('BCACCOUNT');
        localStorage.removeItem('tabelSemester');
        localStorage.removeItem('selectedSemester');
        localStorage.removeItem('tabelGuru');
        localStorage.removeItem('tabelKurikulum');
        localStorage.removeItem('tabelSekolah');
        localStorage.removeItem('tabelSiswa');
        localStorage.removeItem('tabelSiswaAktif');
        localStorage.removeItem('tabelPTKTerdaftar');
        localStorage.removeItem('tabelMapel');
        localStorage.removeItem('tabelKelas');
        localStorage.removeItem('gelarAkademik');
        localStorage.removeItem('tabelTingkatPendidikan');
        localStorage.removeItem('tabelTahunAjaran');
        localStorage.removeItem('selectedTahunAjaran');
        localStorage.removeItem('tabelDashboard');
        localStorage.removeItem('METAMASK_CONNECTED');
        localStorage.removeItem('tabelDns');
        localStorage.removeItem('tabelJurusan');
        localStorage.removeItem('tabelSekolah');
        localStorage.removeItem('CONTRACT');
        localStorage.removeItem('SCIjazah');
        state.token = null;
        state.userRole = null;
        state.user = null;
        state.userProfile = null;
        state.sekolah = null;
        state.refreshToken = null;
    },
    SET_SEKOLAH(state, sekolah) {
        state.sekolah = sekolah;
        localStorage.setItem('sekolah', JSON.stringify(sekolah));
    },
    // SET_LOADING(state, isLoading) {
    //     state.loading = isLoading;
    // },
    SET_ERROR(state, error) {
        state.error = error;
    }
};

const actions = {
    async login({ commit }, payload) {
        try {
            const response = await api.post('/auth/web/login', payload);
            const { status } = response.data;
            if (status) {
                commit('SET_AUTH', response.data.user);
                // commit('setToken', response.data.token);
                commit('SET_USER_ROLE', response.data.user.role);
                commit('SET_USER', response.data.user);
                commit('SET_SEKOLAH', response.data.sekolahTenant);
                const results = {
                    status: true,
                    userRole: response.data.user.role,
                    user: response.data.user,
                    sekolahTenant: response.data.sekolahTenant
                };
                return results; // Indikasi login berhasil
            } else {
                throw new Error(response.data.message || 'Login gagal');
            }
        } catch (error) {
            console.log(error);
            // console.log(error.response);
            throw error.response.data;
        }
    },
    async refreshToken({ commit }) {
        try {
            console.log('refreshToken');
            const res = await api.post('/auth/web/refresh');
            console.log(res);
            // commit('SET_AUTH', res.data.user); // atau abaikan jika hanya refresh token
            return res.data;
        } catch (error) {
            commit('CLEAR_AUTH');
            throw error;
        }
    },

    async logout({ commit }) {
        try {
            const response = await api.post('/auth/web/logout');
            const { status } = response.data;
            if (status) {
                commit('CLEAR_AUTH');
            }
        } finally {
            commit('CLEAR_AUTH');
        }
    },
    async registerAdmin({ commit }, payload) {
        try {
            const response = await api.post('/auth/register', payload);
            console.log('authService/registerAdmin', response);
            if (response.data.ok) {
                commit('setToken', response.data.token);
                // Simpan informasi pengguna setelah login
                commit('SET_USER', response.data.user);
                commit('SET_USER_ROLE', response.data.user.role);
            }
            // console.log('from Register:', response.data);
            return response.data;
        } catch (error) {
            throw error.response.data;
        }
    },
    // Fitur baru ceknpsn
    async ceknpsn({ commit }, npsn) {
        try {
            const response = await api.get(`/sekolah`, {
                params: {
                    npsn: npsn
                }
            });
            commit('SET_SEKOLAH', response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },
    async getSekolahByID({ commit }, sekolahId) {
        try {
            const response = await api.get(`/sekolah`, {
                params: {
                    sekolah_id: sekolahId
                }
            });
            commit('SET_SEKOLAH', response.data);
        } catch (error) {
            throw error;
        }
    },
    //  Ambil Profil Pengguna
    async getUserProfile({ commit }, userID) {
        try {
            const response = await api.get(`/user/${userID}/profile`);
            commit('SET_USER_PROFILE', response.data.userProfile);
            return response.data;
        } catch (error) {
            throw error;
        }
    },

    // Update Profil Pengguna
    async updateUserProfile({ commit }, updatedProfile) {
        console.log('Mengirim data ke server:', updatedProfile);
        try {
            const response = await api.put(`/user/${updatedProfile.userId}/profile`, {
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

            const response = await api.post('/user/profile/photo', formData, {
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
    },
    /*—— BOOTSTRAP di App.vue created() ——*/
    async bootstrap({ commit }) {
        try {
            const { data } = await api.get('/auth/me'); // backend validasi cookie
            commit('SET_USER', data.user);
            commit('SET_SEKOLAH', data.sekolahTenant);
        } catch {
            commit('RESET');
        } // belum login
    }
};

const getters = {
    isAuthenticated: (s) => !!s.user,
    // isAuthenticated: (state) => state.isAuthenticated,
    currentUser: (state) => state.user,
    userRole: (state) => state.userRole,
    getSekolah: (state) => state.sekolah,
    getUserProfile(state) {
        const userData = { ...state.user, ...state.userProfile };
        return userData;
    }
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
