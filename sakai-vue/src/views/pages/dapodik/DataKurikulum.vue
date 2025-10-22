<script setup>
import { useToast } from 'primevue/usetoast';

import JenjangPendidikanComponent from '@/components/sekolah_components/JenjangPendidikanComponent.vue';
import { useKurikulum } from '@/composables/sekolah_composable/useKurikulum';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useUtils } from '@/composables/useUtils';
import { FilterMatchMode } from '@primevue/core/api';
import { computed, onMounted, reactive, ref, watch } from 'vue';

const { getKurikulum } = useKurikulum();
const { formatterDateID } = useUtils();
const { sekolah } = useSekolah();
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS }
});
const toast = useToast();
const kurikulum = reactive({
    kurikulumId: '',
    namakurikulum: '',
    mulaiBerlaku: '',
    jenjangPendidikanId: '',
    sistemSks: '',
    totalSks: '',
    jurusanId: ''
});
const kurikulumList = ref([]);
const selectedkurikulum = ref();
const submitted = ref(false);
const isAddkurikulum = ref(false);
const isLoading = ref(false);
// const handleCreatekurikulum = () => {
//     submitted.value = true;
//     if (kurikulum.tanggalMulai && kurikulum.tanggalSelesai) {
//         try {
//             // Simulasi data baru
//             const newkurikulum = {
//                 namakurikulum: kurikulum.namakurikulum,
//                 tahunAjaranId: kurikulum.tahunAjaranId,
//                 kurikulum: kurikulum.kurikulum,
//                 periodeAktif: 1,
//                 kurikulumId: kurikulum.kurikulumId
//                 // tanggalMulai: kurikulum.tanggalMulai.toISOString().split('T')[0],
//                 // tanggalSelesai: kurikulum.tanggalSelesai.toISOString().split('T')[0]
//             };

//             console.log(newkurikulum);
//             // return
//             kurikulumList.value.unshift(newkurikulum); // ✅ tambah ke awal list

//             toast.add({ severity: 'success', summary: 'Sukses', detail: 'kurikulum ditambahkan', life: 3000 });
//             isAddkurikulum.value = false;
//             resetValue();
//         } catch (error) {
//             toast.add({ severity: 'error', summary: 'Gagal', detail: 'Error saat tambah', life: 3000 });
//         }
//     }
// };

const isDeleteDialog = ref(false);
// const handleDeletekurikulum = async () => {
//     try {
//         const originalkurikulumList = [...kurikulumList.value];
//         const kurikulumDelete = originalkurikulumList.filter((val) => selectedkurikulum.value?.includes(val));
//         const res = await deletekurikulum(kurikulumDelete);
//         if (res.status) {
//             kurikulumList.value = kurikulumList.value.filter((val) => !kurikulumDelete.includes(val));
//             toast.add({
//                 severity: 'success',
//                 summary: 'Successful',
//                 detail: `${res.message}`,
//                 life: 3000
//             });
//         }
//     } catch (error) {
//         console.log('Error saat hapus:', error);
//         toast.add({
//             severity: 'error',
//             summary: 'Failed',
//             detail: `Failed: ${error}`,
//             life: 3000
//         });
//     } finally {
//         isDeleteDialog.value = false;
//         selectedkurikulum.value = [];
//     }
// };

const handleCloseDialog = () => {
    isDeleteDialog.value = false;
    selectedkurikulum.value = null;
};
const judul = computed(() => (isEdit.value ? 'Edit kurikulum' : 'Tambah kurikulum Baru'));
const isEdit = ref(false);
const editDialog = () => {
    isEdit.value = true;
    const selected = selectedkurikulum.value[0];
    Object.assign(kurikulum, selected);
    // kurikulum.tahunAjaranId = new Date(selected.tahunAjaranId.toString());
    kurikulum.tanggalMulai = selected.tanggalMulai ? new Date(selected.tanggalMulai) : '';
    kurikulum.tanggalSelesai = selected.tanggalSelesai ? new Date(selected.tanggalSelesai) : '';
    isAddkurikulum.value = true;
};
// const handleUpdatekurikulum = async () => {
//     try {
//         // kurikulum.tahunAjaranId = extractYearAsUint32(kurikulum.tahunAjaranId);
//         const originalkurikulumList = { ...kurikulum };
//         originalkurikulumList.tahunAjaranId = extractYearAsUint32(kurikulum.tahunAjaranId);
//         const res = await updatekurikulum(originalkurikulumList);

//         if (res.status) {
//             // kurikulumList.value = kurikulumList.value.map((val) => {
//             //     if (val.kurikulumId === kurikulum.kurikulumId) {
//             //         return { ...val, ...kurikulum }; // update data kurikulum yang match
//             //     }
//             //     return val; // sisanya tetap
//             // });
//             toast.add({
//                 severity: 'success',
//                 summary: 'Successful',
//                 detail: `${res.message}`,
//                 life: 3000
//             });
//             selectedkurikulum.value = null; // reset selected
//         }
//     } catch (error) {
//         console.log('Error saat update:', error);
//         toast.add({
//             severity: 'error',
//             summary: 'Failed',
//             detail: `Failed: ${error}`,
//             life: 3000
//         });
//     } finally {
//         isAddkurikulum.value = false;
//     }
// };

