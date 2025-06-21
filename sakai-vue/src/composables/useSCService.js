import { useToast } from 'primevue/usetoast';
import { useStore } from 'vuex';

export function useSCService() {
    const store = useStore();
    const toast = useToast();
    const createMetamaskConnected = (payload) => {
        store.commit('scService/SET_METAMASKCONNECTED', payload);
    };
    const getMetamaskConnected = () => {
        try {
            const response = store.getters['scService/getMetamaskConnected'];
            if (!response) {
                return false;
            }
            return response;
        } catch (error) {
            console.log(error);
        }
    };
    const createSCIjazah = async (payload) => {
        try {
            const response = await store.dispatch('scService/createIjazahBC', { degree_data: payload });
            console.log(response);
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menyimpan data: ${error}`, life: 3000 });
        }
    };

    return {
        createMetamaskConnected,
        getMetamaskConnected,
        createSCIjazah
    };
}
