// src/composables/useTableTenant.js
import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk mengelola data tenant (termasuk schemaname)
 * Menjamin schemaname tersedia sebelum dipakai composable lain.
 *
 * @param {Object} options
 * @param {boolean} [options.autoload=true] - Mulai inisialisasi otomatis saat mounted
 * @returns {Object} API publik
 */
export function useTableTenant({ autoload = true } = {}) {
    // console.info('[COMPOSABLES] useTableTenant dipanggil!');
    const store = useStore();

    // ✅ Reaktif state
    const isLoading = ref(false);
    const error = ref(null);

    // ✅ Derived dari Vuex — single source of truth
    const schemaname = computed(() => {
        return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    });

    // ✅ Status kesiapan: schemaname tersedia & valid
    const isReady = computed(() => Boolean(schemaname.value));

    // ✅ Promise-based readiness guard (bisa di-await berkali-kali)
    let readyPromise = null;
    const ready = () => {
        if (isReady.value) {
            return Promise.resolve(schemaname.value);
        }

        if (readyPromise) {
            return readyPromise;
        }

        // Buat promise baru & pasang watcher
        readyPromise = new Promise((resolve, reject) => {
            const stopWatcher = watch(
                [isReady, error],
                ([isNowReady, currentError]) => {
                    if (isNowReady) {
                        stopWatcher();
                        resolve(schemaname.value);
                    } else if (currentError) {
                        stopWatcher();
                        reject(currentError);
                    }
                },
                { immediate: true } // cek sekali langsung
            );
        }).finally(() => {
            readyPromise = null; // reset agar bisa retry jika gagal
        });

        return readyPromise;
    };

    // ✅ Fetch dengan proteksi race condition & error handling
    const fetchTabelTenant = async () => {
        // Hindari multiple fetch bersamaan
        if (isLoading.value) {
            // Tunggu fetch sedang berjalan selesai
            return new Promise((resolve) => {
                const unwatch = watch(isLoading, (loading) => {
                    if (!loading) {
                        unwatch();
                        resolve(store.getters['sekolahService/getTabeltenant']);
                    }
                });
            });
        }

        // Cek cache dulu
        const cached = store.getters['sekolahService/getTabeltenant'];
        if (cached && schemaname.value) {
            return cached;
        }

        // Mulai fetch
        isLoading.value = true;
        error.value = null;

        try {
            // ✅ Ambil tenantId dari user (pastikan user sudah tersedia!)
            const user = store.getters['authService/user'];
            const tenantId = user?.sekolahAsal?.id;
            // console.log('diambil dari useTableTenant.js', tenantId);
            if (!tenantId) {
                throw new Error('[useTableTenant] tenantId tidak ditemukan — pastikan user sudah login dan memiliki sekolahAsal.id');
            }

            // 🔍 Log untuk debugging
            console.debug(`[useTableTenant] Fetching tenant data for ID: ${tenantId}`);

            // 📡 Panggil API
            const res = await store.dispatch('sekolahService/fetchTabeltenant', tenantId);

            // 🛑 Validasi respons
            if (!res?.status) {
                throw new Error(res?.message || 'Respons tidak valid dari fetchTabeltenant');
            }

            // ✅ Pastikan schemaname ada di data
            if (!res.schemaname) {
                console.warn('[useTableTenant] Respons sukses tapi schemaname tidak ditemukan:', res);
            }

            console.debug(`[useTableTenant] Schema loaded: "${res.schemaname}"`);

            return res;
        } catch (err) {
            error.value = err;
            console.error('[useTableTenant] Gagal memuat data tenant:', err);
            throw err;
        } finally {
            isLoading.value = false;
        }
    };

    // ✅ Inisialisasi eksplisit (bisa dipanggil manual)
    const initialize = async () => {
        if (isReady.value) return;
        await fetchTabelTenant();
    };

    // ✅ Auto-init saat komponen pertama kali mounted
    if (autoload) {
        onMounted(() => {
            // Jalankan secara asinkron, tanpa mengganggu render
            initialize().catch((err) => {
                // Jangan throw — biarkan komponen/pemanggil yang handle error utama
                console.warn('[useTableTenant] Inisialisasi otomatis gagal (autoload)', err.message);
            });
        });
    }

    // ✅ Public API
    return {
        // 🔹 State reaktif
        schemaname,
        isReady,
        isLoading,
        error,

        // 🔹 Kontrol alur
        ready, // ← gunakan ini di composable lain: await ready()
        initialize,
        fetchTabelTenant
    };
}
