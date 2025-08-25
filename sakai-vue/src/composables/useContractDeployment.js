import { ref } from 'vue';
import { useStore } from 'vuex';
export function useContractDeployment() {
    const file = ref(null);
    const status = ref('');
    const isProcessing = ref(false);
    const abiName = ref('');
    const binName = ref('');
    const compileStatus = ref(false);
    const deployStatus = ref(false);
    const store = useStore();
    const handleFileUpload = (event) => {
        file.value = event.target.files[0];
    };
    /**
     * @param {}
     */
    const compileContract = async () => {
        if (!file.value) {
            status.value = 'Silakan pilih file Solidity terlebih dahulu.';
            return;
        }

        const formData = new FormData();
        formData.append('file', file.value);

        isProcessing.value = true;
        status.value = 'Mengunggah dan mendeploy kontrak...';

        try {
            const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/scs/contract/compile-contract`, {
                method: 'POST',
                body: formData
            });

            const result = await response.json();
            if (result.status) {
                compileStatus.value = true;
                status.value = result.message;
                abiName.value = result.abi_filename;
                binName.value = result.bytecode_filename;
            } else {
                status.value = `Gagal: ${result.error || result.message || 'Terjadi kesalahan'}`;
            }
        } catch (err) {
            status.value = `Error jaringan: ${err.message}`;
        } finally {
            isProcessing.value = false;
            file.value = null;
        }
    };

    const cancelBuildContract = () => {
        compileStatus.value = false;
        status.value = '';
        file.value = null;
    };
    const deployContract = async (payload) => {
        try {
            const response = await store.dispatch('scService/deployContract', payload);
            console.log('useContractDeployments', response);
            if (response.status) {
                return response;
            } else {
                return response;
            }
        } catch (error) {
            return error;
        }
    };
    /**
     * Gets a contract for the given owner address
     * @param {string} ownerAddress - The address of the owner
     * @returns {Promise} A promise that resolves with the contract response
     * @throws {Error} If there's an error fetching the contract
     */
    const getContract = async (ownerAddress) => {
        try {
            const response = await store.dispatch('scService/getContract', ownerAddress);
            // console.log(response)
            return response;
        } catch (error) {
            throw error;
        }
    };

    const activeContract = async () => {
        try {
            const response = await store.dispatch('scService/activeContract');
            // console.log(response)
            return response;
        } catch (error) {
            throw error;
        }
    };
    return {
        file,
        status,
        isProcessing,
        abiName,
        binName,
        compileStatus,
        handleFileUpload,
        compileContract,
        cancelBuildContract,
        deployContract,
        deployStatus,
        getContract,
        activeContract
    };
}
