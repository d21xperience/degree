<template>
    <div>
        <Card>
            <template #content>
                <SiswaForm @submit="createSiswa" />
            </template>

            <template #footer>
                <Button label="Kembali" severity="secondary" @click="goBack" />
            </template>
        </Card>
    </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router';
import { useSiswaCrud } from './composables/useSiswaCrud';
const router = useRouter();
const route = useRoute();
const id = computed(route.params.id);


const siswa = useSiswaCrud();
// PrimeVue components
// import Button from 'primevue/button';
// import Card from 'primevue/card';

// External reusable form
import SiswaForm from '@/components/sekolah_components/siswa/SiswaForm.vue';

// const { create } = useSiswaCrud();

const createSiswa = async (formData) => {
    // await create(formData);
    // router.push({ name: 'Siswa.list' });
};

const goBack = () => {
    router.push({ name: 'infoSiswa' });
};
</script>

<style scoped></style>
