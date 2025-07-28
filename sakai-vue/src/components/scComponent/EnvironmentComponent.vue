<template>
    <Select v-model="internalValue" :options="environmentOptions" fluid placeholder="Pilih Environment" :option-label="handleLabelOption" />
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { ref, watch } from 'vue';
const scService = useSCService();
const props = defineProps(['modelValue', 'architecture']);
const emit = defineEmits(['update:modelValue']);

const internalValue = ref(props.modelValue);
const environmentOptions = ref([]);
const handleLabelOption = (newVal) => {
    return `${newVal.Architecture} | ${newVal.Name} - ${newVal.Type}`;
};
watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    },
    { immediate: true } // untuk memuat data saat komponen pertama kali dipasang
);
watch(
    () => props.architecture,
    async (newVal) => {
        // console.log(newVal.name.toLowerCase() == 'hyperledger fabric');
        if (newVal) {
            if (newVal.name.toLowerCase() == 'hyperledger fabric') {
                environmentOptions.value = await scService.fetchBCNetworks({ Architecture: 'NONEVM' });
            } else {
                environmentOptions.value = await scService.fetchBCNetworks({ Architecture: 'EVM' });
            }
        }
    }
);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
</script>
