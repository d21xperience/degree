<script setup>
import IpfsInterface from '@/components/scComponent/IpfsInterface.vue';
import router from '@/router';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

// ============toast============
import { useDns } from '@/composables/sekolah_composable/useDns';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import Toast from 'primevue/toast';
const route = useRoute();
const { fetchSekolah } = useSekolah();
const { updateDns, searchDnsLokal } = useDns();
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
    jenisKelamin: '',
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
        nama: '',
        tempatLahir: '',
        tanggalLahir: '',
        jenisKelamin: '',
        agama: '',
        alamatSiswa: '',
        teleponSiswa: '',
        diterimaTanggal: '',
        namaOrtuWali: '',
        nmIbu: '',
        pekerjaanAyah: '',
        pekerjaanIbu: '',
        nmWali: '',
        pekerjaanWali: '',
        nik: ''
    }
});
const isLoading = ref(false);
const isLoadingSave = ref(false);
// const sekolah = computed(() => fetchSekolah());
// const nmSekolah = sekolah?.nama.toLowerCase().replace(/\s+/g, '');
const cancel = async () => {
    isLoading.value = true;
    const sekolah = await fetchSekolah();
    const nmSekolah = sekolah.sekolah?.nama.toLowerCase().replace(/\s+/g, '');
    setTimeout(() => {
        isLoading.value = false;
        // router.push({ name: 'infoGuru' });
        router.push({ name: 'readIjazah', params: { sekolah: nmSekolah } });
    }, 500);
};
const save = async () => {
    isLoadingSave.value = true;
    // kirim ke server
    const response = await updateDns(dns.value);
    if (response) {
        isLoading.value = false;
        // return
        const sekolah = await fetchSekolah();
        const nmSekolah = sekolah.sekolah?.nama.toLowerCase().replace(/\s+/g, '');
        setTimeout(() => {
            isLoadingSave.value = false;
            // router.push({ name: 'infoGuru' });
            router.push({ name: 'readIjazah', params: { sekolah: nmSekolah } });
        }, 250);
    }
};

