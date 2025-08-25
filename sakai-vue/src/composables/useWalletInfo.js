import store from '@/store';
import { computed, ref } from 'vue';

export function useWalletInfo() {
    // const toast = useToast();

    // State lokal wallet
    const currentWallet = ref({
        address: '',
        chain: {
            chainId: '',
            name: '',
            rpc: '',
            explorer: '',
            nativeCurrency: null
        },
        balance: {
            wei: '',
            formatted: ''
        },
        tokens: [],
        gas: {
            gasPrice: '',
            maxFeePerGas: '',
            maxPriorityFeePerGas: ''
        },
        meta: null
    });

    const isConnected = ref(true);

    // Computed: cek apakah wallet sudah dimuat
    const hasWalletData = computed(() => {
        return !!currentWallet.value.address;
    });

    // Simpan ke Vuex store
    const updateWalletInStore = (walletData) => {
        store.commit('scService/SET_WALLETINFO', walletData);
    };

    // Hanya dispatch untuk membuat wallet (jika ini untuk inisialisasi)
    const createWallet = async (payload) => {
        try {
            if (payload) {
                console.log('[useWalletInfo] Creating wallet...', payload);
                await store.dispatch('scService/createWalletInfo', payload);
                // toast.add({
                //     severity: 'success',
                //     summary: 'Wallet Created',
                //     detail: 'New wallet initialized successfully.',
                //     life: 3000
                // });
            }
        } catch (error) {
            console.error('[useWalletInfo] Failed to create wallet:', error);
            // toast.add({
            //     severity: 'error',
            //     summary: 'Creation Failed',
            //     detail: 'Could not create wallet.',
            //     life: 3000
            // });
        }
    };

    // Ambil dari store atau fetch dari API
    const loadWalletInfo = async (payload) => {
        try {
            // Cek dulu di store
            let walletData = store.getters['scService/getWalletInfo'];
            if (walletData) {
                currentWallet.value = walletData;
                return walletData;
            }

            // Jika tidak ada, fetch
            const response = await fetchWalletInfo(payload);
            if (response?.status && response.walletData) {
                currentWallet.value = response.walletData;
                return response.walletData;
            }

            return null;
        } catch (error) {
            console.error('[useWalletInfo] Error loading wallet:', error);
            return null;
        }
    };

    // Panggil API untuk fetch wallet info
    /**
     * @param {{ public_address: string }} pubAddress
     */
    const fetchWalletInfo = async (pubAddress) => {
        try {
            // console.log(pubAddress);
            const response = await store.dispatch('scService/fetchWalletInfo', pubAddress);
            if (response.status) {
                return response;
            } else {
                return null;
            }
        } catch (error) {
            console.error('[useWalletInfo] API fetch error:', error);
            throw error;
        }
    };
    /**
     * @param {{ public_address: string }} pubAddress
     */
    const getWalletDetail = async (pubAddress) => {
        const { data } = await store.dispatch('scService/fetchWalletInfo', pubAddress);
        console.log(data);
        return data;
    };
    return {
        // State
        currentWallet,
        isConnected,
        hasWalletData,

        // Actions
        updateWalletInStore,
        createWallet,
        loadWalletInfo,
        fetchWalletInfo,
        getWalletDetail
    };
}
