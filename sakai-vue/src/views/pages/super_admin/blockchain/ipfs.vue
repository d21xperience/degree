<template>
    <div class="ipfs-admin">
        <h2><i class="pi pi-cog" style="font-size: 1.5rem;"></i> IPFS Admin Configuration</h2>

        <!-- Status IPFS Node -->
        <div class="status-card" :class="nodeStatus.class">
            <div class="status-header">
                <h3>Node Status</h3>
                <span class="status-badge">{{ nodeStatus.text }}</span>
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
            <Button @click="checkNodeStatus" :disabled="isCheckingStatus" icon="pi pi-sync" :class="{ 'pi pi-spin': isCheckingStatus }" label="Refresh Status" />
        </div>

        <!-- Konfigurasi Koneksi -->
        <div class="config-section">
            <h3><i class="pi pi-server" style="font-size: 1.5rem;"></i> Connection Configuration</h3>
            <form @submit.prevent="updateConnectionConfig">
                <div class="form-group">
                    <label>API Endpoint</label>
                    <input type="text" v-model="connectionConfig.apiEndpoint" required />
                </div>
                <div class="form-group">
                    <label>Gateway URL</label>
                    <input type="text" v-model="connectionConfig.gatewayUrl" required />
                </div>
                <div class="form-group">
                    <label>Swarm Ports</label>
                    <InputText v-model="connectionConfig.swarmPorts" placeholder="e.g., 4001,4002,4003" fluid />
                </div>
                <div class="form-actions">
                    <Button :disabled="isSavingConfig" icon="pi pi-save" label="Save Configuration" />
                    <Button type="button" @click="resetConfig" class="secondary" icon="pi pi-undo" label="Reset" severity="secondary" />
                </div>
            </form>
        </div>

        <!-- Manajemen Peer -->
        <div class="config-section">
            <h3><i class="fas fa-network-wired"></i> Peer Management</h3>
            <div class="peer-controls">
                <input type="text" v-model="newPeerAddress" placeholder="/ip4/127.0.0.1/tcp/4001/p2p/QmPeerID" class="peer-input" />
                <button @click="addPeer" :disabled="!newPeerAddress || isManagingPeers"><i class="fas fa-plus"></i> Add Peer</button>
            </div>

            <div class="peer-list">
                <div v-if="peers.length === 0" class="empty-state">No peers configured</div>
                <div v-for="peer in peers" :key="peer.id" class="peer-item">
                    <div class="peer-info">
                        <span class="peer-id">{{ peer.id }}</span>
                        <span class="peer-address">{{ peer.address }}</span>
                    </div>
                    <button @click="removePeer(peer.id)" class="danger small">
                        <i class="fas fa-trash-alt"></i>
                    </button>
                </div>
            </div>
        </div>

        <!-- Manajemen Pin -->
        <div class="config-section">
            <h3><i class="pi pi-thumbtack" style="font-size: 1.5rem;"></i> Pin Management</h3>
            <div class="pin-controls">
                <input type="text" v-model="pinCid" placeholder="Enter CID to pin" class="cid-input" />
                <Button @click="pinContent" :disabled="!pinCid || isManagingPins" icon="pi pi-thumbtack" label="Pin Content" />
                <Button @click="unpinContent" :disabled="!pinCid || isManagingPins" severity="danger" icon="pi pi-trash-alt" label="Unpin" />
            </div>

            <div class="pin-list">
                <div v-if="pinnedItems.length === 0" class="empty-state">No pinned items</div>
                <div v-for="item in pinnedItems" :key="item.cid" class="pin-item">
                    <span class="pin-cid">{{ item.cid }}</span>
                    <span class="pin-size">{{ formatBytes(item.size) }}</span>
                    <span class="pin-type">{{ item.type || 'unknown' }}</span>
                    <button @click="showPinInfo(item.cid)" class="small"><i class="fas fa-info-circle"></i> Details</button>
                </div>
            </div>
        </div>

        <!-- Advanced Configuration -->
        <div class="config-section advanced">
            <h3 @click="showAdvanced = !showAdvanced">
                <i class="pi" :class="showAdvanced ? 'pi-chevron-down' : 'pi-chevron-right'"></i>
                Advanced Configuration
            </h3>
            <div v-if="showAdvanced" class="advanced-content">
                <div class="form-group">
                    <label>GC Interval (minutes)</label>
                    <input type="number" v-model="advancedConfig.gcInterval" />
                </div>
                <div class="form-group">
                    <label>Routing Type</label>
                    <select v-model="advancedConfig.routingType">
                        <option value="dht">DHT (distributed)</option>
                        <option value="dhtclient">DHT Client</option>
                        <option value="none">None</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Enable Auto-NAT</label>
                    <input type="checkbox" v-model="advancedConfig.autoNAT" />
                </div>
                <div class="form-group">
                    <label>Enable PubSub</label>
                    <input type="checkbox" v-model="advancedConfig.pubSub" />
                </div>
                <button @click="saveAdvancedConfig" :disabled="isSavingAdvanced"><i class="fas fa-save"></i> Save Advanced Config</button>
            </div>
        </div>

        <!-- Modal untuk detail pin -->
        <div v-if="currentPinDetail" class="modal-backdrop" @click.self="currentPinDetail = null">
            <div class="modal">
                <div class="modal-header">
                    <h3>Pin Details</h3>
                    <button @click="currentPinDetail = null" class="close-button">
                        <i class="fas fa-times"></i>
                    </button>
                </div>
                <div class="modal-content">
                    <div class="detail-grid">
                        <div class="detail-row">
                            <label>CID:</label>
                            <span class="monospace">{{ currentPinDetail.cid }}</span>
                        </div>
                        <div class="detail-row">
                            <label>Size:</label>
                            <span>{{ formatBytes(currentPinDetail.size) }}</span>
                        </div>
                        <div class="detail-row">
                            <label>Type:</label>
                            <span>{{ currentPinDetail.type || 'unknown' }}</span>
                        </div>
                        <div class="detail-row">
                            <label>Pinned At:</label>
                            <span>{{ formatDate(currentPinDetail.created) }}</span>
                        </div>
                        <div class="detail-row">
                            <label>References:</label>
                            <span>{{ currentPinDetail.refs || 'none' }}</span>
                        </div>
                    </div>
                </div>
                <div class="modal-actions">
                    <button @click="unpinContent(currentPinDetail.cid)" class="danger"><i class="fas fa-thumbtack"></i> Unpin This Content</button>
                    <button @click="currentPinDetail = null">Close</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { formatBytes, formatDate } from '@/utils/format'; // Anda perlu membuat utility functions ini
