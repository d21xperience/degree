<script setup>
import DialogImport from '@/components/DialogImport.vue';
import { useSekolahService } from '@/composables/sekolah_composable/useSekolah';

import { FilterMatchMode } from '@primevue/core/api';
import { computed, onMounted, ref } from 'vue';
// Kompoenen Tingkat
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    'kelas.nmKelas': { value: null, matchMode: FilterMatchMode.CONTAINS }
});
const sekolahService = useSekolahService();

const isEditKategoriSekolah = ref(false);
const confirmDeleteSelected = () => {};
const selectedSiswa = ref();
const sekolah = computed(() => sekolahService.sekolah.value);
const expandedRows = ref(null);

const siswa = ref([]);
const namaKelas = ref();

const dialogImport = ref(false);
const saveImport = async (e) => {
    // console.log('Data disimpan:', e);
    dialogImport.value = false;
    const cek = await fetchSiswaAktif();
    // console.log(cek);
    siswa.value = cek;
};

const cancelImport = () => {
    // console.log('Import dibatalkan');
    dialogImport.value = false;
};
const jenjangPendidikan = ref(2);

const initJenjang = () => {
    if (sekolah.value.sekolah.jenjangPendidikanId == 6) {
        jenjangPendidikan.value = 6;
    }
    // console.log(jenjangPendidikan.value);
};

onMounted(async () => {
    // console.log(sekolah.value)
    initJenjang();
    siswa.value = await sekolahService.getDns(sekolahService.initSelectedTahunAjaran.value?.tahunAjaranId);
    // console.log(siswa.value)
    namaKelas.value = [...new Set(siswa.value.map((item) => item.kelas?.nmKelas).filter((nm) => nm))].map((nm) => ({
        nama: nm,
        value: nm.toLowerCase()
    }));
    // namaKelas.value = getNmKelas(siswa);
});
</script>

<template>
    <!-- Main Table -->
    <div class="grid grid-cols-1 gap-4">
        <div class=" ">
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Upload nilai'" />
                    </template>
                    <template #end>
                        <Select v-model="filters['kelas.nmKelas'].value" :options="namaKelas" optionLabel="nama" optionValue="value" placeholder="Kelas" class="w-full md:w-48 md:mr-2" checkmark show-clear />

                        <IconField>
                            <InputIcon>
                                <i class="pi pi-search" />
                            </InputIcon>
                            <InputText v-model="filters['global'].value" placeholder="Search..." name="search" id="search" />
                        </IconField>
                    </template>
                </Toolbar>
            </div>
        </div>
        <DataTable
            v-model:expandedRows="expandedRows"
            :value="siswa"
            dataKey="pesertaDidikId"
            striped-rows
            :paginator="true"
            :rows="10"
            size="small"
            :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} siswa"
        >
            <Column expander style="width: 2rem" />
            <Column header="Nama" field="nama" style="width: 20rem" />
            <Column header="nisn" field="nisn" />
            <Column header="Kelas" field="kelas.nmKelas" />
            <!-- <Column header="Tahun Ajaran" field="tahun_ajaran_id" /> -->
            <Column header="Aksi" :hidden="false">
                <template #body="slotProps">
                    <!-- <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="handleDeleteKategoriSekolah(slotProps.data)" size="small" rounded v-tooltip.bottom="'Hapus kompetensi'" /> -->
                    <Button icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" @click="dialogEditKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Edit kelas'" />
                </template>
            </Column>
            <template #expansion="slotProps">
                <DataTable :value="slotProps.data.kategorikelas">
                    <Column header="Mapel" field="tingkat_id" />
                    <template v-for="i in jenjangPendidikan">
                        <Column :header="i" field="jumlah" />
                    </template>
                    <Column header="Aksi" :hidden="!isEditKategoriSekolah">
                        <template #body="slotProps">
                            <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="handleDeleteKategoriKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Hapus kelas'" />
                            <Button icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" @click="dialogEditKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Edit kelas'" />
                        </template>
                    </Column>
                </DataTable>
            </template>
        </DataTable>
    </div>
    <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="transkrip" />
</template>
