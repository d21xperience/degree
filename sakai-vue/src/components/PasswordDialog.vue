<template>
    <Dialog v-model:visible="isVisible" position="top">
        <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-xl mb-2">
            Password
        </label>
        <Password 
            id="password1" 
            v-model="password" 
            placeholder="Password" 
            :toggleMask="true" 
            class="mb-4" 
            fluid 
            :feedback="false"
        />

        <template #footer>
            <Button label="Submit" @click="handleSubmit" />
        </template>
    </Dialog>
</template>

<script setup>
import { computed, ref } from 'vue';
import { Dialog, Password, Button } from 'primevue';
const isVisible = computed({
    get: () => props.visible,
    set : (value)=> emit('update:visible', value)
})
const props = defineProps({
    visible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['update:visible', 'submit']);

const password = ref('');

const handleSubmit = () => {
    emit('submit', password.value);
    emit('update:visible', false);
    password.value = '';
};
</script>

