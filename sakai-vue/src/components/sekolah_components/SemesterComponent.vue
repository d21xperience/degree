<template>
    <Select v-model="internalValue" :options="semesterOptions" optionLabel="namaSemester" placeholder="Tahun Pelajaran" fluid checkmark :disabled="props.isDisabled" />
</template>

<script setup>
import { useToast } from 'primevue/usetoast';
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';
const { fetchSemester, initSelectedSemester } = useSekolahService();
const semesterOptions = ref([]);
const props = defineProps(['modelValue', 'isDisabled']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const toast = useToast();

const internalValue = ref();

watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});

const initial = async () => {
    try {
        semesterOptions.value = await fetchSemester();
        internalValue.value = initSelectedSemester.value
        
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil data semester: ${error}`, life: 3000 });
    }
};
onMounted(async () => {
    await initial();
});
</script>
