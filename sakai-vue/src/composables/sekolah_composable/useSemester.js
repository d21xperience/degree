import { computed, getCurrentInstance, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
import { useTableTenant } from './useTableTenant';
export function useSemester({ autoload = true } = {}) {
    const { schemaname } = useTableTenant();
    if (import.meta.env.DEV && !getCurrentInstance()) {
        console.error('[useSemester] Error: Composable ini harus dipanggil di dalam setup() atau <script setup>.');
    }

    const store = useStore();

    // --- 1. Getter-based computed (source of truth dari store)
    const listSemester = computed(() => store.getters['semesterService/getSemester'] || []);
    const listTahunAjaran = computed(() => store.getters['semesterService/getTahunAjaran'] || []);

    const storeSelectedSemester = computed(() => store.getters['semesterService/getSelectedSemester']);
    const storeSelectedTahunAjaran = computed(() => store.getters['semesterService/getSelectedTahunAjaran']);

    // --- 2. Local UI state (hanya untuk interaksi sementara sebelum commit)
    const selectedSemester = ref(null);
    const selectedTahunAjaran = ref(null);

    // --- 3. Helper: Load dari store, fallback ke backend jika kosong
    const initializeFromStoreOrFetch = async () => {
        // 🔁 Sync `selectedSemester`
        if (storeSelectedSemester.value != null) {
            selectedSemester.value = storeSelectedSemester.value;
        } else {
            // fallback: ambil dari backend & set ke store + local
            try {
                await fetchSemester();
                // Setelah fetch, coba lagi ambil dari getter
                if (storeSelectedSemester.value != null) {
                    selectedSemester.value = storeSelectedSemester.value;
                } else {
                    // fallback ke item pertama (opsional, sesuaikan logika bisnis)
                    selectedSemester.value = listSemester.value[0]?.semesterId ?? null;
                }
            } catch (err) {
                console.warn('Gagal inisialisasi semester:', err.message);
            }
        }

        // 🔁 Sync `selectedTahunAjaran`
        if (storeSelectedTahunAjaran.value != null) {
            selectedTahunAjaran.value = storeSelectedTahunAjaran.value;
        } else {
            try {
                await fetchTahunAjaran();
                if (storeSelectedTahunAjaran.value != null) {
                    selectedTahunAjaran.value = storeSelectedTahunAjaran.value;
                } else {
                    selectedTahunAjaran.value = listTahunAjaran.value[0]?.tahunAjaranId ?? null;
                }
            } catch (err) {
                console.warn('Gagal inisialisasi tahun ajaran:', err.message);
            }
        }
    };

    // --- 4. Watcher: Sync ke store hanya saat local state berubah (one-way: local → store)
    watch(selectedSemester, (newValue) => {
        if (newValue !== storeSelectedSemester.value) {
            store.commit('semesterService/SET_SELECTEDSEMESTER', newValue);
        }
    });

    watch(selectedTahunAjaran, (newValue) => {
        if (newValue !== storeSelectedTahunAjaran.value) {
            store.commit('semesterService/SET_SELECTEDTAHUNAJARAN', newValue);
        }
    });

    // --- 5. Actions (dibuat lebih reusable & clear responsibilities)

    const fetchSemester = async () => {
        try {
            console.log('dipanggil dari useSemester', schemaname.value);
            const result = await store.dispatch('semesterService/fetchSemester', schemaname.value);
            if (!result?.status) throw new Error(result?.message || 'Fetch semester failed');
            return result.semester;
        } catch (error) {
            const msg = error instanceof Error ? error.message : String(error);
            throw new Error(`Gagal mengambil semester: ${msg}`);
        }
    };

    const fetchTahunAjaran = async () => {
        // ✅ Optimasi: hindari fetch ulang jika sudah ada
        if (listTahunAjaran.value.length > 0) return listTahunAjaran.value;

        try {
            const result = await store.dispatch('semesterService/fetchTahunAjaran');
            if (!result?.status) throw new Error(result?.message || 'Fetch tahun ajaran failed');
            return result.tahunAjaran;
        } catch (error) {
            const msg = error instanceof Error ? error.message : String(error);
            throw new Error(`Gagal mendapatkan tahun ajaran: ${msg}`);
        }
    };

    const deleteSemester = async (semesters) => {
        if (!Array.isArray(semesters)) {
            throw new TypeError('Parameter must be an array of semester objects');
        }

        const semesterIds = semesters.map((s) => s.semesterId).filter((id) => typeof id === 'number' || typeof id === 'string');

        if (semesterIds.length === 0) {
            throw new Error('No valid semester IDs provided');
        }

        try {
            const response = await store.dispatch('semesterService/deleteSemester', semesterIds);
            if (!response?.status) throw new Error(response?.message || 'Deletion failed');
            return response;
        } catch (error) {
            const msg = error instanceof Error ? error.message : String(error);
            throw new Error(`Gagal menghapus semester: ${msg}`);
        }
    };

    const updateSemester = async (semester) => {
        if (!semester || typeof semester !== 'object') {
            throw new TypeError('Semester must be a non-null object');
        }

        try {
            const response = await store.dispatch('semesterService/updateSemester', semester);
            if (!response?.status) throw new Error(response?.message || 'Update failed');
            return response;
        } catch (error) {
            const msg = error instanceof Error ? error.message : String(error);
            throw new Error(`Gagal update semester: ${msg}`);
        }
    };

    // --- 6. Auto-init (opsional, bisa dipanggil manual juga)
    if (autoload) {
        onMounted(() => {
            initializeFromStoreOrFetch();
        });
    }

    // 📤 Return API public
    return {
        // State reaktif
        selectedSemester,
        selectedTahunAjaran,
        listSemester,
        listTahunAjaran,

        // Derived (opsional, tapi bisa di-return jika perlu)
        // initSelectedSemester: storeSelectedSemester,  // ← hindari duplikasi; cukup gunakan computed ini jika diperlukan
        // initSelectedTahunAjaran: storeSelectedTahunAjaran,

        // Actions
        fetchSemester,
        fetchTahunAjaran,
        deleteSemester,
        updateSemester,

        // Manual init (untuk edge case)
        initialize: initializeFromStoreOrFetch
    };
}
