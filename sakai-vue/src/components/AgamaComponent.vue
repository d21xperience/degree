<template>
    <Select v-model="internalValue" :options="agamaOptions" option-label="label" placeholder="Pilih Agama" fluid option-value="value" />
</template>
<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, ref, watch } from 'vue';
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

const useFormOption = useFormOptions();
const internalValue = ref(props.modelValue);
const agamaOptions = ref();
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
    agamaOptions.value = useFormOption.agamaOptions.value;
});
</script>
