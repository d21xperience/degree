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
    // ================================================
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
    const getWalletInfo = async (payload) => {
        try {
            let response = store.getters['scService/getWalletInfo'];
            if (!response) {
                response = await store.dispatch('scService/fetchWalletInfo', payload);
                if (response.status) {
                    toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                    return response;
                } else {
                    toast.add({ severity: 'error', summary: 'Gagal', detail: `${response.message}`, life: 3000 });
                    return null;
                }
            }
            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Error', detail: `${error}`, life: 3000 });
        }
    };
    // ================================================
    const fetchBCNetworks = async (networkArchitecture = '') => {
        try {
            let response = store.getters['scService/getBCNETWORK'](networkArchitecture);
            // console.log(response)
            if (!response) {
                response = await store.dispatch('scService/fetchBlockchainNetworks', networkArchitecture);
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.network;
            }
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Berhasil mendapatkan network', life: 3000 });
            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan  BC Network: ${error}`, life: 3000 });
        }
    };

    const updateBCNetwork = async (bcNetwork) => {
        try {
            const { status, message } = await store.dispatch('scService/updateBlockchainNetwork', {
                network: bcNetwork
            });

            if (!status) {
                throw new Error(`Update gagal: ${message}`);
            }

            // Ambil data dari store
            const currentNetworks = store.getters['scService/getBCNETWORK']();
            if (!Array.isArray(currentNetworks)) {
                throw new Error('State BCNETWORK tidak valid (bukan array)');
            }

            // Perbarui elemen yang sesuai dengan Id
            const updatedNetworks = currentNetworks.map((item) => (item.Id === bcNetwork.Id ? { ...item, ...bcNetwork } : item));

            // Commit perubahan ke store
            store.commit('scService/SET_BCNETWORK', updatedNetworks);

            toast.add({
                severity: 'success',
                summary: 'Berhasil',
                detail: `Berhasil update BC Network: ${message}`,
                life: 3000
            });
        } catch (error) {
            console.error('updateBCNetwork error:', error);
            toast.add({
                severity: 'error',
                summary: 'Gagal',
                detail: `Gagal update BC Network: ${error.message || error}`,
                life: 3000
            });
        }
    };
    const deleteBCNetwork = async (bcNetwork = []) => {
        try {
            // kirim ke server
            const { status, message } = await store.dispatch('scService/deleteBlockchainNetwork', bcNetwork);
            // console.log(status);
            if (status) {
                // Validasi awal
                if (!Array.isArray(bcNetwork) || bcNetwork.length === 0) {
                    throw new Error('Data bcNetwork tidak valid atau kosong');
                }

                // Ambil data sekarang dari store
                const currentData = store.getters['scService/getBCNETWORK']();

                if (!Array.isArray(currentData)) {
                    throw new TypeError("Getter 'getBCNETWORK' tidak mengembalikan array");
                }

                // Buat array berisi ID yang ingin dihapus
                const idsToDelete = bcNetwork.map((item) => item.Id);

                // Filter data: hapus item yang ID-nya ada di idsToDelete
                const updatedData = currentData.filter((item) => !idsToDelete.includes(item.Id));

                // Commit perubahan ke store
                store.commit('scService/SET_BCNETWORK', updatedData);
            }

            // Tampilkan notifikasi berhasil
            toast.add({ severity: 'success', summary: 'Berhasil', detail: `Berhasil menghapus ${message} item`, life: 3000 });
        } catch (error) {
            console.error(error);
            toast.add({ severity: 'error', summary: 'Gagal', detail: `Gagal menghapus BC Network: ${error.message}`, life: 3000 });
        }
    };

    const searchBCNetwork = async (param) => {
        try {
            let response = store.getters['scService/getBCNETWORK'](param);
            // console.log(response)
            if (!response) {
                response = await store.dispatch('scService/searchBlockchainNetworks', param);
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.network;
            }
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Berhasil mendapatkan network', life: 3000 });
            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan  BC Network: ${error}`, life: 3000 });
        }
    };
    const setBCNetwork = async (networkPlatform) => {
        await store.dispatch('scService/setBlockchainNetwork', networkPlatform);
    };
    const getBCNetwork = async () => {
        return await store.getters['scService/getBCNetworkSelected'];
    };

    // ================================================
    // Sol Compiler
    // ================================================
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

    // ========================================
    // EVM OR NON EVM CHECK
    // ========================================
    const getNetworkChainId = async (rpcUrl) => {
        try {
            const response = await store.dispatch('scService/getNetworkChainId', rpcUrl);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Berhasil', detail: `Berhasil terhubung: ${response.message}`, life: 3000 });
                return response.networkDetail;
            }
        } catch (error) {
            console.log(error);
            toast.add({ severity: 'error', summary: 'Gagal', detail: `Gagal koneksi ke Network: ${error.message}`, life: 3000 });
        }
    };

    // ========================================

    // ========================================
    // Platform
    // ========================================
    const fetchNetworkPlatform = async () => {
        try {
            const response = await store.dispatch('scService/fetchBCPlatform');
            if (response.status) {
                return response.bcPlatform;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Gagal', detail: `Gagal mendapatkan Platform: ${error.message}`, life: 3000 });
        }
    };
    const setNetwrokPlatform = async (networkPlatform) => {
        await store.dispatch('scService/updateBCPlatformSelected', networkPlatform);
    };
    const getNetowrkPlatform = () => {
        return store.getters['scService/getBCPlatformSelected'];
    };
    const cekNetworkPlatform = (networConfig) => {};
    // ========================================

    // ========================================
    // BC ACCOUNT - Acount Service
    // ========================================
    const fetchBCAccount = async () => {
        try {
            const currentUser = store.getters['authService/currentUser'];
            const response = await store.dispatch('scService/fetchBCAccount', currentUser.username);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Berhasil', detail: `Berhasil mengambil data ${response.message} item`, life: 3000 });
                return response.accounts;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Gagal', detail: `Gagal mendapatkan Platform: ${error.message}`, life: 3000 });
        }
    };
    // ========================================
    // ========================================
    // Blockchain service

    const setBCConfig = async () => {
        try {
            const payload = {
                environment: await getBCNetwork(),
                platform: await getNetowrkPlatform()
            };
            const response = await store.dispatch('scService/setBCConfig', { bc_config: payload });
            console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response;
            }
        } catch (error) {
            console.log(error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal terhubung ke jaringan: ${error}`, life: 3000 });
        }
    };

    const setBCConnected = async (isConnected = false) => {
        store.commit('scService/SET_BcConected', isConnected);
    };
    const getBCConnected = async () => {
        return store.getters['scService/getBcConnected'];
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
        fetchBCNetworks,
        updateBCNetwork,
        deleteBCNetwork,
        getNetworkChainId,
        setNetwrokPlatform,
        fetchNetworkPlatform,
        getNetowrkPlatform,
        fetchBCAccount,
        searchBCNetwork,
        setBCNetwork,
        getBCNetwork,
        setBCConfig,
        setBCConnected,
        getBCConnected

        // batchDeleteBCNetwork
    };
}
