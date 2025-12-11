// src/composables/useIdleDetection.js
import { onMounted, onUnmounted, ref } from 'vue';

export function useIdleDetection({
    timeout = 30 * 60 * 1000, // 30 menit
    warningThreshold = 60 * 1000, // 1 menit sebelum logout
    onIdle,
    onWarning,
    onActivity
} = {}) {
    const isIdle = ref(false);
    const isWarning = ref(false);
    const countdown = ref(0);
    let timeoutId = null;
    let warningId = null;
    let countdownInterval = null;

    const resetTimer = () => {
        // Batalkan semua timer
        if (timeoutId) clearTimeout(timeoutId);
        if (warningId) clearTimeout(warningId);
        if (countdownInterval) clearInterval(countdownInterval);

        isIdle.value = false;
        isWarning.value = false;
        countdown.value = 0;

        // Atur ulang timeout
        timeoutId = setTimeout(() => {
            // Masuk mode warning dulu
            isWarning.value = true;
            countdown.value = warningThreshold / 1000; // detik

            // Countdown mundur
            countdownInterval = setInterval(() => {
                countdown.value -= 1;
                if (countdown.value <= 0) {
                    clearInterval(countdownInterval);
                    onIdle?.();
                }
            }, 1000);

            // Jika tidak ada aktivitas selama `warningThreshold`, logout
            warningId = setTimeout(() => {
                onIdle?.();
            }, warningThreshold);
        }, timeout - warningThreshold);
    };

    const handleActivity = () => {
        if (isWarning.value) {
            // Batalkan logout jika user aktif saat warning
            onActivity?.();
        }
        resetTimer();
    };

    // Event listener untuk aktivitas
    const events = ['mousedown', 'mousemove', 'keypress', 'scroll', 'touchstart'];

    onMounted(() => {
        events.forEach((event) => {
            window.addEventListener(event, handleActivity, { passive: true });
        });
        resetTimer(); // Mulai timer pertama kali
    });

    onUnmounted(() => {
        events.forEach((event) => {
            window.removeEventListener(event, handleActivity);
        });
        if (timeoutId) clearTimeout(timeoutId);
        if (warningId) clearTimeout(warningId);
        if (countdownInterval) clearInterval(countdownInterval);
    });

    return {
        isIdle,
        isWarning,
        countdown,
        reset: resetTimer
    };
}
