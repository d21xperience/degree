<script setup>
import router from '@/router';
// import { debounce } from 'lodash-es';

import { inject, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import DialogImport from '@/components/DialogImport.vue';
import { useUtils } from '@/composables/useUtils';

import { useSiswa } from '@/composables/sekolah_composable/useSiswa';
import { FilterMatchMode } from '@primevue/core/api';

// PROVIDER ============================
const tahunAjaranIdProvider = inject('tahunAjaranProvider');
const initSelectedSemester = inject('selectedSemesterProvider');
// =====================================
// const tahunAjaranId = computed(() => tahunAjaranIdProvider.value?.label);
watch(tahunAjaranIdProvider, () => {
    siswa.value = [];
});

const { fetchSiswaAktif, deleteSiswaAktif, deleteBatchSiswaAktif } = useSiswa();
// const { fetchTingkat } = useSekolah();

// const { initSelectedSemester } = useSemester();
// const initial = async () => {
//     try {
//         const res = await fetchSiswaAktif();
//         if (res.status) {
//         }
//     } catch (error) {}
// };
// const isEdit = ref(false);
const dt = ref();
const deleteSiswaDialog = ref(false);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkatPendidikanId: { value: null, matchMode: FilterMatchMode.EQUALS }
});

const isOpenNew = ref(false);
const openNew = () => {
    router.push({ name: 'createSiswa' });
};
const closeDialog = () => {
    deleteSiswaDialog.value = false;
    selectedSiswa.value = null;
};
const exportCSV = () => {
    dt.value.exportCSV();
};
const deleteSiswa = () => {
    siswa.value = siswa.value.filter((val) => !selectedSiswa.value.includes(val));
    if (selectedSiswa.value.length == 1) {
        deleteSiswaAktif(selectedSiswa.value[0].anggotaRombelId);
    } else if (selectedSiswa.value.length > 1) {
        const ids = selectedSiswa.value.map((item) => item.anggotaRombelId);
        deleteBatchSiswaAktif(ids);
    }
};

const selectedSiswa = ref([]);
const siswa = ref([]);
// ========IMPORT DATA========
const dialogImport = ref(false);
const saveImport = async () => {
    // console.log('Data disimpan:', e);
    dialogImport.value = false;
    const cek = await fetchSiswaAktif();
    // console.log(cek)
    siswa.value = cek;
};

const cancelImport = () => {
    // console.log('Import dibatalkan');
    dialogImport.value = false;
};

const { formatterDateID } = useUtils();
// const tingkatPendidikanOptions = ref([]);
const loadingEdit = ref(false);
const editSiswa = async () => {
    await nextTick();
    loadingEdit.value = true;
    router.push({
        name: 'editSiswa',
        params: { pesertaDidikId: selectedSiswa.value[0]?.pesertaDidikId.toString() }
    });
};

watch(initSelectedSemester, async (newVal) => {
    if (newVal) {
        console.log('watch.........', newVal);
        siswa.value = await fetchSiswaAktif(initSelectedSemester.value.semesterId);
    }
});
onMounted(async () => {
    if (initSelectedSemester.value) {
        siswa.value = await fetchSiswaAktif(initSelectedSemester.value.semesterId);
    }
});
onUnmounted(() => {
    console.log('[UNMOUNT] ReadSiswa dijalankan!');
    if (initSelectedSemester) return;
    siswa.value = [];
});
</script>

<template>
    <div>
        <Toolbar>
            <template #start>
                <div v-show="siswa.length > 0">
                    <Button v-tooltip.bottom="'Edit siswa'" icon="pi pi-pencil" severity="warn" :disabled="selectedSiswa.length == 0 || selectedSiswa.length > 1" class="mr-2" :loading="loadingEdit" @click="editSiswa" />
                    <Button v-tooltip.bottom="'Hapus siswa'" icon="pi pi-trash" severity="danger" class="mr-2 text-sm" :disabled="selectedSiswa.length == 0" @click="deleteSiswaDialog = true" />
                    <Button v-tooltip.bottom="'Download siswa'" icon="pi pi-download" severity="help" class="mr-2 text-sm" @click="exportCSV($event)" />
                </div>
                <div v-show="initSelectedSemester">
                    <Button v-tooltip.bottom="'Tambah Siswa'" icon="pi pi-plus" severity="success" class="mr-2 text-sm" :loading="isOpenNew" :disabled="selectedSiswa.length > 0" @click="openNew" />
                    <Button v-tooltip.bottom="'Upload siswa'" icon="pi pi-upload" severity="info" class="mr-2 text-sm" :disabled="selectedSiswa.length > 0" @click="dialogImport = true" />
                </div>
            </template>
            <template #end>
                <div v-show="siswa.length > 0" class="flex space-x-1">
                    <!-- <TingkatComponent v-model:model-value="filters['tingkatPendidikanId'].value" /> -->
                    <IconField>
                        <InputIcon>
                            <i class="pi pi-search"></i>
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." class="md:w-48" />
                    </IconField>
                </div>
            </template>
        </Toolbar>
        <div v-if="!initSelectedSemester" class="flex justify-center h-32 items-center">
            <h5>Silahkan pilih semester terlebih dahulu</h5>
        </div>
        <div v-else>
            <div v-if="siswa.length === 0" class="flex justify-center h-32 items-center"><h5>Tidak ada data</h5></div>

            <DataTable
                v-else
                ref="dt"
                v-model:selection="selectedSiswa"
                striped-rows
                :value="siswa"
                scrollable
                scroll-height="450px"
                data-key="pesertaDidikId"
                :paginator="true"
                :rows="10"
                :filters="filters"
                paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rows-per-page-options="[10, 20, 50]"
                current-page-report-template="Showing {first} to {last} of {totalRecords} siswa"
            >
                <template #empty>
                    <p class="text-xl flex justify-center font-bold text-red-500">Siswa tidak ditemukan</p>
                </template>

                <Column selection-mode="multiple" style="width: 3rem" :exportable="false" />
                <Column field="nmSiswa" header="Nama" sortable />
                <Column field="jenisKelamin" header="JK" />
                <Column field="nisn" header="NISN" />
                <Column field="nis" header="NIS" sortable />
                <Column field="agama" header="Agama" />
                <Column field="tempatLahir" header="Tpt Lahir" />
                <Column field="tanggalLahir" header="Tgl Lahir">
                    <template #body="slotProps">
                        {{ formatterDateID(slotProps.data.tanggalLahir) }}
                    </template>
                </Column>
                <Column field="tingkatPendidikanId" header="Tingkat" sortable />
                <Column field="nmKelas" header="Rombel" sortable />
            </DataTable>
        </div>
        <DialogImport v-if="initSelectedSemester" v-model:visible="dialogImport" :selected-semester="initSelectedSemester" template-type="siswa" @save="saveImport" @cancel="cancelImport" />

        <!-- end of import data -->
        <DialogConfirmDelete v-model:visible="deleteSiswaDialog" message="Apakah siswa tersebut akan dihapus?" @confirm="deleteSiswa" @close-dialog="closeDialog" />
    </div>
</template>
