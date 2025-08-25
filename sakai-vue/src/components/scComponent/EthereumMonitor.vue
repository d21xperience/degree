<script setup>
import { useSCService } from '@/composables/useSCService';
import { onBeforeUnmount, onMounted, ref } from 'vue';
const scService = useSCService();
const ws = ref(null);
const isConnected = ref(false);
const networkInfo = ref(null);
const latestBlock = ref(null);

// Format timestamp to readable date
const formatTimestamp = (timestamp) => {
    return new Date(timestamp * 1000).toLocaleString();
};

// Initialize WebSocket connection
const initWebSocket = () => {
    ws.value = new WebSocket('ws://localhost:8080/ws');

    ws.value.onopen = () => {
        isConnected.value = true;
        scService.setBCConnected(isConnected.value);
        console.log('Connected to Ethereum monitor');
    };

    ws.value.onmessage = (event) => {
        const data = JSON.parse(event.data);
        console.log('Received:', data);

        switch (data.type) {
            case 'initial_info':
            case 'network_info':
                networkInfo.value = data.data;
                break;
            case 'new_block':
                latestBlock.value = data.data;
                break;
        }
    };

    ws.value.onclose = () => {
        isConnected.value = false;
        scService.setBCConnected(isConnected.value);
        console.log('Connection closed');
    };
};

// Cleanup WebSocket on component unmount
onBeforeUnmount(() => {
    if (ws.value) {
        ws.value.close();
    }
});

// Initialize on component mount
onMounted(() => {
    initWebSocket();
});
</script>

<template>
    <div class="p-4 bg-gray-900 text-white">
        <h1 class="text-2xl font-bold mb-4">Ethereum Network Monitor</h1>

        <!-- Network Info Card -->
        <div v-if="networkInfo" class="bg-gray-800 p-4 rounded-lg mb-4">
            <h2 class="text-xl font-semibold mb-2">Network Status</h2>
            <div class="grid grid-cols-2 gap-2">
                <div>
                    <p class="text-gray-400">Chain ID</p>
                    <p class="font-mono">{{ networkInfo.network_id }}</p>
                </div>
                <div>
                    <p class="text-gray-400">Latest Block</p>
                    <p class="font-mono">{{ networkInfo.latest_block }}</p>
                </div>
                <div>
                    <p class="text-gray-400">Gas Price</p>
                    <p class="font-mono">{{ networkInfo.gas_price }}</p>
                </div>
            </div>
        </div>

        <!-- Connection Status -->
        <div class="flex items-center mb-4">
            <span class="mr-2">Connection:</span>
            <span
                :class="{
                    'text-green-500': isConnected,
                    'text-red-500': !isConnected
                }"
            >
                {{ isConnected ? 'Connected' : 'Disconnected' }}
            </span>
        </div>

        <!-- Latest Block Event -->
        <div v-if="latestBlock" class="bg-gray-800 p-4 rounded-lg">
            <h2 class="text-xl font-semibold mb-2">New Block Mined</h2>
            <p class="font-mono">Block #{{ latestBlock.block_number }}</p>
            <p class="text-gray-400 text-sm">
                {{ formatTimestamp(latestBlock.timestamp) }}
            </p>
        </div>
    </div>
</template>
