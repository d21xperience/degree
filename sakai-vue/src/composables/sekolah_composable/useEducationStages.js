import { computed, onMounted, readonly, ref } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk Jenjang Pendidikan.
 */
export function useEducationStages({ autoload = true, immediate = true } = {}) {
    const store = useStore();

    const isFetching = ref(false);
    const error = ref(null);

    const data = computed(() => {
        const raw = store.getters['sekolahService/getJenjangPendidikan'];
        if (!raw) return [];
        return [...raw];
    });

    const _fetchEducationStages = async () => {
        if (isFetching.value) return;
        isFetching.value = true;
        error.value = null;

        try {
            const cached = store.getters['sekolahService/getJenjangPendidikan'];
            if (cached) return cached;
            const payload = {
                isJenjangLembaga: true,
                jenjangLembaga: 1,
                isJenjangOrang: false,
                jenjangOrang: 0
            };
            return await store.dispatch('sekolahService/fetchJenjangPendidikan', payload);
        } catch (err) {
            error.value = err;
            console.error('[useEducationStages/_fetchEducationStages] Error:', err);
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    const initialize = async () => {
        await _fetchEducationStages();
    };

    if (autoload && immediate) {
        onMounted(() => initialize().catch(console.warn));
    }

    return {
        data: readonly(data),
        isFetching: readonly(isFetching),
        error: readonly(error),
        initialize
        // actions...
    };
}
