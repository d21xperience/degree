<template>
    <!-- Dialog start -->
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Keluar" :modal="true" position="top">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span>Yakin akan keluar?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="isVisible = false" severity="warn" />
            <Button label="Ya" icon="pi pi-check" text @click="onLogout" />
        </template>
    </Dialog>
    <!-- Dialog end -->
</template>
<script setup>
import { computed } from 'vue';

// const dialogSignOut = () => {};
const onLogout = () => {
    emit('confirm')
};
const emit = defineEmits(['update:visible', 'confirm']);
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});
const props = defineProps({
    visible: Boolean
});
</script>
