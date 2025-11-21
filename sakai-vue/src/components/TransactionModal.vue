<script setup>
import { defineEmits, defineProps } from 'vue';

const props = defineProps({
    transaction: {
        type: Object,
        required: true
    }
});

const emit = defineEmits(['close']);

const emitClose = () => {
    emit('close');
};

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
</script>

<template>
    <div class="modal-overlay" @click.self="emitClose">
        <div class="modal-content">
            <div class="modal-header">
                <h2>Detail Transaksi</h2>
                <button class="close-btn" @click="emitClose">&times;</button>
            </div>

            <div class="modal-body">
                <div class="detail-row">
                    <span class="detail-label">ID:</span>
                    <span class="detail-value">{{ transaction.id }}</span>
                </div>

                <div class="detail-row">
                    <span class="detail-label">Blockchain:</span>
                    <span class="detail-value">{{ formatBlockchainName(transaction.blockchain) }}</span>
                </div>

                <div class="detail-row">
                    <span class="detail-label">Dari:</span>
                    <span class="detail-value">{{ transaction.from }}</span>
                </div>

                <div class="detail-row">
                    <span class="detail-label">Ke:</span>
                    <span class="detail-value">{{ transaction.to }}</span>
                </div>

                <div class="detail-row">
                    <span class="detail-label">Nilai:</span>
                    <span class="detail-value">{{ formatValue(transaction.value, transaction.blockchain) }}</span>
                </div>

                <div class="detail-row">
                    <span class="detail-label">Waktu:</span>
                    <span class="detail-value">{{ formatTimestamp(transaction.timestamp) }}</span>
                </div>

                <div v-if="transaction.blockNumber" class="detail-row">
                    <span class="detail-label">Nomor Block:</span>
                    <span class="detail-value">{{ transaction.blockNumber }}</span>
                </div>

                <div v-if="transaction.hash" class="detail-row">
                    <span class="detail-label">Hash:</span>
                    <span class="detail-value">{{ transaction.hash }}</span>
                </div>

                <div v-if="transaction.status" class="detail-row">
                    <span class="detail-label">Status:</span>
                    <span class="detail-value">{{ transaction.status }}</span>
                </div>

                <div v-if="transaction.inputData" class="detail-row">
                    <span class="detail-label">Input Data:</span>
                    <span class="detail-value">{{ transaction.inputData }}</span>
                </div>
            </div>

            <div class="modal-footer">
                <button class="close-btn" @click="emitClose">Tutup</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background-color: white;
    border-radius: 8px;
    width: 90%;
    max-width: 600px;
    max-height: 90vh;
    overflow-y: auto;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    border-bottom: 1px solid #eee;
}

.modal-header h2 {
    margin: 0;
    color: #2c3e50;
}

.close-btn {
    background: none;
    border: none;
    font-size: 24px;
    cursor: pointer;
    color: #666;
}

.close-btn:hover {
    color: #333;
}

.modal-body {
    padding: 20px;
}

.detail-row {
    margin-bottom: 15px;
}

.detail-label {
    font-weight: bold;
    display: inline-block;
    min-width: 120px;
    color: #666;
}

.detail-value {
    word-break: break-all;
}

.modal-footer {
    padding: 15px 20px;
    border-top: 1px solid #eee;
    display: flex;
    justify-content: flex-end;
}

.modal-footer .close-btn {
    padding: 8px 16px;
    background-color: #3498db;
    color: white;
    border-radius: 4px;
    font-size: 16px;
}

.modal-footer .close-btn:hover {
    background-color: #2980b9;
    color: white;
}
</style>
