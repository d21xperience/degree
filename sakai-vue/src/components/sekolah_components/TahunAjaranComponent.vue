<script setup>
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { computed, onMounted, ref, watch } from 'vue';

const {
    listTahunAjaran,
    selectedTahunAjaran,
    isLoading,
    error,
    initialize // ← ✅ gunakan ini!
} = useSemester({ autoload: false }); // ❗ nonaktifkan autoload agar kita kendalikan manual

// const props = defineProps(['modelValue', 'isDisabled']);
const props = defineProps({
    modelValue: {
        type: [Object, String, Number, null],
        default: null
    },
    isDisabled: Boolean
});
const emit = defineEmits(['update:modelValue']);

// 🔁 Sync 2-way dengan props
const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => {
        selectedTahunAjaran.value = value;
        emit('update:modelValue', value);
    }
});

// 📋 Opsional: state lokal untuk UI (jika perlu loading/error di komponen ini)
const tahunAjaranOptions = ref([]);

// ✅ Inisialisasi eksplisit — tunggu sampai selesai
const initComponent = async () => {
    try {
        // 🔁 Jalankan inisialisasi semester (termasuk fetch tahun ajaran)
        await initialize();

        // ✅ Setelah initialize(), data sudah siap
        tahunAjaranOptions.value = listTahunAjaran.value;

        // Jika parent belum tentukan modelValue, fallback ke selected dari store
        // console.log('cek props.modelValue', props.modelValue == null);
        if (props.modelValue == null && selectedTahunAjaran.value != null) {
            internalValue.value = selectedTahunAjaran.value;
            // console.log('selectedTahunAjaran tahun ajaran ...', selectedTahunAjaran.value);
            // console.log('internalValue tahun ajaran ...', internalValue.value);
        }
    } catch (err) {
        console.error('[TahunAjaranComponent] Gagal inisialisasi:', err);
    }
};

// 🚀 Jalankan saat mounted
onMounted(() => {
    initComponent();
});

// 🔁 Watch perubahan list (opsional, jika data bisa berubah dinamis)
watch(listTahunAjaran, (newList) => {
    tahunAjaranOptions.value = newList;
});
</script>

<template>
    <div class="flex min-w-72 items-center space-x-2">
        <div class="w-40">
            <label for="">Tahun Ajaran</label>
        </div>

        <!-- ✅ Handle loading & error secara eksplisit -->
        <template v-if="isLoading">
            <span class="text-sm text-gray-500">Memuat...</span>
        </template>
        <template v-else-if="error">
            <span class="text-sm text-red-500">Gagal memuat</span>
        </template>
        <template v-else>
            <Select v-model="internalValue" :options="tahunAjaranOptions" option-label="label" placeholder="Pilih tahun ajaran" fluid :disabled="props.isDisabled || tahunAjaranOptions.length === 0" class="w-36" />
        </template>
    </div>
</template>
