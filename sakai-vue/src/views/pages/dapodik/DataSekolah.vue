<script setup>
import DataSekolahComponent from '@/components/sekolah_components/DataSekolahComponent.vue';
import KategoriSekolahComponent from '@/components/sekolah_components/KategoriSekolahComponent.vue';
</script>

<template>
    <div class="">
        <Accordion :value="['']" multiple>
            <AccordionPanel value="0">
                <AccordionHeader>
                    <h5>Informasi Sekolah</h5>
                </AccordionHeader>
                <AccordionContent>
                    <DataSekolahComponent />
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="1">
                <AccordionHeader>
                    <h5>Semester</h5>
                </AccordionHeader>
                <AccordionContent>
                    <article class="p-2">
                        <p class="mt-2">Web Service digunakan untuk mengsisi data dapodik melalui aplikasi yang bisa didownload pada web di bawah ini</p>
                        <p>Untuk menggunakan aplikasi tersebut silahkan download aplikasi</p>
                    </article>
                    <!-- <KategoriSekolahComponent /> -->
                </AccordionContent>
            </AccordionPanel>
            <!-- Informasi sekolah -->
            <AccordionPanel value="2">
                <AccordionHeader>
                    <h5>Kompetensi keahlian Dilayani</h5>
                </AccordionHeader>

                <AccordionContent>
                    <KategoriSekolahComponent />
                </AccordionContent>
            </AccordionPanel>

            <AccordionPanel value="3">
                <AccordionHeader>
                    <h5>Web Service</h5>
                    <!-- <Button v-tooltip.bottom="'Edit data sekolah'" icon="pi pi-pencil" :style="isEdit ? 'background-color:blue;border:none' : 'background-color:gray;border:none'" size="small" rounded="" @click="editSekolah" /> -->
                </AccordionHeader>
                <AccordionContent>
                    <article class="p-2">
                        <p class="mt-2">Web Service digunakan untuk mengsisi data dapodik melalui aplikasi yang bisa didownload pada web di bawah ini</p>
                        <p>Untuk menggunakan aplikasi tersebut silahkan download aplikasi</p>
                    </article>
                    <!-- <KategoriSekolahComponent /> -->
                </AccordionContent>
            </AccordionPanel>
        </Accordion>
    </div>
</template>

<style scoped>
edit-class {
    background-color: red;
}
</style>
