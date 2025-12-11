<script setup>
import { useEducationLevel } from '@/composables/sekolah_composable/useEducationLevel';
import { computed, ref, watch } from 'vue';

const { fetchEducationLevel } = useEducationLevel();

const props = defineProps({
    modelValue: {
        type: String,
        default: ''
    },
    jenjangPendidikanId: {
        type: Number,
        default: 0,
        required: true
    },
    isDisabled: {
        type: Boolean,
        default: false
    }
});

// ['modelValue', 'initialValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent

const tingkatPendidikanOptions = ref([]);
const internalValue = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
});

watch(
    () => props.jenjangPendidikanId,
    async (newVal) => {
        if (newVal) {
            tingkatPendidikanOptions.value = await fetchEducationLevel(newVal);
        }
    },
    { immediate: true }
);
</script>

<template>
    <div class="w-52">
        <Select v-model="internalValue" :options="tingkatPendidikanOptions" option-label="nama" option-value="kode" placeholder="Pilih tingkat..." fluid checkmark :disabled="props.isDisabled" :show-clear="true" />
    </div>
</template>
