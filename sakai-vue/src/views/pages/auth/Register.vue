<script setup>
import router from '@/router';
import SekolahService from '@/service/SekolahService';
// import { useAuth } from '@/views/pages/auth/composables/auth';
// const {  } = useAuth();
import Dialog from 'primevue/dialog';
import { onMounted, ref } from 'vue';
import { useStore } from 'vuex';
const store = useStore();

const email = ref('');
const password = ref('');
const loading = ref(false);
const sekolah = ref();
onMounted(() => {
    SekolahService.getSekolah().then((data) => (sekolah.value = data));
});

const statusSekolahTerdaftar = ref(false);
const searchTerm = ref('');
const filteredsekolah = ref();
const search = (event) => {
    setTimeout(() => {
        if (!event.query.trim().length) {
            filteredsekolah.value = [...sekolah.value];
        } else {
            filteredsekolah.value = sekolah.value.filter((country) => country.nama_sekolah.toLowerCase().includes(event.query.toLowerCase()));
        }
    }, 250);
};
const handleKeydown = (event) => {
    if (event.key === ' ') {
        searchTerm.value += ' '; // Menambahkan spasi ke query
    }
};
// const npsn = ref();
// const error = ref();
const cekSekolah = async () => {
    loading.value = true;
    let cek = false;
    // npsn.value = searchTerm.value?.npsn;
    try {
        sekolah.value = null; // Reset data sekolah
        // Panggil fungsi ceknpsn dari Vuex storex
        const data = await store.dispatch('authService/ceknpsn', searchTerm.value.npsn);
        console.log(data);
        if (data) {
            dialogInfo.value = true;
            // searchTerm.value = null
        }
    } catch (e) {
        // console.error('error:', e.response.data.message);
        if (e.response.data.message === 'failed to retrieve school data') {
            cek = true;
        }
        // error.value = 'Terjadi kesalahan saat mengambil data sekolah.';
    } finally {
        if (cek) {
            // sekolah.value = data; //Tampilkan data sekolah
            statusSekolahTerdaftar.value = true;
        }
        loading.value = false;
    }
};
// Fungsi handler submit form
const handleSubmit = async () => {
    loading.value = true;
    let dataReg = {
        user: {
            username: '',
            email: email.value,
            password: password.value,
            role: 'admin'
        },
        sekolah: {
            nama_sekolah: searchTerm.value.nama_sekolah,
            npsn: searchTerm.value.npsn,
            enkrip_id: searchTerm.value.sekolah_id_enkrip,
            kecamatan: searchTerm.value.kecamatan,
            kabupaten: searchTerm.value.kabupaten,
            propinsi: searchTerm.value.propinsi,
            kode_kecamatan: searchTerm.value.kode_kecamatan,
            kode_kab: searchTerm.value.kode_kab,
            kode_prop: searchTerm.value.kode_prop,
            alamat_jalan: searchTerm.value.alamat_jalan
        }
    };
    // if (valid) {
    try {
        dataReg.sekolah = formatValues(dataReg.sekolah);
        const resp = await store.dispatch('authService/registerAdmin', dataReg);
        console.log(resp);
        // Jika sukses, arahkan ke beranda
        if (resp.ok) {
            const result = resp?.sekolahTenant.namaSekolah.toLowerCase().replace(/\s+/g, '');
            router.push({ name: 'dashboard', params: { sekolah: result } });
        }
        // success.value = 'Admin registered successfully!';
    } catch (error) {
        errorDialog.value = true;
        errorInfo.value = error?.message;
        console.error(error);
        // error.value = err.error || 'Registration failed';
    } finally {
        loading.value = false;
    }
    // return { name, email, password, schoolName, register, error, success };
    // }
};

// Fungsi untuk menghapus spasi dan konversi ke string
const formatValues = (obj) => {
    return Object.fromEntries(Object.entries(obj).map(([key, value]) => [key, String(value).trim()]));
};

const dialogInfo = ref(false);
const errorDialog = ref(false);
const errorInfo = ref();
</script>

