<template>
    <div class="">
        <div class="">
            <Toolbar>
                <template #start>
                    <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="isVisible = !isVisible" v-tooltip.bottom="'Tambah Siswa'" />

                    <Button
                        icon="pi pi-trash"
                        severity="danger"
                        class="mr-2 text-lg"
                        @click="dialogBatchDelete"
                        :disabled="!selectedKategoriMapel || !selectedKategoriMapel.length || selectedKategoriMapel.length == 1"
                        v-tooltip.bottom="'Hapus banyak mapel'"
                    />
                    <!--<Button icon="pi pi-pencil" severity="warn" @click="openNew" :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 2" class="mr-2" v-tooltip.bottom="'Edit siswa'" />
                    <Button icon="pi pi-upload" severity="info" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Upload siswa'" v-show="selectedSemester.semester == 1" /> -->
                    <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2 text-sm" v-tooltip.bottom="'Download siswa'" />
                    <Select v-model="selectedKategoriSekolah" :options="kategoriSekolahList" optionLabel="nama_kurikulum" placeholder="Kurikulum" class="mr-2 !w-96" checkmark fluid />
                </template>
                <template #end>
                    <div class="flex">
                        <TingkatComponent v-model="selectedTingkat" :initial-value="initialTingkat" />
                        <!-- <IconField>
                            <InputIcon>
                                <i class="pi pi-search" />
                            </InputIcon>
                            <InputText v-model="filters['global'].value" placeholder="Search..." class="md:w-48" />
                        </IconField> -->
                    </div>
                </template>
            </Toolbar>
            <div class="w-full my-2 container">
                <DataTable
                    ref="dt"
                    v-model:selection="selectedKategoriMapel"
                    stripedRows
                    size="small"
                    :value="kategoriMapelList"
                    dataKey="id"
                    :paginator="true"
                    :rows="10"
                    :filters="filters"
                    :first="first"
                    @page="(e) => (first = e.first)"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 50]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas"
                >
                    <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                    <Column header="No" style="width: 2rem">
                        <template #body="slotProps">
                            <!-- {{ slotProps.index + 1 + (dt.value?.first || 0) }} -->
                            {{ slotProps.index + 1 + first }}
                        </template>
                    </Column>
                    <Column field="nmMapel" header="Nama"></Column>
                    <Column field="tingkatPendidikan" header="Tingkat" sortable></Column>
                    <Column field="" header="Jurusan" sortable>
                        <template #body>
                            {{ selectedKategoriSekolah?.namaJurusan }}
                        </template>
                    </Column>
                    <!-- <Column field="ptk.nama" header="Guru Mapel"></Column> -->
                    <!-- <Column field="ptk.nama" header="Jml.Mapel"></Column> -->
                    <Column field="" header="Aksi">
                        <template #body="{ data }">
                            <!-- <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editMapel(data)" size="small" :rounded="true" /> -->
                            <Button icon="pi pi-trash" outlined rounded severity="danger" @click="dialogDelete(data)" size="small" :rounded="true" />
                        </template>
                    </Column>
                    <!-- <template #expansion="slotProps">
                        <div class="p-4">
                            <DataTable :value="slotProps.data.pembelajaran">
                                <Column field="namaMataPelajaran" header="Mata pelajaran" sortable></Column>
                                <Column field="ptkTerdaftar.ptk.nama" header="Guru Mapel" sortable></Column>
                            </DataTable>
                        </div>
                    </template> -->
                </DataTable>
            </div>
        </div>

        <!-- DIALOGBOX FOR EDIT DATA -->
        <Dialog v-model:visible="mapelDialog" :style="{ width: '50%' }" header="Edit Data" :modal="true" position="top">
            <div class="">
                <div class="my-2">
                    <label for="name" class="block font-bold">Nama Kelas</label>
                    <InputText id="name" v-model.trim="kelas.nmKelas" required="true" autofocus disabled :invalid="submitted && !kelas.nmKelas" fluid class="w-full" />
                    <small v-if="submitted && !kelas.nmKelas" class="text-red-500">NISN is required.</small>
                </div>

                <div>
                    <div class="mb-2">
                        <Toolbar>
                            <template #start>
                                <Button icon="pi pi-plus" severity="success" class="mr-2" @click="openNewMapel" v-tooltip.bottom="'Tambah data'" />
                                <!-- <Button icon="pi pi-pencil" severity="warn" @click="editKelas(selectedKelas)"
                                    :disabled="!selectedKelas || !selectedKelas.length || selectedKelas.length > 1"
                                    class="mr-2" v-tooltip.bottom="'Edit data'" />
                                <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                                    :disabled="!selectedKelas || !selectedKelas.length"
                                    v-tooltip.bottom="'Hapus data'" /> -->
                            </template>
                            <template #end>
                                <Button label="Import" icon="pi pi-download" severity="warn" @click="dialogImport = true" class="mr-2 text-sm" v-tooltip.bottom="'Import siswa'" />
                                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)" class="mr-2 text-sm" />
                            </template>
                        </Toolbar>
                    </div>
                    <!-- <DataTable ref="dt" :value="pembelajaranList" dataKey="pembelajaran_id">
                        <Column field="nama_mata_pelajaran" header="Mata pelajaran"></Column>
                        <Column field="nama" header="Guru Mapel"></Column>
                    </DataTable> -->
                </div>
            </div>

            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
                <Button label="Save" icon="pi pi-check" @click="simpanKeDatabase" />
            </template>
        </Dialog>

        <Dialog v-model:visible="deletemapelDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="product"
                    >Are you sure you want to delete <b>{{ product.name }}</b
                    >?</span
                >
            </div>
            <template #footer>
                <Button label="No" icon="pi pi-times" text @click="deletemapelDialog = false" />
                <Button label="Yes" icon="pi pi-check" @click="deleteMapel" />
            </template>
        </Dialog>

        <Dialog v-model:visible="deleteMapelsDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="product">Apakah data lulusan akan dihapus?</span>
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="deleteMapelsDialog = false" />
                <Button label="Ya" icon="pi pi-check" text @click="deletedataLulusan" />
            </template>
        </Dialog>

        <!-- Dialog Tambah mapel -->
        <Dialog v-model:visible="addMapelDialog" :style="{ width: '450px' }" header="Tambah mapel" :modal="true" position="top">
            <div class="">
                <div class="flex space-x-1">
                    <div class="">
                        <label for="mata-pelajaran">Mata pelajaran</label>
                        <AutoComplete
                            v-model="selectedMapel"
                            :suggestions="filteredMapel"
                            optionLabel="nama"
                            @complete="searchMapel"
                            @keydown.space.prevent="handleKeydown"
                            placeholder="Cari Mapel..."
                            class="w-full"
                            fluid
                            :invalid="submitted && !selectedMapel"
                        />
                        <small v-if="submitted && !selectedMapel" class="text-red-500">Subject is required.</small>
                    </div>
                    <div class="">
                        <label for="mata-pelajaran">Guru Mapel</label>
                        <AutoComplete
                            v-model="selectedGuru"
                            :suggestions="filteredGuru"
                            optionLabel="ptk.nama"
                            @complete="searchGuru"
                            @keydown.space.prevent="handleKeydown"
                            placeholder="Cari Guru..."
                            class="w-full"
                            fluid
                            dropdown
                            :invalid="submitted && !selectedGuru"
                        />
                        <small v-if="submitted && !selectedGuru" class="text-red-500">Teacher is required.</small>
                    </div>
                </div>
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="cancelAddMapel" />
                <Button label="Tambah" icon="pi pi-check" text @click="tambahPembelajaran" />
            </template>
        </Dialog>

        <!-- import data -->
        <!-- DIALOG IMPORT -->
        <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="mapel" />
        <DialogConfirmDelete v-model:visible="isDelete" @confirm="confirmDelete" :message="messageDelete" />
        <DialogConfirmDelete v-model:visible="isBatchDelete" @confirm="confirmBatchDelete" :message="messageBatchDelete" />
        <!-- end of import data -->
        <!-- <DialogMapel v-model:visible="isVisible" /> -->
    </div>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();