import { computed, onMounted, ref } from 'vue';

// API endpoints - sesuaikan dengan backend Golang Anda
const API_BASE = '/api/admin/ipfs';
const STATUS_ENDPOINT = `${API_BASE}/status`;
const CONFIG_ENDPOINT = `${API_BASE}/config`;
const PEERS_ENDPOINT = `${API_BASE}/peers`;
const PINS_ENDPOINT = `${API_BASE}/pins`;

// State untuk status node
const nodeInfo = ref(null);
const isCheckingStatus = ref(false);
const nodeStatus = computed(() => {
    if (!nodeInfo.value) return { text: 'Disconnected', class: 'disconnected' };
    if (nodeInfo.value.peers > 0) return { text: 'Connected', class: 'connected' };
    return { text: 'Standalone', class: 'standalone' };
});

// State untuk konfigurasi koneksi
const connectionConfig = ref({
    apiEndpoint: '',
    gatewayUrl: '',
    swarmPorts: ''
});
const isSavingConfig = ref(false);
const originalConfig = ref(null);

// State untuk manajemen peer
const peers = ref([]);
const newPeerAddress = ref('');
const isManagingPeers = ref(false);

// State untuk manajemen pin
const pinnedItems = ref([]);
const pinCid = ref('');
const isManagingPins = ref(false);
const currentPinDetail = ref(null);

// State untuk advanced config
const advancedConfig = ref({
    gcInterval: 60,
    routingType: 'dht',
    autoNAT: true,
    pubSub: false
});
const isSavingAdvanced = ref(false);
const showAdvanced = ref(false);

// Load data saat komponen dimount
onMounted(async () => {
    await Promise.all([checkNodeStatus(), fetchConfig(), fetchPeers(), fetchPinnedItems()]);
});

// Fungsi untuk memeriksa status node
const checkNodeStatus = async () => {
    isCheckingStatus.value = true;
    try {
        const response = await fetch(STATUS_ENDPOINT);
        if (!response.ok) throw new Error('Failed to fetch node status');
        nodeInfo.value = await response.json();
    } catch (error) {
        console.error('Error checking node status:', error);
        nodeInfo.value = null;
    } finally {
        isCheckingStatus.value = false;
    }
};

