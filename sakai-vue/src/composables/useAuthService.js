import { computed, ref } from 'vue';
import { useStore } from 'vuex';

export function useAuthService() {
    const store = useStore();
    const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
    const fetchTabeltenant = async () => {};
    const currentUser = computed(()=> store.getters['authService/currentUser'])
    const fetchUser = async () => {
        try {
            const payload = {
                tahunAjaranId: selectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            let res = await store.getters['sekolahService/getPTKTerdaftar'];
            if (!res) {
                res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
            } else {
                if (res.tahun_ajaran_id != selectedSemester.value?.tahunAjaranId) {
                    res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
                }
            }
            guruTerdaftarList.value = res.ptkTerdaftar;
            return res.ptkTerdaftar;
        } catch (error) {
            console.error('Gagal mengambil data guru:', error);
        }
    };

    return {
        currentUser
    };
}
