<template>
    <div>
        <div class="">
            <div class="">
                <div class="flex justify-end">
                    <button @click="editProfile" class="hover:text-red-500" :class="{ editProfileClass: isProfileEdit }"><i class="pi pi-user-edit" style="font-size: 1.5rem"></i></button>
                </div>
                <div class="grid grid-cols-2">
                    <div class="flex justify-center">
                        <div class=" ">
                            <div class="relative">
                                <!-- Foto Profil -->
                                <img alt="Profile picture" class="w-32 h-32 rounded-full block border" height="150" :src="previewPhoto || akun.photo_url" width="150" />

                                <!-- Overlay untuk Upload -->
                                <div class="absolute w-full h-full top-0 left-0 flex items-center justify-center opacity-0 rounded-full hover:opacity-100 bg-black bg-opacity-50 text-gray-100 transition ease-in-out duration-500">
                                    <i @click="triggerFileInput" class="pi pi-camera cursor-pointer" style="font-size: 1.5rem"></i>
                                </div>

                                <!-- Input File (Disembunyikan) -->
                                <input ref="fileInput" type="file" accept="image/*" style="display: none" @change="handleFileSelect" />
                            </div>
                            <div class="text-center">
                                <h1 class="text-2xl font-bold">
                                    {{ akun.username }}
                                </h1>
                                <p class="text-gray-600">
                                    {{ akun.role.toUpperCase() }}
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">Personal Information</h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold"> Email: </span>
                                {{ akun.email }}
                            </p>
                            <p>
                                <span class="font-bold"> Phone: </span>
                                +123 456 7890
                            </p>
                            <div>
                                <span class="font-bold"> Nama: </span>
                                <div class="inline-block">
                                    <div v-if="isProfileEdit">
                                        {{ akun.nama }}
                                    </div>
                                    <div v-else>
                                        {{ akun.nama }}
                                    </div>
                                </div>
                            </div>
                            <div>
                                <span class="font-bold"> Tanggal Lahir: </span>
                                <div class="inline-block">
                                    <div v-if="isProfileEdit">
                                        <DatePicker v-model="akun.tglLahir" dateFormat="dd/mm/yy" :inputClass="{ dateClass: isProfileEdit }" />
                                    </div>
                                    <div v-else>
                                        {{ akun.tglLahir }}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <h5 class=" ">Alamat</h5>
                <div class="grid md:grid-cols-2 md:space-x-2 lg:grid-cols-4 lg:space-x-2">
                    <div>
                        <span class=""> Nama Jalan: </span>
                        <InputText fluid type="text" name="nama-jalan" id="nama-jalan" v-model="akun.alamatJalan" :class="{ 'p-inputtext': isProfileEdit }" />
                        <!-- 123 Main St -->
                    </div>
                    <div>
                        <span class=""> Kota/Kab: </span>
                        <InputText fluid type="text" name="kota-kab" id="kota-kab" v-model="akun.kotaKab" :class="{ 'p-inputtext': isProfileEdit }" />
                    </div>
                    <div>
                        <span class=""> Provinsi: </span>
                        <InputText fluid type="text" name="prov" id="prov" v-model="akun.prov" :class="{ 'p-inputtext': isProfileEdit }" />
                    </div>
                    <div>
                        <span class=""> Kode Pos: </span>
                        <InputText fluid type="text" name="kodepos" id="kodepos" v-model="akun.kodePos" :class="{ 'p-inputtext': isProfileEdit }" />
                    </div>
                    <!-- <div class="mt-6">
                        <h2 class="text-xl font-semibold">Orang tua</h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold"> Nama Ayah: </span>
                                <input type="text" name="tgl-lahir" id="nm-ayah" v-model="akun.namaAyah" :class="{ 'p-inputtext': isProfileEdit }" />
                            </p>
                            <p>
                                <span class="font-bold"> Nama Ibu: </span>
                                <input type="text" name="tgl-lahir" id="nama-ibu" v-model="akun.namaIbu" :class="{ 'p-inputtext': isProfileEdit }" />
                            </p>
                        </div>
                    </div> -->
                </div>
            </div>
        </div>
        <h5 class=" ">Riwayat Pendidikan</h5>
        <section class="grid grid-cols-2 space-x-2">
            <!-- <div class="mr-2">
                <label for="email">TK</label>
                <InputText class="w-full" type="email" name="email" id="email" />
            </div> -->
            <div>
                <label for="email">SD</label>
                <InputText class="w-full" type="email" name="email" id="email" />
            </div>
            <div>
                <label for="password">SMP</label>
                <InputText class="w-full" type="password" name="password" id="password" />
            </div>
        </section>
        <h5 class=" ">Change email & password</h5>
        <section class="mt-2">
            <div class="flex justify-between space-x-2">
                <div class="w-1/2">
                    <label for="email">Email</label>
                    <InputText class="w-full" type="email" name="email" id="email" />
                </div>
                <div class="w-1/2">
                    <label for="password">Password</label>
                    <InputText class="w-full" type="password" name="password" id="password" />
                </div>
            </div>
        </section>
        <div v-show="isProfileEdit" class="flex space-x-4 mt-16">
            <Button label="Simpan" class="min-w-32 mr-2" @click="showUpdateProfile" />
            <Button label="Batal" severity="warn" class="min-w-32" @click="editProfile" />
        </div>
    </div>
