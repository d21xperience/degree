<script setup>
import { useToast } from 'primevue/usetoast';

import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import LoadingOverlay from '@/components/LoadingOverlay.vue';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useUtils } from '@/composables/useUtils';
import { FilterMatchMode } from '@primevue/core/api';
import { InputNumber } from 'primevue';
import { computed, onMounted, reactive, ref, watch } from 'vue';

const { deleteSemester, updateSemester, fetchSemester } = useSemester();

const { formatterDateID, extractYearAsUint32 } = useUtils();

const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS }
});
const toast = useToast();
const semester = reactive({
    semesterId: '',
    tahunAjaranId: '',
    namaSemester: '',
    semester: '',
    periodeAktif: '',
    tanggalMulai: new Date(),
    tanggalSelesai: new Date()
});
const semesterList = ref([]);
const selectedSemester = ref();
const submitted = ref(false);
const isAddSemester = ref(false);
const isLoading = ref(false);
const handleCreateSemester = () => {
    submitted.value = true;
    if (semester.tanggalMulai && semester.tanggalSelesai) {
        try {
            // Simulasi data baru
            const newSemester = {
                namaSemester: semester.namaSemester,
                tahunAjaranId: semester.tahunAjaranId,
                semester: semester.semester,
                periodeAktif: 1,
                semesterId: semester.semesterId,
                tanggalMulai: semester.tanggalMulai.toISOString().split('T')[0],
                tanggalSelesai: semester.tanggalSelesai.toISOString().split('T')[0]
            };

            console.log(newSemester);
            // return
            semesterList.value.unshift(newSemester); // ✅ tambah ke awal list

            toast.add({ severity: 'success', summary: 'Sukses', detail: 'Semester ditambahkan', life: 3000 });
            isAddSemester.value = false;
            resetValue();
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Gagal', detail: 'Error saat tambah', life: 3000 });
        }
    }
};

const isDeleteDialog = ref(false);
const handleDeleteSemester = async () => {
    try {
        const originalSemesterList = [...semesterList.value];
        const semesterDelete = originalSemesterList.filter((val) => selectedSemester.value?.includes(val));
        const res = await deleteSemester(semesterDelete);
        if (res.status) {
            semesterList.value = semesterList.value.filter((val) => !semesterDelete.includes(val));
            toast.add({
                severity: 'success',
                summary: 'Successful',
                detail: `${res.message}`,
                life: 3000
            });
        }
    } catch (error) {
        console.log('Error saat hapus:', error);
        toast.add({
            severity: 'error',
            summary: 'Failed',
            detail: `Failed: ${error}`,
            life: 3000
        });
    } finally {
        isDeleteDialog.value = false;
        selectedSemester.value = [];
    }
};

const handleCloseDialog = () => {
    isDeleteDialog.value = false;
    selectedSemester.value = null;
};
const judul = computed(() => (isEdit.value ? 'Edit Semester' : 'Tambah Semester Baru'));
const isEdit = ref(false);
const editDialog = () => {
    isEdit.value = true;
    const selected = selectedSemester.value[0];
    Object.assign(semester, selected);
    // semester.tahunAjaranId = new Date(selected.tahunAjaranId.toString());
    semester.tanggalMulai = selected.tanggalMulai ? new Date(selected.tanggalMulai) : '';
    semester.tanggalSelesai = selected.tanggalSelesai ? new Date(selected.tanggalSelesai) : '';
    isAddSemester.value = true;
};
const handleUpdateSemester = async () => {
    try {
        // semester.tahunAjaranId = extractYearAsUint32(semester.tahunAjaranId);
        const originalSemesterList = { ...semester };
        originalSemesterList.tahunAjaranId = extractYearAsUint32(semester.tahunAjaranId);
        const res = await updateSemester(originalSemesterList);

        if (res.status) {
            // semesterList.value = semesterList.value.map((val) => {
            //     if (val.semesterId === semester.semesterId) {
            //         return { ...val, ...semester }; // update data semester yang match
            //     }
            //     return val; // sisanya tetap
            // });
            toast.add({
                severity: 'success',
                summary: 'Successful',
                detail: `${res.message}`,
                life: 3000
            });
            selectedSemester.value = null; // reset selected
        }
    } catch (error) {
        console.log('Error saat update:', error);
        toast.add({
            severity: 'error',
            summary: 'Failed',
            detail: `Failed: ${error}`,
            life: 3000
        });
    } finally {
        isAddSemester.value = false;
    }
};

