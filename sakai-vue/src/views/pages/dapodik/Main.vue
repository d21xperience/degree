<template>
    <div class="flex justify-between items-center mb-2">
        <div class="text-2xl font-semibold">
            Data <span v-show="namaRoute">{{ `${namaRoute}` }}</span>
        </div>
        <div class="md:flex md:items-center">
            <!-- <div class="min-w-32">Tahun Pelajaran</div> -->
            <SemesterComponent v-model="selectedSemester" :isDisabled="isDisabled" /> 
        </div>
    </div>
    <div class="card">
        <RouterView />
    </div>
</template>

<script setup>
import SemesterComponent from '@/components/sekolah_components/SemesterComponent.vue';
import { useSekolahService } from '@/composables/useSekolahService';
import { computed, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
const {  selectedSemester  } = useSekolahService();

const route = useRoute();
const store = useStore();

// =====================================
const isDisabled = computed(() => route.meta.disableSelect);
const namaRoute = computed(() => route.meta.namaRoute);

// ==============================
onMounted(async () => { 
    fetchTabelTenant();
}); 

// ==================================
// =======DATA SEKOLAH=============
const tabelTenant = ref(null);
const fetchTabelTenant = async () => {
    try {
        tabelTenant.value = store.getters['sekolahService/getTabeltenant']; 
        if (tabelTenant.value == null) {
            await store.dispatch('sekolahService/fetchTabeltenant');
            tabelTenant.value = store.getters['sekolahService/getTabeltenant'];
        }
    } catch (error) {}
};
</script>
