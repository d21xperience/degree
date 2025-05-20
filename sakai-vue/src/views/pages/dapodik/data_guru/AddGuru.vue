<script setup>
import router from '@/router';
import InputText from 'primevue/inputtext';
import Select from 'primevue/select';
import Toast from 'primevue/toast';
import { onMounted, ref, watch } from 'vue';
// ================================
// composable
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
const { searchGuruTerdaftar, updateGuruTerdaftar } = useSekolahService();
// ================================
const { selectedJenisKelamin, jenisKelaminOptions, agamaOptions, selectedAgama, fetchGelarAkademik, selectedGelarAkademikDepan, selectedGelarAkademikBelakang, searchGelar, gelarAkademikDepanOptions, gelarAkademikBelakangOptions } = useFormOptions();

import { useRoute } from 'vue-router';
// Model Peserta Didik Pelengkap
const tabel_ptk_terdaftar = ref({
    ptk_terdaftar_id: '',
    ptk_id: '',
    tahun_ajaran_id: '',
    jenis_keluar_id: '',
    soft_delete: '',
    ptk: {
        ptk_id: '',
        nama: '',
        nip: '',
        jenis_ptk_id: 8,
        jenis_kelamin: '',
        tempat_lahir: '',
        tanggal_lahir: '',
        nuptk: '',
        alamat_jalan: '',
        rt: '',
        rw: '',
        desa_kel: '',
        kec: '',
        kab_kota: '',
        prop: '',
        status_keaktifan_id: 1,
        gelar_belakang: '',
        gelar_depan: ''
    },
    ptk_pelengkap: {
        ptk_pelengkap_id: '',
        ptk_id: '',
        gelar_depan: '',
        gelar_belakang: '',
        nip_niy: '',
        alamat_jalan: '',
        rt: '',
        rw: '',
        desa: '',
        kec: '',
        kab: '',
        prov: ''
    }
});
const route = useRoute();
const ptkTerdaftarId = route.query.ptkTerdaftarId;
const isEdit = ref(false);
onMounted(async () => {
    fetchGelarAkademik();
    if (ptkTerdaftarId) {
        isEdit.value = true;
        const ptkTerdaftar = await searchGuruTerdaftar(ptkTerdaftarId);
        // console.log(ptkTerdaftar);
        const guru = ptkTerdaftar[0].ptk;
        // console.log(guru.ptk);
        // await fetchGuru(ptkId);
        // if (guruList.value && guruList.value.length > 0) {
        //     const guru = guruList.value[0];
        if (guru.jenisKelamin === 'L') {
            selectedJenisKelamin.value = jenisKelaminOptions.value[0];
        } else if (guru.jenisKelamin === 'P') {
            selectedJenisKelamin.value = jenisKelaminOptions.value[1];
        }
        tabel_ptk_terdaftar.value.ptk = {
            ptk_id: guru.ptkId || '',
            nama: guru.nama || '',
            nip: guru.nip || '',
            jenis_ptk_id: guru.jenisPtkId || '',
            tempat_lahir: guru.tempatLahir || '',
            tanggal_lahir: guru.tanggalLahir || '',
            nuptk: guru.nuptk || '',
            alamat_jalan: guru.alamatJalan || '',
            status_keaktifan_id: guru.statusKeaktifanId || '',
            gelar_belakang: guru.gelarBelakang || '',
            gelar_depan: guru.gelarDepan || ''
        };

        //     tabel_ptk_terdaftar.ptk_id = guru.ptkId; // jangan lupa juga isi yang di luar
        // }
    }
});
// Opsi Dropdown
watch(selectedJenisKelamin, () => (tabel_ptk_terdaftar.value.ptk.jenis_kelamin = selectedJenisKelamin.value.value));
watch(selectedGelarAkademikBelakang, (newVal) => {
    tabel_ptk_terdaftar.value.ptk.gelar_belakang = newVal?.map((el) => el.kode).join(', ') || '';
});
watch(selectedGelarAkademikDepan, (newVal) => {
    tabel_ptk_terdaftar.value.ptk.gelar_depan = newVal?.map((el) => el.kode).join(', ') || '';
});

const isLoading = ref(false);
const submitForm = () => {
    submitted.value = true;
    isLoading.value = true;
    if (isEdit) {
        updateGuruTerdaftar(tabel_ptk_terdaftar);
    }
    setTimeout(() => {
        isLoading.value = false;
        router.push({ name: 'infoGuru' });
    }, 2000);
};

