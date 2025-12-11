import { computed, ref } from 'vue';
import { useStore } from 'vuex';

export function useSiswa() {
    const schemaname = computed(() => {
        return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    });
    const store = useStore();
    const siswaAktifList = ref([]);
    // const schemaname = computed(() => {
    //     return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    // });
    // const initSelectedSemester = computed(store.getters['semesterService/getSelectedSemester']);
    // const initSelectedTahunAjaran = computed(store.getters['semesterService/getSelectedTahunAjaran']);

    // ACTION
    const fetchSiswaAktif = async (semesterId = '') => {
        if (semesterId === '') {
            console.error(new Error('semester_id harus diisi'));
            return;
        }

        try {
            const requestData = {
                schemaname: schemaname.value,
                semesterId: semesterId
            };
            const cachedData = store.getters['siswaService/getSiswaAktif'];
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
                schemaname: schemaname,
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
                schemaname: schemaname,
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
    //             schemaname: schemaname,
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
            const response = store.getters['siswaService/getSiswaAktif'];
            if (response) {
                const siswa = response.peserta_didik.find((item) => item.pesertaDidikId.includes(pesertaDidikId));
                return siswa;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian siswa aktif:', error);
        }
    };
    /**
     *
     * @param {String} rombonganBelajarId Wajib
     * @param {String} semesterId Wajib
     * @returns {Array}
     */
    const getSiswaAktifByKelasId = async (rombonganBelajarId = '', semesterId = '') => {
        if (rombonganBelajarId == '' || semesterId == '') {
            throw new Error('Parameter wajib disi');
        }
        try {
            const cachedData = store.getters['siswaService/getSiswaAktif'];
            if (cachedData) {
                if (cachedData.semester_id === semesterId) {
                    return cachedData.peserta_didik.filter((val) => val.rombonganBelajarId === rombonganBelajarId);
                }
            } else {
                const anggotaKelas = await fetchSiswaAktif(semesterId);
                return anggotaKelas.filter((val) => val.rombonganBelajarId === rombonganBelajarId);
            }
        } catch (error) {
            console.error('Gagal mengambil data kelas:', error);
        }
    };

    const fetchBanyakSiswaByTingkatId = async (tingkatPendidikanId, semesterId) => {
        try {
            const payload = {
                schemaname: schemaname,
                semester_id: semesterId,
                tingkat_pendidikan_id: tingkatPendidikanId
            };
            const res = await store.dispatch('siswaService/fetchBanyakSiswaByTingkatId', payload);
            // console.log(res);

            return res;
        } catch (error) {
            throw new Error(`Gagal siswa berdasarkan tingkat: ${error.message}`);
        }
    };
    const fetchBanyakSiswaByRombelId = async (semesterId, rombelId) => {
        // eslint-disable-next-line no-useless-catch
        try {
            const payload = {
                schemaname: schemaname,
                semester_id: semesterId,
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
        fetchBanyakSiswaByRombelId,
        getSiswaAktifByKelasId
    };
}
