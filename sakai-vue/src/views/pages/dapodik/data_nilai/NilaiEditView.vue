<template>
    <div class="p-4">
        <Card>
            <template #title> Edit Nilai </template>

            <template #content>
                <div v-if="loading" class="text-center py-10">
                    <ProgressSpinner />
                </div>

                <div v-else>
                    <NilaiForm :initial-data="Nilai" @submit="updateNilai" />
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
import Button from 'primevue/button';
import Card from 'primevue/card';
import ProgressSpinner from 'primevue/progressspinner';

// Custom composable CRUD
import { useNilaisCrud } from '@/composables/useNilaisCrud';

// External component (form)
import NilaiForm from '@/components/Nilai/NilaiForm.vue';

const route = useRoute();
const router = useRouter();

const { getById, update } = useNilaisCrud();

const Nilai = ref(null);
const loading = ref(true);

onMounted(async () => {
    const id = route.params.id;
    Nilai.value = await getById(id);
    loading.value = false;
});

const updateNilai = async (data) => {
    await update(route.params.id, data);
    router.push({ name: 'Nilai.list' });
};

const goBack = () => {
    router.push({ name: 'Nilai.list' });
};
</script>

<style scoped></style>
