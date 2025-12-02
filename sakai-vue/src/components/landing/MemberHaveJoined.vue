<script setup>
import router from '@/router';
import { ref } from 'vue';
import LazyImage from '../LazyImage.vue';
const useLightBackground = ref(true);
const hoverIndex = ref(-1);

const schools = ref([
    {
        name: 'Universitas Indonesia',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: true
    },
    {
        name: 'Institut Teknologi Bandung',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: true
    },
    {
        name: 'Universitas Gadjah Mada',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: true
    },
    {
        name: 'Universitas Brawijaya',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: true
    },
    {
        name: 'SMA Negeri 1 Jakarta',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: false
    },
    {
        name: 'SMA Negeri 3 Bandung',
        logo: 'https://localhost/static/landingpage/SMAN_5_Bandung.png',
        verified: true
    },
    {
        name: 'SMA Negeri 8 Surabaya',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: false
    },
    {
        name: 'SMA Negeri 2 Yogyakarta',
        logo: 'https://localhost/static/landingpage/pasundan.png',
        verified: true
    }
]);
</script>

<template>
    <section class="partner-schools" :class="{ 'bg-light-blue': useLightBackground }">
        <div class="container">
            <h2 class="section-title">Sekolah Terpercaya yang Telah Bergabung</h2>
            <p class="section-subtitle">Bergabung dengan 500+ institusi pendidikan di Indonesia</p>

            <!-- Logo Grid -->
            <div class="logo-grid">
                <div v-for="(school, index) in schools" :key="index" class="logo-item" :class="{ verified: school.verified }">
                    <LazyImage :src="school.logo" :alt="school.name" container-class="logo-image" @mouseover="hoverIndex = index" @mouseleave="hoverIndex = -1" />
                    <!-- <img :src="school.logo" :alt="school.name" class="logo-image" @mouseover="hoverIndex = index" @mouseleave="hoverIndex = -1" /> -->
                    <span v-if="school.verified" class="verified-badge">✔️</span>
                </div>
            </div>

            <!-- CTA Mini -->
            <div class="cta-container">
                <p class="cta-text">Sekolah Anda belum terdaftar?</p>
                <button class="cta-button" @click="router.push({ name: 'register' })">Ajukan Kolaborasi</button>
            </div>
        </div>
    </section>
</template>

<style scoped>
.partner-schools {
    padding: 4rem 1rem;
    transition: background-color 0.3s ease;
}

.bg-light-blue {
    background-color: #f5f9ff;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
}

.section-title {
    color: #1e3a8a;
    text-align: center;
    font-size: 2rem;
    margin-bottom: 0.5rem;
}

.section-subtitle {
    color: #6b7280;
    text-align: center;
    font-size: 0.9rem;
    margin-bottom: 3rem;
}

.logo-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1.5rem;
    margin-bottom: 3rem;
}

.logo-item {
    position: relative;
    background: white;
    padding: 1.5rem;
    border-radius: 8px;
    border: 1px solid #eeeeee;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 120px;
    transition: all 0.3s ease;
}

.logo-image {
    max-width: 30%;
    max-height: 80px;
    filter: grayscale(100%);
    opacity: 0.8;
    transition: all 0.3s ease;
}

.logo-item:hover .logo-image {
    filter: grayscale(0%);
    opacity: 1;
    transform: scale(1.05);
}

.logo-item:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.verified-badge {
    position: absolute;
    top: 8px;
    right: 8px;
    background: gold;
    color: white;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.cta-container {
    text-align: center;
    margin-top: 2rem;
}

.cta-text {
    color: #6b7280;
    margin-bottom: 1rem;
}

.cta-button {
    background: transparent;
    color: #2563eb;
    border: 2px solid #2563eb;
    padding: 0.5rem 1.5rem;
    border-radius: 50px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
}

.cta-button:hover {
    background: #2563eb;
    color: white;
}

/* Responsive Styles */
@media (max-width: 1024px) {
    .logo-grid {
        grid-template-columns: repeat(3, 1fr);
    }
}

@media (max-width: 768px) {
    .logo-grid {
        grid-template-columns: repeat(2, 1fr);
    }
}

@media (max-width: 480px) {
    .logo-grid {
        grid-template-columns: 1fr;
    }

    .logo-item {
        height: 100px;
    }
}
</style>
