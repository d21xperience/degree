<template>
    <Select v-model="internalValue" :options="programKeahlianOptions" :option-label="handleLabelOption" placeholder="Pilih program keahlian" fluid />
</template>
<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { ref, watch } from 'vue';
const props = defineProps(['modelValue', 'jurusanInduk']);
const emit = defineEmits(['update:modelValue']);

const { fetchProgramKeahlian } = useFormOptions();
const internalValue = ref(props.modelValue);
const programKeahlianOptions = ref();
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
            programKeahlianOptions.value = await fetchProgramKeahlian(newVal);
        }
    }
);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
</script>
