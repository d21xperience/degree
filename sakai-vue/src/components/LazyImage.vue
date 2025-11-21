<script setup>
import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

// Props
const props = defineProps({
    src: {
        type: String,
        required: true
    },
    alt: {
        type: String,
        default: ''
    },
    containerClass: {
        type: String,
        default: 'relative overflow-hidden'
    },
    imageClass: {
        type: String,
        default: 'w-full h-full object-cover transition-opacity duration-300'
    },
    placeholderClass: {
        type: String,
        default: 'absolute inset-0 flex items-center justify-center bg-gray-100'
    },
    errorClass: {
        type: String,
        default: 'absolute inset-0 flex flex-col items-center justify-center bg-gray-50'
    },
    containerStyle: {
        type: Object,
        default: () => ({})
    },
    imageStyle: {
        type: Object,
        default: () => ({})
    },
    placeholderStyle: {
        type: Object,
        default: () => ({})
    },
    errorStyle: {
        type: Object,
        default: () => ({})
    },
    // Intersection Observer options
    rootMargin: {
        type: String,
        default: '50px'
    },
    threshold: {
        type: Number,
        default: 0.1
    },
    // Loading options
    showSpinner: {
        type: Boolean,
        default: true
    },
    nativeLoading: {
        type: Boolean,
        default: false
    },
    errorMessage: {
        type: String,
        default: 'Gagal memuat gambar'
    }
});

// Emits
const emit = defineEmits(['load', 'error', 'visible']);

// Reactive state
const imageContainer = ref(null);
const shouldLoad = ref(false);
const isLoaded = ref(false);
const hasError = ref(false);
const observer = ref(null);

// Computed
const imageSrc = computed(() => {
    // Handle different image sources
    if (props.src.startsWith('https://localhost/')) {
        // Replace localhost with actual domain or use placeholder
        return props.src.replace('https://localhost/', '/api/');
    }
    return props.src;
});

// Methods
const createObserver = () => {
    if (!window.IntersectionObserver) {
        // Fallback for browsers without Intersection Observer
        shouldLoad.value = true;
        return;
    }

    observer.value = new IntersectionObserver(
        (entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    shouldLoad.value = true;
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
};

const startObserving = () => {
    if (imageContainer.value && observer.value) {
        observer.value.observe(imageContainer.value);
    }
};

const stopObserving = () => {
    if (observer.value) {
        observer.value.disconnect();
    }
};

const onImageLoad = (event) => {
    isLoaded.value = true;
    hasError.value = false;
    emit('load', event);
};

const onImageError = (event) => {
    hasError.value = true;
    isLoaded.value = false;
    emit('error', event);
};

// Lifecycle
onMounted(() => {
    createObserver();
    startObserving();
});

onUnmounted(() => {
    stopObserving();
});

// Watch for src changes
watch(
    () => props.src,
    () => {
        isLoaded.value = false;
        hasError.value = false;
        if (!shouldLoad.value) {
            startObserving();
        }
    }
);
</script>

<template>
    <div ref="imageContainer" :class="containerClass" :style="containerStyle">
        <img v-if="shouldLoad" :src="imageSrc" :alt="alt" :class="imageClass" :style="imageStyle" :loading="nativeLoading ? 'lazy' : 'eager'" @load="onImageLoad" @error="onImageError" />

        <!-- Placeholder while loading -->
        <div v-if="!isLoaded && shouldLoad" :class="placeholderClass" :style="placeholderStyle">
            <div v-if="showSpinner" class="loading-spinner">
                <svg class="animate-spin h-8 w-8 text-gray-400" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                </svg>
            </div>
            <div v-else class="placeholder-content">
                <svg class="h-12 w-12 text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M4 4h16v16H4V4zm2 2v12h12V6H6zm2 2h8v8H8V8zm2 2v4h4v-4h-4z" />
                </svg>
            </div>
        </div>

        <!-- Error state -->
        <div v-if="hasError" :class="errorClass" :style="errorStyle">
            <svg class="h-12 w-12 text-red-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z" />
            </svg>
            <p class="text-sm text-red-500 mt-2">{{ errorMessage }}</p>
        </div>
    </div>
</template>

<style scoped>
.loading-spinner {
    @apply flex items-center justify-center;
}

.placeholder-content {
    @apply flex items-center justify-center;
}

.animate-spin {
    animation: spin 1s linear infinite;
}

@keyframes spin {
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
}
</style>
