<script setup>
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
// import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useKategoriSekolah } from '@/composables/sekolah_composable/useKategoriSekolah';
import { useKurikulum } from '@/composables/sekolah_composable/useKurikulum';
import { useToast } from 'primevue';
import { computed, onMounted, reactive, ref, watch } from 'vue';
import TingkatComponent from './TingkatComponent.vue';

const props = defineProps({
    schemaname: {
        type: String,
        required: true
    },
    tahunAjaranId: {
        type: String,
        required: true
    },
    jenjangPendidikanId: {
        type: Number,
        required: true,
        default: 6
    }
});
const toas = useToast();
const { createKategoriSekolah, createProsesKategoriSekolahDanKelas, createKategoriSekolahId, isKurikulumIdDuplicate, isTingkatIdDuplicate, fetchKategoriSekolah } = useKategoriSekolah(props.schemaname);

const { kurikulumList } = useKurikulum();
const kategoriSekolahTabel = ref([]);
const expandedRows = ref({});

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
const errorMessage = ref('');
const currentKategori = reactive({
    bidang_keahlian: null,
    program_keahlian: null,
    jurusan: null,
    jumlah: 0,
    tahun_ajaran_id: props.tahunAjaranId || '',
    tingkat_id: ''
});
// Watch for edit mode changes
watch(isEditKategoriSekolah, (newVal) => {
    isLoadingEditKategoriSekolah.value = true;

    if (newVal) {
        console.log(kategoriSekolahTabel.value);
        if (kategoriSekolahTabel.value.length > 0) {
            expandedRows.value = kategoriSekolahTabel.value.reduce((acc, p) => {
                acc[p.kurikulum_id] = true;
                return acc;
            }, {});
        }
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
            console.log(newVal);
            // formState.selectedKurikulum = kurikulumList.value.find((item) => newVal.jurusanId.includes(item.jurusanId));
            formState.selectedKurikulum = kurikulumList.value.find((item) => {
                if (item.jurusanId === newVal.jurusanId) return item;
            });
            formState.isSubmitDisabled = false;
        }
    }
);

// watch(props.tahunAjaranId, async (newVal) => {
//
// });

const internalValue = computed({
    get: () => props.tahunAjaranId
    // set: (value) => emit('update:modelValue', value)
});

watch(internalValue, async (newVal) => {
    if (newVal) {
        // kategoriSekolahTabel.value = await fetchKategoriSekolah(newVal);
        // console.log(tes);
    }
});
onMounted(async () => {
    kategoriSekolahTabel.value = await fetchKategoriSekolah(internalValue.value);
});
const onRowExpand = (event) => {
    // toas.add({ severity: 'info', summary: 'Product Expanded', detail: event.data.name, life: 3000 });
};
const onRowCollapse = (event) => {
    // toas.add({ severity: 'success', summary: 'Product Collapsed', detail: event.data.name, life: 3000 });
};

// Kategori Sekolah CRUD Operations
// ======================action
const openAddDialog = () => {
    formState.title = `Tambah Kompetensi Keahlian T.A. ${props.tahunAjaranId}`;
    isDialogVisible.addKategoriSekolah = true;
    // Reset currentKategori
    Object.assign(formState, {
        selectedKurikulum: ''
    });
    resetFormState();
};
const resetFormState = () => {
    Object.assign(currentKategori, {
        bidang_keahlian: null,
        program_keahlian: null,
        jurusan: null,
        jumlah: 0,
        tahun_ajaran_id: props.tahunAjaranId || '',
        tingkat_id: ''
    });
};

