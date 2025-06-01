<script setup>
import { computed, ref } from 'vue';

import InputText from 'primevue/inputtext';
import Select from 'primevue/select';

import { useFormOptions } from '@/composables/useFormOptions'; 
const { selectedJenisKelamin, jenisKelaminOptions, agamaOptions, selectedAgama } = useFormOptions();

// const toast = useToast();

// import router from '@/router';
import { useRouter } from 'vue-router';

const router = useRouter();
const alamatLengkap = ref({
    alamatJalan: '',
    rt: '',
    rw: '',
    desa: '',
    kec: '',
    kab: '',
    prov: ''
});
// Model Peserta Didik
const pesertaDidik = ref({
    pesertaDidikId: '',
    nis: '',
    nisn: '',
    nmSiswa: '',
    tempatLahir: '',
    tanggalLahir: '',
    jenisKelamin: '',
    agama: '',
    alamatSiswa: computed(
        () => `${alamatLengkap.value.alamatJalan} RT.${alamatLengkap.value.rt} RW.${alamatLengkap.value.rw} Desa ${alamatLengkap.value.desa} Kec. ${alamatLengkap.value.kec} Kab. ${alamatLengkap.value.kab} Prov. ${alamatLengkap.value.prov}`
    ),
    teleponSiswa: '',
    diterimaTanggal: '',
    nmAyah: '',
    nmIbu: '',
    pekerjaanAyah: '',
    pekerjaanIbu: '',
    nmWali: '',
    pekerjaanWali: ''
});

// Model Peserta Didik Pelengkap
const pesertaDidikPelengkap = ref({
    pelengkapSiswaId: '',
    pesertaDidikId: '',
    statusDalamKel: '',
    anakKe: '',
    sekolahAsal: '',
    diterimaKelas: '',
    alamatOrtu: '',
    teleponOrtu: '',
    alamatWali: '',
    teleponWali: '',
    fotoSiswa: null
});

// Opsi Dropdown
const isValidate = () => {
    if (pesertaDidik.value.nmSiswa.trim().length == 0) {
        return false;
    } else {
        return true;
    }
};

const update = () => {};
// Handle Submit Form
const tambah = () => {
    submitted.value = true;
    if (!isValidate()) {
        alert('Data harus diisi!');
        return;
    }
    isLoadingTambah.value = true;
};

// Handle Upload Foto
const onUpload = (event) => {
    const file = event.files[0];
    pesertaDidikPelengkap.value.fotoSiswa = URL.createObjectURL(file);
    toast.add({ severity: 'info', summary: 'Foto Diunggah', detail: file.name, life: 3000 });
};

