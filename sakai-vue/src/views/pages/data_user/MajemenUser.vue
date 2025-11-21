<template>
    <div class="user-management p-4">
        <Card>
            <template #title>
                <div class="flex justify-between items-center">
                    <div class="flex items-center gap-3">
                        <i class="pi pi-users text-2xl"></i>
                        <span>Manajemen Pengguna</span>
                        <Badge :value="totalUsers" class="ml-2" />
                    </div>
                    <Button label="Tambah Pengguna" icon="pi pi-plus" severity="success" @click="openNew" />
                </div>
            </template>

            <template #content>
                <!-- Toolbar -->
                <div class="flex flex-col md:flex-row gap-4 mb-6">
                    <div class="flex-1">
                        <span class="p-input-icon-left">
                            <i class="pi pi-search"></i>
                            <InputText v-model="filters.global" placeholder="Cari pengguna..." class="w-full" />
                        </span>
                    </div>

                    <div class="flex gap-2">
                        <Dropdown v-model="filters.role" :options="roleOptions" option-label="label" placeholder="Filter Role" class="w-full md:w-40" show-clear />
                        <Dropdown v-model="filters.status" :options="statusOptions" option-label="label" placeholder="Filter Status" class="w-full md:w-40" show-clear />
                        <Button v-tooltip="'Hapus filter'" icon="pi pi-filter-slash" severity="secondary" @click="clearFilters" />
                    </div>
                </div>

                <!-- Data Table -->
                <DataTable
                    :value="filteredUsers"
                    :loading="loading"
                    :paginator="true"
                    :rows="10"
                    :rows-per-page-options="[5, 10, 20, 50]"
                    paginator-template="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    current-page-report-template="Menampilkan {first} sampai {last} dari {totalRecords} pengguna"
                    striped-rows
                    removable-sort
                    class="p-datatable-sm"
                >
                    <Column field="id" header="ID" sortable style="min-width: 80px" />

                    <Column field="name" header="Nama" sortable style="min-width: 200px">
                        <template #body="slotProps">
                            <div class="flex items-center gap-3">
                                <Avatar :image="slotProps.data.avatar" :label="slotProps.data.avatar ? '' : slotProps.data.name.charAt(0)" size="normal" shape="circle" class="bg-primary-500 text-white" />
                                <div>
                                    <div class="font-semibold">{{ slotProps.data.name }}</div>
                                    <div class="text-sm text-gray-600">{{ slotProps.data.email }}</div>
                                </div>
                            </div>
                        </template>
                    </Column>

                    <Column field="role" header="Role" sortable style="min-width: 120px">
                        <template #body="slotProps">
                            <Tag :value="slotProps.data.role" :severity="getRoleSeverity(slotProps.data.role)" />
                        </template>
                    </Column>

                    <Column field="department" header="Departemen" sortable style="min-width: 150px" />

                    <Column field="status" header="Status" sortable style="min-width: 120px">
                        <template #body="slotProps">
                            <Tag :value="slotProps.data.status" :severity="getStatusSeverity(slotProps.data.status)" />
                        </template>
                    </Column>

                    <Column field="lastLogin" header="Login Terakhir" sortable style="min-width: 150px">
                        <template #body="slotProps">
                            {{ formatDate(slotProps.data.lastLogin) }}
                        </template>
                    </Column>

                    <Column field="createdAt" header="Tanggal Dibuat" sortable style="min-width: 150px">
                        <template #body="slotProps">
                            {{ formatDate(slotProps.data.createdAt) }}
                        </template>
                    </Column>

                    <Column header="Aksi" style="min-width: 150px">
                        <template #body="slotProps">
                            <div class="flex gap-1">
                                <Button v-tooltip="'Lihat detail'" icon="pi pi-eye" severity="info" text rounded @click="viewUser(slotProps.data)" />
                                <Button v-tooltip="'Edit pengguna'" icon="pi pi-pencil" severity="warning" text rounded @click="editUser(slotProps.data)" />
                                <Button v-tooltip="'Hapus pengguna'" icon="pi pi-trash" severity="danger" text rounded @click="confirmDeleteUser(slotProps.data)" />
                            </div>
                        </template>
                    </Column>

                    <template #empty>
                        <div class="text-center py-6 text-gray-500">
                            <i class="pi pi-users text-4xl mb-2"></i>
                            <p>Tidak ada data pengguna</p>
                        </div>
                    </template>

                    <template #loading>
                        <div class="text-center py-6">
                            <ProgressSpinner style="width: 50px; height: 50px" />
                            <p class="mt-2">Memuat data...</p>
                        </div>
                    </template>
                </DataTable>
            </template>
        </Card>

        <!-- Dialog Add/Edit User -->
        <Dialog v-model:visible="userDialog" :style="{ width: '600px' }" :header="dialogTitle" :modal="true" class="p-fluid">
            <div class="field grid gap-3">
                <div class="col-12 md:col-6">
                    <label for="name">Nama Lengkap <span class="text-red-500">*</span></label>
                    <InputText id="name" v-model="user.name" required :class="{ 'p-invalid': submitted && !user.name }" />
                    <small v-if="submitted && !user.name" class="p-error"> Nama harus diisi </small>
                </div>

                <div class="col-12 md:col-6">
                    <label for="email">Email <span class="text-red-500">*</span></label>
                    <InputText id="email" v-model="user.email" required type="email" :class="{ 'p-invalid': submitted && !user.email }" />
                    <small v-if="submitted && !user.email" class="p-error"> Email harus diisi </small>
                </div>

                <div class="col-12 md:col-6">
                    <label for="phone">Nomor Telepon</label>
                    <InputText id="phone" v-model="user.phone" />
                </div>

                <div class="col-12 md:col-6">
                    <label for="department">Departemen</label>
                    <Dropdown id="department" v-model="user.department" :options="departments" option-label="name" placeholder="Pilih departemen" />
                </div>

                <div class="col-12 md:col-6">
                    <label for="role">Role <span class="text-red-500">*</span></label>
                    <Dropdown id="role" v-model="user.role" :options="roles" option-label="name" placeholder="Pilih role" :class="{ 'p-invalid': submitted && !user.role }" />
                    <small v-if="submitted && !user.role" class="p-error"> Role harus dipilih </small>
                </div>

                <div class="col-12 md:col-6">
                    <label for="status">Status</label>
                    <Dropdown id="status" v-model="user.status" :options="statuses" option-label="label" placeholder="Pilih status" />
                </div>

                <div class="col-12">
                    <label for="notes">Catatan</label>
                    <Textarea id="notes" v-model="user.notes" rows="3" />
                </div>
            </div>

            <template #footer>
                <Button label="Batal" icon="pi pi-times" text @click="hideDialog" />
                <Button label="Simpan" icon="pi pi-check" :loading="loading" @click="saveUser" />
            </template>
        </Dialog>

        <!-- Delete Confirmation Dialog -->
        <Dialog v-model:visible="deleteUserDialog" :style="{ width: '450px' }" header="Konfirmasi" :modal="true">
            <div class="confirmation-content">
                <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem"></i>
                <span v-if="user">
                    Apakah Anda yakin ingin menghapus <b>{{ user.name }}</b
                    >?
                </span>
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="deleteUserDialog = false" />
                <Button label="Ya" icon="pi pi-check" text severity="danger" @click="deleteUser" />
            </template>
        </Dialog>

        <!-- User Detail Dialog -->
        <Dialog v-model:visible="userDetailDialog" :style="{ width: '700px' }" header="Detail Pengguna" :modal="true">
            <div v-if="selectedUser" class="grid gap-4">
                <div class="flex items-center gap-4">
                    <Avatar :image="selectedUser.avatar" :label="selectedUser.avatar ? '' : selectedUser.name.charAt(0)" size="xlarge" shape="circle" class="bg-primary-500 text-white text-2xl" />
                    <div>
                        <h3 class="text-xl font-bold">{{ selectedUser.name }}</h3>
                        <p class="text-gray-600">{{ selectedUser.email }}</p>
                        <div class="flex gap-2 mt-2">
                            <Tag :value="selectedUser.role" :severity="getRoleSeverity(selectedUser.role)" />
                            <Tag :value="selectedUser.status" :severity="getStatusSeverity(selectedUser.status)" />
                        </div>
                    </div>
                </div>

                <Divider />

                <div class="grid grid-cols-2 gap-4">
                    <div>
                        <h4 class="font-semibold mb-2">Informasi Kontak</h4>
                        <div class="space-y-2">
                            <div><strong>Telepon:</strong> {{ selectedUser.phone || '-' }}</div>
                            <div><strong>Departemen:</strong> {{ selectedUser.department || '-' }}</div>
                        </div>
                    </div>
                    <div>
                        <h4 class="font-semibold mb-2">Informasi Akun</h4>
                        <div class="space-y-2">
                            <div><strong>Terakhir Login:</strong> {{ formatDate(selectedUser.lastLogin) }}</div>
                            <div><strong>Dibuat Pada:</strong> {{ formatDate(selectedUser.createdAt) }}</div>
                        </div>
                    </div>
                </div>

                <div v-if="selectedUser.notes">
                    <h4 class="font-semibold mb-2">Catatan</h4>
                    <p class="text-gray-700">{{ selectedUser.notes }}</p>
                </div>
            </div>
        </Dialog>

        <!-- Toast -->
        <Toast />
    </div>
