<script setup>
import { computed, ref, toRefs, watch } from 'vue';
import { useStore } from 'vuex';

// Store
const store = useStore();

// Props & Emits
const props = defineProps({
    modelValue: {
        type: [String, Number, Object],
        default: null
    },
    isDisabled: {
        type: Boolean,
        default: false
    },
    tahunAjaranId: {
        type: [String, Number],
        default: null
    }
});

const emit = defineEmits(['update:modelValue']);

// Destructure props with toRefs for reactivity
const { modelValue, isDisabled, tahunAjaranId } = toRefs(props);

// Computed
const semesterList = computed(() => store.getters['semesterService/getSemester']);

// Reactive state
const semesterOptions = ref([]);
const internalValue = ref(modelValue.value);

// Watchers
// Watch for external modelValue changes
watch(modelValue, (newValue) => {
    if (newValue !== internalValue.value) {
        internalValue.value = newValue;
    }
});

// Watch for internalValue changes to emit to parent
watch(internalValue, (newValue) => {
    if (newValue !== modelValue.value) {
        emit('update:modelValue', newValue);
    }
});
const isLoading = ref(false);
// Watch for tahunAjaranId changes to filter options
watch(
    tahunAjaranId,
    (newTahunAjaranId) => {
        if (!newTahunAjaranId) {
            semesterOptions.value = [];
            return;
        }
        semesterOptions.value = semesterList.value.filter((semester) => semester.tahunAjaranId == newTahunAjaranId);
    },
    { immediate: true }
);

// Optional: Reset internal value when tahunAjaranId changes
watch(tahunAjaranId, () => {
    internalValue.value = null;
});
</script>

<template>
    <div class="w-full">
        <!-- Loading State -->
        <div v-if="isLoading" class="text-sm text-gray-500">Memuat...</div>

        <!-- Error State -->
        <!-- <div v-else-if="error" class="text-sm text-red-500">Gagal memuat data</div> -->

        <!-- Normal State -->
        <Select
            v-else
            v-model="internalValue"
            :options="semesterOptions"
            option-label="namaSemester"
            placeholder="Pilih Semester"
            :disabled="isDisabled"
            class="w-full md:w-56"
            fluid
            checkmark
            :class="{ 'opacity-50 cursor-not-allowed': isDisabled }"
        />

        <!-- Debug info (remove in production) -->
        <!-- <div v-if="process.env.NODE_ENV === 'development'" class="mt-2 text-xs text-gray-400">Options: {{ semesterOptions.length }}, Selected: {{ internalValue }}</div> -->
    </div>
</template>

<style scoped>
/* Optional: Add any component-specific styles here */
</style>
