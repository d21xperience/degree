import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useDashboard() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();

    const fetchDashboard = async () => {
        try {
            let response = await store.getters['sekolahService/getDashboard'];
            if (response) {
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: initSelectedSemester.value?.semesterId
                };
                if (payload.schemaname === '') {
                    return;
                }
                response = await store.dispatch('sekolahService/fetchDashboard', payload);
            }

            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal medapatkan Dashboard:', error);
        }
    };

    return {
        fetchDashboard
    };
}
