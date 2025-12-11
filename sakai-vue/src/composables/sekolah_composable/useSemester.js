// src/composables/useSemester.js
import { computed, getCurrentInstance, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';
import { useTableTenant } from './useTableTenant';

/**
 * Composable untuk mengelola data semester & tahun ajaran
 * Bergantung pada schemaname dari useTableTenant — menunggu hingga siap.
 *
 * @param {Object} options
 * @param {boolean} [options.autoload=true] - Jalankan inisialisasi otomatis saat mounted
 * @returns {Object} API publik
 */
export function useSemester({ autoload = true } = {}) {
    // console.log('[COMPOSABLES] useSemester dipanggil!!');
    // 🔒 Pastikan dipanggil dalam setup()
    if (import.meta.env.DEV && !getCurrentInstance()) {
        console.error('[useSemester] Error: Composable ini harus dipanggil di dalam setup() atau <script setup>.');
    }

    const store = useStore();
    // 📦 Dependency: schemaname dari tenant
    const { schemaname, isReady: isSchemaReady, ready: ensureSchema } = useTableTenant();

    // === 📊 STATE REAKTIF ===

    // 🔹 Getter-based (source of truth dari store)
    const listSemester = computed(() => store.getters['semesterService/getSemester'] || []);
    const listTahunAjaran = computed(() => store.getters['semesterService/getTahunAjaran'] || []);

    const storeSelectedSemester = computed(() => store.getters['semesterService/getSelectedSemester']);
    const storeSelectedTahunAjaran = computed(() => store.getters['semesterService/getSelectedTahunAjaran']);

    // 🔹 Local UI state (untuk interaksi sementara sebelum commit ke store)
    const selectedSemester = ref(null);
    const selectedTahunAjaran = ref(null);

    // 🔹 Loading & error state (opsional — bisa dipakai di UI)
    const isLoading = ref(false);
    const error = ref(null);

    // === 🔄 SYNC DENGAN STORE ===

    // Sinkronisasi dari store ke local ref (one-way: store → local)
    watch(storeSelectedSemester, (val) => {
        if (val !== selectedSemester.value) selectedSemester.value = val;
    });

    watch(storeSelectedTahunAjaran, (val) => {
        if (val !== selectedTahunAjaran.value) selectedTahunAjaran.value = val;
    });

    // Sinkronisasi balik: local → store (one-way, hanya saat berubah)
    watch(selectedSemester, (newVal) => {
        if (newVal !== storeSelectedSemester.value) {
            store.commit('semesterService/SET_SELECTEDSEMESTER', newVal);
        }
    });

    watch(selectedTahunAjaran, (newVal) => {
        if (newVal !== storeSelectedTahunAjaran.value) {
            store.commit('semesterService/SET_SELECTEDTAHUNAJARAN', newVal);
        }
    });

    // === ⚙️ HELPER: Fetch & Inisialisasi ===

    /**
     * Ambil data semester dari backend (jika belum ada di store)
     * ✅ Menunggu schemaname siap secara eksplisit
     */
    const fetchSemester = async () => {
        try {
            isLoading.value = true;
            error.value = null;

            // ✅ Pastikan schemaname siap — ini kunci utama!
            const schema = await ensureSchema();
            if (!schema) {
                throw new Error('Schema name required but not available');
            }

            const result = await store.dispatch('semesterService/fetchSemester', schema);
            if (!result?.status) {
                throw new Error(result?.message || 'Gagal mengambil data semester');
            }

            return result.semester;
        } catch (err) {
            error.value = err;
            // console.error('[useSemester] fetchSemester error:', err);
            throw err;
        } finally {
            isLoading.value = false;
        }
    };

    /**
     * Ambil data tahun ajaran dari backend
     * 🔁 Hindari fetch ulang jika sudah ada
     */
    const fetchTahunAjaran = async () => {
        if (listTahunAjaran.value.length > 0) {
            return listTahunAjaran.value;
        }

        try {
            isLoading.value = true;
            error.value = null;

            const result = await store.dispatch('semesterService/fetchTahunAjaran');
            if (!result?.status) {
                throw new Error(result?.message || 'Gagal mengambil data tahun ajaran');
            }

            return result.tahunAjaran;
        } catch (err) {
            error.value = err;
            console.error('[useSemester] fetchTahunAjaran error:', err);
            throw err;
        } finally {
            isLoading.value = false;
        }
    };

    /**
     * Inisialisasi state semester/tahun ajaran:
     * 1. Coba dari store (selected)
     * 2. Jika tidak ada → fetch dari backend
     * 3. Fallback ke item pertama jika perlu
     */
    const initializeFromStoreOrFetch = async () => {
        // 🔁 Sync selected semester
        if (storeSelectedSemester.value != null) {
            selectedSemester.value = storeSelectedSemester.value;
        } else {
            try {
                await fetchSemester();
                // Setelah fetch, cek lagi
                if (storeSelectedSemester.value != null) {
                    selectedSemester.value = storeSelectedSemester.value;
                } else if (listSemester.value.length > 0) {
                    // Fallback: semester aktif atau item pertama
                    const active = listSemester.value.find((s) => s.isActive);
                    selectedSemester.value = active?.semesterId ?? listSemester.value[0].semesterId;
                }
            } catch (err) {
                console.warn('[useSemester] Gagal inisialisasi semester:', err.message);
            }
        }

        // 🔁 Sync selected tahun ajaran
        if (storeSelectedTahunAjaran.value != null) {
            selectedTahunAjaran.value = storeSelectedTahunAjaran.value;
        } else {
            try {
                await fetchTahunAjaran();
                if (storeSelectedTahunAjaran.value != null) {
                    selectedTahunAjaran.value = storeSelectedTahunAjaran.value;
                } else if (listTahunAjaran.value.length > 0) {
                    const active = listTahunAjaran.value.find((t) => t.isActive);
                    selectedTahunAjaran.value = active?.tahunAjaranId ?? listTahunAjaran.value[0].tahunAjaranId;
                }
            } catch (err) {
                console.warn('[useSemester] Gagal inisialisasi tahun ajaran:', err.message);
            }
        }
    };

    // === 🧰 MUTATION ACTIONS (CRUD) ===

    const deleteSemester = async (semesters) => {
        if (!Array.isArray(semesters)) {
            throw new TypeError('Parameter `semesters` harus berupa array');
        }

        const semesterIds = semesters.map((s) => s.semesterId).filter((id) => id != null && (typeof id === 'number' || typeof id === 'string'));

        if (semesterIds.length === 0) {
            throw new Error('Tidak ada ID semester yang valid untuk dihapus');
        }

        try {
            isLoading.value = true;
            error.value = null;

            const response = await store.dispatch('semesterService/deleteSemester', semesterIds);
            if (!response?.status) {
                throw new Error(response?.message || 'Gagal menghapus semester');
            }

            return response;
        } catch (err) {
            error.value = err;
            console.error('[useSemester] deleteSemester error:', err);
            throw err;
        } finally {
            isLoading.value = false;
        }
    };

    const updateSemester = async (semester) => {
        if (!semester || typeof semester !== 'object') {
            throw new TypeError('Parameter `semester` harus berupa objek non-null');
        }

        try {
            isLoading.value = true;
            error.value = null;

            const response = await store.dispatch('semesterService/updateSemester', semester);
            if (!response?.status) {
                throw new Error(response?.message || 'Gagal memperbarui semester');
            }

            return response;
        } catch (err) {
            error.value = err;
            console.error('[useSemester] updateSemester error:', err);
            throw err;
        } finally {
            isLoading.value = false;
        }
    };

    // === 🚀 AUTO-INIT (opsional) ===

    if (autoload) {
        onMounted(() => {
            // Jalankan async tanpa mengganggu render
            initializeFromStoreOrFetch().catch((err) => {
                console.warn('[useSemester] Inisialisasi otomatis gagal:', err.message);
            });
        });
    }

    // === 📤 PUBLIC API ===

    return {
        // 🔹 State
        schemaname, // exposed for debugging or derived logic
        selectedSemester,
        selectedTahunAjaran,
        listSemester,
        listTahunAjaran,

        // 🔹 Status
        isLoading,
        error,
        isSchemaReady,

        // 🔹 Actions
        fetchSemester,
        fetchTahunAjaran,
        deleteSemester,
        updateSemester,
        initialize: initializeFromStoreOrFetch,

        // 🔹 Helper (opsional)
        ensureSchema // expose untuk composable lain yang butuh schemaname
    };
}
