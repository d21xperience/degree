<template>
    <div class="">
        <!-- <div class="flex flex-wrap justify-between my-2">
            <h4 class="font-bold text-2xl lg:text-lg my-2">Daftar Jaringan Blockhain </h4>
        </div> -->
        <Toolbar class="mb-6">
            <template #start>
                <Button icon="pi pi-plus" severity="success" class="mr-2" @click="router.push({ name: 'addBCNetworks' })" />
                <Button icon="pi pi-pencil" severity="warn" @click="editBCNetwork" :disabled="!selectedBCNetwork || !selectedBCNetwork.length || selectedBCNetwork.length >= 2" class="mr-2" />
                <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected" :disabled="!selectedBCNetwork || !selectedBCNetwork.length" />
                <Button icon="pi pi-upload" severity="help" @click="exportCSV($event)" class="mr-2" />
                <!-- <Button label="Aktifkan" icon="pi pi-check" severity="info" @click="activeBCNetworkDialog = true"
                    class="mr-2"
                    :disabled="!selectedBCNetwork || !selectedBCNetwork.length || selectedBCNetwork.length >= 2" /> -->
            </template>

            <template #end>
                <IconField>
                    <InputIcon>
                        <i class="pi pi-search" />
                    </InputIcon>
                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                </IconField>
            </template>
        </Toolbar>

        <DataTable
            ref="dt"
            v-model:selection="selectedBCNetwork"
            :value="BCNetworks"
            dataKey="ChainId"
            :paginator="true"
            :rows="10"
            :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[5, 10, 25]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} BCNetworks"
        >
            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="Name" header="Nama Jaringan" />
            <Column field="Type" header="Tipe Jaringan" sortable />
            <Column field="ChainId" header="Chain ID" />
            <Column field="RPCURL" header="RPC URL" />
            <Column field="ExplorerURL" header="Explorer URL" />
        </DataTable>
    </div>

    <!--  -------------------------------------------- -->
    <!--  -------------------------------------------- -->
    <!-- EDIT BCNETWORK DIALOG -->
    <Dialog v-model:visible="BCNetworkDialog" header="Edit Data" :modal="true" :style="{ width: '450px' }">
        <div class="flex flex-col gap-4">
            <label class="font-bold">Nama Jaringan</label>
            <InputText v-model.trim="editingItem.Name" required />

            <label class="font-bold">Chain ID</label>
            <InputText v-model.trim="editingItem.ChainId" required />

            <label class="font-bold">Tipe Jaringan</label>
            <!-- <InputText v-model.trim="editingItem.Type" required /> -->
            <NetworkTypeComponent v-model:modelValue="editingItem.Type" :initialValue="editingItem.Type" />

            <label class="font-bold">RPC URL</label>
            <InputText v-model.trim="editingItem.RPCURL" required />

            <label class="font-bold">Explorer URL</label>
            <InputText v-model.trim="editingItem.ExplorerURL" />
        </div>

        <template #footer>
            <Button label="Cancel" text @click="hideDialog" />
            <Button label="Save" @click="saveNetwork" />
        </template>
    </Dialog>

    <Toast />

    <Dialog v-model:visible="deleteDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="selectedBCNetwork.length == 1"
                >Apakah jaringan <span class="font-medium">{{ selectedBCNetwork[0].Name }} </span> akan dihapus?</span
            >
            <span v-else> Apakah jaringan yang dipilih akan dihapus? </span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="deleteDialog = false" />
            <Button label="Ya" icon="pi pi-check" text @click="deleteBCNetwork" />
        </template>
    </Dialog>
    <!-- END OF DIALOGBOX FOR EDIT DATA -->
    <!--  -------------------------------------------- -->
</template>

<script setup>
import NetworkTypeComponent from '@/components/scComponent/NetworkTypeComponent.vue';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import { onMounted, reactive, ref } from 'vue';
import { useStore } from 'vuex';

const store = useStore();
const toast = useToast();

const BCNetworks = ref([]);
const selectedBCNetwork = ref();
const editingItem = reactive({});
const loading = ref(false);
const BCNetworkDialog = ref(false);
const activeBCNetworkDialog = ref(false);
const deleteDialog = ref(false);
const filters = ref({ global: { value: null, matchMode: FilterMatchMode.CONTAINS } });

const fetchData = async () => {
    loading.value = true;
    try {
        const { network } = await store.dispatch('scService/fetchBlockchainNetworks');
        BCNetworks.value = network;
    } catch (e) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal memuat data', life: 3000 });
    } finally {
        loading.value = false;
    }
};

onMounted(fetchData);

const editBCNetwork = () => {
    // console.log(selectedBCNetwork.value[0]);
    console.log(selectedBCNetwork);
    Object.assign(editingItem, selectedBCNetwork.value[0]);
    BCNetworkDialog.value = true;
};

const hideDialog = () => {
    BCNetworkDialog.value = false;
};

const saveNetwork = async () => {
    try {
        Object.assign(selectedBCNetwork.value[0], editingItem);
        toast.add({ severity: 'success', summary: 'Saved', detail: 'Jaringan diperbarui', life: 3000 });
        hideDialog();
        selectedBCNetwork.value = null;
        return;
        await store.dispatch('scService/updateNetwork', editingItem);
        await fetchData(); // refresh
    } catch {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal menyimpan', life: 3000 });
    }
};

const confirmDeleteSelected = () => {
    deleteDialog.value = true;
};

const deleteBCNetwork = async () => {
    try {
        if (selectedBCNetwork.value.length == 1) {
            BCNetworks.value = BCNetworks.value.filter((n) => n.Id != selectedBCNetwork.value[0].Id);
        } else {
            console.log("cek")
            BCNetworks.value = BCNetworks.value.filter((n) => selectedBCNetwork.value.some((selected) => selected.Id === n.Id));
        }
        // const ids = selectedBCNetwork.value.map((n) => console.log(n));
        // console.log(ids)
        // await store.dispatch('scService/deleteNetworks', ids);
        // await fetchData();
        selectedBCNetwork.value = null;
        toast.add({ severity: 'success', summary: 'Deleted', detail: 'Jaringan dihapus', life: 3000 });
    } catch {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal menghapus', life: 3000 });
    } finally {
        deleteDialog.value = false;
    }
};

// Aktifkan network
const activingBcNetwork = async () => {
    const id = selectedBCNetwork.value[0].network_id;
    try {
        await store.dispatch('scService/activateNetwork', id);
        await fetchData();
        toast.add({ severity: 'success', summary: 'Activated', detail: 'Jaringan diaktifkan', life: 3000 });
    } catch {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal aktivasi', life: 3000 });
    } finally {
        activeBCNetworkDialog.value = false;
        selectedBCNetwork.value = null;
    }
};
</script>
