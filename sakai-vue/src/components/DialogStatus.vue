<template>
    <Dialog
        v-model:visible="modelValue"
        :header="title"
        modal
        :closable="closable"
        :dismissable-mask="dismissableMask"
        :pt="{
            root: '!rounded-xl !shadow-xl !border-0',
            mask: '!backdrop-blur-sm',
            header: '!border-b !border-surface-200 dark:!border-surface-700 !py-4 !px-6 !font-medium',
            content: '!py-6 !px-6',
            footer: '!pt-4 !px-6 !pb-6'
        }"
        @update:visible="onHide"
    >
        <div class="flex flex-col items-center text-center gap-4">
            <!-- Icon Circle -->
            <div class="flex items-center justify-center w-14 h-14 rounded-full" :class="bgClass">
                <i :class="icon" class="text-xl" :style="{ color: iconColor }"></i>
            </div>

            <!-- Content -->
            <div>
                <h3 class="text-lg font-semibold text-900 dark:text-50">
                    {{ title }}
                </h3>
                <p v-if="message" class="mt-2 text-600 dark:text-400 leading-relaxed">
                    {{ message }}
                </p>
            </div>
        </div>

        <template #footer>
            <Button :label="buttonLabel" :severity="buttonSeverity" class="w-full md:w-auto" @click="onConfirm" />
        </template>
    </Dialog>
</template>

<script setup>
import { computed } from 'vue';

// ✅ Gunakan defineModel (Vue 3.4+) — lebih modern dari v-model + emit
const modelValue = defineModel({
    type: Boolean,
    default: false
});

const props = defineProps({
    type: {
        type: String,
        default: 'success',
        validator: (v) => ['success', 'error', 'warning', 'info'].includes(v)
    },
    title: {
        type: String,
        default: ''
    },
    message: {
        type: String,
        default: ''
    },
    buttonLabel: {
        type: String,
        default: 'OK'
    },
    buttonSeverity: {
        type: String,
        default: ''
    },
    closable: {
        type: Boolean,
        default: true
    },
    dismissableMask: {
        type: Boolean,
        default: true
    }
});

const emit = defineEmits(['confirm', 'hide']);

// 🔹 Konfigurasi tema sesuai PrimeVue 4 + Tailwind
const config = computed(() => {
    const base = {
        success: {
            icon: 'pi pi-check',
            bgClass: 'bg-green-100 dark:bg-green-900/20',
            iconColor: 'var(--p-success-color)',
            defaultTitle: 'Success',
            severity: 'success'
        },
        error: {
            icon: 'pi pi-times',
            bgClass: 'bg-red-100 dark:bg-red-900/20',
            iconColor: 'var(--p-error-color)',
            defaultTitle: 'Error',
            severity: 'danger'
        },
        warning: {
            icon: 'pi pi-exclamation-circle',
            bgClass: 'bg-orange-100 dark:bg-orange-900/20',
            iconColor: 'var(--p-warning-color)',
            defaultTitle: 'Warning',
            severity: 'warning'
        },
        info: {
            icon: 'pi pi-info-circle',
            bgClass: 'bg-blue-100 dark:bg-blue-900/20',
            iconColor: 'var(--p-info-color)',
            defaultTitle: 'Information',
            severity: 'secondary'
        }
    };
    return base[props.type] || base.info;
});

const icon = computed(() => config.value.icon);
const bgClass = computed(() => config.value.bgClass);
const iconColor = computed(() => config.value.iconColor);
const resolvedTitle = computed(() => props.title || config.value.defaultTitle);
const buttonSeverity = computed(() => props.buttonSeverity || config.value.severity);

const onConfirm = () => {
    emit('confirm');
    modelValue.value = false;
};

const onHide = () => {
    emit('hide');
};
</script>
