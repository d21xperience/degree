<script setup>
import { useAuthService } from '@/composables/useAuthService';
import { useSCService } from '@/composables/useSCService';
import { computed, onMounted, reactive, ref } from 'vue';
const authService = useAuthService();

const scService = useSCService();
const state = reactive({
    password: '',
    confirmPassword: '',
    isGenerating: false,
    showWallet: false,
    walletData: null,
    error: '',
    wallets: [],
    loadingWallets: false
});
const currentUser = computed(() => authService.currentUser.value);
const form = ref(null);

const validateForm = () => {
    if (!state.password || state.password.length < 8) {
        state.error = 'Password must be at least 8 characters';
        return false;
    }

    if (state.password !== state.confirmPassword) {
        state.error = 'Passwords do not match';
        return false;
    }

    state.error = '';
    return true;
};

const generateWallet = async () => {
    if (!validateForm()) return;

    state.isGenerating = true;
    state.showWallet = false;
    state.error = '';

    try {
        const response = await fetch('http://localhost:8080/api/wallet/generate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                password: state.password
            })
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to generate wallet');
        }

        const data = await response.json();
        state.walletData = data;
        state.showWallet = true;

        // Refresh wallet list
        await loadWallets();
    } catch (error) {
        state.error = 'Failed to generate wallet: ' + error.message;
    } finally {
        state.isGenerating = false;
    }
};

const loadWallets = async () => {
    state.loadingWallets = true;
    try {
        // const response = await fetch('http://localhost:8184/api/v1/scs/blockchainaccount/list');
        // if (response.ok) {
        //     state.wallets = await response.json();
        //     console.log(state.wallets);
        // }
        const response = await scService.fetchBCAccount(currentUser.value.username);
        console.log(response);
        // if (response) {
        state.wallets = response;
        // }
    } catch (error) {
        console.error('Failed to load wallets:', error);
    } finally {
        state.loadingWallets = false;
    }
};

const downloadKeystore = (walletId, filename) => {
    window.open(`http://localhost:8080/api/wallet/${walletId}/download`, '_blank');
};

const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text).then(() => {
        alert('Copied to clipboard!');
    });
};

const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString();
};

onMounted(() => {
    loadWallets();
});
const pvKey = ref();
const isAddDialog = ref(false);
const isImportPvK = ref(false);
const importDialog = () => {
    isAddDialog.value = false;
    isImportPvK.value = true;
};
const importPrivateKey = async () => {
    const payload = {
        private_key: pvKey.value,
        username: currentUser.value.username
    };
    const response = await scService.importBCAccount(payload);
    console.log(response);
    // isAddDialog.value = false;
    // isImportPvK.value = true;
};
</script>

