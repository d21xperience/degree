import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useDashboard() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();

    const fetchDashboard = async () => {
        try {
            let response = store.getters['sekolahService/getDashboard'];
            // console.log('useDashboard', response);
            if (!response || response.semester_id != initSelectedSemester.value?.semesterId) {
                // console.log('useDashboard1', response.semester_id != initSelectedSemester.value?.semesterId);
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: initSelectedSemester.value?.semesterId
                };
                if (payload.schemaname === '') {
                    return null;
                }
                response = await store.dispatch('sekolahService/fetchDashboard', payload);
            }

            return response;
        } catch (error) {
            throw new Error('Gagal medapatkan nilai untuk Dashboard');
        }
    };

    return {
        fetchDashboard
    };
}
