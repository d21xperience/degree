<template>
    <AutoComplete
        v-model="internalValue"
        :suggestions="kurikulumOptions"
        option-label="namaKurikulum"
        @complete="searchKurikulum(query)"
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
// const cek = (data) => {
//     return `${data.kurikulumId} | ${data.namaKurikulum}`;
// };
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

const { fetchKurikulum, kurikulumLoading, searchKurikulum, kurikulumOptions } = useFormOptions();
const internalValue = ref(props.modelValue);
watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    }
);
const query = ref();
watch(internalValue, (newVal) => {
    if (typeof newVal == 'object') {
        query.value = newVal?.namaKurikulum;
    } else if (typeof newVal == 'string') {
        query.value = newVal;
    }
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
