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
                        <div>{{ sekolah.value?.nama }}</div>
                        <div>Jenjang</div>
                        <div>
                            <div v-if="!isEdit">{{ jenjangPendidikanSelected?.nama }}</div>
                            <div v-else><Select v-model="jenjangPendidikanSelected" filter :options="jenjangPendidikanFiltered" optionLabel="nama" placeholder="Pilih jenjang pendidikan" fluid showClear class="w-full md:w-56" /></div>
                        </div>
                        <div>Bentuk Pendidikan</div>
                        <div>
                            <div v-if="!isEdit">{{ bentukPendidikanSelected?.nama }}</div>
                            <div v-else><Select v-model="bentukPendidikanSelected" :options="bentukPendidikan" filter optionLabel="nama" placeholder="Pilih bentuk pendidikan" fluid showClear class="w-full md:w-56" /></div>
                        </div>
                        <div>NSS</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.value?.nss }}</div>
                            <div v-else><InputText type="text" placeholder="Masukan NSS" v-model="sekolah.value.nss" fluid /></div>
                        </div>
                        <div>NPSN</div>
                        <div>{{ sekolah.value?.npsn }}</div>
                    </div>
                </div>

                <div class="my-10">
                    <h3 class="text-lg font-semibold mb-2">Alamat</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Jalan</div>
                        <div>{{ sekolah.value?.alamat }}</div>
                        <div>Desa/Kelurahan</div>
                        <div>{{ sekolah.value?.kelurahan }}</div>
                        <div>Kecamatan</div>
                        <div>{{ sekolah.value?.kecamatan }}</div>
                        <div>Provinsi</div>
                        <div>{{ sekolah.value?.propinsi }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kontak</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Telp./Fax.</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.value?.telepon }}/{{ sekolah.value?.telepon }}</div>
                            <div v-else class="space-y-2">
                                Telp.
                                <InputText type="text" placeholder="Masukan no.tlp" v-model="sekolah.value.telepon" fluid />
                                <div>
                                    Fax.
                                    <InputText type="text" placeholder="Masukan no.fax" v-model="sekolah.value.fax" fluid />
                                </div>
                            </div>
                        </div>
                        <div>email</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.value?.email }}</div>
                            <div v-else><InputText type="text" placeholder="Masukan alamat email" v-model="sekolah.value.email" fluid /></div>
                        </div>
                        <div>website</div>
                        <div>
                            <div v-if="!isEdit">
                                <a v-if="sekolah.value?.website" :href="getWebsiteUrl(sekolah.value?.website)" target="_blank" rel="noopener noreferrer" class="text-blue-700 underline">
                                    {{ sekolah.value?.website }}
                                </a>
                            </div>
                            <div v-else><InputText type="text" placeholder="Masukan alamat website" v-model="sekolah.value.website" fluid /></div>
                        </div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama Kepsek</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.value?.nmKepsek }}</div>
                            <div v-else><InputText v-model="sekolah.value.nmKepsek" fluid name="nmKepsek" /></div>
                        </div>
                        <div>NIP Kepsek</div>
                        <div>
                            <div v-if="!isEdit">{{ sekolah.value?.nipKepsek }}</div>
                            <div v-else><InputText v-model="sekolah.value.nipKepsek" fluid name="nipKepsek" /></div>
                        </div>
                    </div>
                </div>
                <div class="my-10">
                    <div class="flex justify-between">
                        <h3 class="text-lg font-semibold mb-2">Kategori</h3>
                        <div v-show="isEdit">
                            <Button icon="pi pi-plus" @click="isDialogAddKategoriSekolah = true" />
                        </div>
                    </div>
                    <div class="grid grid-cols-1 gap-4">
                        <DataTable :value="kategoriSekolahList" dataKey="id">
                            <Column header="No">
                                <template #body="slotProps">
                                    {{ slotProps.index + 1 }}
                                </template>
                            </Column>
                            <Column header="Kurikulum" field="kurikulum"></Column>
                            <Column header="Jurusan" field="jurusan" />
                            <Column header="Tahun Ajaran" field="tahun_ajaran_id" />
                            <Column header="Aksi" v-if="isEdit">
                                <template #body="slotProps">
                                    <Button icon="pi pi-trash" class="mr-2 !text-sm" severity="danger" @click="deleteKategoriSekolahDialog(slotProps.data)" />
                                    <!-- <Button icon="pi pi-pencil" @click="editKategoriSekolah(slotProps.data)" /> -->
                                </template>
                            </Column>
                        </DataTable>
                    </div>
                </div>
            </div>
            <div class="flex justify-end" v-show="isEdit">
                <button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center" @click="updateData"><i class="fas fa-edit mr-2"></i> Update Data</button>
            </div>
        </div>
        <Toast />

        <!-- <Dialog v-model:visible="isDialogEditKategoriSekolah" header="Edit" :modal="true">
            <div>
                Kuriklum:
                {{ selectedKategoriSekolah?.kurikulum }}
                Jurusan:
                {{ selectedKategoriSekolah?.jurusan }}
            </div>

            <div class="flex justify-end space-x-4 mt-10">
                <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600" @click="confirmEditKategoriSekolah">Simpan</button>
                <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400" @click="isDialogEditKategoriSekolah = false">Batal</button>
            </div>
        </Dialog> -->
        <Dialog v-model:visible="isDialogAddKategoriSekolah" style="width: 450px" header="Tambah" :modal="true">
            <div class="mb-4">
                <label for="kurikulum">Kurikulum</label>
                <kurikulum-component v-model="kurikulumTerpilih" />
            </div>
            <div>
                <label for="">Jurusan</label>
                <jurusan-component v-model="jurusanTerpilih" />
            </div>
            <div class="flex justify-end space-x-4 mt-10">
                <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600" @click="addKategoriSekolah">Tambah</button>
                <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400" @click="isDialogAddKategoriSekolah = false">Batal</button>
            </div>
        </Dialog>

        <DialogConfirmDelete message="Apakah data ini akan dihapus?" v-model:visible="isDialogDeleteKategoriSekolah" @confirm="confirmDeleteKategoriSekolah"></DialogConfirmDelete>
    </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// ambil data dari backend
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import JurusanComponent from '@/components/JurusanComponent.vue';
import KurikulumComponent from '@/components/KurikulumComponent.vue';
import Button from 'primevue/button';
import Select from 'primevue/select';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();
const sekolah = reactive({});
const dataSekolah = ref();

