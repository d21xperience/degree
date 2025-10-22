<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, ref, watch } from 'vue';
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

const useFormOption = useFormOptions();
const internalValue = ref(props.modelValue);
const jenisKelaminOptions = ref();
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
onMounted(() => {
    jenisKelaminOptions.value = useFormOption.jenisKelaminOptions.value;
});
</script>
<template>
    <Select v-model="internalValue" :options="jenisKelaminOptions" option-label="label" placeholder="Pilih Jenis kelamin" fluid option-value="value" />
</template>
