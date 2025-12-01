<script setup>
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import { useKategoriSekolah } from '@/composables/sekolah_composable/useKategoriSekolah';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useFormOptions } from '@/composables/useFormOptions';
import { onMounted, reactive, ref, watch } from 'vue';
import TingkatComponent from './TingkatComponent.vue';

// Services and composables
const { kurikulumList } = useFormOptions();
const { selectedSemester } = useSemester();
const { fetchKategoriSekolah, createKategoriSekolah, deleteKategoriSekolahKurikulum, createProsesKategoriSekolahDanKelas, updateKategoriSekolah, kategoriSekolahList, kategoriSekolahTabel } = useKategoriSekolah();
const { fetchSekolah, fetchTingkat } = useSekolah();
// Data
const sekolah = ref({ sekolah: {} });
const tingkat = ref([]);
// const kategoriSekolahTabel = ref([]);
const expandedRows = ref(null);

// UI State
const isEditKategoriSekolah = ref(false);
const isLoadingEditKategoriSekolah = ref(false);
const isDialogVisible = reactive({
    addKategoriSekolah: false,
    addKategoriKelas: false,
    deleteKategoriSekolah: false,
    deleteKategoriKelas: false,
    prosesKurikulum: false
});

// Form State
const formState = reactive({
    title: '',
    selectedKurikulum: null,
    selectedTingkat: null,
    selectedKategoriKelas: null,
    isEditMode: false,
    isSubmitDisabled: true,
    submitButtonText: 'Tambah'
});

const errorDialog = ref(false);
const messageDelete = ref('');
const tahunAjaran = ref(`${selectedSemester.value?.tahunAjaranId ?? 0}/${(selectedSemester.value?.tahunAjaranId ?? 0) + 1}`);
// Initialize component
const errorMessage = ref('');
const initComponent = async () => {
    try {
        await fetchKategoriSekolah();
        // console.log(kategoriSekolahList.value);
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value;
    } catch (error) {
        errorDialog.value = true;
        errorMessage.value = error;
        // console.error('Initialization error:', error);
        // Handle error (show toast/notification)
    }
};
const currentKategori = reactive({
    bidang_keahlian: null,
    program_keahlian: null,
    jurusan: null,
    jumlah: 0,
    tahun_ajaran_id: selectedSemester.value?.tahunAjaranId || '',
    tingkat_id: ''
});
// Watch for edit mode changes
watch(isEditKategoriSekolah, (newVal) => {
    isLoadingEditKategoriSekolah.value = true;

    if (newVal) {
        expandedRows.value = kategoriSekolahTabel.value.reduce((acc, p) => {
            acc[p.kurikulum_id] = true;
            return acc;
        }, {});
    } else {
        expandedRows.value = null;
    }

    isLoadingEditKategoriSekolah.value = false;
});

watch(
    () => currentKategori.jurusan,
    (newVal) => {
        // console.log('currentKategori.jurusan', newVal);
        if (newVal) {
            // console.log(kurikulumList);
            formState.selectedKurikulum = kurikulumList.value.find((item) => newVal.jurusanId.includes(item.jurusanId));
            // console.log(formState.selectedKurikulum)
            formState.isSubmitDisabled = false;
        }
    }
);

// Kategori Sekolah CRUD Operations
const openAddDialog = () => {
    formState.title = `Tambah Data T.A. ${tahunAjaran.value}`;
    isDialogVisible.addKategoriSekolah = true;
    // Reset currentKategori
    Object.assign(currentKategori, {
        bidang_keahlian: null,
        program_keahlian: null,
        jurusan: null,
        jumlah: 0,
        tahun_ajaran_id: selectedSemester.value?.tahunAjaranId || '',
        tingkat_id: ''
    });
};

