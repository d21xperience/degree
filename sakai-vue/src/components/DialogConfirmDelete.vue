<template>
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span>{{ message }}</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Ya" icon="pi pi-check" text @click="confirm" />
        </template>
    </Dialog>
</template>

<script setup>
import { computed, defineEmits, defineProps, ref, watch } from 'vue';

const props = defineProps({
    visible: Boolean,
    message: {
        type: String,
        default: 'Apakah data akan dihapus?'
    }
});

const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});
const emit = defineEmits(['update:visible', 'confirm', 'closeDialog']);

const localVisible = ref(props.modelValue);

// Sinkronisasi jika parent mengubah nilai modelValue
watch(
    () => props.modelValue,
    (newVal) => {
        localVisible.value = newVal;
    }
);

// // Emit ke parent jika dialog ditutup dari dalam
// watch(localVisible, (newVal) => {
//     emit('update:modelValue', newVal);
// });

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