// Fungsi untuk mengambil konfigurasi
const fetchConfig = async () => {
    try {
        const response = await fetch(CONFIG_ENDPOINT);
        if (!response.ok) throw new Error('Failed to fetch config');
        const config = await response.json();
        connectionConfig.value = config;
        originalConfig.value = JSON.parse(JSON.stringify(config));
        advancedConfig.value = config.advanced || advancedConfig.value;
    } catch (error) {
        console.error('Error fetching config:', error);
    }
};

// Fungsi untuk menyimpan konfigurasi koneksi
const updateConnectionConfig = async () => {
    isSavingConfig.value = true;
    try {
        const response = await fetch(CONFIG_ENDPOINT, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(connectionConfig.value)
        });

        if (!response.ok) throw new Error('Failed to update config');

        const updatedConfig = await response.json();
        connectionConfig.value = updatedConfig;
        originalConfig.value = JSON.parse(JSON.stringify(updatedConfig));
    } catch (error) {
        console.error('Error updating config:', error);
    } finally {
        isSavingConfig.value = false;
    }
};

// Fungsi untuk reset konfigurasi
const resetConfig = () => {
    if (originalConfig.value) {
        connectionConfig.value = JSON.parse(JSON.stringify(originalConfig.value));
    }
};

// Fungsi untuk mengambil daftar peer
const fetchPeers = async () => {
    try {
        const response = await fetch(PEERS_ENDPOINT);
        if (!response.ok) throw new Error('Failed to fetch peers');
        peers.value = await response.json();
    } catch (error) {
        console.error('Error fetching peers:', error);
        peers.value = [];
    }
};

// Fungsi untuk menambahkan peer
const addPeer = async () => {
    isManagingPeers.value = true;
    try {
        const response = await fetch(PEERS_ENDPOINT, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ address: newPeerAddress.value })
        });

        if (!response.ok) throw new Error('Failed to add peer');

        await fetchPeers();
        newPeerAddress.value = '';
    } catch (error) {
        console.error('Error adding peer:', error);
    } finally {
        isManagingPeers.value = false;
    }
};

// Fungsi untuk menghapus peer
const removePeer = async (peerId) => {
    isManagingPeers.value = true;
    try {
        const response = await fetch(`${PEERS_ENDPOINT}/${encodeURIComponent(peerId)}`, {
            method: 'DELETE'
        });

        if (!response.ok) throw new Error('Failed to remove peer');

        await fetchPeers();
    } catch (error) {
        console.error('Error removing peer:', error);
    } finally {
        isManagingPeers.value = false;
    }
};

// Fungsi untuk mengambil daftar pinned items
const fetchPinnedItems = async () => {
    try {
        const response = await fetch(PINS_ENDPOINT);
        if (!response.ok) throw new Error('Failed to fetch pinned items');
        pinnedItems.value = await response.json();
    } catch (error) {
        console.error('Error fetching pinned items:', error);
        pinnedItems.value = [];
    }
};

// Fungsi untuk pin content
const pinContent = async (cid = pinCid.value) => {
    isManagingPins.value = true;
    try {
        const response = await fetch(PINS_ENDPOINT, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ cid })
        });

        if (!response.ok) throw new Error('Failed to pin content');

        await fetchPinnedItems();
        pinCid.value = '';
    } catch (error) {
        console.error('Error pinning content:', error);
    } finally {
        isManagingPins.value = false;
    }
};

// Fungsi untuk unpin content
const unpinContent = async (cid = pinCid.value) => {
    isManagingPins.value = true;
    try {
        const response = await fetch(`${PINS_ENDPOINT}/${encodeURIComponent(cid)}`, {
            method: 'DELETE'
        });

        if (!response.ok) throw new Error('Failed to unpin content');

        await fetchPinnedItems();
        pinCid.value = '';
        if (currentPinDetail.value?.cid === cid) {
            currentPinDetail.value = null;
        }
    } catch (error) {
        console.error('Error unpinning content:', error);
    } finally {
        isManagingPins.value = false;
    }
};

// Fungsi untuk menampilkan detail pin
const showPinInfo = (cid) => {
    const pin = pinnedItems.value.find((p) => p.cid === cid);
    if (pin) {
        currentPinDetail.value = pin;
    }
};

// Fungsi untuk menyimpan advanced config
const saveAdvancedConfig = async () => {
    isSavingAdvanced.value = true;
    try {
        const response = await fetch(`${CONFIG_ENDPOINT}/advanced`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(advancedConfig.value)
        });

        if (!response.ok) throw new Error('Failed to save advanced config');

        const result = await response.json();
        advancedConfig.value = result;
    } catch (error) {
        console.error('Error saving advanced config:', error);
    } finally {
        isSavingAdvanced.value = false;
    }
};
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
