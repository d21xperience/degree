import { useStore } from 'vuex';

export function useSCService() {
    const store = useStore();

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

    return {
        createMetamaskConnected,
        getMetamaskConnected
    };
}
