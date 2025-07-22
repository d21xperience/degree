<template>
    <Select v-model="internalValue" :options="environmentOptions" fluid placeholder="Pilih Environment" :option-label="handleLabelOption"/>
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref } from 'vue';
const scService = useSCService();
const internalValue = ref();
const environmentOptions = ref([]);
const handleLabelOption = (newVal) => {
    // console.log(newVal)
    return `${newVal.Name} - ${newVal.Type}`;
}
onMounted(async () => {
    environmentOptions.value = await scService.fetchBCNetworks();
    // console.log(environmentOptions.value);
});
</script>
