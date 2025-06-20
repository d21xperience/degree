<template>
    <!-- Dialog start -->
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Keluar" :modal="true" position="top">
        <!-- <AutoComplete
            v-model="internalValue"
            :suggestions="kurikulumOptions"
            option-label="namaKurikulum"
            @complete="searchKurikulum(query)"
            @keydown.space.prevent="kurikulumHandleKeydown"
            placeholder="Pilih kurikulum"
            fluid
            dropdown
            :loading="kurikulumLoading"
        /> -->
        <!-- <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="isVisible = false" severity="warn" />
            <Button label="Ya" icon="pi pi-check" text @click="onLogout" />
        </template> -->
    </Dialog>
    <!-- Dialog end -->
</template>
<script setup>
import { useSekolahService } from '@/composables/useSekolahService'; 
import { computed } from 'vue';

// const { fetchKurikulum, kurikulumLoading, searchKurikulum, kurikulumOptions } = useFormOptions();
// const internalValue = ref(props.modelValue);
const sekolahService = useFormOptions();
// const dialogSignOut = () => {};
// const onLogout = () => {
//     emit('confirm');
// };
const props = defineProps('visible','modelValue');
const emit = defineEmits(['update:visible', 'confirm']);

const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
});
// const cek = (data) => {
//     return `${data.kurikulumId} | ${data.namaKurikulum}`;
// };

// watch(
//     () => props.modelValue,
//     (newVal) => {
//         internalValue.value = newVal;
//     }
// );
// const query = ref();
// watch(internalValue, (newVal) => {
//     if (typeof newVal == 'object') {
//         query.value = newVal?.namaKurikulum;
//     } else if (typeof newVal == 'string') {
//         query.value = newVal;
//     }
//     emit('update:modelValue', newVal);
// });

// const kurikulumHandleKeydown = (event) => {
//     if (event.key === ' ') {
//         internalValue.value += ' '; // Menambahkan spasi ke query
//     }
// };
onMounted(async () => {
    await sekolahService.fetchKurikulum();
});
</script>
