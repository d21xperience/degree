<template>
    <div class="flex justify-between items-center mb-2">
        <label class="text-xl md:text-2xl"
            >Data <span v-show="namaRoute">{{ `${namaRoute} - ` }}</span> Dapodik</label
        >
        <div class="md:flex md:items-center">
            <label class="min-w-32">Tahun Pelajaran</label>
            <Select v-model="selectedSemester" :options="listSemester" optionLabel="namaSemester" placeholder="Tahun Pelajaran" class="w-full" :disabled="isDisabled" />
        </div>
    </div>
    <div class="card">
        <RouterView />
    </div>
</template>
<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
const { fetchSemester, initSelectedSemester, selectedSemester, listSemester } = useSekolahService();

const namaRoute = computed(() => route.meta.namaRoute);
const isDisabled = computed(() => route.meta.disableSelect);
</script>
