<template>
    <div class="">
        <Toolbar>
            <template #start>
                <div class="w-56">
                    <KelasComponent class="mr-2" v-model="kelasSelected" />
                </div>
                <div v-show="!!kelasSelected" class="ml-1 flex space-x-1">
                    <!-- <Button icon="pi pi-plus" severity="success" class="text-lg" @click="openNew" v-tooltip.bottom="'Tambah Siswa'" :loading="isOpenNew" /> -->
                    <!-- <Button icon="pi pi-pencil" severity="warn" @click="editNilai" :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 2" class="" v-tooltip.bottom="'Edit nilai'" :loading="loadingEdit" /> -->
                    <!-- <Button icon="pi pi-trash" severity="danger" class="text-lg" @click="deleteSiswaDialog = true" :disabled="!selectedSiswa || !selectedSiswa.length" v-tooltip.bottom="'Hapus Nilai'" /> -->
                    <Button icon="pi pi-upload" severity="warn" @click="exportCSV($event)" class="text-sm" v-tooltip.bottom="'Import Nilai'" />
                </div>
            </template>
            <template #end>
                <div v-show="!!kelasSelected" class="mr-1">
                    <IconField>
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." :disabled="!dataNilaiSiswa.length > 0" />
                    </IconField>
                </div>
                <Button icon="pi pi-refresh" severity="help" class="mr-2 text-lg" @click="loadSiswaAktif" v-tooltip.bottom="'Refresh'" v-show="!!kelasSelected" />
            </template>
        </Toolbar>
        <div v-if="!kelasSelected" class="flex justify-center h-32 items-center">
            <h5>Silahkan pilih kelas terlebih dahulu</h5>
        </div>
        <DataTable
            v-else
            ref="dt"
            v-model:expandedRows="expandedRows"
            stripedRows
            size="small"
            :value="dataNilaiSiswa"
            @rowExpand="onRowExpand"
            @rowCollapse="onRowCollapse"
            dataKey="pesertaDidikId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} siswa"
        >
            <template #header>
                <div class="flex flex-wrap justify-end gap-2">
                    <Button text icon="pi pi-plus" label="Expand All" @click="expandAll" />
                    <Button text icon="pi pi-minus" label="Collapse All" @click="collapseAll" />
                </div>
            </template>
            <template #empty>
                <p class="flex justify-center text-xl">Data Tidak ditemukan.</p>
            </template>
            <Column expander style="width: 5rem" />
            <Column field="nmSiswa" header="Nama" sortable></Column>
            <Column field="tingkatPendidikanId" header="Tingkat"></Column>
            <Column field="nmKelas" header="Nama Kelas"></Column>
            <Column field="" header="Aksi">
                <template #body="{ data }">
                    <Button icon="pi pi-trash" outlined rounded class="mr-2" @click="editNilai(data)" />
                    <!-- <Button icon="pi pi-trash" outlined rounded severity="danger"
                                    @click="confirmdeleteMapel(data)" /> -->
                </template>
            </Column>
            <template #expansion="slotProps">
                <div class="p-4">
                    <DataTable  :value="siswa.nilai" >
                        <Column field="mataPelajaran" header="Mata Pelajaran" class="text-slate-500" />
                        <!-- Kolom Semester Dinamis -->
                        <Column v-for="n in totalSemesters" :key="`semester${n}`" :field="`semester${n}`" :header="`${n}`">
                            <!-- Opsional: Tambahkan warna berdasarkan nilai -->
                            <template #body="{ data }">
                                <span
                                    :class="{
                                        'text-green-600 font-medium': data[`semester${n}`] >= 85,
                                        'text-yellow-600': data[`semester${n}`] >= 75 && data[`semester${n}`] < 85,
                                        'text-red-600': data[`semester${n}`] < 75
                                    }"
                                >
                                    {{ data[`semester${n}`] }}
                                </span>
                            </template>
                        </Column>
                        <template #empty> <p class="text-xl flex justify-center font-bold text-red-500">Nilai tidak ditemukan.</p> </template>
                        <Column field="" header="Edit">
                            <template #body="{ data }">
                                <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editNilai(data)" />
                                <!-- <Button icon="pi pi-trash" outlined rounded severity="danger"
                                    @click="confirmdeleteMapel(data)" /> -->
                            </template>
                        </Column>
                    </DataTable>
                </div>
            </template>
        </DataTable>

        <!-- import data -->
        <!-- DIALOG IMPORT -->
        <DialogImport v-model:visible="isDialogImport" @save="saveImport" @cancel="cancelImport" template-type="siswa" />

        <!-- end of import data -->
        <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay>
    </div>
