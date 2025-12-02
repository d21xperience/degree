<script setup>
import { computed, nextTick, onMounted, onUnmounted, reactive, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import LazyImage from '../LazyImage.vue';

const { t } = useI18n();
// Emits
const emit = defineEmits(['verification-submit', 'dialog-open', 'dialog-close', 'background-load']);

// Lifecycle
onMounted(() => {
    document.addEventListener('keydown', handleKeydown);
});

onUnmounted(() => {
    document.removeEventListener('keydown', handleKeydown);
});
// Methods
// Reactive data
const heroData = computed(() => ({
    headline: t('hero.headline'),
    subtitle: t('hero.subtitle'),
    ctaText: t('hero.ctaText'),
    backgroundImage: {
        src: 'https://localhost/static/landingpage/7.jpg',
        alt: 'Latar belakang teknologi blockchain dengan visualisasi keamanan data'
    }
}));

const dialogData = reactive({
    title: 'Verifikasi Ijazah'
});

// Form state
const formData = reactive({
    nisn: ''
});

const formErrors = reactive({
    nisn: ''
});

// Component state
const showDialog = ref(false);
const isLoading = ref(false);
const isSubmitting = ref(false);
const dialogContent = ref(null);

// Computed
const isFormValid = computed(() => {
    return formData.nisn.length === 10 && /^[0-9]{10}$/.test(formData.nisn) && !formErrors.nisn;
});

const openVerificationDialog = async () => {
    isLoading.value = true;

    // Simulate loading delay
    await new Promise((resolve) => setTimeout(resolve, 500));

    showDialog.value = true;
    isLoading.value = false;

    emit('dialog-open');

    // Focus management
    await nextTick();
    if (dialogContent.value) {
        const firstInput = dialogContent.value.querySelector('input');
        if (firstInput) {
            firstInput.focus();
        }
    }
};

const closeDialog = () => {
    showDialog.value = false;
    formData.nisn = '';
    formErrors.nisn = '';
    emit('dialog-close');
};

const validateNisn = () => {
    formErrors.nisn = '';

    if (!formData.nisn) {
        formErrors.nisn = 'NISN wajib diisi';
        return false;
    }

    if (formData.nisn.length !== 10) {
        formErrors.nisn = 'NISN harus terdiri dari 10 digit';
        return false;
    }

    if (!/^[0-9]{10}$/.test(formData.nisn)) {
        formErrors.nisn = 'NISN hanya boleh berisi angka';
        return false;
    }

    return true;
};

const handleVerification = async () => {
    if (!validateNisn()) {
        return;
    }

    isSubmitting.value = true;

    try {
        // Simulate API call
        await new Promise((resolve) => setTimeout(resolve, 2000));

        emit('verification-submit', {
            nisn: formData.nisn
        });

        closeDialog();
    } catch (error) {
        console.error('Verification error:', error);
        formErrors.nisn = 'Terjadi kesalahan saat verifikasi. Silakan coba lagi.';
    } finally {
        isSubmitting.value = false;
    }
};

const onBackgroundLoad = (event) => {
    emit('background-load', event);
};

const onBackgroundError = (event) => {
    console.error('Failed to load hero background:', event);
};

// Keyboard event handlers
const handleKeydown = (event) => {
    if (event.key === 'Escape' && showDialog.value) {
        closeDialog();
    }
};
</script>

<template>
    <div class="hero-wrapper">
        <!-- Navigation Bar -->
        <!-- <nav class="navigation-bar" role="navigation" aria-label="Main navigation">
            <div class="nav-container">
                <TopbarWidget />
            </div>
        </nav> -->

        <!-- Hero Section -->
        <section class="hero-section" role="banner" aria-labelledby="hero-title">
            <!-- Background Image with Lazy Loading -->
            <!-- <LazyImage :src="heroData.backgroundImage.src" :alt="heroData.backgroundImage.alt" container-class="hero-overlay" image-class="hero-background-image" :show-spinner="false" @load="onBackgroundLoad" @error="onBackgroundError" /> -->
            <LazyImage
                :src="heroData.backgroundImage.src"
                :alt="heroData.backgroundImage.alt"
                :use-native-lazy="false"
                :show-spinner="false"
                container-class="hero-overlay"
                image-class="hero-background-image opacity-0"
                loaded-class="opacity-70 transition-opacity duration-500"
            />
            <!-- Hero Content -->
            <div class="hero-content">
                <h1 id="hero-title" class="hero-headline">
                    {{ heroData.headline }}
                </h1>
                <p class="hero-subtitle">
                    {{ heroData.subtitle }}
                </p>
                <button class="hero-cta" :disabled="isLoading" :aria-describedby="isLoading ? 'loading-description' : null" @click="openVerificationDialog">
                    <span v-if="!isLoading">{{ heroData.ctaText }}</span>
                    <span v-else class="loading-content">
                        <svg class="loading-spinner" viewBox="0 0 24 24">
                            <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" opacity="0.25" />
                            <path fill="currentColor" opacity="0.75" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                        </svg>
                        Memuat...
                    </span>
                </button>
                <div v-if="isLoading" id="loading-description" class="sr-only">Sedang memuat dialog verifikasi</div>
            </div>

            <!-- Decorative Pattern -->
            <div class="blockchain-pattern" aria-hidden="true"></div>
        </section>

        <!-- Verification Dialog -->
        <Teleport to="body">
            <div v-if="showDialog" class="dialog-overlay" role="dialog" aria-modal="true" aria-labelledby="dialog-title" aria-describedby="dialog-description" @click="closeDialog">
                <div ref="dialogContent" class="dialog-content" @click.stop>
                    <header class="dialog-header">
                        <h2 id="dialog-title" class="dialog-title">
                            {{ dialogData.title }}
                        </h2>
                        <button class="dialog-close" aria-label="Tutup dialog" @click="closeDialog">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18" />
                                <line x1="6" y1="6" x2="18" y2="18" />
                            </svg>
                        </button>
                    </header>

                    <div id="dialog-description" class="dialog-body">
                        <form class="verification-form" @submit.prevent="handleVerification">
                            <div class="form-group">
                                <label for="nisn-input" class="form-label"> NISN (Nomor Induk Siswa Nasional) </label>
                                <input
                                    id="nisn-input"
                                    v-model="formData.nisn"
                                    type="text"
                                    class="form-input"
                                    :class="{ error: formErrors.nisn }"
                                    placeholder="Masukkan NISN Anda"
                                    maxlength="10"
                                    pattern="[0-9]{10}"
                                    :disabled="isSubmitting"
                                    aria-describedby="nisn-error nisn-help"
                                    required
                                    @input="validateNisn"
                                    @blur="validateNisn"
                                />
                                <div id="nisn-help" class="form-help">NISN terdiri dari 10 digit angka</div>
                                <div v-if="formErrors.nisn" id="nisn-error" class="form-error" role="alert">
                                    {{ formErrors.nisn }}
                                </div>
                            </div>

                            <div class="form-actions">
                                <button type="button" class="btn-secondary" :disabled="isSubmitting" @click="closeDialog">Batal</button>
                                <button type="submit" class="btn-primary" :disabled="!isFormValid || isSubmitting">
                                    <span v-if="!isSubmitting">Verifikasi</span>
                                    <span v-else class="loading-content">
                                        <svg class="loading-spinner" viewBox="0 0 24 24">
                                            <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" opacity="0.25" />
                                            <path fill="currentColor" opacity="0.75" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                                        </svg>
                                        Memverifikasi...
                                    </span>
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </Teleport>
    </div>
</template>

<style scoped>
.hero-wrapper {
    position: relative;
    min-height: 100vh;
}

.navigation-bar {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    z-index: 50;
    background: transparent;
}

.nav-container {
    padding: 1.5rem 1.5rem;
    margin: 0 auto;
    max-width: 1200px;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.hero-section {
    position: relative;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
    background: linear-gradient(135deg, #f0f7ff 0%, #e1edff 100%);
}

.hero-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
}

.hero-background-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    opacity: 0.7;
}

.blockchain-pattern {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100px;
    background: linear-gradient(90deg, transparent, rgba(30, 58, 138, 0.1), transparent);
    z-index: 2;
}

.hero-content {
    position: relative;
    z-index: 10;
    max-width: 800px;
    text-align: center;
    padding: 2rem;
    background: rgba(255, 255, 255, 0.95);
    border-radius: 16px;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.hero-headline {
    color: #1e3a8a;
    font-size: 2.5rem;
    font-weight: 700;
    line-height: 1.2;
    margin-bottom: 1.5rem;
    font-family: 'Poppins', sans-serif;
}

.hero-subtitle {
    color: #4b5563;
    font-size: 1.25rem;
    line-height: 1.6;
    margin-bottom: 2.5rem;
    max-width: 700px;
    margin-left: auto;
    margin-right: auto;
    font-family: 'Inter', sans-serif;
}

.hero-cta {
    background: linear-gradient(135deg, #d85305 0%, #0a2fd4 100%);
    color: white;
    font-weight: 600;
    padding: 1rem 2.5rem;
    border: none;
    border-radius: 50px;
    font-size: 1.1rem;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
    font-family: 'Inter', sans-serif;
    min-width: 200px;
    min-height: 56px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
}

.hero-cta:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.25);
}

.hero-cta:disabled {
    opacity: 0.7;
    cursor: not-allowed;
}

.loading-content {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.loading-spinner {
    width: 20px;
    height: 20px;
    animation: spin 1s linear infinite;
}

/* Dialog Styles */
.dialog-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    padding: 1rem;
    backdrop-filter: blur(4px);
}

.dialog-content {
    background: white;
    border-radius: 12px;
    box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
    max-width: 500px;
    width: 100%;
    max-height: 90vh;
    overflow-y: auto;
    animation: dialogSlideIn 0.3s ease-out;
}

.dialog-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1.5rem 1.5rem 0;
    border-bottom: 1px solid #e5e7eb;
    margin-bottom: 1.5rem;
}

