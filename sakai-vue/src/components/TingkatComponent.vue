<template>
    <Select v-model="internalValue" :options="tingkatPendidikanOptions" optionLabel="nama" optionValue="kode" placeholder="Pilih tingkat..." fluid checkmark />
</template>

<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';
const tingkatPendidikanOptions = ref([]);
const props = defineProps(['modelValue', 'initialValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { fetchTingkat } = useSekolahService(); 
const internalValue = ref();

watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
}); 
watch(
    () => props.initialValue,
    async (newVal) => { 
        if (newVal) {
            tingkatPendidikanOptions.value = await fetchTingkat(); 
            internalValue.value = tingkatPendidikanOptions.value.find(item => item.kode.includes(newVal)).kode 
        }
    },
    { immediate: true }
);

onMounted(async () => {
    tingkatPendidikanOptions.value = await fetchTingkat(); 
});
</script>
