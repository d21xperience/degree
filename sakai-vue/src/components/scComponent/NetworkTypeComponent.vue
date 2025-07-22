<template>
    <Select v-model="internalValue" :options="networkTypeOptions" fluid placeholder="Pilih tipe Network" option-label="Type" option-value="Type" />
</template>

<script setup>
import { ref, watch } from 'vue';
const internalValue = ref();
const props = defineProps(['modelValue', 'initialValue']); // props dari parent
const emit = defineEmits(['update:modelValue']); // emit update ke parent
const networkTypeOptions = ref([
    {
        Id: 1,
        Type: 'mainnet'
    },
    {
        Id: 2,
        Type: 'testnet'
    },
    {
        Id: 3,
        Type: 'local'
    },
    {
        Id: 4,
        Type: 'private'
    },
]);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});

watch(
    () => props.initialValue,
    async (newVal) => {
        // console.log(newVal);
        if (newVal) {
            // networkTypeOptions.value = await fetchTingkat();
            internalValue.value = networkTypeOptions.value.find((item) => item.Type.includes(newVal)).Type;
        }
    },
    { immediate: true }
);
</script>
