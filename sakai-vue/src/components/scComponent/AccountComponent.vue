<template>
    <Select v-model="internalValue" :options="accountOptions" fluid placeholder="Pilih Account" :option-label="handleLabelOption" />
</template>

<script setup>
import { useSCService } from '@/composables/useSCService';
import { useUtils } from '@/composables/useUtils';
import { onMounted, ref, watch } from 'vue';
const scService = useSCService();
const utils = useUtils();
const internalValue = ref(null);
const accountOptions = ref([]);
// const isPasswordEnter = ref(false);
// const passwordInput = ref('');
// const passwordError = ref('');
// const isPasswordInvalid = ref(true); // Awalnya dianggap invalid
const props = defineProps(['modelValue']); // props dari parent
const emit = defineEmits(['update:modelValue', 'walletInfo']); // emit update ke parent

const handleLabelOption = (newVal) => {
    return `${utils.ringkasHash(newVal.address, 4, 4)}`;
};

watch(internalValue, async (newVal) => {
    if (newVal) {
        try {
            const response = await scService.getWalletInfo({ public_address: internalValue.value.address });
            // console.log(response)
            if (response) {
                console.log(response);
                // Object.assign(walletInfo,)
                // kirim ke parent component
                emit('walletInfo', response.walletData);
                // emit('update:modelValue', response.wallet);
            }
        } catch (error) {
            console.log(error);
        }
    }
});
// watch(isPasswordEnter, (newVal) => {
//     if (!newVal && isPasswordInvalid.value) {
//         internalValue.value = null;
//     }
// });

onMounted(async () => {
    accountOptions.value = await scService.fetchBCAccount();
});
</script>