const addKategoriSekolah = async () => {
    if (!formState.selectedKurikulum) {
        alert('Kurikulum harus dipilih!');
        return;
    }

    try {
        const newKategori = {
            kurikulum_id: formState.selectedKurikulum.kurikulumId,
            nama_kurikulum: formState.selectedKurikulum.namaKurikulum,
            jurusan_id: currentKategori.jurusan?.jurusanId || '',
            nama_jurusan: currentKategori.jurusan?.namaJurusan || '',
            nama_bidang_keahlian: currentKategori.bidang_keahlian?.namaJurusan || '',
            nama_program_keahlian: currentKategori.program_keahlian?.namaJurusan || '',
            jenjang_pendidikan_id: formState.selectedKurikulum.jenjangPendidikanId,
            tahun_ajaran_id: selectedSemester.value?.tahunAjaranId,
            jumlah: 0
        };
        await createKategoriSekolah(newKategori);
        isDialogVisible.addKategoriSekolah = false;
        // await fetchKategoriSekolah()
        kategoriSekolahList.value.push(newKategori);
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value;
    } catch (error) {
        console.error('Error adding kategori sekolah:', error);
        // Handle error
    }
};

const selectedItemToDelete = ref(null);
const handleDeleteKategoriSekolah = async (e) => {
    isDialogVisible.deleteKategoriSekolah = true;
    selectedItemToDelete.value = e;
    messageDelete.value = `Kurikulum <b>${e.nama_kurikulum}</b> akan dihapus?<br><br>Semua kelas pada kurikulum tersebut juga akan dihapus!`;
};
const deleteKategoriSekolah = async (item) => {
    try {
        await deleteKategoriSekolahKurikulum(item.kurikulum_id);
        kategoriSekolahList.value = kategoriSekolahList.value.filter((k) => k.kurikulum_id !== item.kurikulum_id);
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value;
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value.filter((k) => k.kurikulum_id !== item.kurikulum_id);
    } catch (error) {
        console.error('Error deleting kategori sekolah:', error);
        // Handle error
    } finally {
        isDialogVisible.deleteKategoriSekolah = false;
    }
};
const dialogAddKelas = (kategoriSekolah) => {
    formState.title = `Tambah kelas | ${kategoriSekolah.nama_jurusan}`;
    formState.selectedKategoriKelas = kategoriSekolah;

    // Reset form state
    formState.selectedTingkat = null;
    formState.isEditMode = false;
    currentKategori.jumlah = 0;
    currentKategori.tingkat_id = '';

    isDialogVisible.addKategoriKelas = true;
};

const addKelas = async () => {
    if (!formState.selectedTingkat) {
        alert('Tingkat harus dipilih!');
        return;
    }

    try {
        const newKelas = formState.selectedKategoriKelas;
        newKelas.tingkat_id = formState.selectedTingkat;
        newKelas.jumlah = currentKategori.jumlah;
        kategoriSekolahList.value.push(newKelas);
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value;
        await createKategoriSekolah(newKelas);

        isDialogVisible.addKategoriKelas = false;
        return;
    } catch (error) {
        console.error('Error adding kelas:', error);
        // Handle error
    }
};
const dialogEditKelas = (kategoriKelas) => {
    const cek = kategoriSekolahList.value.find((item) => kategoriKelas.kategori_sekolah_id == item.kategori_sekolah_id);
    formState.title = `Edit kelas | ${cek.nama_jurusan}`;
    formState.selectedKategoriKelas = cek;
    formState.isEditMode = true;
    formState.selectedTingkat = cek.tingkat_id;
    currentKategori.jumlah = cek.jumlah;
    currentKategori.tingkat_id = cek.tingkat_id;
    isDialogVisible.addKategoriKelas = true;
};

const updateKelas = async () => {
    if (!formState.selectedTingkat) {
        alert('Tingkat harus dipilih!');
        return;
    }
    try {
        const updatedKelas = {
            ...formState.selectedKategoriKelas,
            tingkat_id: formState.selectedTingkat,
            jumlah: currentKategori.jumlah
        };
        // Update di local state
        const index = kategoriSekolahList.value.findIndex((item) => item.kategori_sekolah_id === updatedKelas.kategori_sekolah_id);
        if (index !== -1) {
            kategoriSekolahList.value[index] = updatedKelas;
        }
        // Update tabel
        kategoriSekolahTabel.value = [...kategoriSekolahTabel.value];

        // Kirim ke API
        await updateKategoriSekolah(updatedKelas);

        // Tutup dialog
        isDialogVisible.addKategoriKelas = false;
    } catch (error) {
        console.error('Error updating kelas:', error);
    }
};

