<script setup>
import { useSekolahService } from '@/composables/useSekolahService';
import { computed, onMounted, ref } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
const username = computed(() => store.getters['authService/getUserProfile']?.username);
const { fetchDashboard } = useSekolahService();
// Ambil data dashboard
const dashboard = ref();
onMounted(async () => {
    dashboard.value = await fetchDashboard();
});
</script>

<template>
    <div>
        <h3><span class="font-normal">Selamat datang</span> {{ username }} 🖐</h3>
        <div class="grid grid-cols-12 gap-8">
            <StatsWidget label="Siswa" :target-number="Number(dashboard?.countSiswa)" icon="pi pi-users" url="infoSiswa"/>
            <StatsWidget label="Guru" :target-number="Number(dashboard?.countGuru)" icon="pi pi-users" url="infoGuru"/>
            <StatsWidget label="Kelas" :target-number="Number(dashboard?.countKelas)" icon="pi pi-building-columns" url="infoKelas"/>
            <!-- <div class="col-span-12 xl:col-span-6">
            <RecentSalesWidget />
            <BestSellingWidget />
        </div>
        <div class="col-span-12 xl:col-span-6">
            <RevenueStreamWidget />
            <NotificationsWidget />
        </div> -->
        </div>
    </div>
</template>