const initial = async () => {
    isLoading.value = true;
    try {
        console.log(sekolah);
        // kurikulumList.value = await getKurikulum(jenjangPendidikanId);
        // console.log(kurikulumList.value);
    } catch (error) {
        console.log('Error saat inisialisasi:', error);
        toast.add({
            severity: 'error',
            summary: 'Failed',
            detail: `Failed: ${error}`,
            life: 3000
        });
    } finally {
        isLoading.value = false;
    }
};

const resetValue = () => {
    isEdit.value = false;
    // isTahunAjaranAvailable.value = true;
    submitted.value = false;
    Object.keys(kurikulum).forEach((key) => {
        kurikulum[key] = ''; // semua jadi string kosong
    });
    selectedkurikulum.value = null;
};
// const getActive = (value) => {
//     return value === 1;
// };

// const setActive = (data, newValue) => {
//     data.periodeAktif = newValue ? 1 : 0;
//     // Opsional: panggil API atau emit event jika perlu
//     alert('hello world!');
//     // updatePeriodeAktif(data.kurikulumId, data.periodeAktif)
// };

// const setDefaultTanggal = (tahun) => {
//     console.log(tahun);
//     if (!tahun) return;
//     if (kurikulum.kurikulum === 1) {
//         kurikulum.tanggalMulai = new Date(`${tahun - 1}-07-07`); // 07 Juli thn ajaran
//         kurikulum.tanggalSelesai = new Date(`${tahun - 1}-12-12`); // 12 Desember thn ajaran
//         kurikulum.namakurikulum = `${tahun} / ${tahun + 1} Ganjil`;
//         kurikulum.kurikulumId = `${tahun}1`;
//         kurikulum.tahunAjaranId = `${tahun}`;
//     } else if (kurikulum.kurikulum === 2) {
//         kurikulum.tanggalMulai = new Date(`${tahun}-01-01`); // 01 Januari tahun berikutnya
//         kurikulum.tanggalSelesai = new Date(`${tahun}-06-06`); // 06 Juni tahun berikutnya
//         kurikulum.namakurikulum = `${tahun} / ${tahun + 1} Genap`;
//         kurikulum.kurikulumId = `${tahun}2`;
//         kurikulum.tahunAjaranId = `${tahun}`;
//     }
// };
// watch(
//     () => kurikulum.kurikulum,
//     () => {
//         console.log(kurikulum.kurikulum);
//         if (!isEdit.value && kurikulum.tahunAjaranId) {
//             const year = new Date(kurikulum.tahunAjaranId).getFullYear();
//             setDefaultTanggal(year);
//         }
//     }
// );

// const kurikulumOptions = ref([
//     {
//         value: 1,
//         label: 'kurikulum 1'
//     },
//     {
//         value: 2,
//         label: 'kurikulum 2'
//     }
// ]);

// const isTahunAjaranAvailable = ref(true);
// const onTahunAjaranChange = (val) => {
//     if (!val) return;
//     isTahunAjaranAvailable.value = false;
// };
const jenjangPendidikan = ref();
watch(jenjangPendidikan, async (newVal) => {
    if (newVal) {
        const response = await getKurikulum(newVal?.jenjangPendidikanId);
        if (response.status) {
            kurikulumList.value = response.kurikulum;
            toast.add({ severity: 'success', summary: 'Sukses', detail: `${response.message}`, life: 3000 });
        }
    }
});
onMounted(async () => {
    // initial();
});
</script>

