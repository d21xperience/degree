<script setup>
// import LoadingOverlay from '@/components/LoadingOverlay.vue';
import DialogImport from '@/components/DialogImport.vue';
import KelasComponent from '@/components/sekolah_components/KelasComponent.vue';
import { useNilai } from '@/composables/sekolah_composable/useNilai';
import { useSiswa } from '@/composables/sekolah_composable/useSiswa';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import { computed, inject, ref, watch } from 'vue';
const pembelajaran = ref({});
const dataNilaiSiswa = ref([]);
const kelasSelected = ref(null);
const expandedRows = ref();
// watch(selectedSemester, async (newVal, oldVal) => {
//     // console.log(newVal)
//     dataNilaiSiswa.value = await fetchNilaiSiswa();
// });
const { searchNilai } = useNilai();
// const { initSelectedSemester } = useSemester();
const { getSiswaAktifByKelasId } = useSiswa();
const tahunAjaranId = inject('tahunAjaranProvider');
const initSelectedSemester = inject('selectedSemesterProvider');
// const selectedSemester = computed(() => initSelectedSemester.value?.label);
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

// const exportCSV = () => {
//     // isLoading.value = true;
//     // alert("hello")
//     // dt.value.exportCSV();
// };

const isDialogImport = ref(false);
const saveImport = () => {
    // console.log("Data disimpan:", e);
    isDialogImport.value = false;
};

const cancelImport = () => {
    console.log('Import dibatalkan');
    isDialogImport.value = false;
};
// ===========================================
let jenjang = '';
const totalSemesters = computed(() => {
    if (jenjang === 'TK / sederajat') return 2;
    if (jenjang === 'SD / sederajat') return 12;
    if (jenjang === 'SMP / sederajat') return 6;
    if (jenjang === 'SMA / sederajat') return 6;
    return 6; // default
});

// Data siswa
const siswa = ref({
    jenjang: '',
    nilaiMapel: []
});
// const siswa = ref()
const resetSiswa = () => {
    dataNilaiSiswa.value = [];
    siswa.value.jenjang = '';
    siswa.value.nilaiMapel = [];
};
const onRowExpand = async (event) => {
    try {
        const res = await searchNilai(event.data.pesertaDidikId);
        if (res.status) {
            Object.assign(siswa.value, res.nilai);
            jenjang = siswa.value.jenjang;
            toast.add({ severity: 'info', summary: 'Success', detail: `${res.message}`, life: 3000 });
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: 'Gagal mengambil nilai', life: 3000 });
    }
};
const onRowCollapse = (event) => {
    toast.add({ severity: 'success', summary: 'Product Collapsed', detail: event, life: 3000 });
};
const expandAll = () => {
    expandedRows.value = dataNilaiSiswa.value.reduce((acc, p) => (acc[p.rombonganBelajarId] = true) && acc, {});
};
const collapseAll = () => {
    expandedRows.value = null;
};

watch(kelasSelected, (newVal) => {
    // console.log(!newVal);
    if (!newVal) {
        resetSiswa();
    } else {
        loadSiswaAktif();
    }
});
const loadSiswaAktif = async () => {
    isLoading.value = true;
    try {
        dataNilaiSiswa.value = await getSiswaAktifByKelasId(kelasSelected.value?.rombonganBelajarId, initSelectedSemester.value?.semesterId);
        // console.log(dataNilaiSiswa.value);
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Failled', detail: 'Gagal mendapatkan nilai siswa', life: 3000 });
    } finally {
        isLoading.value = false;
    }
};

const importNIlai = () => {
    isDialogImport.value = true;
};

const paramNilai = computed(() => {
    let tes = {
        rombelId: kelasSelected.value?.rombonganBelajarId,
        kurikulumId: String(kelasSelected.value?.kurikulumId),
        tahunAjaranId: tahunAjaranId.value.label,
        tingkatPendidikanId: String(kelasSelected.value?.tingkatPendidikanId)
    };
    return tes;
});
</script>

