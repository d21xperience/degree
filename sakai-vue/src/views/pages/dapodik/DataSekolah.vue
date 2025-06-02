<template>
    <div>
        <div class="flex justify-between items-center mb-2">
            <h1 class="text-2xl font-bold">Data Sekolah</h1>
            <div>
                <Button icon="pi pi-pencil" @click="editSekolah" :style="isEdit ? 'background-color:blue;border:none' : 'background-color:gray;border:none'" v-tooltip.bottom="'Edit data'" />
            </div>
        </div>
        <div>
            <div class="p-2">
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Identitas</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama</div>
                        <div>{{ sekolah?.sekolah.nama }}</div>
                        <div>Jenjang</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.jenjangPendidikanStr }}</div>
                            <div v-else>
                                <Select v-model="selectedJenjangPendidikan" filter :options="jenjangPendidikanFiltered" optionLabel="nama" placeholder="Pilih jenjang pendidikan" fluid showClear class="w-full md:w-56" />
                            </div>
                        </div>
                        <div>Bentuk Pendidikan</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.bentukPendidikanStr }}</div>
                            <div v-else>
                                <Select v-model="selectedBentukPendidikan" :options="bentukPendidikan" filter optionLabel="nama" placeholder="Pilih bentuk pendidikan" fluid showClear class="w-full md:w-56" />
                            </div>
                        </div>
                        <div>NSS</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.sekolah.nss }}</div>
                            <div v-else><InputText type="text" placeholder="Masukan NSS" v-model="sekolah.sekolah.nss" fluid /></div>
                        </div>
                        <div>NPSN</div>
                        <div>{{ sekolah?.sekolah.npsn }}</div>
                    </div>
                </div>
                <div class="my-10">
                    <h3 class="text-lg font-semibold mb-2">Alamat</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Jalan</div>
                        <div>{{ sekolah?.sekolah.alamat }}</div>
                        <div>Desa/Kelurahan</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.sekolah.kelurahan }}</div>
                            <InputText v-model="sekolah.sekolah.kelurahan" placeholder="Masukan nama Desa/Kelurahan" fluid v-else />
                        </div>
                        <div>Kecamatan</div>
                        <div>{{ sekolah?.sekolah.kecamatan }}</div>
                        <div>Provinsi</div>
                        <div>{{ sekolah?.sekolah.propinsi }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kontak</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Telp./Fax.</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.sekolah.telepon }}/{{ sekolah?.sekolah.telepon }}</div>
                            <div v-else class="space-y-2">
                                Telp.
                                <InputText type="text" placeholder="Masukan no.tlp" v-model="sekolah.sekolah.telepon" fluid />
                                <div>
                                    Fax.
                                    <InputText type="text" placeholder="Masukan no.fax" v-model="sekolah.sekolah.fax" fluid />
                                </div>
                            </div>
                        </div>
                        <div>email</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah?.sekolah.email }}</div>
                            <div v-else><InputText type="text" placeholder="Masukan alamat email" v-model="sekolah.sekolah.email" fluid /></div>
                        </div>
                        <div>website</div>
                        <div>
                            <div v-if="!isEdit">
                                <a v-if="sekolah?.sekolah.website" :href="getWebsiteUrl(sekolah?.sekolah.website)" target="_blank" rel="noopener noreferrer" class="text-blue-700 underline">
                                    {{ sekolah?.sekolah.website }}
                                </a>
                            </div>
                            <div v-else><InputText type="text" placeholder="Masukan alamat website" v-model="sekolah.sekolah.website" fluid /></div>
                        </div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama Kepsek</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.sekolah?.nmKepsek }}</div>
                            <div v-else><InputText v-model="sekolah.sekolah.nmKepsek" fluid name="nmKepsek" /></div>
                        </div>
                        <div>NIP Kepsek</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.sekolah?.nipKepsek }}</div>
                            <div v-else><InputText v-model="sekolah.sekolah.nipKepsek" fluid name="nipKepsek" /></div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex justify-end" v-show="isEdit">
                <Button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center" @click="updateData" label="Update Data" icon="pi pi-save" :loading="loadingUpdate" />
            </div>

            <div class="my-10">
                <div class="flex justify-between">
                    <div class="flex">
                        <h3 class="text-lg font-semibold mb-2 mr-2">Kategori</h3>
                        <Button v-show="isEditKategoriSekolah" icon="pi pi-plus" @click="openAddDialog" size="small" rounded v-tooltip.bottom="`Tambah kategori`" />
                    </div>
                    <div>
                        <Button icon="pi pi-pencil" @click="isEditKategoriSekolah = !isEditKategoriSekolah" severity="secondary" v-tooltip.bottom="'Edit kategori'" />
                    </div>
                </div>
                <div class="grid grid-cols-1 gap-4">
                    <DataTable v-model:expandedRows="expandedRows" :value="kategoriSekolahList" @row-expand="onRowExpand" @row-collapse="onRowCollapse" dataKey="kategoriSekolahId" striped-rows>
                        <template #header>
                            <div class="flex flex-wrap justify-end gap-2">
                                <Button text icon="pi pi-plus" label="Expand All" @click="expandAll" />
                                <Button text icon="pi pi-minus" label="Collapse All" @click="collapseAll" />
                            </div>
                        </template>
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
                <Button class="px-4 py-2 rounded flex items-center" @click="prosesKategoriSekolahDanKelas = false" label="Proses" severity="help" :loading="isLoadingProsesKategoriSekolahDanKelas" />
            </div>
        </div>

        <Toast />
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
    </div>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// ambil data dari backend
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import JurusanComponent from '@/components/JurusanComponent.vue';
import KurikulumComponent from '@/components/KurikulumComponent.vue';
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
const { fetchSekolah, fetchKategoriSekolah, createKategoriSekolah, deleteKategoriSekolah, updateKategoriSekolah, initSelectedSemester, updateSekolah, fetchTingkat, createProsesKategoriSekolahDanKelas } = useSekolahService();

