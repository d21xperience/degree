// src/i18n/index.js
import { createI18n } from 'vue-i18n';

const messages = {
    en: {
        platform: 'Platform',
        unselected: 'Unselected',
        network: 'Network',
        disconnected: 'Disconnected',
        connected: 'Connected',
        connectTitle: 'Connect',
        connectMessage: 'Connect to platform {platform} with network {network}?',
        yes: 'Yes',
        config: 'Config',
        contact: 'Contact',
        about: 'About',
        hero: {
            headline: 'Verified Diploma, Guaranteed Future',
            subtitle: 'Ensure your diploma authenticity with secure and transparent blockchain technology.',
            ctaText: 'Verify Now'
        }
    },
    id: {
        platform: 'Platform',
        unselected: 'Belum dipilih',
        network: 'Jaringan',
        disconnected: 'Terputus',
        connected: 'Terhubung',
        connectTitle: 'Hubungkan',
        connectMessage: 'Hubungkan ke platform {platform} dengan jaringan {network}?',
        yes: 'Ya',
        config: 'Konfigurasi',
        contact: 'Kontak',
        about: 'Tentang',
        hero: {
            headline: 'Ijazah Terverifikasi, Masa Depan Terjamin',
            subtitle: 'Pastikan keaslian ijazah Anda dengan teknologi blockchain yang aman dan transparan.',
            ctaText: 'Verifikasi Sekarang'
        }
    }
};

const i18n = createI18n({
    locale: 'id', // default language
    fallbackLocale: 'en',
    messages
});

export default i18n;
