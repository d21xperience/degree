<template>
    <Select v-model="internalValue" :options="kurikulumOptionsSekolah" option-label="namaKurikulum" placeholder="Pilih kurikulum" fluid showClear />
</template>
<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, ref, watch } from 'vue';

const kurikulumOptionsSekolah = ref([]);
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);
const { fetchSekolah, fetchKategoriSekolah, createKategoriSekolah, deleteKategoriSekolah, updateKategoriSekolah, initSelectedSemester, fetchSemester } = useSekolahService();

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
    kurikulumOptionsSekolah.value = await fetchKategoriSekolah();
});
</script>
