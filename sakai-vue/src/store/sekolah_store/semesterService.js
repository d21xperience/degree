/* eslint-disable no-unused-vars */
import api from '../api';
// ================================================
// Semester
// ================================================
const state = {
    selectedTahunAjaran: JSON.parse(localStorage.getItem('selectedTahunAjaran')) || null,
    selectedSemester: JSON.parse(localStorage.getItem('selectedSemester')) || null,
    tabelSemester: JSON.parse(localStorage.getItem('tabelSemester')) || null,
    tabelTahunAjaran: JSON.parse(localStorage.getItem('tabelTahunAjaran')) || null
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
    resetState(state) {
        state.selectedTahunAjaran = null;
        state.selectedSemester = null;
        state.tabelSemester = null;
        state.tabelTahunAjaran = null;
    }
};
const actions = {
    async fetchSemester({ commit }, schemaname = 'ref') {
        try {
            const { data } = await api.get(`ss/semester`, {
                params: {
                    schemaname: schemaname
                }
            });

            const listTahunAjaran = [...new Set(data.semester.map((item) => item.tahunAjaranId))].map((id) => ({ tahunAjaranId: id, label: String(id) }));
            // 1. Ambil yang periodeAktif = 1
            const active = data.semester.filter((item) => item.periodeAktif === 1)[0];

            // 2. Bentuk array baru
            // const selectedSemester = active ? { semesterId: active.semesterId, label: String(active.namaSemester) } : null;
            const selectedSemester = data.semester.find((item) => item.periodeAktif === 1) || null;

            const selectedTahunAjaran = active ? { tahunAjaranId: active.tahunAjaranId, label: String(active.tahunAjaranId) } : null;

            commit('SET_TABELSEMESTER', data.semester);
            commit('SET_TABELTAHUNAJARAN', listTahunAjaran);

            commit('SET_SELECTEDSEMESTER', selectedSemester);
            commit('SET_SELECTEDTAHUNAJARAN', selectedTahunAjaran);
            return {
                status: data.status,
                message: data.message,
                semester: data.semester,
                tahunAjaran: listTahunAjaran
            };
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
    async fetchTahunAjaran({ commit, dispatch }, tahun_ajaran_id) {
        const { data } = await dispatch.fetchSemester();
        return {
            status: data.status,
            message: data.message,
            tahunAjaran: data.tahunAjaran
        };
        // try {
        //     const { data } = await api.get(`/ss/tahun-ajaran`, {
        //         params: {
        //             tahun_ajaran_id: tahun_ajaran_id
        //         }
        //     });
        //     commit('SET_TABELTAHUNAJARAN', data.tahunAjaran);
        //     const selectedTahunAjaran = data.tahunAjaran.reduce((max, item) => (item.tahunAjaranId > max.tahunAjaranId ? item : max), data.tahunAjaran[0]);
        //     commit('SET_SELECTEDTAHUNAJARAN', selectedTahunAjaran);
        //     return data;
        // } catch (error) {
        //     throw new Error(`Gagal mendapatkan tahun ajaran: ${error}`);
        // }
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
