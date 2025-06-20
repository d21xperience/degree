<template>
    <div>
        <div class="">
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <!-- <Button icon="pi pi-plus" severity="success" class="mr-2" @click="openNew" v-tooltip.bottom="'Tambah data'" v-show="selectedSemester.semester == 1" /> -->
                        <Button icon="pi pi-pencil" severity="warn" @click="editKelas" :disabled="!selectedKelas || !selectedKelas.length || selectedKelas.length > 1" class="mr-2" v-tooltip.bottom="'Edit data'" />
                        <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected" :disabled="!selectedKelas || !selectedKelas.length" v-tooltip.bottom="'Hapus data'" />
                        <!-- <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2" v-tooltip.bottom="'Upload'" v-show="selectedSemester.semester == 1" /> -->
                        <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2" v-tooltip.bottom="'Download'" />

                        <div v-show="selectedSemester.semester == 2">
                            <Button label="Lulus" severity="help" class="mr-2 text-sm" @click="isDialogKelulusan = true" :disabled="!isLulus || !selectedKelas.length || selectedKelas.length > 1" v-tooltip.bottom="'Pilih tingkat tertinggi'" />
                        </div>
                    </template>
                    <template #end>
                        <div class="flex flex-wrap gap-2 items-center justify-between">
                            <div class="flex">
                                <Select
                                    v-model="filters['tingkatPendidikanId'].value"
                                    :options="tingkatPendidikanOptions"
                                    optionLabel="nama"
                                    optionValue="kode"
                                    placeholder="Tingkat"
                                    class="w-full md:w-48"
                                    checkmark
                                    show-clear
                                    v-show="kelasList.length > 0"
                                />
                            </div>
                        </div>
                    </template>
                </Toolbar>
            </div>

            <DataTable
                ref="dt"
                v-model:selection="selectedKelas"
                stripedRows
                size="small"
                :value="kelasList"
                scrollable
                scrollHeight="400px"
                dataKey="rombonganBelajarId"
                :paginator="true"
                :rows="10"
                :filters="filters"
                tableStyle="min-width: 50rem"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[10, 20, 30]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas"
                class="mt-2"
            >
                <!-- <template #empty> No customers found. </template> -->
                <!-- <template #loading> Loading customers data. Please wait. </template> -->
                <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                <Column field="nmKelas" header="Nama" style="width: 7rem">
                    <template #loading>
                        <div class="flex items-center" :style="{ height: '17px', 'flex-grow': '1', overflow: 'hidden' }">
                            <Skeleton width="40%" height="1rem" />
                        </div>
                    </template>
                </Column>
                <Column field="tingkatPendidikanId" header="Tingkat" sortable>
                    <template #loading>
                        <div class="flex items-center" :style="{ height: '17px', 'flex-grow': '1', overflow: 'hidden' }">
                            <Skeleton width="40%" height="1rem" />
                        </div>
                    </template>
                </Column>
                <Column field="kurikulum.namaKurikulum" header="Kurikulum"></Column>
                <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
                <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                    <Column field="namaJurusanSp" header="Jurusan" sortable></Column>
                </div>
                <Column header="Anggota">
                    <template #body="slotProps">
                        <Button icon="pi pi-bullseye" outlined rounded class="mr-2" @click="dialogAnggotaRombel(slotProps.data)" />
                    </template>
                </Column>
                <Column field="jumlahAnggota" header="Jml."></Column>
            </DataTable>
        </div>

        <DialogImport v-model:visible="dialogImport" template-type="kelas" />
        <DialogConfirmDelete v-model:visible="deleteKelasDialog" message="Apakah kelas tersebut akan dihapus?" @confirm="deleteKelas" @closeDialog="closeDialog" />
        <Dialog v-model:visible="showAnggotaKelas" style="width: 450px; height: max-content" header="Anggota Kelas" close-icon="pi pi-times" maximizable>
            <AnggotaKelas :rombongan-belajar-id="rombonganBelajarId" />
        </Dialog>
        <Dialog v-model:visible="isDialogKelulusan">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                Apakah
                <template v-if="selectedKelas.length == 1"> kelas {{ selectedKelas[0]?.nmKelas }} </template>
                <template v-else> semua kelas </template>
                akan diluluskan?
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="closeDialogKelulusan" />
                <Button label="Ya" icon="pi pi-check" text @click="sendToDns" />
            </template>
        </Dialog>

        <!-- end of import data -->
    </div>
</template>

