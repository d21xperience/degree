<script setup>
import { ref } from 'vue';

const emit = defineEmits(['change']);
const previewUrl = ref(null);
const dragging = ref(false);

function onDrop(e) {
    e.preventDefault();
    dragging.value = false;
    const file = e.dataTransfer.files[0];
    handleFile(file);
}

function onSelect(e) {
    const file = e.target.files[0];
    handleFile(file);
}

function handleFile(file) {
    if (!file) return;

    if (!file.type.startsWith('image/')) {
        alert('File harus berupa gambar.');
        return;
    }

    previewUrl.value = URL.createObjectURL(file);
    emit('change', file);
}

function onDragOver(e) {
    e.preventDefault();
    dragging.value = true;
}

function onDragLeave() {
    dragging.value = false;
}
</script>

<template>
    <div class="w-full space-y-4">
        <!-- Dropzone -->
        <div
            class="border-2 border-dashed rounded-xl p-8 flex flex-col items-center justify-center cursor-pointer transition-all select-none"
            :class="dragging ? 'border-blue-500 bg-blue-50 shadow-inner' : 'border-gray-300 bg-white hover:bg-gray-50'"
            @dragover="onDragOver"
            @dragleave="onDragLeave"
            @drop="onDrop"
            @click="$refs.fileInput.click()"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-12 h-12 text-gray-400 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6H16a5 5 0 010 10H7z" />
            </svg>

            <p class="text-gray-700 font-medium">Drag & Drop gambar</p>
            <p class="text-gray-500 text-sm">atau klik untuk memilih file</p>

            <input ref="fileInput" type="file" accept="image/*" class="hidden" @change="onSelect" />
        </div>

        <!-- Preview -->
        <div v-if="previewUrl" class="flex items-start space-x-4 p-4 bg-gray-50 border rounded-lg shadow-sm">
            <img :src="previewUrl" alt="preview" class="w-32 h-32 object-cover rounded-lg border shadow" />

            <div class="flex flex-col justify-between">
                <p class="text-gray-700 text-sm font-medium">Preview Gambar</p>
                <button class="mt-2 inline-flex items-center px-3 py-1.5 bg-red-500 text-white text-xs rounded hover:bg-red-600 transition" @click="previewUrl = null">Hapus Preview</button>
            </div>
        </div>
    </div>
</template>
