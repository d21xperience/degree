<script setup>
import FileUpload from 'primevue/fileupload';
import InputText from 'primevue/inputtext';
import Tab from 'primevue/tab';
import TabList from 'primevue/tablist';
import TabPanel from 'primevue/tabpanel';
import TabPanels from 'primevue/tabpanels';
import Tabs from 'primevue/tabs';
import { computed, ref } from 'vue';
import { useStore } from 'vuex';
const store = useStore();

// ========================
// Props dari parent
const props = defineProps({
    visible: Boolean,
    pesertaDidik: Array
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

const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
]);

const alamatLengkap = ref({
    alamatJalan: '',
    rt: '',
    rw: '',
    desa: '',
    kec: '',
    kab: '',
    prov: ''
});

const selectedBCNetwork = computed(() => store.getters['scService/getBCPlatformSelected']);
</script>

<template>
    <div>
        <Tabs value="0">
            <TabList>
                <Tab value="0">Ijazah</Tab>
                <Tab value="1">Transkrip</Tab>
            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <div class="container bg-white p-2 shadow-md">
                        <section class="mb-2">
                            <h2 class="text-xl font-semibold mb-4">Informasi diri</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="nmSiswa">Nama</label>
                                    <InputText id="nmSiswa" v-model="pesertaDidik[0].pesertaDidik.nmSiswa" fluid name="nmSiswa" placeholder="Masukan nama" disabled="" />
                                </div>
                                <div class="w-full">
                                    <label class="block text-gray-700">Jenis Kelamin</label>
                                    <Select v-model="pesertaDidik[0].pesertaDidik.jenisKelamin" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin" option-label="label" option-value="value" class="w-full" />
                                </div>
                                <div>
                                    <div class="md:flex md:space-x-1">
                                        <div class="w-full">
                                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                                            <InputText id="tempatLahir" v-model="pesertaDidik[0].pesertaDidik.tempatLahir" fluid name="tempatLahir" placeholder="Masukan tempat lahir" class="w-full md:w-64" disabled />
                                        </div>
                                        <div>
                                            <label class="block text-gray-700">Tgl Lahir</label>
                                            <input v-model="pesertaDidik[0].pesertaDidik.tanggalLahir" type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" />
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nmWali">Nama Wali</label>
                                    <InputText id="nmSiswa" v-model="pesertaDidik[0].namaOrtuWali" fluid name="nmSiswa" placeholder="Masukan nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nis">NIS</label>
                                    <InputText id="nis" v-model="pesertaDidik[0].pesertaDidik.nis" fluid name="nis" placeholder="Masukan NIS" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nisn">NISN</label>
                                    <InputText id="nisn" v-model="pesertaDidik[0].pesertaDidik.nisn" fluid name="nisn" placeholder="Masukan NISN" />
                                </div>
                            </div>
                            <h2 class="text-xl font-semibold mb-4">Informasi Sekolah</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="prov">Prov.</label>
                                    <InputText id="prov" v-model="alamatLengkap.prov" fluid name="prov" placeholder="Masukan nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="kab">Kab</label>
                                    <InputText id="kab" v-model="alamatLengkap.kab" fluid name="kab" placeholder="Masukan nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="no-ijazah">No. Ijazah</label>
                                    <InputText id="no-ijazah" v-model="pesertaDidik[0].nomorIjazah" fluid name="no-ijazah" placeholder="Masukan no ijazah" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tpt-ijazah">Tempat Ijazah</label>
                                    <InputText id="tpt-ijazah" v-model="pesertaDidik[0].tempatIjazah" fluid name="tpt-ijazah" placeholder="Masukan nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tgl-dikelurkan-ijazah">Tgl dikeluarkan Ijazah</label>
                                    <input v-model="pesertaDidik[0].tanggalIjazah" type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" />
                                </div>
                            </div>
                            <h2 class="text-xl font-semibold mb-4">Informasi Blockhain</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="cid-uri">CID URL</label>
                                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".jpg" :max-file-size="2000000" :custom-upload="true" severity="secondary" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="blockhain-type">Blockhain Type</label>
                                    <InputText id="blockhain-type" v-model="selectedBCNetwork.name" fluid name="blockhain-type" placeholder="Masukan nama" disabled />
                                </div>
                            </div>
                        </section>
                        <div class="flex justify-end space-x-4">
                            <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600" @click="submitForm">Simpan</button>
                            <!-- <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400"
                            @click="closeDialog">Batal</button> -->
                        </div>
                    </div>
                </TabPanel>
                <TabPanel value="1">
                    <div>
                        <!-- <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="kelasList"
                        scrollable scrollHeight="400px" dataKey="rombonganBelajarId" :paginator="true" :rows="10"
                        :filters="filters" tableStyle="min-width: 50rem"
                        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                        :rowsPerPageOptions="[10, 20, 30]"
                        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-2">
                        <Column field="nmKelas" header="Nama Mapel"></Column>
                        <Column field="tingkatPendidikanId" header="SMT 1"></Column>
                        <Column field="kurikulum.namaKurikulum" header="SMT 2"></Column>

                        <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                            <Column field="namaJurusanSp" header="Jurusan"></Column>
                        </div>
                        <Column field="ptk.nama" header="SMT 3"></Column>
                        <Column field="jumlahAnggota" header="SMT 4"></Column>
                        <Column field="jumlahAnggota" header="SMT 5"></Column>
                        <Column field="jumlahAnggota" header="SMT 6"></Column>
                        <Column header="Rata-rata">
                            <template #body="slotProps">
                                <Button icon="pi pi-bullseye" outlined rounded class="mr-2"
                                    @click="dialogAnggotaRombel(slotProps.data)" />
                            </template>
</Column>

</DataTable> -->
                    </div>
                </TabPanel>
            </TabPanels>
        </Tabs>
    </div>
</template>
