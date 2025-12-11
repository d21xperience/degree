<script setup>
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { onMounted, ref, watch } from 'vue';
import BentukPendidikanComponent from './BentukPendidikanComponent.vue';
import JenjangPendidikanComponent from './JenjangPendidikanComponent.vue';

const sekolah = ref({
    sekolah: {}
});

const props = defineProps({
    schemaname: {
        type: String,
        required: true
    },
    storeSekolah: {
        type: Object,
        required: true
    }
});

watch(
    () => props.storeSekolah,
    (baru) => {
        sekolah.value = baru;
    }
);

const { isFetching, updateSekolah, isError } = useSekolah({ schemaname: props.schemaname, autoload: false });

onMounted(async () => {
    if (!props.storeSekolah) {
        // await initialize();
        sekolah.value = props.storeSekolah;
    } else {
        sekolah.value = props.storeSekolah;
    }
});
const selectedBentukPendidikan = ref();
const selectedJenjangPendidikan = ref();

const isEdit = ref(false);
const messageError = ref('');
const dialogError = ref(false);
const handleEditSekolah = async () => {
    isEdit.value = !isEdit.value;
    selectedBentukPendidikan.value = sekolah.value?.bentukPendidikanStr;
    selectedJenjangPendidikan.value = sekolah.value?.jenjangPendidikanStr;

    // jenjangPendidikanFiltered.value = await fetchJenjangPendidikan();
    // bentukPendidikanFilter.value = await fetchBe;
};
const getWebsiteUrl = (url) => {
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
        return `https://${url}`; // Tambahkan https jika belum ada
    }
    return url;
};

const isUpdateSekolah = ref(false);
const updateData = async () => {
    isUpdateSekolah.value = true;
};
const handleUpdate = async () => {
    isUpdateSekolah.value = false;
    await updateSekolah(sekolah.value);
};

watch(isError, (val) => {
    if (val) {
        messageError.value = val;
        dialogError.value = true;
    } else {
        messageError.value = '';
        dialogError.value = false;
    }
});
</script>

<template>
    <div class="p-2">
        <div class="mb-4">
            <div class="flex justify-between">
                <h3 class="text-lg font-semibold mb-2">Identitas</h3>
                <Button
                    v-tooltip.bottom="'Edit data sekolah'"
                    :disabled="isFetching"
                    icon="pi pi-pencil"
                    :style="isEdit ? 'background-color:blue;border:none' : 'background-color:gray;border:none'"
                    size="small"
                    rounded=""
                    @click="handleEditSekolah"
                />
            </div>

            <div class="grid grid-cols-2 gap-4">
                <div>Nama</div>
                <div>{{ sekolah?.sekolah.nama }}</div>
                <div>Jenjang</div>
                <div>
                    <div v-if="!isEdit">{{ sekolah?.jenjangPendidikanStr }}</div>
                    <div v-else>
                        <JenjangPendidikanComponent v-model:model-value="selectedJenjangPendidikan" />
                    </div>
                </div>
                <div>Bentuk Pendidikan</div>
                <div>
                    <div v-if="!isEdit">{{ sekolah?.bentukPendidikanStr }}</div>
                    <div v-else>
                        <BentukPendidikanComponent v-model:model-value="selectedBentukPendidikan" />
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
        <div class="mb-4">
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
                    <div v-if="!isEdit">{{ sekolah?.sekolah.nmKepsek }}</div>
                    <div v-else><InputText v-model="sekolah.sekolah.nmKepsek" fluid name="nmKepsek" /></div>
                </div>
                <div>NIP Kepsek</div>
                <div>
                    <div v-if="!isEdit">{{ sekolah?.sekolah.nipKepsek }}</div>
                    <div v-else><InputText v-model="sekolah.sekolah.nipKepsek" fluid name="nipKepsek" /></div>
                </div>
            </div>
        </div>
        <div v-show="isEdit" class="flex justify-end">
            <Button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center" label="Update Data" icon="pi pi-save" :loading="isFetching" @click="updateData" />
        </div>

        <Dialog v-model:visible="isUpdateSekolah" :modal="true" :style="{ width: '450px' }">
            Apakah Data sekolah akan disimpan?
            <template #header><b>Update Data</b></template>
            <template #footer>
                <Button label="Ya" severity="danger" class="w-32" @click="handleUpdate" />
                <Button label="Tidak" class="w-32" @click="isUpdateSekolah = false" />
            </template>
        </Dialog>
        <LoadingOverlay :visible="isFetching" />
        <Dialog v-model:visible="dialogError" :style="{ width: '450px' }" header="Warning" :modal="true" position="top">
            <p>{{ messageError }}</p>
        </Dialog>
    </div>
</template>

<style scoped>
edit-class {
    background-color: red;
}
</style>
