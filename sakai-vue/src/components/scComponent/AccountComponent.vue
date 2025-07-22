<template>
    <Select v-model="internalValue" :options="accountOptions" fluid placeholder="Pilih Account" :option-label="handleLabelOption" />
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref } from 'vue';
const scService = useSCService();
const internalValue = ref();
const accountOptions = ref([]);
const handleLabelOption = (newVal) => {
    // console.log(newVal)
    // return `${newVal.Name} - ${newVal.Type}`;
};
onMounted(async () => {
    accountOptions.value = await scService.fetchBCNetworks();
    // console.log(accountOptions.value);
});
</script>
