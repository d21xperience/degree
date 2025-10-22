/* eslint-disable no-unused-vars */
import api from '../api';

const state = {
    tabelSiswaAktif: JSON.parse(localStorage.getItem('tabelSiswaAktif')) || null
};
const mutations = {
    SET_TABELSISWAAKTIF(state, tabelSiswaAktif) {
        state.tabelSiswaAktif = tabelSiswaAktif;
        localStorage.setItem('tabelSiswaAktif', JSON.stringify(tabelSiswaAktif));
    },
    RESET_STATE(state) {
        state.tabelSiswaAktif = null;
    }
};
const actions = {
    // ==================================
    // SISWA
    // ==================================
    async createBanyakSiswa({ commit }, payload) {
        try {
            const response = await api.post(`/ss/${payload.schemaname}/siswa/create-banyak`, payload);
            console.log(response.data);
            return response.data;
        } catch (error) {
            return null;
        }
    },
    async fetchSiswa({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/siswa`, {
                params: {
                    page: payload.page,
                    perpage: payload.perpage
                    // schemaname: schemaname,
                }
            });
            // console.log(response);

            return response.data;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async fetchSiswaAktif({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas`, {
                params: {
                    semester_id: payload.semesterId,
                    schemaname: payload.schemaname
                }
            });
            const results = {
                semester_id: payload.semesterId,
                peserta_didik: response.data.anggotaKelas
            };
            commit('SET_TABELSISWAAKTIF', results);
            return results;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    // async searchSiswa({ commit }, payload) {
    //     try {
    //         const response = await api.get(`/ss/${payload.schemaname}/siswa/search`, {
    //             params: {
    //                 nm_siswa: payload.nm_siswa
    //             }
    //         });
    //         // console.log("results",response.data);
    //         return response.data;
    //     } catch (error) {
    //         commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
    //         console.log(error);
    // throw new Error('Gagal menghapus Kategori Mapel:', error);
    //     }
    // },
    async searchSiswaAktifById({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/siswa/search`, {
                params: {
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            // console.log("results",response.data);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async fetchBanyakSiswaByTingkatId({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/filter`, {
                params: {
                    semester_id: payload.semester_id,
                    tingkat_pendidikan_id: payload.tingkat_pendidikan_id
                }
            });
            console.log(response);
            return response.data.anggotaKelas;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async fetchBanyakSiswaByRombelId({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/filter`, {
                params: {
                    semester_id: payload.semester_id,
                    rombongan_belajar_id: payload.rombongan_belajar_id
                }
            });
            console.log(response);
            return response.data.anggotaKelas;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    }
};
const getters = {
    getSiswaAktif: (state) => state.tabelSiswaAktif
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