.dialog-title {
    font-size: 1.5rem;
    font-weight: 600;
    color: #1e3a8a;
    margin: 0;
    font-family: 'Poppins', sans-serif;
}

.dialog-close {
    background: none;
    border: none;
    cursor: pointer;
    padding: 0.5rem;
    border-radius: 6px;
    color: #6b7280;
    transition: all 0.2s ease;
}

.dialog-close:hover {
    background: #f3f4f6;
    color: #374151;
}

.dialog-close svg {
    width: 20px;
    height: 20px;
}

.dialog-body {
    padding: 0 1.5rem 1.5rem;
}

.verification-form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.form-label {
    font-weight: 500;
    color: #374151;
    font-family: 'Inter', sans-serif;
}

.form-input {
    padding: 0.75rem 1rem;
    border: 2px solid #d1d5db;
    border-radius: 8px;
    font-size: 1rem;
    transition: border-color 0.2s ease;
    font-family: 'Inter', sans-serif;
}

.form-input:focus {
    outline: none;
    border-color: #1e3a8a;
    box-shadow: 0 0 0 3px rgba(30, 58, 138, 0.1);
}

.form-input.error {
    border-color: #ef4444;
}

.form-input:disabled {
    background: #f9fafb;
    cursor: not-allowed;
}

.form-help {
    font-size: 0.875rem;
    color: #6b7280;
    font-family: 'Inter', sans-serif;
}