const addKategoriSekolah = async () => {
    if (!formState.selectedKurikulum) {
        alert('Kurikulum harus dipilih!');
        return;
    }

    try {
        console.log(kategoriSekolahTabel.value);

        const newKategori = {
            kategori_sekolah_id: createKategoriSekolahId(kategoriSekolahTabel.value),
            kurikulum_id: formState.selectedKurikulum.kurikulumId,
            nama_kurikulum: formState.selectedKurikulum.namaKurikulum,
            jurusan_id: currentKategori.jurusan?.jurusanId || '',
            nama_jurusan: currentKategori.jurusan?.namaJurusan || '',
            nama_bidang_keahlian: currentKategori.bidang_keahlian?.namaJurusan || '',
            nama_program_keahlian: currentKategori.program_keahlian?.namaJurusan || '',
            jenjang_pendidikan_id: formState.selectedKurikulum.jenjangPendidikanId,
            tahun_ajaran_id: props.tahunAjaranId,
            jumlah: 0,
            kategorikelas: [],
            is_added: false
        };
        if (isKurikulumIdDuplicate(kategoriSekolahTabel.value, newKategori.kurikulum_id)) {
            console.error(`❌ Gagal menambahkan data: kurikulum_id ${newKategori.kurikulum_id} sudah ada!`);
            errorDialog.value = true;
            errorMessage.value = `Kurikulum ${newKategori.nama_kurikulum} sudah ditambahkan, silahkan tambahkan kurikulum lainnya. `;
            return;
            // return {
            //     success: false,
            //     message: `kurikulum_id ${newKategori.kurikulum_id} sudah terdaftar`,
            //     data: kategoriSekolahTabel.value // Kembalikan array asli tanpa perubahan
            // };
        }
        kategoriSekolahTabel.value.push(newKategori);
        // localStorage.setItem('kategoriSekolahTemporary', JSON.stringify(newKategori));
        toas.add({ severity: 'success', summary: 'Sukses', detail: 'Kurikulum berhasil ditambahkan', life: 3000 });
        isDialogVisible.addKategoriSekolah = false;
        isComplete.value = false;
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
        // await deleteKategoriSekolahKurikulum(item.kurikulum_id);
        kategoriSekolahTabel.value = kategoriSekolahTabel.value.filter((k) => k.kurikulum_id !== item.kurikulum_id);
    } catch (error) {
        console.error('Error deleting kategori sekolah:', error);
        errorDialog.value = true;
        errorMessage.value = error;
        // Handle error
    } finally {
        isDialogVisible.deleteKategoriSekolah = false;
    }
};
const currentKategoriSekolah = ref();
const dialogAddKelas = (e) => {
    formState.title = `Tambah kelas | ${e.nama_jurusan}`;
    formState.selectedKategoriKelas = { ...e };
    currentKategoriSekolah.value = { ...e };
    // Reset form state
    formState.selectedTingkat = null;
    formState.isEditMode = false;
    currentKategori.jumlah = 0;
    currentKategori.tingkat_id = 0;
    isDialogVisible.addKategoriKelas = true;
};

const addKelas = async () => {
    if (!formState.selectedTingkat) {
        alert('Tingkat harus dipilih!');
        return;
    }

    try {
        const newKelas = {
            tingkat_id: formState.selectedTingkat,
            jumlah: currentKategori.jumlah,
            kategori_sekolah_id: currentKategoriSekolah.value.kategori_sekolah_id
        };
        if (isTingkatIdDuplicate(formState.selectedKategoriKelas.kategorikelas, newKelas.tingkat_id)) {
            console.error(`❌ Gagal menambahkan data: tingkat_id ${newKelas.tingkat_id} sudah ada!`);
            errorDialog.value = true;
            errorMessage.value = `Kelas  ${newKelas.tingkat_id} sudah ada pada kurkikulum ${currentKategoriSekolah.value.nama_kurikulum}! Silahkan tambahkan kelas yang lain `;
            return;
        }
        formState.selectedKategoriKelas.kategorikelas.push(newKelas);
        // localStorage.setItem('kategoriSekolahTemporary', JSON.stringify(kategoriSekolahTabel.value));
        toas.add({ severity: 'success', summary: 'Sukses', detail: 'Kelas berhasil ditambahkan', life: 3000 });
        isDialogVisible.addKategoriKelas = false;
        // return;
    } catch (error) {
        console.error('Error adding kelas:', error);
    }
};
const isDisableKelasOptions = ref(false);
const dialogEditKelas = (kategoriKelas) => {
    formState.title = 'Edit kelas';
    isDisableKelasOptions.value = true;
    formState.selectedKategoriKelas = kategoriKelas;
    formState.isEditMode = true;
    formState.selectedTingkat = kategoriKelas.tingkat_id;
    currentKategori.jumlah = kategoriKelas.jumlah;
    currentKategori.tingkat_id = kategoriKelas.tingkat_id;
    isDialogVisible.addKategoriKelas = true;
};

