<template>
    <div class="p-4">
        <Card>
            <template #title> Edit Guru </template>

            <template #content>
                <div v-if="loading" class="text-center py-10">
                    <ProgressSpinner />
                </div>

                <div v-else>
                    <GuruForm :initial-data="Guru" @submit="updateGuru" />
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

// PrimeVue components
// import Button from 'primevue/button';
// import Card from 'primevue/card';
// import ProgressSpinner from 'primevue/progressspinner';

// Custom composable CRUD
import { useGurusCrud } from '@/composables/useGurusCrud';

// External component (form)
import GuruForm from '@/components/Guru/GuruForm.vue';

const route = useRoute();
const router = useRouter();

const { getById, update } = useGurusCrud();

const Guru = ref(null);
const loading = ref(true);

onMounted(async () => {
    const id = route.params.id;
    Guru.value = await getById(id);
    loading.value = false;
});

const updateGuru = async (data) => {
    await update(route.params.id, data);
    router.push({ name: 'Guru.list' });
};

const goBack = () => {
    router.push({ name: 'Guru.list' });
};
</script>

<style scoped></style>
