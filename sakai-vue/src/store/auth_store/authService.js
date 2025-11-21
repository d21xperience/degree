import api from '../api';
const decodeToken = (token) => {
    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return {
            id: payload.user_id,
            username: payload.username || '',
            name: payload.name || '',
            roles: payload.roles || '',
            asalSekolah: payload.asal_sekolah || '',
            sekolah_tenant_id: payload.sekolah_tenant_id || ''
            // exp: payload.exp ? new Date(payload.exp * 1000) : null,
            // raw: payload // simpan semua jika perlu
        };
    } catch (e) {
        console.error('Gagal decode token:', e);
        return null;
    }
};

const state = {
    user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')) : null,
    sekolah: JSON.parse(localStorage.getItem('sekolah')) || null,
    userRole: localStorage.getItem('userRole') ? JSON.parse(localStorage.getItem('userRole')) : null,
    userProfile: JSON.parse(localStorage.getItem('userProfile')) || null, // Ambil dari localStorage
    accessToken: null
};

const mutations = {
    SET_USER(state, user) {
        state.user = user;
        localStorage.setItem('user', JSON.stringify(user)); // Simpan user ke localStorage
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
    SET_TOKEN(state, value) {
        state.accessToken = value;
        localStorage.setItem('accessToken', JSON.stringify(value));
    },
    RESET(state) {
        state.user = null;
        state.sekolah = null;
        localStorage.clear(); // opsional – pilih kunci spesifik
    }
};

const actions = {
    async login({ commit }, credentials) {
        try {
            const { data } = await api.post('/as/auth:login', credentials);
            if (data.accessToken) {
                commit('SET_TOKEN', data.accessToken);
                const userClaims = decodeToken(data.accessToken);
                commit('SET_USER', {
                    username: userClaims.username,
                    asalSekolah: userClaims.asalSekolah
                });
                commit('SET_USER_ROLE', userClaims.roles[0]);
                // user.value = { id: payload.user_id };
                console.log(userClaims);
                // console.log('Id:', userClaims.id, 'Role:', userClaims.roles, 'Asal sekolah:', userClaims.asalSekolah);
                // const user = await api.get('/as/users/');
                // console.log(user);
                return {
                    status: true,
                    userRole: userClaims.roles[0],
                    sekolahTenant: {
                        namaSekolah: userClaims.asalSekolah,
                        sekolahTenantId: userClaims.sekolah_tenant_id
                    }
                };
            }
            return { status: false };
            // const { status, user, sekolahTenant } = data;
            // if (status) {
            //     commit('SET_USER', user);

            //     commit('SET_SEKOLAH', sekolahTenant);
            //     return {
            //         status: true,
            //         userRole: user.role,
            //         user,
            //         sekolahTenant
            //     };
            // }
        } catch (error) {
            alert('error');
            console.log(error);
            // throw error.response?.data || error;
            throw error;
        }
    },
    async refreshToken() {
        try {
            const { data } = await api.post('/as/auth/web/refresh');
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
    async bootstrap({ commit }) {
        try {
            const data = await actions.me();
            if (data.status) {
                commit('SET_USER', data.user);
                commit('SET_USER_ROLE', data.user.role);
                commit('SET_SEKOLAH', data.sekolahTenant);
            }
            return true;
        } catch (_) {
            commit('RESET'); // clear user if not logged in
            return false;
        }
    },
    async logout({ commit }) {
        try {
            // await api.post('/as/auth/web/logout');
            localStorage.removeItem('accessToken');
        } finally {
            commit('CLEAR_AUTH');
            commit('RESET');
        }
    },
    async registerAdmin({ commit }, payload) {
        console.log(payload);
        try {
            const response = await api.post('/as/auth/register', payload);
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
            console.log(error);
            throw error.response.data;
        }
    },
    // Fitur baru ceknpsn
    async ceknpsn({ commit }, npsn) {
        // console.log(npsn);
        try {
            const { data } = await api.get(`/as/sekolah`, {
                params: {
                    npsn: npsn
                }
            });
            // console.log(data);
            if (data.status) {
                commit('SET_SEKOLAH', data.sekolahData);
            }
            return data; // Mengembalikan data sekolah
        } catch (error) {
            // console.log(error);
            throw new Error('Sekolah tidak ditemukan', error);
        }
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
    isAuthenticated: (state) => !!state.user,
    userRole: (state) => state.userRole,
    currentUser: (state) => state.user,
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