<template>
    <div class="">
        <Toolbar>
            <template #start>
                <div class="w-56">
                    <KelasComponent v-if="initSelectedSemester" v-model="kelasSelected" class="mr-2" :init-selected-semester="initSelectedSemester" />
                </div>
                <div v-show="!!kelasSelected" class="ml-1 flex space-x-1">
                    <!-- <Button icon="pi pi-plus" severity="success" class="text-lg" @click="openNew" v-tooltip.bottom="'Tambah Siswa'" :loading="isOpenNew" /> -->
                    <!-- <Button icon="pi pi-pencil" severity="warn" @click="editNilai" :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 2" class="" v-tooltip.bottom="'Edit nilai'" :loading="loadingEdit" /> -->
                    <!-- <Button icon="pi pi-trash" severity="danger" class="text-lg" @click="deleteSiswaDialog = true" :disabled="!selectedSiswa || !selectedSiswa.length" v-tooltip.bottom="'Hapus Nilai'" /> -->
                    <Button v-tooltip.bottom="'Import Nilai'" icon="pi pi-upload" severity="" class="text-sm" @click="importNIlai" />
                </div>
            </template>
            <template #end>
                <div v-show="!!kelasSelected" class="mr-1">
                    <IconField>
                        <InputIcon>
                            <i class="pi pi-search"></i>
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." :disabled="!dataNilaiSiswa || (Array.isArray(dataNilaiSiswa) && !dataNilaiSiswa.length > 0)" />
                    </IconField>
                </div>
                <Button v-show="!!kelasSelected" v-tooltip.bottom="'Refresh'" icon="pi pi-refresh" severity="help" class="mr-2 text-sm" @click="loadSiswaAktif" />
            </template>
        </Toolbar>
        <div v-if="!initSelectedSemester" class="flex justify-center h-32 items-center">
            <h5>Silahkan pilih semester terlebih dahulu</h5>
        </div>
        <div v-else>
            <div v-if="!kelasSelected" class="flex justify-center h-32 items-center">
                <h5>Silahkan pilih kelas</h5>
            </div>
            <DataTable
                v-else
                ref="dt"
                v-model:expanded-rows="expandedRows"
                striped-rows
                size="small"
                :value="dataNilaiSiswa"
                data-key="pesertaDidikId"
                :paginator="true"
                :rows="10"
                :filters="filters"
                paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rows-per-page-options="[10, 20, 50]"
                current-page-report-template="Showing {first} to {last} of {totalRecords} siswa"
                @row-expand="onRowExpand"
                @row-collapse="onRowCollapse"
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
                <Column field="nmSiswa" header="Nama" sortable />
                <Column field="tingkatPendidikanId" header="Tingkat" />
                <Column field="nmKelas" header="Nama Kelas" />
                <Column field="" header="Aksi">
                    <template #body="{ data }">
                        <Button icon="pi pi-trash" outlined rounded class="mr-2" @click="editNilai(data)" />
                        <!-- <Button icon="pi pi-trash" outlined rounded severity="danger"
                                    @click="confirmdeleteMapel(data)" /> -->
                    </template>
                </Column>
                <template #expansion="">
                    <div class="p-4">
                        <DataTable :value="siswa.nilaiMapel">
                            <Column field="MataPelajaran" header="Mata Pelajaran" class="text-slate-500" />
                            <!-- Kolom Semester Dinamis -->
                            <Column v-for="n in totalSemesters" :key="`Semester${n}`" :field="`Semester${n}`" :header="`${n}`">
                                <!-- Opsional: Tambahkan warna berdasarkan nilai -->
                                <template #body="{ data }">
                                    <span
                                        :class="{
                                            'text-green-600 font-medium': data[`Semester${n}`] >= 85,
                                            'text-yellow-600': data[`Semester${n}`] >= 75 && data[`Semester${n}`] < 85,
                                            'text-red-600': data[`Semester${n}`] < 75
                                        }"
                                    >
                                        {{ data[`Semester${n}`] }}
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
        </div>

        <!-- import data -->
        <!-- DIALOG IMPORT -->
        <DialogImport
            v-if="initSelectedSemester"
            v-model:visible="isDialogImport"
            template-type="nilai"
            message="untuk nilai siswa sesuai dengan kelas yang dipilih."
            :selected-semester="initSelectedSemester"
            :param-nilai="paramNilai"
            @save="saveImport"
            @cancel="cancelImport"
        />

        <!-- end of import data -->
        <!-- <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay> -->
    </div>
</template>
