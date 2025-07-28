<script setup>
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { InputText } from 'primevue';
import { computed, onMounted, ref } from 'vue';

const file = ref(null);
const status = ref('');
const isProcessing = ref(false);

const handleFileUpload = (event) => {
    file.value = event.target.files[0];
};

const scService = useSCService();
const solcVersion = ref('');
onMounted(async () => {
    solcVersion.value = await scService.getSolVersion();
    unitSelected.value = unitOptions.value[0].nama;
});
const deployContract = async () => {
    if (!file.value) {
        alert('Silakan pilih file Solidity terlebih dahulu.');
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

        console.log(response)
        const result = await response.json();
        console.log(result)
        // if (response.ok) {
        //     status.value = `Kontrak berhasil dideploy! Address: ${result.contractAddress}`;
        //     deployStatus.value = true;
        // } else {
        //     status.value = `Gagal: ${result.error || 'Terjadi kesalahan'}`;
        // }
    } catch (err) {
        status.value = `Error jaringan: ${err.message}`;
    } finally {
        isProcessing.value = false;
    }
};
const deployStatus = ref(false);
const showDialog = ref(false);
const batalBuildContract = () => {
    deployStatus.value = false;
};
const buildContract = () => {
    showDialog.value = true;
};

const isConected = ref(true);
const unitOptions = ref([
    {
        id: 1,
        nama: 'WEI'
    },
    {
        id: 2,
        nama: 'GWEI'
    },
    {
        id: 3,
        nama: 'ETH'
    }
]);
const unitSelected = ref();
const defaultValue = computed(() => {
    const defVal = 3000000;
    return defVal.toLocaleString('id-ID');
});
const { formatBalance } = useUtils();
const walletInfo = ref();
const getWallet = (e) => {
    console.log(e);
    walletInfo.value = e;
    if (walletInfo.value) {
        isConected.value = false;
    }
};
</script>

<template>
    <div class="contract-deployer">
        <h2>Deploy Smart Contract</h2>
        <div v-if="!deployStatus">
            <div>Compiler: {{ solcVersion ? solcVersion : '' }}</div>
            <input type="file" accept=".sol" @change="handleFileUpload" />
            <Button icon="pi pi-refresh" @click="deployContract" :disabled="isProcessing || !file" :label="isProcessing ? 'Memproses...' : 'Deploy Kontrak'" />
            <p v-if="status">{{ status }}</p>
        </div>
        <div v-else>
            <p>Build contract</p>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    Wallet addres:
                    <AccountComponent @wallet-info="getWallet" />
                    <div class="my-2 text-gray-500">
                        <span class="text-base text-green-600">Wallet detail:</span>
                        <div class="flex justify-between my-2">
                            <span class="text-gray-400">Balance</span>
                            <span class="font-bold">{{ formatBalance(walletInfo?.Balance.Wei) }} ETH</span>
                        </div>
                        <!-- <div class="flex justify-between my-2">
                            <span class="text-gray-400">Transaction</span>
                            <span class="font-bold">{{walletInfo}}</span>
                        </div> -->
                        <div class="flex justify-between my-2">
                            <span class="text-gray-400">Contract</span>
                            <span class="font-bold">{{ walletInfo?.isContract }}</span>
                        </div>
                        <div class="flex justify-between my-2">
                            <span class="text-gray-400">Created at</span>
                            <span class="font-bold">{{ walletInfo?.createdAt }}</span>
                        </div>
                    </div>
                </div>
                <div>
                    <div class="mb-2">
                        Gas limit:
                        <InputText fluid :default-value="defaultValue" />
                    </div>
                    <!-- <div class="mb-2">
                        Gas price:
                        <InputText fluid/>
                    </div> -->
                    <div class="mb-2">
                        Value:
                        <div class="flex space-x-1">
                            <InputText fluid default-value="0" />
                            <Select :options="unitOptions" option-label="nama" option-value="nama" v-model="unitSelected" />
                        </div>
                    </div>
                    <div>
                        Contract address:
                        <InputText fluid />
                    </div>
                </div>
            </div>
            <div class="flex space-x-1 justify-end">
                <div>
                    <Button label="Batal" @click="batalBuildContract" class="w-32" />
                </div>
                <div>
                    <Button label="Build" @click="buildContract" class="w-32" severity="warn" :disabled="isConected" />
                    <!-- <Button v-else label="Build" @click="buildContract" class="w-32" severity="warn" /> -->
                </div>
            </div>
        </div>

        <Dialog v-model:visible="showDialog" position="top">
            <h3>Apakah akan dibatalkan</h3>

            <template #footer>
                <Button label="Simpan" icon="pi pi-save" severity="info" />
            </template>
        </Dialog>
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
p {
    margin-top: 1rem;
    color: green;
}
</style>
