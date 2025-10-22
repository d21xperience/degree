/* eslint-disable no-unused-vars */
import api from '../api';

const state = {
    tabelNilaiakhir: JSON.parse(localStorage.getItem('tabelNilaiakhir')) || null
};
const mutations = {
    SET_TABELNILAIAKHIR(state, value) {
        state.tabelNilaiakhir = value;
        localStorage.setItem('tabelNilaiakhir', JSON.stringify(value));
    },
    RESET_STATE(state) {
        state.tabelNilaiakhir = null;
    }
};

const actions = {
    async fetchNilaiSiswa({ commit }, payload) {
        try {
            const response = await api.get(`ss/${payload.schemaname}/nilai-akhir`, {
                params: {
                    semester_id: payload.semesterId,
                    peserta_didik_id: payload.semester_id
                }
            });
            commit('SET_TABELNILAIAKHIR', response.data?.nilaiSiswa);
            return response.data.nilaiSiswa; // Mengembalikan nilai siswa
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal nilai siswa:', error);
            return null;
        }
    },

    async searchNilai({ commit }, payload) {
        try {
            const { data } = await api.get(`ss/${payload.schemaname}/nilai-akhir/search`, {
                params: {
                    semester_id: payload.semesterId,
                    peserta_didik_id: payload.pesertaDidikId
                }
            });
            console.log(data);
            return data;
        } catch (error) {
            console.error(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    }
};
const getters = {
    getNilaiSiswa: (state) => state.tabelNilaiakhir
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
