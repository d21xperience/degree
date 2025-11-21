<script setup>
import SekolahService from '@/service/SekolahService';
import { onMounted, ref, watch } from 'vue';

const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const sekolah = ref();
const sekolahOptions = ref();
const loading = ref(false);
// const { ptkSearch, ptkOptions, ptkLoading } = useFormOptions();
const internalValue = ref(props.modelValue);
const search = (event) => {
    setTimeout(() => {
        if (!event.query.trim().length) {
            sekolahOptions.value = [...sekolah.value];
        } else {
            loading.value = true;
            sekolahOptions.value = sekolah.value.filter((country) => country.nama_sekolah.toLowerCase().includes(event.query.toLowerCase()));
            loading.value = false;
        }
    }, 250);
};
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
onMounted(() => {
    SekolahService.getSekolah().then((data) => (sekolah.value = data));
});
</script>

<template>
    <!-- <AutoComplete v-model="internalValue" :suggestions="ptkOptions" optionLabel="nama" @complete="ptkSearch" @keydown.space.prevent="handleKeydown" placeholder="Masukan nama..." class="w-full" fluid :loading="ptkLoading" /> -->

    <FloatLabel variant="on">
        <IconField>
            <AutoComplete v-model="internalValue" option-label="nama_sekolah" :suggestions="sekolahOptions" fluid :loading="loading" @complete="search" @keydown.space.prevent="handleKeydown" />
            <InputIcon class="pi pi-building-columns" />
        </IconField>
        <label>NPSN/Nama Sekolah</label>
    </FloatLabel>
</template>
