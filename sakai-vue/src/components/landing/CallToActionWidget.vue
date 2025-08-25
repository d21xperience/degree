<script setup>
// Tidak perlu logic tambahan untuk komponen sederhana ini
</script>

<template>
    <section class="cta-section">
        <div class="cta-content">
            <h2 class="cta-title">Sekolah Anda Belum Terhubung?</h2>
            <p class="cta-subtitle">Daftarkan institusi Anda untuk integrasi verifikasi blockchain.</p>
            <button class="cta-button">Daftarkan Sekolah</button>
        </div>
    </section>
</template>

<style scoped>
.cta-section {
    /* background: linear-gradient(135deg, #e6f2ff, #b3d9ff); */
    padding: 3rem 1rem;
    text-align: center;
    border-radius: 8px;
    margin: 2rem 0;
}

.cta-content {
    max-width: 800px;
    margin: 0 auto;
}

/* .cta-title {
    font-size: 2rem;
    color: #333;
    margin-bottom: 1rem;
} */

/* .cta-subtitle {
    font-size: 1.2rem;
    color: #555;
    margin-bottom: 2rem;
} */

.cta-button {
    background: linear-gradient(135deg, #ffd700, #4caf50);
    color: white;
    font-weight: bold;
    padding: 0.8rem 2rem;
    border: none;
    border-radius: 50px;
    font-size: 1.1rem;
    cursor: pointer;
    transition:
        transform 0.3s,
        box-shadow 0.3s;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.cta-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}
</style>
