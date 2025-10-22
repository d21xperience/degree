/* eslint-disable no-unused-vars */
import api from '../api';

const state = {
    tabelAnggotaKelas: JSON.parse(localStorage.getItem('tabelAnggotaKelas')) || []
};
const mutations = {
    SET_TABELANGGOTAKELAS(state, value) {
        state.tabelAnggotaKelas = value;
        localStorage.setItem('tabelAnggotaKelas', JSON.stringify(value));
    },
    RESETE_STATE(state) {
        state.tabelAnggotaKelas = [];
    }
};
const actions = {
    // =======================================
    // ANGGOTA KELAS
    // =======================================
    async searchAnggotaKelas({ commit }, payload) {
        try {
            // console.log({commit}, payload);
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/search`, {
                params: {
                    semester_id: payload.semester_id,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            return response.data.anggotaKelas;
        } catch (error) {
            if (error.code === 'ECONNABORTED') {
                console.error('Permintaan terlalu lama, server lambat atau tidak merespons.');
            } else {
                console.error('Terjadi kesalahan:', error.message);
            }
            return null;
        }
    },
    async fetchAnggotaKelas({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas`, {
                params: {
                    semester_id: payload.semester_id,
                    rombongan_belajar_id: payload.rombongan_belajar_id
                }
            });
            return response.data.anggotaKelas;
        } catch (error) {
            console.error('Gagal mendapatkan anggota kelas:', error);
            return null;
        }
    },

    async deleteAnggotaKelas({ commit }, payload) {
        try {
            const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete`, {
                params: {
                    schemaname: payload.schemaname,
                    anggota_rombel_id: payload.anggota_rombel_id
                }
            });
            const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => item.anggotaRombelId != payload.anggota_rombel_id);
            const updateState = {
                semester_id: state.selectedSemester.semesterId,
                peserta_didik: updateAnggotaRombel
            };
            // console.log(updateState);
            commit('SET_TABELSISWAAKTIF', updateState);
            return response.data; // Mengembalikan data
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteBatchAnggotaKelas({ commit }, payload) {
        try {
            const filteredIds = payload.anggota_rombel_id.filter((id) => id && id.trim() !== '');
            const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete-batch`, {
                params: {
                    schemaname: payload.schemaname,
                    anggota_rombel_id: filteredIds
                }
            });
            if (response.data.status) {
                const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => !filteredIds.includes(item.anggotaRombelId));
                const updateState = {
                    semester_id: state.selectedSemester.semesterId,
                    peserta_didik: updateAnggotaRombel
                };
                // console.log(updateState);
                commit('SET_TABELSISWAAKTIF', updateState);
                return response.data; // Mengembalikan data
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async createAnggotaKelas({ commit }, payload) {
        try {
            const response = await api.post(`/ss/${payload.schemaname}/anggota-kelas/delete`, payload);

            // const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => item.anggotaRombelId != payload.anggota_rombel_id);
            // console.log(updateState);
            // commit('SET_TABELSISWAAKTIF', updateState);
            return response.data; // Mengembalikan data
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    }
};
const getters = {
    getTabelAnggotaKelas: (state) => state.tabelAnggotaKelas
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
