import { computed, ref, watch } from 'vue';
import { useStore } from 'vuex';
export function useSemester() {
    const store = useStore();
    const listTahunAjaran = computed(() => rawlistTahunAjaran.value || []);
    const rawlistTahunAjaran = ref();
    const listSemester = store.getters['semesterService/getSemester'];
    const initSelectedSemester = computed(() => store.getters['semesterService/getSelectedSemester']);
    const selectedSemester = ref();
    const initSelectedTahunAjaran = computed(() => store.getters['semesterService/getSelectedTahunAjaran']);
    const selectedTahunAjaran = ref();
    const fetchSemester = async () => {
        try {
            const results = await store.dispatch('semesterService/fetchSemester');
            if (results.status) {
                // const periodeAktif = results.semester
                return results.semester;
            }
        } catch (error) {
            throw new Error(`Gagal mengambil semester: ${error.message}`);
        }
    };

    /**
     * Gets a contract for the given owner address
     * @param {Array} semester
     * @returns {Promise} A promise that resolves with the contract response
     * @throws {Error} If there's an error fetching the contract
     */
    /**
     * Deletes semesters by their IDs
     * @param {Array<Object>} semesters - Array of semester objects containing semesterId
     * @returns {Promise<Object>} A promise that resolves with the deletion response
     * @throws {Error} If the deletion request fails
     *
     * @example
     * await deleteSemester([{ semesterId: 1 }, { semesterId: 2 }]);
     */
    const deleteSemester = async (semesters) => {
        // Validasi input
        if (!Array.isArray(semesters)) {
            throw new Error('Parameter must be an array of semester objects');
        }

        // Ekstrak semesterId secara langsung
        const semesterIds = semesters
            .map((semester) => {
                if (semester.semesterId == null) {
                    console.warn('Semester object missing semesterId:', semester);
                }
                return semester.semesterId;
            })
            .filter((id) => id != null); // Filter null/undefined

        if (semesterIds.length === 0) {
            throw new Error('No valid semester IDs provided');
        }
        console.log(semesterIds);
        // return;
        try {
            // Tunggu hasil dispatch dengan await
            const response = await store.dispatch('semesterService/deleteSemester', semesterIds);
            console.log('deleteSemester', response);
            // Asumsi response memiliki struktur { status: true, data: ... } atau sejenisnya
            if (response.status) {
                return response;
            } else {
                throw new Error(response?.message || 'Failed to delete semesters');
            }
        } catch (error) {
            // Tambahkan konteks error
            console.error('Error deleting semesters:', error);
            throw new Error(`Failed to delete semesters: ${error.message || 'Unknown error'}`);
        }
    };

    const updateSemester = async (semester) => {
        try {
            console.log('updateSemester', semester);
            // return
            const res = await store.dispatch('semesterService/updateSemester', semester);
            return res;
        } catch (error) {
            throw new Error(`Gagal update semester: ${error.message}`);
        }
    };

    const fetchTahunAjaran = async () => {
        try {
            rawlistTahunAjaran.value = store.getters['semesterService/getTahunAjaran'];
            if (!rawlistTahunAjaran.value || rawlistTahunAjaran.value.length == 0) {
                const results = await store.dispatch('semesterService/fetchTahunAjaran');
                if (results.status) {
                    rawlistTahunAjaran.value = results.tahunAjaran;
                    // toast.add({ severity: 'success', summary: 'Successful', detail: `${results.message}`, life: 3000 });
                }
            }
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mendapatkan tahun ajaran: ${error.message}`);
        }
    };

    watch(selectedTahunAjaran, (e) => {
        store.commit('semesterService/SET_SELECTEDTAHUNAJARAN', e);
    });
    watch(selectedSemester, (e) => {
        store.commit('semesterService/SET_SELECTEDSEMESTER', e);
    });
    return {
        fetchSemester,
        deleteSemester,
        updateSemester,
        selectedSemester,
        initSelectedSemester,
        listTahunAjaran,
        listSemester,
        fetchTahunAjaran,
        selectedTahunAjaran,
        initSelectedTahunAjaran
    };
}
