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
        // console.log('EnvComponent', newVal);
        if (newVal) {
            if (newVal.name.toLowerCase() == 'hyperledger fabric') {
                environmentOptions.value = await scService.fetchBCNetworks('NONEVM');
            } else {
                environmentOptions.value = await scService.fetchBCNetworks('EVM');
                // console.log(environmentOptions.value);
                internalValue.value = environmentOptions.value.find((item) => item.Activate == true);
            }
        }
    }
);
watch(internalValue, (newVal) => {
    emit('update:modelValue', newVal);
});
</script>

<template>
    <Select v-model="internalValue" :options="environmentOptions" fluid placeholder="Pilih Environment" :option-label="handleLabelOption" :show-clear="true" :checkmark="true" />
</template>
