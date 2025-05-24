<script setup>
import InputText from 'primevue/inputtext';
import Select from 'primevue/select';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();

const kelasId = route.query.kelasId;
const toast = useToast();

import AnggotaKelas from '@/components/AnggotaKelas.vue';
import router from '@/router';
import { useStore } from 'vuex';
const store = useStore();

const isEdit = ref(false);

// const sekolah = computed(() => store.getters['sekolahService/getSekolah']);
const selectedTingkat = ref(null);
const tingkatPendidikanOptions = ref([]);
const submitted = ref(false);
// Model Kelas
const rombel = ref({
    rombonganBelajarId: '',
    sekolahId: '',
    semesterId: '',
    jurusanId: '',
    ptkId: '',
    nmKelas: '',
    tingkatPendidikanId: '',
    jenisRombel: '',
    namaJurusanSp: '',
    jurusan_sp_id: '',
    kurikulumId: '',
    tingkatPendidikan: null,
    kurikulum: null,
    jurusan: null
});

const fetchTingkat = async () => {
    const payload = {
        jenjang_pendidikan_id: sekolah.value?.jenjangPendidikanId
    };
    const response = await store.dispatch('sekolahService/fetchTingkatPendidikan', payload);
    tingkatPendidikanOptions.value = response;
};

// ===================================
// KURIKULUM
// ===================================
const selectedKurikulum = ref(null);
const kurikulumList = ref([]);
// const fetchKurikulum = async () => {
//     try {
//         const payload = {
//             // jurusan_id: selectedJurusan.value?.jurusanId
//         };
//         const response = await store.dispatch('sekolahService/fetchKurikulum', payload);
//         kurikulumList.value = response;
//     } catch (error) {
//         console.error(error);
//     }
// };
// ===================================
// ================================
// composable
// ================================
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
const { fetchKelas, kelasList, fetchGuruTerdaftar, guruTerdaftarList, sekolah } = useSekolahService(schemaname, selectedSemester);
const { fetchKurikulum, kurikulumLoading, kurikulumSearch, kurikulumOptions } = useFormOptions();
// ================================

// ===================================
// JURUSAN
// ===================================
const selectedJurusan = ref(null);
const jurusanList = ref([]);
const filteredJurusan = ref([]);
const fetchJurusan = async () => {
    const payload = {};
    if (sekolah.value?.jenjangPendidikanId == 6) {
        switch (sekolah.value?.bentukPendidikanId) {
            case 15:
                payload.param = 'untuk_smk';
                break;
            case 13:
                payload.param = 'untuk_sma';
                break;

            default:
                break;
        }
    }
    payload.jenjang_pendidikan_id = sekolah.value?.jenjangPendidikanId;
    const response = await store.dispatch('sekolahService/fetchJurusan', payload);
    jurusanList.value = response;
};
const searchJurusan = (event) => {
    setTimeout(() => {
        let query = event.query.toLowerCase();

        filteredJurusan.value = jurusanList.value.filter((jurusan) => jurusan.namaJurusan.toLowerCase().includes(query));
    }, 250);
};
const handleKeydown = (event) => {
    if (event.key === ' ') {
        selectedJurusan.value += ' '; // Menambahkan spasi ke query
    }
};
const kurikulumHandleKeydown = (event) => {
    if (event.key === ' ') {
        selectedKurikulum.value += ' '; // Menambahkan spasi ke query
    }
};

watch(selectedJurusan, async (newVal) => {
    // console.log("tes Jursauan")
    // if (typeof newVal === 'object') {
    //     if (newVal) {
    //         // await fetchKurikulum();
    //         const kur = kurikulumList._rawValue;
    //         selectedKurikulum.value = kur.find((item) => item.jurusanId === newVal.jurusanId);
    //     }
    // }
});
// ===================================

const selectedGuruTerdaftar = ref(null);

