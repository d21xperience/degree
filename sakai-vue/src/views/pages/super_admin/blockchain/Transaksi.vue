<script setup>
import TransactionModal from '@/components/TransactionModal.vue';
import { debounce } from 'lodash-es';
import { onMounted, ref } from 'vue';
// State
const transactions = ref([]);
const isLoading = ref(false);
const error = ref(null);
const selectedBlockchain = ref('');
const searchQuery = ref('');
const currentPage = ref(1);
const totalPages = ref(1);
const selectedTransaction = ref(null);

// Konstanta
const ITEMS_PER_PAGE = 10;
const API_BASE_URL = 'http://localhost:8080/api/transactions'; // Sesuaikan dengan URL backend Golang Anda

// Format nama blockchain untuk tampilan
const formatBlockchainName = (blockchain) => {
    const names = {
        ethereum: 'Ethereum',
        quorum: 'Quorum',
        fabric: 'Hyperledger Fabric'
    };
    return names[blockchain.toLowerCase()] || blockchain;
};

// Format nilai transaksi berdasarkan blockchain
const formatValue = (value, blockchain) => {
    if (blockchain.toLowerCase() === 'fabric') {
        return value;
    }
    // Untuk Ethereum dan Quorum, konversi dari wei ke ether
    return `${(parseInt(value) / 1e18).toFixed(4)} ETH`;
};

// Format timestamp
const formatTimestamp = (timestamp) => {
    return new Date(timestamp).toLocaleString();
};

// Fungsi untuk mengambil data transaksi dari backend Golang
const fetchTransactions = async () => {
    try {
        isLoading.value = true;
        error.value = null;

        const params = new URLSearchParams();
        if (selectedBlockchain.value) params.append('blockchain', selectedBlockchain.value);
        if (searchQuery.value) params.append('search', searchQuery.value);
        params.append('page', currentPage.value);
        params.append('limit', ITEMS_PER_PAGE);

        const response = await fetch(`${API_BASE_URL}?${params.toString()}`);

        if (!response.ok) {
            throw new Error(`Gagal memuat data: ${response.statusText}`);
        }

        const data = await response.json();
        transactions.value = data.transactions || [];
        totalPages.value = Math.ceil(data.total / ITEMS_PER_PAGE);
    } catch (err) {
        console.error('Error fetching transactions:', err);
        error.value = err.message || 'Terjadi kesalahan saat memuat data transaksi';
        transactions.value = [];
    } finally {
        isLoading.value = false;
    }
};

// Debounce pencarian
const debouncedSearch = debounce(() => {
    currentPage.value = 1;
    fetchTransactions();
}, 500);

// Ganti halaman
const changePage = (newPage) => {
    if (newPage >= 1 && newPage <= totalPages.value) {
        currentPage.value = newPage;
        fetchTransactions();
    }
};

// Tampilkan detail transaksi
const showTransactionDetails = (tx) => {
    selectedTransaction.value = tx;
};

// Lifecycle hook
onMounted(() => {
    fetchTransactions();
});
</script>

