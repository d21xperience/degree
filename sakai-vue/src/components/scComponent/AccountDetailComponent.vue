<script setup>
import { ref } from 'vue';
import { useWebSocket } from '@/composables/useWebSocket';

// Account data
const account = ref({
    address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    balance: '1.542389',
    transactionCount: 42,
    tokenCount: 7
});

const transactions = ref([
    {
        hash: '0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b',
        direction: 'in',
        value: '0.25',
        timestamp: 1678901234,
        confirmations: 12
    },
    {
        hash: '0x99df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713955c',
        direction: 'out',
        value: '1.0',
        timestamp: 1678800234,
        confirmations: 24
    }
]);

// WebSocket connection
const { isConnected } = useWebSocket('ws://localhost:8080/ws');

// Utility functions
const shortenAddress = (address) => {
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
};

const shortenHash = (hash) => {
    return `${hash.substring(0, 8)}...`;
};

const formatBalance = (balance) => {
    return parseFloat(balance).toFixed(4);
};

const formatTimestamp = (timestamp) => {
    return new Date(timestamp * 1000).toLocaleString();
};

const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    // Add toast notification in real implementation
};
</script>

<template>
    <div class="max-w-4xl mx-auto p-6 bg-gray-900 text-gray-100 rounded-lg shadow-lg">
        <!-- Header Section -->
        <div class="flex items-center justify-between">
            <div>
                <h1 class="text-2xl font-bold">Blockchain Account Details</h1>
                <!-- <p class="text-gray-400">Ethereum Mainnet</p> -->
            </div>
        </div>

        <!-- Address Card -->
        <div class="bg-gray-800 p-2 rounded-lg mb-2">
            <div class="flex items-center justify-between">
                <div>
                    <h2 class="text-lg font-semibold mb-2">Wallet Address</h2>
                    <div class="flex items-center space-x-2">
                        <span class="font-mono bg-gray-700 px-3 py-1 rounded">
                            {{ shortenAddress(account.address) }}
                        </span>
                        <button @click="copyToClipboard(account.address)" class="text-blue-400 hover:text-blue-300">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3"
                                />
                            </svg>
                        </button>
                    </div>
                </div>
                <div class="text-right">
                    <p class="text-gray-400">QR Code</p>
                    <div class="bg-white p-1 rounded">
                        <!-- QR Code Placeholder -->
                        <div class="h-16 w-16 bg-gray-200 flex items-center justify-center">
                            <span class="text-xs text-gray-600">QR</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Balance and Stats -->
        <div class="grid grid-cols-1 md:grid-cols-1 gap-4 mb-2">
            <div class="bg-gray-800 p-2 rounded-lg">
                <p class="text-gray-400">Balance</p>
                <p class="text-2xl font-bold">{{ formatBalance(account.balance) }} ETH</p>
            </div>
            <div class="bg-gray-800 p-4 rounded-lg">
                <p class="text-gray-400">Transactions</p>
                <p class="text-2xl font-bold">{{ account.transactionCount }}</p>
            </div>
            <div class="bg-gray-800 p-4 rounded-lg">
                <p class="text-gray-400">Token Holdings</p>
                <p class="text-2xl font-bold">{{ account.tokenCount }}</p>
            </div>
        </div>

        <!-- Transaction History -->
        <!-- <div class="bg-gray-800 p-6 rounded-lg">
            <h2 class="text-lg font-semibold mb-4">Recent Transactions</h2>
            <div v-if="transactions.length > 0">
                <div v-for="tx in transactions" :key="tx.hash" class="border-b border-gray-700 py-4 last:border-0">
                    <div class="flex justify-between items-center">
                        <div class="flex items-center space-x-2">
                            <div class="h-8 w-8 rounded-full flex items-center justify-center" :class="tx.direction === 'in' ? 'bg-green-900' : 'bg-red-900'">
                                <span v-if="tx.direction === 'in'">↓</span>
                                <span v-else>↑</span>
                            </div>
                            <div>
                                <p class="font-mono text-sm">{{ shortenHash(tx.hash) }}</p>
                                <p class="text-xs text-gray-400">
                                    {{ formatTimestamp(tx.timestamp) }}
                                </p>
                            </div>
                        </div>
                        <div class="text-right">
                            <p :class="tx.direction === 'in' ? 'text-green-400' : 'text-red-400'">{{ tx.direction === 'in' ? '+' : '-' }}{{ formatBalance(tx.value) }} ETH</p>
                            <p class="text-xs text-gray-400">{{ tx.confirmations }} confirmations</p>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else class="text-center py-8 text-gray-400">
                <p>No transactions found</p>
            </div>
        </div> -->
    </div>
</template>