</template>

<script setup>
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, reactive, ref } from 'vue';

// PrimeVue Components
import Avatar from 'primevue/avatar';
import Badge from 'primevue/badge';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import Dialog from 'primevue/dialog';
import Divider from 'primevue/divider';
import Dropdown from 'primevue/dropdown';
import InputText from 'primevue/inputtext';
import ProgressSpinner from 'primevue/progressspinner';
import Tag from 'primevue/tag';
import Textarea from 'primevue/textarea';
import Toast from 'primevue/toast';

const toast = useToast();
const loading = ref(false);
const userDialog = ref(false);
const deleteUserDialog = ref(false);
const userDetailDialog = ref(false);
const submitted = ref(false);

// Data
const users = ref([]);
const user = ref({});
const selectedUser = ref(null);

// Filters
const filters = reactive({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    role: { value: null, matchMode: FilterMatchMode.EQUALS },
    status: { value: null, matchMode: FilterMatchMode.EQUALS }
});

// Options
const roles = ref([
    { name: 'Administrator', code: 'admin' },
    { name: 'Manager', code: 'manager' },
    { name: 'User', code: 'user' },
    { name: 'Guest', code: 'guest' }
]);

const statuses = ref([
    { label: 'Aktif', value: 'active' },
    { label: 'Nonaktif', value: 'inactive' },
    { label: 'Tertunda', value: 'pending' }
]);

