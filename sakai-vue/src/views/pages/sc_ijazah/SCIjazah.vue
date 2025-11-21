<script setup>
import { useSekolahService } from '@/composables/sekolah_composable/useSekolah';
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { FilterMatchMode } from '@primevue/core/api';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import Toolbar from 'primevue/toolbar';
import { computed, onMounted, ref, watch } from 'vue';
const utils = useUtils();
const visible = ref(false);
const tingkatPendidikanOptions = ref();
const selectedSiswa = ref();
const siswa = ref([]);
const bentukPendidikan = ref('smk');
const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    'kelas.nmKelas': { value: null, matchMode: FilterMatchMode.CONTAINS }
});
// const { getDns, deleteDns, initSelectedTahunAjaran, fetchSekolah } = useSekolahService();

const scService = useSCService();
const sekolahService = useSekolahService();

const sekolah = computed(() => sekolahService.sekolah.value);
const tahunAjaranId = computed(() => sekolahService.initSelectedTahunAjaran.value.tahunAjaranId);

// ==================================
onMounted(async () => {
    // console.log(tahunAjaranId.value);
    // console.log(sekolah.value);
    // console.log(siswa.value);
    initialFirst();
    // namaKelas.value = [...new Set(siswa.value.map((item) => item.kelas?.nmKelas).filter((nm) => nm))].map((nm) => ({
    //     nama: nm,
    //     value: nm.toLowerCase()
    // }));
    // namaKelas.value = getNmKelas(siswa);
});

watch(tahunAjaranId, async (e) => {
    initialFirst();
    // console.log(`${e.tahunAjaranId}`);
    // siswa.value = await getDns(`${e.tahunAjaranId}`);
    // namaKelas.value = getNmKelas(siswa);
});

// ==================================
const scData = ref({
    degreeData: null,
    sekolah: null,
    ipfsUrl: null,
    transcript: null
});
// const selectedJurusan = ref();
import DialogConfirmDelete from '@/components/DialogConfirmDelete.vue';
import router from '@/router';
// Dummy data (bisa kamu ambil dari API atau input form)
// const degreeData = ref({
//     nama: '',
//     nisn: '',
//     nik: '3211142109820004',
//     tahun_lulus: 2024,
//     major: 'Rekayasa Perangkat Lunak'
// });

const ipfsUrl = ref('https://ipfs.io/ipfs/Qm...examplehash');
const transcript = ref({
    subjects: ['Matematika', 'Pemrograman', 'Basis Data'],
    grades: [85, 90, 88]
});
const contract = null;

watch(selectedSiswa, (newVal) => {
    // if (newVal.length === 1) {
    //     console.log(newVal[0].pesertaDidik.nmSiswa);
    //     degreeData.value.nama = newVal[0].pesertaDidik.nmSiswa;
    //     degreeData.value.nisn = newVal[0].pesertaDidik.nisn;
    //     degreeData.value.tahun_lulus = 2023;
    // }
});
const namaKelas = ref();

const getNmKelas = (data) => {
    return [...new Set(data.value.map((item) => item.kelas?.nmKelas).filter((nm) => nm))].map((nm) => ({
        nama: nm,
        value: nm.toLowerCase()
    }));
};

const editIjazah = async () => {
    // const sekolah = await fetchSekolah();
    const nmSekolah = sekolah.value.sekolah?.nama.toLowerCase().replace(/\s+/g, '');
    // console.log(selectedSiswa.value);
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
        //     const ids = selectedSiswa.value.map((item) => item.anggotaRombelId);
        //     deleteBatchSiswaAktif(ids);
    }
};
const closeDialog = () => {
    selectedSiswa.value = null;
};
const onSubmitIjazah = () => {
    deleteData();
};

const initialFirst = async () => {
    siswa.value = await scService.getSCIjazah({ tahun_ajaran_id: tahunAjaranId.value, sekolah_id: sekolah.value.sekolah.sekolah_id });
};

const handleLinkBlockscout = (e) => {
    const url = `http://localhost:26000/tx/${e.txHash}/internal-transactions`;
    window.open(url, '_blank'); // Buka di tab baru
};
</script>

<template>
    <div class="">
        <div class=" ">
            <div class="mb-2">
                <Toolbar>
                    <template #end>
                        <!-- <Select v-model="filters['kelas.nmKelas'].value" :options="namaKelas" optionLabel="nama" optionValue="value" placeholder="Kelas" class="w-full md:w-48 md:mr-2" checkmark show-clear /> -->
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
            <!-- <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column> -->
            <!-- <Column field="kelas.nmKelas" header="Kelas" style="width: 5rem"></Column> -->
            <Column field="ijazah.nisn" header="NISN" />
            <Column field="ijazah.nama" header="Nama" sortable />
            <Column field="ijazah.nomorIjazah" header="No Ijazah" />
            <Column field="" header="Txt Hash">
                <template #body="slotProps">
                    {{ utils.ringkasHash(slotProps.data.txHash) }}
                </template>
            </Column>
            <Column field="" header="Ijazah Hash">
                <template #body="slotProps">
                    {{ utils.ringkasHash(slotProps.data.degreeHash) }}
                </template>
            </Column>
            <Column field="" header="Tgl Trx">
                <template #body="slotProps">
                    {{ slotProps.data.tglTransaksi }}
                </template>
            </Column>
            <Column field="namaOrtuWali" header="Aksi">
                <template #body="slotProps">
                    <Button v-tooltip.bottom="'block explorer'" icon="pi pi-ethereum" class="mr-2 !text-sm" severity="danger" size="small" rounded @click="handleLinkBlockscout(slotProps.data)" />
                    <!-- <Button icon="pi pi-pencil" class="mr-2 !text-sm" severity="warn" @click="dialogEditKelas(slotProps.data)" size="small" rounded v-tooltip.bottom="'Edit kelas'" /> -->
                </template>
            </Column>
            <!-- <Column field="cidUrl" header="CID Ijazah"></Column> -->
            <!-- <Column field="nomorIjazah" header="No. Ijazah"></Column> -->
            <!-- <Column field="nis" header="NIS"></Column> -->
            <!-- <Column field="" header="Status">
                <template #body="slotProps">
                    <span :class="{ 'text-red-600': !slotProps.data.isComplete }">{{ slotProps.data.isComplete ? '✔' : 'X' }}</span>
                </template>
            </Column> -->
        </DataTable>
        <!-- <Dialog v-model:visible="visible" modal header="Data ijazah" :style="{ width: '60rem', height: '100rem' }">
            <DialogIjazah :peserta-didik="selectedSiswa" :visible="visible" />
        </Dialog>  -->
        <DialogConfirmDelete v-model:visible="visible" message="Apakah data ini akan dihapus?" @confirm="deleteData" @close-dialog="closeDialog" />
        <!-- <DialogImport :visible="dialogImport" /> -->
    </div>
</template>