const updateKelas = async () => {
    if (!formState.selectedTingkat) {
        alert('Tingkat harus dipilih!');
        return;
    }
    try {
        const jumlahBaru = currentKategori.jumlah;

        const index = kategoriSekolahTabel.value.findIndex((item) => formState.selectedKategoriKelas.kategori_sekolah_id === item.kategori_sekolah_id);
        console.log(kategoriSekolahTabel.value[index].kategorikelas);
        const index2 = kategoriSekolahTabel.value[index].kategorikelas.findIndex((item) => formState.selectedKategoriKelas.tingkat_id === item.tingkat_id);
        if (index !== -1) {
            kategoriSekolahTabel.value[index].kategorikelas[index2].jumlah = jumlahBaru;
        }
        // Update tabel
        kategoriSekolahTabel.value = [...kategoriSekolahTabel.value];

        // Kirim ke API
        // await updateKategoriSekolah(newKelas);
        isDisableKelasOptions.value = false;

        toas.add({ severity: 'success', summary: 'Sukses', detail: 'Kelas berhasil diupdate', life: 3000 });

        // Tutup dialog
        isDialogVisible.addKategoriKelas = false;
    } catch (error) {
        console.error('Error updating kelas:', error);
    }
};

const handleDeleteKategoriKelas = async (e) => {
    console.log(e);
    isDialogVisible.deleteKategoriKelas = true;
    selectedItemToDelete.value = e;
    messageDelete.value = `Kelas <b>${e.tingkat_id}</b> akan dihapus?<br>`;
};
const deleteKelas = async (item) => {
    try {
        // await deleteKategoriSekolah(item.kategori_sekolah_id);
        // kategoriSekolahTabel.value.kategoriSekolah = kategoriSekolahTabel.value.kategoriSekolah.filter((k) => k.tingkat_id !== item.tingkat_id);
        const tes = hapusKelasByTingkatId(kategoriSekolahTabel, item.tingkat_id);
        console.log(tes);
    } catch (error) {
        console.log(error);
        // Handle error
    } finally {
        isDialogVisible.deleteKategoriSekolah = false;
    }
};

const hapusKelasByTingkatId = (dataArray, tingkatIdYangAkanDihapus) => {
    // Buat salinan baru untuk immutability
    console.log(dataArray);
    // const dataBaru = JSON.parse(JSON.stringify(dataArray));
    const dataBaru = [...dataArray];

    // Loop melalui setiap sekolah
    dataBaru.forEach((sekolah) => {
        // Filter kategorikelas: simpan yang tingkat_id-nya TIDAK sama dengan yang akan dihapus
        sekolah.kategorikelas = sekolah.kategorikelas.filter((kelas) => kelas.tingkat_id !== tingkatIdYangAkanDihapus);

        // Update total_kelas
        sekolah.total_kelas = sekolah.kategorikelas.length;
    });

    return dataBaru;
};

const newKategoriKelas = ref();
const dialogProsesKurikulum = (e) => {
    newKategoriKelas.value = { ...e };
    console.log(newKategoriKelas.value);
    messageDelete.value = `Anda akan memproses kurikulum. <br>Pastikan kurikulum dan jumlah kelas sudah sesuai dengan Satuan Pendidikan yang akan diproses. Lanjutkan!`;
    isDialogVisible.prosesKurikulum = true;
};
const addProsesKurikulum = async () => {
    // Cek apakah kelas sudah diisi?
    console.log(newKategoriKelas.value);
    const jmlKelas = newKategoriKelas.value?.kategorikelas.length;
    // console.log(jmlKelas);
    if (!jmlKelas) {
        alert(`Data tidak bisa disimpan karena kelas masih kosong`);
        return;
    }
    const res = await createKategoriSekolah(newKategoriKelas.value);
    // Setelah berhasil membuat kategori kelas, lanjut pembuatan kurikulum
    // createProsesKategoriSekolahDanKelas();
    if (res.status) {
        await createProsesKategoriSekolahDanKelas();
        toas.add({ severity: 'success', summary: 'Berhaisl', detail: 'Sukes Kompetensi ', life: 2000 });
    }
    return;
};

const isComplete = ref(true);
const resetKelas = () => {
    isDialogVisible.addKategoriKelas = false;
    isDisableKelasOptions.value = false;
};
</script>

