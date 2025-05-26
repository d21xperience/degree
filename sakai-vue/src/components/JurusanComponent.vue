<template>
    <AutoComplete
        v-model="internalValue"
        :suggestions="jurusanOptions"
        optionLabel="namaJurusan"
        @complete="jurusanSearch"
        @keydown.space.prevent="handleKeydown"
        placeholder="Cari Jurusan..."
        class="w-full"
        fluid
        dropdown
        :loading="jurusanLoading"
    />
</template>

<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, ref, watch } from 'vue';

const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { fetchJurusan, jurusanLoading, jurusanSearch, jurusanOptions } = useFormOptions();
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

const handleKeydown = (event) => {
    if (event.key === ' ') {
        internalValue.value += ' ';
    }
};

onMounted(async () => {
    await fetchJurusan();
});
</script>
