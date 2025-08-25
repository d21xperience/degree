<script setup>
import { computed, ref } from 'vue';

// API endpoints - sesuaikan dengan backend Golang Anda
const API_BASE_URL = 'http://localhost:8080'; // atau URL backend Anda
const RETRIEVE_ENDPOINT = `${API_BASE_URL}/api/ipfs/retrieve`;

// State untuk retrieve
const cidInput = ref('');
const isRetrieving = ref(false);
const retrievedFile = ref(null);
const fileContent = ref('');
const filePreviewUrl = ref('');
const retrieveError = ref(null);

// Retrieve file dari IPFS melalui backend Golang
const retrieveFile = async () => {
    if (!cidInput.value) return;

    isRetrieving.value = true;
    retrieveError.value = null;
    retrievedFile.value = null;
    filePreviewUrl.value = '';

    try {
        const response = await fetch(`${RETRIEVE_ENDPOINT}/${cidInput.value}`);

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        // Dapatkan metadata dari header
        const contentType = response.headers.get('content-type');
        const contentDisposition = response.headers.get('content-disposition');
        let filename = 'file';

        if (contentDisposition) {
            const filenameMatch = contentDisposition.match(/filename="?(.+)"?/i);
            if (filenameMatch && filenameMatch[1]) {
                filename = filenameMatch[1];
            }
        }

        // Handle response berdasarkan tipe konten
        if (contentType.startsWith('text/')) {
            fileContent.value = await response.text();
            retrievedFile.value = {
                name: filename,
                type: contentType,
                size: fileContent.value.length
            };
        } else {
            const blob = await response.blob();
            filePreviewUrl.value = URL.createObjectURL(blob);
            retrievedFile.value = {
                name: filename,
                type: contentType,
                size: blob.size
            };
        }
    } catch (error) {
        console.error('Retrieve error:', error);
        retrieveError.value = `Retrieve failed: ${error.message}`;
    } finally {
        isRetrieving.value = false;
    }
};

const formatBytes = (bytes, decimals = 2) => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm) + ' ' + sizes[i]);
};

// Computed properties untuk tipe file
const isImageFile = computed(() => {
    return retrievedFile.value?.type?.startsWith('image/');
});

const isVideoFile = computed(() => {
    return retrievedFile.value?.type?.startsWith('video/');
});

const isAudioFile = computed(() => {
    return retrievedFile.value?.type?.startsWith('audio/');
});

const isTextFile = computed(() => {
    return retrievedFile.value?.type?.startsWith('text/') || fileContent.value !== '';
});
</script>

<template>
    <div class="ipfs-container">
        <!-- Download Section -->
        <div class="section">
            <h3>Retrieve File</h3>
            <input type="text" v-model="cidInput" placeholder="Enter IPFS CID" class="cid-input" />
            <button @click="retrieveFile" :disabled="!cidInput || isRetrieving">
                {{ isRetrieving ? 'Retrieving...' : 'Retrieve from IPFS' }}
            </button>

            <div v-if="retrievedFile" class="file-preview">
                <h4>File Preview</h4>
                <p>Name: {{ retrievedFile.name }}</p>
                <p>Type: {{ retrievedFile.type }}</p>
                <p>Size: {{ formatBytes(retrievedFile.size) }}</p>

                <!-- Preview berdasarkan tipe file -->
                <img v-if="isImageFile" :src="filePreviewUrl" alt="IPFS File Preview" class="preview-image" />
                <video v-else-if="isVideoFile" controls class="preview-video">
                    <source :src="filePreviewUrl" :type="retrievedFile.type" />
                </video>
                <audio v-else-if="isAudioFile" controls class="preview-audio">
                    <source :src="filePreviewUrl" :type="retrievedFile.type" />
                </audio>
                <div v-else-if="isTextFile" class="text-preview">
                    <pre>{{ fileContent }}</pre>
                </div>
                <div v-else>
                    <p>Binary file - no preview available</p>
                    <a :href="filePreviewUrl" download>Download File</a>
                </div>
            </div>

            <div v-if="retrieveError" class="error">
                {{ retrieveError }}
            </div>
        </div>
    </div>
</template>

<style scoped>
.ipfs-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
}

.section {
    margin-bottom: 30px;
    padding: 20px;
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    background-color: #f9f9f9;
}

h2,
h3,
h4 {
    color: #2c3e50;
}

input[type='file'],
.cid-input {
    display: block;
    width: 100%;
    padding: 10px;
    margin: 10px 0;
    border: 1px solid #ddd;
    border-radius: 4px;
}

button {
    background-color: #3498db;
    color: white;
    padding: 10px 15px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.3s;
}

button:hover {
    background-color: #2980b9;
}

button:disabled {
    background-color: #95a5a6;
    cursor: not-allowed;
}

.result,
.error {
    margin-top: 15px;
    padding: 15px;
    border-radius: 4px;
}

.result {
    background-color: #e8f8f5;
    border: 1px solid #a3e4d7;
}

.error {
    background-color: #fdedec;
    border: 1px solid #f5b7b1;
    color: #c0392b;
}

.file-preview {
    margin-top: 20px;
    padding: 15px;
    background-color: #fff;
    border: 1px solid #eee;
    border-radius: 4px;
}

.preview-image {
    max-width: 100%;
    max-height: 400px;
    display: block;
    margin: 10px auto;
    border-radius: 4px;
}

.preview-video,
.preview-audio {
    width: 100%;
    margin-top: 10px;
    border-radius: 4px;
}

.text-preview {
    margin-top: 10px;
    padding: 10px;
    background-color: #f5f5f5;
    border-radius: 4px;
    max-height: 300px;
    overflow-y: auto;
}

pre {
    white-space: pre-wrap;
    word-wrap: break-word;
}

code {
    font-family: monospace;
    background-color: #f5f5f5;
    padding: 2px 4px;
    border-radius: 3px;
}

a {
    color: #3498db;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}
</style>
