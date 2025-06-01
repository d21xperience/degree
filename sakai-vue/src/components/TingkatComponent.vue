<template>
    <Select v-model="internalValue" :options="tingkatPendidikanOptions" optionLabel="nama" optionValue="kode" placeholder="Pilih tingkat..." fluid showClear checkmark />
</template>

<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';
const tingkatPendidikanOptions = ref([]);
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { fetchTingkat } = useSekolahService();
const internalValue = ref(props.modelValue);

watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    }
);

watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});

onMounted(async () => {
    tingkatPendidikanOptions.value = await fetchTingkat();
});
</script>
