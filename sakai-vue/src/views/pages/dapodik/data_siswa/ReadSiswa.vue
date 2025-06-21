<template>
    <div>
        <div>
            <div>
                <div class="mb-2">
                    <Toolbar>
                        <template #start>
                            <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="openNew" v-tooltip.bottom="'Tambah Siswa'" :loading="isOpenNew" />
                            <Button icon="pi pi-pencil" severity="warn" @click="openNew" :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 2" class="mr-2" v-tooltip.bottom="'Edit siswa'" />
                            <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg" @click="deleteSiswaDialog = true" :disabled="!selectedSiswa || !selectedSiswa.length" v-tooltip.bottom="'Hapus siswa'" />
                            <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Upload siswa'" v-show="selectedSemester.semester == 1" />
                            <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2 text-sm" v-tooltip.bottom="'Download siswa'" />
                        </template>
                        <template #end>
                            <div class="flex">
                                <Select v-model="filters['tingkatPendidikanId'].value" :options="tingkatPendidikanOptions" optionLabel="nama" optionValue="kode" placeholder="Tingkat" class="w-full md:w-36 mr-2" checkmark show-clear />
                                <IconField>
                                    <InputIcon>
                                        <i class="pi pi-search" />
                                    </InputIcon>
                                    <InputText v-model="filters['global'].value" placeholder="Search..." class="md:w-48" />
                                </IconField>
                            </div>
                        </template>
                    </Toolbar>
                </div>
            </div>

            <DataTable
                ref="dt"
                v-model:selection="selectedSiswa"
                stripedRows
                :value="siswa"
                scrollable
                scrollHeight="450px"
                dataKey="pesertaDidikId"
                :paginator="true"
                :rows="10"
                :filters="filters"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[10, 20, 50]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} siswa"
            >
                <template #empty>
                    <p class="text-xl flex justify-center font-bold text-red-500">Siswa tidak ditemukan</p>
                </template>

                <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                <Column field="nmSiswa" header="Nama" sortable></Column>
                <Column field="jenisKelamin" header="JK"></Column>
                <Column field="nisn" header="NISN"></Column>
                <Column field="nis" header="NIS" sortable></Column>
                <Column field="agama" header="Agama"></Column>
                <Column field="tempatLahir" header="Tpt Lahir"></Column>
                <Column field="tanggalLahir" header="Tgl Lahir">
                    <template #body="slotProps">
                        {{ formatterDateID(slotProps.data.tanggalLahir) }}
                    </template>
                </Column>
                <Column field="tingkatPendidikanId" header="Tingkat" sortable></Column>
                <Column field="nmKelas" header="Rombel" sortable></Column>
            </DataTable>
        </div>

        <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="siswa" />

        <!-- end of import data -->
        <DialogConfirmDelete v-model:visible="deleteSiswaDialog" message="Apakah siswa tersebut akan dihapus?" @confirm="deleteSiswa" @closeDialog="closeDialog" />
    </div>
</template>

<script setup>
import { debounce } from 'lodash-es';

import { useSekolahService } from '@/composables/useSekolahService';
import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// import FileUpload from 'primevue/fileupload';

import Column from 'primevue/column';
import DataTable from 'primevue/datatable';

import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import DialogImport from '@/components/DialogImport.vue';
import { useUtils } from '@/composables/useUtils';
import router from '@/router';
import { FilterMatchMode } from '@primevue/core/api';
import Button from 'primevue/button';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import Toolbar from 'primevue/toolbar';

// ==============================
onMounted(async () => {
    siswa.value = await fetchSiswaAktif();
    tingkatPendidikanOptions.value = await fetchTingkat();
});
const isEdit = ref(false);
const dt = ref();
const deleteSiswaDialog = ref(false);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkatPendidikanId: { value: null, matchMode: FilterMatchMode.EQUALS }
});

const isOpenNew = ref(false);
const openNew = debounce(() => {
    router.push({ name: 'inputSiswa' });
}, 250);
const closeDialog = () => {
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

import Select from 'primevue/select';

const selectedSiswa = ref();
const siswa = ref([]);
// ================================
// composable
// ================================
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
// const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
const { fetchSiswaAktif, fetchTingkat, deleteSiswaAktif, deleteBatchSiswaAktif } = useSekolahService();
// ================================
watch(selectedSemester, async (e, b) => {
    siswa.value = await fetchSiswaAktif();
});
// ========IMPORT DATA========
const dialogImport = ref(false);
const saveImport = async (e) => {
    console.log('Data disimpan:', e);
    dialogImport.value = false;
    const cek = await fetchSiswaAktif();
    console.log(cek)
    siswa.value = cek
};

const cancelImport = () => {
    // console.log('Import dibatalkan');
    dialogImport.value = false;
};

const { formatterDateID } = useUtils();
const tingkatPendidikanOptions = ref([]);
</script>
