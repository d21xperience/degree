<script setup>
import { computed } from 'vue';

// Props
const props = defineProps({
    modelValue: { type: Object, required: true },
    loading: { type: Boolean, default: false },
    submitLabel: { type: String, default: 'Save' }
});

// Emit
const emit = defineEmits(['update:modelValue', 'submit']);

// Form model
const form = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val)
});

// Example dropdown options (jurusan)
const majors = [
    { label: 'Computer Science', value: 'cs' },
    { label: 'Information System', value: 'is' },
    { label: 'Accounting', value: 'acc' }
];

const submitForm = () => {
    emit('submit');
};
</script>

<template>
    <form class="space-y-4 p-4 bg-white rounded-lg shadow-md" @submit.prevent="submitForm">
        <!-- NAME -->
        <div class="flex flex-col gap-1">
            <label class="font-medium">Name</label>
            <InputText v-model="form.name" placeholder="Enter student name" class="w-full" />
        </div>

        <!-- EMAIL -->
        <div class="flex flex-col gap-1">
            <label class="font-medium">Email</label>
            <InputText v-model="form.email" type="email" placeholder="Student email" class="w-full" />
        </div>

        <!-- MAJOR -->
        <div class="flex flex-col gap-1">
            <label class="font-medium">Major</label>
            <Dropdown v-model="form.major" :options="majors" option-label="label" option-value="value" placeholder="Select major" class="w-full" />
        </div>

        <!-- ACTIONS -->
        <div class="flex justify-end mt-4">
            <Button type="submit" :loading="loading" :label="submitLabel" class="px-4 py-2" />
        </div>
    </form>
</template>

<style scoped></style>
