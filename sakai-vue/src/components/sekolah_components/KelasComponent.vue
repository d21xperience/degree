<template>
    <Select v-model="internalValue" :options="kelasOptions" optionLabel="nmKelas" placeholder="Pilih Kelas..." fluid checkmark :showClear="true" />
</template>

<script setup>
import { useToast } from 'primevue/usetoast';

import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';
const kelasOptions = ref([]);
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { fetchKelas } = useSekolahService();
const internalValue = ref();
const toast = useToast();

watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});

const initial = async () => {
    try {
        const response = await fetchKelas();
        if (response.status) {
            kelasOptions.value = response.kelas;
            toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil kelas: ${error}`, life: 3000 });
    }
};
onMounted(async () => {
    await initial();
});
</script>
