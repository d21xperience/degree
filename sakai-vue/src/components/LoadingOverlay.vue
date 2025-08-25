<script setup>
import ProgressSpinner from 'primevue/progressspinner';

defineProps({
    visible: {
        type: Boolean,
        default: false
    }
});
</script>

<template>
    <!-- Overlay fullscreen (bisa di-disable jika hanya ingin loading lokal) -->
    <div v-if="visible" class="fixed inset-0 z-[1000] flex items-center justify-center bg-black/50 backdrop-blur-sm">
        <div class="flex flex-col items-center justify-center px-10 py-12 gap-6 bg-white/10 backdrop-blur-lg border border-white/20 rounded-3xl shadow-2xl min-w-[320px] text-center">
            <!-- Spinner -->
            <div class="relative">
                <ProgressSpinner :style="{ width: '80px', height: '80px' }" strokeWidth="6" fill="none" animationDuration="1.2s" class="text-white drop-shadow-lg" />
                <div class="absolute inset-0 rounded-full border-2 border-blue-300/30 animate-ping" style="width: 80px; height: 80px"></div>
            </div>

            <!-- Teks -->
            <div class="text-white font-medium text-lg leading-tight drop-shadow">
                <slot>
                    <span>Sedang memuat data...</span>
                </slot>
            </div>

            <!-- Garis animasi -->
            <div class="w-16 h-0.5 bg-white/30 rounded-full">
                <div class="h-full bg-gradient-to-r from-transparent via-white to-transparent rounded-full animate-pulse" style="width: 60%; margin: 0 auto; animation-duration: 1.5s"></div>
            </div>
        </div>
    </div>
</template>