const departments = ref([
    { name: 'IT & Development', code: 'IT' },
    { name: 'Human Resources', code: 'HR' },
    { name: 'Marketing', code: 'MKT' },
    { name: 'Finance', code: 'FIN' },
    { name: 'Operations', code: 'OPS' }
]);

const roleOptions = ref([
    { label: 'Administrator', value: 'Administrator' },
    { label: 'Manager', value: 'Manager' },
    { label: 'User', value: 'User' },
    { label: 'Guest', value: 'Guest' }
]);

const statusOptions = ref([
    { label: 'Aktif', value: 'active' },
    { label: 'Nonaktif', value: 'inactive' },
    { label: 'Tertunda', value: 'pending' }
]);

// Computed
const filteredUsers = computed(() => {
    let result = users.value;

    if (filters.global.value) {
        const search = filters.global.value.toLowerCase();
        result = result.filter((user) => user.name.toLowerCase().includes(search) || user.email.toLowerCase().includes(search) || user.department?.toLowerCase().includes(search));
    }

    if (filters.role.value) {
        result = result.filter((user) => user.role === filters.role.value.value);
    }

    if (filters.status.value) {
        result = result.filter((user) => user.status === filters.status.value.value);
    }

    return result;
});

const totalUsers = computed(() => filteredUsers.value.length);

const dialogTitle = computed(() => {
    return user.value.id ? 'Edit Pengguna' : 'Tambah Pengguna Baru';
});

// Methods
const openNew = () => {
    user.value = {};
    submitted.value = false;
    userDialog.value = true;
};

const hideDialog = () => {
    userDialog.value = false;
    submitted.value = false;
};

