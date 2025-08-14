<template>
    <section id="how-it-works" class="how-it-works">
        <div class="container">
            <h2 class="section-title">Bagaimana Cara Verifikasi?</h2>

            <div class="steps-container">
                <div class="step-card" v-for="(step, index) in steps" :key="index">
                    <div class="step-number">{{ index + 1 }}</div>
                    <div class="step-icon" :class="step.iconClass">
                        <component :is="icons[step.iconName]" />
                    </div>
                    <h3 class="step-title">{{ step.title }}</h3>
                    <p class="step-description">{{ step.description }}</p>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import { FileText as DocumentText, Network as ShareNetwork, ShieldCheck } from 'lucide-vue-next';
import { shallowRef } from 'vue';

// Create reactive references for icons
const icons = shallowRef({
    DocumentText,
    ShareNetwork,
    ShieldCheck
});

const steps = [
    {
        title: 'Upload Ijazah',
        description: 'Unggah scan ijazah Anda.',
        iconName: 'DocumentText',
        iconClass: 'blue-icon'
    },
    {
        title: 'Verifikasi Blockchain',
        description: 'Data diverifikasi di jaringan terdesentralisasi.',
        iconName: 'ShareNetwork',
        iconClass: 'gold-icon'
    },
    {
        title: 'Lihat Hasil',
        description: 'Dapatkan laporan keaslian dalam hitungan menit.',
        iconName: 'ShieldCheck',
        iconClass: 'green-icon'
    }
];
</script>

<style scoped>
/* Your existing styles remain the same */
.how-it-works {
    padding: 4rem 1rem;
    background-color: #f8fafc;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
}

.section-title {
    text-align: center;
    color: #2563eb;
    font-size: 2rem;
    margin-bottom: 3rem;
}

.steps-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 2rem;
}

.step-card {
    flex: 1;
    min-width: 280px;
    max-width: 350px;
    background: white;
    border-radius: 12px;
    padding: 2rem;
    text-align: center;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
    transition: transform 0.3s ease;
    position: relative;
    overflow: hidden;
}

.step-card:hover {
    transform: translateY(-5px);
}

.step-number {
    position: absolute;
    top: 10px;
    left: 10px;
    font-size: 1.5rem;
    font-weight: bold;
    color: #e2e8f0;
    z-index: 1;
}

.step-icon {
    width: 80px;
    height: 80px;
    margin: 0 auto 1.5rem;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    z-index: 2;
}

.step-icon svg {
    width: 40px;
    height: 40px;
}

.blue-icon {
    background-color: #dbeafe;
    color: #3b82f6;
}

.gold-icon {
    background-color: #fef3c7;
    color: #d97706;
}

.green-icon {
    background-color: #dcfce7;
    color: #16a34a;
}

.step-title {
    color: #1e40af;
    font-size: 1.25rem;
    margin-bottom: 1rem;
}

.step-description {
    color: #4b5563;
    line-height: 1.6;
}

@media (max-width: 768px) {
    .steps-container {
        flex-direction: column;
        align-items: center;
    }

    .step-card {
        width: 100%;
        max-width: 400px;
    }
}
</style>
