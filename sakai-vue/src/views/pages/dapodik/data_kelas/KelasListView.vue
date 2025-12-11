<script setup>
import TingkatComponent from '@/components/sekolah_components/TingkatComponent.vue';
import { useDns } from '@/composables/sekolah_composable/useDns';
import { useKelas } from '@/composables/sekolah_composable/useKelas';
import router from '@/router';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue';
import { computed, inject, ref, watch } from 'vue';

const sekolah = inject('sekolahProvider');
const { fetchKelas } = useKelas();
const { addDns } = useDns();
const sekolahSlug = inject('sekolahSlugProvider');
// const slug = computed(() => store.getters['authService/getUser'])?.sekolahSlug;
// const { selectedSemester } = useSemester();
const kelasList = ref([]);
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
// const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
// watch(selectedSemester, async () => {
//     await initial();
// });
const toast = useToast();
const initial = async (ObjSemester) => {
    try {
        kelasList.value = await fetchKelas(ObjSemester);
    } catch (error) {
        console.log('Gagal initialisasi kelas: ', error);
        toast.add({ severity: 'error', summary: 'Gagal initialisasi', detail: `${error}`, life: 20000 });
    }
};

const dt = ref();
const deleteKelasDialog = ref(false);
const selectedKelas = ref([]);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkatPendidikanId: { value: null, matchMode: FilterMatchMode.CONTAINS }
});
// const submitted = ref(false);

