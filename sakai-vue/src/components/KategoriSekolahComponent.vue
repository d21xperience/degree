<template>
    <div class="">
        <div class="flex justify-between">
            <div>
                <Button v-show="isEditKategoriSekolah" icon="pi pi-plus" @click="openAddDialog" size="small" rounded v-tooltip.bottom="`Tambah kategori`" />
            </div>
            <div>
                <Button icon="pi pi-pencil" @click="isEditKategoriSekolah = !isEditKategoriSekolah" severity="secondary" v-tooltip.bottom="'Edit kategori'" />
            </div>
        </div>
        <div class="grid grid-cols-1 gap-4">
            <DataTable v-model:expandedRows="expandedRows" :value="kategoriSekolahList" @row-expand="onRowExpand" @row-collapse="onRowCollapse" dataKey="kategoriSekolahId" striped-rows>
                <!-- <template #header>
                    <div class="flex flex-wrap justify-end gap-2">
                        <Button text icon="pi pi-plus" label="Expand All" @click="expandAll" />
                        <Button text icon="pi pi-minus" label="Collapse All" @click="collapseAll" />
                    </div>
                </template> -->
                <Column expander style="width: 2rem" />
                <Column header="No" style="width: 2rem">
                    <template #body="slotProps">
                        {{ slotProps.index + 1 }}
                    </template>
                </Column>
                <Column header="Kurikulum" field="namaKurikulum" style="width: 20rem"></Column>
                <Column header="Jurusan" field="namaJurusan" />
                <!-- <Column header="Total Kelas"></Column> -->
                <Column header="Tahun Ajaran" field="tahunAjaranId" />
                <Column header="Aksi" :hidden="!isEditKategoriSekolah">
                    <template #body="slotProps">
                        <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="deleteKategoriSekolahDialog(slotProps.data)" size="small" rounded />
                        <Button icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" @click="openEditDialog(slotProps.data)" size="small" rounded />
                    </template>
                </Column>

                <template #expansion="slotProps">
                    <DataTable :value="slotProps.data.kategoriKelas">
                        <Column></Column>
                        <Column header="Kelas" field="tingkatId"> </Column>
                        <Column header="Jml.Kelas" field="jumlah"></Column>
                    </DataTable>
                </template>
            </DataTable>
        </div>
    </div>
    <div class="flex justify-end" v-show="isEditKategoriSekolah">
        <Button class="px-4 py-2 rounded flex items-center" @click="prosesKategoriSekolahDanKelas" label="Proses" severity="help" :loading="isLoadingProsesKategoriSekolahDanKelas" />
    </div>

    <!-- DIALOG -->
    <Dialog v-model:visible="isDialogAddKategoriSekolah" style="width: 550px" :header="title" :modal="true">
        <div class="mb-4 md:grid md:space-x-2" :class="{ 'md:grid-cols-2': sekolah.sekolah?.jenjangPendidikanId == 6 }">
            <div class="">
                Kurikulum
                <kurikulum-component v-model="kurikulumTerpilih" fluid id="kurikulum" />
            </div>
            <div v-show="sekolah.sekolah?.jenjangPendidikanId == 6">
                Jurusan
                <jurusan-component v-model="jurusanTerpilih" id="jurusan" />
            </div>
        </div>
        <h6>Jumlah Kelas yg menggunakan kurikulum tersebut</h6>
        <div class="my-2 grid md:grid-cols-4 md:space-x-1">
            <div v-for="(tkt, index) in tingkat" :key="tkt.kode">
                {{ tkt.nama }} :
                <InputNumber placeholder="e.g. 1 (max: 10)" fluid v-model="getKategoriKelas(tkt.kode).jumlah" />
            </div>
        </div>
        <div class="flex justify-end space-x-4 mt-10">
            <Button @click="editKategoriSekolah" label="Update" severity="warn" class="w-20" v-if="isEditDataKategoriSekolah" />
            <Button @click="addKategoriSekolah" label="Tambah" severity="warn" class="w-20" v-else />
            <Button @click="cancelKategoriSekolah" severity="help" label="Batal" class="w-20" />
        </div>
    </Dialog>

    <DialogConfirmDelete message="Apakah data ini akan dihapus?" v-model:visible="isDialogDeleteKategoriSekolah" @confirm="confirmDeleteKategoriSekolah"></DialogConfirmDelete>

    <!-- ----- -->
</template>

<script setup>
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import JurusanComponent from '@/components/JurusanComponent.vue';
import KurikulumComponent from '@/components/KurikulumComponent.vue';
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
import { onMounted, reactive, ref, watch } from 'vue';