const fetchSekolah = async () => {
    // get tabel tenant
    // ambil sekolah id dari user yang sedang login
    const sekolahID = await store.state.authService.user?.sekolahTenantId;
    const tTenant = await store.dispatch('sekolahService/fetchTabeltenant', sekolahID);
    // console.log(tTenant)
    dataSekolah.value = await store.dispatch('sekolahService/fetchSekolah', { schemaname: tTenant.schemaname, namaSekolah: tTenant.namaSekolah });
    // console.log(dataSekolah)
    sekolah.value = dataSekolah;
};

const fetchRefTable = async () => {
    bentukPendidikan.value = await store.dispatch('sekolahService/fetchBentukPendidikan');
    jenjangPendidikan.value = await store.dispatch('sekolahService/fetchJenjangPendidikan');
};

const jenjangPendidikanFiltered = computed(() => {
    return jenjangPendidikan.value.filter((item) => item.jenjangLembaga === 1);
});

onMounted(async () => {
    await fetchSekolah();
    await fetchRefTable();
    jenjangPendidikanSelected.value = jenjangPendidikan.value.find((item) => item.jenjangPendidikanId == sekolah.value.jenjangPendidikanId);
    bentukPendidikanSelected.value = bentukPendidikan.value.find((item) => item.bentukPendidikanId == sekolah.value.bentukPendidikanId);
});

// Edit
const isEdit = ref(false);
const editSekolah = () => {
    isEdit.value = !isEdit.value;
};
// watch(isEdit, async (newVal) => {
//     if (newVal) {
//         await fetchKurikulum();
//     }
// });
const bentukPendidikan = ref([]);
const jenjangPendidikan = ref([]);
const tahunAjaranId = ref('20231'); // Misal nilai default, bisa juga pakai dropdown tahun ajaran

const bentukPendidikanSelected = ref();
const jenjangPendidikanSelected = ref();