<template>
    <div class="kategori-sekolah-container">
        <!-- Header Section -->
        <div>
            <div class="flex justify-between mb-4">
                <div>
                    <Button v-show="isEditKategoriSekolah" v-tooltip.bottom="`Tambah kompetensi keahlian`" label="Tambah" icon="pi pi-plus" size="small" class="rounded-full" @click="openAddDialog" />
                </div>
                <div>
                    <Button v-tooltip.bottom="'Edit kategori'" icon="pi pi-pencil" severity="secondary" :loading="isLoadingEditKategoriSekolah" @click="isEditKategoriSekolah = !isEditKategoriSekolah" />
                </div>
            </div>
            <!-- Main Table -->
            <div class="grid grid-cols-1 gap-4">
                <DataTable v-if="kategoriSekolahTabel.length > 0" v-model:expanded-rows="expandedRows" :value="kategoriSekolahTabel" data-key="kurikulum_id" striped-rows @row-expand="onRowExpand" @row-collapse="onRowCollapse">
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
                            <div class="flex justify-center space-x-2">
                                <Button v-tooltip.bottom="'Tambah kelas'" icon="pi pi-plus" class="!text-sm" severity="warn" size="small" rounded @click="dialogAddKelas(slotProps.data)" />
                                <Button v-tooltip.bottom="'Hapus kompetensi'" icon="pi pi-trash" class="!text-sm" severity="danger" size="small" rounded @click="handleDeleteKategoriSekolah(slotProps.data)" />
                                <Button v-show="!slotProps.data.is_added" v-tooltip.bottom="'Simpan'" icon="pi pi-save" class="!text-sm" severity="success" size="small" rounded @click="dialogProsesKurikulum(slotProps.data)" />
                            </div>
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
                                    <Button v-tooltip.bottom="'Edit kelas'" icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" size="small" rounded @click="dialogEditKelas(slotProps.data)" />
                                    <Button v-tooltip.bottom="'Hapus kelas'" icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" size="small" rounded @click="handleDeleteKategoriKelas(slotProps.data)" />
                                </template>
                            </Column>
                        </DataTable>
                    </template>
                </DataTable>
                <div v-else class="flex justify-center items-center">Tidak ada data! Silahkan tambah dengan memilih edit kategori</div>
            </div>
            <!-- <div class="flex justify-end mt-10">
                <Button v-show="isEditKategoriSekolah" icon="pi pi-process" label="Simpan" :disabled="kategoriSekolahTabel.length == 0 || kategoriSekolahTabel[0]?.kategorikelas.length == 0" @click="dialogProsesKurikulum" />
            </div> -->
        </div>
        <!-- Add Kategori Sekolah Dialog -->
        <Dialog v-model:visible="isDialogVisible.addKategoriSekolah" style="width: 550px" :header="formState.title" :modal="true">
            <div class="mb-4 md:grid space-y-2">
                <div v-show="jenjangPendidikanId == 6">
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
                <Button :label="formState.submitButtonText" severity="warn" :disabled="formState.isSubmitDisabled" @click="addKategoriSekolah" />
                <Button severity="help" label="Batal" class="w-20" @click="isDialogVisible.addKategoriSekolah = false" />
            </div>
        </Dialog>

        <!-- Edit Kategori Kelas Dialog -->
        <Dialog v-model:visible="isDialogVisible.addKategoriKelas" style="width: 550px" :header="formState.title" :modal="true">
            <div class="mb-4 md:grid space-y-2">
                <div>
                    <div class="font-semibold">Tingkat</div>
                    <TingkatComponent v-model="formState.selectedTingkat" :jenjang-pendidikan-id="props.jenjangPendidikanId" :model-value="currentKategori.tingkat_id" class="w-full" :is-disabled="isDisableKelasOptions" />
                </div>
                <div>
                    <div class="font-semibold">Jumlah Kelas</div>
                    <InputNumber v-model="currentKategori.jumlah" placeholder="Masukan jumlah kelas" class="block" fluid :min="1" />
                </div>
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <div>
                    <Button v-if="formState.isEditMode == true" label="Update" severity="warn" @click="updateKelas" />
                    <Button v-else label="Tambah" severity="warn" @click="addKelas" />
                </div>
                <Button severity="help" label="Batal" class="w-20" @click="resetKelas" />
            </div>
        </Dialog>

        <!-- Delete Confirmation Dialog -->
        <DialogConfirmDelete v-model:visible="isDialogVisible.deleteKategoriSekolah" :message="messageDelete" @confirm="deleteKategoriSekolah(selectedItemToDelete)" />
        <DialogConfirmDelete v-model:visible="isDialogVisible.deleteKategoriKelas" :message="messageDelete" @confirm="deleteKelas(selectedItemToDelete)" />
        <DialogConfirmDelete v-model:visible="isDialogVisible.prosesKurikulum" :message="messageDelete" judul="Proses" @confirm="addProsesKurikulum(selectedItemToDelete)" />
        <!-- <Dialog :visible="true">"hello"</Dialog> -->
        <Dialog v-model:visible="errorDialog" style="width: 550px" header="Error" :modal="true" position="top">
            <div>{{ errorMessage }}</div>
        </Dialog>
    </div>
</template>
