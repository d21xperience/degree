<script setup>
import router from '@/router';
import { useAuth } from '@/views/pages/auth/composables/auth';
import { onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
const { login, isAuthLoading, authError } = useAuth();
// const username = ref('');
// const password = ref('');
const checked = ref(false);
// const loading = ref(false);
// const dialogError = ref(false);
// const messageError = ref('');
const route = useRoute();
onMounted(() => {
    if (route.query.from === 'register-success') {
        registerSuccessMessage.value = 'Silakan login untuk melanjutkan.';
        registerSuccessStatus.value = true;
    }
});

const registerSuccessMessage = ref('');
const registerSuccessStatus = ref(false);
// Fungsi handler submit form
// const handleSubmit = async () => {
//     loading.value = true;
//     try {
//         // Cek apakah username dan password diisi
//         if (username.value == '' || password.value == '') {
//             messageError.value = 'Username dan password tidak boleh kosong';
//             dialogError.value = true;
//             return;
//         }
//         // Kirim data form ke onLogin
//         await onLogin({
//             values: {
//                 username: username.value,
//                 password: password.value,
//                 rememberMe: checked.value
//             }
//         });
//     } catch (error) {
//         alert('gagal');
//         // messageError.value = error?.message;
//         // dialogError.value = true;
//     } finally {
//         // setTimeout(() => (loading.value[index] = false), 1000);
//         loading.value = false;
//     }
// };

// State UI
const username = ref('');
const password = ref('');
const rememberMe = ref(false);
const loading = ref(false);
const messageError = ref('');
const dialogError = ref(false);
// Sinkronkan loading & error ke UI (opsional — bisa juga pakai langsung ref dari composable)
watch(isAuthLoading, (val) => (loading.value = val));
watch(authError, (val) => {
    if (val) {
        messageError.value = val;
        dialogError.value = true;
    } else {
        messageError.value = '';
        dialogError.value = false;
    }
});
const handleSubmit = async () => {
    const result = await login({
        username: username.value,
        password: password.value,
        rememberMe: rememberMe.value
    });

    if (result.success && result.redirectRoute) {
        await router.push(result.redirectRoute);
    }
    // Jika gagal, UI sudah otomatis update via `authError` dan `dialogError`
};
</script>

<template>
    <!-- <FloatingConfigurator /> -->
    <div class="bg-surface-50 dark:bg-surface-950 flex items-center justify-center min-h-screen min-w-[100vw] overflow-hidden">
        <div class="flex flex-col items-center justify-center">
            <div style="border-radius: 56px; padding: 0.3rem; background: linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)">
                <div class="w-full bg-surface-0 dark:bg-surface-900 py-20 px-8 sm:px-20" style="border-radius: 53px">
                    <div class="text-center mb-8">
                        <!-- <div class="text-surface-900 dark:text-surface-0 text-3xl font-medium mb-4">Welcome to PrimeLand!</div> -->
                        <!-- <span class="text-muted-color font-medium">Sign in to continue</span> -->
                    </div>

                    <div>
                        <label for="email1" class="block text-surface-900 dark:text-surface-0 text-xl font-medium mb-2">Email atau Username</label>
                        <InputText id="email1" v-model="username" name="email1" type="text" placeholder="Masukan email atau username" class="w-full md:w-[30rem] mb-8" />

                        <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-xl mb-2">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" :toggle-mask="true" class="mb-4" fluid :feedback="false" />

                        <div class="flex items-center justify-between mt-2 mb-8 gap-8">
                            <div class="flex items-center">
                                <Checkbox id="rememberme1" v-model="checked" binary class="mr-2" name="rememberme1" />
                                <label for="rememberme1">Remember me</label>
                            </div>
                            <span class="font-medium no-underline ml-2 text-right cursor-pointer text-primary">Forgot password?</span>
                        </div>
                        <!-- <Button label="Sign In" class="w-full" as="router-link" to="/"></Button> -->
                        <Button label="Sign In" class="w-full" :loading="loading" @click="handleSubmit" />
                    </div>
                    <div class="mt-6 flex justify-between">
                        <div><router-link to="/" class="text-blue-600 hover:underline">Kembali</router-link></div>
                        <div>Belum punya akun?<router-link to="/auth/register" class="text-blue-600 hover:underline">Daftar</router-link></div>
                    </div>
                </div>
            </div>
        </div>

        <Dialog v-model:visible="dialogError" :style="{ width: '450px' }" header="Warning" :modal="true" position="top">
            <p>{{ messageError }}</p>
        </Dialog>
        <!-- <Dialog v-model:visible="registerSuccessStatus" :style="{ width: '450px' }" header="Akun berhasil dibuat" :modal="true" pt:mask:class="backdrop-blur-sm">
            <p class="text-xl">{{ registerSuccessMessage }}</p>
        </Dialog> -->

        <Dialog v-model:visible="registerSuccessStatus" modal :draggable="false" :breakpoints="{ '960px': '450px', '640px': '90vw' }" pt:mask:class="backdrop-blur-sm" header="Akun berhasil dibuat">
            <template #header>
                <div class="flex items-center gap-3 py-2">
                    <i class="pi pi-check-circle text-green-500 text-3xl"></i>
                    <span class="text-xl font-semibold">Selamat! 😀</span>
                </div>
            </template>

            <div class="py-4">
                <p class="text-lg leading-relaxed text-gray-700">Akun berhasil dibuat. {{ registerSuccessMessage }}</p>
            </div>

            <template #footer>
                <Button label="OK" icon="pi pi-check" class="p-button-success px-5 py-2" @click="registerSuccessStatus = false" />
            </template>
        </Dialog>
    </div>
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
