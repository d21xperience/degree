<script setup>
import { ref } from 'vue';
const UPLOAD_ENDPOINT = `${API_BASE_URL}/api/ipfs/upload`;

// State untuk upload
const selectedFile = ref(null);
const isUploading = ref(false);
const uploadResult = ref(null);
const uploadError = ref(null);
const fileInput = ref(null);

// Handle pemilihan file
const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
        selectedFile.value = file;
        uploadResult.value = null;
        uploadError.value = null;
    }
};
// Upload file ke IPFS melalui backend Golang
const uploadFile = async () => {
    if (!selectedFile.value) return;

    isUploading.value = true;
    uploadError.value = null;

    try {
        const formData = new FormData();
        formData.append('file', selectedFile.value);

        const response = await fetch(UPLOAD_ENDPOINT, {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const result = await response.json();
        uploadResult.value = {
            cid: result.cid,
            size: result.size,
            name: selectedFile.value.name
        };

        // Reset input file
        if (fileInput.value) {
            fileInput.value.value = '';
        }
        selectedFile.value = null;
    } catch (error) {
        console.error('Upload error:', error);
        uploadError.value = `Upload failed: ${error.message}`;
    } finally {
        isUploading.value = false;
    }
};
// Helper functions
const gatewayUrl = (cid) => {
    return `https://ipfs.io/ipfs/${cid}`;
};
</script>
<template>
    <!-- Upload Section -->
    <div class="section">
        <h3>Upload File</h3>
        <input type="file" @change="handleFileChange" ref="fileInput" />
        <button @click="uploadFile" :disabled="!selectedFile || isUploading">
            {{ isUploading ? 'Uploading...' : 'Upload to IPFS' }}
        </button>
        <div v-if="uploadResult" class="result">
            <p>File uploaded successfully!</p>
            <p>
                CID: <code>{{ uploadResult.cid }}</code>
            </p>
            <p>Size: {{ formatBytes(uploadResult.size) }}</p>
            <a :href="gatewayUrl(uploadResult.cid)" target="_blank">View on IPFS Gateway</a>
        </div>
        <div v-if="uploadError" class="error">
            {{ uploadError }}
        </div>
    </div>
</template>
