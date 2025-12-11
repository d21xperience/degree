// composable/useDashboard.js
import { computed } from 'vue';
import { useStore } from 'vuex';

export function useDashboard() {
    const store = useStore();
    const dashboard = computed(() => {
        return store.getters['sekolahService/getDashboard'] || null;
    });

    // 🚀 Action utama
    const fetchDashboardSekolah = async (schemaname = '', semesterId = '') => {
        // 1. Cek cache dulu
        const cached = store.getters['sekolahService/getDashboard'];
        if (cached) return cached;

        // 2. Pastikan dependensi tersedia
        const schema = schemaname;
        // const semesterId = selectedSemester.value;

        if (!schema) {
            throw new Error('Schema name belum tersedia');
        }
        if (schema === '') {
            throw new Error('Schema name kosong');
        }
        if (semesterId == null) {
            throw new Error('Semester belum dipilih');
        }

        console.log(semesterId);
        // 3. Fetch
        const payload = {
            schemaname: schema,
            semester_id: semesterId
        };
        console.log(payload);
        return await store.dispatch('sekolahService/fetchDashboard', payload);
    };

    // 🔁 Init manual
    const initialize = async (schemaname, semesterId) => {
        if (!store.getters['sekolahService/getDashboard']) {
            await fetchDashboardSekolah(schemaname, semesterId);
        }
    };

    return {
        dashboard,
        fetchDashboardSekolah,
        initialize
    };
}
