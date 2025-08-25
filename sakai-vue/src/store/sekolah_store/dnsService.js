import api from '../api';

const state = {
    tabelDns: JSON.parse(localStorage.getItem('tabelDns')) || []
};

const mutations = {
    SET_TABELDNS(state, value) {
        state.tabelDns = value;
        localStorage.setItem('tabelDns', JSON.stringify(value));
    },
    RESET_STATE(state) {
        state.tabelDns = [];
    }
};
const actions = {
    // =============================================
    //  DNS
    // =============================================
    async createDns(payload) {
        try {
            const response = await api.post(`ss/ijazah/data-nominasi_sementara`, payload);
            return response.data;
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async updateDns(payload) {
        try {
            const response = await api.put(`ss/ijazah/data-nominasi_sementara/update`, payload);

            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async fetchDns({ commit }, payload) {
        try {
            const response = await api.get(`/ss/data-nominasi_sementara`, {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    is_complete: payload.is_complete
                }
            });
            // console.log(response.data);
            if (response.status) {
                const results = {
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    dataNominasiSementara: response.data.dataNominasiSementara
                };
                commit('SET_TABELDNS', results);
                results.message = response.data.message;
                results.status = response.data.status;
                console.log(results);
                return results;
            }
            return null;
        } catch (error) {
            // console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async searchDns(payload) {
        try {
            const response = await api.get('/ss/data-nominasi_sementara/search', {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteDns({ commit }, payload) {
        try {
            const response = await api.delete(`ss/data-nominasi_sementara`, {
                params: {
                    schemaname: payload.schemaname,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            if (response.status) {
                const cek = state.tabelDns.dataNominasiSementara.filter((item) => item.pesertaDidikId != payload.peserta_didik_id);
                const results = {
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    dataNominasiSementara: cek
                };
                commit('SET_TABELDNS', results);
                results.message = response.data.message;
                results.status = response.data.status;
                // console.log(results);
                return results;
            }
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    }
};

const getters = {
    getDns: (state) => state.tabelDns
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
