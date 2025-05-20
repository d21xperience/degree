<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import router from '@/router';
import InputText from 'primevue/inputtext';
import Tab from 'primevue/tab';
import TabList from 'primevue/tablist';
import TabPanel from 'primevue/tabpanel';
import TabPanels from 'primevue/tabpanels';
import Tabs from 'primevue/tabs';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
const store = useStore();
const { fetchSekolah, searchDns } = useSekolahService();
// ============toast============
import Toast from 'primevue/toast';
const route = useRoute();
// ========================
const pesertaDidikId = route.query.pesertaDidikId;
// const selectedjenisKelaminOptions = ref();
const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
]);

// const alamatLengkap = ref({
//     alamatJalan: '',
//     rt: '',
//     rw: '',
//     desa: '',
//     kec: '',
//     kab: '',
//     prov: ''
// });
const selectedBCNetwork = ref({});
// const pesertaDidik = ref();
const dns = ref({
    ID: '',
    pesertaDidikId: '',
    rombonganBelajarId: '',
    programKeahlian: '',
    paketKeahlian: '',
    sekolahId: '',
    npsn: '',
    kabupatenKota: '',
    provinsi: '',
    nama: '',
    tempatLahir: '',
    tanggalLahir: '',
    nis: '',
    nisn: '',
    namaOrtuWali: '',
    sekolahPenyelenggaraUjianUs: '',
    sekolahPenyelenggaraUjianUn: '',
    asalSekolah: '',
    nomorIjazah: '',
    tempatIjazah: '',
    tanggalIjazah: '',
    isComplete: '',
    tahunAjaranId: '',
    kelas: {
        rombonganBelajarId: '',
        sekolahId: '',
        semesterId: '',
        jurusanId: '',
        ptkId: '',
        nmKelas: '',
        tingkatPendidikanId: '',
        jenisRombel: '',
        namaJurusanSp: '',
        kurikulumId: '',
        anggotaKelas: '',
        pembelajaran: '',
        ptk: '',
        jurusan: '',
        kurikulum: '',
        tingkatPendidikan: '',
        jumlahAnggota: ''
    },
    siswa: {
        pesertaDidikId: '',
        nis: '',
        nisn: '',
        nmSiswa: '',
        tempatLahir: '',
        tanggalLahir: '',
        jenisKelamin: '',
        agama: '',
        alamatSiswa: '',
        teleponSiswa: '',
        diterimaTanggal: '',
        nmAyah: '',
        nmIbu: '',
        pekerjaanAyah: '',
        pekerjaanIbu: '',
        nmWali: '',
        pekerjaanWali: '',
        nik: ''
    }
});
onMounted(async () => {
    const cek = await searchDns(pesertaDidikId);
    if (cek) {
        dns.value.siswa = { ...cek.siswa };
        dns.value.kelas = { ...cek.kelas };
    }
});
const isLoading = ref(false);
const isLoadingSave = ref(false);
// const sekolah = computed(() => fetchSekolah());
// const nmSekolah = sekolah?.nama.toLowerCase().replace(/\s+/g, '');
const cancel = async () => {
    isLoading.value = true;
    const sekolah = await fetchSekolah();
    const nmSekolah = sekolah?.nama.toLowerCase().replace(/\s+/g, '');
    setTimeout(() => {
        isLoading.value = false;
        // router.push({ name: 'infoGuru' });
        router.push({ name: 'readIjazah', params: { sekolah: nmSekolah } });
    }, 500);
};
const save = async () => {
    isLoadingSave.value = true;
    const sekolah = await fetchSekolah();
    const nmSekolah = sekolah?.nama.toLowerCase().replace(/\s+/g, '');
    setTimeout(() => {
        isLoadingSave.value = false;
        // router.push({ name: 'infoGuru' });
        router.push({ name: 'readIjazah', params: { sekolah: nmSekolah } });
    }, 250);
};
const submit = ref(false);
</script>

