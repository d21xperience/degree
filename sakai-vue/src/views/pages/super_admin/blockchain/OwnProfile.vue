<template>
    <div class="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-md">
        <h1 class="text-2xl font-bold text-gray-800 mb-6">Pengaturan Profil Website</h1>

        <form @submit.prevent="saveProfile" class="space-y-6">
            <!-- Informasi Dasar -->
            <div class="space-y-4">
                <h2 class="text-xl font-semibold text-gray-700 border-b pb-2">Informasi Dasar</h2>

                <div>
                    <label for="siteName" class="block text-sm font-medium text-gray-700">Nama Website</label>
                    <input v-model="profile.siteName" type="text" id="siteName" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Nama website Anda" />
                </div>

                <div>
                    <label for="siteDescription" class="block text-sm font-medium text-gray-700">Deskripsi Website</label>
                    <textarea
                        v-model="profile.siteDescription"
                        id="siteDescription"
                        rows="3"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                        placeholder="Deskripsi singkat tentang website Anda"
                    ></textarea>
                </div>

                <div>
                    <label for="logoUrl" class="block text-sm font-medium text-gray-700">URL Logo</label>
                    <input v-model="profile.logoUrl" type="url" id="logoUrl" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="https://example.com/logo.png" />
                    <div v-if="profile.logoUrl" class="mt-2">
                        <p class="text-sm text-gray-500">Pratinjau Logo:</p>
                        <img :src="profile.logoUrl" alt="Logo Preview" class="h-16 mt-1" />
                    </div>
                </div>
            </div>

            <!-- Informasi Kontak -->
            <div class="space-y-4">
                <h2 class="text-xl font-semibold text-gray-700 border-b pb-2">Informasi Kontak</h2>

                <div>
                    <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
                    <input v-model="profile.contact.email" type="email" id="email" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="contact@example.com" />
                </div>

                <div>
                    <label for="phone" class="block text-sm font-medium text-gray-700">Telepon</label>
                    <input v-model="profile.contact.phone" type="tel" id="phone" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="+62 123 4567 890" />
                </div>

                <div>
                    <label for="address" class="block text-sm font-medium text-gray-700">Alamat</label>
                    <textarea
                        v-model="profile.contact.address"
                        id="address"
                        rows="2"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                        placeholder="Alamat fisik perusahaan/organisasi"
                    ></textarea>
                </div>
            </div>

            <!-- Media Sosial -->
            <div class="space-y-4">
                <h2 class="text-xl font-semibold text-gray-700 border-b pb-2">Media Sosial</h2>

                <div v-for="(social, index) in profile.socialMedia" :key="index" class="flex items-end space-x-4">
                    <div class="flex-1">
                        <label :for="'socialPlatform' + index" class="block text-sm font-medium text-gray-700">Platform</label>
                        <select v-model="social.platform" :id="'socialPlatform' + index" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500">
                            <option value="">Pilih Platform</option>
                            <option value="facebook">Facebook</option>
                            <option value="twitter">Twitter/X</option>
                            <option value="instagram">Instagram</option>
                            <option value="linkedin">LinkedIn</option>
                            <option value="youtube">YouTube</option>
                            <option value="tiktok">TikTok</option>
                            <option value="whatsapp">WhatsApp</option>
                            <option value="telegram">Telegram</option>
                        </select>
                    </div>

                    <div class="flex-1">
                        <label :for="'socialUrl' + index" class="block text-sm font-medium text-gray-700">URL/Username</label>
                        <input v-model="social.url" type="text" :id="'socialUrl' + index" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Username atau URL lengkap" />
                    </div>

                    <button type="button" @click="removeSocial(index)" class="mb-1 px-3 py-1 bg-red-100 text-red-700 rounded-md hover:bg-red-200 transition">Hapus</button>
                </div>

                <button type="button" @click="addSocial" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition flex items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
                    </svg>
                    Tambah Media Sosial
                </button>
            </div>

            <!-- Informasi Tambahan -->
            <div class="space-y-4">
                <h2 class="text-xl font-semibold text-gray-700 border-b pb-2">Informasi Tambahan</h2>

                <div>
                    <label for="businessHours" class="block text-sm font-medium text-gray-700">Jam Operasional</label>
                    <textarea
                        v-model="profile.businessHours"
                        id="businessHours"
                        rows="2"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                        placeholder="Contoh: Senin-Jumat: 08:00-17:00"
                    ></textarea>
                </div>

                <div>
                    <label for="privacyPolicy" class="block text-sm font-medium text-gray-700">Kebijakan Privasi</label>
                    <textarea
                        v-model="profile.privacyPolicy"
                        id="privacyPolicy"
                        rows="3"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                        placeholder="Ringkasan kebijakan privasi atau link ke halaman lengkap"
                    ></textarea>
                </div>

                <div>
                    <label for="termsConditions" class="block text-sm font-medium text-gray-700">Syarat & Ketentuan</label>
                    <textarea
                        v-model="profile.termsConditions"
                        id="termsConditions"
                        rows="3"
                        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                        placeholder="Ringkasan syarat dan ketentuan atau link ke halaman lengkap"
                    ></textarea>
                </div>
            </div>

            <!-- Tombol Aksi -->
            <div class="flex justify-end space-x-4 pt-4">
                <button type="button" @click="resetForm" class="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 transition">Reset</button>
                <button type="submit" class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition">Simpan Profil</button>
            </div>
        </form>

        <!-- Notifikasi -->
        <div v-if="notification.show" :class="['mt-4 p-4 rounded-md', notification.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
            {{ notification.message }}
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';

// Data profil
const profile = ref({
    siteName: '',
    siteDescription: '',
    logoUrl: '',
    contact: {
        email: '',
        phone: '',
        address: ''
    },
    socialMedia: [{ platform: '', url: '' }],
    businessHours: '',
    privacyPolicy: '',
    termsConditions: ''
});

// Notifikasi
const notification = ref({
    show: false,
    message: '',
    success: false
});

// Memuat data dari localStorage
const loadProfile = () => {
    const savedProfile = localStorage.getItem('websiteProfile');
    if (savedProfile) {
        profile.value = JSON.parse(savedProfile);
    }
};

// Menyimpan data ke localStorage
const saveProfile = () => {
    try {
        localStorage.setItem('websiteProfile', JSON.stringify(profile.value));
        showNotification('Profil berhasil disimpan!', true);
    } catch (error) {
        showNotification('Gagal menyimpan profil: ' + error.message, false);
    }
};

// Menampilkan notifikasi
const showNotification = (message, isSuccess) => {
    notification.value = {
        show: true,
        message,
        success: isSuccess
    };
    setTimeout(() => {
        notification.value.show = false;
    }, 5000);
};

// Menambahkan media sosial baru
const addSocial = () => {
    profile.value.socialMedia.push({ platform: '', url: '' });
};

// Menghapus media sosial
const removeSocial = (index) => {
    if (profile.value.socialMedia.length > 1) {
        profile.value.socialMedia.splice(index, 1);
    } else {
        showNotification('Minimal harus ada satu media sosial', false);
    }
};

// Reset form
const resetForm = () => {
    if (confirm('Apakah Anda yakin ingin mengembalikan ke nilai default?')) {
        profile.value = {
            siteName: '',
            siteDescription: '',
            logoUrl: '',
            contact: {
                email: '',
                phone: '',
                address: ''
            },
            socialMedia: [{ platform: '', url: '' }],
            businessHours: '',
            privacyPolicy: '',
            termsConditions: ''
        };
        showNotification('Form telah direset', true);
    }
};

// Muat data saat komponen di-load
onMounted(() => {
    loadProfile();
});
</script>