const initial = async () => {
    isLoading.value = true;
    try {
        semesterList.value = await fetchSemester();
        // console.log(semesterList.value);
    } catch (error) {
        alert('error');
    } finally {
        isLoading.value = false;
    }
};

const resetValue = () => {
    isEdit.value = false;
    isTahunAjaranAvailable.value = true;
    submitted.value = false;
    Object.keys(semester).forEach((key) => {
        semester[key] = ''; // semua jadi string kosong
    });
    selectedSemester.value = null;
};
const getActive = (value) => {
    return value === 1;
};

const setActive = (data, newValue) => {
    data.periodeAktif = newValue ? 1 : 0;
    // Opsional: panggil API atau emit event jika perlu
    alert('hello world!');
    // updatePeriodeAktif(data.semesterId, data.periodeAktif)
};

const setDefaultTanggal = (tahun) => {
    console.log(tahun);
    if (!tahun) return;
    if (semester.semester === 1) {
        semester.tanggalMulai = new Date(`${tahun - 1}-07-07`); // 07 Juli thn ajaran
        semester.tanggalSelesai = new Date(`${tahun - 1}-12-12`); // 12 Desember thn ajaran
        semester.namaSemester = `${tahun} / ${tahun + 1} Ganjil`;
        semester.semesterId = `${tahun}1`;
        semester.tahunAjaranId = `${tahun}`;
    } else if (semester.semester === 2) {
        semester.tanggalMulai = new Date(`${tahun}-01-01`); // 01 Januari tahun berikutnya
        semester.tanggalSelesai = new Date(`${tahun}-06-06`); // 06 Juni tahun berikutnya
        semester.namaSemester = `${tahun} / ${tahun + 1} Genap`;
        semester.semesterId = `${tahun}2`;
        semester.tahunAjaranId = `${tahun}`;
    }
};
watch(
    () => semester.semester,
    () => {
        // console.log(newVal);
        if (isEdit.value && semester.tahunAjaranId) {
            const year = new Date(semester.tahunAjaranId).getFullYear();
            setDefaultTanggal(year);
        }
    }
);

const semesterOptions = ref([
    {
        value: 1,
        label: 'Semester 1'
    },
    {
        value: 2,
        label: 'Semester 2'
    }
]);

const isTahunAjaranAvailable = ref(true);
const onTahunAjaranChange = (val) => {
    if (!val) return;
    isTahunAjaranAvailable.value = false;
};

onMounted(async () => {
    initial();
});
</script>

