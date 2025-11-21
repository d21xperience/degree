<script setup>
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useToast } from 'primevue/usetoast';

import { computed, onMounted, ref } from 'vue';
const jenjangPendidikanOptions = ref([]);
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const loadingJenjangPendidikan = ref(false);
// const { initSelectedSemester } = useSemester();

const toast = useToast();
const { fetchJenjangPendidikan } = useSekolah();
const initial = async () => {
    loadingJenjangPendidikan.value = true;
    try {
        const response = await fetchJenjangPendidikan({ isJenjangLembaga: true, jenjangLembaga: 1 });
        if (response.status) {
            jenjangPendidikanOptions.value = response.jenjang;
            toast.add({ severity: 'success', summary: 'Success', detail: `Berhasil mengambil data jenjang`, life: 3000 });
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil kelas: ${error}`, life: 3000 });
    } finally {
        loadingJenjangPendidikan.value = false;
    }
};
// watch(internalValue, (newVal) => {
//     emit('update:modelValue', newVal);
// });

const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

onMounted(async () => {
    await initial();
});
</script>

<template>
    <Select
        v-model="internalValue"
        :options="jenjangPendidikanOptions"
        option-label="nama"
        :placeholder="!loadingJenjangPendidikan ? 'Pilih Jenjang Pendidikan...' : 'Memuat data..'"
        fluid
        checkmark
        :show-clear="true"
        :loading="loadingJenjangPendidikan"
    />
</template>
