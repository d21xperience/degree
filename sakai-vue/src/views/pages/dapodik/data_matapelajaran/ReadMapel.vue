<script setup>
import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// import FileUpload from 'primevue/fileupload';

import Column from 'primevue/column';
import DataTable from 'primevue/datatable';

import Button from 'primevue/button';

import Dialog from 'primevue/dialog';

import { FilterMatchMode } from '@primevue/core/api';
import InputText from 'primevue/inputtext';
import Toolbar from 'primevue/toolbar';
import { useToast } from 'primevue/usetoast';
// =============UJI FITUR FOTO========================
// import Image from 'primevue/image';
// =====================================
const pembelajaran = ref({});
const pembelajaranList = ref([]);
const guruList = ref();
const rombel = ref();

const fetchGuru = async () => {
    try {
        const cekGuru = await store.getters['sekolahService/getPTKTerdaftar'];
        if (cekGuru == null) {
            let payload = {
                tahunAjaranId: selectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            const response = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
            // console.log(response)
            guruList.value = response;
        } else {
            guruList.value = cekGuru;
        }
    } catch (error) {
        console.error(error);
    }
};

const fetchMapel = async () => {
    try {
        const cekMapel = await store.getters['sekolahService/getMapel'];
        // console.log(cekMapel)
        if (cekMapel == null) {
            const payload = {
                mapel_id: ''
            };
            const response = await store.dispatch('sekolahService/fetchMapel', payload);
            // console.log(response)
            mapelList.value = response;
        } else {
            mapelList.value = cekMapel;
        }
    } catch (error) {
        console.log(error);
    }
};

// ==============================
const kelasList = ref([]);
onMounted(async () => {
    await fetchK();
});
const fetchK = async () => {
    try {
        const res = store.getters['sekolahService/getKelas'](selectedSemester.value?.semesterId);
        if (!res || res.length === 0) {
            // console.log("fetch to backend!")
            kelasList.value = await fetchKelas();
        } else {
            kelasList.value = res;
        }
    } catch (error) {
        console.log(error);
    }
};
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService';
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
const { fetchKelas } = useSekolahService(schemaname, selectedSemester);
// ================================
watch(selectedSemester, async (newVal, oldVal) => {
    // console.log(newVal)
    kelasList.value = await fetchKelas();
});
// import DialogLoading from "@/components/DialogLoading.vue";

const isLoading = ref(false);

const dataConnected = ref(true);
const toast = useToast();
const dt = ref();
// const products = ref();
const mapelDialog = ref(false);
const mapelList = ref([]);
const deletemapelDialog = ref(false);
const deleteMapelsDialog = ref(false);
const product = ref({});
const dataLulusan = ref();
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    jurusan: { value: null, matchMode: FilterMatchMode.CONTAINS },
    tingkat: { value: null, matchMode: FilterMatchMode.GREATER_THAN }
});
const submitted = ref(false);
// const openNew = () => {
//     mapelDialog.value = true
//     // router.push({ name: "inputMapel" })
// };
const addMapelDialog = ref(false);
const openNewMapel = async () => {
    addMapelDialog.value = true;
    fetchMapel();
    fetchGuru();
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
const findIndexById = (id) => {
    let index = -1;
    for (let i = 0; i < products.value.length; i++) {
        if (products.value[i].id === id) {
            index = i;
            break;
        }
    }

    return index;
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
import DialogImport from '@/components/DialogImport.vue';
const expandedRows = ref();
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

const onRowExpand = (event) => {
    toast.add({ severity: 'info', summary: 'Product Expanded', detail: event.data.nmKelas, life: 3000 });
    // Ambil data mapel untuk kelas tertentu
    console.log(event);
};
const onRowCollapse = (event) => {
    toast.add({ severity: 'success', summary: 'Product Collapsed', detail: event.data.nmKelas, life: 3000 });
};
const expandAll = () => {
    expandedRows.value = kelasList.value.reduce((acc, p) => (acc[p.rombonganBelajarId] = true) && acc, {});
};
const collapseAll = () => {
    expandedRows.value = null;
};

import AutoComplete from 'primevue/autocomplete';
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
</script>
<template>
    <div class="">
        <div class="">
            <div class="w-full my-2 container">
                <h2 class="text-xl mb-2">Data Mapel</h2>
                <DataTable
                    ref="dt"
                    v-model:expandedRows="expandedRows"
                    stripedRows
                    size="small"
                    :value="kelasList"
                    @rowExpand="onRowExpand"
                    @rowCollapse="onRowCollapse"
                    dataKey="rombonganBelajarId"
                    :paginator="true"
                    :rows="10"
                    :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 50]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas"
                >
                    <template #header>
                        <div class="flex flex-wrap justify-end gap-2">
                            <Button text icon="pi pi-plus" label="Expand All" @click="expandAll" />
                            <Button text icon="pi pi-minus" label="Collapse All" @click="collapseAll" />
                        </div>
                    </template>
                    <Column expander style="width: 5rem" />
                    <Column field="nmKelas" header="Name"></Column>
                    <Column field="tingkatPendidikanId" header="Tingkat" sortable></Column>
                    <Column field="jurusan.namaJurusan" header="Jurusan" sortable></Column>
                    <Column field="ptk.nama" header="Wali Kelas"></Column>
                    <!-- <Column field="ptk.nama" header="Jml.Mapel"></Column> -->
                    <Column field="" header="Edit">
                        <template #body="{ data }">
                            <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editMapel(data)" />
                            <!-- <Button icon="pi pi-trash" outlined rounded severity="danger"
                                    @click="confirmdeleteMapel(data)" /> -->
                        </template>
                    </Column>
                    <template #expansion="slotProps">
                        <div class="p-4">
                            <DataTable :value="slotProps.data.pembelajaran">
                                <!-- <p>{{ slotProps.data.pembelajaran }}</p> -->
                                <Column field="namaMataPelajaran" header="Mata pelajaran" sortable></Column>
                                <Column field="ptkTerdaftar.ptk.nama" header="Guru Mapel" sortable></Column>
                                <!-- <Column field="date" header="Date" sortable></Column> -->
                            </DataTable>
                        </div>
                    </template>
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
                    <DataTable ref="dt" :value="pembelajaranList" dataKey="pembelajaran_id">
                        <Column field="nama_mata_pelajaran" header="Mata pelajaran"></Column>
                        <Column field="nama" header="Guru Mapel"></Column>
                    </DataTable>
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

        <!-- end of import data -->
    </div>
</template>
