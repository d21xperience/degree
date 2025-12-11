/* eslint-disable no-unused-vars */
import api from '../api';

const state = {
    tabelKelas: JSON.parse(localStorage.getItem('tabelKelas')) || null
};
const mutations = {
    SET_TABELKELAS(state, value) {
        state.tabelKelas = value;
        localStorage.setItem('tabelKelas', JSON.stringify(value));
    },
    resetState(state) {
        state.tabelKelas = null;
    }
};

const actions = {
    // =======================================
    // KELAS
    // =======================================

    async fetchKelas({ commit }, payload) {
        try {
            console.log('----[kelasService]----', payload);
            const response = await api.get(`/ss/${payload.schemaname}/kelas`, {
                params: {
                    semester_id: payload.semester_id,
                    kelas_id: payload.kelas_id,
                    tingkat_pendidikan_id: payload.tingkat_pendidikan_id
                }
            });
            const data = {
                semesterId: payload.semester_id,
                kelas: response.data.kelas
            };
            commit('SET_TABELKELAS', data);
            return response.data;
        } catch (error) {
            console.error('Gagal mengambil kelas:', error);
        }
    },

    async createKelas({ commit }, payload) {
        try {
            const response = await api.post(`/ss/${payload.schemaname}/tambah-kelas`, payload);
            console.log('sekolahService/createKelas', response);
            return response.data;
        } catch (error) {
            console.log(error);
        }
    },
    async editKelas({ commit }, payload) {
        try {
            const response = await api.put(`/ss/${payload.schemaname}/kelas`, payload);
            // console.log(response.data);
            // commit("SET_TABELSEMESTER", response.data.semester);
            return response.data;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal edit kelas:', error);
            return null;
        }
    },
    async deleteKelas() {
        // try {
        //     const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete`, {
        //         params: {
        //             schemaname: payload.schemaname,
        //             anggota_rombel_id: payload.anggota_rombel_id
        //         }
        //     });
        //     const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => item.anggotaRombelId != payload.anggota_rombel_id);
        //     const updateState = {
        //         semester_id: state.selectedSemester.semesterId,
        //         peserta_didik: updateAnggotaRombel
        //     };
        //     // console.log(updateState);
        //     commit('SET_TABELSISWAAKTIF', updateState);
        //     return response.data; // Mengembalikan data
        // } catch (error) {
        //     console.log(error);
        // throw new Error('Gagal menghapus Kategori Mapel:', error);
        // }
    },
    async deleteBatchKelas() {
        // try {
        //     const filteredIds = payload.anggota_rombel_id.filter((id) => id && id.trim() !== '');
        //     const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete-batch`, {
        //         params: {
        //             schemaname: payload.schemaname,
        //             anggota_rombel_id: filteredIds
        //         }
        //     });
        //     if (response.data.status) {
        //         const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => !filteredIds.includes(item.anggotaRombelId));
        //         const updateState = {
        //             semester_id: state.selectedSemester.semesterId,
        //             peserta_didik: updateAnggotaRombel
        //         };
        //         // console.log(updateState);
        //         commit('SET_TABELSISWAAKTIF', updateState);
        //         return response.data; // Mengembalikan data
        //     }
        // } catch (error) {
        //     console.log(error);
        // throw new Error('Gagal menghapus Kategori Mapel:', error);
        // }
    }
};

const getters = {
    getKelas: (state) => state.tabelKelas
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
