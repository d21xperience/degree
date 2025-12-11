<template>
    <div>
        <!-- {{ cekTahunajaranId || '' }} -->
        <div v-if="isCompleted || route.name == 'infoSekolah' || route.name == 'semesterInfo'">
            <div class="flex justify-between items-center mb-2">
                <div class="text-2xl font-semibold">
                    Data <span v-show="namaRoute">{{ `${namaRoute}` }}</span>
                </div>
                <div class="md:flex md:items-center">
                    <!-- <div class="min-w-32">Tahun Pelajaran</div> -->
                    <TahunAjaranComponent v-model:model-value="cekTahunajaranId" :is-disabled="isDisabled" />
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
                        <router-link :to="{ name: 'infoSekolah', query: { lengkapi: 2 } }" class="font-medium rounded-md hover:text-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition"
                            >Kompetensi keahlian dilayani</router-link
                        >
                    </ol>
                </div>
            </div>
        </div>

        <!-- Global Dialog (cukup 1x di App.vue) -->
        <DialogStatus v-model="dialog.visible" :type="dialog.type" :title="dialog.title" :message="dialog.message" :button-label="dialog.buttonLabel" @confirm="dialog.onConfirm" />
    </div>
</template>

<script setup>
import TahunAjaranComponent from '@/components/sekolah_components/TahunAjaranComponent.vue';
import { useKategoriSekolah } from '@/composables/sekolah_composable/useKategoriSekolah';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useDialogStatus } from '@/composables/useDialogStatus';
import store from '@/store';
import { computed, provide, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();

// Mendapatkan url parameter
// =====================================
const isDisabled = computed(() => route.meta.disableSelect);
const namaRoute = computed(() => route.meta.namaRoute);
// =====================================
const cekTahunajaranId = ref(null);
const isCompleted = computed(() => store.getters['sekolahService/getIsKategoriSekolahCompleted']);
const user = computed(() => store.getters['authService/user']);
const schemaname = computed(() => {
    return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
});

const { currentSekolah } = useSekolah({ schemaname: schemaname.value, autoload: true });
const { fetchKategoriSekolah } = useKategoriSekolah(schemaname.value);

// Mengirimkan data ke seluruh keluarga
// const currentSekolah = computed(() => store.getters['sekolahService/getSekolah']);
provide('sekolahSlugProvider', user.value.sekolahSlug);
provide('sekolahProvider', currentSekolah);
provide('schemanameProvider', schemaname.value);
provide(
    'tahunAjaranProvider',
    computed(() => cekTahunajaranId.value)
);
// console.log('currentSekolah!!', currentSekolah.value);
watch(cekTahunajaranId, (newVal) => {
    if (newVal) {
        fetchKategoriSekolah(newVal.label);
    }
});
const { dialogState } = useDialogStatus();
const dialog = computed({
    get: () => dialogState.value,
    set: (val) => {
        dialogState.value.visible = val.visible;
    }
});
</script>
