<!-- src/components/LazyImage.vue -->
<script setup>
/**
 * LazyImage - Komponen gambar dengan lazy loading & error handling robust
 *
 * Fitur:
 * - Lazy loading via Intersection Observer (fallback ke eager jika IO tidak didukung)
 * - Penanganan error + retry opsional
 * - Aksesibilitas penuh (a11y)
 * - Optimasi performa dan UX
 *
 * Contoh penggunaan:
 * <LazyImage
 *   src="/static/landingpage/7.jpg"
 *   alt="Latar belakang teknologi blockchain"
 *   class="hero-bg"
 *   @load="handleLoad"
 *   @error="handleError"
 * />
 */

import { computed, nextTick, onBeforeUnmount, onMounted, ref, watch } from 'vue';

const props = defineProps({
    /** URL sumber gambar */
    src: {
        type: String,
        required: true
    },
    /** Teks alternatif (wajib untuk a11y) */
    alt: {
        type: String,
        default: ''
    },
    /** Kelas untuk container */
    containerClass: {
        type: String,
        default: 'relative overflow-hidden'
    },
    /** Kelas untuk elemen <img> */
    imageClass: {
        type: String,
        default: 'w-full h-full object-cover transition-opacity duration-300'
    },
    /** Kelas saat loading */
    loadingClass: {
        type: String,
        default: 'opacity-0'
    },
    /** Kelas saat berhasil dimuat */
    loadedClass: {
        type: String,
        default: 'opacity-100'
    },
    /** Kelas saat error */
    errorClass: {
        type: String,
        default: 'flex flex-col items-center justify-center bg-gray-50 text-gray-500'
    },
    /** Apakah tampilkan spinner */
    showSpinner: {
        type: Boolean,
        default: true
    },
    /** Gunakan native `loading="lazy"` (lebih ringan, direkomendasikan) */
    useNativeLazy: {
        type: Boolean,
        default: true
    },
    /** Threshold Intersection Observer (0–1) */
    threshold: {
        type: Number,
        default: 0.01 // cukup 1% terlihat → load
    },
    /** Margin untuk trigger IO */
    rootMargin: {
        type: String,
        default: '50px'
    },
    /** Pesan error default */
    errorMessage: {
        type: String,
        default: 'Gambar tidak dapat dimuat'
    },
    /** Jumlah retry saat error (0 = matikan retry) */
    retryCount: {
        type: Number,
        default: 1
    }
});

const emit = defineEmits({
    load: (e) => e instanceof Event,
    error: (e) => e instanceof Event,
    visible: () => true
});

// State
const containerRef = ref(null);
const isLoading = ref(true);
const isLoaded = ref(false);
const hasError = ref(false);
const observer = ref(null);
const retryAttempt = ref(0);

// Gunakan src langsung — **tidak ada manipulasi hardcoded**
const safeSrc = computed(() => props.src);

// Intersection Observer
const setupObserver = () => {
    if (typeof window === 'undefined' || !window.IntersectionObserver) {
        // Jika IO tidak tersedia (SSR, browser lama), load langsung
        isLoading.value = false;
        return;
    }

    observer.value = new IntersectionObserver(
        (entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    isLoading.value = false;
                    emit('visible');
                    observer.value?.unobserve(entry.target);
                }
            });
        },
        {
            rootMargin: props.rootMargin,
            threshold: props.threshold
        }
    );

    if (containerRef.value) {
        observer.value.observe(containerRef.value);
    }
};

// Preload manual (opsional, hanya jika tidak pakai native lazy)
const preloadImage = () => {
    if (!safeSrc.value) return;

    const img = new Image();
    let aborted = false;

    const cleanup = () => {
        img.onload = null;
        img.onerror = null;
        if (observer.value) observer.value.disconnect();
    };

    img.onload = () => {
        if (aborted) return;
        cleanup();
        isLoaded.value = true;
        hasError.value = false;
        emit('load', { target: img });
    };

    img.onerror = (e) => {
        if (aborted) return;
        cleanup();
        handleImageError(e);
    };

    img.src = safeSrc.value;

    // Batalkan saat unmount
    onBeforeUnmount(() => {
        aborted = true;
        cleanup();
    });
};

