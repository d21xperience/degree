<script setup>
import { useKelas } from '@/composables/sekolah_composable/useKelas';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useToast } from 'primevue/usetoast';

import { computed, onMounted, ref, watch } from 'vue';
const kelasOptions = ref([]);
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const loadingKelas = ref(false);
const { fetchKelas } = useKelas();
const { initSelectedSemester } = useSemester();
// const internalValue = ref();
const toast = useToast();

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
// watch(internalValue, (newVal) => {
//     emit('update:modelValue', newVal);
// });

const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

watch(initSelectedSemester, async (newVal) => {
    console.log(newVal);
    if (newVal) {
        loadingKelas.value = true;
        try {
            // kelasOptions.value = []
            internalValue.value = null;
            await initial();
        } catch (error) {
            console.error(error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil kelas: ${error}`, life: 3000 });
        } finally {
            loadingKelas.value = false;
        }
    }
});

onMounted(async () => {
    await initial();
});
</script>

<template>
    <Select v-model="internalValue" :options="kelasOptions" option-label="nmKelas" :placeholder="!loadingKelas ? 'Pilih Kelas...' : 'Memuat data..'" fluid checkmark :show-clear="true" :loading="loadingKelas" />
</template>
