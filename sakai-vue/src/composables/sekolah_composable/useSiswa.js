import { ref } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useSiswa() {
    const store = useStore();
    const siswaAktifList = ref([]);
    const { schemaname } = useTableTenant();
    const { initSelectedSemester, initSelectedTahunAjaran } = useSemester();
    const fetchSiswaAktif = async (semesterId = null) => {
        try {
            const requestData = {
                schemaname: schemaname.value,
                semesterId: semesterId || initSelectedSemester.value.semesterId
            };
            const cachedData = await store.getters['siswaService/getSiswaAktif'];
            const shouldFetchNewData = !cachedData || !cachedData?.peserta_didik?.length || cachedData.semester_id !== requestData.semesterId;

            let studentData = cachedData;
            if (shouldFetchNewData) {
                studentData = await store.dispatch('siswaService/fetchSiswaAktif', requestData);
            }

            // Update reactive data
            siswaAktifList.value = studentData.peserta_didik;

            return studentData.peserta_didik;
        } catch (error) {
            console.error('Failed to fetch active students:', error);
            throw error;
        }
    };

    const deleteSiswaAktif = async (anggotaRombelId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelId
            };
            const response = await store.dispatch('siswaService/createAnggotaKelas', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus siswa aktif:', error);
        }
    };
    const deleteBatchSiswaAktif = async (anggotaRombelIds) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelIds
            };
            const response = await store.dispatch('siswaService/deleteBatchAnggotaKelas', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus siswa aktif:', error);
        }
    };
    // const addSiswaAktif = async (anggotaRombelIds) => {
    //     try {
    //         const payload = {
    //             schemaname: schemaname.value,
    //             anggota_rombel_id: anggotaRombelIds
    //         };
    //         const response = await store.dispatch('siswaService/deleteBatchAnggotaKelas', payload);
    //         if (response.status) {
    //             toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
    //         }
    //     } catch (error) {
    //         toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus banyak siswa: ${error}`, life: 3000 });
    //     }
    // };
    const searchSiswaAktif = async (pesertaDidikId) => {
        try {
            const response = await store.getters['siswaService/getSiswaAktif'];
            if (response) {
                const siswa = response.peserta_didik.find((item) => item.pesertaDidikId.includes(pesertaDidikId));
                return siswa;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian siswa aktif:', error);
        }
    };
    const fetchBanyakSiswaByTingkatId = async (tingkatPendidikanId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: `${initSelectedTahunAjaran.value?.tahunAjaranId}2`,
                tingkat_pendidikan_id: tingkatPendidikanId
            };
            const res = await store.dispatch('siswaService/fetchBanyakSiswaByTingkatId', payload);
            // console.log(res);

            return res;
        } catch (error) {
            throw new Error(`Gagal siswa berdasarkan tingkat: ${error.message}`);
        }
    };
    const fetchBanyakSiswaByRombelId = async (rombelId) => {
        // eslint-disable-next-line no-useless-catch
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: `${initSelectedSemester.value?.semesterId}`,
                rombongan_belajar_id: rombelId
            };
            // return
            const res = await store.dispatch('siswaService/fetchBanyakSiswaByRombelId', payload);

            return res;
        } catch (error) {
            throw error;
        }
    };
    return {
        fetchSiswaAktif,
        deleteSiswaAktif,
        deleteBatchSiswaAktif,
        searchSiswaAktif,
        fetchBanyakSiswaByTingkatId,
        fetchBanyakSiswaByRombelId
    };
}
