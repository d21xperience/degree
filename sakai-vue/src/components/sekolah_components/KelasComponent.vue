<script setup>
import { useKelas } from '@/composables/sekolah_composable/useKelas';

import { computed, onMounted, ref, watch } from 'vue';
const kelasOptions = ref([]);
const props = defineProps({
    modelValue: {
        type: [Object, String, Number, null],
        default: null
    },
    initSelectedSemester: {
        type: Object,
        required: true
    }
}); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const loadingKelas = ref(false);
const { fetchKelas, isFetching, isError } = useKelas();
// const { initSelectedSemester } = useSemester();
// const internalValue = ref();

const initial = async () => {
    const response = await fetchKelas(props.initSelectedSemester);
    console.log(response);
    kelasOptions.value = response;
};
// watch(internalValue, (newVal) => {
//     emit('update:modelValue', newVal);
// });

const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

watch(props.initSelectedSemester, async (newVal) => {
    console.log(newVal);
    if (newVal) {
        loadingKelas.value = true;
        try {
            // kelasOptions.value = []
            internalValue.value = null;
            await initial();
        } catch (error) {
            console.error(error);
        } finally {
            loadingKelas.value = false;
        }
    }
});

onMounted(async () => {
    await initial();
});
</script>

<template>
    <div class="flex min-w-72 items-center space-x-2">
        <template v-if="isFetching">
            <span class="text-sm text-gray-500">Memuat...</span>
        </template>
        <template v-else-if="isError">
            <span class="text-sm text-red-500">Gagal memuat</span>
        </template>
        <template v-else>
            <Select v-model="internalValue" :options="kelasOptions" option-label="nmKelas" :placeholder="!loadingKelas ? 'Pilih Kelas...' : 'Memuat data..'" fluid checkmark :show-clear="true" :loading="loadingKelas" />
        </template>
    </div>
</template>
