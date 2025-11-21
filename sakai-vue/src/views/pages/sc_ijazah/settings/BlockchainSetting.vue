<script setup>
import { ref } from 'vue';

const submitForm = () => {};
const batal = () => {};
const isLoading = ref(false);
</script>
<template>
    <h1>helelo from setting blockchain</h1>
    <div class="card">
        <h5>Smart Contract</h5>
        <div>
            <div>
                Smart Contarct Address
                <InputText fluid placeholder="Masukan alamat smart contract" />
            </div>

            <div class="flex justify-end mt-8 space-x-4">
                <Button label="Simpan" severity="success" :loading="isLoading" class="min-w-28" @click="submitForm" />
                <Button label="Batal" severity="secondary" class="min-w-28" @click="batal" />
            </div>
        </div>
    </div>
</template>
