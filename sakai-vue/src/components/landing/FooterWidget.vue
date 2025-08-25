<script setup>
import { Facebook, Instagram, Linkedin, Twitter, Youtube } from 'lucide-vue-next';

const links = [
    { text: 'Beranda', url: '#' },
    { text: 'Tentang', url: '#' },
    { text: 'Fitur', url: '#' },
    { text: 'Harga', url: '#' },
    { text: 'Kontak', url: '#' }
];

const socialMedia = [
    { icon: Facebook, url: 'https://facebook.com' },
    { icon: Twitter, url: 'https://twitter.com' },
    { icon: Instagram, url: 'https://instagram.com' },
    { icon: Linkedin, url: 'https://linkedin.com' },
    { icon: Youtube, url: 'https://youtube.com' }
];
</script>

<template>
    <footer class="footer-section">
        <div class="footer-container">
            <!-- Left Column -->
            <div class="footer-left">
                <img src="https://via.placeholder.com/80x30?text=LOGO" alt="Verifikasi Ijazah Logo" class="footer-logo" />
                <p class="copyright-text">© 2023 VerifikasiIjazah.id. All rights reserved.</p>
            </div>

            <!-- Middle Column -->
            <div class="footer-middle">
                <h4 class="footer-title">Tautan Cepat</h4>
                <ul class="footer-links">
                    <li v-for="(link, index) in links" :key="index">
                        <a :href="link.url" class="footer-link">{{ link.text }}</a>
                    </li>
                    <li>
                        <a href="#" class="footer-link">Kebijakan Privasi</a>
                    </li>
                </ul>
            </div>

            <!-- Right Column -->
            <div class="footer-right">
                <h4 class="footer-title">Hubungi Kami</h4>
                <div class="social-icons">
                    <a v-for="(social, index) in socialMedia" :key="index" :href="social.url" class="social-icon" target="_blank">
                        <component :is="social.icon" />
                    </a>
                </div>
                <div class="footer-email">
                    <span class="email-icon">✉️</span>
                    <a href="mailto:info@verifikasiijazah.id" class="email-link">info@verifikasiijazah.id</a>
                </div>
            </div>
        </div>
    </footer>
</template>

<style scoped>
.footer-section {
    background-color: #f8fafc;
    padding: 3rem 1rem;
    border-top: 1px solid #e2e8f0;
}

.footer-container {
    max-width: 1200px;
    margin: 0 auto;
    display: flex;
    flex-wrap: wrap;
    gap: 2rem;
    justify-content: space-between;
}

.footer-left,
.footer-middle,
.footer-right {
    flex: 1;
    min-width: 200px;
}

.footer-logo {
    width: 80px;
    height: auto;
    margin-bottom: 1rem;
}

.copyright-text {
    color: #64748b;
    font-size: 0.875rem;
}

.footer-title {
    color: #1e3a8a;
    font-size: 1.125rem;
    margin-bottom: 1.25rem;
    font-weight: 600;
}

.footer-links {
    list-style: none;
    padding: 0;
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
}

.footer-link {
    color: #475569;
    text-decoration: none;
    transition: color 0.2s ease;
    font-size: 0.9375rem;
}

.footer-link:hover {
    color: #2563eb;
}

.social-icons {
    display: flex;
    gap: 1rem;
    margin-bottom: 1.5rem;
}

.social-icon {
    color: #2563eb;
    width: 24px;
    height: 24px;
    transition: color 0.2s ease;
}

.social-icon:hover {
    color: #1e40af;
}

.footer-email {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    margin-top: 0.5rem;
}

.email-icon {
    font-size: 1rem;
}

.email-link {
    color: #475569;
    text-decoration: none;
    transition: color 0.2s ease;
    font-size: 0.9375rem;
}

.email-link:hover {
    color: #2563eb;
}

@media (max-width: 768px) {
    .footer-container {
        flex-direction: column;
        gap: 2.5rem;
    }

    .footer-left,
    .footer-middle,
    .footer-right {
        min-width: 100%;
    }
}
</style>