onMounted(async () => {
    const cek = await searchDnsLokal(pesertaDidikId);
    console.log(cek);
    console.log(cek.jenisKelamin);
    if (cek) {
        dns.value = { ...cek };
        // dns.value.jenisKelamin = jenisKelaminOptions.value.filter(item => item.value == cek.jenisKelamin)
    }
    // console.log(dns.value);
});
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
                                    <label class="block text-gray-700" for="nama">Nama</label>
                                    <InputText id="nama" v-model="dns.nama" fluid name="nama" placeholder="Diisi nama" disabled />
                                </div>
                                <div class="w-full">
                                    <label class="block text-gray-700">Jenis Kelamin</label>
                                    <Select v-model="dns.jenisKelamin" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin" option-label="label" option-value="value" class="w-full" disabled />
                                </div>
                                <div>
                                    <div class="md:flex md:space-x-1">
                                        <div class="w-full">
                                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                                            <InputText id="tempatLahir" v-model="dns.tempatLahir" fluid name="tempatLahir" placeholder="Diisi tempat lahir" class="w-full md:w-64" disabled />
                                        </div>
                                        <div>
                                            <label class="block text-gray-700">Tgl Lahir</label>
                                            <input v-model="dns.tanggalLahir" type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" disabled />
                                        </div>
                                    </div>
                                </div>
                                <div class="flex space-x-2">
                                    <div class="w-full">
                                        <label class="block text-gray-700" for="nis">NIS</label>
                                        <InputText id="nis" v-model="dns.nis" fluid name="nis" placeholder="Diisi NIS" disabled />
                                    </div>
                                    <div class="w-full">
                                        <label class="block text-gray-700" for="nisn">NISN</label>
                                        <InputText id="nisn" v-model="dns.nisn" fluid name="nisn" placeholder="Diisi NISN" disabled />
                                    </div>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="jurusan">Program Keahlian</label>
                                    <InputText id="jurusan" v-model="dns.programKeahlian" fluid name="jurusan" placeholder="Diisi nama" disabled />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="kab">Kompetensi Keahlian</label>
                                    <InputText id="kab" v-model="dns.paketKeahlian" fluid name="kab" placeholder="Diisi nama" disabled />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="jurusan">Kab/Kota</label>
                                    <InputText id="jurusan" v-model="dns.kabupatenKota" fluid name="jurusan" placeholder="Diisi nama" disabled />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="kab">Provinsi</label>
                                    <InputText id="kab" v-model="dns.provinsi" fluid name="kab" placeholder="Diisi nama" disabled />
                                </div>
                            </div>
                            <h2 class="text-xl font-normal mb-4">Informasi Pengisian Ijazah</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="no-ijazah">No. Ijazah</label>
                                    <InputText id="no-ijazah" v-model="dns.nomorIjazah" fluid name="no-ijazah" placeholder="Contoh: M-SMK/K13-3/23/0000001" :invalid="!dns.nomorIjazah" />
                                    <small v-if="!dns.nomorIjazah" class="text-red-500">No Ijazah harus diisi.</small>
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tpt-ijazah">Tempat Penerbitan</label>
                                    <InputText id="tpt-ijazah" v-model="dns.tempatIjazah" fluid name="tpt-ijazah" placeholder="Diisi nama" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="tgl-dikelurkan-ijazah">Tgl Penerbitan</label>
                                    <input v-model="dns.tanggalIjazah" type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" />
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="nmWali">Nama Ortu/Wali</label>
                                    <InputText id="nmWali" v-model="dns.namaOrtuWali" fluid name="nmWali" placeholder="Diisi nama Wali" :invalid="!dns.namaOrtuWali" />
                                    <small v-if="!dns.namaOrtuWali" class="text-red-500">Nama Ortu/Wali harus diisi.</small>
                                </div>
                            </div>
                            <h2 class="text-xl font-normal mb-4">Informasi Blockhain</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                                <div>
                                    <label class="block text-gray-700" for="cid-uri">CID URL</label>
                                    <IpfsInterface />
                                    <!-- <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".jpg" :maxFileSize="2000000" :customUpload="true" severity="secondary" /> -->
                                </div>
                                <div>
                                    <label class="block text-gray-700" for="blockhain-type">Blockhain Type</label>
                                    <InputText id="blockhain-type" v-model="selectedBCNetwork.name" fluid name="blockhain-type" placeholder="Diisi nama" disabled />
                                </div>
                            </div>
                        </section>
                        <div class="flex justify-end space-x-4">
                            <Button class="w-32" label="Update" :loading="isLoadingSave" @click="save" />
                            <Button class="w-32" label="Batal" severity="info" :loading="isLoading" @click="cancel" />
                        </div>
                    </div>
                </TabPanel>
                <TabPanel value="1">
                    <div>
                        <h1>Hello world!</h1>
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
                            <Column field="nmKelas" header="Nama Mapel" />
                            <Column field="tingkatPendidikanId" header="SMT 1" />
                            <Column field="kurikulum.namaKurikulum" header="SMT 2" />

                            <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                                <Column field="namaJurusanSp" header="Jurusan" />
                            </div>
                            <Column field="ptk.nama" header="SMT 3" />
                            <Column field="jumlahAnggota" header="SMT 4" />
                            <Column field="jumlahAnggota" header="SMT 5" />
                            <Column field="jumlahAnggota" header="SMT 6" />
                            <Column header="Rata-rata">
                                <template #body="slotProps">
                                    <Button icon="pi pi-bullseye" outlined rounded class="mr-2" @click="dialogAnggotaRombel(slotProps.data)" />
                                </template>
                            </Column>
                        </DataTable>
                    </div>
                </TabPanel>
            </TabPanels>
        </Tabs>
    </div>
</template>