const changedData = ref({});
watch(
    sekolah,
    (newVal) => {
        changedData.value = {}; // Reset perubahan
        for (const key in newVal) {
            if (newVal[key] !== dataSekolah.value[key]) {
                changedData.value[key] = newVal[key]; // Simpan hanya data yang berubah
            }
        }
        // console.log("Perubahan terdeteksi:", changedData.value);
    },
    { deep: true } // Perlu `deep: true` untuk melacak perubahan dalam objek
);

watch(jenjangPendidikanSelected, (newVal) => (sekolah.value.jenjangPendidikanId = newVal.jenjangPendidikanId));
watch(bentukPendidikanSelected, (newVal) => (sekolah.value.bentukPendidikanId = newVal.bentukPendidikanId));
const getWebsiteUrl = (url) => {
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
        return `https://${url}`; // Tambahkan https jika belum ada
    }
    return url;
};

const updateData = async () => {
    if (Object.keys(changedData.value).length > 0) {
        // Kirim `changedData.value` ke backend via API
        const payload = {
            schemaname: store.getters['sekolahService/getTabeltenant']?.schemaname,
            sekolah: sekolah.value
        };
        const resp = await store.dispatch('sekolahService/updateSekolah', payload);
        toast.add({ severity: 'info', summary: 'Info', detail: resp?.message, life: 3000 });
        // console.log(resp)
    } else {
        toast.add({ severity: 'warn', summary: 'Info', detail: resp?.message, life: 3000 });

        console.log('Tidak ada perubahan, tidak perlu update.');
    }
    isEdit.value = false;
};
const kurikulumTerpilih = ref(null);
const jurusanTerpilih = ref(null);
const kategoriSekolahList = ref([
    {
        id: 1,
        kurikulum: 'Kurikulum TKJ',
        jurusan: 'Teknik Komputer dan Jaringan',
        tahun_ajaran_id: '20231'
    },
    {
        id: 2,
        kurikulum: 'Kurikulum Akuntansi',
        jurusan: 'Akuntansi Keuangan Lembaga',
        tahun_ajaran_id: '20231'
    },
    {
        id: 3,
        kurikulum: 'Kurikulum Bisnis',
        jurusan: 'Manajemen Perkantoran',
        tahun_ajaran_id: '20231'
    }
]);
const selectedKategoriSekolah = ref();

const isDialogAddKategoriSekolah = ref(false);

watch(isDialogAddKategoriSekolah, (newVal) => {
    console.log(newVal);
    if (!newVal) {
        jurusanTerpilih.value = null;
        kurikulumTerpilih.value = null;
    }
});

// ========================================
// Add
const addKategoriSekolah = () => {
    if (!kurikulumTerpilih.value || !jurusanTerpilih.value) {
        alert('Kurikulum dan jurusan harus dipilih!');
        return;
    }

    const newId = kategoriSekolahList.value.length ? Math.max(...kategoriSekolahList.value.map((item) => item.id)) + 1 : 1;

    const kategoriBaru = {
        id: newId,
        kurikulum: kurikulumTerpilih.value.namaKurikulum || kurikulumTerpilih.value,
        jurusan: jurusanTerpilih.value.namaJurusan || jurusanTerpilih.value,
        tahun_ajaran_id: tahunAjaranId.value
    };

    kategoriSekolahList.value.push(kategoriBaru);

    // Reset input dan tutup dialog
    kurikulumTerpilih.value = null;
    jurusanTerpilih.value = null;
    isDialogAddKategoriSekolah.value = false;
};

// ========================================

// const isDeleteKategoriSekolah = ref(false);
// ========================================
// Edit
const isDialogEditKategoriSekolah = ref(false);
const editKategoriSekolah = (e) => {
    isDialogEditKategoriSekolah.value = true;
    selectedKategoriSekolah.value = e;
};
const confirmEditKategoriSekolah = () => {
    addKategoriSekolah.value = true;
};

// ========================================
// Delete
const isDialogDeleteKategoriSekolah = ref(false);
const deleteKategoriSekolahDialog = (e) => {
    isDialogDeleteKategoriSekolah.value = true;
    selectedKategoriSekolah.value = e;
};

const confirmDeleteKategoriSekolah = () => {
    kategoriSekolahList.value = kategoriSekolahList.value.filter((item) => item.id != selectedKategoriSekolah.value.id);
    isDialogDeleteKategoriSekolah.value = false;
};
// ========================================
</script>

<style scoped>
edit-class {
    background-color: red;
}
</style>
