<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { useUtils } from '@/composables/useUtils';
import { nextTick, onMounted, ref, watch } from 'vue';
const { formatterDateID } = useUtils();
// import FileUpload from 'primevue/fileupload';

import Column from 'primevue/column';
import DataTable from 'primevue/datatable';

import Button from 'primevue/button';

import Dialog from 'primevue/dialog';

import { FilterMatchMode } from '@primevue/core/api';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import Toolbar from 'primevue/toolbar';
import { useToast } from 'primevue/usetoast';
import { useRouter } from 'vue-router';

const router = useRouter();
const loading = ref(false);

const toast = useToast();
const dt = ref();

const deleteGuruDialog = ref(false);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS }
});

const openNew = async () => {
    router.push({ name: 'inputGuru' });
};

const exportCSV = () => {
    dt.value.exportCSV();
};

// import EmptyData from '@/components/EmptyData.vue';
onMounted(async () => {
    await fetchGuruTerdaftar();
});

// ==================================
// =======composable=============
const selectedGuru = ref();
const { initSelectedSemester, guruTerdaftarList, fetchGuruTerdaftar, deleteGuruTerdaftar } = useSekolahService();
// ==================================

watch(initSelectedSemester, (e, b) => {
    fetchGuruTerdaftar();
});

// ========IMPORT DATA========
const dialogImport = ref(false);

const deletGuru = () => {
    guruTerdaftarList.value = guruTerdaftarList.value.filter((val) => val.ptkTerdaftarId !== selectedGuru.value[0].ptkTerdaftarId);
    deleteGuruDialog.value = false;
    console.log(selectedGuru.value[0]);
    // update ke backend
    deleteGuruTerdaftar(selectedGuru.value[0].ptkTerdaftarId);
    selectedGuru.value = [];
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Guru dihapus', life: 3000 });
};

const editGuru = async () => {
    await nextTick();
    router.push({
        name: 'editGuru',
        query: { ptkTerdaftarId: selectedGuru.value[0]?.ptkTerdaftarId.toString() }
    });
};
</script>

<template>
    <div class="">
        <div class="">
            <div class="w-full">
                <div class=" ">
                    <div class="mb-2">
                        <Toolbar>
                            <template #start>
                                <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="openNew" v-tooltip.bottom="'Tambah Guru Baru'" />
                                <Button icon="pi pi-pencil" severity="warn" @click="editGuru" :disabled="!selectedGuru || !selectedGuru.length || selectedGuru.length > 1" class="mr-2" v-tooltip.bottom="'Edit Guru'" />
                                <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg" @click="deleteGuruDialog = true" :disabled="!selectedGuru || !selectedGuru.length" v-tooltip.bottom="'Hapus Guru'" :loading="loading" />
                                <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Import Guru'" v-show="initSelectedSemester.semester == 1"/>
                                <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2 text-sm" v-tooltip.bottom="'Download Guru'" />
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
                    </div>

                    <!-- <Toolbar>
                        <template #end>
                            <IconField>
                                <InputIcon>
                                    <i class="pi pi-search" />
                                </InputIcon>
                                <InputText v-model="filters['global'].value" placeholder="Search..." />
                            </IconField>
                        </template>
                    </Toolbar> -->
                </div>

                <DataTable
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
                    <Column field="ptk.gelarBelakang" header="Gelar belakang"> </Column>
                    <Column field="ptk.jenisKelamin" header="JK"> </Column>
                    <Column field="ptk.nip" header="NIP"> </Column>
                    <Column field="ptk.nuptk" header="NUPTK"> </Column>
                    <Column field="ptk.tempatLahir" header="Tpt.Lahir"> </Column>
                    <Column field="ptk.tanggalLahir" header="Tgl.Lahir">
                        <template #body="slotProps">
                            {{ formatterDateID(slotProps.data.ptk.tanggalLahir) }}
                        </template>
                    </Column>
                </DataTable>
            </div>
        </div>

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

        <!-- import data -->
        <!-- <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="guru" :schema-name="schemaname" /> -->
        <DialogImport v-model:visible="dialogImport" template-type="kelas" />

        <!-- end of import data -->
    </div>
</template>
