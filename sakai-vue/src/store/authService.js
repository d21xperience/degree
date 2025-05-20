import axios from 'axios';
const api = axios.create({
    baseURL: import.meta.env.VITE_API_AUTH_BASE_URL, //'http://localhost:8182/api/v1',
    withCredentials: true, // Untuk mengirim cookie atau credensial
    headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer your_token'
    }
});

const state = {
    token: localStorage.getItem('token') || null,
    userRole: localStorage.getItem('userRole') || null,
    user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')) : null,
    userProfile: JSON.parse(localStorage.getItem('userProfile')) || null, // Ambil dari localStorage
    sekolah: JSON.parse(localStorage.getItem('sekolah')) || null, // Ambil dari localStorage
    refreshToken: null
};

const mutations = {
    setUser(state, user) {
        state.user = user;
        localStorage.setItem('user', JSON.stringify(user)); // Simpan user ke localStorage
    },
    setUserRole(state, userRole) {
        state.userRole = userRole;
        localStorage.setItem('userRole', userRole);
    },
    setUserProfile(state, userProfile) {
        state.userProfile = userProfile;
        localStorage.setItem('userProfile', JSON.stringify(userProfile));
    },

    setToken(state, token) {
        state.token = token;
        localStorage.setItem('token', token);
    },
    setRefreshToken(state, refreshToken) {
        state.refreshToken = refreshToken;
    },
    clearAuthData(state) {
        state.token = null;
        state.refreshToken = null;
        state.user = null;
        localStorage.removeItem('token'); // Hapus token saat logout
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
            const response = await api.post('/auth/login', payload);
            const { ok } = response.data;
            if (ok) {
                commit('setToken', response.data.token);
                commit('setUser', response.data.user);
                commit('setUserRole', response.data.user.role);
                commit('SET_SEKOLAH', response.data.sekolahTenant);
                const results = {
                    user: response.data.user,
                    sekolahTenant: response.data.sekolahTenant
                };
                return results; // Indikasi login berhasil
            } else {
                throw new Error(response.data.message || 'Login gagal');
            }
        } catch (error) {
            // console.log(error.response);
            throw error.response.data;
        }
    },
    async logout({ commit }) {
        try {
            commit('clearAuthData');
        } finally {
            commit('clearAuthData');
        }
    },
    async registerAdmin({ commit }, payload) {
        try {
            const response = await api.post('/auth/register', payload);
            if (response.data.ok) {
                commit('setToken', response.data.token);
                // Simpan informasi pengguna setelah login
                commit('setUser', response.data.user);
                commit('setUserRole', response.data.user.role);
            }
            console.log('from Register:', response.data);
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
            commit('setUserProfile', response.data.userProfile);
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
                commit('setUserProfile', response.data.user_profile);
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

            commit('setUser', response.data);
            return response.data;
        } catch (error) {
            console.error('Gagal mengunggah foto profil:', error);
            return null;
        }
    }
};

const getters = {
    isAuthenticated: (state) => {
        if (state.token) {
            return true;
        }
        return false;
    },
    userRole(state) {
        return state.userRole;
    },
    getSekolah(state) {
        return state.sekolah;
    },
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
