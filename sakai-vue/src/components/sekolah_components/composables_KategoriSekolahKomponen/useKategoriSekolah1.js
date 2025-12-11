import { computed, onMounted, onUnmounted, ref } from 'vue';

/**
 * useSomething composable
 * Deskripsi singkat tentang apa yang dilakukan composable ini.
 *
 * @returns {
 *   state: object,
 *   actions: object,
 *   computed: object,
 * }
 */
export function useKategoriSekolahRev() {
    // =============================================================
    // State
    // =============================================================
    const data = ref(null);

    // =============================================================
    // Computed
    // =============================================================
    const computedValue = computed(() => {
        return data.value;
    });

    // =============================================================
    // Actions / Methods
    // =============================================================
    function fetchData() {
        // implementasi
    }

    // =============================================================
    // Lifecycle Hooks
    // =============================================================
    onMounted(() => {
        fetchData();
    });

    onUnmounted(() => {
        // cleanup jika diperlukan
    });

    // =============================================================
    // Return API (Ordered, Clean)
    // =============================================================
    return {
        // state
        data,

        // computed
        computedValue,

        // methods
        fetchData
    };
}
