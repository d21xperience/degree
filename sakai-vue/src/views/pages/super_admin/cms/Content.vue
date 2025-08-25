<script setup>
import { ref } from 'vue';

const content = ref({
    hero: {
        title: '',
        subtitle: '',
        buttonText: ''
    },
    features: []
});

// Ambil konten dari API atau store
async function fetchContent() {
    const res = await fetch('/api/content');
    content.value = await res.json();
}

// Simpan ke API
async function saveContent() {
    await fetch('/api/content', {
        method: 'POST',
        body: JSON.stringify(content.value),
        headers: { 'Content-Type': 'application/json' }
    });
    alert('Konten berhasil disimpan!');
}

fetchContent();
</script>

<template>
    <div class="p-4">
        <h1 class="text-2xl font-bold mb-4">Admin Panel</h1>

        <!-- Hero Section -->
        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-2">Hero</h2>
            <input v-model="content.hero.title" placeholder="Title" class="block w-full mb-2 p-2 border rounded" />
            <input v-model="content.hero.subtitle" placeholder="Subtitle" class="block w-full mb-2 p-2 border rounded" />
            <input v-model="content.hero.buttonText" placeholder="Button Text" class="block w-full mb-2 p-2 border rounded" />
        </section>

        <!-- Features -->
        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-2">Features</h2>
            <div v-for="(feature, index) in content.features" :key="index" class="mb-4 border p-2 rounded">
                <input v-model="feature.title" placeholder="Feature Title" class="block w-full mb-1 p-2 border rounded" />
                <input v-model="feature.description" placeholder="Description" class="block w-full mb-1 p-2 border rounded" />
            </div>
        </section>

        <!-- Simpan -->
        <button @click="saveContent" class="bg-blue-500 text-white px-4 py-2 rounded">Simpan</button>
    </div>
</template>