const { fetchJurusan, fetchKurikulum } = useFormOptions();

const fetchRefTable = async () => {
    bentukPendidikan.value = await store.dispatch('sekolahService/fetchBentukPendidikan');
    jenjangPendidikan.value = await store.dispatch('sekolahService/fetchJenjangPendidikan');
};

// Evaluasi variabel di bawah ini:
const jenjangPendidikanFiltered = computed(() => {
    return jenjangPendidikan.value.filter((item) => item.jenjangLembaga === 1);
});

const sekolah = ref({
    sekolah: {}
});
const initKategoriSekolah = async () => {
    kategoriSekolahList.value = await fetchKategoriSekolah();

    kategoriSekolahList.value.forEach((el) => {
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
const initFirst = async () => {
    // fetchSemester;
    sekolah.value = await fetchSekolah();
    initKategoriSekolah();
    // console.log(kategoriSekolahList.value);
};
const tingkat = ref();
onMounted(async () => {
    initFirst();
    tingkat.value = await fetchTingkat();

    // console.log(tingkat.value);
    // const data = await fetchSekolah();
    // Object.assign(sekolah, data);
});
watch(initSelectedSemester, () => {
    // console.log('update semester');
    initFirst();
});
const selectedBentukPendidikan = ref();
const selectedJenjangPendidikan = ref();
watch(selectedBentukPendidikan, (newVal) => {
    if (newVal) {
        sekolah.value.sekolah.bentukPendidikanId = newVal.bentukPendidikanId;
        sekolah.value.bentukPendidikanStr = newVal.nama;
    }
});
watch(selectedJenjangPendidikan, (newVal) => {
    if (newVal) {
        sekolah.value.sekolah.jenjangPendidikanId = newVal.jenjangPendidikanId;
        sekolah.value.jenjangPendidikanStr = newVal.nama;
    }
});
// Edit
const isEdit = ref(false);
const isEditKategoriSekolah = ref(false);
const editSekolah = () => {
    isEdit.value = !isEdit.value;
};
watch(isEdit, async (newVal) => {
    if (newVal) {
        await fetchRefTable();
        selectedBentukPendidikan.value = bentukPendidikan.value.find((item) => item.nama === sekolah.value.bentukPendidikanStr);
        selectedJenjangPendidikan.value = jenjangPendidikan.value.find((item) => item.nama === sekolah.value.jenjangPendidikanStr);
    }
});

const bentukPendidikan = ref([]);
const jenjangPendidikan = ref([]);
const tahunAjaranId = initSelectedSemester.value?.tahunAjaranId;
const getWebsiteUrl = (url) => {
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
        return `https://${url}`; // Tambahkan https jika belum ada
    }
    return url;
};
const loadingUpdate = ref(false);
const updateData = async () => {
    loadingUpdate.value = true;
    // console.log(sekolah.value);
    const payload = {
        schemaname: store.getters['sekolahService/getTabeltenant']?.schemaname,
        sekolah: sekolah.value.sekolah,
        bentukPendidikanStr: selectedBentukPendidikan.value.nama,
        jenjangPendidikanStr: selectedJenjangPendidikan.value.nama
    };
    await updateSekolah(payload);
    selectedBentukPendidikan.value = null;
    loadingUpdate.value = false;
    isEdit.value = false;
};
const kurikulumTerpilih = ref(null);
const jurusanTerpilih = ref(null);
const kategoriSekolahList = ref([]);
const selectedKategoriSekolah = ref({
    kategoriKelas: []
});

const isDialogAddKategoriSekolah = ref(false);

watch(isDialogAddKategoriSekolah, (newVal) => {
    // console.log(newVal);
    if (!newVal) {
        jurusanTerpilih.value = null;
        kurikulumTerpilih.value = null;
    }
});

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
        id: newId,
        namaKurikulum: kurikulumTerpilih.value.namaKurikulum || kurikulumTerpilih.value,
        namaJurusan: jurusanTerpilih.value.namaJurusan || jurusanTerpilih.value,
        jurusanId: jurusanTerpilih.value.jurusanId,
        kurikulumId: kurikulumTerpilih.value.kurikulumId,
        tahunAjaranId: tahunAjaranId,
        kategoriKelas: {}
    };

    kategoriSekolahList.value.push(kategoriBaru);
    // Kirim ke server
    createKategoriSekolah(kategoriBaru);
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

    const kategoriBaru = {
        kategoriSekolahId: selectedKategoriSekolah.value.kategoriSekolahId,
        namaKurikulum: kurikulumTerpilih.value.namaKurikulum || kurikulumTerpilih.value,
        namaJurusan: jurusanTerpilih.value.namaJurusan || jurusanTerpilih.value,
        jurusanId: jurusanTerpilih.value.jurusanId,
        kurikulumId: kurikulumTerpilih.value.kurikulumId,
        tahunAjaranId: tahunAjaranId,
        kategoriKelas: selectedKategoriSekolah.value.kategoriKelas.map((k) => {
            return {
                kategori_sekolah_id: k.kategoriSekolahId,
                tingkat_id: k.tingkatId,
                jumlah: k.jumlah,
                tahun_ajaran_id: k.tahunAjaranId
            };
        })
    };
    console.log(kategoriBaru);
    // return;

    // Cari index dari kategori yang diedit
    const index = kategoriSekolahList.value.findIndex((k) => k.kategoriSekolahId === kategoriBaru.kategoriSekolahId);

    if (index !== -1) {
        // Ganti data lama dengan data baru
        kategoriSekolahList.value.splice(index, 1, kategoriBaru);
    } else {
        console.warn('Kategori sekolah tidak ditemukan di list');
    }

    // Kirim ke server
    updateKategoriSekolah(kategoriBaru);

    // Reset input dan tutup dialog
    kurikulumTerpilih.value = null;
    jurusanTerpilih.value = null;
    isEditDataKategoriSekolah.value = false;
    isDialogAddKategoriSekolah.value = false;
    initKategoriSekolah();
};

// ========================================
// Delete
const isDialogDeleteKategoriSekolah = ref(false);
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

const getKategoriKelas = computed(() => {
    return (kode) => {
        const kelas = selectedKategoriSekolah.value.kategoriKelas.find((k) => k.tingkatId == kode);
        return kelas || { jumlah: 0 };
    };
});

const isLoadingProsesKategoriSekolahDanKelas = ref(false);
const prosesKategoriSekolahDanKelas = () => {
    isLoadingProsesKategoriSekolahDanKelas.value = true;
    createProsesKategoriSekolahDanKelas();
    isLoadingProsesKategoriSekolahDanKelas.value = false;
    isEditKategoriSekolah.value = false;
};
</script>

<style scoped>
edit-class {
    background-color: red;
}
</style>
