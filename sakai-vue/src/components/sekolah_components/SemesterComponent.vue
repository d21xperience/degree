<script setup>
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref } from 'vue';
const { listSemester, selectedSemester } = useSemester();
const semesterOptions = ref([]);
const props = defineProps(['modelValue', 'isDisabled']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const toast = useToast();

// const internalValue = ref();

// watch(internalValue, (newVal) => {
//     emit('update:modelValue', newVal);
// });

const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

const initial = async () => {
    try {
        semesterOptions.value = listSemester.value;
        // console.log(semesterOptions.value);
        internalValue.value = selectedSemester;
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil data semester: ${error}`, life: 3000 });
    }
};
onMounted(async () => {
    await initial();
});
</script>

<template>
    <Select v-model="internalValue" :options="semesterOptions" option-label="namaSemester" placeholder="Semester" fluid checkmark :disabled="props.isDisabled" class="w-full md:w-56" />
</template>
