<script setup>
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import IssueDegreeButton from '@/components/IssueDegreeButton.vue';
import { useDns } from '@/composables/sekolah_composable/useDns';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useUtils } from '@/composables/useUtils';
import router from '@/router';
import { FilterMatchMode } from '@primevue/core/api';
import { computed, onMounted, ref, watch } from 'vue';
const { formatterDateID } = useUtils();

const { initSelectedTahunAjaran } = useSemester();
const { getDns, deleteDns } = useDns();
const { fetchSekolah } = useSekolah();
const visible = ref(false);
const selectedSiswa = ref();
const siswa = ref([]);
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    'kelas.nmKelas': { value: null, matchMode: FilterMatchMode.CONTAINS }
});
onMounted(async () => {
    siswa.value = await getDns(initSelectedTahunAjaran.value?.tahunAjaranId);
    // console.log(siswa.value)
    namaKelas.value = [...new Set(siswa.value.map((item) => item.kelas?.nmKelas).filter((nm) => nm))].map((nm) => ({
        nama: nm,
        value: nm.toLowerCase()
    }));
    // namaKelas.value = getNmKelas(siswa)
});
const tahunAjaranId = computed(() => initSelectedTahunAjaran.value.tahunAjaranId);
watch(initSelectedTahunAjaran, async (e) => {
    // console.log(`${e.tahunAjaranId}`)
    siswa.value = await getDns(`${e.tahunAjaranId}`);
    namaKelas.value = getNmKelas(siswa);
});

// ==================================
// const scData = ref({
//     degreeData: null,
//     sekolah: null,
//     ipfsUrl: null,
//     transcript: null
// });
// const selectedJurusan = ref()
// Dummy data (bisa kamu ambil dari API atau input form)
// const degreeData = ref({
//     nama: '',
//     nisn: '',
//     nik: '3211142109820004',
//     tahun_lulus: 2024,
//     major: 'Rekayasa Perangkat Lunak'
// })
const sekolah = ref('SMK PASUNDAN JATINANGOR');

const ipfsUrl = ref('https://ipfs.io/ipfs/Qm...examplehash');
const transcript = ref({
    subjects: ['Matematika', 'Pemrograman', 'Basis Data'],
    grades: [85, 90, 88]
});
const contract = null;

// watch(selectedSiswa, (newVal) => {
//     // if (newVal.length === 1) {
//     //     console.log(newVal[0].pesertaDidik.nmSiswa)
//     //     degreeData.value.nama = newVal[0].pesertaDidik.nmSiswa
//     //     degreeData.value.nisn = newVal[0].pesertaDidik.nisn
//     //     degreeData.value.tahun_lulus = 2023
//     // }
// });
const namaKelas = ref();

const getNmKelas = (data) => {
    return [...new Set(data.value.map((item) => item.kelas?.nmKelas).filter((nm) => nm))].map((nm) => ({
        nama: nm,
        value: nm.toLowerCase()
    }));
};

const editIjazah = async () => {
    const sekolah = await fetchSekolah();
    const nmSekolah = sekolah.sekolah?.nama.toLowerCase().replace(/\s+/g, '');
    // console.log(selectedSiswa.value)
    router.push({
        name: 'editIjazah',
        query: {
            pesertaDidikId: selectedSiswa.value[0]?.pesertaDidikId.toString()
        },
        params: {
            sekolah: nmSekolah
        }
    });
};
const confirmDeleteSelected = () => {
    visible.value = true;
};
const deleteData = () => {
    siswa.value = siswa.value.filter((val) => !selectedSiswa.value.includes(val));
    if (selectedSiswa.value.length == 1) {
        deleteDns(selectedSiswa.value[0].pesertaDidikId);
        // } else if (selectedSiswa.value.length > 1) {
        //     const ids = selectedSiswa.value.map((item) => item.anggotaRombelId)
        //     deleteBatchSiswaAktif(ids)
    }
};
const closeDialog = () => {
    selectedSiswa.value = null;
};
const onSubmitIjazah = () => {
    deleteData();
};
// ==================================
</script>

<template>
    <div class="">
        <div class=" ">
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <!-- <Button icon="pi pi-plus" severity="success" class="mr-2" @click="visible = true"
                            v-tooltip.bottom="'Tambah data'" /> -->
                        <Button v-tooltip.bottom="'Edit data'" icon="pi pi-pencil" severity="warn" :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 1" class="mr-2" @click="editIjazah" />
                        <Button v-tooltip.bottom="'Delete data'" icon="pi pi-trash" severity="danger" class="mr-2" :disabled="!selectedSiswa || !selectedSiswa.length" @click="confirmDeleteSelected" />
                        <!-- <Button icon="pi pi-download" severity="help" @click="exportCSV($event)" class="mr-2" v-tooltip.bottom="'Download data'" /> -->
                        <IssueDegreeButton
                            :degree-data="selectedSiswa"
                            :sekolah="sekolah"
                            :ipfs-url="ipfsUrl"
                            :transcript="transcript"
                            :tahun-ajaran-id="`${tahunAjaranId}`"
                            :contract="contract"
                            class="bg-blue-600 p-3 rounded-lg text-white"
                            :disabled="!selectedSiswa"
                            :class="{ 'bg-slate-500': !selectedSiswa || selectedSiswa.length === 0 || selectedSiswa.length > 2 }"
                            @submit="onSubmitIjazah"
                        />
                    </template>
                    <template #end>
                        <Select v-model="filters['kelas.nmKelas'].value" :options="namaKelas" option-label="nama" option-value="value" placeholder="Kelas" class="w-full md:w-48 md:mr-2" checkmark show-clear />

                        <IconField>
                            <InputIcon>
                                <i class="pi pi-search"></i>
                            </InputIcon>
                            <InputText id="search" v-model="filters['global'].value" placeholder="Search..." name="search" />
                        </IconField>
                    </template>
                </Toolbar>
            </div>
        </div>
        <DataTable
            ref="dt"
            v-model:selection="selectedSiswa"
            striped-rows
            size="small"
            :value="siswa"
            scrollable
            scroll-height="29rem"
            data-key="pesertaDidikId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rows-per-page-options="[10, 20, 50]"
            current-page-report-template="Showing {first} to {last} of {totalRecords} siswa"
        >
            <Column selection-mode="multiple" style="width: 3rem" :exportable="false" />
            <Column field="kelas.nmKelas" header="Kelas" style="width: 5rem" />
            <Column field="nisn" header="NISN" />
            <Column field="nama" header="Nama" sortable />
            <Column field="nomorIjazah" header="No Ijazah" />
            <Column field="tempatLahir" header="Tpt. Lahir" />
            <Column field="" header="Tgl. Lahir">
                <template #body="slotProps">
                    {{ formatterDateID(slotProps.data.tanggalLahir) }}
                </template>
            </Column>
            <Column field="namaOrtuWali" header="Nama Wali" />
        </DataTable>
        <!-- <Dialog v-model:visible="visible" modal header="Data ijazah" :style="{ width: '60rem', height: '100rem' }">
            <DialogIjazah :peserta-didik="selectedSiswa" :visible="visible" />
        </Dialog>  -->
        <DialogConfirmDelete v-model:visible="visible" message="Apakah data ini akan dihapus?" @confirm="deleteData" @close-dialog="closeDialog" />
    </div>
</template>
