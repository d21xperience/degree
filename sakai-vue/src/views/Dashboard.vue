<script setup>
import TahunAjaranComponent from '@/components/sekolah_components/TahunAjaranComponent.vue';
import { useDashboard } from '@/composables/sekolah_composable/useDashboard';
import { useSemester } from '@/composables/sekolah_composable/useSemester';
import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';
import { computed, onMounted, watch } from 'vue';
import { useAuth } from './pages/auth/composables/auth';

// useSemester
// 🔗 Composables
const { user } = useAuth();
const { selectedSemester } = useSemester({ autoload: false });
const { schemaname } = useTableTenant({ autoload: true });
const { dashboard, initialize } = useDashboard({ autoload: false });

// ✅ Derived state
const namaSekolah = computed(() => user?.sekolahAsal?.namaSekolah || 'Sekolah');

// 🔄 Fetch dashboard saat dependensi siap
watch(
    [schemaname, selectedSemester],
    ([schema, semesterId]) => {
        if (schema && schema !== '' && semesterId != null) {
            initialize().catch((err) => {
                console.error('[Dashboard] Gagal memuat data:', err.message);
                // Opsional: tampilkan notifikasi ke user
            });
        }
    },
    { immediate: true }
);

// 🎯 Opsional: fallback jika dashboard tetap null setelah beberapa detik
// (misal: user pertama kali, belum ada data)
onMounted(() => {
    setTimeout(() => {
        if (!dashboard.value && schemaname.value && selectedSemester.value) {
            console.warn('[Dashboard] Data masih kosong — mungkin belum ada entri di semester ini.');
        }
    }, 3000);
});
</script>

<template>
    <div>
        <div class="flex justify-between items-center">
            <div>
                Selamat datang
                <h3>{{ namaSekolah }}!</h3>
            </div>
            <!-- Semester/Tahun Ajaran selector -->
            <TahunAjaranComponent />
        </div>

        <!-- 📊 Dashboard stats (langsung dari computed `dashboard`) -->
        <div v-if="dashboard" class="grid grid-cols-12 gap-8 my-4">
            <StatsWidget label="Siswa" :target-number="Number(dashboard.data?.countSiswa || 0)" icon="pi pi-users" url="infoSiswa" />
            <StatsWidget label="Guru" :target-number="Number(dashboard.data?.countGuru || 0)" icon="pi pi-users" url="infoGuru" />
            <StatsWidget label="Kelas" :target-number="Number(dashboard.data?.countKelas || 0)" icon="pi pi-building-columns" url="infoKelas" />
        </div>

        <!-- 📝 Widget lain (pastikan tidak tergantung `dashboard`) -->
        <div class="grid grid-cols-12 gap-8 my-4">
            <div class="col-span-12 xl:col-span-6">
                <NotificationsWidget />
                <RevenueStreamWidget />
            </div>
            <!-- Kolom kanan bisa diisi lainnya -->
            <div class="col-span-12 xl:col-span-6">
                <!-- Opsional: placeholder atau widget tambahan -->
            </div>
        </div>
    </div>
</template>
