<template>
    <div class="ipfs-admin max-w-6xl mx-auto p-6">
        <!-- Judul Utama -->
        <h2 class="flex items-center gap-3 text-2xl font-semibold border-b border-gray-200 pb-3 mb-6">
            <i class="pi pi-cog text-2xl"></i>
            Blockchain Admin Configuration
        </h2>

        <!-- Konfigurasi Koneksi -->
        <div class="config-section bg-white rounded-lg p-6 mb-6 shadow-sm border border-gray-100">
            <EthereumMonitor />

            <!-- Subjudul -->
            <h3 class="flex items-center gap-2 text-lg font-medium text-gray-800 mb-4">
                <i class="pi pi-server text-xl"></i>
                Connection Configuration
            </h3>

            <!-- Platform & Environment + Save/Edit Button -->
            <div class="flex justify-between items-center space-x-2 mb-4">
                <div class="flex w-full space-x-1">
                    <PlatformComponent v-model="platformSelected" :disabled="isNetworkPlatform" class="flex-1" />
                    <EnvironmentComponent v-model="environmentSelected" :architecture="platformSelected" :disabled="isNetworkPlatform" class="flex-1" />
                </div>
                <div>
                    <Button v-if="!isNetworkPlatform" icon="pi pi-save" @click="saveConfiguration" severity="danger" class="p-button-danger" />
                    <Button v-else icon="pi pi-pencil" @click="editConfiguration" severity="warn" class="p-button-warning" />
                </div>
            </div>

            <!-- Tombol Sync -->
            <div class="flex justify-start mt-2 space-x-2">
                <Button icon="pi pi-sync" @click="checknetworkStatus" :label="networkStatus.text" :severity="networkStatus.severity" :disabled="!isNetworkPlatform" class="p-button-sm" />
            </div>
        </div>
    </div>
</template>

<script setup>
import EnvironmentComponent from '@/components/scComponent/EnvironmentComponent.vue';
import EthereumMonitor from '@/components/scComponent/EthereumMonitor.vue';
import PlatformComponent from '@/components/scComponent/PlatformComponent.vue';
import { useSCService } from '@/composables/useSCService';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref, watchEffect } from 'vue';

const scService = useSCService();
const toast = useToast();

// State untuk status node
const nodeInfo = ref(null);
const isCheckingStatus = ref(false);
const networkStatus = computed(() => {
    if (nodeInfo.value) return { text: 'Disconnect', class: 'disconnected', severity: 'help' };
    if (!nodeInfo.value) return { text: 'Connect', class: 'connected', severity: 'warn' };
    // if (nodeInfo.value.peers > 0) return { text: 'Connected', class: 'connected' };
    return { text: 'Standalone', class: 'standalone' };
});

// Fungsi untuk memeriksa status node
const checknetworkStatus = async () => {
    isCheckingStatus.value = true;
    try {
        const response = await scService.setBCConfig();
        if (!response.status) throw new Error('Failed to fetch node status');
        // console.log(response);
        toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
        nodeInfo.value = response;
    } catch (error) {
        console.error('Error checking node status:', error);
        toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal terhubung ke jaringan: ${error}`, life: 3000 });

        nodeInfo.value = null;
    } finally {
        isCheckingStatus.value = false;
    }
};

const isNetworkPlatform = ref(false);

const environmentSelected = ref();

const platformSelected = ref(null);

// watchEffect(async () => {

//     //     environmentSelected.value = await scService.getBCNetwork();
//     //     // console.log(environmentSelected.value);
//     if (platformSelected.value) {
//             platformSelected.value = await scService.getNetowrkPlatform();
//             // console.log(platformSelected.value);
//             // isNetworkPlatform.value = true;
//         }
// });

const saveConfiguration = () => {
    scService.setNetwrokPlatform(platformSelected.value);
    scService.setBCNetwork(environmentSelected.value);
    isNetworkPlatform.value = true;
};

const editConfiguration = () => {
    isNetworkPlatform.value = false;
};

const isConnected = ref(true);
// Load data saat komponen dimount
onMounted(async () => {
    platformSelected.value = {
    "id": "f45865b2-9dd9-4085-942c-89a8d1847674",
    "name": "Ethereum",
    "active": false
} 
    // environmentSelected.value = await scService.getBCNetwork();
    
    // console.log(environmentSelected.value)
    // await Promise.all([checknetworkStatus()]);
    // await Promise.all([checknetworkStatus(), fetchConfig(), fetchPeers(), fetchPinnedItems()]);
});
</script>

<style scoped></style>
