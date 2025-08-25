<script setup>
import { onMounted, ref } from 'vue';
const displayNumber = ref(0);
const duration = 10000; // in ms
const props = defineProps({
    label: {
        type: String,
        default: 'Stat'
    },
    icon: String,
    targetNumber: {
        type: Number,
        default: 0
    },
    url: String
});
onMounted(() => {
    const startTime = performance.now();
    const update = (currentTime) => {
        const elapsed = currentTime - startTime;
        const progress = Math.min(elapsed / duration, 1);
        displayNumber.value = Math.floor(progress * props.targetNumber);

        if (progress < 1) {
            requestAnimationFrame(update);
        } else {
            displayNumber.value = props.targetNumber;
        }
    };

    requestAnimationFrame(update);
});
</script>

<template>
    <div class="stat-box">
        <div class="stat-number">{{ displayNumber }}+</div>
        <div class="stat-label">{{ props.label }}</div>
    </div>
</template>
<style scoped>
.stat-box {
    text-align: center;
    padding: 1.5rem;
    border-radius: 8px;
    background-color: white;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    min-width: 200px;
}

.stat-number {
    color: #ffd700;
    font-family: 'Poppins', sans-serif;
    font-weight: 700;
    font-size: 2rem;
    margin-bottom: 0.5rem;
}

.stat-label {
    color: #1e3a8a;
    font-family: 'Inter', sans-serif;
    font-size: 0.9rem;
}
</style>
