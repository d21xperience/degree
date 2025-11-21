<script setup>
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { useWalletInfo } from '@/composables/useWalletInfo';
import { onMounted, ref, watch } from 'vue';

const scService = useSCService();
const utils = useUtils();
const { fetchWalletInfo } = useWalletInfo();

const internalValue = ref(null);
const accountOptions = ref([]);

const props = defineProps({
    modelValue: {
        type: [Object, String, null],
        default: null
    }
});

const emit = defineEmits(['update:modelValue', 'walletInfo']);

// Fungsi label untuk tampilan opsi
const handleLabelOption = (option) => {
    return utils.ringkasHash(option.address, 4, 4);
};

// Sinkronisasi dari parent (v-model) ke internalValue
watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    },
    { immediate: true }
);

// Ketika user memilih akun dari Select
const onSelectionChange = async (selected) => {
    // console.log("selected", selected)
    if (!selected) return;

    try {
        // console.log("selected", selected.address)
        const response = await fetchWalletInfo({ public_address: selected.address });
        console.log('response', response);
        if (response?.walletData) {
            // Emit wallet info untuk keperluan komponen lain
            emit('walletInfo', response.walletData);

            // Hanya emit address/selected object jika ingin sinkron ke parent
            // Pastikan tipe datanya sesuai: apakah parent ingin objek atau hanya address?
            emit('update:modelValue', selected); // <-- kirim objek akun yang dipilih
        }
    } catch (error) {
        console.error('Gagal mengambil data wallet:', error);
        emit('update:modelValue', null); // reset jika gagal
    }
};

// Load daftar akun saat komponen mount
onMounted(async () => {
    try {
        const accounts = await scService.fetchBCAccount();
        accountOptions.value = Array.isArray(accounts) ? accounts : [];
    } catch (error) {
        console.error('Gagal memuat daftar akun:', error);
        accountOptions.value = [];
    }
});
</script>

<template>
    <Select v-model="internalValue" :options="accountOptions" fluid placeholder="Pilih Account" :option-label="handleLabelOption" @update:model-value="onSelectionChange" />
</template>
