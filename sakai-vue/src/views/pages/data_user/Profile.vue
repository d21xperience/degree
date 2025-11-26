<template>
    <div class="profile-container p-4">
        <Card class="max-w-4xl mx-auto">
            <template #title>
                <div class="flex items-center gap-3">
                    <i class="pi pi-user text-2xl"></i>
                    <span>Profile Pengguna</span>
                </div>
            </template>

            <template #content>
                <!-- Header Profile -->
                <div class="flex flex-col md:flex-row items-center gap-6 mb-8">
                    <div class="relative">
                        <Avatar :image="userData.avatar" :label="userData.avatar ? '' : userData.name.charAt(0)" size="xlarge" class="bg-primary-500 text-white text-2xl" shape="circle" />
                        <Button icon="pi pi-camera" class="absolute bottom-0 right-0 w-8 h-8" severity="secondary" @click="changeAvatar" />
                    </div>

                    <div class="text-center md:text-left">
                        <h2 class="text-2xl font-bold mb-2">{{ userData.name }}</h2>
                        <p class="text-gray-600 mb-2">{{ userData.email }}</p>
                        <Chip :label="userData.role" class="bg-primary-100 text-primary-800" />
                    </div>
                </div>

                <!-- Form Edit Profile -->
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <!-- Kolom Kiri -->
                    <div class="space-y-4">
                        <div class="field">
                            <label for="name" class="font-semibold block mb-2">Nama Lengkap</label>
                            <InputText id="name" v-model="userData.name" class="w-full" placeholder="Masukkan nama lengkap" />
                        </div>

                        <div class="field">
                            <label for="email" class="font-semibold block mb-2">Email</label>
                            <InputText id="email" v-model="userData.email" class="w-full" placeholder="Masukkan email" type="email" />
                        </div>

                        <div class="field">
                            <label for="phone" class="font-semibold block mb-2">Nomor Telepon</label>
                            <InputText id="phone" v-model="userData.phone" class="w-full" placeholder="Masukkan nomor telepon" />
                        </div>
                    </div>

                    <!-- Kolom Kanan -->
                    <div class="space-y-4">
                        <div class="field">
                            <label for="department" class="font-semibold block mb-2">Departemen</label>
                            <Select id="department" v-model="userData.department" :options="departments" option-label="name" placeholder="Pilih departemen" class="w-full" />
                        </div>

                        <div class="field">
                            <label for="position" class="font-semibold block mb-2">Posisi</label>
                            <InputText id="position" v-model="userData.position" class="w-full" placeholder="Masukkan posisi" />
                        </div>

                        <div class="field">
                            <label for="joinDate" class="font-semibold block mb-2">Tanggal Bergabung</label>
                            <DatePicker id="joinDate" v-model="userData.joinDate" class="w-full" date-format="dd/mm/yy" show-icon />
                            <!-- show-icon /> -->
                        </div>
                    </div>

                    <!-- Bio Section -->
                    <div class="mt-6">
                        <label for="bio" class="font-semibold block mb-2">Bio</label>
                        <Textarea id="bio" v-model="userData.bio" rows="3" class="w-full" placeholder="Ceritakan sedikit tentang diri Anda..." />
                    </div>
                </div>
                <!-- Action Buttons -->
                <div class="flex gap-3 justify-end mt-6 pt-6 border-t">
                    <Button label="Batal" severity="secondary" outlined @click="resetForm" />
                    <Button label="Simpan Perubahan" icon="pi pi-check" :loading="loading" @click="saveProfile" />
                </div>
            </template>
        </Card>

        <!-- Additional Info Cards -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6 max-w-4xl mx-auto">
            <Card>
                <template #title>
                    <div class="flex items-center gap-2">
                        <i class="pi pi-chart-line text-primary-500"></i>
                        <span>Statistik</span>
                    </div>
                </template>
                <template #content>
                    <div class="space-y-3">
                        <div class="flex justify-between">
                            <span>Projects</span>
                            <span class="font-semibold">12</span>
                        </div>
                        <div class="flex justify-between">
                            <span>Tasks</span>
                            <span class="font-semibold">47</span>
                        </div>
                        <div class="flex justify-between">
                            <span>Completed</span>
                            <span class="font-semibold">38</span>
                        </div>
                    </div>
                </template>
            </Card>

            <Card>
                <template #title>
                    <div class="flex items-center gap-2">
                        <i class="pi pi-cog text-primary-500"></i>
                        <span>Pengaturan</span>
                    </div>
                </template>
                <template #content>
                    <div class="space-y-2">
                        <Button label="Ubah Password" icon="pi pi-key" text class="w-full justify-start" />
                        <Button label="Preferensi Notifikasi" icon="pi pi-bell" text class="w-full justify-start" />
                        <Button label="Privasi & Keamanan" icon="pi pi-shield" text class="w-full justify-start" />
                    </div>
                </template>
            </Card>

            <Card>
                <template #title>
                    <div class="flex items-center gap-2">
                        <i class="pi pi-info-circle text-primary-500"></i>
                        <span>Informasi</span>
                    </div>
                </template>
                <template #content>
                    <div class="space-y-2 text-sm">
                        <div class="flex justify-between">
                            <span class="text-gray-600">Status</span>
                            <Tag value="Aktif" severity="success" />
                        </div>
                        <div class="flex justify-between">
                            <span class="text-gray-600">Terakhir Login</span>
                            <span>2 jam yang lalu</span>
                        </div>
                        <div class="flex justify-between">
                            <span class="text-gray-600">Member sejak</span>
                            <span>Jan 2023</span>
                        </div>
                    </div>
                </template>
            </Card>
        </div>
    </div>
