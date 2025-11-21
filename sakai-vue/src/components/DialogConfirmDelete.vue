<script setup>
import { computed } from 'vue';

const emit = defineEmits(['update:visible', 'confirm', 'closeDialog']);
const props = defineProps({
    visible: Boolean,
    message: {
        type: String,
        default: 'data'
    },
    judul: {
        type: String,
        default: 'Hapus data'
    }
});

const isVisible = computed({
    get: () => props.visible,
    set: (value) => {
        emit('update:visible', value);
        emit('closeDialog');
    }
});

// Function untuk menutup dialog
const closeDialog = () => {
    emit('closeDialog');
    isVisible.value = false;
};

const confirm = () => {
    emit('confirm');
    isVisible.value = false;
};
</script>

<template>
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" :header="judul" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl"></i>
            <span>Apakah <span v-html="message"></span></span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Ya" icon="pi pi-check" text @click="confirm" />
        </template>
    </Dialog>
</template>
