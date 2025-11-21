<script setup>
import ContractDeployer from '@/components/scComponent/ContractDeployer.vue';
import { useContractDeployment } from '@/composables/useContractDeployment';
import { useWalletInfo } from '@/composables/useWalletInfo';
import { onMounted, reactive, ref } from 'vue';
const { currentWallet, loadWalletInfo } = useWalletInfo();
const state = reactive({
    password: '',
    confirmPassword: '',
    isGenerating: false,
    showsmartContract: false,
    smartContractData: null,
    error: '',
    smartContracts: [],
    loadingsmartContracts: false
});
const { getContract } = useContractDeployment();
const isAddDialog = ref(false);
const handleAddDialog = async () => {
    // const res = await loadWalletInfo(currentWallet.value.address);
    // Object.assign(currentWallet.value, res);
    // console.log(currentWallet.value);
    if (currentWallet.value?.address == '') {
        alert('Wallet belum dipilih, silahkan pilih terlebih dahulu!');
    } else {
        isAddDialog.value = true;
    }
};
const loadsmartContracts = async () => {
    state.loadingsmartContracts = true;
    try {
        const response = await getContract(currentWallet.value.address);
        console.log(response);
        if (response.status) {
            state.loadingsmartContracts = false;
            state.smartContracts = response.contract;
        }
    } catch (error) {
    } finally {
        state.loadingsmartContracts = false;
    }
};
onMounted(async () => {
    const res = await loadWalletInfo(currentWallet.value.address);
    Object.assign(currentWallet.value, res);
    await loadsmartContracts();
});
</script>
<template>
    <div class="rounded-xl">
        <div class="flex justify-between items-center mb-6">
            <div class="flex space-x-2">
                <h2 class="text-2xl font-semibold">Generated Smart Contract</h2>
                <div>
                    <Button icon="pi pi-plus" size="small" severity="secondary" @click="handleAddDialog" />
                </div>
            </div>
            <button :disabled="state.loadingsmartContracts" class="px-4 py-2 hover:bg-gray-200 rounded-lg text-sm font-medium" @click="loadsmartContracts">
                {{ state.loadingsmartContracts ? 'Loading...' : 'Refresh' }}
            </button>
        </div>

        <div v-if="state.loadingsmartContracts" class="text-center py-8">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>

        <div v-else-if="state.smartContracts.length === 0" class="text-center py-8 text-gray-500">No smart contracts generated yet</div>

        <div v-else class="space-y-4">
            <div v-for="smartContract in state.smartContracts" :key="smartContract.id" class="border border-gray-200 rounded-lg p-4 hover:bg-gray-50 transition">
                <div class="flex justify-between items-start">
                    <div class="flex-1">
                        <div class="font-mono text-sm text-gray-600 mb-1">Contract Name: {{ smartContract?.contractName }}</div>
                        <div class="font-mono text-sm text-gray-600 mb-1">Contract Address: {{ smartContract?.contractAddress }}</div>
                        <div class="font-mono text-sm text-gray-600 mb-1">Contract Owner: {{ smartContract?.contractOwner }}</div>
                        <div class="font-mono text-sm text-gray-600 mb-1">Owner Address: {{ smartContract?.ownerAddress }}</div>
                        <!-- <div class="text-xs text-gray-500">Created: {{ formatDate(smartContract.createdAt) }} | Filename: {{ smartContract.filename }}</div> -->
                    </div>
                    <button class="ml-4 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium py-2 px-3 rounded-lg transition" @click="downloadKeystore(smartContract.id, smartContract.filename)">Activate</button>
                </div>
            </div>
        </div>

        <Dialog v-model:visible="isAddDialog" position="top" header="Deploy Smart Contract" :modal="true">
            <div class="p-2">
                <ContractDeployer />
            </div>
            <template #footer></template>
        </Dialog>
    </div>
</template>
