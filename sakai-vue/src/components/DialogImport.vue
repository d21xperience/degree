<script setup>
import FileUpload from 'primevue/fileupload';
import { computed, onMounted, ref, watch } from 'vue';

const { schemaname } = useTableTenant();
const { initSelectedSemester } = useSemester();
// ============toast============
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DialogLoading from './DialogLoading.vue';
const toast = useToast();

const isLoading = ref(false);

const baseUrl = `${import.meta.env.VITE_API_SEKOLAH_BASE_URL}/ss`;
const templateUrl = computed(() => {
    return `${baseUrl}/download/template?template_type=${props.templateType}&schemaname=${schemaname.value}&semesterId=${selectedTahunAjaran.value?.tahunAjaranId}`;
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
    // console.log(file.name.includes(props.templateType) && file.name.includes(`${selectedTahunAjaran.value.tahunAjaranId}`))
    // Cek file yang akan diupload apakah sudah sesuai dengan ketentuan
    if (!(file.name.includes(props.templateType) && file.name.includes(`${selectedTahunAjaran.value.tahunAjaranId}`))) {
        toast.add({ severity: 'warn', summary: 'Gagal', detail: 'Silakan unggah file sesuai dengan template yang telah disediakan!', life: 3000 });
        uploadedFiles.value.files = null;
        return;
    }
    const formData = new FormData();
    formData.append('file', file);
    formData.append('upload_type', props.templateType);
    formData.append('schemaname', JSON.stringify(schemaname.value));
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
        console.log(error);
        toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
    }
};
watch(initSelectedSemester, (newVal) => {
    // console.log("newVal",newVal)
    selectedTahunAjaran.value = newVal;
    console.log('selectedTahun ', selectedTahunAjaran.value);
});
onMounted(async () => {
    // await fetchTahunAjaran();
    // console.log(initSelectedSemester.value);
    selectedTahunAjaran.value = initSelectedSemester.value;
    // console.log(selectedTahunAjaran.value);
    // console.log(`${selectedTahunAjaran.value?.tahunAjaranId}`)
});
</script>

<template>
    <Toast />
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700"> Unggah File Excel (Pastikan sesuai dengan Template yang disediakan) </label>
                <!-- <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000"
                        :customUpload="true" @before-upload="onBeforeUpload" @upload="onUpload" severity="secondary" />
                </div> -->
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000" :customUpload="true" severity="secondary" />
                </div>
            </div>
            <div class="mb-4 flex justify-between">
                <div class="mt-2 text-sm text-gray-500">
                    Unduh Template Import data
                    <a href="#" @click.prevent="downloadTemplate" class="text-indigo-600 hover:text-indigo-500"
                        >disini <span class="text-gray-500">untuk tahun ajaran {{ selectedTahunAjaran.namaSemester }}</span></a
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
