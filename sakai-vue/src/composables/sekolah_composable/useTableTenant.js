import { computed } from 'vue';
import { useStore } from 'vuex';
export function useTableTenant() {
    const store = useStore();
    const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);

    const fetchTabelTenant = async () => {
        try {
            let response = store.getters['sekolahService/getTabeltenant'];
            if (!response) {
                await store.dispatch('sekolahService/fetchTabeltenant');
                response = store.getters['sekolahService/getTabeltenant'];
            }
            return response;
        } catch (error) {
            console.log(error);
        }
    };
    return {
        schemaname,
        fetchTabelTenant
    };
}