</template>

<script setup>
import LoadingOverlay from '@/components/LoadingOverlay.vue';
import KelasComponent from '@/components/sekolah_components/KelasComponent.vue';
import { useSekolahService } from '@/composables/useSekolahService';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref, watch } from 'vue';
const sekolahService = useSekolahService();
const pembelajaran = ref({});
const dataNilaiSiswa = ref([]);
const kelasSelected = ref(null);
const expandedRows = ref();
// watch(selectedSemester, async (newVal, oldVal) => {
//     // console.log(newVal)
//     dataNilaiSiswa.value = await fetchNilaiSiswa();
// });

const isLoading = ref(false);
const toast = useToast();
const dt = ref();
const mapelDialog = ref(false);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    jurusan: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkat: { value: null, matchMode: FilterMatchMode.GREATER_THAN }
});
const kelas = ref({});
const editNilai = (mapel) => {
    kelas.value = { ...mapel };
    mapelDialog.value = true;
    pembelajaran.value.rombongan_belajar_id = kelas.value.rombonganBelajarId;
    pembelajaran.value.semester_id = kelas.value.semesterId;
};

const exportCSV = () => {
    isLoading.value = true;
    // alert("hello")
    // dt.value.exportCSV();
};

const isDialogImport = ref(false);
const saveImport = (e) => {
    // console.log("Data disimpan:", e);
    isDialogImport.value = false;
};

const cancelImport = () => {
    console.log('Import dibatalkan');
    isDialogImport.value = false;
};
// ===========================================

const onRowExpand = (event) => {
    toast.add({ severity: 'info', summary: 'Product Expanded', detail: event, life: 3000 });
    // Ambil data nilai akhir

    // console.log(event)
};
const onRowCollapse = (event) => {
    // toast.add({ severity: 'success', summary: 'Product Collapsed', detail: event, life: 3000 });
};
const expandAll = () => {
    expandedRows.value = dataNilaiSiswa.value.reduce((acc, p) => (acc[p.rombonganBelajarId] = true) && acc, {});
};
const collapseAll = () => {
    expandedRows.value = null;
};

watch(kelasSelected, (newVal) => {
    loadSiswaAktif();
});
const loadSiswaAktif = async () => {
    isLoading.value = true;
    try {
        dataNilaiSiswa.value = await sekolahService.fetchAnggotaKelas(kelasSelected.value?.rombonganBelajarId, sekolahService.initSelectedSemester.value?.semesterId);
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: 'Gagal mendapatkan nilai siswa', life: 3000 });
    } finally {
        isLoading.value = false;
    }
};

const jenjang = 'SMA';
const totalSemesters = computed(() => {
    //   const jenjang = props.jenjang; // Harus dikirim dari parent

    if (jenjang === 'TK') return 2;
    if (jenjang === 'SD') return 12;
    if (jenjang === 'SMP') return 6;
    if (jenjang === 'SMA') return 6;
    return 6; // default
});

// Data siswa
const siswa = {
    nama: 'Ahmad',
    jenjang: 'SMA',
    nilai: [
        { mataPelajaran: 'Bahasa Indonesia', semester1: 80, semester2: 85, semester3: 78, semester4: 82, semester5: 88, semester6: 90 },
        { mataPelajaran: 'Matematika', semester1: 75, semester2: 70, semester3: 80, semester4: 85, semester5: 83, semester6: 87 },
        { mataPelajaran: 'Bahasa Inggris', semester1: 88, semester2: 90, semester3: 92, semester4: 89, semester5: 91, semester6: 93 }
    ]
};
onMounted(() => {
    // loadSiswaAktif();
});
</script>
