<template>
    <div>
        <div class="py-6 px-6 mx-0 md:mx-12 lg:mx-20 lg:px-20 flex items-center justify-between bg-transparent text-white lg:static z-50">
            <TopbarWidget />
        </div>
        <section class="hero-section h-screen flex justify-center align-middle">
            <div class="hero-overlay"></div>
            <div class="hero-content">
                <h1 class="hero-headline">Ijazah Terverifikasi, Masa Depan Terjamin</h1>
                <p class="hero-subtitle">Pastikan keaslian ijazah Anda dengan teknologi blockchain yang aman dan transparan.</p>
                <button class="hero-cta" @click="verificationNow">Verifikasi Sekarang</button>
            </div>
            <div class="blockchain-pattern"></div>
        </section>

        <Dialog header="Verifikasi Ijazah" v-model:visible="showDialog" style="width: 32rem;">
            <div class="grid grid-cols-1 xl:grid-cols-2 gap-9 px-4 py-3">
                <div>NISN</div>
                <div class="w-full">
                    <InputText fluid placeholder="Masukan NISN"/>
                </div>
            </div>
            <div class="flex justify-end mt-6">
                <Button label="Verifikasi" class="w-32"/>
            </div>
        </Dialog>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const showDialog = ref(false);
const verificationNow = () => {
    showDialog.value = true;
};
</script>

<style scoped>
.hero-section {
    /* position: relative; */
    /* min-height: 500px; */
    /* display: flex; */
    /* align-items: center; */
    /* justify-content: center; */
    /* padding: 2rem; */
    /* overflow: hidden; */
    background: linear-gradient(135deg, #f0f7ff 0%, #e1edff 100%);
}

.hero-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-image: url('https://localhost/static/landingpage/7.jpg');
    background-repeat: no-repeat;
    background-size: cover;
    opacity: 0.7;
}

.blockchain-pattern {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100px;
    /* background-image: url('https://localhost/static/landingpage/1.jpg'); */
    /* background-repeat: repeat-x; */
    /* background-size: 1000px 100px; */
}

.hero-content {
    position: relative;
    z-index: 2;
    max-width: 800px;
    text-align: center;
    padding: 2rem;
}

.hero-headline {
    /* color: #05257e; */
    font-size: 2.5rem;
    font-weight: 700;
    line-height: 1.2;
    margin-bottom: 1.5rem;
}

.hero-subtitle {
    /* color: #2e3f5a; */
    font-size: 1.25rem;
    line-height: 1.6;
    margin-bottom: 2.5rem;
    max-width: 700px;
    margin-left: auto;
    margin-right: auto;
}

.hero-cta {
    background: linear-gradient(135deg, #d85305 0%, #0a2fd4 100%);
    color: white;
    font-weight: 600;
    font-size: 1.1rem;
    padding: 0.8rem 2.5rem;
    border: none;
    border-radius: 50px;
    cursor: pointer;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
}

.hero-cta:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.hero-cta::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0) 100%);
}

@media (max-width: 768px) {
    .hero-headline {
        font-size: 2rem;
    }

    .hero-subtitle {
        font-size: 1.1rem;
    }

    .hero-section {
        min-height: 400px;
    }
}

@media (max-width: 480px) {
    .hero-headline {
        font-size: 1.8rem;
    }
}
</style>
