<template>
    <div class="flex justify-between items-center mb-2">
        <div class="text-2xl font-semibold">
            Data <span v-show="namaRoute">{{ `${namaRoute} - ` }}</span> Dapodik
        </div>
        <div class="md:flex md:items-center">
            <div class="min-w-32">Tahun Pelajaran</div>
            <Select v-model="selectedSemester" :options="listSemester" optionLabel="namaSemester" placeholder="Tahun Pelajaran" class="w-full" :disabled="isDisabled" />
        </div>
    </div>
    <div class="card">
        <RouterView />
    </div>
</template>

<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import Select from 'primevue/select';
import { computed, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
const { fetchSemester, initSelectedSemester, selectedSemester, listSemester } = useSekolahService();

const route = useRoute();
const store = useStore();

// =====================================
const isDisabled = computed(() => route.meta.disableSelect);
const namaRoute = computed(() => route.meta.namaRoute);

// ==============================
onMounted(async () => {
    selectedSemester.value = initSelectedSemester.value;
    fetchSemester();
    fetchTabelTenant();
});
// const dataConnected = ref(true)

// ==================================
// =======DATA SEKOLAH=============
const tabelTenant = ref(null);
const fetchTabelTenant = async () => {
    try {
        tabelTenant.value = store.getters['sekolahService/getTabeltenant'];
        // console.log(tabelTenant.value)
        if (tabelTenant.value == null) {
            await store.dispatch('sekolahService/fetchTabeltenant');
            tabelTenant.value = store.getters['sekolahService/getTabeltenant'];
        }
    } catch (error) {}
};
</script>
