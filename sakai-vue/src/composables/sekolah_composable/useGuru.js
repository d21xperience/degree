import { ref } from 'vue';
import { useStore } from 'vuex';

/**
 *
 * @param {String} schemaname Wajib
 * @param {Object} initSelectedSemester Wajib
 * @returns
 */
export function useGuru(schemaname, initSelectedSemester) {
    const store = useStore();
    // const schemaname = computed(() => {
    //     return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    // });
    // const initSelectedSemester = computed(() => store.getters['semesterService/getSelectedSemester']);
    const guruList = ref([]);
    const guruTerdaftarList = ref([]);

    // ACTION
    const fetchGuruTerdaftar = async () => {
        try {
            const payload = {
                tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            let res = store.getters['guruService/getPTKTerdaftar'];
            // console.log('useguruService/fetchGuruTerdaftar', res);
            if (!res || res.length == 0) {
                // console.log(payload);
                res = await store.dispatch('guruService/fetchPTKTerdaftar', payload);
            } else {
                if (res.tahun_ajaran_id != initSelectedSemester.value?.tahunAjaranId) {
                    res = await store.dispatch('guruService/fetchPTKTerdaftar', payload);
                }
            }
            guruTerdaftarList.value = res.ptkTerdaftar;
            return res.ptkTerdaftar;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan guru terdaftar:', error);
        }
    };
    const searchGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };
            const response = await store.dispatch('guruService/searchPTKTerdaftar', payload);
            return response.ptkTerdaftar;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian guru:', error);
        }
    };
    const deleteGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };

            // console.log('useguruService', payload);
            const response = await store.dispatch('guruService/deletePTKTerdaftar', payload);
            return response;
        } catch (error) {
            console.error('Gagal menghapus data guru:', error);
        }
    };
    const deleteBatchGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };

            // console.log('useguruService', payload);
            const response = await store.dispatch('guruService/deleteBatchPTKTerdaftar', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus banyak guru:', error);
        }
    };
    const updateGuruTerdaftar = async (ptkTerdaftar) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar: [ptkTerdaftar._rawValue]
            };
            const response = await store.dispatch('guruService/updatePTKTerdaftar', payload);

            await store.dispatch('guruService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal update guru terdaftar:', error);
        }
    };
    const guruTerdaftarLoading = ref(false);
    const addGuruTerdaftar = async (ptkTerdaftar) => {
        guruTerdaftarLoading.value = true;
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar: ptkTerdaftar._rawValue
            };
            // console.log(payload);
            const response = await store.dispatch('guruService/addPTKTerdaftar', payload);
            await store.dispatch('guruService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
            // toast.add({ severity: 'success', summary: 'Success', detail: 'Berhasil menambah data guru', life: 3000 });

            return response;
        } catch (error) {
            console.error('Gagal update data guru:', error);
        } finally {
            guruTerdaftarLoading.value = false;
        }
    };

    const fetchGuru = async (ptkId = null) => {
        try {
            const payload = {
                schemaname: schemaname.value
            };

            if (ptkId) {
                payload.ptk_id = ptkId;
            }
            const response = await store.dispatch('guruService/fetchGuru', payload);
            guruList.value = response;
        } catch (error) {
            console.error('Gagal mengambil data guru:', error);
        }
    };

    return {
        fetchGuru,
        fetchGuruTerdaftar,
        searchGuruTerdaftar,
        deleteGuruTerdaftar,
        updateGuruTerdaftar,
        guruList,
        guruTerdaftarList,
        deleteBatchGuruTerdaftar,
        addGuruTerdaftar,
        guruTerdaftarLoading
    };
}
