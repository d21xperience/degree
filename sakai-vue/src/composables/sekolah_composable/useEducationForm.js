import { computed, onMounted, readonly, ref } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk Bentuk Pendidikan.
 */
export function useEducationForm({ autoload = true, immediate = true } = {}) {
    const store = useStore();

    const isFetching = ref(false);
    const error = ref(null);

    const data = computed(() => {
        const raw = store.getters['sekolahService/getBentukPendidikan'];
        if (!raw) return [];
        return [...raw];
    });

    const _fetchEducationForm = async () => {
        if (isFetching.value) return;
        isFetching.value = true;
        error.value = null;

        try {
            const cached = store.getters['sekolahService/getBentukPendidikan'];
            if (cached) return cached;

            // const tes = await store.dispatch('sekolahService/fetchBentukPendidikan');
            // console.log(tes);
            return await store.dispatch('sekolahService/fetchBentukPendidikan', {
                /* payload */
            });
        } catch (err) {
            error.value = err;
            console.error('[useEducationForm/_fetchEducationForm] Error:', err);
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    const initialize = async () => {
        await _fetchEducationForm();
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
