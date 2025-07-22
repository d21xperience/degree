<template>
    <section class="trust-section">
        <h2 class="trust-title">Mengapa Memilih Kami?</h2>
        <div class="benefits-grid">
            <div v-for="benefit in benefits" :key="benefit.title" class="benefit-card">
                <div class="benefit-icon">{{ benefit.icon }}</div>
                <h3 class="benefit-title">{{ benefit.title }}</h3>
                <p class="benefit-description">{{ benefit.description }}</p>
            </div>
        </div>
    </section>
</template>

<script setup>
const benefits = [
    {
        icon: '🔒',
        title: 'Aman',
        description: 'Enkripsi blockchain tingkat tinggi.'
    },
    {
        icon: '👁️',
        title: 'Transparan',
        description: 'Riwayat verifikasi tidak bisa diubah.'
    },
    {
        icon: '⚡',
        title: 'Cepat',
        description: 'Proses otomatis dalam 5 menit.'
    },
    {
        icon: '🌐',
        title: 'Akses Nasional',
        description: 'Terintegrasi dengan sekolah seluruh Indonesia.'
    }
];
</script>

<style scoped>
.trust-section {
    padding: 4rem 1rem;
    max-width: 1200px;
    margin: 0 auto;
}

.trust-title {
    text-align: center;
    color: #2563eb;
    font-size: 2rem;
    margin-bottom: 3rem;
}

.benefits-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    padding: 0 1rem;
}

.benefit-card {
    background: white;
    border-radius: 12px;
    padding: 2rem;
    text-align: center;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    transition:
        transform 0.3s ease,
        box-shadow 0.3s ease;
}

.benefit-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 15px rgba(0, 0, 0, 0.1);
}

.benefit-icon {
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

.benefit-title {
    color: #1e40af;
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
}

.benefit-description {
    color: #4b5563;
    line-height: 1.5;
}

@media (max-width: 640px) {
    .benefits-grid {
        grid-template-columns: 1fr;
    }
}
</style>
