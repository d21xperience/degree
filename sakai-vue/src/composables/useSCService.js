import { useToast } from 'primevue/usetoast';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { useSekolahService } from './useSekolahService';
export function useSCService() {
    const store = useStore();
    const toast = useToast();
    const sekolahService = useSekolahService();
    const createMetamaskConnected = (payload) => {
        store.commit('scService/SET_METAMASKCONNECTED', payload);
    };
    const schemaname = computed(() => sekolahService.schemaname.value);
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

    const getBCTransaction = async () => {
        const payload = {
            schemaname: schemaname.value
        };

        // console.log(payload);
        // return;
        let response = await store.dispatch('scService/fetchBCTransaction', payload);
        if (response.status) {
            return response.bcTransaction;
        }
    };

    const createWalletInfo = async (payload) => {
        try {
            if (payload) {
                console.log(payload);
                await store.dispatch('scService/createWalletInfo', payload);
            }
            // return;
        } catch (error) {}
    };
    const getWalletInfo = async () => {
        try {
            const response = store.getters['scService/getWalletInfo'];
            if (response) {
                return response;
            }
        } catch (error) {}
    };

    const fetchBCNetworks = async () => {
        try {
            let response = store.getters['scService/getBCNETWORK'];
            if (!response || !Array.isArray(response) || response.length === 0) {
                response = await store.dispatch('scService/fetchBlockchainNetworks');
                return response.network;
                // console.log(network);
                // // BCNetworks.value = data.filter(item => !item.activate)
                // BCNetworks.value = network;

                // let obj ={...data}
                // console.log(obj)
                // BCNetwork.value = obj
            }
            return response;
        } catch (error) {
            if (error.response) {
                console.error(`Error ${error.response.status}: ${error.response.data}`);
            } else {
                console.error('Error:', error.message);
            }
        }
    };

    const getSolVersion = async () => {
        const { data } = await store.dispatch('scService/getsolVersion');
        if (data.status) {
            return data.message;
        }
        return '';
    };
    const deployContract = async (formData) => {
        const response = await store.dispatch('scService/deployContract', formData);
        console.log(response);
        // if (response.status){
        //     return `Kontrak berhasil dideploy! Address:` //${data.contractAddress}` //data.message
        //     //     status.value = `Kontrak berhasil dideploy! Address: ${result.contractAddress}`;
        // } else {
        //     return  `Gagal: ${result.error || 'Terjadi kesalahan'}`;
        // }
    };
    return {
        createMetamaskConnected,
        getMetamaskConnected,
        createSCIjazah,
        getContract,
        getSCIjazah,
        getBCTransaction,
        createWalletInfo,
        getWalletInfo,
        getSolVersion,
        deployContract,
        fetchBCNetworks
    };
}
