<script setup>
import TopbarWidget from '@/components/landing/TopbarWidget.vue';
import { ref } from 'vue';
const shwoVerifikasi = ref(false);
</script>
<template>
    <!-- <div
        id="hero"
        class="flex flex-col pt-6 px-6 lg:px-20 overflow-hidden"
        style="background: linear-gradient(0deg, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.2)), radial-gradient(77.36% 256.97% at 77.36% 57.52%, rgb(238, 239, 175) 0%, rgb(195, 227, 250) 100%); clip-path: ellipse(150% 87% at 93% 13%)"
    > -->
    <div class="relative top-0 h-screen hero">
        <div class="py-6 px-6 mx-0 md:mx-12 lg:mx-20 lg:px-20 flex items-center justify-between bg-transparent text-white lg:static z-50">
            <TopbarWidget />
        </div>
        <!-- <div class="mx-6 md:mx-20 md:mt-6 mt-28"> -->
        <div class="absolute mx-6 md:mx-20 lg:top-1/3">
            <h1 class="text-6xl text-primary font-bold leading-tight"><span class="text-primary font-light block">Verifikasi Ijazah</span>dengan Blockchain</h1>
            <p class="font-normal text-2xl leading-6 md:mt-4 text-primary">Pastikan keaslian ijazah Anda dengan teknologi blockchain yang terpercaya dan transparan...</p>
            <Button label="Mulai verifikasi" @click="shwoVerifikasi = true" rounded class="!text-xl mt-8 !px-4"></Button>
        </div>
        <!-- <div class="flex justify-center md:justify-end">
            <img src="/demo/images/landing/screen-1.png"  alt="Hero Image" class="w-9/12 md:w-auto" />
        </div> -->
    </div>

    <Dialog v-model:visible="shwoVerifikasi">
        <div>
            <h3>Verifikasi Ijazah</h3>
            <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Voluptatem eveniet fuga aliquid quidem pariatur explicabo sapiente, exercitationem quam ad error veniam, nemo suscipit iure nulla maiores vitae nobis saepe ea.</p>
        </div>
    </Dialog>
</template>
<style>
.hero {
    background: url('/demo/images/landing/screen-2.png') center center/cover no-repeat;
    /* top: 0; */
    /* z-index: 999; */
}
</style>
