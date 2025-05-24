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
                <Button label="Simpan" severity="success" @click="submitForm" :loading="isLoading" class="min-w-28" />
                <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" />
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';

const submitForm = () => {};
const batal = () => {};
const isLoading = ref(false);
</script>