// import FileUpload from 'primevue/fileupload';
import DialogImport from '@/components/DialogImport.vue';
import TingkatComponent from '@/components/TingkatComponent.vue';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';

import Button from 'primevue/button';

import Dialog from 'primevue/dialog';

import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import DialogMapel from '@/components/DialogMapel.vue';
import { FilterMatchMode } from '@primevue/core/api';
import AutoComplete from 'primevue/autocomplete';
import InputText from 'primevue/inputtext';
import Toolbar from 'primevue/toolbar';
import { useToast } from 'primevue/usetoast';
// =============UJI FITUR FOTO========================
// import Image from 'primevue/image';
// =====================================
const pembelajaran = ref({});
const pembelajaranList = ref([]);
const guruList = ref();

// const fetchMapel = async () => {
//     try {
//         const cekMapel = await store.getters['sekolahService/getMapel'];
//         // console.log(cekMapel)
//         if (cekMapel == null) {
//             const payload = {
//                 mapel_id: ''
//             };
//             const response = await store.dispatch('sekolahService/fetchMapel', payload);
//             // console.log(response)
//             mapelList.value = response;
//         } else {
//             mapelList.value = cekMapel;
//         }
//     } catch (error) {
//         console.log(error);
//     }
// };

// ==============================
const kelasList = ref([]);
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService';
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);

