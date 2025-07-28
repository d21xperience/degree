<template>
    <Select v-model="internalValue" :options="platformOptions" fluid placeholder="Pilih platform" option-label="name" />
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref } from 'vue';
const scService = useSCService();
const internalValue = ref();
const platformOptions = ref([]);
// const handleLabelOption = (newVal) => {
//     return `${newVal.name} - ${newVal.Type}`;
// };
onMounted(async () => {
    platformOptions.value = await scService.fetchNetworkPlatform();
    console.log(platformOptions.value)
});
</script>
