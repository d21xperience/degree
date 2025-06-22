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

    // BC IJAZAH
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

    const getSCIjazah = async (payload) => {
        try {
            let response = await store.getters['scService/getSCIjazah'];
            // console.log(response);
            // return
            if (!response || !Array.isArray(response.degreeData) || response.degreeData.length == 0 || response.tahun_ajaran_id != payload.tahun_ajaran_id) {
                response = await store.dispatch('scService/fetchSCIjazah', payload);
                // console.log(response);
                if (response) {
                    toast.add({ severity: 'success', summary: 'Successful', detail: `"${response.message}"`, life: 3000 });
                    return response.degreeData;
                } else {
                    toast.add({ severity: 'info', summary: 'Failled', detail: `Silahkan reload aplikasi: "${response.message}"`, life: 3000 });
                    return [];
                }
            }
            return response.degreeData;
        } catch (error) {
            console.log(error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan  BC Ijazah: ${error}`, life: 3000 });
        }
    };
    const getContract = () => {
        try {
            let response = store.getters['scService/getContract'];
            if (!response) {
                response = store.dispatch('scService/fetchContract');
                if (response.status) {
                    return response.contract;
                }
            }
            return response;
        } catch (error) {
            console.log(error);
        }
    };

    return {
        createMetamaskConnected,
        getMetamaskConnected,
        createSCIjazah,
        getContract,
        getSCIjazah
    };
}