.form-error {
    font-size: 0.875rem;
    color: #ef4444;
    font-family: 'Inter', sans-serif;
}

.form-actions {
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
    margin-top: 1rem;
}

.btn-primary,
.btn-secondary {
    padding: 0.75rem 1.5rem;
    border-radius: 8px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
    font-family: 'Inter', sans-serif;
    min-width: 100px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
}

.btn-primary {
    background: #1e3a8a;
    color: white;
    border: 2px solid #1e3a8a;
}

.btn-primary:hover:not(:disabled) {
    background: #1e40af;
    border-color: #1e40af;
}

.btn-primary:disabled {
    background: #9ca3af;
    border-color: #9ca3af;
    cursor: not-allowed;
}

.btn-secondary {
    background: white;
    color: #374151;
    border: 2px solid #d1d5db;
}

.btn-secondary:hover:not(:disabled) {
    background: #f9fafb;
    border-color: #9ca3af;
}

.btn-secondary:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

/* Screen reader only */
.sr-only {
    position: absolute;
    width: 1px;
    height: 1px;
    padding: 0;
    margin: -1px;
    overflow: hidden;
    clip: rect(0, 0, 0, 0);
    white-space: nowrap;
    border: 0;
}

/* Animations */
@keyframes spin {
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
}

@keyframes dialogSlideIn {
    from {
        opacity: 0;
        transform: translateY(-20px) scale(0.95);
    }
    to {
        opacity: 1;
        transform: translateY(0) scale(1);
    }
}

/* Responsive Design */
@media (max-width: 768px) {
    .hero-headline {
        font-size: 2rem;
    }

    .hero-subtitle {
        font-size: 1.1rem;
    }

    .hero-content {
        margin: 1rem;
        padding: 1.5rem;
    }

    .dialog-content {
        margin: 1rem;
    }

    .form-actions {
        flex-direction: column;
    }
}

@media (max-width: 480px) {
    .hero-headline {
        font-size: 1.75rem;
    }

    .hero-subtitle {
        font-size: 1rem;
    }

    .nav-container {
        padding: 1rem;
    }
}

/* Accessibility improvements */
@media (prefers-reduced-motion: reduce) {
    .hero-cta,
    .dialog-content,
    .loading-spinner {
        transition: none;
        animation: none;
    }
}

/* High contrast mode support */
@media (prefers-contrast: high) {
    .hero-headline {
        color: #000;
    }

    .hero-subtitle {
        color: #000;
    }

    .form-input {
        border-color: #000;
    }

    .btn-primary {
        background: #000;
        border-color: #000;
    }
}
</style>
