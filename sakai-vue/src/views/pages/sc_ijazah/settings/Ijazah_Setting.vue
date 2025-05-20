<template>
    <div class="card">
        <div class="lg:flex lg:justify-between">
            <h4>Seting Pengisian Ijazah & Transkrip Nilai</h4>
            <div class="md:flex md:items-center md:space-x-2">
                <label class="text-slate-500 md:text-base text-sm">Tahun Ajaran</label>
                <div>
                    <Select v-model="selectedTahunAjaran" :options="listTahunAjaran" optionLabel="nama" placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />
                </div>
            </div>
        </div>
        <section class="lg:mt-0 mt-10">
            <h5 class="text-xl font-semibold mb-4">Informasi Penerbitan Ijazah</h5>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <!-- <div>
                    <label class="block" for="prov">Provinsi</label>
                    <InputText v-model="ijazahParam.prov" fluid name="prov" id="prov" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block" for="kab">Kabupaten</label>
                    <InputText v-model="ijazahParam.kab" fluid name="kab" id="kab" placeholder="Masukan nama" />
                </div> -->
                <div>
                    <label class="block" for="tpt-ijazah">Tempat Penerbitan</label>
                    <InputText v-model="ijazahParam.tempat_dikeluarkan_ijazah" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block" for="tgl-dikelurkan-ijazah">Tgl. Penerbitan</label>
                    <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" v-model="ijazahParam.tgl_dikeluarkan_ijazah" />
                </div>
                <div>
                    <label class="block" for="ptk-id">Kepala Sekolah</label>
                    <AutoComplete v-model="selectedPtk" dropdown :suggestions="ptk" optionLabel="nama" fluid @complete="getPtk($event)" @keydown.space.prevent="handleKeydown"   />
                </div>
            </div>
        </section>
        <section>
            <h5 class="text-xl font-semibold mb-4">Informasi Transkrip</h5>
            <div class="flex flex-col mb-2 card border-2 space-x-2">
                <div class="w-full">
                    <label class="block" for="prov">Kop Sekolah</label>
                    <FileUpload name="demo[]" @uploader="onUpload" :multiple="true" accept="image/*" :maxFileSize="1000000" customUpload />
                </div>
                <div class="w-full">
                    <p>Preview</p>
                    <div class="text-center">
                        <Image alt="Image" preview>
                            <template #previewicon>
                                <i class="pi pi-search"></i>
                            </template>
                            <template #image>
                                <img src="https://primefaces.org/cdn/primevue/images/galleria/galleria11.jpg" alt="image" width="250" />
                            </template>
                            <template #preview="slotProps">
                                <img src="https://primefaces.org/cdn/primevue/images/galleria/galleria11.jpg" alt="preview" :style="slotProps.style" @click="slotProps.onClick" />
                            </template>
                        </Image>
                    </div>
                </div>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <!-- <div>
                    <label class="block" for="no-transkrip">No Transkrip</label>
                    <InputText v-model="ijazahParam.kab" fluid name="no-transkrip" id="no-transkrip" placeholder="Masukan No. Transkrip" />
                </div> -->
                <!-- <div>
                    <label class="block" for="tpt-ijazah">Tempat Ijazah</label>
                    <InputText v-model="ijazahParam.tempat_dikeluarkan_ijazah" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block" for="tgl-dikelurkan-ijazah">Tgl dikeluarkan Ijazah</label>
                    <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded" v-model="ijazahParam.tgl_dikeluarkan_ijazah" />
                </div>
                <div>
                    <label class="block" for="tgl-dikelurkan-ijazah">Kepala Sekolah</label>
                    <InputText v-model="ijazahParam.ptk_id" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Masukan nama" />
                </div> -->
            </div>
        </section>
        <div class="flex justify-end space-x-4">
            <Button label="Simpan" severity="success" @click="save" :loading="isLoading" class="min-w-28" />
            <Button label="Batal" severity="secondary" @click="batal" class="min-w-28" />
        </div>
    </div>
</template>

<script setup>
import { useFormOptions } from '@/composables/useFormOptions';
import { useSekolahService } from '@/composables/useSekolahService';
import InputText from 'primevue/inputtext';
import { useToast } from 'primevue/usetoast';
import { onMounted, ref } from 'vue';
const { listTahunAjaran, selectedTahunAjaran, initSelectedTahunAjaran, createInfoIjazah, fetchGuruTerdaftar, guruTerdaftarList } = useSekolahService();
const { handleKeydown } = useFormOptions();
const loading = ref(false);
const isLoading = ref(false);
const selectedPtk = ref();
const ptk = ref();
const ptkList = ref();
const ijazahParam = ref({
    tahun_ajaran_id: initSelectedTahunAjaran.value?.tahunAjaranId,
    tempat_dikeluarkan_ijazah: '',
    tgl_dikeluarkan_ijazah: '',
    ptk_id: '',
    kop_sekolah_url: ''
});
onMounted(async () => {
    selectedTahunAjaran.value = initSelectedTahunAjaran.value;
    await fetchGuruTerdaftar();
    ptkList.value = guruTerdaftarList.value.map((item) => item.ptk);
    // console.log(tes)
    // console.log(ptk.value);
});

const toast = useToast();
const fileupload = ref();

function upload() {
    fileupload.value.upload();
}

function onUpload() {
    toast.add({ severity: 'info', summary: 'Success', detail: 'File Uploaded', life: 3000 });
}

const save = () => {
    createInfoIjazah(ijazahParam.value);
};
const batal = () => {};
const getPtk = (event) => {
    setTimeout(() => {
        if (!event.query.trim().length) {
            ptk.value = ptkList.value;
        } else {
            ptk.value = ptkList.value.filter((item) => item.nama.toLowerCase().includes(event.query.toLowerCase()));
        }
    }, 250);
};
</script>