<template>
    <div class="">
        <div class="flex justify-between items-center">
            <h3>Data Kurikulum</h3>
            <JenjangPendidikanComponent class="max-w-[215px]" v-model="jenjangPendidikan" />
        </div>

        <div v-if="!jenjangPendidikan">Silahkan pilih jenjang pendidikan terlebih dahulu</div>
        <div v-else>
            <Toolbar>
                <template #start>
                    <div class="flex">
                        <Button
                            icon="pi pi-plus"
                            severity="success"
                            class="mr-2 text-lg"
                            @click="isAddkurikulum = true"
                            v-tooltip.bottom="'Tambah kurikulum'"
                            :disabled="selectedkurikulum && !(Array.isArray(selectedkurikulum) && selectedkurikulum.length === 0)"
                        />

                        <Button icon="pi pi-pencil" severity="warn" @click="editDialog" :disabled="!selectedkurikulum || !selectedkurikulum.length || selectedkurikulum.length > 1" class="mr-2" v-tooltip.bottom="'Edit kurikulum'" />
                        <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg" @click="isDeleteDialog = true" :disabled="!selectedkurikulum || !selectedkurikulum.length" v-tooltip.bottom="'Hapus kurikulum'" />
                    </div>
                </template>
                <template #end>
                    <IconField class="mr-2">
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." :disabled="Array.isArray(kurikulumList) && kurikulumList.length === 0" />
                    </IconField>
                    <Button icon="pi pi-refresh" severity="help" class="mr-2 text-lg" @click="initial" v-tooltip.bottom="'Refresh'" />
                </template>
            </Toolbar>
            <div v-if="!kurikulumList || (Array.isArray(kurikulumList) && kurikulumList.length === 0)" class="h-96">
                <div class="flex justify-center">
                    <h3>Tidak ada data kurikulum</h3>
                </div>
            </div>
            <DataTable
                v-else
                ref="dt"
                v-model:selection="selectedkurikulum"
                stripedRows
                size="small"
                :value="kurikulumList"
                dataKey="kurikulumId"
                :paginator="true"
                :rows="10"
                :filters="filters"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[10, 20, 50]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kurikulum"
            >
                <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                <Column field="namaKurikulum" header="Nama"> </Column>
                <Column field="tanggalMulai" header="Mulai Berlaku">
                    <template #body="slotProps">
                        {{ formatterDateID(slotProps.data.mulaiBerlaku) }}
                    </template>
                </Column>
            </DataTable>
        </div>

        <!-- Dialog Tambah dan edit kurikulum -->
        <Dialog v-model:visible="isAddkurikulum" :style="{ width: '450px' }" :header="judul" :modal="true" position="top" @hide="resetValue">
            <div class="grid gap-y-2">
                <div>
                    <label for="thn-ajaran">Tahun Ajaran</label>
                    <DatePicker v-model="kurikulum.tahunAjaranId" view="year" dateFormat="yy" :inline="false" fluid @update:modelValue="onTahunAjaranChange" />
                </div>
                <div v-if="isEdit == true" class="grid gap-y-2">
                    <div class="" v-show="isEdit == true">
                        <label for="mata-pelajaran">Nama </label>
                        <InputText fluid v-model="kurikulum.namakurikulum" />
                        <!-- <small v-if="submitted && !selectedMapel" class="text-red-500">Subject is required.</small> -->
                    </div>
                    <div>
                        <label for="mata-pelajaran">kurikulum</label>
                        <InputNumber fluid :min="1" :max="2" v-model="kurikulum.kurikulum" />
                        <!-- <small v-if="submitted && !selectedkurikulum" class="text-red-500">Teacher is required.</small> -->
                    </div>
                    <div>
                        <label for="tgl-mulai">Tanggal Mulai</label>
                        <DatePicker v-model="kurikulum.tanggalMulai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !kurikulum.tanggalMulai" />
                        <small v-if="submitted && !kurikulum.tanggalMulai" class="text-red-500">Tanggal mulai harus diisi.</small>
                    </div>
                    <div>
                        <label for="tgl-mulai">Tanggal Selesai</label>
                        <DatePicker v-model="kurikulum.tanggalSelesai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !kurikulum.tanggalSelesai" />
                        <small v-if="submitted && !kurikulum.tanggalSelesai" class="text-red-500">Tanggal selesai harus diisi.</small>
                    </div>
                </div>
                <div v-else class="grid gap-y-2">
                    <div class="flex space-x-2">
                        <div class="w-full">
                            kurikulum
                            <Select :options="kurikulumOptions" optionLabel="label" optionValue="value" fluid :disabled="isTahunAjaranAvailable" v-model:modelValue="kurikulum.kurikulum" />
                        </div>
                        <div class="w-full">
                            Nama
                            <InputText fluid :disabled="isTahunAjaranAvailable" v-model:modelValue="kurikulum.namakurikulum" />
                        </div>
                    </div>
                    <div class="flex space-x-2">
                        <div>
                            <label for="tgl-mulai">Tanggal Mulai</label>
                            <DatePicker v-model="kurikulum.tanggalMulai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !kurikulum.tanggalMulai" :disabled="isTahunAjaranAvailable" />
                            <small v-if="submitted && !kurikulum.tanggalMulai" class="text-red-500">Tanggal mulai harus diisi.</small>
                        </div>
                        <div>
                            <label for="tgl-mulai">Tanggal Selesai</label>
                            <DatePicker v-model="kurikulum.tanggalSelesai" showIcon iconDisplay="input" class="block" fluid dateFormat="dd-mm-yy" :invalid="submitted && !kurikulum.tanggalSelesai" :disabled="isTahunAjaranAvailable" />
                            <small v-if="submitted && !kurikulum.tanggalSelesai" class="text-red-500">Tanggal selesai harus diisi.</small>
                        </div>
                    </div>
                </div>
            </div>
            <template #footer>
                <Button v-if="isEdit == true" label="Update" icon="pi pi-save" @click="handleUpdatekurikulum" />
                <Button v-else label="Tambah" icon="pi pi-save" @click="handleCreatekurikulum" />
            </template>
        </Dialog>

        <!-- <DialogConfirmDelete message="data ini akan dihapus?" v-model:visible="isDeleteDialog" @confirm="handleDeletekurikulum" @closeDialog="handleCloseDialog" judul="Hapus kurikulum" /> -->
    </div>
</template>