const editKelas = async () => {
    console.log(sekolahSlug);
    router.push({
        name: 'editKelas',
        params: { sekolah: sekolahSlug, id: selectedKelas.value[0]?.rombonganBelajarId.toString() }
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

// const isNaik = ref(false);
// const selectedKelasLulus = ref();
const isLulus = ref(false);
const isKejuruan = computed(() => ['SMK', 'MAK'].includes(sekolah.value.bentukPendidikanStr ?? false));

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
// const isDialogKenaikan = ref(false);
const isDialogKelulusan = ref(false);
// const luluskan = async () => {
//     try {
//         isDialogKelulusan.value = false;
//         const anggotaKelas = selectedKelas.value.flatMap((kelas) => kelas?.anggotaKelas || []);
//         const payload = {
//             schemaname: schemaname,
//             tahun_ajaran_id: `${selectedSemester.value?.tahunAjaranId + 1}`,
//             anggota_kelas: anggotaKelas,
//             sekolah_id: store.getters['sekolahService/getSekolah']?.sekolah_id
//         };
//         const res = await store.dispatch('sekolahService/createProsesIjazah', payload);
//         // console.log(res);
//         if (res) {
//             toast.add({ severity: 'success', summary: 'Successful', detail: 'Data Ijazah berhasil ditambahkan', life: 3000 });
//             selectedKelas.value = [];
//         }
//     } catch (error) {
//         toast.add({ severity: 'error', summary: 'Gagal', detail: 'Gagal menambahkan data', life: 3000 });
//     }
//     isDialogKelulusan.value = false;
// };
// const dialogImport = ref(false);

const sendToDns = async () => {
    const anggotaKelas = selectedKelas.value[0].anggotaKelas.map((item) => ({
        peserta_didik_id: item.pesertaDidikId || '',
        rombongan_belajar_id: item.rombonganBelajarId || '',
        program_keahlian: selectedKelas.value[0].namaJurusanSp || '',
        paket_keahlian: selectedKelas.value[0].namaJurusanSp || '',
        sekolah_id: sekolah.value.sekolah.sekolah_id || '',
        npsn: sekolah.value.sekolah.npsn || '', // jika ada, ambil dari sekolah
        kabupaten_kota: sekolah.value.sekolah.kabKota || '',
        provinsi: sekolah.value.sekolah.propinsi || '',
        nama: item.pesertaDidik.nmSiswa || '',
        tempat_lahir: item.pesertaDidik.tempatLahir || '',
        tanggal_lahir: item.pesertaDidik.tanggalLahir || '',
        jenis_kelamin: item.pesertaDidik.jenisKelamin,
        nis: item.pesertaDidik.nis || '',
        nisn: item.pesertaDidik.nisn || '',
        nama_ortu_wali: item.pesertaDidik.nmAyah || '',
        sekolah_penyelenggara_ujian_us: sekolah.value.sekolah.nama || '',
        sekolah_penyelenggara_ujian_un: sekolah.value.sekolah.nama || '',
        asal_sekolah: sekolah.value.sekolah.nama || '',
        nomor_ijazah: '',
        tempat_ijazah: sekolah.value.sekolah.kabKota || '',
        tanggal_ijazah: '',
        tahun_ajaran_id: `${selectedSemester.value.tahunAjaranId}`,
        is_complete: false
    }));
    // console.log(anggotaKelas);
    // return;
    isDialogKelulusan.value = false;
    addDns(anggotaKelas);
};

// const selectedTahunAjaran = computed(() => store.getters['semesterService/getSelectedTahunAjaran']);
const selectedSemester = inject('selectedSemesterProvider');

watch(selectedSemester, async (newVal) => {
    console.log('----[selectedSemester]----', newVal);
    if (newVal) {
        await initial(newVal);
    } else {
        kelasList.value = [];
    }
});
</script>

<template>
    <div class="">
        <Toolbar>
            <template #start>
                <div v-show="kelasList?.length > 0">
                    <!-- <Button icon="pi pi-plus" severity="success" class="mr-2" @click="openNew" v-tooltip.bottom="'Tambah data'" v-show="selectedSemester.semester == 1" /> -->
                    <Button v-tooltip.bottom="'Edit data'" icon="pi pi-pencil" severity="warn" :disabled="!selectedKelas || !selectedKelas.length || selectedKelas.length > 1" class="mr-2" @click="editKelas" />
                    <Button v-tooltip.bottom="'Hapus data'" icon="pi pi-trash" severity="danger" class="mr-2" :disabled="!selectedKelas || !selectedKelas.length" @click="confirmDeleteSelected" />
                    <!-- <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2" v-tooltip.bottom="'Upload'" v-show="selectedSemester.semester == 1" /> -->
                    <Button v-tooltip.bottom="'Download'" icon="pi pi-download" severity="help" class="mr-2" @click="exportCSV($event)" />

                    <Button
                        v-show="selectedSemester?.semester == 2"
                        v-tooltip.bottom="'Pilih tingkat tertinggi'"
                        label="Lulus"
                        severity="help"
                        class="mr-2 text-sm"
                        :disabled="!isLulus || !selectedKelas.length || selectedKelas.length > 1"
                        @click="isDialogKelulusan = true"
                    />
                </div>
            </template>
            <template #end>
                <div class="flex flex-wrap gap-2 items-center justify-between">
                    <div class="flex flex-wrap items-center space-x-2">
                        <label>Tingkat</label>
                        <TingkatComponent v-model:model-value="filters['tingkatPendidikanId'].value" :jenjang-pendidikan-id="sekolah?.sekolah.jenjangPendidikanId" :is-disabled="!selectedSemester || kelasList.length === 0" />
                    </div>
                    <!-- <div>
                        <Button v-tooltip.bottom="'Refresh'" icon="pi pi-refresh" severity="help" class="mr-2 text-lg" @click="initial" />
                    </div> -->
                </div>
            </template>
        </Toolbar>

        <div v-if="kelasList.length === 0" class="flex justify-center h-32 items-center"><h5>Silahkan pilih semseter</h5></div>
        <DataTable
            v-else
            ref="dt"
            v-model:selection="selectedKelas"
            striped-rows
            size="small"
            :value="kelasList"
            scrollable
            scroll-height="400px"
            data-key="rombonganBelajarId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            table-style="min-width: 50rem"
            paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rows-per-page-options="[10, 20, 30]"
            current-page-report-template="Showing {first} to {last} of {totalRecords} kelas"
            class="mt-2"
        >
            <Column selection-mode="multiple" style="width: 3rem" :exportable="false" />
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
            <Column field="kurikulum.namaKurikulum" header="Kurikulum" />
            <div v-if="isKejuruan">
                <Column field="namaJurusanSp" header="Jurusan" sortable />
            </div>
            <Column header="Anggota">
                <template #body="slotProps">
                    <Button icon="pi pi-bullseye" outlined rounded class="mr-2" @click="dialogAnggotaRombel(slotProps.data)" />
                </template>
            </Column>
            <Column field="jumlahAnggota" header="Jml." />
        </DataTable>

        <!-- <DialogImport v-model:visible="dialogImport" template-type="kelas" /> -->
        <DialogConfirmDelete v-model:visible="deleteKelasDialog" message="Apakah kelas tersebut akan dihapus?" @confirm="deleteKelas" @close-dialog="closeDialog" />
        <Dialog v-model:visible="showAnggotaKelas" style="width: 450px; height: max-content" header="Anggota Kelas" close-icon="pi pi-times" maximizable>
            <AnggotaKelas :rombongan-belajar-id="rombonganBelajarId" />
        </Dialog>
        <Dialog v-model:visible="isDialogKelulusan">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl"></i>
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
