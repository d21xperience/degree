<template>
    <div class="kategori-sekolah-container">
        <!-- Header Section -->
        <div class="flex justify-between mb-4">
            <div>
                <Button label="Tambah" v-show="isEditKategoriSekolah" icon="pi pi-plus" @click="openAddDialog" size="small" v-tooltip.bottom="`Tambah kompetensi`" />
            </div>
            <div>
                <Button icon="pi pi-pencil" @click="isEditKategoriSekolah = !isEditKategoriSekolah" severity="secondary" v-tooltip.bottom="'Edit kategori'" :loading="isLoadingEditKategoriSekolah" />
            </div>
        </div>

        <!-- Main Table -->
        <div class="grid grid-cols-1 gap-4">
            <DataTable v-model:expandedRows="expandedRows" :value="kategoriSekolahTabel" dataKey="kurikulum_id" striped-rows>
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
                        <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="handleDeleteKategoriSekolah(slotProps.data)" size="small" rounded v-tooltip.bottom="'Hapus kompetensi'" />
                        <Button icon="pi pi-plus" class="mr-2 !text-sm" severity="success" @click="dialogAddKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Tambah kelas'" />
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
                            <template #body="slotProps">
                                <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="handleDeleteKategoriKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Hapus kelas'" />
                                <Button icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" @click="dialogEditKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Edit kelas'" />
                            </template>
                        </Column>
                    </DataTable>
                </template>
            </DataTable>
        </div>
        <div class="flex justify-end mt-10">
            <Button icon="pi pi-process" label="Proses kurikulum" severity="danger" @click="dialogProsesKurikulum" v-show="isEditKategoriSekolah"/>
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
                        <JurusanComponent v-model="currentKategori.jurusan" id="jurusan" :disabled="!currentKategori.program_keahlian" :jurusan-induk="currentKategori.program_keahlian?.jurusanId ?? ''" />
                    </div>
                </div>
                <div class="">
                    <div class="font-semibold">Kurikulum</div>
                    <KurikulumComponent v-model="formState.selectedKurikulum" fluid />
                </div>
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <Button v-if="formState.isEditMode" @click="editKategoriSekolah" label="Update" severity="warn" class="w-20" :disabled="formState.isSubmitDisabled" />
                <Button v-else @click="addKategoriSekolah" :label="formState.submitButtonText" severity="warn" :disabled="formState.isSubmitDisabled" />
                <Button @click="isDialogVisible.addKategoriSekolah = false" severity="help" label="Batal" class="w-20" />
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
                    <InputNumber placeholder="Masukan jumlah kelas" class="block" fluid v-model="currentKategori.jumlah" :min="1" />
                </div>
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <div>
                    <Button @click="updateKelas" label="Update" severity="warn" v-if="formState.isEditMode == true" />
                    <Button @click="addKelas" label="Simpan" severity="warn" v-else />
                </div>
                <Button @click="isDialogVisible.addKategoriKelas = false" severity="help" label="Batal" class="w-20" />
            </div>
        </Dialog>

        <!-- Delete Confirmation Dialog -->
        <DialogConfirmDelete :message="messageDelete" v-model:visible="isDialogVisible.deleteKategoriSekolah" @confirm="deleteKategoriSekolah(selectedItemToDelete)" />
        <DialogConfirmDelete :message="messageDelete" v-model:visible="isDialogVisible.deleteKategoriKelas" @confirm="deleteKelas(selectedItemToDelete)" />
        <DialogConfirmDelete :message="messageDelete" v-model:visible="isDialogVisible.prosesKurikulum" @confirm="addProsesKurikulum(selectedItemToDelete)" />
    </div>
</template>

<script setup>
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, reactive, ref, watch } from 'vue';
import TingkatComponent from './TingkatComponent.vue';

// Services and composables
const { kurikulumList } = useFormOptions();
const sekolahService = useSekolahService();

// Data
const sekolah = ref({ sekolah: {} });
const tingkat = ref([]);
const kategoriSekolahTabel = ref([]);
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

const messageDelete = ref('');
const tahunAjaran = ref(`${sekolahService.initSelectedSemester.value?.tahunAjaranId}/${sekolahService.initSelectedSemester.value?.tahunAjaranId + 1}`);
// Initialize component
const initComponent = async () => {
    try {
        await sekolahService.fetchKategoriSekolah();
        // console.log(sekolahService.kategoriSekolahList.value);
        kategoriSekolahTabel.value = sekolahService.kategoriSekolahTabel.value;
    } catch (error) {
        console.error('Initialization error:', error);
        // Handle error (show toast/notification)
    }
};
const currentKategori = reactive({
    bidang_keahlian: null,
    program_keahlian: null,
    jurusan: null,
    jumlah: 0,
    tahun_ajaran_id: sekolahService.initSelectedSemester.value?.tahunAjaranId || '',
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
            formState.selectedKurikulum = kurikulumList.find((item) => newVal.jurusanId.includes(item.jurusanId));
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
        tahun_ajaran_id: sekolahService.initSelectedSemester.value?.tahunAjaranId || '',
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
            tahun_ajaran_id: sekolahService.initSelectedSemester.value?.tahunAjaranId,
            jumlah: 0
        };
        await sekolahService.createKategoriSekolah(newKategori);
        isDialogVisible.addKategoriSekolah = false;
        // await sekolahService.fetchKategoriSekolah()
        sekolahService.kategoriSekolahList.value.push(newKategori);
        kategoriSekolahTabel.value = sekolahService.kategoriSekolahTabel.value;
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
        await sekolahService.deleteKategoriSekolahKurikulum(item.kurikulum_id);
        sekolahService.kategoriSekolahList.value = sekolahService.kategoriSekolahList.value.filter((k) => k.kurikulum_id !== item.kurikulum_id);
        kategoriSekolahTabel.value = sekolahService.kategoriSekolahTabel.value;
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
        sekolahService.kategoriSekolahList.value.push(newKelas);
        kategoriSekolahTabel.value = sekolahService.kategoriSekolahTabel.value;
        await sekolahService.createKategoriSekolah(newKelas);

        isDialogVisible.addKategoriKelas = false;
        return;
    } catch (error) {
        console.error('Error adding kelas:', error);
        // Handle error
    }
};
const dialogEditKelas = (kategoriKelas) => {
    const cek = sekolahService.kategoriSekolahList.value.find((item) => kategoriKelas.kategori_sekolah_id == item.kategori_sekolah_id);
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
        const index = sekolahService.kategoriSekolahList.value.findIndex((item) => item.kategori_sekolah_id === updatedKelas.kategori_sekolah_id);
        if (index !== -1) {
            sekolahService.kategoriSekolahList.value[index] = updatedKelas;
        }
        // Update tabel
        kategoriSekolahTabel.value = [...sekolahService.kategoriSekolahTabel.value];

        // Kirim ke API
        await sekolahService.updateKategoriSekolah(updatedKelas);

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
        await sekolahService.deleteKategoriSekolah(item.kategori_sekolah_id);
        sekolahService.kategoriSekolahList.value = sekolahService.kategoriSekolahList.value.filter((k) => k.kategori_sekolah_id !== item.kategori_sekolah_id);
        kategoriSekolahTabel.value = sekolahService.kategoriSekolahTabel.value;
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
    sekolahService.createProsesKategoriSekolahDanKelas();
};
// Lifecycle
watch(sekolahService.initSelectedSemester, (newVal) => {
    initComponent();
});
onMounted(async () => {
    sekolah.value = await sekolahService.fetchSekolah();
    tingkat.value = await sekolahService.fetchTingkat();
    initComponent();
});
</script>
