<script setup>
import { useEducationForm } from '@/composables/sekolah_composable/useEducationForm';
import { computed, onMounted, ref } from 'vue';

// ✅ Ambil API lengkap — termasuk `initialize` dan state loading/error
const { initialize, isFetching, data, error } = useEducationForm({ autoload: false }); // ❗ nonaktifkan autoload agar kita kendalikan manual

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
        // selectedTahunAjaran.value = value;
        emit('update:modelValue', value);
    }
});

// 📋 Opsional: state lokal untuk UI (jika perlu loading/error di komponen ini)
const bentukPendidikanOptions = ref([]);

// ✅ Inisialisasi eksplisit — tunggu sampai selesai
const initComponent = async () => {
    try {
        // 🔁 Jalankan inisialisasi bentuk pendidikan
        await initialize();

        // ✅ Setelah initialize(), data sudah siap
        bentukPendidikanOptions.value = data.value;
        // console.log(data.value);
    } catch (err) {
        console.error('[BentukPendidikanComponent] Gagal inisialisasi:', err);
    }
};

// 🚀 Jalankan saat mounted
onMounted(() => {
    initComponent();
});
</script>

<template>
    <div class="flex min-w-72 items-center space-x-2">
        <!-- ✅ Handle loading & error secara eksplisit -->
        <template v-if="isFetching">
            <span class="text-sm text-gray-500">Memuat...</span>
        </template>
        <template v-else-if="error">
            <span class="text-sm text-red-500">Gagal memuat</span>
        </template>
        <template v-else>
            <Select v-model="internalValue" :options="bentukPendidikanOptions" option-label="nama" option-value="nama" placeholder="Pilih bentuk pendidikan" fluid :disabled="props.isDisabled || bentukPendidikanOptions.length === 0" class="w-36" />
        </template>
    </div>
</template>
