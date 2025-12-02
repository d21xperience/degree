<script setup>
import { ref } from 'vue';
import LazyImage from './LazyImage.vue';

const fullWidthLayout = ref(false); // Set to true for full-width blue background

// const heroData = computed(() => ({
//     backgroundImage: {
//         src: 'https://localhost/static/landingpage/7.jpg',
//         alt: 'Latar belakang teknologi blockchain dengan visualisasi keamanan data'
//     }
// }));
</script>

<template>
    <section id="about" class="about-section" :class="{ 'full-width': fullWidthLayout }">
        <div class="container">
            <div class="section-header">
                <h2 class="section-title">Tentang Verifikasi Ijazah</h2>
                <p class="section-subtitle">Membangun Masa Depan dengan Kepercayaan dan Teknologi Blockchain</p>
            </div>

            <div class="content-wrapper">
                <div class="image-column">
                    <LazyImage src="https://localhost/static/landingpage/4.jpg" alt="Tes" @load="console.log('✅ Load sukses:', $event)" @error="console.error('❌ Error:', $event)" />

                    <!-- <LazyImage alt="Tes" src="https://localhost/static/landingpage/4.jpg" :use-native-lazy="true" :show-spinner="true" image-class="hero-background-image opacity-0" loaded-class="opacity-70 transition-opacity duration-500" /> -->
                    <!-- <img src="https://localhost/static/landingpage/4.jpg" alt="Ilustrasi verifikasi blockchain" class="feature-image" /> -->
                </div>

                <div class="text-column">
                    <div class="description">
                        <p class="paragraph"><strong>Visi:</strong> "Menjadi platform verifikasi ijazah berbasis blockchain terpercaya di Indonesia."</p>
                        <p class="paragraph"><strong>Misi:</strong> "Memerangi ijazah palsu dengan transparansi teknologi terdesentralisasi."</p>
                        <p class="paragraph"><strong>Keunikan:</strong> "Kolaborasi dengan Kemdikbud, sekolah, dan perusahaan untuk ekosistem pendidikan yang aman."</p>
                    </div>

                    <a href="#" class="learn-more"> Pelajari Lebih Lanjut <span class="arrow">→</span> </a>

                    <div class="key-points">
                        <div class="point">
                            <span class="point-icon">🛡️</span>
                            <span class="point-text">Data tidak bisa dimanipulasi</span>
                        </div>
                        <div class="point">
                            <span class="point-icon">🌍</span>
                            <span class="point-text">Jangkauan seluruh Indonesia</span>
                        </div>
                        <div class="point">
                            <span class="point-icon">🤝</span>
                            <span class="point-text">Dibangun oleh pakar pendidikan dan teknologi</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="stats-container">
                <StatWidget label="Ijazah Terverifikasi" :target-number="Number(100000)" />
                <StatWidget label="Sekolah Mitra" :target-number="Number(500)" />
                <StatWidget label="Kepuasan Pengguna" :target-number="Number(99)" />
                <!-- <div class="stat-box">
                    
                    <div class="stat-number">10.000+</div>
                    <div class="stat-label">Ijazah Terverifikasi</div>
                </div> -->
                <!-- <div class="stat-box">
                    <div class="stat-number">500+</div>
                    <div class="stat-label">Sekolah Mitra</div>
                </div>
                <div class="stat-box">
                    <div class="stat-number">99%</div>
                    <div class="stat-label">Kepuasan Pengguna</div>
                </div> -->
            </div>

            <div class="partners-container">
                <h3 class="partners-title">Mitra Resmi Kami</h3>
                <div class="partners-grid">
                    <img src="https://localhost/static/landingpage/logo_pusat.jpg" alt="Kemdikbud" class="partner-logo" />
                    <img src="https://localhost/static/landingpage/komdigi.png" alt="Asosiasi Sekolah" class="partner-logo" />
                    <img src="https://localhost/static/landingpage/jabar.png" alt="Mitra Teknologi" class="partner-logo" />
                    <img src="https://localhost/static/landingpage/disdik8.png" alt="Mitra Teknologi" class="partner-logo" />
                </div>
            </div>
        </div>
    </section>
</template>

<style scoped>
.about-section {
    padding: 4rem 0;
    /* background-color: white; */
}

.about-section.full-width {
    background-color: #f5f9ff;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 5rem; /* 80px left-right */
}

.section-header {
    text-align: center;
    margin-bottom: 3rem;
}

.section-title {
    color: #1e3a8a;
    font-family: 'Poppins', sans-serif;
    font-weight: 700;
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

.section-subtitle {
    color: #666666;
    font-family: 'Inter', sans-serif;
    font-size: 1.25rem;
    max-width: 700px;
    margin: 0 auto;
}

.content-wrapper {
    display: flex;
    gap: 3rem;
    align-items: center;
    margin-bottom: 4rem;
}

.image-column,
.text-column {
    flex: 1;
}

.feature-image {
    width: 100%;
    height: auto;
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.description {
    margin-bottom: 2rem;
}

.paragraph {
    color: #666666;
    font-family: 'Inter', sans-serif;
    line-height: 1.6;
    margin-bottom: 1rem;
}

.learn-more {
    display: inline-block;
    color: #1e3a8a;
    font-family: 'Inter', sans-serif;
    font-weight: 600;
    text-decoration: none;
    margin-bottom: 2rem;
    transition: color 0.3s ease;
}

.learn-more:hover {
    color: #ffd700;
}

.arrow {
    margin-left: 0.5rem;
    transition: transform 0.3s ease;
}

.learn-more:hover .arrow {
    transform: translateX(4px);
}

.key-points {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.point {
    display: flex;
    align-items: center;
    gap: 0.8rem;
}

.point-icon {
    font-size: 1.2rem;
}

.point-text {
    color: #666666;
    font-family: 'Inter', sans-serif;
}

.stats-container {
    display: flex;
    justify-content: space-around;
    margin: 4rem 0;
}

.partners-container {
    margin-top: 4rem;
    text-align: center;
}

.partners-title {
    color: #1e3a8a;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    margin-bottom: 2rem;
}

.partners-grid {
    display: flex;
    justify-content: center;
    gap: 3rem;
    flex-wrap: wrap;
}

.partner-logo {
    height: 100px;
    width: auto;
    object-fit: contain;
    filter: grayscale(100%);
    opacity: 0.7;
    transition: all 0.3s ease;
}

.partner-logo:hover {
    filter: grayscale(0%);
    opacity: 1;
}

@media (max-width: 1024px) {
    .container {
        padding: 0 3rem;
    }

    .content-wrapper {
        flex-direction: column;
    }

    .image-column,
    .text-column {
        width: 100%;
    }

    .stats-container {
        flex-wrap: wrap;
        gap: 1rem;
    }

    .stat-box {
        flex: 1;
        min-width: 160px;
    }
}

@media (max-width: 768px) {
    .container {
        padding: 0 2rem;
    }

    .section-title {
        font-size: 2rem;
    }

    .section-subtitle {
        font-size: 1.1rem;
    }

    .partners-grid {
        gap: 1.5rem;
    }
}
</style>
