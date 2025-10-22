/* eslint-disable no-unused-vars */
import api from '../api';
// ================================================
// Semester
// ================================================
const state = {
    selectedTahunAjaran: JSON.parse(localStorage.getItem('selectedTahunAjaran')) || [],
    tabelSemester: JSON.parse(localStorage.getItem('tabelSemester')) || null,
    selectedSemester: JSON.parse(localStorage.getItem('selectedSemester')) || null,
    tabelTahunAjaran: JSON.parse(localStorage.getItem('tabelTahunAjaran')) || []
};
const mutations = {
    SET_TABELSEMESTER(state, value) {
        state.tabelSemester = value;
        localStorage.setItem('tabelSemester', JSON.stringify(value));
    },
    SET_TABELTAHUNAJARAN(state, value) {
        state.tabelTahunAjaran = value;
        localStorage.setItem('tabelTahunAjaran', JSON.stringify(value));
    },
    SET_SELECTEDSEMESTER(state, value) {
        state.selectedSemester = value;
        localStorage.setItem('selectedSemester', JSON.stringify(value));
    },
    SET_SELECTEDTAHUNAJARAN(state, value) {
        state.selectedTahunAjaran = value;
        localStorage.setItem('selectedTahunAjaran', JSON.stringify(value));
    },
    RESET_STATE(state) {
        state.tabelSemester = null;
        state.selectedSemester = null;
        state.tabelTahunAjaran = [];
    }
};
const actions = {
    async fetchSemester({ commit }) {
        try {
            const { data, status } = await api.get(`ss/semester`);
            if (status) {
                commit('SET_TABELSEMESTER', data.semester);
            }
            const tahunAjaran = state.selectedTahunAjaran?.tahunAjaranId;
            // const selectedSemester = response.data.semester.reduce((max, item) => (item.semesterId > max.semesterId ? item : max), response.data.semester[0]);
            // if (!state.selectedSemester) {
            //     const selectedSemester = data.semester.filter((item) => item.tahunAjaranId == tahunAjaran);
            //     commit('SET_SELECTEDSEMESTER', selectedSemester[0]);
            // }
            return data;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mendapatkan semester: ${error}`);
        }
    },

    async deleteSemester({ commit }, payload) {
        try {
            console.log({ commit }, payload);
            const { data } = await api.delete('ss/semester', {
                params: {
                    semester_ids: payload
                }
            });
            return data;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },

    async updateSemester({ commit }, payload) {
        try {
            // console.log("sekolalhService",payload)
            // return
            const { data } = await api.put('ss/semester', { semester: payload });
            console.log(data);
            return data;
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    async fetchSelectedSemester({ commit }, payload) {
        // commit('SET_SELECTEDSEMESTER', payload);
    },

    // ================================================
    // Tahun Ajaran
    // ================================================
    async fetchTahunAjaran({ commit }, tahun_ajaran_id) {
        try {
            const response = await api.get(`/ss/tahun-ajaran`, {
                params: {
                    tahun_ajaran_id: tahun_ajaran_id
                }
            });
            commit('SET_TABELTAHUNAJARAN', response.data.tahunAjaran);
            const selectedTahunAjaran = response.data.tahunAjaran.reduce((max, item) => (item.tahunAjaranId > max.tahunAjaranId ? item : max), response.data.tahunAjaran[0]);
            commit('SET_SELECTEDTAHUNAJARAN', selectedTahunAjaran);
            return response.data;
        } catch (error) {
            throw new Error(`Gagal mendapatkan tahun ajaran: ${error}`);
        }
    },

    async fetchSelectedTahunAjaran({ commit }, payload) {
        commit('SET_SELECTEDTAHUNAJARAN', payload);
    }
};

const getters = {
    getSemester: (state) => state.tabelSemester,
    getSelectedSemester: (state) => state.selectedSemester,
    getTahunAjaran: (state) => state.tabelTahunAjaran,
    getSelectedTahunAjaran: (state) => state.selectedTahunAjaran
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
