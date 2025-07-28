<template>
    <Select v-model="internalValue" :options="accountOptions" fluid placeholder="Pilih Account" :option-label="handleLabelOption" />

    <Dialog v-model:visible="isPasswordEnter" position="top" :modal="true" style="width: 23rem">
        <div class="p-2">
            <div class="mb-2">Masukan Password</div>
            <InputText fluid v-model="passwordInput" type="password" placeholder="Ketik password blockhain Anda" :class="{ 'border-red-500': passwordError }" @keyup.enter="handleSubmitPassword" />
            <div v-if="passwordError" class="text-red-500 text-sm mt-1">
                {{ passwordError }}
            </div>
        </div>
        <template #footer>
            <div class="flex justify-end gap-2 p-2">
                <Button label="Batal" severity="secondary" @click="cancelPassword" />
                <Button label="Ok" @click="handleSubmitPassword" />
            </div>
        </template>
    </Dialog>
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { onMounted, ref, watch } from 'vue';

const scService = useSCService();
const utils = useUtils();
const internalValue = ref(null);
const accountOptions = ref([]);
const isPasswordEnter = ref(false);
const passwordInput = ref('');
const passwordError = ref('');
const isPasswordInvalid = ref(true); // Awalnya dianggap invalid
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue', 'wallet']); // emit update ke parent

const handleLabelOption = (newVal) => {
    return `${utils.ringkasHash(newVal.address, 4, 4)}`;
};

const wallet = ref({
    address: '',
    balance: '0',
    createdAt: '',
    isContract: false,
    label: ''
});

const handleSubmitPassword = async () => {
    if (!passwordInput.value) {
        passwordError.value = 'Password tidak boleh kosong';
        return;
    }

    // Jika password valid
    passwordError.value = '';
    isPasswordInvalid.value = false;
    isPasswordEnter.value = false;
    // console.log(payload);
    try {
        const payload = {
            password: passwordInput.value,
            account_address: internalValue.value.address
        };
        const response = await scService.getWalletInfo(payload);
        // console.log(response)
        if (response) {
            console.log(response);
            // kirim ke parent component
            emit('wallet', response.wallet);
            // emit('update:modelValue', response.wallet);
        } else {
            internalValue.value = null;
            return;
        }
    } catch (error) {
        console.log(error);
    }
};

const cancelPassword = () => {
    passwordInput.value = '';
    passwordError.value = '';
    internalValue.value = null;
    isPasswordEnter.value = false;
};

watch(internalValue, (newVal) => {
    if (newVal && isPasswordInvalid.value) {
        isPasswordEnter.value = true;
    }
});
watch(isPasswordEnter, (newVal) => {
    if (!newVal && isPasswordInvalid.value) {
        internalValue.value = null;
    }
});

onMounted(async () => {
    accountOptions.value = await scService.fetchBCAccount();
});
</script>
