// composables/useWebSocket.js
import { onUnmounted, ref } from 'vue';

export function useWebSocket(url) {
    const ws = ref(null);
    const isConnected = ref(false);
    const networkInfo = ref(null);
    const latestBlock = ref(null);

    const initWebSocket = () => {
        ws.value = new WebSocket(url);

        ws.value.onopen = () => {
            isConnected.value = true;
            console.log('Connected to Ethereum monitor');
        };

        ws.value.onmessage = (event) => {
            const data = JSON.parse(event.data);

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
            console.log('Connection closed');
        };
    };

    // Auto-close on unmount
    onUnmounted(() => {
        if (ws.value) {
            ws.value.close();
        }
    });

    // Initialize WebSocket
    initWebSocket();

    return {
        isConnected,
        networkInfo,
        latestBlock
    };
}
