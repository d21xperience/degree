<script setup>
import KategoriSekolahComponent from '@/components/sekolah_components/KategoriSekolahComponent.vue';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';

import { computed, onMounted, ref, watch } from 'vue';
const { initSelectedSemester } = useSemester();
const { fetchBentukPendidikan, fetchJenjangPendidikan, fetchSekolah, fetchTingkat, updateSekolah } = useSekolah();
const fetchRefTable = async () => {
    bentukPendidikan.value = await fetchBentukPendidikan(); //getBentukPendidikan();
    jenjangPendidikan.value = await fetchJenjangPendidikan(); //getJenjangPendidikan();
    console.log(jenjangPendidikan.value);
};
const { schemaname } = useTableTenant();
// Evaluasi variabel di bawah ini:
const jenjangPendidikanFiltered = computed(() => {
    return jenjangPendidikan.value.filter((item) => item.jenjangLembaga === 1);
});

const sekolah = ref({
    sekolah: {}
});

const initFirst = async () => {
    sekolah.value = await fetchSekolah();
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

const editSekolah = () => {
    isEdit.value = !isEdit.value;
};

const bentukPendidikan = ref([]);
const jenjangPendidikan = ref([]);
watch(isEdit, async (newVal) => {
    if (newVal) {
        await fetchRefTable();
        selectedBentukPendidikan.value = bentukPendidikan.value.find((item) => item.nama === sekolah.value.bentukPendidikanStr);
        selectedJenjangPendidikan.value = jenjangPendidikan.value.find((item) => item.nama === sekolah.value.jenjangPendidikanStr);
    }
});
const getWebsiteUrl = (url) => {
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
        return `https://${url}`; // Tambahkan https jika belum ada
    }
    return url;
};
const loadingUpdate = ref(false);
const updateData = async () => {
    loadingUpdate.value = true;
    try {
        const payload = {
            schemaname: schemaname.value,
            sekolah: sekolah.value.sekolah,
            bentukPendidikanStr: selectedBentukPendidikan.value.nama,
            jenjangPendidikanStr: selectedJenjangPendidikan.value.nama
        };
        await updateSekolah(payload);
        selectedBentukPendidikan.value = null;
    } catch (error) {
        console.log(error);
    } finally {
        loadingUpdate.value = false;
        isEdit.value = false;
    }
};
</script>

<template>
    <div>
        <div>
            <div class="flex justify-between items-center mb-2">
                <h5>Data Sekolah</h5>
                <div>
                    <Button v-tooltip.bottom="'Edit data sekolah'" icon="pi pi-pencil" :style="isEdit ? 'background-color:blue;border:none' : 'background-color:gray;border:none'" size="small" rounded="" @click="editSekolah" />
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
                                    <Select v-model="selectedJenjangPendidikan" filter :options="jenjangPendidikanFiltered" option-label="nama" placeholder="Pilih jenjang pendidikan" fluid show-clear class="w-full md:w-56" />
                                </div>
                            </div>
                            <div>Bentuk Pendidikan</div>
                            <div>
                                <div v-if="!isEdit">{{ sekolah?.bentukPendidikanStr }}</div>
                                <div v-else>
                                    <Select v-model="selectedBentukPendidikan" :options="bentukPendidikan" filter option-label="nama" placeholder="Pilih bentuk pendidikan" fluid show-clear class="w-full md:w-56" />
                                </div>
                            </div>
                            <div>NSS</div>
                            <div>
                                <div v-if="!isEdit">{{ sekolah?.sekolah.nss }}</div>
                                <div v-else><InputText v-model="sekolah.sekolah.nss" type="text" placeholder="Masukan NSS" fluid /></div>
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
                                <InputText v-else v-model="sekolah.sekolah.kelurahan" placeholder="Masukan nama Desa/Kelurahan" fluid />
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
                                    <InputText v-model="sekolah.sekolah.telepon" type="text" placeholder="Masukan no.tlp" fluid />
                                    <div>
                                        Fax.
                                        <InputText v-model="sekolah.sekolah.fax" type="text" placeholder="Masukan no.fax" fluid />
                                    </div>
                                </div>
                            </div>
                            <div>email</div>
                            <div>
                                <div v-if="!isEdit">{{ sekolah?.sekolah.email }}</div>
                                <div v-else><InputText v-model="sekolah.sekolah.email" type="text" placeholder="Masukan alamat email" fluid /></div>
                            </div>
                            <div>website</div>
                            <div>
                                <div v-if="!isEdit">
                                    <a v-if="sekolah?.sekolah.website" :href="getWebsiteUrl(sekolah?.sekolah.website)" target="_blank" rel="noopener noreferrer" class="text-blue-700 underline">
                                        {{ sekolah?.sekolah.website }}
                                    </a>
                                </div>
                                <div v-else><InputText v-model="sekolah.sekolah.website" type="text" placeholder="Masukan alamat website" fluid /></div>
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
                <div v-show="isEdit" class="flex justify-end">
                    <Button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center" label="Update Data" icon="pi pi-save" :loading="loadingUpdate" @click="updateData" />
                </div>
            </div>
        </div>
        <!-- Informasi sekolah -->
        <div>
            <div>
                <h5>Kompetensi keahlian Dilayani</h5>
                <KategoriSekolahComponent />
            </div>
        </div>
    </div>
</template>

<style scoped>
edit-class {
    background-color: red;
}
</style>
