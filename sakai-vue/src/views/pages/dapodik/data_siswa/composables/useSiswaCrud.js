import { computed, onMounted, readonly, ref } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk description.
 */
export function useSiswaCrud({ autoload = true, immediate = true } = {}) {
    const store = useStore();

    const isFetching = ref(false);
    const error = ref(null);

    const data = computed(() => {
        const raw = store.getters['namespace/getData'];
        if (!raw) return null;
        return { ...raw };
    });

    const _fetchName = async () => {
        if (isFetching.value) return;
        isFetching.value = true;
        error.value = null;

        try {
            const cached = store.getters['namespace/getData'];
            if (cached) return cached;

            return await store.dispatch('namespace/fetchData', {
                /* payload */
            });
        } catch (err) {
            error.value = err;
            console.error('[useName/_fetchName] Error:', err);
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    const initialize = async () => {
        await _fetchName();
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
