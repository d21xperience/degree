<script setup>
import axios from 'axios';
import Button from 'primevue/button';
import InputTextarea from 'primevue/inputtextarea';
import Message from 'primevue/message';
import { ref } from 'vue';

// Reactive state
const abiJson = ref('');
const bytecodeHex = ref('');
const loading = ref(false);
const result = ref(null);
const errorMsg = ref('');

// Configure Axios – fallback to localhost when .env not set
const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
    headers: { 'Content-Type': 'application/json' }
});

async function deploy() {
    errorMsg.value = '';
    result.value = null;

    if (!abiJson.value || !bytecodeHex.value) {
        errorMsg.value = 'ABI dan bytecode wajib diisi.';
        return;
    }

    loading.value = true;
    try {
        const { data } = await api.post('/utils/deploy', {
            abiJSON: abiJson.value,
            bytecodeHex: bytecodeHex.value
        });

        // Sesuaikan properti sesuai respons backend Anda
        result.value = {
            address: data.address || data.Address,
            txHash: data.txHash || data.TxHash
        };
    } catch (err) {
        errorMsg.value = err.response?.data?.error || err.message;
    } finally {
        loading.value = false;
    }
}
</script>

<template>
    <div class="max-w-3xl mx-auto p-6">
        <h2 class="text-2xl font-semibold mb-6">Deploy Smart Contract</h2>

        <div class="mb-4">
            <label class="block font-medium mb-1">ABI (JSON)</label>
            <!-- PrimeVue InputTextarea automatically resizes -->
            <InputTextarea v-model="abiJson" rows="8" auto-resize placeholder="{ ... }" class="w-full" />
        </div>

        <div class="mb-4">
            <label class="block font-medium mb-1">Bytecode (hex, 0x…)</label>
            <InputTextarea v-model="bytecodeHex" rows="5" auto-resize placeholder="0x600360…" class="w-full" />
        </div>

        <Button :loading="loading" label="Deploy" class="w-full" @click="deploy" />

        <!-- Result -->
        <div v-if="result" class="mt-6 p-4 border rounded-lg bg-gray-50 dark:bg-gray-800">
            <p class="font-medium">Contract Address:</p>
            <p class="break-all">{{ result.address }}</p>
            <p class="font-medium mt-3">Transaction Hash:</p>
            <p class="break-all">{{ result.txHash }}</p>
        </div>

        <!-- Error -->
        <Message v-if="errorMsg" severity="error" class="mt-4">{{ errorMsg }}</Message>
    </div>
</template>

<style scoped></style>
