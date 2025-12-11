<template>
    <div class="p-4">
        <Card>
            <template #title> Edit Kelas </template>

            <template #content>
                <div v-if="loading" class="text-center py-10">
                    <ProgressSpinner />
                </div>

                <div v-else>
                    <KelasForm :initial-data="Kelas" @submit="updateKelas" />
                </div>
            </template>

            <template #footer>
                <Button label="Kembali" severity="secondary" @click="goBack" />
            </template>
        </Card>
    </div>
</template>

<script setup>
import KelasForm from '@/components/sekolah_components/kelas/KelasForm.vue';
import { useKelas } from '@/composables/sekolah_composable/useKelas';
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

// PrimeVue components
const kelasId = computed(() => route.params.id);
console.log('kelas Id', kelasId.value);
// Custom composable CRUD

// External component (form)

const { getById, update } = useKelas();

const Kelas = ref(null);
const loading = ref(true);

onMounted(async () => {
    const id = route.params.id;
    Kelas.value = await getById(id);
    loading.value = false;
});

const updateKelas = async (data) => {
    await update(route.params.id, data);
    router.push({ name: 'Kelas.list' });
};

const goBack = () => {
    router.push({ name: 'infoKelas' });
};
</script>

<style scoped></style>
