<script setup>
import router from '@/router';
import { debounce } from 'lodash-es';
import InputText from 'primevue/inputtext';
import Select from 'primevue/select';
import Toast from 'primevue/toast';
import { onMounted, ref, watch } from 'vue';
// ================================
// composable
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
const { searchGuruTerdaftar, updateGuruTerdaftar, addGuruTerdaftar, initSelectedTahunAjaran } = useSekolahService();
// ================================
const { selectedJenisKelamin, jenisKelaminOptions, agamaOptions, selectedAgama, fetchGelarAkademik, selectedGelarAkademikDepan, selectedGelarAkademikBelakang, searchGelar, gelarAkademikDepanOptions, gelarAkademikBelakangOptions } = useFormOptions();

import { useRoute } from 'vue-router';
// Model Peserta Didik Pelengkap
const tabel_ptk_terdaftar = ref({
    ptk_terdaftar_id: '',
    ptk_id: '',
    tahun_ajaran_id: `${initSelectedTahunAjaran.value?.tahunAjaranId}`,
    jenis_keluar_id: '',
    ptk: {
        ptk_id: '',
        nama: '',
        jenis_kelamin: '',
        agama: '',
        tempat_lahir: '',
        tanggal_lahir: '',
        isDapodik: false,
        jenis_ptk_id: 4
    },
    ptk_pelengkap: {
        gelar_depan: '',
        gelar_belakang: '',
        nik: '',
        no_KK: '',
        nuptk: '',
        niy: '',
        nip: '',
        alamat_jalan: '',
        rt: '',
        rw: '',
        desa_kelurahanurahan: '',
        kec: '',
        kab_kota: '',
        propinsi: '',
        kode_pos: '',
        no_telepon_rumah: '',
        no_hp: '',
        email: ''
    }
});
const route = useRoute();
const ptkTerdaftarId = route.query.ptkTerdaftarId;
const isEdit = ref(false);
onMounted(async () => {
    if (ptkTerdaftarId) {
        isEdit.value = true;

        const ptkTerdaftar = await searchGuruTerdaftar(ptkTerdaftarId);
        const ptkPelengkap = ptkTerdaftar.ptkPelengkap;
        const guru = ptkTerdaftar.ptk;
        if (guru.jenisKelamin === 'L') {
            selectedJenisKelamin.value = jenisKelaminOptions.value[0];
        } else if (guru.jenisKelamin === 'P') {
            selectedJenisKelamin.value = jenisKelaminOptions.value[1];
        }
        tabel_ptk_terdaftar.value.ptk = {
            ptk_id: guru.ptkId || '',
            nama: guru.nama || '',
            tanggal_lahir: guru.tanggalLahir || '',
            tempat_lahir: guru.tempatLahir || '',
            agama: guru.agama || '',
            jenis_kelamin: guru.jenisKelamin
        };
        tabel_ptk_terdaftar.value.ptk_pelengkap = {
            gelar_depan: ptkPelengkap.gelarDepan,
            gelar_belakang: ptkPelengkap.gelarDepan,
            alamat_jalan: ptkPelengkap.alamatJalan,
            rt: ptkPelengkap.rt,
            rw: ptkPelengkap.rw,
            desa_kelurahan: ptkPelengkap.desaKelurahan,
            kab_kota: ptkPelengkap.kabKota,
            kec: ptkPelengkap.kec,
            propinsi: ptkPelengkap.propinsi,
            nuptk: ptkPelengkap.nuptk
        };
        // selectedGelarAkademikBelakang.value = ptkPelengkap.gelarDepan;

        //     tabel_ptk_terdaftar.ptk_id = guru.ptkId; // jangan lupa juga isi yang di luar
        // }
    }
    fetchGelarAkademik();
});
// Opsi Dropdown
// watch(selectedJenisKelamin, () => (tabel_ptk_terdaftar.value.ptk.jenis_kelamin = selectedJenisKelamin.value.value));
watch(selectedGelarAkademikBelakang, (newVal) => {
    tabel_ptk_terdaftar.value.ptk_pelengkap.gelar_belakang = newVal?.map((el) => el.kode).join(', ') || '';
});
watch(selectedGelarAkademikDepan, (newVal) => {
    tabel_ptk_terdaftar.value.ptk_pelengkap.gelar_depan = newVal?.map((el) => el.kode).join(', ') || '';
});

const loadingUpdate = ref(false);
const loadingTambah = ref(false);
const update = debounce(async () => {
    submitted.value = true;
    loadingUpdate.value = true;
    if (isValidate()) {
        alert('Data harus diisi!');
        return;
    }

    await updateGuruTerdaftar(tabel_ptk_terdaftar);
    // setTimeout(() => {
    loadingUpdate.value = false;
    router.push({ name: 'infoGuru' });
    // }, 2000);
}, 250);

const isValidate = () => {
    if (
        !tabel_ptk_terdaftar.value.ptk.jenis_kelamin ||
        !tabel_ptk_terdaftar.value.ptk.agama ||
        !tabel_ptk_terdaftar.value.ptk.tanggal_lahir ||
        tabel_ptk_terdaftar.value.ptk.nama.trim().length == 0 ||
        tabel_ptk_terdaftar.value.ptk.tempat_lahir.trim().length == 0
    ) {
        return false;
    } else {
        return true;
    }
};

const tambah = debounce(async () => {
    submitted.value = true;
    if (!isValidate()) {
        alert('Data harus diisi!');
        return;
    }
    loadingTambah.value = true; // return;
    const tes = await addGuruTerdaftar(tabel_ptk_terdaftar);
    if (tes) {
        loadingTambah.value = false;
        router.push({ name: 'infoGuru' });
    }
}, 250);
const batal = () => {
    loadingBatal.value = true;
    router.push({ name: 'infoGuru' });
};
const submitted = ref(false);
const loadingBatal = ref(false);
const ptkModel = ref();

