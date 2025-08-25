<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { ref, watch } from 'vue';
const props = defineProps(['modelValue', 'jurusanInduk']);
const emit = defineEmits(['update:modelValue']);

const { fetchJurusan } = useFormOptions();
const internalValue = ref(props.modelValue);
const jurusanOptions = ref();
const handleLabelOption = (newVal) => {
    return `${newVal.jurusanId} - ${newVal.namaJurusan}`;
};
watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    },
    { immediate: true } // untuk memuat data saat komponen pertama kali dipasang
);
watch(
    () => props.jurusanInduk,
    async (newVal) => {
        if (newVal) {
            jurusanOptions.value = await fetchJurusan(newVal);
        }
    }
);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
</script>
<template>
    <Select v-model="internalValue" :options="jurusanOptions" :option-label="handleLabelOption" placeholder="Pilih kompetensi keahlian" fluid />
</template>
