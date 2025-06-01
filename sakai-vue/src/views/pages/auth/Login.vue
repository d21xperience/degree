<script setup>
import FloatingConfigurator from '@/components/FloatingConfigurator.vue';
import { useAuth } from '@/views/pages/auth/composables/auth';
import { ref } from 'vue';
const { onLogin } = useAuth();

const username = ref('');
const password = ref('');
const checked = ref(false);
const loading = ref(false);
// Fungsi handler submit form
const handleSubmit = async () => {
    loading.value = true;
    try {
        // Kirim data form ke onLogin
        await onLogin({
            values: {
                username: username.value,
                password: password.value
            }
        });
    } catch (error) {
        alert(error?.message);
    } finally {
        // setTimeout(() => (loading.value[index] = false), 1000);
        loading.value = false;
    }
};
</script>

<template>
    <FloatingConfigurator />
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
                        <InputText id="email1" name="email1" type="text" placeholder="Masukan email atau username" class="w-full md:w-[30rem] mb-8" v-model="username" />

                        <label for="password1" class="block text-surface-900 dark:text-surface-0 font-medium text-xl mb-2">Password</label>
                        <Password id="password1" v-model="password" placeholder="Password" :toggleMask="true" class="mb-4" fluid :feedback="false"></Password>

                        <div class="flex items-center justify-between mt-2 mb-8 gap-8">
                            <div class="flex items-center">
                                <Checkbox v-model="checked" id="rememberme1" binary class="mr-2"></Checkbox>
                                <label for="rememberme1">Remember me</label>
                            </div>
                            <span class="font-medium no-underline ml-2 text-right cursor-pointer text-primary">Forgot password?</span>
                        </div>
                        <!-- <Button label="Sign In" class="w-full" as="router-link" to="/"></Button> -->
                        <Button label="Sign In" class="w-full" @click="handleSubmit" :loading="loading"></Button>
                    </div>
                    <!-- <div class="mt-6 flex justify-between">
                        <div><router-link to="/" class="text-blue-600 hover:underline">Kembali</router-link></div>
                        <div>Belum punya akun?<router-link to="/auth/register" class="text-blue-600 hover:underline">Daftar</router-link></div>
                    </div> -->
                </div>
            </div>
        </div>
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
