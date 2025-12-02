<script setup>
import TahunAjaranComponent from '@/components/sekolah_components/TahunAjaranComponent.vue';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
// import LoadingOverlay from '@/components/LoadingOverlay.vue';
// import { isLoading } from '@/router';
import { computed, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
const { selectedSemester } = useSemester();
// const { fetchTabelTenant } = useTableTenant();
const route = useRoute();
// =====================================
const isDisabled = computed(() => route.meta.disableSelect);
const namaRoute = computed(() => route.meta.namaRoute);

// const tabelTenant = ref(null);

// ==============================
// onMounted(async () => {
//     // tabelTenant.value = await fetchTabelTenant();
// });

const isCompleted = ref(false);
const listDataPersiapan = reactive([
    {
        id: 0,
        keterangan: 'Semester',
        link: 'infoSekolah'
    },
    {
        id: 1,
        keterangan: 'Kompetensi keahlian dilayani',
        link: 'infoSekolah'
    }
]);
</script>

<template>
    <div>
        <div v-if="isCompleted || route.name == 'infoSekolah'">
            <div class="flex justify-between items-center mb-2">
                <div class="text-2xl font-semibold">
                    Data <span v-show="namaRoute">{{ `${namaRoute}` }}</span>
                </div>
                <div class="md:flex md:items-center">
                    <!-- <div class="min-w-32">Tahun Pelajaran</div> -->
                    <TahunAjaranComponent v-model="selectedSemester" :is-disabled="isDisabled" />
                </div>
            </div>
            <div class="card">
                <RouterView />
            </div>
        </div>
        <div v-else>
            <div class="flex flex-col items-center justify-center p-6 bg-blue-50 border border-blue-200 rounded-lg shadow-sm text-center max-w-md mx-auto">
                <div class="mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-blue-500 mx-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                    </svg>
                </div>
                <p class="text-gray-700">Silakan lengkapi data di bawah ini terlebih dahulu pada menu <strong>Data Sekolah.</strong></p>
                <div class="flex">
                    <ol class="list-decimal list-inside">
                        <li v-for="(value, index) in listDataPersiapan" :key="index" class="mb-2">
                            <router-link :to="{ name: value.link }" class="font-medium rounded-md hover:text-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition">{{ value.keterangan }}</router-link>
                        </li>
                    </ol>
                </div>
            </div>
        </div>
    </div>
</template>
