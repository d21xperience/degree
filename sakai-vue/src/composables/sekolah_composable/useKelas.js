import { computed, ref } from 'vue';
import { useStore } from 'vuex';
// =================================================
// KELAS
// =================================================
export function useKelas() {
    const store = useStore();
    const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']);
    const isFetching = ref(null);
    const isError = ref(false);
    // const selectedSemester = computed(() => store.getters['semesterService/getSelectedSemester']);
    // console.log('call useKelas', schemaname.value);
    const kelasList = ref([]);
    /**
     *
     * @param {Object} selectedSemester Wajib
     * @param {String} kelasId Optional
     * @param {String} tingkatPendidikanId Optional
     * @returns
     */
    const fetchKelas = async (selectedSemester = null, kelasId = '', tingkatPendidikanId = '') => {
        isFetching.value = true;
        try {
            const cached = computed(() => store.getters['kelasService/getKelas']);
            console.log('----[fetchKelas cached]----', typeof cached.value);
            if (cached.value) {
                console.log('----[blok1]----');
                if (cached.value.semesterId === selectedSemester?.semesterId) {
                    console.log('----[blok2]----');
                    return cached.value.kelas;
                }
            }

            console.log('----[blok cek]----', schemaname.value);
            const payload = {
                schemaname: schemaname.value.schemaname,
                semester_id: selectedSemester?.semesterId
            };
            if (kelasId) {
                payload.kelas_id = kelasId;
            }
            if (tingkatPendidikanId) {
                payload.tingkat_pendidikan_id = tingkatPendidikanId;
            }
            console.log('----[blok3]----');
            const response = await store.dispatch('kelasService/fetchKelas', payload);
            return response.kelas;
        } catch (error) {
            console.log(error);
            isError.value = true;
            throw new Error(`Gagal mengambil kelas: ${error.message}`);
        } finally {
            isFetching.value = false;
        }
    };
    // const getKelas = async () => {
    //     try {
    //         let response = store.getters['kelasService/getKelas'];
    //         if (!response || Array.isArray(response.kelas) || response.kelas.length == 0 || selectedSemester?.semesterId != response?.semesterId) {
    //             response = await fetchKelas();
    //         }
    //         kelasList.value = response.kelas;
    //         return response;
    //     } catch (error) {
    //         throw new Error(`Gagal mendapatkan kelas: ${error.message}`);
    //     }
    // };

    /**
     *
     * @param {Object} selectedSemester Wajib
     * @param {String} kelasId Optional
     * @returns
     */
    const getById = async (selectedSemester = null, kelasId = null) => {
        try {
            let response = store.getters['kelasService/getKelas'];

            if (!response || response.kelas.length == 0 || selectedSemester?.semesterId != response?.semesterId) {
                const payload = {
                    schemaname: schemaname.value.schemaname,
                    semester_id: selectedSemester?.semesterId
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
                schemaname: schemaname.value.schemaname,
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

    const update = () => {};

    /**
     *
     * @param {String} rombonganBelajarId
     * @param {String} semesterId
     * @returns
     */
    const fetchAnggotaKelas = async (rombonganBelajarId = '', semesterId = '') => {
        console.log(rombonganBelajarId);
        console.log(semesterId);
        if (rombonganBelajarId == '' || semesterId == '') {
            throw new Error('Parameter wajib disi');
        }
        try {
            const cachedData = store.getters['siswaService/getSiswaAktif'];
            console.log('cachaedData', cachedData);
            console.log(semesterId);
            // if (cachedData.semester_id === semesterId) {
            //     // const anggotaKelas = cachedData.peserta_didik.filter((val) => val.rombonganBelajarId === rombonganBelajarId);
            //     // return anggotaKelas;
            // }
            // return null;
        } catch (error) {
            console.error('Gagal mengambil data kelas:', error);
        }
    };

    return {
        fetchKelas,
        isError,
        // getKelas,
        addKelas,
        getById,
        update,
        kelasList,
        fetchAnggotaKelas
    };
}
