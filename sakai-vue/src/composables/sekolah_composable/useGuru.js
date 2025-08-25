import { ref } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useGuru() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();
    const guruList = ref([]);
    const guruTerdaftarList = ref([]);

    const fetchGuruTerdaftar = async () => {
        try {
            const payload = {
                tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            let res = await store.getters['sekolahService/getPTKTerdaftar'];
            // console.log('useSekolahService/fetchGuruTerdaftar', res);
            if (!res || res.length == 0) {
                // console.log(payload);
                res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
            } else {
                if (res.tahun_ajaran_id != initSelectedSemester.value?.tahunAjaranId) {
                    res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
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
            const response = await store.dispatch('sekolahService/searchPTKTerdaftar', payload);
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

            // console.log('useSekolahService', payload);
            const response = await store.dispatch('sekolahService/deletePTKTerdaftar', payload);
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

            // console.log('useSekolahService', payload);
            const response = await store.dispatch('sekolahService/deleteBatchPTKTerdaftar', payload);
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
            const response = await store.dispatch('sekolahService/updatePTKTerdaftar', payload);

            await store.dispatch('sekolahService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
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
            const response = await store.dispatch('sekolahService/addPTKTerdaftar', payload);
            await store.dispatch('sekolahService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
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
            const response = await store.dispatch('sekolahService/fetchGuru', payload);
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
