<template>
    <div class="">
        <!-- {{ storeSekolah }} -->
        <Accordion :value="[lengkapi]" multiple>
            <!-- Informasi sekolah -->
            <AccordionPanel value="0">
                <AccordionHeader>
                    <h5>Informasi Sekolah</h5>
                </AccordionHeader>
                <AccordionContent>
                    <div v-if="!storeSekolah">
                        <Skeleton class="mb-2" />
                        <Skeleton width="10rem" class="mb-2" />
                        <Skeleton width="5rem" class="mb-2" />
                        <Skeleton height="2rem" class="mb-2" />
                        <Skeleton width="10rem" height="4rem" />
                    </div>
                    <DataSekolahForm v-else :schemaname="schemaname" :store-sekolah="storeSekolah" />
                </AccordionContent>
            </AccordionPanel>
            <!--  KategoriSekolahComponent > -->
            <AccordionPanel value="2">
                <AccordionHeader>
                    <h5>Kompetensi keahlian Dilayani</h5>
                </AccordionHeader>
                <AccordionContent>
                    <KategoriSekolahComponent v-if="tahunAjaranAktif" :schemaname="schemaname" :tahun-ajaran-id="tahunAjaranAktif?.label" :jenjang-pendidikan-id="6" />
                </AccordionContent>
            </AccordionPanel>
        </Accordion>
    </div>
</template>

<script setup>
import DataSekolahForm from '@/components/sekolah_components/DataSekolahForm.vue';
import KategoriSekolahComponent from '@/components/sekolah_components/KategoriSekolahComponent.vue';
import { computed, inject } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();

// inject
const tahunAjaranAktif = inject('tahunAjaranProvider');
const schemaname = inject('schemanameProvider');
const storeSekolah = inject('sekolahProvider');
// const jenjangPendidikanId = computed(() => storeSekolah.value?.sekolah.jenjangPendidikanId);
// cek param
const lengkapi = computed(() => {
    if (route.query.lengkapi) {
        return String(route.query.lengkapi);
    } else {
        return '0';
    }
});
</script>
<style scoped>
edit-class {
    background-color: red;
}
</style>
