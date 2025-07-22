<script setup>
import { useSCService } from '@/composables/useSCService';
import { InputText } from 'primevue';
import { onMounted, ref } from 'vue';

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
    console.log(solcVersion.value);
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

        const result = await response.json();

        if (response.ok) {
            status.value = `Kontrak berhasil dideploy! Address: ${result.contractAddress}`;
            deployStatus.value = true;
        } else {
            status.value = `Gagal: ${result.error || 'Terjadi kesalahan'}`;
        }
    } catch (err) {
        status.value = `Error jaringan: ${err.message}`;
    } finally {
        isProcessing.value = false;
    }
};
const deployStatus = ref(true);
const showDialog = ref(false);
const batalBuildContract = () => {
    deployStatus.value = false;
};
const buildContract = () => {
    showDialog.value = true;
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
                    Environment:
                    <EnvironmentComponent />
                </div>
                <div>
                    Account:
                    <AccountComponent />
                </div>
                <div>
                    Gas limit:
                    <InputText fluid />
                </div>
                <div>
                    Value:
                    <InputText fluid />
                </div>
                <div>
                    Contract address:
                    <InputText fluid />
                </div>
            </div>
            <div class="flex space-x-1 justify-end">
                <div>
                    <Button label="Batal" @click="batalBuildContract" class="w-32" />
                </div>
                <div>
                    <Button label="Build" @click="buildContract" class="w-32" severity="danger" />
                </div>
            </div>
        </div>

        <Dialog v-model:visible="showDialog">
            <h3>Apakah akan dibatalkan</h3>
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
