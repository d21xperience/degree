import api from '../api';

const state = {
    tabelGuru: JSON.parse(localStorage.getItem('tabelGuru')) || null,
    tabelPTKTerdaftar: JSON.parse(localStorage.getItem('tabelPTKTerdaftar')) || null
};
const mutations = {
    SET_TABELGURU(state, tabelGuru) {
        state.tabelGuru = tabelGuru;
        localStorage.setItem('tabelGuru', JSON.stringify(tabelGuru));
    },

    SET_TABELPTKTERDAFTAR(state, value) {
        state.tabelPTKTerdaftar = value;
        localStorage.setItem('tabelPTKTerdaftar', JSON.stringify(value));
    },
    RESET_STATE(state) {
        state.tabelGuru = null;
        state.tabelPTKTerdaftar = null;
    }
};
const actions = {
    // ================================================
    // Service Guru
    // ================================================
    async fetchGuru(payload) {
        try {
            const response = await api.get('/ss/ptk', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_id: payload.ptk_id
                }
            });
            // console.log(response.data);
            // commit("SET_TABELGURU", response.data.PTK);
            return response.data.PTK;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async saveGuru(payload) {
        try {
            const response = await api.post('/ss/ptk/create', payload);
            return response.data.status;
        } catch (error) {
            console.log(error);
        }
    },

    async searchPTKByName(payload) {
        try {
            const response = await api.get('/ss/ptk/search', {
                params: {
                    schemaname: payload.schemaname,
                    nama: payload.nama
                }
            });
            // console.log(response.data)
            return response.data.PTK;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async fetchPTKTerdaftar({ commit }, payload) {
        try {
            const response = await api.get('/ss/ptk-terdaftar', {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahunAjaranId
                }
            });
            const results = {
                tahun_ajaran_id: payload.tahunAjaranId,
                ptkTerdaftar: response.data.ptkTerdaftar
            };
            // console.log(response.data);
            commit('SET_TABELPTKTERDAFTAR', results);
            return results;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deletePTKTerdaftar({ commit }, payload) {
        try {
            const response = await api.delete('/ss/ptk-terdaftar/delete', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_terdaftar_id: payload.ptk_terdaftar_id
                }
            });
            // update tabelPtkTerdaftarId
            if (response.status) {
                console.log(state.tabelPTKTerdaftar);
                const updatePTK = {
                    tahun_ajaran_id: state.selectedSemester?.semesterId,
                    ptkTerdaftar: state.tabelPTKTerdaftar.ptkTerdaftar.filter((item) => item.ptkTerdaftarId != payload.ptk_terdaftar_id)
                };
                console.log(updatePTK);
                commit('SET_TABELPTKTERDAFTAR', updatePTK);
            }
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteBatchPTKTerdaftar({ commit }, payload) {
        try {
            const response = await api.delete('/ss/ptk-terdaftar/delete-batch', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_terdaftar_id: payload.ptk_terdaftar_id
                }
            });
            // update tabelPtkTerdaftarId
            if (response.status) {
                console.log(state.tabelPTKTerdaftar);
                const updatePTK = {
                    tahun_ajaran_id: state.selectedSemester?.semesterId,
                    ptkTerdaftar: state.tabelPTKTerdaftar.ptkTerdaftar.filter((item) => item.ptkTerdaftarId != payload.ptk_terdaftar_id)
                };
                console.log(updatePTK);
                commit('SET_TABELPTKTERDAFTAR', updatePTK);
            }
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async searchPTKTerdaftar(payload) {
        try {
            // console.log('searchGuruTerdaftar', payload);
            const response = await api.get('/ss/ptk-terdaftar/search', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_terdaftar_id: payload.ptk_terdaftar_id
                }
            });
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async addPTKTerdaftar(payload) {
        try {
            console.log(payload);
            const response = await api.post(`/ss/${payload.schemaname}/ptk-terdaftar/create`, payload);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async updatePTKTerdaftar(payload) {
        try {
            console.log(payload);
            // return;
            const response = await api.put(`/ss/ptk-terdaftar/update`, payload);
            // console.log(response);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    }
};
const getters = {
    getGuru: (state) => state.tabelGuru,
    getPTKTerdaftar: (state) => state.tabelPTKTerdaftar
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
