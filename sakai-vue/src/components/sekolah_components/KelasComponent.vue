<template>
    <Select v-model="internalValue" :options="kelasOptions" optionLabel="nmKelas" optionValue="nmKelas" placeholder="Pilih Kelas..." fluid checkmark />
</template>

<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';
const kelasOptions = ref([]);
const props = defineProps(['modelValue', 'initialValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { fetchKelas } = useSekolahService();
const internalValue = ref();

watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
watch(
    () => props.initialValue,
    async (newVal) => {
        if (newVal) {
            kelasOptions.value = await fetchKelas();
            internalValue.value = kelasOptions.value.find((item) => item.kode.includes(newVal)).kode;
        }
    },
    { immediate: true }
);

onMounted(async () => {
    kelasOptions.value = await fetchKelas();
});
</script>