const saveUser = async () => {
    submitted.value = true;

    if (!user.value.name || !user.value.email || !user.value.role) {
        return;
    }

    loading.value = true;

    try {
        // Simulate API call
        await new Promise((resolve) => setTimeout(resolve, 1000));

        if (user.value.id) {
            // Update existing user
            const index = users.value.findIndex((u) => u.id === user.value.id);
            users.value[index] = { ...user.value };
            toast.add({
                severity: 'success',
                summary: 'Berhasil',
                detail: 'Pengguna berhasil diperbarui',
                life: 3000
            });
        } else {
            // Add new user
            const newUser = {
                ...user.value,
                id: generateId(),
                createdAt: new Date(),
                lastLogin: null,
                status: user.value.status || 'active'
            };
            users.value.push(newUser);
            toast.add({
                severity: 'success',
                summary: 'Berhasil',
                detail: 'Pengguna berhasil ditambahkan',
                life: 3000
            });
        }

        userDialog.value = false;
        user.value = {};
    } catch (error) {
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Gagal menyimpan pengguna',
            life: 3000
        });
    } finally {
        loading.value = false;
    }
};

const editUser = (userData) => {
    user.value = { ...userData };
    userDialog.value = true;
};

const viewUser = (userData) => {
    selectedUser.value = { ...userData };
    userDetailDialog.value = true;
};

const confirmDeleteUser = (userData) => {
    user.value = userData;
    deleteUserDialog.value = true;
};

const deleteUser = async () => {
    loading.value = true;

    try {
        // Simulate API call
        await new Promise((resolve) => setTimeout(resolve, 500));

        users.value = users.value.filter((u) => u.id !== user.value.id);
        deleteUserDialog.value = false;
        user.value = {};

        toast.add({
            severity: 'success',
            summary: 'Berhasil',
            detail: 'Pengguna berhasil dihapus',
            life: 3000
        });
    } catch (error) {
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Gagal menghapus pengguna',
            life: 3000
        });
    } finally {
        loading.value = false;
    }
};

const clearFilters = () => {
    filters.global.value = null;
    filters.role.value = null;
    filters.status.value = null;
};

const getRoleSeverity = (role) => {
    switch (role) {
        case 'Administrator':
            return 'danger';
        case 'Manager':
            return 'warning';
        case 'User':
            return 'info';
        case 'Guest':
            return 'secondary';
        default:
            return null;
    }
};

const getStatusSeverity = (status) => {
    switch (status) {
        case 'active':
            return 'success';
        case 'inactive':
            return 'danger';
        case 'pending':
            return 'warning';
        default:
            return null;
    }
};

const formatDate = (date) => {
    if (!date) return '-';
    return new Date(date).toLocaleDateString('id-ID', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
};

const generateId = () => {
    return Math.floor(Math.random() * 1000) + 1;
};

// Mock data
const initializeUsers = () => {
    users.value = [
        {
            id: 1,
            name: 'John Doe',
            email: 'john.doe@example.com',
            phone: '+62 812-3456-7890',
            role: 'Administrator',
            department: 'IT & Development',
            status: 'active',
            lastLogin: new Date('2024-01-15 08:30:00'),
            createdAt: new Date('2023-01-15'),
            notes: 'Super administrator dengan akses penuh'
        },
        {
            id: 2,
            name: 'Jane Smith',
            email: 'jane.smith@example.com',
            phone: '+62 813-4567-8901',
            role: 'Manager',
            department: 'Marketing',
            status: 'active',
            lastLogin: new Date('2024-01-14 14:20:00'),
            createdAt: new Date('2023-03-20'),
            notes: 'Manager departemen marketing'
        },
        {
            id: 3,
            name: 'Bob Johnson',
            email: 'bob.johnson@example.com',
            phone: '+62 814-5678-9012',
            role: 'User',
            department: 'Finance',
            status: 'active',
            lastLogin: new Date('2024-01-13 09:15:00'),
            createdAt: new Date('2023-06-10')
        },
        {
            id: 4,
            name: 'Alice Brown',
            email: 'alice.brown@example.com',
            role: 'Guest',
            department: null,
            status: 'inactive',
            lastLogin: new Date('2024-01-10 16:45:00'),
            createdAt: new Date('2023-08-05')
        }
    ];
};

onMounted(() => {
    initializeUsers();
});
</script>

<style scoped>
.user-management {
    min-height: 100vh;
    background: #f8f9fa;
}

:deep(.p-card) {
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

:deep(.p-datatable) {
    border-radius: 8px;
    overflow: hidden;
}

.confirmation-content {
    display: flex;
    align-items: center;
    gap: 1rem;
}
</style>
