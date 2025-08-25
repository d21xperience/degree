import api from './api';
// ==================================
// IJAZAH SERVICE
// ==================================

const actions = {
    async createProsesIjazah(payload) {
        try {
            const response = await api.post(`ss/ijazah/create`, payload);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus proses Ijazah:', error);
        }
    },
    async fetchProsesIjazah(payload) {
        try {
            const response = await api.get(`ss/proses-ijazah`, {
                params: {
                    schemaname: payload.schemaname,
                    semester_id: payload.tahun_ajaran_id,
                    ijazah_id: payload.ijazah_id
                }
            });
            return response.data.anggotaKelas;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus proses Ijazah:', error);
        }
    },

    async createInfoIjazah(payload) {
        console.log(payload);

        const response = await api.post(`ss/ijazah/seting-ijazah`, payload);
        console.log(response);
    }
};
export default {
    namespaced: true,
    // state,
    // mutations,
    actions
    // getters
};
