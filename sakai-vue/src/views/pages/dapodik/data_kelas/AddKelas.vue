<script setup>
import router from '@/router';
import { debounce } from 'lodash-es';
import { computed, onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();

const kelasId = route.query.kelasId;

import JurusanComponent from '@/components/JurusanComponent.vue';
import KurikulumSekolahComponent from '@/components/KurikulumSekolahComponent.vue';
import TingkatComponent from '@/components/TingkatComponent.vue';
import { useSekolahService } from '@/composables/useSekolahService';
import { useStore } from 'vuex';
const store = useStore();

const isEdit = ref(false);

const selectedTingkat = ref(null);
const tingkatPendidikanOptions = ref([]);
const submitted = ref(false);
// Model Kelas
const rombel = ref({
    rombongan_belajar_id: '',
    sekolah_id: '',
    semester_id: '',
    jurusan_id: '',
    ptk_id: '',
    nm_kelas: '',
    tingkat_pendidikan_id: '',
    jenis_rombel: '',
    nama_jurusan_sp: '',
    // jurusan_sp_id: '',
    kurikulum_id: ''
});

// ===================================
// KURIKULUM
// ===================================
const kurikulumTerpilih = ref(null);
const jurusanTerpilih = ref(null);
const tingkatTerpilih = ref(null);
const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
const { fetchKelas, kelasList, fetchGuruTerdaftar, guruTerdaftarList, addKelas } = useSekolahService();
// ================================
const selectedGuruTerdaftar = ref(null);
watch(kurikulumTerpilih, (newVal) => {
    // console.log(newVal);
    jurusanTerpilih.value = newVal.namaJurusan;
    rombel.value.kurikulum_id = newVal.kurikulumId;
    rombel.value.jurusan_id = newVal.jurusanId;
    rombel.value.semester_id = `${newVal.tahunAjaranId}`;
    rombel.value.nama_jurusan_sp = newVal.namaJurusan;
    rombel.value.jenis_rombel = 1;
});

onMounted(async () => {
    if (kelasId) {
        isEdit.value = true;
        await fetchKelas(kelasId);
        if (kelasList.value && kelasList.value.length > 0) {
            const kelas = kelasList.value[0];
            rombel.value = {
                rombonganBelajarId: kelas.rombonganBelajarId || '',
                sekolahId: kelas.sekolahId || '',
                semesterId: kelas.semesterId || '',
                jurusanId: kelas.jurusanId || '',
                ptkId: kelas.ptkId || '',
                nmKelas: kelas.nmKelas || '',
                tingkatPendidikanId: kelas.tingkatPendidikanId || '',
                jenisRombel: kelas.jenisRombel || '',
                namaJurusanSp: kelas.namaJurusan || '',
                jurusan_sp_id: kelas.namaJurusanSpId || '',
                kurikulumId: kelas.kurikulumId || ''
                // tingkatPendidikan: null,
                // kurikulum: null,
                // jurusan: null,
            };
            selectedTingkat.value = tingkatPendidikanOptions.value.find((item) => item.tingkatPendidikanId === kelas.tingkatPendidikanId);
            selectedJurusan.value = jurusanList.value.find((item) => item.jurusanId === kelas.jurusanId);
            selectedGuruTerdaftar.value = guruTerdaftarList.value.find((item) => item.ptk.ptkId === kelas.ptk.ptkId);
        }
    }
});

const isLoadingTambah = ref(false);
const isLoadingCancel = ref(false);
const isValidate = () => {
    if (!rombel.value.tingkat_pendidikan_id || rombel.value.nm_kelas.trim().length == 0) {
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
    isLoadingTambah.value = true; // return;
    const tes = await addKelas(rombel);
    if (tes) {
        isLoadingTambah.value = false;
        router.push({ name: 'infoKelas' });
    }
}, 250);
const batal = async () => {
    // await nextTick();
    router.push({ name: 'infoKelas' });
};
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center mb-1">
            <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Form Edit Kelas' : 'Form Tambah Kelas' }}</h1>
            <Button icon="pi pi-times" severity="danger" @click="batal" rounded size="small" :loading="isLoadingCancel" v-tooltip.bottom="'Batal'" />
        </div>
        <section class="mb-6">
            <h2 class="text-xl font-semibold mb-4">Informasi Kelas</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div class="flex space-x-1">
                    <div class="w-full">
                        <label class="block text-gray-700" for="nmKelas">Nama Kelas</label>
                        <InputText v-model="rombel.nm_kelas" fluid name="nmKelas" id="nmKelas" placeholder="contoh: TBSM A (tanpa tingkat)" :invalid="submitted && !rombel.nm_kelas" />
                        <small v-if="submitted && rombel.nm_kelas.trim().length == 0" class="text-red-500">Nama kelas harus diisi.</small>
                    </div>
                    <div class="w-full md:w-40">
                        <label class="block text-gray-700">Tingkat</label>
                        <tingkat-component v-model="rombel.tingkat_pendidikan_id" />
                        <!-- <Select v-model="selectedTingkat" :options="tingkatPendidikanOptions" placeholder="Pilih tingkat" optionLabel="nama" class="w-full" :invalid="submitted && !selectedTingkat" />
                        <small v-if="submitted && !selectedTingkat" class="text-red-500">Tingkat harus diisi.</small> -->
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700">Kurikulum</label>
                    <div class="relative">
                        <kurikulum-sekolah-component v-model="kurikulumTerpilih" />
                    </div>
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Jurusan</label>
                    <jurusan-component v-model="jurusanTerpilih" />
                </div>
                <!-- <div class="">
                    <label class="block text-gray-700">Wali kelas</label>
                    <div class="relative">
                        <Select v-model="selectedGuruTerdaftar" :options="guruTerdaftarList" placeholder="Pilih Wali kelas" optionLabel="ptk.nama" class="w-full" :invalid="submitted && !selectedGuruTerdaftar">
                            <template #footer>
                                <div class="p-3">
                                    <Button label="Tambah Guru" fluid severity="secondary" text size="small" icon="pi pi-plus" @click="addGuru" />
                                </div>
                            </template>
                        </Select>
                        <small v-if="submitted && !selectedGuruTerdaftar" class="text-red-500">Wali kelas harus diisi.</small>
                    </div>
                </div> -->
            </div>
        </section>

        <!-- <Toast /> -->
        <!-- Daftar anggota rombel -->
        <!-- <div v-show="isEdit"> -->
        <!--<h2 class="text-xl font-semibold mb-4">Anggota Kelas</h2>-->
        <!-- Anggota Kelas -->
        <!-- <AnggotaKelas :rombongan-belajar-id="kelasId" :is-edit="isEdit" /> -->
        <!-- End of Anggota Kelas -->
        <!-- </div> -->

        <div class="flex justify-end mt-8 space-x-4">
            <Button label="Simpan" severity="success" @click="tambah" :loading="isLoadingTambah" class="min-w-28" />
            <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" :loading="isLoadingCancel" />
        </div>
    </div>
</template>
