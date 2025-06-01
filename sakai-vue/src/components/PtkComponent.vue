<template>
    <AutoComplete v-model="internalValue" :suggestions="ptkOptions" optionLabel="nama" @complete="ptkSearch" @keydown.space.prevent="handleKeydown" placeholder="Masukan nama..." class="w-full" fluid :loading="ptkLoading" />
</template>

<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { ref, watch } from 'vue';

const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const { ptkSearch, ptkOptions, ptkLoading } = useFormOptions();
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
// onMounted(async () => {
//     await fetchJurusan();
// });
</script>
