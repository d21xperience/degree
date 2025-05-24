<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import FileUpload from 'primevue/fileupload';
import Select from 'primevue/select';
import { computed, onMounted, ref } from 'vue';
import { useStore } from 'vuex';
const store = useStore();

const { schemaname, listTahunAjaran, fetchTahunAjaran } = useSekolahService();
// ============toast============
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DialogLoading from './DialogLoading.vue';
const toast = useToast();
onMounted(async () => {
    await fetchTahunAjaran();
    // console.log(schemaname.value);
});
const isLoading = ref(false);

const baseUrl = `${import.meta.env.VITE_API_SEKOLAH_BASE_URL}/ss`; //'http://localhost:8183/api/v1/ss'; // Disimpan di child
const templateUrl = computed(() => {
    return `${baseUrl}/download/template?template_type=${props.templateType}&schemaname=${schemaname}&semesterId=${selectedTahunAjaran.value?.tahunAjaranId}`;
});
const selectedTahunAjaran = ref();
// ========================
// Props dari parent
const props = defineProps({
    visible: Boolean,
    templateType: String
});

// Emit event ke parent
const emit = defineEmits(['update:visible', 'save', 'cancel']);

// Menggunakan computed agar bisa mengupdate prop.visible
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});

// Function untuk menutup dialog
const closeDialog = () => {
    isVisible.value = false;
    selectedTahunAjaran.value = null;
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
    const formData = new FormData();
    formData.append('file', file);
    formData.append('upload_type', props.templateType);
    formData.append('schemaname', JSON.stringify(schemaname.value));
    // console.log(schemaname.value);
    // Debugging
    console.log('Upload URL:', uploadUrl);
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

        const result = await response.json();
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

// const upload = () => {
//     uploadedFiles.value.upload();
// };
// // Handle sebelum upload (validasi file)
// const onBeforeUpload = (event) => {
//     const file = event.files[0]; // Ambil file pertama
//     const allowedExtensions = ['xlsx'];
//     const maxFileSize = 2 * 1024 * 1024; // 2MB

//     const fileExtension = file.name.split('.').pop().toLowerCase();
//     if (!allowedExtensions.includes(fileExtension)) {
//         toast.add({ severity: 'error', summary: 'Error', detail: 'Format file harus .xlsx!', life: 3000 });
//         return false;
//     }

//     if (file.size > maxFileSize) {
//         toast.add({ severity: 'error', summary: 'Error', detail: 'Ukuran file tidak boleh lebih dari 2MB!', life: 3000 });
//         return false;
//     }

//     return true; // Izinkan unggahan
// };

// // Handle upload file
// const onUpload = (event) => {
//     console.log('sedang upload files');
//     console.log(event.xhr.response);
//     // uploadedFiles.value = event.files;
//     toast.add({ severity: 'info', summary: 'Success', detail: 'File berhasil diunggah!', life: 3000 });
// };
// Function untuk mengunduh template
const isErr = ref(false);
const submitted = ref(false);
const downloadTemplate = async () => {
    // console.log(!selectedSemester.value);
    // return;
    submitted.value = true;
    if (!selectedTahunAjaran.value) {
        // alert('Pilih tahun pelajaran');
        isErr.value = true;
        return;
    }
    try {
        const response = await fetch(templateUrl.value, {
            method: 'GET',
            headers: {
                Accept: 'application/octet-stream'
            }
        });

        if (!response.ok) {
            throw new Error('Gagal mengunduh file');
        }

        // Coba ambil nama file dari header Content-Disposition
        const contentDisposition = response.headers.get('Content-Disposition');
        let fileName = 'downloaded_file.xlsx'; // Default jika tidak ditemukan
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
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
    }
};
</script>

<template>
    <Toast />
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700"> Tahun Pelajaran <span class="text-red-500">*</span> </label>
                <Select v-model="selectedTahunAjaran" :options="listTahunAjaran.value" optionLabel="nama" placeholder="Pilih Tahun Pelajaran" class="w-full mr-2" :invalid="submitted && !selectedTahunAjaran" />
                <small v-if="submitted && !selectedTahunAjaran" class="text-red-500">Pilih tahun ajaran.</small>
            </div>

            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700"> Unggah File Excel (Pastikan sesuai dengan Template yang disediakan) </label>
                <!-- <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000"
                        :customUpload="true" @before-upload="onBeforeUpload" @upload="onUpload" severity="secondary" />
                </div> -->
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000" :customUpload="true" severity="secondary" />
                </div>
                <p class="mt-2 text-sm text-gray-500">
                    Unduh Template Import data
                    <a href="#" @click.prevent="downloadTemplate" class="text-indigo-600 hover:text-indigo-500">Disini</a>
                </p>
            </div>
        </div>

        <template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Simpan" icon="pi pi-check" text @click="saveData" />
        </template>
    </Dialog>

    <DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>
    <Dialog v-model:visible="isErr" header="Warning!">
        <div>Pilih <b>Tahun Pelajaran</b> terlebih dahulu!</div>
    </Dialog>
</template>
