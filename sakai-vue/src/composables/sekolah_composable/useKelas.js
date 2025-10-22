import { ref } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
// =================================================
// KELAS
// =================================================
export function useKelas() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();

    const kelasList = ref([]);

    const fetchKelas = async (kelasId = null, tingkatPendidikanId = null) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: initSelectedSemester.value?.semesterId
            };
            if (kelasId) {
                payload.kelas_id = kelasId;
            }
            if (tingkatPendidikanId) {
                payload.tingkat_pendidikan_id = tingkatPendidikanId;
            }
            const response = await store.dispatch('kelasService/fetchKelas', payload);
            return response;
        } catch (error) {
            throw new Error(`Gagal mengambil kelas: ${error.message}`);
        }
    };
    const getKelas = async () => {
        try {
            let response = store.getters['kelasService/getKelas'];
            if (!response || Array.isArray(response.kelas) || response.kelas.length == 0 || initSelectedSemester.value?.semesterId != response?.semesterId) {
                response = await fetchKelas();
            }
            kelasList.value = response.kelas;
            return response;
        } catch (error) {
            throw new Error(`Gagal mendapatkan kelas: ${error.message}`);
        }
    };

    const searchKelas = async (kelasId = null) => {
        try {
            let response = store.getters['kelasService/getKelas'];

            if (!response || response.kelas.length == 0 || initSelectedSemester.value?.semesterId != response?.semesterId) {
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: initSelectedSemester.value?.semesterId
                };
                response = await store.dispatch('kelasService/fetchKelas', payload);
            }
            const result = response.kelas.find((item) => item.rombonganBelajarId == kelasId);
            return result;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian kelas:', error);
        }
    };
    const addKelas = async (kelas, anggotaKelas = null) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kelas: [kelas._rawValue],
                anggota_kelas: anggotaKelas
            };
            console.log(payload);
            const response = await store.dispatch('kelasService/createKelas', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menambhkan kelas:', error);
        }
    };

    /**
     *
     * @param {String} rombonganBelajarId
     * @param {String} semesterId
     * @returns
     */
    const fetchAnggotaKelas = async (rombonganBelajarId = '', semesterId = '') => {
        try {
            const cachedData = await store.getters['siswaService/getSiswaAktif'];
            if (cachedData.semester_id === semesterId) {
                const anggotaKelas = cachedData.peserta_didik.filter((val) => val.rombonganBelajarId === rombonganBelajarId);
                return anggotaKelas;
            }
            return null;
        } catch (error) {
            console.error('Gagal mengambil data kelas:', error);
        }
    };

    return {
        fetchKelas,
        getKelas,
        addKelas,
        searchKelas,
        kelasList,
        fetchAnggotaKelas
    };
}