const messageInfo = ref('');
const isDialogMessageInfo = ref(false);
watch(ptkModel, (newVal) => {
    // console.log(typeof newVal);
    if (typeof newVal == 'object') {
        messageInfo.value = `PTK a.n. ${newVal.nama} sudah ada di sistem! Apakah ptk tersebut akan ditambahkan untuk tahun pelajaran ini?`;
        isDialogMessageInfo.value = true;
    } else if (typeof newVal == 'string') {
        tabel_ptk_terdaftar.value.ptk.nama = newVal || '';
    }
    // }
});
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center mb-1">
            <h1 class="text-2xl font-bold mb-2">{{ isEdit ? 'Form Edit Guru' : 'Form Tambah Guru' }}</h1>
            <Button icon="pi pi-times" severity="danger" @click="batal" :loading="loadingBatal" rounded size="small" v-tooltip.bottom="'Batal'" />
        </div>
        <section class="mb-2">
            <h2 class="text-xl font-normal mb-4">Informasi Guru</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 md:gap-4 mb-4">
                <div class="md:flex md:justify-between md:space-x-2 items-center">
                    <div class="w-full mb-2 md:mb-0">
                        Nama Lengkap
                        <div v-if="isEdit">
                            <InputText v-model="tabel_ptk_terdaftar.ptk.nama" fluid name="nmGuru" id="nmGuru" placeholder="Masukan nama" :invalid="submitted && !tabel_ptk_terdaftar.ptk.nama" />
                            <small v-if="submitted && !tabel_ptk_terdaftar.ptk.nama" class="text-red-500">Nama harus diisi.</small>
                        </div>
                        <ptk-component v-model="ptkModel" v-else :invalid="submitted && !ptkModel" />
                        <small v-if="submitted && !ptkModel" class="text-red-500">Nama harus diisi.</small>
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
                    <div class="w-full">
                        Jenis Kelamin <small v-if="submitted && !tabel_ptk_terdaftar.ptk.jenis_kelamin" class="text-red-500">Jenis kelamin harus diisi.</small>
                        <Select
                            v-model="tabel_ptk_terdaftar.ptk.jenis_kelamin"
                            :options="jenisKelaminOptions"
                            optionLabel="label"
                            option-value="value"
                            class="w-full"
                            placeholder="Pilih jenis kelamin"
                            :invalid="submitted && !tabel_ptk_terdaftar.ptk.jenis_kelamin"
                            fluid
                        />
                    </div>
                    <div class="w-full">
                        Agama <small v-if="submitted && !tabel_ptk_terdaftar.ptk.agama" class="text-red-500">Agama harus diisi.</small>
                        <Select v-model="tabel_ptk_terdaftar.ptk.agama" :options="agamaOptions" placeholder="Pilih Agama" optionLabel="label" class="w-full" :invalid="submitted && !tabel_ptk_terdaftar.ptk.agama" optionValue="value" />
                    </div>
                </div>
                <div>
                    <div class="md:flex md:space-x-2">
                        <div class="md:w-[70%] w-full">
                            Tpt Lahir
                            <InputText v-model="tabel_ptk_terdaftar.ptk.tempat_lahir" fluid name="tempatLahir" id="tempatLahir" placeholder="Masukan tempat lahir" :invalid="submitted && tabel_ptk_terdaftar.ptk.tempat_lahir.trim().length == 0" />
                            <small v-if="submitted && tabel_ptk_terdaftar.ptk.tempat_lahir.trim().length == 0" class="text-red-500">Tpt.Lahir harus diisi.</small>
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
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.alamat_jalan" fluid name="alamat-jalan" id="alamat-jalan" placeholder="Masukan alamat" />
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
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.propinsi" fluid name="prov" id="prov" />
                    </div>
                    <div>
                        Kab/Kota
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.kab_kota" fluid name="kab" id="kab" />
                    </div>
                </div>
                <div class="flex justify-between space-x-2">
                    <div>
                        Kecamatan
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.kec" placeholder="Kecamatan" fluid name="kec" id="kec" />
                    </div>
                    <div>
                        Desa/Kel
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.desa_kelurahan" placeholder="Desa" fluid name="desa" id="desa" />
                    </div>
                </div>
                <div class="flex justify-between space-x-2">
                    <div>
                        NUPTK
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.nuptk" fluid name="nuptk" id="nuptk" />
                    </div>
                    <div>
                        NIP
                        <InputText v-model="tabel_ptk_terdaftar.ptk_pelengkap.nip" fluid name="nip" id="nip" />
                    </div>
                </div>
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
            <Button label="Update" severity="success" @click="update" :loading="loadingUpdate" class="min-w-28" icon="pi pi-save" v-if="isEdit" />
            <Button label="Tambah" severity="success" @click="tambah" :loading="loadingTambah" class="min-w-28" icon="pi pi-save" v-else />
            <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" />
        </div>
        <Toast />

        <Dialog v-model:visible="isDialogMessageInfo" style="width: 400px" modal>
            <div><span></span>{{ messageInfo }}</div>
            <div class="flex justify-self-end mt-10">
                <Button icon="pi pi-check" label="Ya" severity="info" class="w-32 mr-2" />
                <Button icon="pi pi-times" label="Tidak" @click="isDialogMessageInfo = false" class="w-32" />
            </div>
        </Dialog>
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
