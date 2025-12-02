<script setup>
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref } from 'vue';
const { listTahunAjaran, selectedTahunAjaran, fetchSemester } = useSemester({ autoload: false });
const tahunAjaranOptions = ref([]);
const props = defineProps(['modelValue', 'isDisabled']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const toast = useToast();
const { schemaname } = useTableTenant({ autoload: true });
// const internalValue = ref();

// watch(internalValue, (newVal) => {
//     emit('update:modelValue', newVal);
// });

const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

const initial = async () => {
    console.log('memanggil initial tahunAjaranComponent');
    try {
        tahunAjaranOptions.value = listTahunAjaran.value;
        internalValue.value = selectedTahunAjaran.value;
        // console.log(internalValue.value);
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil data tahunAjaran: ${error}`, life: 3000 });
    }
};
onMounted(async () => {
    // await initial();
    await fetchSemester();
});
</script>

<template>
    <div class="flex min-w-72 items-center space-x-2">
        <div class="w-40">
            <label for="">Tahun Ajaran</label>
        </div>
        <Select v-model="internalValue" :options="tahunAjaranOptions" option-label="tahunAjaranId" placeholder="" fluid checkmark :disabled="props.isDisabled" class="w-36" />
    </div>
</template>
