<script setup>
import { debounce } from 'lodash-es';
import { onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
const props = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue', 'addMapel']);
const internalValue = ref(props.modelValue);
const mapelOptions = ref([]);
const isLoading = ref(false);

watch(
    () => props.modelValue,
    (newVal) => {
        internalValue.value = newVal;
    }
);
watch(internalValue, (newVal) => {
    // console.log(newVal);
    // if (typeof newVal == 'object') {
    //     // query.value = newVal?.namaKurikulum;
    // } else if (typeof newVal == 'string') {
    //     // query.value = newVal;
    // }
    emit('update:modelValue', newVal);
});

const handleKeydown = (event) => {
    if (event.key === ' ') {
        internalValue.value += ' '; // Menambahkan spasi ke query
    }
};
const filterMapel = debounce(async (searchTerm) => {
    try {
        const response = await store.dispatch('sekolahService/filterMapel', { query: searchTerm.query.toLowerCase() });
        mapelOptions.value = response;
        console.log(mapelOptions.value);
    } catch (error) {
        console.log(error);
    } finally {
        // ptkLoading.value = false;
    }
}, 250);
const addMapel = () => {
    emit('addMapel', internalValue);
};
onMounted(async () => {});
</script>
<template>
    <AutoComplete v-model="internalValue" :suggestions="mapelOptions" option-label="nama" @complete="filterMapel" @keydown.space.prevent="handleKeydown" placeholder="Cari mapel..." fluid :loading="isLoading" />
    <div class="flex justify-end mt-6">
        <!-- <Button label="Batal" icon="pi pi-times" text severity="warn" /> -->
        <Button label="Tambah" icon="pi pi-check" text @click="addMapel" />
    </div>
    <!-- </Dialog> -->
    <!-- Dialog end -->
</template>
