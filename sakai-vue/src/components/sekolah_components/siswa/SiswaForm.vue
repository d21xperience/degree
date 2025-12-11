<template>
    <div class="">
        <form @submit.prevent="onSubmit">
            <div class="grid gap-4">
                <section class="mb-8">
                    <h2 class="text-xl font-semibold mb-4">Informasi Siswa</h2>
                    <div class="md:grid grid-cols-2 grid-rows-5 gap-x-4 gap-y-2">
                        <div>
                            <label class="block text-gray-700" for="nmSiswa">Nama Lengkap</label>
                            <InputText id="nmSiswa" v-model="pesertaDidik.nmSiswa" fluid name="nmSiswa" placeholder="Masukan nama" :invalid="submitted && pesertaDidik.nmSiswa.trim().length == 0" />
                            <small v-if="submitted && pesertaDidik.nmSiswa.trim().length == 0" class="text-red-500">Nama harus diisi.</small>
                        </div>
                        <div class="col-start-1 row-start-2">
                            <label class="block text-gray-700">Jenis Kelamin</label>
                            <JKComponent v-model="pesertaDidik.jenisKelamin" />
                            <small v-if="submitted && !pesertaDidik.jenisKelamin" class="text-red-500">Jenis kelalmin harus diisi.</small>
                        </div>
                        <div class="col-start-1 row-start-3">
                            <div class="">
                                <label class="block text-gray-700" for="tempatLahir">Tempat Lahir</label>
                                <InputText id="tempatLahir" v-model="pesertaDidik.tempatLahir" fluid name="tempatLahir" placeholder="Masukan tempat lahir" class="w-full md:w-64" :invalid="submitted && pesertaDidik.tempatLahir.trim().length == 0" />
                                <small v-if="submitted && pesertaDidik.tempatLahir.trim().length == 0" class="text-red-500">Tempat lahir harus diisi.</small>
                            </div>
                        </div>
                        <div class="col-start-1 row-start-4">
                            <div>
                                <label class="block text-gray-700">Tanggal Lahir</label>
                                <input
                                    v-model="pesertaDidik.tanggalLahir"
                                    type="date"
                                    placeholder="YYYY-MM-DD"
                                    class="w-full p-2 border border-gray-300 rounded"
                                    :class="{ 'border-red-400': submitted && !pesertaDidik.tanggalLahir, 'text-red-400': submitted && !pesertaDidik.tanggalLahir }"
                                />
                                <small v-if="submitted && !pesertaDidik.tanggalLahir" class="text-red-500">Tgl.Lahir harus diisi.</small>
                            </div>
                        </div>
                        <div class="col-start-1 row-start-5">
                            <label class="block text-gray-700">Agama</label>
                            <AgamaComponent v-model="pesertaDidik.agama" />
                        </div>
                        <div class="row-span-5 col-start-2 row-start-1">
                            <label class="block text-gray-700">Foto Siswa</label>
                            <div class=" h-full content-around p-8">
                                <div>
                                    <div class="">
                                        <ImageDropZone @change="onImageSelected" />

                                        <!-- <pre class="mt-4 text-xs bg-gray-100 p-3 rounded"
                                            >{{ imageFile }}
    </pre
                                        > -->
                                        <!-- <input type="file" class="w-full p-2 border border-gray-300 rounded" />
                                        <i class="fas fa-upload absolute right-3 top-3 text-gray-400"></i> -->
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="mt-2 grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                        <div>
                            <label class="block text-gray-700" for="nis">NIS</label>
                            <InputText id="nis" v-model="pesertaDidik.nis" fluid name="nis" placeholder="Kosongkan jika belum ada" />
                        </div>
                        <div>
                            <label class="block text-gray-700" for="nisn">NISN</label>
                            <InputText id="nisn" v-model="pesertaDidik.nisn" fluid name="nisn" placeholder="Masukan NISN" />
                        </div>
                        <div>
                            <label class="block text-gray-700">Asal Sekolah</label>
                            <input type="text" placeholder="Masukan asal sekolah" class="w-full p-2 border border-gray-300 rounded" />
                        </div>
                        <div class="flex justify-between space-x-1">
                            <div class="w-full">
                                <label class="block text-gray-700">Tgl diterima di sekolah</label>
                                <input v-model="pesertaDidik.diterimaTanggal" type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" />
                            </div>
                            <div class="w-full">
                                <label class="block text-gray-700">Diterima di kelas</label>
                                <InputNumber placeholder="contoh: 10, 7" fluid />
                                <!-- <input type="number" placeholder="contoh: 10, 7" class="w-full p-2 border border-gray-300 rounded" /> -->
                            </div>
                        </div>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                        <div>
                            <label class="block text-gray-700" for="alamat-siswa">Alamat Jalan</label>
                            <InputText id="alamat-siswa" v-model="alamatLengkap.alamatJalan" fluid name="alamat-siswa" placeholder="Masukan nama" />
                        </div>
                        <div class="flex space-x-1">
                            <div class="w-1/2">
                                <label class="block text-gray-700" for="rt">RT</label>
                                <InputText id="rt" v-model="alamatLengkap.rt" fluid name="rt" placeholder="Masukan RT" />
                            </div>
                            <div class="w-1/2">
                                <label class="block text-gray-700" for="rw">RW</label>
                                <InputText id="rw" v-model="alamatLengkap.rw" fluid name="rw" placeholder="Masukan RW" />
                            </div>
                        </div>
                        <div>
                            <label class="block text-gray-700" for="prov">Prov.</label>
                            <InputText id="prov" v-model="alamatLengkap.prov" fluid name="prov" placeholder="Masukan nama" />
                        </div>
                        <div>
                            <label class="block text-gray-700" for="kab">Kab</label>
                            <InputText id="kab" v-model="alamatLengkap.kab" fluid name="kab" placeholder="Masukan nama" />
                        </div>
                        <div>
                            <label class="block text-gray-700" for="kec">Kecamatan</label>
                            <InputText id="kec" v-model="alamatLengkap.kec" fluid name="kec" placeholder="Masukan nama kecamatan" />
                        </div>
                        <div>
                            <label class="block text-gray-700" for="desa">Desa</label>
                            <InputText id="desa" v-model="alamatLengkap.desa" fluid name="desa" placeholder="Masukan nama desa" />
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
                            <InputText id="nmAyah" v-model="pesertaDidik.nmAyah" fluid name="nmAyah" placeholder="Masukan nama" />
                        </div>
                        <div>
                            <label class="block text-gray-700">Pekerjaan Ayah</label>
                            <InputText placeholder="Masukan perkerjaan Ayah" class="w-full p-2 border border-gray-300 rounded" />
                        </div>
                        <div>
                            <label class="block text-gray-700">Nama Ibu Kandung</label>
                            <InputText v-model="pesertaDidik.nmIbu" placeholder="Isi nama ibu kandung" fluid />
                        </div>
                        <div>
                            <label class="block text-gray-700">Pekerjaan Ibu</label>
                            <input type="text" placeholder="Masukan pekerjaan Ibu" class="w-full p-2 border border-gray-300 rounded" />
                        </div>
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
                    </div>
                </section>
            </div>

            <div class="flex justify-end gap-3 mt-6">
                <Button type="button" label="Cancel" severity="secondary" @click="dialogVisible = false" />
                <Button type="submit" :label="editingItem ? 'Update' : 'Save'" :loading="submitting" :disabled="submitting" />
            </div>
        </form>
    </div>
</template>
<script setup>
// Custom composable CRUD

// import { useSiswaCrud } from '@/composables/title2/Siswa/Crud';
import ImageDropZone from '@/components/ImageDropZone.vue';
import { computed, ref } from 'vue';

// State
const count = ref(0);
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
const imageFile = ref(null);
const onImageSelected = (file) => {
    imageFile.value = file;
    console.log('File diterima:', file);
};
</script>
