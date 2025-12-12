<script setup>
import FileUpload from 'primevue/fileupload';
import { computed, ref } from 'vue';

import store from '@/store';
import { useToast } from 'primevue/usetoast';
import DialogLoading from './DialogLoading.vue';
const toast = useToast();
const schemaname = computed(() => {
    return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
});
// ========================
// Props dari parent
const props = defineProps({
    visible: Boolean,
    templateType: {
        type: String,
        required: true
    },
    selectedSemester: {
        type: Object,
        required: true
    },
    paramNilai: {
        type: Object,
        default: null
    },
    message: {
        type: String,
        default: ''
    }
});
// Emit event ke parent
const emit = defineEmits(['update:visible', 'save', 'cancel']);
const isLoading = ref(false);
const selectedTahunAjaran = computed(() => props.selectedSemester?.tahunAjaranId);
// let rombelQuery = '';
const semesterAktif = computed(() => {
    let tes = '';
    switch (props.templateType) {
        case 'siswa':
            tes = selectedTahunAjaran.value;
            break;
        default:
            tes = props.selectedSemester?.semesterId;
            break;
    }
    return tes;
});

const baseUrl = `${import.meta.env.VITE_API_BASE_URL}/ss`;
const templateUrl = computed(() => {
    return `${baseUrl}/download/template?template_type=${props.templateType}&schemaname=${schemaname.value}&semesterId=${semesterAktif.value}`;
});

// Menggunakan computed agar bisa mengupdate prop.visible
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});

// Function untuk menutup dialog
const closeDialog = () => {
    isVisible.value = false;
};

// Refs untuk FileUpload dan file yang diunggah
// const fileupload = ref();
const uploadedFiles = ref();
// const uploadUrl = `${baseUrl}/upload/rest?upload_type=${props.templateType}`
const uploadUrl = `${baseUrl}/upload/rest`;
const saveData = async () => {
    if (uploadedFiles.value.files.length == 0) {
        toast.add({ severity: 'warn', summary: 'Gagal', detail: 'Silakan unggah file terlebih dahulu!', life: 3000 });
        return;
    }

    const file = uploadedFiles.value.files[0];
    // console.log(file);
    // console.log(file.name.includes(props.templateType) && file.name.includes(`${props.selectedTahunAjaran}`))
    // Cek file yang akan diupload apakah sudah sesuai dengan ketentuan
    if (!(file.name.includes(props.templateType) && file.name.includes(semesterAktif))) {
        toast.add({ severity: 'warn', summary: 'Gagal', detail: 'Silakan unggah file sesuai dengan template yang telah disediakan!', life: 3000 });
        uploadedFiles.value.files = null;
        return;
    }
    const formData = new FormData();
    formData.append('file', file);
    formData.append('upload_type', props.templateType);
    formData.append('schemaname', JSON.stringify(schemaname));
    // console.log('Upload URL:', uploadUrl);
    for (let pair of formData.entries()) {
        console.log(pair[0] + ': ', pair[1]);
    }

    isLoading.value = true;

    try {
        const response = await fetch(uploadUrl, {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error('Gagal mengunggah file');
        }

        // const result = await response.json();
        toast.add({ severity: 'success', summary: 'Sukses', detail: 'File berhasil diunggah!', life: 3000 });
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal mengunggah file', life: 3000 });
        console.error('Upload error:', error);
    } finally {
        isLoading.value = false;
        if (uploadedFiles.value) {
            uploadedFiles.value.clear();
        }
        emit('save', uploadedFiles.value);
        closeDialog();
    }
};
// Function untuk mengunduh template
const isErr = ref(false);
const submitted = ref(false);
const downloadTemplate = async () => {
    submitted.value = true;
    if (!semesterAktif.value) {
        isErr.value = true;
        return;
    }

    try {
        // Siapkan data yang akan dikirim
        let requestData = {};
        console.log(props.paramNilai);
        if (props.paramNilai) {
            requestData = props.paramNilai;
        }
        // return;
        // Ganti GET dengan POST dan kirim data di body
        const response = await fetch(templateUrl.value, {
            method: 'POST', // Ubah ke POST
            headers: {
                'Content-Type': 'application/json',
                Accept: 'application/octet-stream'
            },
            body: JSON.stringify(requestData) // Kirim data di body
        });

        if (!response.ok) {
            throw new Error('Gagal mengunduh file');
        }

        // Coba ambil nama file dari header Content-Disposition
        const contentDisposition = response.headers.get('Content-Disposition');
        let fileName = 'template_nilai.xlsx';
        if (contentDisposition) {
            const match = contentDisposition.match(/filename="([^"]+)"/);
            if (match && match[1]) {
                fileName = match[1];
            }
        }

        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = fileName;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);

        window.URL.revokeObjectURL(url);

        // Reset loading state jika ada
        submitted.value = false;
    } catch (error) {
        console.log(error);
        submitted.value = false;
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Terjadi kesalahan saat mengunduh file',
            life: 3000
        });
    }
};
</script>

<template>
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label class="block text-sm font-bold text-gray-700"> Unggah File Excel (Pastikan sesuai dengan Template yang disediakan) </label>
                <p class="text-sm">Upload masal {{ props.message }}</p>
                <!-- <p class="text-sm">Upload banyak siswa hanya untuk semester 1. Jika pada semester 2 ada penambahan siswa tambahkan secara manual.</p> -->
                <!-- <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000"
                        :customUpload="true" @before-upload="onBeforeUpload" @upload="onUpload" severity="secondary" />
                </div> -->
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :max-file-size="2000000" :custom-upload="true" severity="secondary" />
                </div>
            </div>
            <div class="mb-4 flex justify-between">
                <div class="mt-2 text-sm text-gray-500">
                    Unduh Template Import data
                    <a href="#" class="text-indigo-600 hover:text-indigo-500" @click.prevent="downloadTemplate"
                        >disini <span class="text-gray-500">untuk tahun ajaran {{ selectedTahunAjaran }}</span></a
                    >
                </div>
                <!-- <div class="">
                    <Select v-model="selectedTahunAjaran" :options="listTahunAjaran.value" optionLabel="nama" placeholder="Pilih Tahun Pelajaran" class="text-sm" :invalid="submitted && !selectedTahunAjaran" fluid />
                    <small v-if="submitted && !selectedTahunAjaran" class="text-red-500">Pilih tahun ajaran.</small>
                </div> -->
            </div>
        </div>

        <template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Upload" icon="pi pi-upload" severity="warn" text @click="saveData" />
        </template>
    </Dialog>

    <DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>
    <Dialog v-model:visible="isErr" header="Warning!">
        <div>Pilih <b>Tahun Pelajaran</b> terlebih dahulu!</div>
    </Dialog>
</template>