const handleDeleteKategoriKelas = async (e) => {
    isDialogVisible.deleteKategoriKelas = true;
    selectedItemToDelete.value = e;
    messageDelete.value = `Kelas <b>${e.tingkat_id}</b> akan dihapus?<br>`;
};
const deleteKelas = async (item) => {
    try {
        await deleteKategoriSekolah(item.kategori_sekolah_id);
        kategoriSekolahList.value = kategoriSekolahList.value.filter((k) => k.kategori_sekolah_id !== item.kategori_sekolah_id);
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value;
        // kategoriSekolahTabel.value = kategoriSekolahTabel.value.filter((k) => k.kurikulum_id !== item.kurikulum_id);
    } catch (error) {
        console.error('Error deleting kategori sekolah:', error);
        // Handle error
    } finally {
        isDialogVisible.deleteKategoriSekolah = false;
    }
};

const dialogProsesKurikulum = () => {
    messageDelete.value = `Anda akan memproses kurikulum? <br>Pastikan kurikulum sudah sesuai dengan Satuan Pendidikan yang akan diproses!`;
    isDialogVisible.prosesKurikulum = true;
};
const addProsesKurikulum = () => {
    createProsesKategoriSekolahDanKelas();
};
// Lifecycle
watch(selectedSemester, () => {
    initComponent();
});
onMounted(async () => {
    sekolah.value = await fetchSekolah();
    tingkat.value = await fetchTingkat();
    initComponent();
});

const isError = ref(false);
</script>