<template>
    <div class="">
        <Toolbar>
            <template #start>
                <div>
                    <Button
                        icon="pi pi-plus"
                        severity="success"
                        class="mr-2 text-lg"
                        @click="isAddSemester = true"
                        v-tooltip.bottom="'Tambah Semester'"
                        :disabled="selectedSemester && !(Array.isArray(selectedSemester) && selectedSemester.length === 0)"
                    />

                    <Button icon="pi pi-pencil" severity="warn" @click="editDialog" :disabled="!selectedSemester || !selectedSemester.length || selectedSemester.length > 1" class="mr-2" v-tooltip.bottom="'Edit Semester'" />
                    <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg" @click="isDeleteDialog = true" :disabled="!selectedSemester || !selectedSemester.length" v-tooltip.bottom="'Hapus Semester'" />
                </div>
            </template>
            <template #end>
                <IconField class="mr-2">
                    <InputIcon>
                        <i class="pi pi-search" />
                    </InputIcon>
                    <InputText v-model="filters['global'].value" placeholder="Search..." :disabled="Array.isArray(semesterList) && semesterList.length === 0" />
                </IconField>
                <Button icon="pi pi-refresh" severity="help" class="mr-2 text-lg" @click="initial" v-tooltip.bottom="'Refresh'" />
            </template>
        </Toolbar>
        <div v-if="!semesterList || (Array.isArray(semesterList) && semesterList.length === 0)">
            <div class="flex justify-center">
                <h3>Tidak ada data semester</h3>
            </div>
        </div>
        <DataTable
            v-else
            ref="dt"
            v-model:selection="selectedSemester"
            stripedRows
            size="small"
            :value="semesterList"
            dataKey="semesterId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} Semester"
        >
            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="namaSemester" header="Nama"> </Column>
            <Column field="tahunAjaranId" header="Tahun Ajaran" sortable> </Column>
            <Column field="semester" header="Semester"> </Column>
            <Column field="tanggalMulai" header="Tanggal Mulai">
                <template #body="slotProps">
                    {{ formatterDateID(slotProps.data.tanggalMulai) }}
                    <!-- {{ slotProps.data.tanggalMulai }} -->
                </template>
            </Column>
            <Column field="tanggalSelesai" header="Tanggal Selesai">
                <template #body="slotProps">
                    {{ formatterDateID(slotProps.data.tanggalSelesai) }}
                    <!-- {{ slotProps.data.tanggalSelesai }} -->
                </template>
            </Column>
            <Column field="periodeAktif" header="Status">
                <template #body="slotProps">
                    <ToggleButton :modelValue="getActive(slotProps.data.periodeAktif)" @update:modelValue="setActive(slotProps.data, $event)" class="w-24" onLabel="On" offLabel="Off" />
                </template>
            </Column>
        </DataTable>

        <!-- Dialog Tambah dan edit semester -->
        <Dialog v-model:visible="isAddSemester" :style="{ width: '450px' }" :header="judul" :modal="true" position="top" @hide="resetValue">
            <div class="grid gap-y-2">
                <div>
                    <label for="thn-ajaran">Tahun Ajaran</label>
                    <DatePicker v-model="semester.tahunAjaranId" view="year" dateFormat="yy" :inline="false" fluid @update:modelValue="onTahunAjaranChange" />
                </div>
                <div v-if="isEdit == true" class="grid gap-y-2">
                    <div class="" v-show="isEdit == true">
                        <label for="mata-pelajaran">Nama </label>
                        <InputText fluid v-model="semester.namaSemester" />
                        <!-- <small v-if="submitted && !selectedMapel" class="text-red-500">Subject is required.</small> -->
                    </div>
                    <div>
                        <label for="mata-pelajaran">Semester</label>
                        <InputNumber fluid :min="1" :max="2" v-model="semester.semester" />
                        <!-- <small v-if="submitted && !selectedSemester" class="text-red-500">Teacher is required.</small> -->
                    </div>
                    <div>
                        <label for="tgl-mulai">Tanggal Mulai</label>
                        <DatePicker v-model="semester.tanggalMulai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !semester.tanggalMulai" />
                        <small v-if="submitted && !semester.tanggalMulai" class="text-red-500">Tanggal mulai harus diisi.</small>
                    </div>
                    <div>
                        <label for="tgl-mulai">Tanggal Selesai</label>
                        <DatePicker v-model="semester.tanggalSelesai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !semester.tanggalSelesai" />
                        <small v-if="submitted && !semester.tanggalSelesai" class="text-red-500">Tanggal selesai harus diisi.</small>
                    </div>
                </div>
                <div v-else class="grid gap-y-2">
                    <div class="flex space-x-2">
                        <div class="w-full">
                            Semester
                            <Select :options="semesterOptions" optionLabel="label" optionValue="value" fluid :disabled="isTahunAjaranAvailable" v-model:modelValue="semester.semester" />
                        </div>
                        <div class="w-full">
                            Nama
                            <InputText fluid :disabled="isTahunAjaranAvailable" v-model:modelValue="semester.namaSemester" />
                        </div>
                    </div>
                    <div class="flex space-x-2">
                        <div>
                            <label for="tgl-mulai">Tanggal Mulai</label>
                            <DatePicker v-model="semester.tanggalMulai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !semester.tanggalMulai" :disabled="isTahunAjaranAvailable" />
                            <small v-if="submitted && !semester.tanggalMulai" class="text-red-500">Tanggal mulai harus diisi.</small>
                        </div>
                        <div>
                            <label for="tgl-mulai">Tanggal Selesai</label>
                            <DatePicker v-model="semester.tanggalSelesai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !semester.tanggalSelesai" :disabled="isTahunAjaranAvailable" />
                            <small v-if="submitted && !semester.tanggalSelesai" class="text-red-500">Tanggal selesai harus diisi.</small>
                        </div>
                    </div>
                </div>
            </div>
            <template #footer>
                <Button v-if="isEdit == true" label="Update" icon="pi pi-save" @click="handleUpdateSemester" />
                <Button v-else label="Tambah" icon="pi pi-save" @click="handleCreateSemester" />
            </template>
        </Dialog>

        <!-- end of import data -->
        <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay>

        <DialogConfirmDelete message="data ini akan dihapus?" v-model:visible="isDeleteDialog" @confirm="handleDeleteSemester" @closeDialog="handleCloseDialog" judul="Hapus semester" />
    </div>
</template>
