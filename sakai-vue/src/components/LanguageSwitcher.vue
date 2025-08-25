<!-- components/LanguageSwitcher.vue -->
<script setup>
import { ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';

const { locale } = useI18n();
const currentLang = ref(locale.value);

// Ganti bahasa
const switchTo = (lang) => {
    locale.value = lang;
    currentLang.value = lang;
    localStorage.setItem('preferred-language', lang);
};

// Sinkronkan jika locale berubah di tempat lain
watch(locale, (newVal) => {
    currentLang.value = newVal;
});

// Inisialisasi dari localStorage
(function init() {
    const saved = localStorage.getItem('preferred-language');
    if (saved === 'id' || saved === 'en') {
        locale.value = saved;
        currentLang.value = saved;
    }
})();
</script>

<template>
    <div class="language-switcher">
        <!-- Tombol untuk ganti ke Bahasa Inggris -->
        <button v-if="currentLang === 'id'" @click="switchTo('en')" class="flex items-center space-x-2 px-4 py-2 hover:bg-gray-50 border-none transition-all duration-200 text-sm font-medium text-gray-700">
            <!-- Bendera Indonesia → Inggris -->
            <span class="text-lg" aria-label="UK flag">🇬🇧</span>
            <span>English</span>
        </button>

        <!-- Tombol untuk ganti ke Bahasa Indonesia -->
        <button v-else @click="switchTo('id')" class="flex items-center space-x-2 px-4 py-2 hover:bg-gray-50 border-none transition-all duration-200 outline-none focus:ring-2 text-sm font-medium text-gray-700">
            <!-- Bendera Inggris → Indonesia -->
            <span class="text-lg" aria-label="Indonesia flag">🇮🇩</span>
            <span>Indonesia</span>
        </button>
    </div>
</template>
