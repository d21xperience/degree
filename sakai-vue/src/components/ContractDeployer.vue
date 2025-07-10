<script setup>
import { ref } from 'vue';

const file = ref(null);
const status = ref('');
const isProcessing = ref(false);

const handleFileUpload = (event) => {
    file.value = event.target.files[0];
};

const deployContract = async () => {
    if (!file.value) {
        alert('Silakan pilih file Solidity terlebih dahulu.');
        return;
    }

    const formData = new FormData();
    formData.append('contract', file.value);

    isProcessing.value = true;
    status.value = 'Mengunggah dan mendeploy kontrak...';

    try {
        const response = await fetch('http://localhost:8080/deploy', {
            method: 'POST',
            body: formData
        });

        const result = await response.json();

        if (response.ok) {
            status.value = `Kontrak berhasil dideploy! Address: ${result.contractAddress}`;
        } else {
            status.value = `Gagal: ${result.error || 'Terjadi kesalahan'}`;
        }
    } catch (err) {
        status.value = `Error jaringan: ${err.message}`;
    } finally {
        isProcessing.value = false;
    }
};
</script>

<template>
    <div class="contract-deployer">
        <h2>Deploy Smart Contract</h2>
        <input type="file" accept=".sol" @change="handleFileUpload" />
        <button @click="deployContract" :disabled="isProcessing || !file">
            {{ isProcessing ? 'Memproses...' : 'Deploy Kontrak' }}
        </button>
        <p v-if="status">{{ status }}</p>
    </div>
</template>

<style scoped>
.contract-deployer {
    max-width: 500px;
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
