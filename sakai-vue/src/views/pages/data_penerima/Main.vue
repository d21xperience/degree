<template>
    <div class="">
        <div class="">
            <div class="flex justify-between p-2">
                <div class="text-xl" md:text-2xl>Data Penerima Ijazah</div>
                <!-- <div class="flex items-center space-x-2">
                    <label class="text-gray-500">Platform</label>
                    <button @click="dialogSelectplatforms = true" class="rounded-full bg-slate-300 py-2 px-4 hover:opacity-80">
                        <span v-if="platformsActivate">{{ platformsActivate?.name ?? 'Pilih Platform' }}</span>
                        <i class="ml-2 pi pi-angle-up"></i>
                    </button>
                </div> -->
                <div class="md:flex md:items-center md:space-x-2">
                    <label class="text-slate-500 md:text-base text-sm">Tahun Pelajaran</label>
                    <div>
                        <Select v-model="sekolahService.selectedTahunAjaran" :options="listTahunAjaran" optionLabel="nama" placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" :disabled="isDisabled" />
                    </div>
                </div>
            </div>
        </div>
        <div class="card">
            <RouterView></RouterView>
        </div>

        <Dialog v-model:visible="dialogSelectplatforms" modal header="Pilih Platform" position="top">
            <div>
                <div v-for="platform in platforms" :key="platform.id" class="my-2 flex space-x-1">
                    <input type="radio" :name="platform.id" :id="platform.id" :value="platform.name" v-model="selectedPlatform" />
                    <label :for="platform.id">{{ platform?.name }}</label>
                </div>

                <div class="flex justify-center">
                    <button type="button" class="bg-blue-400 p-2 rounded-md text-white" @click="saveSelection">Simpan</button>
                </div>
            </div>
        </Dialog>
    </div>
</template>

<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { computed, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
const store = useStore();
const route = useRoute();
// Dialog
import Dialog from 'primevue/dialog';
import Select from 'primevue/select';

// const { selectedTahunAjaran, initSelectedTahunAjaran, listTahunAjaran } = useSekolahService();
const sekolahService = useSekolahService();
const platformsActivate = ref({});
const platforms = ref(null);

const dialogSelectplatforms = ref(false);
const selectedPlatform = ref();
const fetchPlatforms = async () => {
    payload.schemaname = store.getters['sekolahService/getTabeltenant']?.schemaname;

    const res = await store.dispatch('scService/fetchBCPlatform', payload);
    // console.log(res)
    platforms.value = res.bcPlatform;
    platforms.value.filter((platform) => {
        if (platform.active) {
            selectedPlatform.value = platform.name;
            platformsActivate.value = platform;
        }
    });
};
let payload = {};
// Fungsi untuk mengubah platform yang aktif
const setPlatform = async () => {
    if (!platformsActivate.value) return;

    payload.bc_platform = { ...platformsActivate.value }; // Pastikan objek tidak reaktif
    // console.log("📡 Payload yang dikirim:", JSON.stringify(payload));

    try {
        // console.log(payload.bc_platform)
        store.commit('scService/setBCPlatformSelected', payload.bc_platform);

        const res = await store.dispatch('scService/setBCPlatform', payload);
        console.log('✅ Response backend:', res);
    } catch (error) {
        // console.error("❌ Error mengirim ke backend:", error);
    }
};
// Fungsi untuk memilih platform baru
const saveSelection = async () => {
    dialogSelectplatforms.value = false;
    // Set semua platform ke `false`
    platforms.value.forEach((platform) => (platform.active = false));
    // Temukan platform yang dipilih dan aktifkan
    platformsActivate.value = platforms.value.find((p) => p.name === selectedPlatform.value) || null;
    if (platformsActivate.value) {
        platformsActivate.value.active = true;
    }

    // Update ke backend
    await setPlatform();
};
// Set diaktive current Blockchain network
const platformDiactive = async () => {
    if (!platformsActivate.value) return;

    // Set semua platform ke `false`
    platforms.value.forEach((platform) => (platform.active = false));

    // Pastikan platform yang aktif diubah menjadi false
    platformsActivate.value.active = false;

    console.log('🚀 Sebelum mengirim ke backend:', JSON.stringify(platformsActivate.value));

    // Kirim perubahan ke backend

    await setPlatform();

    // Pastikan platform ter-reset setelah update sukses
    platformsActivate.value = {};
    selectedPlatform.value = '';
};
const listTahunAjaran = ref([]);
// const selectedTahunAjaran = ref();
onMounted(async () => {
    await sekolahService.fetchTahunAjaran();
    listTahunAjaran.value = sekolahService.listTahunAjaran.value._rawValue;
    sekolahService.selectedTahunAjaran.value = sekolahService.initSelectedTahunAjaran.value;
    // // await fetchPlatforms();
    // await fetchSemester();
}); 
// const semester = ref();
// const fetchSemester = async () => {
//     try {
//         let cek = await store.getters['sekolahService/getTahunAjaran'];
//         if (!cek || cek.length === 0) {
//             semester.value = await store.getters['sekolahService/getSemester'];
//             if (!semester.value) {
//                 semester.value = await store.dispatch('sekolahService/fetchSemester');
//             }
//             cek = getTahunAjaran(semester.value);
//         }
//         tahunAjaranOptions.value = cek;
//         store.commit('sekolahService/SET_TABELTAHUNAJARAN', cek);

//         // Ambil tahun ajaran terbaru berdasarkan tahun terbesar
//         selectedTahunAjaran.value = store.getters['sekolahService/getSelectedTahunAjaran'];
//         // console.log(selectedTahunAjaran.value)
//         if (!selectedTahunAjaran) {
//             selectedTahunAjaran.value = tahunAjaranOptions.value.reduce((latest, current) => (current.tahunAjaranId > latest.tahunAjaranId ? current : latest));
//         }
//     } catch (error) {
//         console.log(error);
//     }
// };

const getTahunAjaran = (semesterArray) => {
    const unique = new Set();
    // console.log(semesterArray)
    return semesterArray
        .filter((item) => {
            if (!unique.has(item.tahunAjaranId)) {
                unique.add(item.tahunAjaranId);
                return true;
            }
            return false;
        })
        .map((item) => ({
            label: item.tahunAjaranId,
            value: item.tahunAjaranId + '2'
        }));
};

// watch(selectedTahunAjaran, (val) => {
//     // tetapkan tahun ajaran yang dipilih
//     store.commit('sekolahService/SET_SELECTEDTAHUNAJARAN', val);
// });
const isDisabled = computed(() => route.meta.disableSelect);
</script>

<style lang="scss" scoped></style>