</template>

<script setup>
import { useToast } from 'primevue/usetoast';
import { reactive, ref } from 'vue';

// Components PrimeVue
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Chip from 'primevue/chip';
import InputText from 'primevue/inputtext';
import Select from 'primevue/Select';
import Tag from 'primevue/tag';
import Textarea from 'primevue/textarea';

const toast = useToast();
const loading = ref(false);

// Data user
const userData = reactive({
    name: 'John Doe',
    email: 'john.doe@example.com',
    phone: '+62 812-3456-7890',
    department: null,
    position: 'Senior Developer',
    joinDate: new Date('2023-01-15'),
    bio: 'Full-stack developer dengan pengalaman 5+ tahun dalam pengembangan web application.',
    role: 'Administrator',
    avatar: null
});

// Options untuk Select
const departments = ref([
    { name: 'IT & Development', code: 'IT' },
    { name: 'Human Resources', code: 'HR' },
    { name: 'Marketing', code: 'MKT' },
    { name: 'Finance', code: 'FIN' },
    { name: 'Operations', code: 'OPS' }
]);

// Set default department
userData.department = departments.value[0];

// Methods
const changeAvatar = () => {
    toast.add({
        severity: 'info',
        summary: 'Info',
        detail: 'Fitur upload avatar akan segera tersedia',
        life: 3000
    });
};

const saveProfile = async () => {
    loading.value = true;

    // Simulate API call
    try {
        await new Promise((resolve) => setTimeout(resolve, 1000));

        toast.add({
            severity: 'success',
            summary: 'Berhasil',
            detail: 'Profile berhasil diperbarui',
            life: 3000
        });
    } catch (error) {
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Gagal memperbarui profile',
            life: 3000
        });
    } finally {
        loading.value = false;
    }
};

const resetForm = () => {
    // Reset logic here
    toast.add({
        severity: 'info',
        summary: 'Info',
        detail: 'Form telah direset',
        life: 3000
    });
};
</script>

<style scoped>
.profile-container {
    min-height: 100vh;
    /* background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); */
}

:deep(.p-card) {
    box-shadow:
        0 10px 25px -5px rgba(0, 0, 0, 0.1),
        0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

:deep(.p-avatar) {
    border: 4px solid white;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}
</style>
