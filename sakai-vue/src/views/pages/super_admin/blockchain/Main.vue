<template>
    <div>
        <div class="">
            <div class="flex justify-between p-2">
                <div class="flex items-center space-x-2">
                    <label class="text-gray-500 font-bold">Platform:</label>
                    <span :class="!platformSelected ? 'text-red-500' : ''">{{ !platformSelected ? 'Unselected' : platformSelected.name }}</span>
                    <div class="flex items-center space-x-2">
                        <div class="h-3 w-3 rounded-full" :class="isConnected ? 'bg-green-500' : 'bg-red-500'"></div>
                        <span>{{ isConnected ? 'Connected' : 'Disconnected' }}</span>
                    </div>
                </div>
                <div class="md:flex md:items-center md:space-x-2">
                    <label class="text-gray-500 font-bold">Network:</label>
                    <span :class="!networkSelected ? 'text-red-500' : ''">{{ !networkSelected ? 'Disconected' : networkSelected.Name }}</span>
                </div>
            </div>
        </div>
        <div class="card">
            <RouterView />
        </div>
    </div>
</template>
<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref, watchEffect } from 'vue';
const scService = useSCService();
const networkSelected = ref(null);
const platformSelected = ref(null);
const isConnected = ref(false)
watchEffect(async () => {
    platformSelected.value = await scService.getNetowrkPlatform();
    networkSelected.value = await scService.getBCNetwork();
    isConnected.value = await scService.getBCConnected()
    // console.log(isConnected.value)
});

onMounted(async () => {
    isConnected.value = await scService.getBCConnected()
    // platforms.value = await scService.fetchNetworkPlatform();
});
</script>