<template>
    <div class="blockchain-transactions">
        <h1>Daftar Transaksi Blockchain</h1>

        <div class="filter-controls">
            <select v-model="selectedBlockchain" @change="fetchTransactions">
                <option value="">Semua Blockchain</option>
                <option value="ethereum">Ethereum</option>
                <option value="quorum">Quorum</option>
                <option value="fabric">Hyperledger Fabric</option>
            </select>

            <input v-model="searchQuery" type="text" placeholder="Cari transaksi..." @input="debouncedSearch" />
        </div>

        <div v-if="isLoading" class="loading">Memuat data transaksi...</div>

        <div v-else-if="error" class="error">
            {{ error }}
        </div>

        <div v-else-if="transactions.length === 0" class="empty">Tidak ada transaksi ditemukan</div>

        <div v-else class="transactions-list">
            <div v-for="tx in transactions" :key="tx.id" class="transaction-card">
                <div class="tx-header">
                    <span class="tx-id">{{ tx.id }}</span>
                    <span class="tx-chain" :class="tx.blockchain.toLowerCase()">
                        {{ formatBlockchainName(tx.blockchain) }}
                    </span>
                </div>

                <div class="tx-details">
                    <div class="tx-row">
                        <span class="label">Dari:</span>
                        <span class="value">{{ tx.from }}</span>
                    </div>
                    <div class="tx-row">
                        <span class="label">Ke:</span>
                        <span class="value">{{ tx.to }}</span>
                    </div>
                    <div class="tx-row">
                        <span class="label">Nilai:</span>
                        <span class="value">{{ formatValue(tx.value, tx.blockchain) }}</span>
                    </div>
                    <div class="tx-row">
                        <span class="label">Waktu:</span>
                        <span class="value">{{ formatTimestamp(tx.timestamp) }}</span>
                    </div>
                    <div v-if="tx.blockNumber" class="tx-row">
                        <span class="label">Block:</span>
                        <span class="value">{{ tx.blockNumber }}</span>
                    </div>
                </div>

                <div class="tx-footer">
                    <button class="details-btn" @click="showTransactionDetails(tx)">Detail</button>
                </div>
            </div>
        </div>

        <div v-if="totalPages > 1" class="pagination">
            <button :disabled="currentPage === 1" @click="changePage(currentPage - 1)">Sebelumnya</button>
            <span>Halaman {{ currentPage }} dari {{ totalPages }}</span>
            <button :disabled="currentPage === totalPages" @click="changePage(currentPage + 1)">Berikutnya</button>
        </div>

        <TransactionModal v-if="selectedTransaction" :transaction="selectedTransaction" @close="selectedTransaction = null" />
    </div>
</template>

<style scoped>
.blockchain-transactions {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
    margin-bottom: 30px;
    color: #2c3e50;
}

.filter-controls {
    display: flex;
    gap: 15px;
    margin-bottom: 20px;
}

.filter-controls select,
.filter-controls input {
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 16px;
}

.filter-controls select {
    min-width: 200px;
}

.filter-controls input {
    flex-grow: 1;
}

.transactions-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 20px;
}

.transaction-card {
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 15px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: transform 0.2s;
}

.transaction-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.tx-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
    padding-bottom: 10px;
    border-bottom: 1px solid #eee;
}

.tx-id {
    font-weight: bold;
    font-size: 14px;
    color: #555;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: 70%;
}

.tx-chain {
    padding: 3px 8px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: bold;
    text-transform: uppercase;
}

.tx-chain.ethereum {
    background-color: #627eea;
    color: white;
}

.tx-chain.quorum {
    background-color: #5848a2;
    color: white;
}

.tx-chain.fabric {
    background-color: #a163ff;
    color: white;
}

.tx-details {
    margin-bottom: 15px;
}

.tx-row {
    display: flex;
    margin-bottom: 8px;
    font-size: 14px;
}

.tx-row .label {
    font-weight: bold;
    min-width: 60px;
    color: #666;
}

.tx-row .value {
    flex-grow: 1;
    word-break: break-all;
}

.tx-footer {
    display: flex;
    justify-content: flex-end;
}

.details-btn {
    padding: 6px 12px;
    background-color: #3498db;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.2s;
}

.details-btn:hover {
    background-color: #2980b9;
}

.loading,
.error,
.empty {
    text-align: center;
    padding: 40px;
    font-size: 18px;
    color: #666;
}

.error {
    color: #e74c3c;
}

.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 15px;
    margin-top: 30px;
    padding: 15px 0;
}

.pagination button {
    padding: 8px 16px;
    background-color: #f8f9fa;
    border: 1px solid #ddd;
    border-radius: 4px;
    cursor: pointer;
}

.pagination button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.pagination button:not(:disabled):hover {
    background-color: #e9ecef;
}
</style>
