<template>
    <div class="col-span-12 lg:col-span-6 xl:col-span-4">
        <div class="card mb-0 p-0">
            <div class="flex justify-between">
                <div class="flex items-center">
                    <h2 class="text-5xl">{{ props.targetNumber ? displayNumber : 0 }}</h2>
                </div>
                <div class="flex flex-col items-center">
                    <div class="flex items-center justify-center bg-blue-100 dark:bg-blue-400/10 rounded-border mb-2" style="width: 4rem; height: 4rem">
                        <i class="text-blue-500 !text-5xl" :class="[props.icon]"></i>
                    </div>
                    <span class="block text-muted-color font-medium mb-2">{{ props.label }}</span>
                </div>
            </div>
            <div v-if="props.url" class="text-center cursor-pointer">
                <router-link :to="{ name: props.url }" class="hover:text-slate-900">Details</router-link>
            </div>
        </div>
    </div>
</template>
<script setup>
import { onMounted, ref } from 'vue';

const displayNumber = ref(0);
const duration = 2000; // in ms
const props = defineProps({
    label: String,
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