<template>
    <!-- <FloatingConfigurator /> -->
    <div class="bg-surface-50 dark:bg-surface-950 flex items-center justify-center min-h-screen min-w-[100vw] overflow-hidden">
        <div style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
            <div v-if="!statusSekolahTerdaftar" class="flex flex-col items-center justify-center">
                <div class="w-full bg-surface-0 dark:bg-surface-900 py-20 px-8" style="border-radius: 53px">
                    <h3 class="text-2xl font-bold text-center">Form. Register</h3>
                    <ol class="list-decimal space-y-2">
                        <li>Formulir Register diperuntukan untuk <strong>Admin Sekolah.</strong></li>
                        <li>Admin Sekolah adalah <strong>Guru</strong> atau <strong> Tendik</strong> yang telah terdaftar di Dapodik.</li>
                        <li>Isi nama Sekolah atau NPSN pada kolom di bawah kemudian click tombol <strong>cek</strong>.</li>
                        <li>Jika sekolah <strong>belum terdaftar</strong> maka akan dilanjutkan untuk mengsisi formulir register.</li>
                    </ol>

                    <div class="my-6">
                        <div class="flex justify-between items-center space-x-8">
                            <div class="w-full">
                                <FloatLabel variant="on">
                                    <IconField>
                                        <AutoComplete optionLabel="nama_sekolah" v-model="searchTerm" :suggestions="filteredsekolah" @complete="search" fluid @keydown.space.prevent="handleKeydown" :disabled="loading" />
                                        <InputIcon class="pi pi-building-columns" />
                                    </IconField>
                                    <label>NPSN/Nama Sekolah</label>
                                </FloatLabel>
                            </div>
                            <div>
                                <Button label="Cek" class="w-24" @click="cekSekolah" :loading="loading" :disabled="searchTerm.length <= 0"></Button>
                            </div>
                        </div>
                    </div>

                    <div class="flex justify-between mt-6">
                        <div class="flex justify-center flex-col">
                            <RouterLink to="/" class="xs:text-[10px] text-blue-600 hover:underline">Ke halaman utama </RouterLink>
                        </div>
                        <div class="flex justify-center flex-col">
                            <p class="">Sudah punya akun? <RouterLink :to="{ name: 'login' }" class="text-blue-600 font-semibold hover:underline ml-1">Login disini</RouterLink></p>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else class="flex flex-col items-center justify-center">
                <!-- <h2>Register</h2> -->
                <div class="w-full bg-surface-0 dark:bg-surface-900 py-20 px-8 sm:px-20" style="border-radius: 53px">
                    <div class="text-center mb-8">
                        <!-- <div class="text-surface-900 dark:text-surface-0 text-3xl font-medium mb-4">Welcome to PrimeLand!</div> -->
                        <!-- <span class="text-muted-color font-medium">Sign in to continue</span> -->
                    </div>

                    <div>
                        <label for="sekolah" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Sekolah</label>
                        <InputText id="sekolah" name="sekolah" type="text" placeholder="Masukan sekolah" class="w-full md:w-[30rem] mb-3" v-model="filteredsekolah[0].nama_sekolah" disabled />
                        <label for="email1" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Email</label>
                        <InputText id="email1" name="email1" type="text" placeholder="Masukan email" class="w-full md:w-[30rem] mb-3" v-model="email" />
                        <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-lg">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" :toggleMask="true" class="mb-4" fluid :feedback="false"></Password>
                        <Button label="Sign Up" class="w-full" @click="handleSubmit" :loading="loading"></Button>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Dialog start -->
    <Dialog v-model:visible="dialogInfo" :style="{ width: '450px' }" header="Warning" :modal="true" position="top">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span class="font-semibold">{{ searchTerm?.nama_sekolah }}</span> sudah terdaftar
        </div>
        <template #footer>
            <Button label="Ok" icon="pi pi-times" @click="dialogInfo = false" severity="warn" />
        </template>
    </Dialog>

    <Dialog v-model:visible="errorDialog">
        {{ errorInfo }}
    </Dialog>




    
    <!-- Dialog end -->
</template>

<style scoped>
.pi-eye {
    transform: scale(1.6);
    margin-right: 1rem;
}

.pi-eye-slash {
    transform: scale(1.6);
    margin-right: 1rem;
}
</style>
