<template>
    <div class="p-4">
        <Card>
            <template #title> Edit user </template>

            <template #content>
                <div v-if="loading" class="text-center py-10">
                    <ProgressSpinner />
                </div>

                <div v-else>
                    <userForm :initial-data="user" @submit="updateuser" />
                </div>
            </template>

            <template #footer>
                <Button label="Kembali" severity="secondary" @click="goBack" />
            </template>
        </Card>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// Custom composable CRUD
import { useusersCrud } from '@/composables/useusersCrud';

// External component (form)
import userForm from '@/components/user/userForm.vue';

const route = useRoute();
const router = useRouter();

const { getById, update } = useusersCrud();

const user = ref(null);
const loading = ref(true);

onMounted(async () => {
    const id = route.params.id;
    user.value = await getById(id);
    loading.value = false;
});

const updateuser = async (data) => {
    await update(route.params.id, data);
    router.push({ name: 'user.list' });
};

const goBack = () => {
    router.push({ name: 'user.list' });
};
</script>

<style scoped></style>
