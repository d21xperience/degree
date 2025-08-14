<script setup>
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { InputText, useToast } from 'primevue';
import { computed, onMounted, reactive, ref, watch } from 'vue';

import { useContractDeployment } from '@/composables/useContractDeployment';
import { useWalletInfo } from '@/composables/useWalletInfo';

const toast = useToast();
const scService = useSCService();
const { formatBalance, shortenAddress } = useUtils();

// Contract Deployment Logic
const { file, status, isProcessing, abiName, binName, compileStatus, handleFileUpload, compileContract, cancelBuildContract, deployContract, deployStatus } = useContractDeployment();

// Wallet Info Logic
const { isConnected, isWalletInfoAvailable, getWalletDetail, currentWallet, loadWalletInfo } = useWalletInfo();

// Password Dialog Logic
const showPasswordDialog = ref(false);
const password = ref('');

const handlePasswordSubmit = (submittedPassword) => {
    password.value = submittedPassword;
    // Here you would typically use the password for further actions, e.g., actual deployment
    console.log('Password submitted:', password.value);
    // For now, we'll just close the dialog
    showPasswordDialog.value = false;
};

const contractRequest = reactive({
    name: 'Default Contract',
    contract_address: '',
    contract_owner: 'Default Owner',
    owner_address: '',
    tx_hash: '',
    network_id: ''
});
const handleForward = () => {
    showPasswordDialog.value = true;
};
const handleDeploy = async () => {
    // This function would contain the actual contract deployment logic after password submission
    console.log('Initiating deploy...');
    deployStatus.value = true;
    try {
        console.log(currentWallet.value);
        const payload = {
            password: '12345', //password.value,
            contract_request: {
                name: contractRequest.name,
                contract_address: currentWallet.value.address,
                contract_owner: contractRequest.contract_owner,
                owner_address: contractRequest.owner_address,
                // network_id: contractRequest.
            },
            username: 'superadmin', //currentWallet.value.Username,
            abi_name: abiName.value,
            bin_name: binName.value
        };
        const response = await deployContract(payload);
        console.log(response);
        if (response.status == true) {
            alert(response?.message);
        } else {
            alert(response?.message);
        }
    } catch (error) {
        console.log(error);
    } finally {
        deployStatus.value = false;
    }
};

const handleReset = () => {
    cancelBuildContract();
    password.value = '';
};
// Other component specific states
const solcVersion = ref('');
const unitOptions = ref([
    { id: 1, nama: 'WEI' },
    { id: 2, nama: 'GWEI' },
    { id: 3, nama: 'ETH' }
]);
const unitSelected = ref();

const defaultValue = computed(() => {
    const defVal = 3000000;
    return defVal.toLocaleString('id-ID');
});

const handleCompile = async () => {
    // cek apakah jaringan sudah berjalan
    await compileContract();
    // await getWallet();
};
watch(compileStatus, async (newVal) => {
    if (newVal) {
        try {
            console.log(currentWallet.value);
            // const response = await getWalletDetail({ public_address: currentWallet.value.address });
            // console.log(response);
            // if (currentWallet.value) {
            //     Object.assign(currentWallet.value, response);
            //     toast.add({
            //         severity: 'success',
            //         summary: 'Success',
            //         detail: response.message,
            //         life: 3000
            //     });
            // }
        } catch (error) {
            console.error('[useWalletInfo] API fetch error:', error);
            toast.add({
                severity: 'error',
                summary: 'Network Error',
                detail: 'Unable to reach server.',
                life: 3000
            });
        }
    }
});
onMounted(async () => {
    solcVersion.value = await scService.getSolVersion();
    unitSelected.value = unitOptions.value[0].nama;
    await loadWalletInfo();
});
</script>

