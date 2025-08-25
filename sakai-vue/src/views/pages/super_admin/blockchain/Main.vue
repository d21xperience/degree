<script setup>
import { useSCService } from '@/composables/useSCService';
import { useWebSocket } from '@/composables/useWebSocket';
import { onMounted, ref, watchEffect } from 'vue';
const scService = useSCService();
const networkSelected = ref(null);
const platformSelected = ref(null);
// const isConnected = ref(false);
watchEffect(async () => {
    platformSelected.value = await scService.getNetowrkPlatform();
    networkSelected.value = await scService.getBCNetwork();
    isConnected.value = await scService.getBCConnected();
    // console.log(isConnected.value)
});
const isShowConnected = ref(false);
const showConnected = () => {
    isShowConnected.value = true;
};
const handleCancel = () => {
    isShowConnected.value = false;
};
const handleConnect = async () => {
    const response = await scService.setBCConfig();
    if (response.status) {
        isShowConnected.value = false;
        isConnected.value = true;
    }
    // console.log(response)
};
const { isConnected } = useWebSocket('ws://localhost:8080/ws');
onMounted(async () => {
    isConnected.value = await scService.getBCConnected();
    console.log(isConnected.value);
    // platforms.value = await scService.fetchNetworkPlatform();
});

// // Cleanup WebSocket on component unmount
// onBeforeUnmount(() => {
//     if (ws.value) {
//         ws.value.close();
//     }
// });

// // Initialize on component mount
// onMounted(() => {
//     initWebSocket();
// });
</script>
<template>
    <div>
        <div class="">
            <div class="flex justify-between p-2">
                <div class="flex items-center space-x-2">
                    <label class="text-gray-500 font-bold">Platform:</label>
                    <span :class="!platformSelected ? 'text-red-500' : ''">{{ !platformSelected ? 'Unselected' : platformSelected.name }}</span>
                </div>
                <div class="md:flex md:items-center md:space-x-2">
                    <label class="text-gray-500 font-bold">Network:</label>
                    <span :class="!networkSelected ? 'text-red-500' : ''">{{ !networkSelected ? 'Unselected' : networkSelected.Name }}</span>
                    <label class="text-gray-500 font-bold">Status:</label>
                    <div class="flex items-center space-x-2 hover:text-red-500 transition-colors duration-300 hover:cursor-pointer" @click="showConnected">
                        <div class="h-3 w-3 rounded-full" :class="isConnected ? 'bg-green-500' : 'bg-red-500'"></div>
                        <span>{{ isConnected ? 'Connected' : 'Disconnected' }}</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="card">
            <RouterView />
        </div>

        <Dialog header="Connect" v-model:visible="isShowConnected" position="top">
            <div>Hubungkan ke platform {{ platformSelected.name }} dengan jaringan {{ networkSelected.Name }} ?</div>
            <template #footer>
                <div class="space-x-2">
                    <Button label="Ya" @click="handleConnect" class="w-24" />
                    <Button severity="help" label="Config" @click="handleCancel" icon="pi pi-cog" class="w-24" />
                </div>
            </template>
        </Dialog>
    </div>
</template>