<template>
    <div class="max-full mx-auto">
        <!-- Generated Wallet Display -->
        <div v-if="state.showWallet && state.walletData" class="bg-white rounded-xl shadow-lg p-6 mb-8">
            <h3 class="text-xl font-semibold text-gray-800 mb-4">Your New Wallet</h3>

            <div class="space-y-4">
                <div class="p-4 bg-green-50 rounded-lg border border-green-200">
                    <div class="flex justify-between items-start mb-2">
                        <h4 class="font-medium text-green-800">Wallet Address</h4>
                        <button @click="copyToClipboard(state.walletData.address)" class="text-green-600 hover:text-green-800 text-sm">Copy</button>
                    </div>
                    <p class="font-mono text-green-700 break-all">{{ state.walletData.address }}</p>
                </div>

                <div class="p-4 bg-blue-50 rounded-lg border border-blue-200">
                    <h4 class="font-medium text-blue-800 mb-2">Keystore File</h4>
                    <p class="text-blue-700 mb-3">Filename: {{ state.walletData.filename }}</p>
                    <p class="text-blue-700 mb-3">Created: {{ formatDate(state.walletData.createdAt) }}</p>
                    <button @click="downloadKeystore(state.walletData.id, state.walletData.filename)" class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition">Download Keystore</button>
                </div>
            </div>

            <div class="mt-6 p-4 bg-red-50 border border-red-200 rounded-lg">
                <h4 class="font-bold text-red-800 mb-2">⚠️ Important Security Notice</h4>
                <ul class="text-red-700 space-y-1 text-sm">
                    <li>• Store your keystore file and password in separate secure locations</li>
                    <li>• Never share your private key or keystore file with anyone</li>
                    <li>• Make multiple backups of your keystore file</li>
                    <li>• Test your backup before storing large amounts of cryptocurrency</li>
                </ul>
            </div>
        </div>

        <!-- Wallet List -->
        <div class="bg-white rounded-xl">
            <div class="flex justify-between items-center mb-6">
                <div class="flex space-x-2">
                    <h2 class="text-2xl font-semibold text-gray-800">Generated Wallets</h2>
                    <div>
                        <Button icon="pi pi-plus" size="small" severity="secondary" @click="isAddDialog = true" />
                    </div>
                </div>
                <button @click="loadWallets" :disabled="state.loadingWallets" class="px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-lg text-sm font-medium">
                    {{ state.loadingWallets ? 'Loading...' : 'Refresh' }}
                </button>
            </div>

            <div v-if="state.loadingWallets" class="text-center py-8">
                <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            </div>

            <div v-else-if="state.wallets.length === 0" class="text-center py-8 text-gray-500">No wallets generated yet</div>

            <div v-else class="space-y-4">
                <div v-for="wallet in state.wallets" :key="wallet.id" class="border border-gray-200 rounded-lg p-4 hover:bg-gray-50 transition">
                    <div class="flex justify-between items-start">
                        <div class="flex-1">
                            <div class="font-mono text-sm text-gray-600 mb-1">{{ wallet?.address }}</div>
                            <div class="text-xs text-gray-500">Created: {{ formatDate(wallet.createdAt) }} | Filename: {{ wallet.filename }}</div>
                        </div>
                        <button @click="downloadKeystore(wallet.id, wallet.filename)" class="ml-4 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium py-2 px-3 rounded-lg transition">Download</button>
                    </div>
                </div>
            </div>
        </div>

        <Dialog v-model:visible="isAddDialog" header="New Wallet" position="top" :modal="true" style="width: 24rem">
            <form ref="form" @submit.prevent="generateWallet" class="space-y-6">
                <div>
                    <label for="password" class="block text-sm font-medium text-gray-700 mb-2"> Password </label>
                    <input
                        id="password"
                        v-model="state.password"
                        type="password"
                        required
                        minlength="8"
                        class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                        placeholder="Enter password for keystore encryption"
                    />
                    <p class="mt-1 text-sm text-gray-500">Minimum 8 characters</p>
                </div>

                <div>
                    <label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-2"> Confirm Password </label>
                    <input
                        id="confirmPassword"
                        v-model="state.confirmPassword"
                        type="password"
                        required
                        class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                        placeholder="Confirm your password"
                    />
                </div>

                <div v-if="state.error" class="p-3 bg-red-50 text-red-700 rounded-lg">
                    {{ state.error }}
                </div>

                <button type="submit" :disabled="state.isGenerating" class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 text-white font-semibold py-3 px-6 rounded-lg transition duration-200 flex items-center justify-center">
                    <span v-if="state.isGenerating" class="mr-2">
                        <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                    </span>
                    {{ state.isGenerating ? 'Generating...' : 'Generate Wallet' }}
                </button>
                <div class="flex justify-center text-sm">Or</div>
                <Button label="Import Private Key" fluid @click="importDialog" severity="secondary" />
            </form>
        </Dialog>

        <!-- Privatekey -->
        <Dialog header="Import Private Key" v-model:visible="isImportPvK" style="width: 24rem" position="top">
            <InputText placeholder="masukan private key" fluid v-model="pvKey" />
            <div class="mt-2">
                <Button label="Import" icon="pi pi-upload" fluid @click="importPrivateKey" />
            </div>
            <template #footer>
                <div class="mt-6 p-4 bg-red-50 border border-red-200 rounded-lg">
                    <h4 class="font-bold text-red-800 mb-2">⚠️ Important Security Notice</h4>
                    <ul class="text-red-700 space-y-1 text-sm list-disc px-2">
                        <!-- <li>• Store your keystore file and password in separate secure locations</li> -->
                        <li>Never share your private key or keystore file with anyone</li>
                        <li>Make multiple backups of your keystore file</li>
                        <li>Test your backup before storing large amounts of cryptocurrency</li>
                    </ul>
                </div>
            </template>
        </Dialog>
    </div>
</template>
