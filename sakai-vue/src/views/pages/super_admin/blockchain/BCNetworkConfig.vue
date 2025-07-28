<template>
    <div class="ipfs-admin">
        <h2><i class="pi pi-cog" style="font-size: 1.5rem"></i> Blockhain Admin Configuration</h2>

        <!-- Status IPFS Node -->
        <!-- <div class="status-card" :class="networkStatus.class">
            <div class="status-header">
                <h3>Network Status</h3>
                <span class="status-badge">{{ networkStatus.text }}</span>
            </div>
            <div v-if="nodeInfo" class="status-details">
                <div class="detail-item">
                    <label>ID:</label>
                    <span class="monospace">{{ nodeInfo.id }}</span>
                </div>
                <div class="detail-item">
                    <label>Version:</label>
                    <span>{{ nodeInfo.version }}</span>
                </div>
                <div class="detail-item">
                    <label>Peers:</label>
                    <span>{{ nodeInfo.peers }}</span>
                </div>
                <div class="detail-item">
                    <label>Repository:</label>
                    <span>{{ formatBytes(nodeInfo.repoSize) }} / {{ formatBytes(nodeInfo.repoSizeMax) }}</span>
                </div>
            </div>
            <Button @click="checknetworkStatus" :disabled="isCheckingStatus" icon="pi pi-sync" :class="{ 'pi pi-spin': isCheckingStatus }" label="Refresh Status" />
        </div> -->

        <EthereumMonitor />
        <!-- Konfigurasi Koneksi -->
        <div class="config-section">
            <h3><i class="pi pi-server" style="font-size: 1.5rem"></i> Connection Configuration</h3>
            <div class="flex justify-between space-x-2">
                <div class="flex w-full space-x-1">
                    <PlatformComponent v-model="platformSelected" :disabled="isNetworkPlatform" />
                    <EnvironmentComponent v-model="environmentSelected" :architecture="platformSelected" :disabled="isNetworkPlatform" />
                </div>
                <div>
                    <Button icon="pi pi-save" class="" @click="saveConfiguration" v-if="!isNetworkPlatform" severity="danger"/>
                    <Button icon="pi pi-pencil" class="" @click="editConfiguration" severity="warn" v-else />

                </div>
            </div>
            <div class="flex justify-start mt-2 space-x-2">
                <Button label="Connect" icon="pi pi-sync" class="w-32" @click="checknetworkStatus" />
            </div>
        </div>
    </div>
</template>

<script setup>
import EnvironmentComponent from '@/components/scComponent/EnvironmentComponent.vue';
import EthereumMonitor from '@/components/scComponent/EthereumMonitor.vue';
import PlatformComponent from '@/components/scComponent/PlatformComponent.vue';
import { useSCService } from '@/composables/useSCService';
import { computed, onMounted, ref, watchEffect } from 'vue';
const scService = useSCService();

// API endpoints - sesuaikan dengan backend Golang Anda
const API_BASE = '/api/admin/ipfs';
const STATUS_ENDPOINT = `${API_BASE}/status`;

// State untuk status node
const nodeInfo = ref(null);
const isCheckingStatus = ref(false);
const networkStatus = computed(() => {
    if (!nodeInfo.value) return { text: 'Disconnected', class: 'disconnected' };
    if (nodeInfo.value) return { text: 'Connected', class: 'connected' };
    // if (nodeInfo.value.peers > 0) return { text: 'Connected', class: 'connected' };
    return { text: 'Standalone', class: 'standalone' };
});

// Fungsi untuk memeriksa status node
const checknetworkStatus = async () => {
    isCheckingStatus.value = true;
    try {
        const response = await scService.setBCConfig();
        // if (!response.status) throw new Error('Failed to fetch node status');
        console.log(response);
        nodeInfo.value = response;
    } catch (error) {
        console.error('Error checking node status:', error);
        nodeInfo.value = null;
    } finally {
        isCheckingStatus.value = false;
    }
};

const isNetworkPlatform = ref(false);

const environmentSelected = ref();

const platformSelected = ref(null);
watchEffect(async () => {
    platformSelected.value = await scService.getNetowrkPlatform();
    environmentSelected.value = await scService.getBCNetwork();
    // console.log(environmentSelected.value);
    if (platformSelected.value) {
        // console.log(platformSelected.value);
        isNetworkPlatform.value = true;
    }
});

const saveConfiguration = () => {
    scService.setNetwrokPlatform(platformSelected.value);
    scService.setBCNetwork(environmentSelected.value);
    isNetworkPlatform.value = true;
};

const editConfiguration = () => {
    isNetworkPlatform.value = false;
};

// Load data saat komponen dimount
onMounted(async () => {
    // await Promise.all([checknetworkStatus()]);
    // await Promise.all([checknetworkStatus(), fetchConfig(), fetchPeers(), fetchPinnedItems()]);
});
</script>

<style scoped>
.ipfs-admin {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
    /* font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; */
}

h2,
h3 {
    /* color: #2c3e50; */
    margin-bottom: 20px;
}

