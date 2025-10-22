// import { useToast } from 'primevue/usetoast';
import { useStore } from 'vuex';
import { useTableTenant } from './sekolah_composable/useTableTenant';
export function useSCService() {
    const store = useStore();
    const { schemaname } = useTableTenant();
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
    // ================================================
    // BC IJAZAH
    const createSCIjazah = async (payload) => {
        try {
            const response = await store.dispatch('scService/createIjazahBC', { degree_data: payload });
            return response;
        } catch (error) {
            console.error('Create Ijazah failed:', error);
            throw new Error(`Gagal membuat ijazah: ${error.message}`);
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
                    return response.degreeData;
                } else {
                    return [];
                }
            }
            return response.degreeData;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mendapatkan ijazah: ${error.message}`);
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

    // =============================================

    // ================================================
    const fetchBCNetworks = async (payload = '') => {
        try {
            let response = store.getters['scService/getBCNETWORK'](payload);
            // console.log(response);
            if (!response || !Array.isArray(response) || response.length == 0) {
                response = await store.dispatch('scService/fetchBlockchainNetworks', payload);
                return response.network;
            }
            return response;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mengambil BC Network: ${error.message}`);
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
        } catch (error) {
            console.error('updateBCNetwork error:', error);
        }
    };
    const deleteBCNetwork = async (bcNetwork = []) => {
        try {
            // kirim ke server
            const { status } = await store.dispatch('scService/deleteBlockchainNetwork', bcNetwork);
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
        } catch (error) {
            console.error(error);
        }
    };

    const searchBCNetwork = async (param) => {
        try {
            let response = store.getters['scService/getBCNETWORK'](param);
            // console.log(response)
            if (!response) {
                response = await store.dispatch('scService/searchBlockchainNetworks', param);
            }
            return response;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal melakukan pencarian BC Netowrk: ${error.message}`);
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
                return response.networkDetail;
            }
        } catch (error) {
            console.log(error);
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
            throw new Error(`Gagal mendapatkan platform network: ${error.message}`);
        }
    };
    const setNetwrokPlatform = async (networkPlatform) => {
        await store.dispatch('scService/updateBCPlatformSelected', networkPlatform);
    };
    const getNetowrkPlatform = () => {
        return store.getters['scService/getBCPlatformSelected'];
    };

    // ========================================
    // BC ACCOUNT - Acount Service
    // ========================================
    const fetchBCAccount = async () => {
        try {
            const currentUser = store.getters['authService/currentUser'];
            const response = await store.dispatch('scService/fetchBCAccount', currentUser.username);
            if (response.status) {
                return response.accounts;
            }
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mendapatkan BC Akun: ${error.message}`);
        }
    };

    const importBCAccount = async (payload) => {
        try {
            const response = await store.dispatch('scService/importBCAccount', payload);
            if (response.status) {
                return response.accounts;
            }
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal mengimpor BC Akun: ${error.message}`);
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
            // console.log(response);
            if (response.status) {
                return response;
            }
        } catch (error) {
            console.log(error);
            throw error;
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
        getBCConnected,
        importBCAccount
    };
}
