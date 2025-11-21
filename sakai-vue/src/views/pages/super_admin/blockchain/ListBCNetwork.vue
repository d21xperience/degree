<script setup>
import NetworkTypeComponent from '@/components/scComponent/NetworkTypeComponent.vue';
import { useSCService } from '@/composables/useSCService';
import { FilterMatchMode } from '@primevue/core/api';
import { Select } from 'primevue';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, reactive, ref, watch } from 'vue';
// import AddBCNetwork from './AddBCNetwork.vue';

const toast = useToast();
const scService = useSCService();
const BCNetworks = ref([]);
const selectedBCNetwork = ref([]);
const editingItem = reactive({});
const BCNetworkDialog = ref(false);
// const activeBCNetworkDialog = ref(false);
const deleteDialog = ref(false);
const BCPlatformSelected = computed(() => scService.getNetowrkPlatform());
watch(BCPlatformSelected, () => {
    // console.log(newVal);
    filters.value['architecture'];
});
onMounted(async () => {
    BCNetworks.value = await scService.fetchBCNetworks();
});
const headerTitle = ref('');
const isEdit = ref(false);
const editBCNetwork = () => {
    isEdit.value = true;
    headerTitle.value = 'Edit Jaringan';
    Object.assign(editingItem, selectedBCNetwork.value[0]);
    BCNetworkDialog.value = true;
};

const hideDialog = () => {
    BCNetworkDialog.value = false;
    resetData();
};
const resetData = () => {
    Object.keys(editingItem).forEach((key) => delete editingItem[key]);
    selectedBCNetwork.value = [];
};
const addBCNetwork = () => {
    headerTitle.value = 'Tambah Jaringan';
    resetData();
    BCNetworkDialog.value = true;
};

const handleAddBCNetwork = () => {};
const handleEditBCNetwork = () => {
    scService.updateBCNetwork(editingItem);
    hideDialog();
};
// const saveNetwork = async () => {
//     try {
//         Object.assign(selectedBCNetwork.value[0], editingItem);
//         toast.add({ severity: 'success', summary: 'Saved', detail: 'Jaringan diperbarui', life: 3000 });
//         hideDialog();
//         selectedBCNetwork.value = null;
//         return;
//         await store.dispatch('scService/updateNetwork', editingItem);
//         await fetchData(); // refresh
//     } catch {
//         toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal menyimpan', life: 3000 });
//     }
// };

const confirmDeleteSelected = () => {
    deleteDialog.value = true;
};

const handleDeleteBCNetwork = async () => {
    try {
        if (selectedBCNetwork.value.length == 1) {
            BCNetworks.value = BCNetworks.value.filter((n) => n.Id != selectedBCNetwork.value[0].Id);
            scService.deleteBCNetwork([selectedBCNetwork.value[0]]);
            // console.log("handle", )
        } else {
            // console.log('cek');
            BCNetworks.value = BCNetworks.value.filter((n, index) => selectedBCNetwork.value[index].Id != n.Id);
            scService.deleteBCNetwork(selectedBCNetwork.value);
        }
        // const ids = selectedBCNetwork.value.map((n) => console.log(n));
        // console.log(ids)
        // await store.dispatch('scService/deleteNetworks', ids);
        // await fetchData();
        selectedBCNetwork.value = [];
        toast.add({ severity: 'success', summary: 'Deleted', detail: 'Jaringan dihapus', life: 3000 });
    } catch {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal menghapus', life: 3000 });
    } finally {
        deleteDialog.value = false;
    }
};

const testNetwork = async () => {
    const rpcUrl = editingItem.RPCURL;
    const res = await scService.getNetworkChainId(rpcUrl);
    editingItem.ChainId = res?.ChainId;
};

const filters = ref({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    architecture: { value: null, matchMode: FilterMatchMode.EQUALS }
});
// Aktifkan network
// const activingBcNetwork = async () => {
//     const id = selectedBCNetwork.value[0].network_id;
//     try {
//         await store.dispatch('scService/activateNetwork', id);
//         await fetchData();
//         toast.add({ severity: 'success', summary: 'Activated', detail: 'Jaringan diaktifkan', life: 3000 });
//     } catch {
//         toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal aktivasi', life: 3000 });
//     } finally {
//         activeBCNetworkDialog.value = false;
//         selectedBCNetwork.value = null;
//     }
// };

