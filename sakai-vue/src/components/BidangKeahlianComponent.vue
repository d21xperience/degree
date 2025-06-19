<template>
    <Select v-model="internalValue" :options="bidangKeahlianOptions" :option-label="handleLabelOption" placeholder="Pilih bidang keahlian" fluid show-clear />
</template>
<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, ref, watch } from 'vue';
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

const { fetchBidangKeahlian } = useFormOptions();
const internalValue = ref(props.modelValue);
const bidangKeahlianOptions = ref();
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
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
onMounted(async () => {
    bidangKeahlianOptions.value = await fetchBidangKeahlian();
});
</script>