<template>
    <div class="max-w[600px] rounded-lg">
        <!-- <h2>Deploy Smart Contract</h2> -->
        <div v-if="!compileStatus">
            <p class="">Step 1: Compile Contract</p>
            <div>Compiler: {{ solcVersion ? solcVersion : '' }}</div>
            <input type="file" accept=".sol" @change="handleFileUpload" />
            <Button icon="pi pi-refresh" @click="handleCompile" :disabled="isProcessing || !file" :label="isProcessing ? 'Memproses...' : 'Compile Contract'" />
            <p v-if="status">{{ status }}</p>
        </div>
        <div v-else>
            <p class="text-lg">Step 2: Build Contract</p>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div class="grid grid-cols-1">
                    <div class="border p-2 rounded">
                        <p class="text-green-600">Contract Detail:</p>
                        <div>
                            Name:
                            <InputText fluid placeholder="Isi nama kontrak" v-model="contractRequest.name"/>
                        </div>
                        <div class="my-1">
                            Owner:
                            <InputText fluid placeholder="Masukan nama pemilik kontrak" v-model="contractRequest.contract_owner"/>
                        </div>
                        <!-- <div class="my-1">
                            Address:
                            <InputText fluid :disabled="true" />
                        </div> -->
                    </div>
                    <div class="mt-3 border p-2 rounded">
                        <p class="text-green-600">Wallet Info:</p>
                        <div class="my-2 text-gray-500">
                            <div class="flex justify-between my-2">
                                <span>Address:</span>
                                <span>{{ shortenAddress(currentWallet.address) }}</span>
                            </div>
                            <!-- <span class="text-base text-green-600">Wallet detail:</span> -->
                            <div class="flex justify-between my-2">
                                <span class="text-gray-400">Balance</span>
                                <span class="font-bold">{{ formatBalance(currentWallet.balance?.wei) }} ETH</span>
                            </div>
                            <!-- <div class="flex justify-between my-2">
                                <span class="text-gray-400">Contract</span>
                                <span class="font-bold">{{ walletInfo?.isContract }}</span>
                            </div> -->
                            <div class="flex justify-between my-2">
                                <span class="text-gray-400">Created at</span>
                                <span class="font-bold">{{ currentWallet?.createdAt }}</span>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="border p-2 rounded">
                    <p class="text-green-600">Transaction Detail:</p>
                    <div class="mb-2">
                        Gas limit:
                        <InputText fluid :default-value="defaultValue" />
                    </div>
                    <div class="mb-2">
                        Gas price:
                        <div class="border p-2 rounded-lg bg-slate-200">{{ 20000000000 / 1e9 }} GWei</div>
                        <!-- <span class="block">{{ formatBalance(currentWallet.gas?.gasPrice) }}</span> -->
                    </div>
                    <div class="mb-2">
                        Value:
                        <div class="flex space-x-1">
                            <InputText fluid default-value="0" />
                            <Select :options="unitOptions" option-label="nama" option-value="nama" v-model="unitSelected" />
                        </div>
                    </div>
                    <div class="flex justify-between space-x-1 flex-col">
                        <div class="w-full p-2 font-semibold">
                            ABI: <span class="font-normal">{{ abiName ?? '' }}</span>
                        </div>
                        <div class="w-full p-2 font-semibold">
                            BIN: <span class="font-normal">{{ binName ?? '' }}</span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex space-x-1 justify-end">
                <div>
                    <Button label="Batal" @click="handleReset" class="w-32" />
                </div>
                <div>
                    <Button label="Deploy" severity="warn" @click="handleDeploy" class="w-32" />
                    <!-- <Button v-if="!password" label="Berikutnya" severity="warn" @click="handleForward" :disabled="!isWalletInfoAvailable" class="w-32" />
                    <Button v-else label="Deploy" severity="info" @click="handleDeploy" class="w-32" /> -->
                </div>
            </div>
        </div>

        <!-- <PasswordDialog :visible="showPasswordDialog" @update:visible="showPasswordDialog = $event" @submit="handlePasswordSubmit" /> -->
        <!-- <PasswordDialog :visible="showPasswordDialog" @submit="handlePasswordSubmit" /> -->
    </div>
</template>

<style scoped>
.contract-deployer {
    max-width: 600px;
    margin: auto;
    padding: 2rem;
    border: 1px solid #ccc;
    border-radius: 8px;
}
button {
    margin-top: 1rem;
    padding: 0.5rem 1rem;
}

.pi-eye {
    transform: scale(1.6);
    margin-right: 1rem;
}

.pi-eye-slash {
    transform: scale(1.6);
    margin-right: 1rem;
}
</style>