const architecture = ref([
    {
        name: 'EVM'
    },
    {
        name: 'Non EVM'
    }
]);
</script>

<template>
    <div class="">
        <!-- <div class="flex flex-wrap justify-between my-2">
            <h4 class="font-bold text-2xl lg:text-lg my-2">Daftar Jaringan Blockhain </h4>
        </div> -->
        <Toolbar class="mb-6">
            <template #start>
                <Button icon="pi pi-plus" severity="success" class="mr-2" :disabled="selectedBCNetwork.length" @click="addBCNetwork" />
                <Button icon="pi pi-pencil" severity="warn" :disabled="!selectedBCNetwork || !selectedBCNetwork.length || selectedBCNetwork.length >= 2" class="mr-2" @click="editBCNetwork" />
                <Button icon="pi pi-trash" severity="danger" class="mr-2" :disabled="!selectedBCNetwork || !selectedBCNetwork.length" @click="confirmDeleteSelected" />
            </template>

            <template #end>
                <IconField>
                    <InputIcon>
                        <i class="pi pi-search"></i>
                    </InputIcon>
                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                </IconField>
            </template>
        </Toolbar>

        <DataTable
            ref="dt"
            v-model:selection="selectedBCNetwork"
            :value="BCNetworks"
            data-key="Type"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rows-per-page-options="[5, 10, 25]"
            current-page-report-template="Showing {first} to {last} of {totalRecords} BCNetworks"
        >
            <Column selection-mode="multiple" style="width: 3rem" :exportable="false" />
            <Column field="Name" header="Nama Jaringan" body-style="width:20rem" />
            <Column field="Architecture" header="Arsitektur" />
            <Column field="Type" header="Tipe Jaringan" sortable />
            <Column field="ChainId" header="Chain ID" />
            <Column field="RPCURL" header="RPC URL" />
            <Column field="ExplorerURL" header="Explorer URL" />
        </DataTable>
    </div>

    <!--  -------------------------------------------- -->
    <!--  -------------------------------------------- -->
    <!-- EDIT BCNETWORK DIALOG -->
    <Dialog v-model:visible="BCNetworkDialog" :header="headerTitle" :modal="true" :style="{ width: '450px' }">
        <div class="flex flex-col gap-4">
            <label class="font-bold">Arsitektur</label>
            <Select :options="architecture" option-label="name" placeholder="Pilih EVM atau Non EVM" />

            <label class="font-bold">Tipe Jaringan</label>
            <!-- <InputText v-model.trim="editingItem.Type" required /> -->
            <NetworkTypeComponent v-model:modelValue="editingItem.Type" :initial-value="editingItem.Type" />

            <label class="font-bold">Nama Jaringan</label>
            <InputText v-model.trim="editingItem.Name" required />

            <label class="font-bold">RPC URL</label>
            <div class="flex justify-between space-x-2">
                <div class="w-full">
                    <InputText v-model.trim="editingItem.RPCURL" required fluid />
                </div>
                <Button label="Test" severity="secondary" @click="testNetwork" />
            </div>

            <label class="font-bold">Chain ID</label>
            <InputText v-model.trim="editingItem.ChainId" required :disabled="true" />

            <label class="font-bold">Explorer URL</label>
            <InputText v-model.trim="editingItem.ExplorerURL" />
        </div>

        <template #footer>
            <Button label="Cancel" text @click="hideDialog" />
            <Button v-if="isEdit" icon="pi pi-save" label="Update" @click="handleEditBCNetwork" />
            <Button v-else icon="pi pi-save" label="Save" @click="handleAddBCNetwork" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl"></i>
            <span v-if="selectedBCNetwork.length == 1"
                >Apakah jaringan <span class="font-medium">{{ selectedBCNetwork[0].Name }} </span> akan dihapus?</span
            >
            <span v-else> Apakah jaringan yang dipilih akan dihapus? </span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="deleteDialog = false" />
            <Button label="Ya" icon="pi pi-check" text @click="handleDeleteBCNetwork" />
        </template>
    </Dialog>
    <!-- END OF DIALOGBOX FOR EDIT DATA -->
    <!--  -------------------------------------------- -->
</template>
