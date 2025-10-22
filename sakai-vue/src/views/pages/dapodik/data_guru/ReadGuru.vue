<script setup>
import { useUtils } from '@/composables/useUtils';
import { nextTick, onMounted, ref, watch } from 'vue';
const { formatterDateID } = useUtils();
// import FileUpload from 'primevue/fileupload';

import { useGuru } from '@/composables/sekolah_composable/useGuru';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import router from '@/router';
import { FilterMatchMode } from '@primevue/core/api';
const loading = ref(false);
const { initSelectedSemester } = useSemester();
const dt = ref();

const deleteGuruDialog = ref(false);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS }
});

const openNew = async () => {
    loadingAdd.value = true;
    router.push({ name: 'inputGuru', params: { sekolah: 'smkspasundanjatinangor' } }).catch((err) => {
        console.error('Router error:', err);
    });
    loadingAdd.value = false;
};

const exportCSV = () => {
    dt.value.exportCSV();
};

onMounted(async () => {
    await fetchGuruTerdaftar();
});

// ==================================
// =======composable=============
const selectedGuru = ref([]);
const { guruTerdaftarList, fetchGuruTerdaftar, deleteGuruTerdaftar, deleteBatchGuruTerdaftar } = useGuru();

// ==================================

watch(initSelectedSemester, () => {
    fetchGuruTerdaftar();
});

// ========IMPORT DATA========
const dialogImport = ref(false);

const deletGuru = () => {
    guruTerdaftarList.value = guruTerdaftarList.value.filter((val) => !selectedGuru.value.includes(val));
    if (selectedGuru.value.length == 1) {
        deleteGuruTerdaftar(selectedGuru.value[0].ptkTerdaftarId);
    } else if (selectedGuru.value.length > 1) {
        const ids = selectedGuru.value.map((item) => item.ptkTerdaftarId);
        deleteBatchGuruTerdaftar(ids);
    }
};

const editGuru = async () => {
    await nextTick();
    loadingEdit.value = true;
    router.push({
        name: 'editGuru',
        query: { ptkTerdaftarId: selectedGuru.value[0]?.ptkTerdaftarId.toString() }
    });
};

const afterUpload = async (e) => {
    console.log('cek', e);
    await fetchGuruTerdaftar();
    // if (!e) {
    // }
};
const loadingEdit = ref(false);
const loadingAdd = ref(false);
</script>

<template>
    <div class="">
        <Toolbar>
            <template #start>
                <Button icon="pi pi-pencil" severity="warn" @click="editGuru" :disabled="!selectedGuru || !selectedGuru.length || selectedGuru.length > 1" class="mr-2" v-tooltip.bottom="'Edit Guru'" :loading="loadingEdit" />
                <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg" @click="deleteGuruDialog = true" :disabled="!selectedGuru || !selectedGuru.length" v-tooltip.bottom="'Hapus Guru'" :loading="loading" />
                <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2 text-sm" v-tooltip.bottom="'Download Guru'" />
                <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="openNew" v-tooltip.bottom="'Tambah Guru Baru'" :loading="loadingAdd" :disabled="selectedGuru.length" />
                <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Upload Guru'" />
            </template>
            <template #end>
                <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" /> -->
                <IconField>
                    <InputIcon>
                        <i class="pi pi-search" />
                    </InputIcon>
                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                </IconField>
            </template>
        </Toolbar>

        <div v-if="guruTerdaftarList.length === 0" class="flex justify-center h-32 items-center"><h5>Tidak ada data</h5></div>
        <DataTable
            v-else
            ref="dt"
            v-model:selection="selectedGuru"
            stripedRows
            size="small"
            :value="guruTerdaftarList"
            dataKey="ptkTerdaftarId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} Guru"
        >
            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="ptk.nama" header="Nama" sortable> </Column>
            <Column field="ptkPelengkap.gelarBelakang" header="Gelar belakang"> </Column>
            <Column field="ptk.jenisKelamin" header="JK"> </Column>
            <Column field="ptkPelengkap.nip" header="NIP"> </Column>
            <Column field="ptkPelengkap.nuptk" header="NUPTK"> </Column>
            <Column field="ptk.tempatLahir" header="Tpt.Lahir"> </Column>
            <Column field="ptk.tanggalLahir" header="Tgl.Lahir">
                <template #body="slotProps">
                    {{ formatterDateID(slotProps.data.ptk.tanggalLahir) }}
                </template>
            </Column>
        </DataTable>

        <Dialog v-model:visible="deleteGuruDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                Yakin akan menghapus<span v-if="selectedGuru.length == 1"
                    ><b>{{ selectedGuru[0].ptk.nama }}</b
                    >?</span
                >
            </div>
            <template #footer>
                <Button label="No" icon="pi pi-times" text @click="deleteGuruDialog = false" />
                <Button label="Yes" icon="pi pi-check" @click="deletGuru" />
            </template>
        </Dialog>

        <DialogImport v-model:visible="dialogImport" template-type="guru" @save="afterUpload" />
    </div>
</template>