</template>

<script setup>
import DatePicker from 'primevue/datepicker';

import { onMounted, ref } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
// ==========[PROFILE]-----------
const fetchUserProfile = async () => {
    try {
        const userId = store.state.authService.user?.userId;
        if (!userId) throw new Error('User ID not found');
        // Dispatch untuk mendapatkan profil pengguna
        await store.dispatch('authService/getUserProfile', userId);

        // Ambil data terbaru dari store
        akun.value = store.getters['authService/getUserProfile'];
    } catch (error) {
        console.error('Failed to fetch user profile:', error.message);
    }
};
const showUpdateProfile = async () => {
    try {
        await store.dispatch('authService/updateUserProfile', akun.value);
    } catch (error) {
        console.log(error);
    }
};

onMounted(fetchUserProfile);
// ==============================
// State untuk menyimpan file yang dipilih
const selectedFile = ref(null);
// Biodata
// const akun = computed(() => store.getters["authService/getUserProfile"])
const akun = ref({
    userId: '',
    username: '',
    email: '',
    role: '',
    sekolahId: '',
    nama: '',
    jk: '',
    phone: '',
    tptLahir: '',
    tglLahir: '',
    alamatJalan: '',
    kotaKab: '',
    prov: '',
    kodePos: '',
    namaAyah: '',
    namaIbu: ''
    // photo_url: "@/assets/default_profile.jpg"
});
const isProfileEdit = ref(false);
const editProfile = () => {
    isProfileEdit.value = !isProfileEdit.value;
};
// const 'p-inputtext' = ref('edit')
// const editProfileClass = ref('tes')

const fileInput = ref(null);
const previewPhoto = ref(null);

// Data pengguna (contoh akun dengan photo_url dari backend)
// const akun = ref({
//     photo_url: "https://via.placeholder.com/150" // Ganti dengan URL default atau dari backend
// });

// Trigger input file saat ikon kamera diklik
const triggerFileInput = () => {
    fileInput.value.click();
};

// Handle file yang dipilih oleh pengguna
const handleFileSelect = (event) => {
    const file = event.target.files[0];
    if (file) {
        // Tampilkan preview foto sebelum upload
        previewPhoto.value = URL.createObjectURL(file);

        // Kirim ke backend
        uploadPhoto(file);
    }
};

// Upload foto ke backend
const uploadPhoto = async (file) => {
    const formData = new FormData();
    formData.append('photo', file);

    try {
        const response = await axios.post('http://localhost:8080/api/upload/photo', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });

        // Update photo_url dengan respons dari server
        akun.value.photo_url = response.data.photo_url;
    } catch (error) {
        console.error('Gagal upload foto:', error);
    }
};
</script>

<style scoped>
.p-inputext {
    padding: 0;
}

.editProfileClass {
    color: red;
}
</style>
