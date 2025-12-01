import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useDashboard() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { selectedSemester } = useSemester();

    const fetchDashboard = async () => {
        try {
            let response = store.getters['sekolahService/getDashboard'];
            if (!response || response.semester_id != selectedSemester.value) {
                // console.log('useDashboard1', response.semester_id != storeSelectedSemester.value?.semesterId);
                console.log(schemaname.value);
                console.log(selectedSemester.value);
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: selectedSemester.value
                };
                if (payload.schemaname === '') {
                    return null;
                }
                response = await store.dispatch('sekolahService/fetchDashboard', payload);
            }

            return response;
        } catch (error) {
            throw new Error(error);
        }
    };

    return {
        fetchDashboard
    };
}
