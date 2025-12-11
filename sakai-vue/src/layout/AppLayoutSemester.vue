<template>
    <div>
        <div class="flex justify-end">
            <div class="flex items-center space-x-2">
                <div><span class="font-semibold">Semester</span></div>
                <div>
                    <SemesterComponent v-model="selectedSemester" :tahun-ajaran-id="tahunAjaranProvider?.label" :is-disabled="isDisabled" />
                </div>
            </div>
        </div>
        <div class="mt-4">
            <RouterView />
        </div>
    </div>
</template>
<script setup>
import SemesterComponent from '@/components/sekolah_components/SemesterComponent.vue';
import store from '@/store';
import { computed, inject, provide, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const isDisabled = computed(() => route.meta.disableSemester);

const selectedSemester = ref();
const tahunAjaranProvider = inject('tahunAjaranProvider');
// const selectedTahunAjaran = ref();
provide('selectedSemesterProvider', selectedSemester);

// watch(tahunAjaranProvider, (newVal) => {
//     if (newVal) {
//         selectedTahunAjaran.value = newVal.label;
//     }
// });
watch(selectedSemester, async (newVal) => {
    console.log(newVal);
    store.commit('semesterService/SET_SELECTEDSEMESTER', newVal);
});
</script>
