<script setup>
import { computed, ref } from 'vue';

// ============toast============
import DialogLoading from './DialogLoading.vue';

// import { isError } from "lodash";

const isLoading = ref(false);

// ========================
// Props dari parent
const props = defineProps({
    visible: Boolean
    // templateType: String,
    // schemaName: String,
});

// Emit event ke parent
const emit = defineEmits(['update:visible', 'save', 'cancel']);

// Menggunakan computed agar bisa mengupdate prop.visible
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});

// Function untuk menutup dialog
const closeDialog = () => {
    isVisible.value = false;
};

// Function untuk mengunduh template
const isErr = ref(false);
</script>

<template>
    <Dialog v-model:visible="isVisible" header="Tambah Data" :modal="true">
        <div>
            <DataTable
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
                <Column field="nmKelas" header="Nama Siswa" />
                <Column field="tingkatPendidikanId" header="Jk" sortable />
                <Column field="kurikulum.namaKurikulum" header="nis" />
                <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
                <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                    <Column field="namaJurusanSp" header="Jurusan" sortable />
                </div>
                <Column field="ptk.nama" header="Wali kelas" />
                <Column field="jumlahAnggota" header="Jml.Anggota" />
                <Column header="Anggota Kelas">
                    <template #body="slotProps">
                        <Button icon="pi pi-bullseye" outlined rounded class="mr-2" @click="dialogAnggotaRombel(slotProps.data)" />
                    </template>
                </Column>
            </DataTable>
        </div>

        <template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Simpan" icon="pi pi-check" text @click="uploadToBackend" />
        </template>
    </Dialog>

    <DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>
    <Dialog v-model:visible="isErr" header="Warning!">
        <div>Pilih <b>Tahun Pelajaran</b> terlebih dahulu!</div>
    </Dialog>
</template>
