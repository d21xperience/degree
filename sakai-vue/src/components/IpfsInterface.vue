<template>
    <div class="ipfs-interface">
        <h1>IPFS Interface via Golang</h1>

        <div class="section">
            <h2>Node Info</h2>
            <button @click="getNodeInfo">Get Node Info</button>
            <div v-if="nodeInfo">
                <pre>{{ nodeInfo }}</pre>
            </div>
        </div>

        <div class="section">
            <h2>Upload File</h2>
            <input type="file" @change="handleFileUpload" />
            <button @click="uploadFile" :disabled="!selectedFile">Upload</button>
            <div v-if="uploadResult">
                <p>File uploaded with CID: {{ uploadResult.cid }}</p>
                <button @click="viewFile(uploadResult.cid)">View File</button>
            </div>
        </div>

        <div class="section" v-if="fileContent">
            <h2>File Content</h2>
            <pre>{{ fileContent }}</pre>
        </div>

        <div class="status" v-if="status">
            {{ status }}
        </div>
        <div class="error" v-if="error">Error: {{ error }}</div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const nodeInfo = ref(null);
const selectedFile = ref(null);
const uploadResult = ref(null);
const fileContent = ref(null);
const status = ref('');
const error = ref('');

const apiBaseUrl = 'http://localhost:8080/api/ipfs';

async function getNodeInfo() {
    try {
        status.value = 'Fetching node info...';
        error.value = '';

        const [versionRes, idRes] = await Promise.all([fetch(`${apiBaseUrl}/version`), fetch(`${apiBaseUrl}/id`)]);

        if (!versionRes.ok || !idRes.ok) {
            throw new Error('Failed to fetch node info');
        }

        const version = await versionRes.json();
        const id = await idRes.json();

        nodeInfo.value = { version, id };
        status.value = 'Node info fetched successfully';
    } catch (err) {
        error.value = err.message;
        status.value = '';
        console.error('Error fetching node info:', err);
    }
}

function handleFileUpload(event) {
    selectedFile.value = event.target.files[0];
}

async function uploadFile() {
    if (!selectedFile.value) return;

    try {
        status.value = 'Uploading file...';
        error.value = '';

        const formData = new FormData();
        formData.append('file', selectedFile.value);

        const response = await fetch(`${apiBaseUrl}/add`, {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error('File upload failed');
        }

        uploadResult.value = await response.json();
        status.value = 'File uploaded successfully';
    } catch (err) {
        error.value = err.message;
        status.value = '';
        console.error('Error uploading file:', err);
    }
}

async function viewFile(cid) {
    try {
        status.value = 'Fetching file content...';
        error.value = '';

        const response = await fetch(`${apiBaseUrl}/cat/${cid}`);

        if (!response.ok) {
            throw new Error('Failed to fetch file content');
        }

        if (selectedFile.value.type.startsWith('text/') || selectedFile.value.type === 'application/json') {
            fileContent.value = await response.text();
        } else {
            fileContent.value = 'Binary file content not displayed';
        }

        status.value = 'File content fetched successfully';
    } catch (err) {
        error.value = err.message;
        status.value = '';
        console.error('Error fetching file content:', err);
    }
}
</script>

<style scoped>
.ipfs-interface {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

.section {
    margin-bottom: 30px;
    padding: 15px;
    border: 1px solid #ddd;
    border-radius: 5px;
}

button {
    margin: 10px 0;
    padding: 8px 16px;
    background-color: #42b983;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

button:disabled {
    background-color: #cccccc;
    cursor: not-allowed;
}

.status {
    color: #42b983;
    margin-top: 10px;
}

.error {
    color: #ff4444;
    margin-top: 10px;
}

pre {
    background-color: #f5f5f5;
    padding: 10px;
    border-radius: 4px;
    overflow-x: auto;
}
</style>