const { fetchJurusan, fetchKurikulum } = useFormOptions();
const { fetchSekolah, fetchKategoriSekolah, createKategoriSekolah, deleteKategoriSekolah, updateKategoriSekolah, initSelectedSemester, fetchTingkat, createProsesKategoriSekolahDanKelas } = useSekolahService();

const kurikulumTerpilih = ref(null);
const jurusanTerpilih = ref(null);
const tahunAjaranId = ref();
const sekolah = ref({
    sekolah: {}
});
const tingkat = ref();
const kategoriSekolahList = ref([]);
const selectedKategoriSekolah = ref({
    kategoriKelas: []
});

const isLoadingProsesKategoriSekolahDanKelas = ref(false);
const isEditKategoriSekolah = ref(false);
const isDialogDeleteKategoriSekolah = ref(false);
const isDialogAddKategoriSekolah = ref(false);
const initFirst = async () => {
    sekolah.value = await fetchSekolah();
    // initKategoriSekolah();
    kategoriSekolahList.value = await fetchKategoriSekolah();
};
const initKategoriSekolah = async (valList = []) => {
    valList.forEach((el) => {
        if (el.kategoriKelas.length === 0) {
            el.kategoriKelas = tingkat.value.map((tingkatItem) => ({
                jumlah: 0,
                kategoriSekolahId: el.kategoriSekolahId,
                tahunAjaranId: el.tahunAjaranId,
                tingkatId: tingkatItem.kode
            }));
        }

        // Jika jumlah kategoriKelas kurang dari jumlah tingkat, tambahkan sisanya
        if (el.kategoriKelas.length < tingkat.value.length) {
            const existingTingkatIds = el.kategoriKelas.map((k) => k.tingkatId);
            const missingTingkat = tingkat.value.filter((t) => !existingTingkatIds.includes(t.kode));

            missingTingkat.forEach((t) => {
                el.kategoriKelas.push({
                    jumlah: 0,
                    kategoriSekolahId: el.kategoriSekolahId,
                    tahunAjaranId: el.tahunAjaranId,
                    tingkatId: t.kode
                });
            });
        }
    });
};

const prosesKategoriSekolahDanKelas = () => {
     
    isLoadingProsesKategoriSekolahDanKelas.value = true;
    createProsesKategoriSekolahDanKelas();
    isLoadingProsesKategoriSekolahDanKelas.value = false;
    isEditKategoriSekolah.value = false;
};

// ========================================
// Add
const title = ref('');
const addKategoriSekolah = () => {
    if (!kurikulumTerpilih.value || !jurusanTerpilih.value) {
        alert('Kurikulum dan jurusan harus dipilih!');
        return;
    }
    const newId = kategoriSekolahList.value.length ? Math.max(...kategoriSekolahList.value.map((item) => item.id)) + 1 : 1;
    const kategoriBaru = {
        // id: newId,
        // kategoriSekolahId: newId,
        namaKurikulum: kurikulumTerpilih.value.namaKurikulum || kurikulumTerpilih.value,
        namaJurusan: jurusanTerpilih.value.namaJurusan || jurusanTerpilih.value,
        jurusanId: jurusanTerpilih.value.jurusanId,
        kurikulumId: kurikulumTerpilih.value.kurikulumId,
        tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
        kategoriKelas: []
    };
    initKategoriSekolah([kategoriBaru]);
    kategoriBaru.kategoriKelas = kategoriBaru.kategoriKelas.map((item) => {
        const jumlah = kategoriKelas[item.tingkatId]?.jumlah || 0;
        return {
            ...item,
            jumlah
        };
    });
    kategoriSekolahList.value.push(kategoriBaru);

    // Kirim ke server
    createKategoriSekolah(kategoriBaru);
    // reset kategoriKelas
    Object.keys(kategoriKelas).forEach((key) => delete kategoriKelas[key]);
    // Reset input dan tutup dialog
    kurikulumTerpilih.value = null;
    jurusanTerpilih.value = null;
    isDialogAddKategoriSekolah.value = false;
};

