<script setup>
import TahunAjaranComponent from '@/components/sekolah_components/TahunAjaranComponent.vue';
import { useDashboard } from '@/composables/sekolah_composable/useDashboard';
import { useTableTenant } from '@/composables/sekolah_composable/useTableTenant';
import { computed, onMounted, ref } from 'vue';
import { useAuth } from './pages/auth/composables/auth';

const { user } = useAuth();

const { schemaname } = useTableTenant();
const { dashboard, initialize } = useDashboard();

// ✅ Derived state
const namaSekolah = computed(() => user?.sekolahAsal?.namaSekolah || 'Sekolah');
const cekTahunajaranId = ref();

onMounted(() => {
    setTimeout(() => {
        // console.log('schemaname', schemaname);
        // console.log('cekTahunajaranId', cekTahunajaranId);
        initialize(schemaname.value, cekTahunajaranId.value.label);
        if (!dashboard.value && schemaname.value && cekTahunajaranId.value.tahunAjaranId) {
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
            <TahunAjaranComponent v-model:model-value="cekTahunajaranId" />
        </div>

        <!-- 📊 Dashboard stats (langsung dari computed `dashboard`) -->
        <div v-if="dashboard" class="grid grid-cols-12 gap-8 my-4">
            <StatsWidget label="Siswa" :target-number="Number(dashboard.data?.countSiswa || 0)" icon="pi pi-users" url="infoSiswa" />
            <!-- <StatsWidget label="Guru" :target-number="Number(dashboard.data?.countGuru || 0)" icon="pi pi-users" url="infoGuru" /> -->
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
