<script setup>
import { ref } from 'vue';
// import { useStore } from 'vuex';
// const store = useStore();

import SekolahComponent from '@/components/sekolah_components/SekolahComponent.vue';
import { isObject } from '@/utils/format';
import { useAuth } from '@/views/pages/auth/composables/auth';
const { cekSekolahByNPSN, onRegisterAdmin } = useAuth();
const email = ref('');
const password = ref('');
const loading = ref(false);

const dialogInfo = ref(false);
const errorDialog = ref(false);
const errorInfo = ref();
const statusSekolahTerdaftar = ref(false);
const searchTerm = ref('');
// const npsn = ref();
// const error = ref();
const resetSearchTerm = () => {
    searchTerm.value = '';
    dialogInfo.value = false;
};

const cekSekolah = async () => {
    statusSekolahTerdaftar.value = await cekSekolahByNPSN(searchTerm.value?.npsn);
    if (!statusSekolahTerdaftar.value) {
        // console.log(statusSekolahTerdaftar.value);
        dialogInfo.value = true;
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
    try {
        dataReg.sekolah = formatValues(dataReg.sekolah);
        const response = await onRegisterAdmin(dataReg);
        //     // console.log(response);
        // Jika sukses, arahkan ke beranda
        if (response.ok) {
            // const result = response?.sekolahTenant.namaSekolah.toLowerCase().replace(/\s+/g, '');
            // router.push({ name: 'dashboard', params: { sekolah: result } });
            // await store.dispatch('sekolahService/fetchTahunAjaran');
            // await store.dispatch('sekolahService/fetchSemester');
            // const result = response?.sekolahTenant.namaSekolah.toLowerCase().replace(/\s+/g, '');
            // await store.dispatch('sekolahService/fetchTabeltenant', response?.user.sekolahTenantId);
            // await router.push({ name: 'dashboard', params: { sekolah: result } });
        }
        //     // success.value = 'Admin registered successfully!';
    } catch (error) {
        //     errorDialog.value = true;
        //     errorInfo.value = error?.message;
        //     console.error(error);
        //     // error.value = err.error || 'Registration failed';
    } finally {
        loading.value = false;
    }
    // return { name, email, password, schoolName, register, error, success };
};

// Fungsi untuk menghapus spasi dan konversi ke string
const formatValues = (obj) => {
    return Object.fromEntries(Object.entries(obj).map(([key, value]) => [key, String(value).trim()]));
};
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
                                <SekolahComponent v-model:modelValue="searchTerm" />
                            </div>
                            <div>
                                <Button label="Cek" class="w-24" @click="cekSekolah" :loading="loading" :disabled="!isObject(searchTerm)"></Button>
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
                    <!-- <div class="text-center mb-8"> 
                    </div> -->

                    <!-- <div>
                        <label for="sekolah" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Sekolah</label>
                        <InputText id="sekolah" name="sekolah" type="text" placeholder="Masukan sekolah" class="w-full md:w-[30rem] mb-3" v-model="searchTerm.nama_sekolah" disabled />
                        <label for="email1" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Email</label>
                        <InputText id="email1" name="email1" type="text" placeholder="Masukan email" class="w-full md:w-[30rem] mb-3" v-model="email" />
                        <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-lg">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" :toggleMask="true" class="mb-4" fluid :feedback="false"></Password>
                        <Button label="Sign Up" class="w-full" @click="handleSubmit" :loading="loading"></Button>
                    </div> -->

                    <form @submit.prevent="handleSubmit">
                        <div>
                            <label for="sekolah" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Sekolah</label>
                            <InputText id="sekolah" name="sekolah" type="text" placeholder="Masukan sekolah" class="w-full md:w-[30rem] mb-3" v-model="searchTerm.nama_sekolah" disabled />

                            <label for="email1" class="block text-surface-900 dark:text-surface-0 text-lg font-medium">Email</label>
                            <InputText id="email1" name="email1" type="email" placeholder="Masukan email" class="w-full md:w-[30rem] mb-3" v-model="email" />

                            <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-lg">Password</label>
                            <Password id="password1" name="password1" v-model="password" placeholder="Password" :toggleMask="true" class="mb-4" fluid :feedback="false" autocomplete="new-password"></Password>

                            <Button type="submit" label="Sign Up" class="w-full" :loading="loading"></Button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>

    <!-- Dialog start -->
    <Dialog v-model:visible="dialogInfo" :style="{ width: '450px' }" header="Warning" :modal="true" position="top">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span class="font-semibold">{{ searchTerm?.nama_sekolah }}</span
            >sudah terdaftar!!
        </div>
        <template #footer>
            <Button label="Ok" icon="pi pi-times" @click="resetSearchTerm" severity="warn" />
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