const angotaKelas = computed(() => JSON.parse(localStorage.getItem('unsavedPesertaDidik')) || null);
const simpanKelas = async () => {
    try {
        rombel.value.tingkatPendidikanId = selectedTingkat.value?.tingkatPendidikanId;
        rombel.value.semesterId = await store.getters['sekolahService/getSelectedSemester']?.semesterId;
        rombel.value.jurusanId = selectedJurusan.value?.jurusanId;
        rombel.value.ptkId = selectedGuruTerdaftar.value.ptk?.ptkId;
        rombel.value.jenisRombel = Number(rombel.value.jenisRombel) || 1;
        rombel.value.namaJurusanSp = selectedJurusan.value?.namaJurusan;
        rombel.value.kurikulumId = selectedKurikulum.value?.kurikulumId;
        rombel.value.sekolahId = sekolah.value?.sekolah_id;
        let result = null;
        const payload = {
            schemaname: schemaname.value,
            kelas: rombel.value
            // anggota_kelas: [anggotaKelas.value]
        };
        if (isEdit.value) {
            // Cek apakah anggota kelas diisi
            if (angotaKelas.value) {
                payload.anggota_kelas = angotaKelas.value;
            }
            result = await store.dispatch('sekolahService/editKelas', payload);
            if (result) {
                toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil diperbaharui', life: 3000 });
            }
        } else {
            result = await store.dispatch('sekolahService/createKelas', payload);
            if (result) {
                toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
                submitted.value = false;
                Object.keys(rombel.value).forEach((key) => {
                    rombel.value[key] = '';
                });
                selectedJurusan.value = null;
                selectedKurikulum.value = null;
                selectedGuruTerdaftar.value = {};
                selectedTingkat.value = '';
            }
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Gagal', detail: 'Data gagal disimpan', life: 3000 });

        console.error(error);
    }
};

const batal = async () => {
    // await nextTick();
    router.push({ name: 'infoKelas' });
};

onMounted(async () => {
    // console.log(sekolah.value);
    await fetchTingkat();
    await fetchKurikulum();
    await fetchJurusan();
    // await fetchGuruTerdaftar();
    // if (kelasId) {
    //     isEdit.value = true;
    //     await fetchKelas(kelasId);
    //     if (kelasList.value && kelasList.value.length > 0) {
    //         const kelas = kelasList.value[0];
    //         rombel.value = {
    //             rombonganBelajarId: kelas.rombonganBelajarId || '',
    //             sekolahId: kelas.sekolahId || '',
    //             semesterId: kelas.semesterId || '',
    //             jurusanId: kelas.jurusanId || '',
    //             ptkId: kelas.ptkId || '',
    //             nmKelas: kelas.nmKelas || '',
    //             tingkatPendidikanId: kelas.tingkatPendidikanId || '',
    //             jenisRombel: kelas.jenisRombel || '',
    //             namaJurusanSp: kelas.namaJurusan || '',
    //             jurusan_sp_id: kelas.namaJurusanSpId || '',
    //             kurikulumId: kelas.kurikulumId || ''
    //             // tingkatPendidikan: null,
    //             // kurikulum: null,
    //             // jurusan: null,
    //         };
    //         selectedTingkat.value = tingkatPendidikanOptions.value.find((item) => item.tingkatPendidikanId === kelas.tingkatPendidikanId);
    //         selectedJurusan.value = jurusanList.value.find((item) => item.jurusanId === kelas.jurusanId);
    //         selectedGuruTerdaftar.value = guruTerdaftarList.value.find((item) => item.ptk.ptkId === kelas.ptk.ptkId);
    //     }
    // }
});

const addGuru = () => {
    alert('tambah guru');
    // ==================================

    // ==================================
};
const isLoadingSave = ref(false);
const isLoadingCancel = ref(false);
const submitForm = ref(false);
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center mb-1">
            <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Form Edit Kelas' : 'Form Tambah Kelas' }}</h1>
            <Button icon="pi pi-times" severity="danger" @click="batal" />
        </div>
        <section class="mb-6">
            <h2 class="text-xl font-semibold mb-4">Informasi Kelas</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div class="flex space-x-1">
                    <div class="w-full">
                        <label class="block text-gray-700" for="nmKelas">Nama Kelas</label>
                        <InputText v-model="rombel.nmKelas" fluid name="nmKelas" id="nmKelas" placeholder="contoh: x tbsm a" :invalid="submitted && !rombel.nmKelas" />
                        <small v-if="submitted && !rombel.nmKelas" class="text-red-500">Nama harus diisi.</small>
                    </div>
                    <div class="w-full md:w-40">
                        <label class="block text-gray-700">Tingkat</label>
                        <Select v-model="selectedTingkat" :options="tingkatPendidikanOptions" placeholder="Pilih tingkat" optionLabel="nama" class="w-full" :invalid="submitted && !selectedTingkat" />
                        <small v-if="submitted && !selectedTingkat" class="text-red-500">Tingkat harus diisi.</small>
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700">Kurikulum</label>
                    <div class="relative">
                        <!-- <Select v-model="selectedKurikulum" :options="kurikulumList" placeholder="Pilih Kurikulum" optionLabel="namaKurikulum" class="w-full" :invalid="submitted && !selectedKurikulum" /> -->
                        <AutoComplete
                            v-model="selectedKurikulum"
                            :suggestions="kurikulumOptions"
                            option-label="namaKurikulum"
                            placeholder="Pilih kurikulum"
                            dropdown
                            @complete="kurikulumSearch"
                            :loading="kurikulumLoading"
                            @keydown.space.prevent="kurikulumHandleKeydown"
                            fluid
                        />
                        <small v-if="submitted && !selectedKurikulum" class="text-red-500">Kurikulum harus diisi.</small>
                    </div>
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div v-show="sekolah.bentukPendidikanId == 13 || sekolah.bentukPendidikanId == 15">
                    <label class="block text-gray-700">Jurusan</label>
                    <div class="relative">
                        <AutoComplete
                            v-model="selectedJurusan"
                            :suggestions="filteredJurusan"
                            optionLabel="namaJurusan"
                            @complete="searchJurusan"
                            @keydown.space.prevent="handleKeydown"
                            placeholder="Cari Jurusan..."
                            class="w-full"
                            fluid
                            dropdown
                            :invalid="submitted && !selectedJurusan"
                            :disabled="!selectedKurikulum"
                        />
                        <small v-if="submitted && !selectedJurusan" class="text-red-500">Jurusan harus diisi.</small>
                    </div>
                </div>
                <div class="">
                    <label class="block text-gray-700">Wali kelas</label>
                    <div class="relative">
                        <Select v-model="selectedGuruTerdaftar" :options="guruTerdaftarList" placeholder="Pilih Wali kelas" optionLabel="ptk.nama" class="w-full" :invalid="submitted && !selectedGuruTerdaftar">
                            <template #footer>
                                <div class="p-3">
                                    <Button label="Tambah Guru" fluid severity="secondary" text size="small" icon="pi pi-plus" @click="addGuru" />
                                </div>
                            </template>
                        </Select>
                        <small v-if="submitted && !selectedGuruTerdaftar" class="text-red-500">Wali kelas harus diisi.</small>
                    </div>
                </div>
            </div>
        </section>

        <Toast />
        <!-- Daftar anggota rombel -->
        <!-- <div v-show="isEdit"> -->
        <h2 class="text-xl font-semibold mb-4">Anggota Kelas</h2>
        <!-- Anggota Kelas -->
        <AnggotaKelas :rombongan-belajar-id="kelasId" :is-edit="isEdit" />
        <!-- End of Anggota Kelas -->
        <!-- </div> -->

        <div class="flex justify-end mt-8 space-x-4">
            <Button label="Simpan" severity="success" @click="submitForm" :loading="isLoadingSave" class="min-w-28" />
            <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" :loading="isLoadingCancel" />
        </div>
    </div>
</template>
