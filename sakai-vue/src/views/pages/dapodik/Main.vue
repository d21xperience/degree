<script setup>
import SemesterComponent from '@/components/sekolah_components/SemesterComponent.vue';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
// import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';
// import LoadingOverlay from '@/components/LoadingOverlay.vue';
// import { isLoading } from '@/router';
import { computed } from 'vue';
import { useRoute } from 'vue-router';
const { selectedSemester } = useSemester();
// const { fetchTabelTenant } = useTableTenant();
const route = useRoute();
// =====================================
const isDisabled = computed(() => route.meta.disableSelect);
const namaRoute = computed(() => route.meta.namaRoute);

// const tabelTenant = ref(null);

// // ==============================
// onMounted(async () => {
//     tabelTenant.value = await fetchTabelTenant();
// });
</script>

<template>
    <div>
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
        <!-- <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay> -->
    </div>
</template>