const editKategoriSekolah = () => {
    if (!kurikulumTerpilih.value || !jurusanTerpilih.value) {
        alert('Kurikulum dan jurusan harus dipilih!');
        return;
    }

    const kategoriEdit = {
        kategoriSekolahId: selectedKategoriSekolah.value.kategoriSekolahId,
        namaKurikulum: kurikulumTerpilih.value.namaKurikulum || kurikulumTerpilih.value,
        namaJurusan: jurusanTerpilih.value.namaJurusan || jurusanTerpilih.value,
        jurusanId: jurusanTerpilih.value.jurusanId,
        kurikulumId: kurikulumTerpilih.value.kurikulumId,
        tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
        kategoriKelas: []
    };
    initKategoriSekolah([kategoriEdit]);
    kategoriEdit.kategoriKelas.id = 0
    kategoriEdit.kategoriKelas = kategoriEdit.kategoriKelas.map((item, idx) => {
        const jumlah = kategoriKelas[item.tingkatId]?.jumlah || 0;
        const id = selectedKategoriSekolah.value.kategoriKelas[idx].id;
        return {
            ...item,
            jumlah,
            id
        };
    });
    // kategoriSekolahList.value.push(kategoriEdit);
    console.log(kategoriEdit);
    // return;

    // Cari index dari kategori yang diedit
    const index = kategoriSekolahList.value.findIndex((k) => k.kategoriSekolahId === kategoriEdit.kategoriSekolahId);

    if (index !== -1) {
        // Ganti data lama dengan data baru
        kategoriSekolahList.value.splice(index, 1, kategoriEdit);
    } else {
        console.warn('Kategori sekolah tidak ditemukan di list');
    }

    // Kirim ke server
    updateKategoriSekolah(kategoriEdit);
    // reset kategoriKelas
    Object.keys(kategoriKelas).forEach((key) => delete kategoriKelas[key]);
    // Reset input dan tutup dialog
    kurikulumTerpilih.value = null;
    jurusanTerpilih.value = null;
    isEditDataKategoriSekolah.value = false;
    isDialogAddKategoriSekolah.value = false;
    // initKategoriSekolah();
};

// ========================================
// Delete

const deleteKategoriSekolahDialog = (e) => {
    isDialogDeleteKategoriSekolah.value = true;
    selectedKategoriSekolah.value = e;
};

const confirmDeleteKategoriSekolah = () => {
    kategoriSekolahList.value = kategoriSekolahList.value.filter((item) => item.kategoriSekolahId != selectedKategoriSekolah.value.kategoriSekolahId);
    isDialogDeleteKategoriSekolah.value = false;
    deleteKategoriSekolah(selectedKategoriSekolah.value.kategoriSekolahId);
};
// ========================================
const isEditDataKategoriSekolah = ref(false);
const openAddDialog = () => {
    title.value = 'Tambah Data';
    isDialogAddKategoriSekolah.value = true;
};
const openEditDialog = async (e) => {
    title.value = 'Edit Data';
    isEditDataKategoriSekolah.value = true;
    selectedKategoriSekolah.value = e;
    let tes = await fetchKurikulum();
    kurikulumTerpilih.value = tes.find((item) => item.namaKurikulum == selectedKategoriSekolah.value?.namaKurikulum);
    tes = await fetchJurusan();
    jurusanTerpilih.value = tes.find((item) => item.namaJurusan == selectedKategoriSekolah.value?.namaJurusan);
    isDialogAddKategoriSekolah.value = true;
    selectedKategoriSekolah.value.kategoriKelas.map((item) => {
        kategoriKelas[item.tingkatId] = { jumlah: item.jumlah };
    });
};
const cancelKategoriSekolah = () => {
    // selectedKategoriSekolah.value = null;
    isDialogAddKategoriSekolah.value = false;
    isEditDataKategoriSekolah.value = false;
};

const expandedRows = ref();
const onRowExpand = (event) => {
    // Ambil data mapel untuk kelas tertentu
    //   console.log(event);
};
const onRowCollapse = (event) => {
    //   console.log(event);
};
const expandAll = () => {
    expandedRows.value = kategoriSekolahList.value.reduce((acc, p) => (acc[p.kategoriSekolahId] = true) && acc, {});
};
const collapseAll = () => {
    expandedRows.value = null;
};

const kategoriKelas = reactive({});
const getKategoriKelas = (kode) => {
    if (!kategoriKelas[kode]) {
        kategoriKelas[kode] = { jumlah: 0 };
    }
    return kategoriKelas[kode];
};

onMounted(async () => {
    initFirst();
    tingkat.value = await fetchTingkat();
});

watch(isDialogAddKategoriSekolah, (newVal) => {
    // console.log(newVal);
    if (!newVal) {
        jurusanTerpilih.value = null;
        kurikulumTerpilih.value = null;
    }
});
watch(initSelectedSemester, (newVal) => {
    initKategoriSekolah();
});
</script>