<template>
    <div class="kategori-sekolah-container">
        <div v-if="!isError" class="flex flex-col items-center justify-center p-6 bg-blue-50 border border-blue-200 rounded-lg shadow-sm text-center max-w-md mx-auto">
            <!-- <div class="mb-4">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-blue-500 mx-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
            </div> -->
            <p class="text-gray-700 mb-4">Data tidak ditemukan, <strong>Klik tambah untuk menambahkan</strong>.</p>
            <Button label="Tambah" />
        </div>
        <!-- Header Section -->
        <div v-else>
            <div class="flex justify-between mb-4">
                <div>
                    <Button v-show="isEditKategoriSekolah" v-tooltip.bottom="`Tambah kompetensi`" label="Tambah" icon="pi pi-plus" size="small" class="rounded-full" @click="openAddDialog" />
                </div>
                <div>
                    <Button v-tooltip.bottom="'Edit kategori'" icon="pi pi-pencil" severity="secondary" :loading="isLoadingEditKategoriSekolah" @click="isEditKategoriSekolah = !isEditKategoriSekolah" />
                </div>
            </div>

            <!-- Main Table -->
            <div class="grid grid-cols-1 gap-4">
                <DataTable v-model:expandedRows="expandedRows" :value="kategoriSekolahTabel" data-key="kurikulum_id" striped-rows>
                    <Column expander style="width: 2rem" />
                    <Column header="No" style="width: 2rem">
                        <template #body="slotProps">
                            {{ slotProps.index + 1 }}
                        </template>
                    </Column>
                    <Column header="Kurikulum" field="nama_kurikulum" style="width: 20rem" />
                    <Column header="Jurusan" field="nama_jurusan" />
                    <Column header="Total Kelas" field="total_kelas" />
                    <Column header="Tahun Ajaran" field="tahun_ajaran_id" />
                    <Column header="Aksi" :hidden="!isEditKategoriSekolah">
                        <template #body="slotProps">
                            <Button v-tooltip.bottom="'Hapus kompetensi'" icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" size="small" rounded @click="handleDeleteKategoriSekolah(slotProps.data)" />
                            <Button v-tooltip.bottom="'Tambah kelas'" icon="pi pi-plus" class="mr-2 !text-sm" severity="success" size="small" rounded @click="dialogAddKelas(slotProps.data)" />
                        </template>
                    </Column>

                    <!-- Expanded Row Content -->
                    <template #expansion="slotProps">
                        <DataTable :value="slotProps.data.kategorikelas">
                            <Column>
                                <!-- {{ slotProps.data }} -->
                            </Column>
                            <Column header="Kelas" field="tingkat_id" />
                            <Column header="Jml.Kelas" field="jumlah" />
                            <Column header="Aksi" :hidden="!isEditKategoriSekolah">
                                <!-- eslint-disable-next-line vue/no-template-shadow -->
                                <template #body="slotProps">
                                    <Button v-tooltip.bottom="'Hapus kelas'" icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" size="small" rounded @click="handleDeleteKategoriKelas(slotProps.data)" />
                                    <Button v-tooltip.bottom="'Edit kelas'" icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" size="small" rounded @click="dialogEditKelas(slotProps.data)" />
                                </template>
                            </Column>
                        </DataTable>
                    </template>
                </DataTable>
            </div>
            <div class="flex justify-end mt-10">
                <Button v-show="isEditKategoriSekolah" icon="pi pi-process" label="Proses kurikulum" severity="danger" @click="dialogProsesKurikulum" />
            </div>
        </div>
        <!-- Add Kategori Sekolah Dialog -->
        <Dialog v-model:visible="isDialogVisible.addKategoriSekolah" style="width: 550px" :header="formState.title" :modal="true">
            <div class="mb-4 md:grid space-y-2">
                <div v-show="sekolah.sekolah?.jenjangPendidikanId == 6">
                    <div>
                        <div class="font-semibold">Bidang Keahlian</div>
                        <BidangKeahlianComponent v-model="currentKategori.bidang_keahlian" />
                    </div>
                    <div class="my-2">
                        <div class="font-semibold">Program Keahlian</div>
                        <ProgramKeahlianComponent v-model="currentKategori.program_keahlian" :disabled="!currentKategori.bidang_keahlian" :jurusan-induk="currentKategori.bidang_keahlian?.jurusanId ?? ''" />
                    </div>
                    <div>
                        <div class="font-semibold">Kompetensi Keahlian/ Jurusan</div>
                        <JurusanComponent id="jurusan" v-model="currentKategori.jurusan" :disabled="!currentKategori.program_keahlian" :jurusan-induk="currentKategori.program_keahlian?.jurusanId ?? ''" />
                    </div>
                </div>
                <div class="">
                    <div class="font-semibold">Kurikulum</div>
                    <KurikulumComponent v-model="formState.selectedKurikulum" fluid />
                </div>
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <!-- <Button v-if="formState.isEditMode" @click="editKategoriSekolah" label="Update" severity="warn" class="w-20" :disabled="formState.isSubmitDisabled" /> -->
                <Button :label="formState.submitButtonText" severity="warn" :disabled="formState.isSubmitDisabled" @click="addKategoriSekolah" />
                <Button severity="help" label="Batal" class="w-20" @click="isDialogVisible.addKategoriSekolah = false" />
            </div>
        </Dialog>

        <!-- Edit Kategori Kelas Dialog -->
        <Dialog v-model:visible="isDialogVisible.addKategoriKelas" style="width: 550px" :header="formState.title" :modal="true">
            <div class="mb-4 md:grid space-y-2">
                <div>
                    <div class="font-semibold">Tingkat</div>
                    <TingkatComponent v-model="formState.selectedTingkat" :initial-value="currentKategori.tingkat_id" />
                </div>
                <div>
                    <div class="font-semibold">Jumlah Kelas</div>
                    <InputNumber v-model="currentKategori.jumlah" placeholder="Masukan jumlah kelas" class="block" fluid :min="1" />
                </div>
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <div>
                    <Button v-if="formState.isEditMode == true" label="Update" severity="warn" @click="updateKelas" />
                    <Button v-else label="Simpan" severity="warn" @click="addKelas" />
                </div>
                <Button severity="help" label="Batal" class="w-20" @click="isDialogVisible.addKategoriKelas = false" />
            </div>
        </Dialog>

        <!-- Delete Confirmation Dialog -->
        <DialogConfirmDelete v-model:visible="isDialogVisible.deleteKategoriSekolah" :message="messageDelete" @confirm="deleteKategoriSekolah(selectedItemToDelete)" />
        <DialogConfirmDelete v-model:visible="isDialogVisible.deleteKategoriKelas" :message="messageDelete" @confirm="deleteKelas(selectedItemToDelete)" />
        <DialogConfirmDelete v-model:visible="isDialogVisible.prosesKurikulum" :message="messageDelete" @confirm="addProsesKurikulum(selectedItemToDelete)" />
        <!-- <Dialog :visible="true">"hello"</Dialog> -->
        <Dialog v-model:visible="errorDialog" style="width: 550px" header="Error" :modal="true" position="top">
            <div>{{ errorMessage }}</div>
        </Dialog>
    </div>
</template>
