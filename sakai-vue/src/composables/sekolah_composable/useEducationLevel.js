import { computed, readonly, ref } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk description.
 */
export function useEducationLevel() {
    const store = useStore();

    const isFetching = ref(false);
    const error = ref(null);

    const data = computed(() => {
        const raw = store.getters['namespace/getData'];
        if (!raw) return null;
        return { ...raw };
    });

    /**
     *
     * @param {Number} jenjangPendidikanId Wajib
     * @returns {Object} API Payload
     */
    const fetchEducationLevel = async (jenjangPendidikanId = 0) => {
        if (isFetching.value) return;
        isFetching.value = true;
        error.value = null;

        try {
            const cached = store.getters['sekolahService/getTingkatPendidikan'];
            if (cached) return cached;
            if (jenjangPendidikanId == 0) throw new Error('jenjangPendidikanId tidak ditemukan');

            return await store.dispatch('sekolahService/fetchTingkatPendidikan', {
                jenjang_pendidikan_id: jenjangPendidikanId
            });
        } catch (err) {
            error.value = err;
            console.error('[useName/_fetchName] Error:', err);
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    return {
        data: readonly(data),
        isFetching: readonly(isFetching),
        error: readonly(error),
        fetchEducationLevel
        // actions...
    };
}
