// composable/useTableTenant.js
import { computed, onMounted } from 'vue';
import { useStore } from 'vuex';

export function useTableTenant({ autoload = true } = {}) {
    const store = useStore();

    const schemaname = computed(() => {
        return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    });

    const fetchTabelTenant = async () => {
        const cached = store.getters['sekolahService/getTabeltenant'];
        if (cached) return cached;

        const tenantId = store.getters['authService/user'].sekolahAsal.id;
        console.log('tenant Id', tenantId);
        return await store.dispatch('sekolahService/fetchTabeltenant', tenantId);
    };

    const initialize = async () => {
        if (!store.getters['sekolahService/getTabeltenant']) {
            await fetchTabelTenant();
        }
    };

    // ✅ Auto-init saat komponen mounted — hanya jika autoload: true (default)
    if (autoload) {
        onMounted(() => {
            // Jalankan tanpa await (fire-and-forget), error ditangani di komponen
            initialize().catch((err) => {
                console.warn('[useTableTenant] Gagal inisialisasi otomatis:', err.message);
                // Catatan: Jangan throw di sini — biarkan komponen yang handle error utama
            });
        });
    }

    return {
        schemaname,
        fetchTabelTenant,
        initialize
    };
}