const batal = () => {
    router.push({ name: 'infoGuru' });
};
const submitted = ref(false);
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center mb-1">
            <h1 class="text-2xl font-bold mb-2">{{ isEdit ? 'Form Edit Guru' : 'Form Tambah Guru' }}</h1>
            <Button icon="pi pi-times" severity="danger" @click="batal" />
        </div>
        <section class="mb-2">
            <h2 class="text-xl font-normal mb-4">Informasi Guru</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 md:gap-4 mb-4">
                <div class="md:flex md:justify-between md:space-x-2 items-center">
                    <div class="w-full mb-2 md:mb-0">
                        Nama Lengkap
                        <!-- <label class="block font-normal text-sm" for="nmGuru">Nama Lengkap</label> -->
                        <InputText v-model="tabel_ptk_terdaftar.ptk.nama" fluid name="nmGuru" id="nmGuru" placeholder="Masukan nama" :invalid="submitted && !tabel_ptk_terdaftar.ptk.nama" />
                        <small v-if="submitted && !tabel_ptk_terdaftar.ptk.nama" class="text-red-500">Nama harus diisi.</small>
                    </div>
                </div>
                <div class="w-full">
                    Gelar Dpn
                    <AutoComplete v-model="selectedGelarAkademikDepan" :suggestions="gelarAkademikDepanOptions" optionLabel="kode" optionValue="kode" placeholder="Gelar dpn" dropdown multiple @complete="searchGelar(1, $event)" fluid />
                </div>
                <div class="w-full">
                    Gelar Blk
                    <AutoComplete v-model="selectedGelarAkademikBelakang" :suggestions="gelarAkademikBelakangOptions" optionLabel="kode" placeholder="Gelar blk" dropdown multiple @complete="searchGelar(2, $event)" fluid />
                </div>
                <div class="md:flex justify-between items-center md:space-x-2">
                    <div class="md:w-1/2">
                        Jenis Kelamin
                        <Select v-model="selectedJenisKelamin" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin" optionLabel="label" class="w-full" :invalid="submitted && !selectedJenisKelamin" />
                        <small v-if="submitted && !selectedJenisKelamin" class="text-red-500">Jenis kelamin harus diisi.</small>
                    </div>
                    <!-- <div class="md:w-1/2">
                        Agama
                        <Select v-model="selectedAgama" :options="agamaOptions" placeholder="Pilih Agama" optionLabel="label" class="w-full" :invalid="submitted && !selectedAgama" />
                    </div> -->
                </div>
                <div>
                    <div class="md:flex md:space-x-2">
                        <div class="md:w-[70%] w-full">
                            Tpt Lahir
                            <InputText v-model="tabel_ptk_terdaftar.ptk.tempat_lahir" fluid name="tempatLahir" id="tempatLahir" placeholder="Masukan tempat lahir" :invalid="submitted && !tabel_ptk_terdaftar.ptk.tempat_lahir" />
                            <small v-if="submitted && !tabel_ptk_terdaftar.ptk.tempat_lahir" class="text-red-500">Tpt.Lahir harus diisi.</small>
                        </div>
                        <div class="md:w-[30%] w-full">
                            Tgl Lahir
                            <input
                                type="date"
                                placeholder="YYYY-MM-DD"
                                class="w-full p-2 border border-gray-300 rounded"
                                v-model="tabel_ptk_terdaftar.ptk.tanggal_lahir"
                                :class="{ 'border-red-400': submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir, 'text-red-400': submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir }"
                            />
                            <small v-if="submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir" class="text-red-500">Tgl.Lahir harus diisi.</small>
                        </div>
                    </div>
                </div>
                <div class="md:flex md:space-x-2">
                    <div class="md:w-[80%]">
                        Alamat Jalan
                        <InputText v-model="tabel_ptk_terdaftar.ptk.alamat_jalan" fluid name="alamat-jalan" id="alamat-jalan" placeholder="Masukan alamat" />
                    </div>
                    <div class="md:flex md:space-x-2">
                        <div>
                            RT
                            <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.rt" fluid name="rt" id="rt" />
                        </div>
                        <div>
                            RW
                            <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.rw" fluid name="rw" id="rw" />
                        </div>
                    </div>
                </div>

                <div class="flex justify-between space-x-2">
                    <div>
                        Provinsi
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.prop" fluid name="prov" id="prov" />
                    </div>
                    <div>
                        Kab
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.kab" fluid name="kab" id="kab" />
                    </div>
                </div>
                <div class="flex justify-between space-x-2">
                    <div>
                        Kecamatan
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.kec" placeholder="Kecamatan" fluid name="kec" id="kec" />
                    </div>
                    <div>
                        Desa/Kota
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.desa" placeholder="Desa" fluid name="desa" id="desa" />
                    </div>
                </div>
                <!-- <div class="flex justify-between space-x-2">
                    <div>
                        No.Hp
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap." fluid name="phone-number" id="phone-number" />
                    </div>
                    <div>
                        Alamat email
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.rw" fluid name="email-address" id="email-address" />
                    </div>
                </div> -->
                <!-- <div class="flex justify-between space-x-2">
                    <div>
                        Tuga tambahanan
                        <InputText v-model="tabel_ptk_terdaftar.ptk.kab_kota" fluid />
                    </div>
                    <div>
                        Mengajar Mapel
                        <InputText v-model="tabel_ptk_terdaftar.ptk.kab_kota" fluid />
                    </div>
                </div> -->
            </div>
        </section>
        <div class="flex justify-end mt-8 space-x-4">
            <Button label="Simpan" severity="success" @click="submitForm" :loading="isLoading" class="min-w-28" />
            <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" />
        </div>
        <Toast />
    </div>
</template>

<style scoped>
label {
    font-weight: bold;
    display: block;
    margin-bottom: 0.5rem;
}

select option:disabled {
    color: gray;
    font-weight: bold;
}
</style>
