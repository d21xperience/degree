<script setup>
import { useDashboard } from '@/composables/sekolah_composable/useDashboard';
import { useSekolah } from '@/composables/sekolah_composable/useSekolah';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { onMounted, ref } from 'vue';
const { fetchDashboard } = useDashboard();

const { sekolah } = useSekolah();
const { initSelectedSemester } = useSemester();
// Ambil data dashboard
const dashboard = ref({
    data: {
        countSiswa: 0,
        countGuru: 0,
        countKelas: 0
    }
});
onMounted(async () => {
    dashboard.value = await fetchDashboard();
    alert('hello');
});
</script>

<template>
    <div>
        <div class="flex justify-between items-center">
            <div>
                Selamat datang
                <h3>{{ sekolah.sekolah.nama }}!</h3>
            </div>
            <h4>T.A. {{ initSelectedSemester?.namaSemester }}</h4>
        </div>
        <div class="grid grid-cols-12 gap-8">
            <StatsWidget label="Siswa" :target-number="Number(dashboard.data?.countSiswa)" icon="pi pi-users" url="infoSiswa" />
            <StatsWidget label="Guru" :target-number="Number(dashboard.data?.countGuru)" icon="pi pi-users" url="infoGuru" />
            <StatsWidget label="Kelas" :target-number="Number(dashboard.data?.countKelas)" icon="pi pi-building-columns" url="infoKelas" />
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