const batal = () => {
    // isLoadingBatal.value = true;
    router.push({ name: 'infoSiswa' });
};
const isEdit = ref(false);
const isLoadingBatal = ref(false);
const isLoadingTambah = ref(false);
const isLoadingUpdate = ref(false);
const submitted = ref(false);
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center mb-1">
            <h1 class="text-2xl font-bold mb-2">{{ isEdit ? 'Form Edit Siswa' : 'Form Tambah Siswa' }}</h1>
            <Button icon="pi pi-times" severity="danger" @click="batal" :loading="isLoadingBatal" rounded size="small" />
        </div>

        <section class="mb-8">
            <h2 class="text-xl font-normal mb-4">Informasi Siswa</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmSiswa">Nama Lengkap</label>
                    <InputText v-model="pesertaDidik.nmSiswa" fluid name="nmSiswa" id="nmSiswa" placeholder="Masukan nama" :invalid="submitted && pesertaDidik.nmSiswa.trim().length == 0" />
                    <small v-if="submitted && pesertaDidik.nmSiswa.trim().length == 0" class="text-red-500">Nama harus diisi.</small>
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Jenis Kelamin</label>
                    <Select v-model="pesertaDidik.jenisKelamin" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin" optionLabel="label" optionValue="value" fluid :invalid="submitted && !pesertaDidik.jenisKelamin" />
                    <small v-if="submitted && !pesertaDidik.jenisKelamin" class="text-red-500">Jenis kelalmin harus diisi.</small>
                </div>
                <div>
                    <div class="md:flex md:space-x-1">
                        <div class="w-full">
                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                            <InputText v-model="pesertaDidik.tempatLahir" fluid name="tempatLahir" id="tempatLahir" placeholder="Masukan tempat lahir" class="w-full md:w-64" :invalid="submitted && pesertaDidik.tempatLahir.trim().length == 0" />
                            <small v-if="submitted && pesertaDidik.tempatLahir.trim().length == 0" class="text-red-500">Tempat lahir harus diisi.</small>
                        </div>
                        <div>
                            <label class="block text-gray-700">Tgl Lahir</label>
                            <input
                                type="date"
                                placeholder="YYYY-MM-DD"
                                class="w-full p-2 border border-gray-300 rounded"
                                v-model="pesertaDidik.tanggalLahir"
                                :class="{ 'border-red-400': submitted && !pesertaDidik.tanggalLahir, 'text-red-400': submitted && !pesertaDidik.tanggalLahir }"
                            />
                            <small v-if="submitted && !pesertaDidik.tanggalLahir" class="text-red-500">Tgl.Lahir harus diisi.</small>
                        </div>
                    </div>
                </div>

                <div>
                    <label class="block text-gray-700">Agama</label>
                    <Select v-model="selectedAgama" :options="agamaOptions" placeholder="Pilih Agama" optionLabel="label" fluid class="w-full" />
                </div>
                <div>
                    <label class="block text-gray-700" for="nis">NIS</label>
                    <InputText v-model="pesertaDidik.nis" fluid name="nis" id="nis" placeholder="Masukan NIS" />
                </div>
                <div>
                    <label class="block text-gray-700" for="nisn">NISN</label>
                    <InputText v-model="pesertaDidik.nisn" fluid name="nisn" id="nisn" placeholder="Masukan NISN" />
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="alamat-siswa">Alamat Jalan</label>
                    <InputText v-model="alamatLengkap.alamatJalan" fluid name="alamat-siswa" id="alamat-siswa" placeholder="Masukan nama" />
                </div>
                <div class="flex space-x-1">
                    <div class="w-1/2">
                        <label class="block text-gray-700" for="rt">RT</label>
                        <InputText v-model="alamatLengkap.rt" fluid name="rt" id="rt" placeholder="Masukan RT" />
                    </div>
                    <div class="w-1/2">
                        <label class="block text-gray-700" for="rw">RW</label>
                        <InputText v-model="alamatLengkap.rw" fluid name="rw" id="rw" placeholder="Masukan RW" />
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700" for="prov">Prov.</label>
                    <InputText v-model="alamatLengkap.prov" fluid name="prov" id="prov" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block text-gray-700" for="kab">Kab</label>
                    <InputText v-model="alamatLengkap.kab" fluid name="kab" id="kab" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block text-gray-700" for="kec">Kecamatan</label>
                    <InputText v-model="alamatLengkap.kec" fluid name="kec" id="kec" placeholder="Masukan nama kecamatan" />
                </div>
                <div>
                    <label class="block text-gray-700" for="desa">Desa</label>
                    <InputText v-model="alamatLengkap.desa" fluid name="desa" id="desa" placeholder="Masukan nama desa" />
                </div>

                <!-- <div class="mb-4">
                    <label class="block text-gray-700">Address</label>
                    <textarea placeholder="Enter student's address"
                        class="w-full p-2 border border-gray-300 rounded"></textarea>
                </div> -->
            </div>
            <!-- <div class="mb-4">
                <label class="block text-gray-700">Phone Number</label>
                <div class="relative">
                    <input type="text" placeholder="Enter student's phone number" class="w-full p-2 border border-gray-300 rounded" />
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Email</label>
                <div class="relative">
                    <input type="text" placeholder="Enter student's phone number" class="w-full p-2 border border-gray-300 rounded" />
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div> -->
        </section>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Orang tua</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmAyah">Nama Ayah Kandung</label>
                    <InputText v-model="pesertaDidik.nmAyah" fluid name="nmAyah" id="nmAyah" placeholder="Masukan nama" />
                </div>
                <!-- <div>
                    <label class="block text-gray-700">Pekerjaan Ayah</label>
                    <InputText placeholder="Enter father's occupation" class="w-full p-2 border border-gray-300 rounded" />
                </div> -->
                <div>
                    <label class="block text-gray-700">Nama Ibu Kandung</label>
                    <InputText placeholder="Isi nama ibu kandung" v-model="pesertaDidik.nmIbu" fluid />
                </div>
                <!-- <div>
                    <label class="block text-gray-700">Pekerjaan Ibu</label>
                    <input type="text" placeholder="Enter mother's occupation" class="w-full p-2 border border-gray-300 rounded" />
                </div> -->
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Alamat Orang tua</label>
                <textarea placeholder="Enter parents' address (if different from student)" class="w-full p-2 border border-gray-300 rounded"></textarea>
            </div>
            <div>
                <label class="block text-gray-700">No.Tlp Ortu</label>
                <div class="relative">
                    <input type="text" placeholder="Enter parents' phone number" class="w-full p-2 border border-gray-300 rounded" />
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
        </section>

        <!-- <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Wali (Optional)</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Nama Wali</label>
                    <input type="text" placeholder="Enter guardian's name (if applicable)" class="w-full p-2 border border-gray-300 rounded" />
                </div>
                <div>
                    <label class="block text-gray-700">Pekerjaan wali</label>
                    <input type="text" placeholder="Enter guardian's occupation" class="w-full p-2 border border-gray-300 rounded" />
                </div>
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Alamat Wali</label>
                <textarea placeholder="Enter guardian's address" class="w-full p-2 border border-gray-300 rounded"></textarea>
            </div>
            <div>
                <label class="block text-gray-700">No.Tlp. Wali</label>
                <div class="relative">
                    <input type="text" placeholder="Enter guardian's phone number" class="w-full p-2 border border-gray-300 rounded" />
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
        </section> -->

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Tambahan</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Status dalam Keluarga</label>
                    <select class="w-full p-2 border border-gray-300 rounded">
                        <option>Select status</option>
                    </select>
                </div>
                <div>
                    <label class="block text-gray-700">Anak Ke-</label>
                    <input type="number" placeholder="Masukan anak ke (contoh: 1, 2)" class="w-full p-2 border border-gray-300 rounded" />
                </div>
                <div>
                    <label class="block text-gray-700">Asal Sekolah</label>
                    <input type="text" placeholder="Enter previous school name" class="w-full p-2 border border-gray-300 rounded" />
                </div>
                <div>
                    <label class="block text-gray-700">Tanggal diterima di sekolah</label>
                    <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" v-model="pesertaDidik.diterimaTanggal" />
                </div>
                <div>
                    <label class="block text-gray-700">Diterima di kelas</label>
                    <input type="number" placeholder="contoh: 10, 7" class="w-full p-2 border border-gray-300 rounded" />
                </div>
            </div>
            <div>
                <label class="block text-gray-700">Foto Siswa</label>
                <div class="relative">
                    <input type="file" class="w-full p-2 border border-gray-300 rounded" />
                    <i class="fas fa-upload absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
        </section>

        <div class="flex justify-end space-x-2">
            <Button label="Update" severity="success" @click="update" :loading="isLoadingUpdate" class="min-w-28" icon="pi pi-save" v-if="isEdit" />
            <Button label="Tambah" severity="success" @click="tambah" :loading="isLoadingTambah" class="min-w-28" icon="pi pi-save" v-else />
            <Button class="w-32" @click="batal" label="Batal" :loading="isLoadingBatal" />
        </div>
    </div>

    <Toast />
</template>

<style scoped>
/* label {
    font-weight: bold;
    display: block;
    margin-bottom: 0.5rem;
}

select option:disabled {
    color: gray;
    font-weight: bold;
} */
</style>
