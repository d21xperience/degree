import api from '../api';
// ================================================
// Semester
// ================================================
const state = {
    selectedTahunAjaran: JSON.parse(localStorage.getItem('selectedTahunAjaran')) || [],
    tabelSemester: JSON.parse(localStorage.getItem('tabelSemester')) || null,
    selectedSemester: JSON.parse(localStorage.getItem('selectedSemester')) || null
};
const mutations = {
    SET_TABELSEMESTER(state, value) {
        state.tabelSemester = value;
        localStorage.setItem('tabelSemester', JSON.stringify(value));
    },
    SET_SELECTEDSEMESTER(state, value) {
        state.selectedSemester = value;
        localStorage.setItem('selectedSemester', JSON.stringify(value));
    },
    RESET_STATE(state) {
        state.tabelSemester = null;
        state.selectedSemester = null;
    }
};
const actions = {
    async fetchSemester({ commit }, semester_id) {
        try {
            const response = await api.get(`/ss/semester`, {
                params: {
                    semester_id: semester_id
                }
            });
            // console.log('sekolahService', response);
            commit('SET_TABELSEMESTER', response.data.semester);
            const tahunAjaran = state.selectedTahunAjaran?.tahunAjaranId;
            // const selectedSemester = response.data.semester.reduce((max, item) => (item.semesterId > max.semesterId ? item : max), response.data.semester[0]);
            if (!state.selectedSemester) {
                const selectedSemester = response.data.semester.filter((item) => item.tahunAjaranId == tahunAjaran);
                commit('SET_SELECTEDSEMESTER', selectedSemester[0]);
            }
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan semester:', error);
        }
    },

    async deleteSemester(payload) {
        try {
            console.log(payload);
            const { data } = await api.delete('ss/semester', {
                params: {
                    semester_ids: payload
                }
            });
            return data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async updateSemester(payload) {
        try {
            // console.log("sekolalhService",payload)
            // return
            const { data } = await api.put('ss/semester', { semester: payload });
            console.log(data);
            return data;
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async fetchSelectedSemester({ commit }, payload) {
        commit('SET_SELECTEDSEMESTER', payload);
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
