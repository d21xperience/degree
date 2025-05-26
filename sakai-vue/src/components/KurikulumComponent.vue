<template>
    <AutoComplete
        v-model="internalValue"
        :suggestions="kurikulumOptions"
        option-label="namaKurikulum"
        @complete="kurikulumSearch"
        @keydown.space.prevent="kurikulumHandleKeydown"
        placeholder="Pilih kurikulum"
        fluid
        dropdown
        :loading="kurikulumLoading"
    />
</template>
<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, ref, watch } from 'vue';

const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

const { fetchKurikulum, kurikulumLoading, kurikulumSearch, kurikulumOptions } = useFormOptions();
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

const kurikulumHandleKeydown = (event) => {
    if (event.key === ' ') {
        internalValue.value += ' '; // Menambahkan spasi ke query
    }
};
onMounted(async () => {
    await fetchKurikulum();
});
</script>