const handleImageError = (event) => {
    retryAttempt.value++;
    if (retryAttempt.value <= props.retryCount) {
        // Coba ulang setelah delay
        setTimeout(() => {
            if (import.meta.env.DEV) {
                console.warn(`[LazyImage] Retry ${retryAttempt.value}/${props.retryCount} for ${safeSrc.value}`);
            }
            preloadImage();
        }, 300 * retryAttempt.value);
        return;
    }

    hasError.value = true;
    isLoaded.value = false;
    emit('error', event);

    if (import.meta.env.DEV) {
        console.error('[LazyImage] Gagal memuat gambar:', safeSrc.value, event);
    }
};

// Lifecycle
onMounted(() => {
    if (props.useNativeLazy) {
        // Gunakan native `loading="lazy"` — lebih ringan & didukung browser modern
        isLoading.value = false; // langsung render <img>, biarkan browser handle lazy
    } else {
        setupObserver();
    }
});

onBeforeUnmount(() => {
    if (observer.value) {
        observer.value.disconnect();
    }
});

// Watch perubahan src
watch(
    () => props.src,
    async (newSrc, oldSrc) => {
        if (newSrc === oldSrc) return;

        // Reset state
        isLoaded.value = false;
        hasError.value = false;
        isLoading.value = props.useNativeLazy ? false : true;
        retryAttempt.value = 0;

        await nextTick();

        if (!props.useNativeLazy) {
            setupObserver();
        }
    }
);
</script>

<template>
    <div
        ref="containerRef"
        :class="containerClass"
        :aria-busy="isLoading || (hasError && retryAttempt <= retryCount)"
        :aria-live="hasError ? 'polite' : 'off'"
        role="img"
        :aria-label="alt || undefined"
        :aria-labelledby="!alt ? 'lazy-img-title' : undefined"
        :style="{
            '--lazy-border-radius': '0px',
            '--lazy-shadow': 'none',
            ...$attrs.style
        }"
    >
        <!-- Judul tersembunyi untuk a11y jika alt kosong -->
        <span v-if="!alt" id="lazy-img-title" class="sr-only">Gambar konten</span>

        <!-- Gambar utama -->
        <img
            v-if="!hasError"
            :src="safeSrc"
            :alt="alt"
            :class="[imageClass, isLoaded ? loadedClass : loadingClass]"
            :loading="useNativeLazy ? 'lazy' : 'eager'"
            @load="
                isLoaded = true;
                isLoading = false;
                emit('load', $event);
            "
            @error="handleImageError"
        />

        <!-- Loading spinner (hanya jika bukan native lazy atau butuh kontrol penuh) -->
        <div v-if="!isLoaded && !hasError && showSpinner && !useNativeLazy" class="absolute inset-0 flex items-center justify-center bg-gray-100" aria-hidden="true">
            <svg class="animate-spin h-8 w-8 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" aria-hidden="true">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
            </svg>
            <span class="sr-only">Memuat gambar...</span>
        </div>

        <!-- Error state -->
        <div v-if="hasError" :class="['absolute inset-0', errorClass]" role="alert">
            <svg class="h-12 w-12 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M4 12a8 8 0 018-8v0a8 8 0 018 8v0a8 8 0 01-8 8v0a8 8 0 01-8-8v0z" />
            </svg>
            <p class="mt-2 text-sm font-medium">
                {{ errorMessage }}
            </p>
            <p v-if="retryAttempt > 0" class="mt-1 text-xs text-gray-500">Telah dicoba {{ retryAttempt }}/{{ retryCount + 1 }} kali</p>
            <button
                v-if="retryCount > 0"
                type="button"
                class="mt-2 px-3 py-1 text-sm bg-red-50 text-red-600 rounded hover:bg-red-100 focus:outline-none focus:ring-2 focus:ring-red-500"
                :disabled="retryAttempt > retryCount"
                @click="
                    () => {
                        retryAttempt = 0;
                        handleImageError({});
                    }
                "
            >
                Coba Lagi
            </button>
        </div>
    </div>
</template>

<style scoped>
/* Animasi smooth dengan reduced motion support */
img {
    transition: opacity 0.3s ease-in-out;
}

@media (prefers-reduced-motion: reduce) {
    img {
        transition: none;
    }
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

/* Animasi spinner */
.animate-spin {
    animation: spin 1s linear infinite;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}
</style>