const sekolahService = useSekolahService();

const kategoriSekolahList = ref([]);
const kategoriMapelList = ref([]);
const selectedKategoriSekolah = ref();
const selectedTingkat = ref([]);
const initialTingkat = ref();
const initKategoriMapel = async () => {
    const payload = {
        kurikulumId: selectedKategoriSekolah.value?.kurikulum_id,
        tingkatPendidikan: Number(selectedTingkat.value)
    };
    // console.log(payload);
    // return;
    if (payload.kurikulumId && payload.tingkatPendidikan) {
        // alert("Data error")
    }

    kategoriMapelList.value = await sekolahService.fetchKategoriMapel(payload);
    // console.log(kategoriMapelList.value);
};

const initFirst = async () => {
    await sekolahService.fetchKategoriSekolah();
    kategoriSekolahList.value = sekolahService.kategoriSekolahTabel.value;
    selectedKategoriSekolah.value = kategoriSekolahList.value[0];

    const results = await sekolahService.fetchTingkat();
    initialTingkat.value = results[0].kode;
};

watch(sekolahService.initSelectedSemester, () => {
    // console.log(newVal);
    initFirst();
    //    return sekolahService.initSelectedSemester.value?.tahunAjaranId;
});

watch(selectedKategoriSekolah, () => {
    initKategoriMapel();
});
const isLoading = ref(false);
watch(selectedTingkat, () => {
    initKategoriMapel();
});
// const dataConnected = ref(true);
const toast = useToast();
const dt = ref(null);
// const products = ref();
const mapelDialog = ref(false);
const mapelList = ref([]);
const deletemapelDialog = ref(false);
const deleteMapelsDialog = ref(false);
const product = ref({});
const dataLulusan = ref();
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkatPendidikanId: { value: null, matchMode: FilterMatchMode.EQUALS }
});
const submitted = ref(false);
const addMapelDialog = ref(false);
const openNewMapel = async () => {
    addMapelDialog.value = true;
    fetchMapel();
    // fetchGuru();
};
const hideDialog = () => {
    mapelDialog.value = false;
    // selectedMapel.value =
    submitted.value = false;
};
const tambahPembelajaran = () => {
    console.log('submitted', submitted.value);
    // console.log("selecteGuru",!selectedGuru.value)
    // console.log(submitted.value && !selectedGuru.value)
    submitted.value = true;
    if (selectedMapel?.value.nama?.trim()) {
        if (pembelajaran.value.pembelajaran_id) {
            pembelajaran.value.inventoryStatus = pembelajaran.value.inventoryStatus.value ? pembelajaran.value.inventoryStatus.value : pembelajaran.value.inventoryStatus;
            // products.value[findIndexById(product.value.id)] = product.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        } else {
            pembelajaran.value.pembelajaran_id = generateUUID();
            pembelajaran.value.ptk_terdaftar_id = selectedGuru.value?.ptkTerdaftarId;
            pembelajaran.value.nama = selectedGuru.value?.ptk.nama;
            pembelajaran.value.mata_pelajaran_id = selectedMapel.value.mataPelajaranId;
            pembelajaran.value.nama_mata_pelajaran = selectedMapel.value.nama;
            // product.value.code = createId();
            // product.value.image = 'product-placeholder.svg';
            // product.value.inventoryStatus = product.value.inventoryStatus ? product.value.inventoryStatus.value : 'INSTOCK';
            pembelajaranList.value.push(pembelajaran.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
        }

        // console.log(pembelajaran.value)

        selectedGuru.value = {};
        selectedMapel.value = {};
        // pembelajaran.value = {};
    }
};
const kelas = ref({});
const editMapel = (mapel) => {
    kelas.value = { ...mapel };
    mapelDialog.value = true;
    pembelajaran.value.rombongan_belajar_id = kelas.value.rombonganBelajarId;
    pembelajaran.value.semester_id = kelas.value.semesterId;
};
const deleteMapel = () => {
    products.value = products.value.filter((val) => val.id !== product.value.id);
    deletemapelDialog.value = false;
    product.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
};
const exportCSV = () => {
    isLoading.value = true;
    // alert("hello")
    // dt.value.exportCSV();
};
const deletedataLulusan = () => {
    products.value = products.value.filter((val) => !dataLulusan.value.includes(val));
    deleteMapelsDialog.value = false;
    dataLulusan.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};
// ==================================
// ========IMPORT DATA========
const dialogImport = ref(false);
const saveImport = (e) => {
    // console.log("Data disimpan:", e);
    dialogImport.value = false;
};

const cancelImport = () => {
    console.log('Import dibatalkan');
    dialogImport.value = false;
};
// ===========================================

// import { debounce } from 'lodash';

const selectedMapel = ref();
const filteredMapel = ref();
const searchMapel = (event) => {
    setTimeout(() => {
        let query = event.query.toLowerCase();
        filteredMapel.value = mapelList.value.filter((mapel) => mapel.nama.toLowerCase().includes(query));
    }, 250);
};
const selectedGuru = ref();
const filteredGuru = ref();
const searchGuru = (event) => {
    setTimeout(() => {
        let query = event.query.toLowerCase();

        filteredGuru.value = guruList.value.filter((guru) => guru.ptk.nama.toLowerCase().includes(query));
    }, 250);
};
const handleKeydown = (event) => {
    if (event.key === ' ') {
        selectedMapel.value += ' '; // Menambahkan spasi ke query
    }
};

const cancelAddMapel = () => {
    addMapelDialog.value = false;
    selectedGuru.value = {};
    selectedMapel.value = {};
};

const generateUUID = () => crypto.randomUUID();

const saveToDB = (req_Object, endpoint_String) => {
    console.log(endpoint_String);
    try {
        const payload = {
            schemaname: req_Object?.schemaname,
            pembelajaran: req_Object?.data
        };
        const results = store.dispatch(endpoint_String, payload);
        if (results) {
            return results;
        } else {
            null;
        }
    } catch (error) {
        console.log(error);
    }
};

const simpanKeDatabase = () => {
    // mapelDialog.value = false;
    // simpan di localstorage
    const aData = [];
    aData.push(pembelajaran.value);
    const req_Object = {
        schemaname: schemaname.value,
        data: aData
    };
    // console.log(req_Object)
    const endpoint = 'sekolahService/createPembelajaran';
    saveToDB(req_Object, endpoint);
    // localStorage.setItem("unsavedPembelajaran", JSON.stringify(pembelajaran.value));
};

// ======================================
//  DELETE
// ======================================
const isDelete = ref(false);
const messageDelete = ref('');
const selectedKategoriMapel = ref();
const dialogDelete = (data) => {
    isDelete.value = true;
    messageDelete.value = `Apakah mapel ${data.nmMapel} akan dihapus pada tahun ajaran ini?`;
    selectedKategoriMapel.value = data;
};

const confirmDelete = () => {
    kategoriMapelList.value = kategoriMapelList.value.filter((item) => item.id != selectedKategoriMapel.value.id);
    isDelete.value = false;
    sekolahService.deleteKategoriMapel(selectedKategoriMapel.value?.id);
    selectedKategoriMapel.value = null;
};

const isBatchDelete = ref(false);
const messageBatchDelete = ref('');
const dialogBatchDelete = () => {
    isBatchDelete.value = true;
    messageBatchDelete.value = `Apakah ${selectedKategoriMapel.value.length} mapel yang dipilih akan dihapus pada tahun ajaran ini?`;
};
const confirmBatchDelete = () => {
    if (selectedKategoriMapel.value.length > 1) {
        kategoriMapelList.value = kategoriMapelList.value.filter((item) => !selectedKategoriMapel.value.includes(item));
        let ids = selectedKategoriMapel.value.map((item) => item.id);
        sekolahService.deleteBatchKategoriMapel(ids);
        selectedKategoriMapel.value = null;
    }
};
const first = ref(0);
const isVisible = ref(false);
onMounted(async () => {
    // await fetchK();
    initFirst();
});
</script>
