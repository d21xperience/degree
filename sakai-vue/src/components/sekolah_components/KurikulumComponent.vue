<script setup>
import { useKurikulum } from '@/composables/sekolah_composable/useKurikulum';
import { debounce } from 'lodash-es';
import { computed, onMounted, ref, watch } from 'vue';

const props = defineProps({
    modelValue: {
        type: [Object, String, Number, null],
        default: null
    }
});

const emit = defineEmits(['update:modelValue']);

const { kurikulumList, error, isFetching, initialize } = useKurikulum({ autoload: false });

/* -----------------------------
   INTERNAL VALUE (v-model)
--------------------------------*/
const internalValue = ref(props.modelValue);

// Sync props -> internal
watch(
    () => props.modelValue,
    (value) => (internalValue.value = value)
);

// Sync internal -> parent
watch(internalValue, (value) => emit('update:modelValue', value));

/* -----------------------------
   FETCH DATA
--------------------------------*/
const init = async () => {
    try {
        await initialize();
    } catch (err) {
        console.error('[KurikulumComponent] Error init:', err);
    }
};

onMounted(init);

/* -----------------------------
   FILTERING (AutoComplete)
--------------------------------*/
const searchTerm = ref('');

const filteredOptions = computed(() => {
    if (!searchTerm.value) return kurikulumList.value || [];

    return (kurikulumList.value || []).filter((item) => item.namaKurikulum.toLowerCase().includes(searchTerm.value.toLowerCase()));
});

// PrimeVue sends event.query
const searchKurikulum = debounce((event) => {
    searchTerm.value = event.query || '';
}, 250);

/* -----------------------------
   KEYDOWN HANDLER
--------------------------------*/
// const kurikulumHandleKeydown = (event) => {
//     if (event.key === ' ') {
//         event.preventDefault();
//         searchTerm.value += ' ';
//     }
// };
const kurikulumHandleKeydown = (event) => {
    if (event.key === ' ') {
        // Jika value masih berupa object (pilihan), ubah ke string dulu
        if (typeof internalValue.value === 'object' && internalValue.value !== null) {
            internalValue.value = internalValue.value.namaKurikulum + ' ';
        } else {
            internalValue.value = (internalValue.value || '') + ' ';
        }

        // Update searchTerm agar filtering tetap sinkron
        searchTerm.value = internalValue.value;

        // Tidak perlu preventDefault karena kita ingin input menerima spasi
        return;
    }
};

// const kurikulumHandleKeydown = (event) => {
//     if (event.key === ' ') {
//         internalValue.value += ' '; // Menambahkan spasi ke query
//     }
// };
</script>

<template>
    <div class="w-full">
        <template v-if="isFetching">
            <span class="text-sm text-gray-500">Memuat...</span>
        </template>

        <template v-else-if="error">
            <span class="text-sm text-red-500">Gagal memuat data</span>
        </template>

        <AutoComplete
            v-model="internalValue"
            :suggestions="filteredOptions"
            option-label="namaKurikulum"
            placeholder="Pilih kurikulum"
            :virtual-scroller-options="{ itemSize: 34, lazy: true }"
            dropdown
            fluid
            @complete="searchKurikulum"
            @keydown="kurikulumHandleKeydown"
        />
    </div>
</template>
