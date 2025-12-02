// composable/useDashboard.js
import { computed } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';

export function useDashboard({ autoload = false } = {}) {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { selectedSemester } = useSemester();

    // ✅ Derived state (opsional, tapi berguna)
    const dashboard = computed(() => {
        return store.getters['sekolahService/getDashboard'] || null;
    });

    // 🚀 Action utama
    const fetchDashboardSekolah = async () => {
        // 1. Cek cache dulu
        const cached = store.getters['sekolahService/getDashboard'];
        if (cached) return cached;

        // 2. Pastikan dependensi tersedia
        const schema = schemaname.value;
        const semesterId = selectedSemester.value;

        if (!schema) {
            throw new Error('Schema name belum tersedia');
        }
        if (schema === '') {
            throw new Error('Schema name kosong');
        }
        if (semesterId == null) {
            throw new Error('Semester belum dipilih');
        }

        // 3. Fetch
        const payload = {
            schemaname: schema,
            semester_id: semesterId
        };

        return await store.dispatch('sekolahService/fetchDashboard', payload);
    };

    // 🔁 Init manual
    const initialize = async () => {
        if (!store.getters['sekolahService/getDashboard']) {
            await fetchDashboardSekolah();
        }
    };

    // 🔄 Auto-init (opsional & aman)
    if (autoload) {
        // Karena dependensi async (semester/tenant), lebih baik **tidak auto-fetch di sini**,
        // tapi biarkan komponen yang kendalikan setelah dependensi siap.
        // Contoh: di komponen, panggil `initialize()` setelah `selectedSemester` terisi.
        // → Jadi `autoload: true` di sini **tidak otomatis fetch**, hanya sediakan kemudahan.
        // (Kita tetap sediakan opsi, tapi biarkan user yang trigger)
    }

    return {
        dashboard,
        fetchDashboardSekolah,
        initialize
    };
}