<script setup>
import { computed, nextTick, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// import DialogImport from '@/components/DialogImport.vue'
import { FilterMatchMode } from '@primevue/core/api';
import Button from 'primevue/button';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import Select from 'primevue/select';
import Toolbar from 'primevue/toolbar';
import { useToast } from 'primevue/usetoast';
// import DialogLoading from '@/components/DialogLoading.vue';
import router from '@/router';

import AnggotaKelas from '@/components/AnggotaKelas.vue';
import Skeleton from 'primevue/skeleton';
// ================================
// composable
// ================================
// import DialogAnggotaKelas from '@/components/dapodik/AnggotaKelas.vue';
import { useSekolahService } from '@/composables/useSekolahService';

const { schemaname, fetchKelas, fetchTingkat, sekolah, addDns } = useSekolahService();
// ================================
const kelasList = ref([]);
const isLoading = ref(false);
const tingkatPendidikanOptions = ref();

const openNew = async () => {
    // console.log(sekolah.value)
    await nextTick();
    router.push({ name: 'inputKelas', params: { sekolah: sekolah.value?.uri } });
};
const closeDialog = () => {
    selectedKelas.value = null;
};
const closeDialogKelulusan = () => {
    isDialogKelulusan.value = false;
    isLulus.value = false;
    selectedKelas.value = null;
};
const exportCSV = () => {
    dt.value.exportCSV();
};

const deleteKelas = () => {
    kelasList.value = kelasList.value.filter((val) => !selectedKelas.value.includes(val));
    // if (selectedKelas.value.length == 1) {
    //     deleteSiswaAktif(selectedKelas.value[0].anggotaRombelId);
    // } else if (selectedKelas.value.length > 1) {
    //     const ids = selectedKelas.value.map((item) => item.anggotaRombelId);
    //     deleteBatchSiswaAktif(ids);
    // }
};
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
watch(selectedSemester, async () => {
    await fetchK();
});
const fetchK = async () => {
    kelasList.value = await fetchKelas();
    console.log(kelasList.value)
    if (kelasList.value.length > 0) {
        tingkatPendidikanOptions.value = await fetchTingkat();
    }
};

const toast = useToast();
const dt = ref();
const deleteKelasDialog = ref(false);
const selectedKelas = ref([]);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkatPendidikanId: { value: null, matchMode: FilterMatchMode.CONTAINS }
});
const submitted = ref(false);

const editKelas = async () => {
    router.push({
        name: 'editKelas',
        params: { sekolah: sekolah.value?.uri },
        query: { kelasId: selectedKelas.value[0]?.rombonganBelajarId.toString() }
    });
};

const confirmDeleteSelected = () => {
    deleteKelasDialog.value = true;
};

const showAnggotaKelas = ref(false);
const rombonganBelajarId = ref();
const dialogAnggotaRombel = (d) => {
    console.log(d);
    selectedKelas.value.push(d);
    showAnggotaKelas.value = true;
    rombonganBelajarId.value = d?.rombonganBelajarId;
};
watch(showAnggotaKelas, (e) => {
    if (!e) {
        selectedKelas.value = [];
    }
});

const bentukPendidikan = ref('smk');

const isLulus = ref(false);
const isNaik = ref(false);
// const selectedKelasLulus = ref();

watch(selectedKelas, (newVal) => {
    if (!newVal) {
        return;
    }
    if (newVal.length === 1) {
        if (newVal[0]?.tingkatPendidikanId == 12) {
            isLulus.value = true;
        }
    } else if (newVal.length > 1) {
        isLulus.value = true;
        const kelas12 = newVal.filter((item) => item.tingkatPendidikanId == 12);
        if (kelas12) {
            isLulus.value = true;
        }
    }
});
const isDialogKenaikan = ref(false);
const isDialogKelulusan = ref(false);
const luluskan = async () => {
    try {
        isDialogKelulusan.value = false;
        const anggotaKelas = selectedKelas.value.flatMap((kelas) => kelas?.anggotaKelas || []);
        const payload = {
            schemaname: schemaname,
            tahun_ajaran_id: `${selectedSemester.value?.tahunAjaranId + 1}`,
            anggota_kelas: anggotaKelas,
            sekolah_id: await store.getters['sekolahService/getSekolah']?.sekolah_id
        };
        const res = await store.dispatch('sekolahService/createProsesIjazah', payload);
        console.log(res);
        if (res) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Data Ijazah berhasil ditambahkan', life: 3000 });
            selectedKelas.value = [];
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Gagal', detail: 'Gagal menambahkan data', life: 3000 });
    }
    isDialogKelulusan.value = false;
};
const dialogImport = ref(false);

const sendToDns = async () => {
    const sekolah = store.getters['sekolahService/getSekolah'];
    const anggotaKelas = selectedKelas.value[0].anggotaKelas.map((item) => ({
        peserta_didik_id: item.pesertaDidikId || '',
        rombongan_belajar_id: item.rombonganBelajarId || '',
        program_keahlian: selectedKelas.value[0].namaJurusanSp || '',
        paket_keahlian: selectedKelas.value[0].namaJurusanSp || '',
        sekolah_id: selectedKelas.value[0].sekolahId || '',
        npsn: sekolah.npsn || '', // jika ada, ambil dari sekolah
        kabupaten_kota: sekolah.kabKota || '',
        provinsi: sekolah.propinsi || '',
        nama: item.pesertaDidik.nmSiswa || '',
        tempat_lahir: item.pesertaDidik.tempatLahir || '',
        tanggal_lahir: item.pesertaDidik.tanggalLahir || '',
        nis: item.pesertaDidik.nis || '',
        nisn: item.pesertaDidik.nisn || '',
        nama_ortu_wali: item.pesertaDidik.nmAyah || '',
        sekolah_penyelenggara_ujian_us: sekolah.nama || '',
        sekolah_penyelenggara_ujian_un: sekolah.nama || '',
        asal_sekolah: sekolah.nama || '',
        nomor_ijazah: '',
        tempat_ijazah: sekolah.kabKota || '',
        tanggal_ijazah: '',
        tahun_ajaran_id: selectedSemester.tahunAjaranId,
        is_complete: false
    }));

    isDialogKelulusan.value = false;
    addDns(anggotaKelas);
};
onMounted(async () => {
    await fetchK();
});
</script>
