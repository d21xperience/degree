<script setup>
import { useDashboard } from '@/composables/sekolah_composable/useDashboard';
import { computed, onMounted, ref } from 'vue';
import { useAuth } from './pages/auth/composables/auth';
const { fetchDashboard } = useDashboard();

onMounted(async () => {
    try {
        dashboard.value = await fetchDashboard();
        isSekolahDashboard.value = true;
        isSCDashboard.value = true;
    } catch (error) {
        console.log(error);
    }
    // alert('hello');
});
// const { sekolah } = useSekolah();
const { currentUser } = useAuth();
// const { initSelectedSemester } = useSemester();
// Ambil data dashboard
const dashboard = ref({
    countSiswa: 0,
    countGuru: 0,
    countKelas: 0
});
const isSekolahDashboard = ref(false);
const isSCDashboard = ref(false);
const namaSekolah = computed(() => currentUser.value.asalSekolah);
</script>

<template>
    <div>
        <div class="flex justify-between items-center">
            <div>
                Selamat datang
                <h3>{{ namaSekolah }}!</h3>
                <!-- <h3 v-else>Sekolah tes!</h3> -->
            </div>
            <!-- <h4>T.A. {{ initSelectedSemester?.namaSemester }}</h4> -->
        </div>
        <div v-show="isSekolahDashboard" class="grid grid-cols-12 gap-8 my-4">
            <StatsWidget label="Siswa" :target-number="Number(dashboard.data?.countSiswa)" icon="pi pi-users" url="infoSiswa" />
            <StatsWidget label="Guru" :target-number="Number(dashboard.data?.countGuru)" icon="pi pi-users" url="infoGuru" />
            <StatsWidget label="Kelas" :target-number="Number(dashboard.data?.countKelas)" icon="pi pi-building-columns" url="infoKelas" />
        </div>
        <div class="col-span-12 xl:col-span-6">
            <NotificationsWidget />
            <RevenueStreamWidget />
        </div>
        <div class="col-span-12 xl:col-span-6">
            <RecentSalesWidget />
            <BestSellingWidget />
        </div>
    </div>
</template>
