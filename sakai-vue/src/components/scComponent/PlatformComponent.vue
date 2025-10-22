<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref, watch } from 'vue';
const scService = useSCService();
const internalValue = ref();
const platformOptions = ref();
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);

watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    },
    { immediate: true }
);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});

onMounted(async () => {
    platformOptions.value = await scService.fetchNetworkPlatform();
    internalValue.value = platformOptions.value.find((item) => item.active == true);
});
</script>

<template>
    <Select v-model="internalValue" :options="platformOptions" fluid placeholder="Pilih platform" option-label="name" :show-clear="true" :checkmark="true" />
</template>