<template>
    <div>
        <Toast />
        <Tabs value="0">
            <TabList>
                <Tab value="0">Ijazah</Tab>
                <Tab value="1">Transkrip</Tab>
            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <div class="container bg-white p-2 shadow-md">
                        <section class="mb-2">
                            <h2 class="text-xl font-normal mb-4">Informasi diri</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="nmSiswa">Nama</label>
                                    <InputText v-model="dns.siswa.nmSiswa" fluid name="nmSiswa" id="nmSiswa" placeholder="Diisi nama" disabled="" />
                                </div>
                                <div class="w-full">
                                    <label class="block text-gray-700">Jenis Kelamin</label>
                                    <Select v-model="dns.siswa.jenisKelamin" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin" optionLabel="label" option-value="value" class="w-full" />
                                </div>
                                <div>
                                    <div class="md:flex md:space-x-1">
                                        <div class="w-full">
                                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                                            <InputText v-model="dns.siswa.tempatLahir" fluid name="tempatLahir" id="tempatLahir" placeholder="Diisi tempat lahir" class="w-full md:w-64" disabled />
                                        </div>
                                        <div>
                                            <label class="block text-gray-700">Tgl Lahir</label>
                                            <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" v-model="dns.siswa.tanggalLahir" disabled />
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nmWali">Nama Ortu/Wali</label>
                                    <InputText v-model="dns.siswa.nmAyah" fluid name="nmWali" id="nmWali" placeholder="Diisi nama Wali" :invalid="!dns.siswa.nmAyah" />
                                    <small v-if="!dns.siswa.nmAyah" class="text-red-500">Nama Ortu/Wali harus diisi.</small>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nis">NIS</label>
                                    <InputText v-model="dns.siswa.nis" fluid name="nis" id="nis" placeholder="Diisi NIS" disabled />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nisn">NISN</label>
                                    <InputText v-model="dns.siswa.nisn" fluid name="nisn" id="nisn" placeholder="Diisi NISN" disabled />
                                </div>
                            </div>
                            <h2 class="text-xl font-normal mb-4">Informasi Pengisian Ijazah</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <!-- untuk SMK -->
                                <div>
                                    <label class="block text-gray-700" for="jurusan">Program Keahlian</label>
                                    <InputText v-model="dns.kelas.namaJurusanSp" fluid name="jurusan" id="jurusan" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="kab">Kompetensi Keahlian</label>
                                    <InputText fluid name="kab" id="kab" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="jurusan">Kab/Kota</label>
                                    <InputText fluid name="jurusan" id="jurusan" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="kab">Provinsi</label>
                                    <InputText fluid name="kab" id="kab" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="no-ijazah">No. Ijazah</label>
                                    <InputText v-model="dns.nomorIjazah" fluid name="no-ijazah" id="no-ijazah" placeholder="Contoh: M-SMK/K13-3/23/0000001" :invalid="!dns.nomorIjazah" />
                                    <small v-if="!dns.nomorIjazah" class="text-red-500">No Ijazah harus diisi.</small>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tpt-ijazah">Tempat Penerbitan</label>
                                    <InputText v-model="dns.tempatIjazah" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tgl-dikelurkan-ijazah">Tgl Penerbitan</label>
                                    <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" v-model="dns.tanggalIjazah" />
                                </div>
                            </div>
                            <!-- <h2 class="text-xl font-normal mb-4">Informasi Blockhain</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="cid-uri">CID URL</label>
                                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".jpg" :maxFileSize="2000000" :customUpload="true" severity="secondary" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="blockhain-type">Blockhain Type</label>
                                    <InputText v-model="selectedBCNetwork.name" fluid name="blockhain-type" id="blockhain-type" placeholder="Diisi nama" disabled />
                                </div>
                            </div> -->
                        </section>
                        <div class="flex justify-end space-x-4">
                            <Button class="w-32" @click="save" label="Simpan" :loading="isLoadingSave" />
                            <Button class="w-32" @click="cancel" label="Batal" severity="info" :loading="isLoading" />
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