h2 {
    border-bottom: 2px solid #eee;
    padding-bottom: 10px;
    display: flex;
    align-items: center;
    gap: 10px;
}

h3 {
    font-size: 1.2em;
    display: flex;
    align-items: center;
    gap: 8px;
}

.config-section {
    /* background: white; */
    border-radius: 8px;
    padding: 20px;
    margin-bottom: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.advanced h3 {
    cursor: pointer;
    user-select: none;
}

.advanced-content {
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid #eee;
}

.status-card {
    /* background: white; */
    border-radius: 8px;
    padding: 20px;
    margin-bottom: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
    border-left: 5px solid #ccc;
}

.status-card.connected {
    border-left-color: #2ecc71;
}

.status-card.standalone {
    border-left-color: #f39c12;
}

.status-card.disconnected {
    border-left-color: #e74c3c;
}

.status-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
}

.status-badge {
    padding: 5px 10px;
    border-radius: 20px;
    font-size: 0.8em;
    font-weight: bold;
    text-transform: uppercase;
}

.status-card.connected .status-badge {
    background: #2ecc71;
    color: white;
}

.status-card.standalone .status-badge {
    background: #f39c12;
    color: white;
}

.status-card.disconnected .status-badge {
    background: #e74c3c;
    color: white;
}

.status-details {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 15px;
    margin-bottom: 15px;
}

.detail-item {
    display: flex;
    flex-direction: column;
}

.detail-item label {
    font-size: 0.9em;
    color: #7f8c8d;
    margin-bottom: 3px;
}

.form-group {
    margin-bottom: 15px;
}

.form-group label {
    display: block;
    margin-bottom: 5px;
    font-weight: 500;
}

.form-group input,
.form-group select {
    width: 100%;
    padding: 8px 12px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1em;
}

.form-actions {
    display: flex;
    gap: 10px;
    margin-top: 20px;
}

.peer-controls,
.pin-controls {
    display: flex;
    gap: 10px;
    margin-bottom: 15px;
}

.peer-input,
.cid-input {
    flex: 1;
    padding: 8px 12px;
    border: 1px solid #ddd;
    border-radius: 4px;
}

.peer-list,
.pin-list {
    border: 1px solid #eee;
    border-radius: 4px;
    max-height: 300px;
    overflow-y: auto;
}

.empty-state {
    padding: 20px;
    text-align: center;
    color: #95a5a6;
}

.peer-item,
.pin-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 15px;
    border-bottom: 1px solid #eee;
}

.peer-item:last-child,
.pin-item:last-child {
    border-bottom: none;
}

.peer-info {
    flex: 1;
    display: flex;
    flex-direction: column;
}

.peer-id {
    font-weight: 500;
    font-family: monospace;
    font-size: 0.9em;
}

.peer-address {
    font-size: 0.8em;
    color: #7f8c8d;
    word-break: break-all;
}

.pin-cid {
    font-family: monospace;
    flex: 2;
}

.pin-size {
    flex: 1;
    text-align: right;
    padding-right: 15px;
    color: #7f8c8d;
}

.pin-type {
    flex: 1;
    text-align: center;
    color: #7f8c8d;
    text-transform: capitalize;
}

/* button {
    padding: 8px 15px;
    border: none;
    border-radius: 4px;
    background: #3498db;
    color: white;
    cursor: pointer;
    font-size: 0.9em;
    display: inline-flex;
    align-items: center;
    gap: 5px;
    transition: background 0.2s;
} */

button:hover {
    background: #2980b9;
}

button.secondary {
    background: #95a5a6;
}

button.secondary:hover {
    background: #7f8c8d;
}

button.danger {
    background: #e74c3c;
}

button.danger:hover {
    background: #c0392b;
}

button.small {
    padding: 5px 10px;
    font-size: 0.8em;
}

button:disabled {
    background: #bdc3c7;
    cursor: not-allowed;
}

.monospace {
    font-family: monospace;
    word-break: break-all;
}

/* Modal styles */
.modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal {
    background: white;
    border-radius: 8px;
    width: 90%;
    max-width: 600px;
    max-height: 90vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}

.modal-header {
    padding: 15px 20px;
    border-bottom: 1px solid #eee;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.modal-content {
    padding: 20px;
    overflow-y: auto;
}

.modal-actions {
    padding: 15px 20px;
    border-top: 1px solid #eee;
    display: flex;
    justify-content: flex-end;
    gap: 10px;
}

.close-button {
    background: none;
    border: none;
    color: #7f8c8d;
    font-size: 1.2em;
    padding: 5px;
}

.detail-grid {
    display: grid;
    grid-template-columns: max-content 1fr;
    gap: 10px 15px;
}

.detail-row {
    display: contents;
}

.detail-row label {
    font-weight: 500;
    color: #7f8c8d;
}

/* Font Awesome icons */
.fa-spin {
    animation: fa-spin 2s infinite linear;
}

@keyframes fa-spin {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(359deg);
    }
}
</style>
