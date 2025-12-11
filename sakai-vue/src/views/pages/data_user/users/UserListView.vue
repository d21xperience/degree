<script setup>
import { useUserCrud } from '@/composables/user/crud';
import { useToast } from 'primevue/usetoast';
import { onMounted, ref } from 'vue';

const toast = useToast();
const { list, fetchList, loading } = useUserCrud();

// Delete dialog state
const deletingData = ref(null);
const deleteDialogVisible = ref(false);

const openDelete = (row) => {
    deletingData.value = row;
    deleteDialogVisible.value = true;
};

onMounted(() => {
    fetchList();
});
</script>

<template>
    <div class="p-4 space-y-4">
        <Toast />
        <div class="flex justify-between items-center">
            <h2 class="text-xl font-semibold">User List</h2>
            <RouterLink :to="{ name: 'user-create' }" class="px-3 py-2 rounded-lg bg-primary text-white hover:bg-primary-600 transition"> Create User </RouterLink>
        </div>

        <DataTable :value="list" :loading="loading" data-key="id" class="rounded-lg shadow-md">
            <Column field="id" header="ID" style="width: 80px" />
            <Column field="name" header="Name" />

            <Column header="Actions" style="width: 150px">
                <template #body="{ data }">
                    <div class="flex items-center gap-2">
                        <RouterLink :to="{ name: 'user-edit', params: { id: data.id } }" class="px-2 py-1 rounded bg-blue-500 text-white hover:bg-blue-600 text-xs"> Edit </RouterLink>

                        <button class="px-2 py-1 rounded bg-red-500 text-white hover:bg-red-600 text-xs" @click="openDelete(data)">Delete</button>
                    </div>
                </template>
            </Column>
        </DataTable>

        <!-- Delete Dialog -->
        <DeleteDialog v-if="deleteDialogVisible" v-model:visible="deleteDialogVisible" :data="deletingData" crud="User" @success="fetchList()" />
    </div>
</template>

<style scoped></style>
