import { computed } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './sekolah_composable/useSemester';
import { useTableTenant } from './sekolah_composable/useTableTenant';

export function useAuthService() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();
    const currentUser = computed(() => store.getters['authService/currentUser']);
    const fetchUser = async () => {
        try {
            const payload = {
                tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            let res = await store.getters['sekolahService/getPTKTerdaftar'];
            if (!res) {
                res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
            } else {
                if (res.tahun_ajaran_id != initSelectedSemester.value?.tahunAjaranId) {
                    res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
                }
            }
            // guruTerdaftarList.value = res.ptkTerdaftar;
            return res.ptkTerdaftar;
        } catch (error) {
            console.error('Gagal mengambil data guru:', error);
        }
    };

    return {
        currentUser,
        fetchUser
    };
}
