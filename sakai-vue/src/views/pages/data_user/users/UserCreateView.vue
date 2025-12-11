<template>
    <div class="p-4">
        <Card>
            <template #title> Tambah user </template>

            <template #content>
                <userForm @submit="createuser" />
            </template>

            <template #footer>
                <Button label="Kembali" severity="secondary" @click="goBack" />
            </template>
        </Card>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

// PrimeVue components
import Button from 'primevue/button';
import Card from 'primevue/card';

// Custom composable CRUD
import { useuserCrud } from '@/composables/sekolah/user/Crud';

// External reusable form
import userForm from '@/components/user/userForm.vue';

const router = useRouter();

const { create } = useuserCrud();

const createuser = async (formData) => {
    await create(formData);
    router.push({ name: 'user.list' });
};

const goBack = () => {
    router.push({ name: 'user.list' });
};
</script>

<style scoped></style>
