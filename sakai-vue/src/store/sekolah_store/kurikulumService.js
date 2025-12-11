/* eslint-disable no-unused-vars */
import api from '../api';

const state = {
    tabelKurikulum: JSON.parse(localStorage.getItem('tabelKurikulum')) || null
    // tabelKurikulum: JSON.parse(localStorage.getItem('tabelKurikulum')) || null,
};
const mutations = {
    SET_TABELKURIKULUM(state, value) {
        state.tabelKurikulum = value;
        localStorage.setItem('tabelKurikulum', JSON.stringify(value));
    },
    // SET_TABELKURIKULUM(state, value) {
    //     state.tabelKurikulum = value;
    //     localStorage.setItem('tabelKurikulum', JSON.stringify(value));
    // },
    resetState(state) {
        state.tabelKurikulum = null;
    }
};

const actions = {
    /**
     *
     * @param {any} param0 commit
     * @param {String} payload payload
     * @returns
     */
    async fetchKurikulum({ commit }, payload) {
        try {
            const { data } = await api.get('ss/ref/kurikulum', {
                params: {
                    jenjang_pendidikan_id: payload
                }
            });
            // console.log(data);
            if (data.status) {
                commit('SET_TABELKURIKULUM', { jenjangPendidikanId: payload, kurikulum: data?.kurikulum });
                return data;
            }
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal mendapatkan kurikulum:', error);
            throw new Error(error.message);
        }
    }

    // async searchNilai({ commit }, payload) {
    //     try {
    //         const { data } = await api.get(`ss/${payload.schemaname}/nilai-akhir/search`, {
    //             params: {
    //                 semester_id: payload.semesterId,
    //                 peserta_didik_id: payload.pesertaDidikId
    //             }
    //         });
    //         console.log(data);
    //         return data;
    //     } catch (error) {
    //         console.error(error);
    //         throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
    //     }
    // }
    // async fetchKurikulum({ commit }, payload) {
    //     try {
    //         const response = await api.get(`/ss/ref/kurikulum`, {
    //             params: {
    //                 jenjang_pendidikan_id: payload.jenjangPendidikanId,
    //                 jenjang_pendidikan_str: payload.jenjangPendidikanStr
    //             }
    //         });
    //         if (response.data.status) {
    //             commit('SET_TABELKURIKULUM', response.data.kurikulum);
    //             return response.data;
    //         }
    //     } catch (error) {
    //         throw new Error(`Gagal mendapatkan kurikulum: ${error}`);
    //     }
    // },
};
const getters = {
    getKurikulum: (state) => state.tabelKurikulum
    // getKurikulum: (state) => state.tabelKurikulum,
};
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
