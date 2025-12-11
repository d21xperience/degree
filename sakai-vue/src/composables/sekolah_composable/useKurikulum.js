import { computed, onMounted, readonly, ref } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk useKurikulum.
 */
export function useKurikulum({ autoload = true, immediate = true } = {}) {
    const store = useStore();
    const isFetching = ref(false);
    const error = ref(null);

    const kurikulumList = computed(() => {
        const raw = store.getters['kurikulumService/getKurikulum'];
        // console.log('raw', raw);
        if (!raw) return [];
        return [...raw.kurikulum];
    });

    /**
     *
     * @param {Number} jenjang_pendidikan_id
     * @returns Array
     */
    const _fetchKurikulum = async (jenjang_pendidikan_id = 6) => {
        if (isFetching.value) return;
        isFetching.value = true;
        error.value = null;

        try {
            const cached = store.getters['kurikulumService/getKurikulum'];
            // console.log('sebelum useKurikulum', cached);
            if (cached) return cached;
            // console.log('sesudah useKurikulum', cached);
            return await store.dispatch('kurikulumService/fetchKurikulum', jenjang_pendidikan_id);
        } catch (err) {
            error.value = err;
            console.error('[useKurikulum/fetchKurikulum] Error:', err);
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    const initialize = async (jenjang_pendidikan_id) => {
        await _fetchKurikulum(jenjang_pendidikan_id);
    };

    if (autoload && immediate) {
        onMounted(() => initialize().catch(console.warn));
    }

    // action

    return {
        kurikulumList: readonly(kurikulumList),
        isFetching: readonly(isFetching),
        error: readonly(error),
        initialize
        // actions...
    };
}

// import store from '@/store';

// export function useKurikulum() {

//     const fetchKurikulum = async (jenjang_pendidikan_id) => {
//         try {
//             // console.log(jenjang_pendidikan_id);
//             const response = await store.dispatch('kurikulumService/fetchKurikulum', { jenjangPendidikanId: jenjang_pendidikan_id });
//             return response;
//         } catch (error) {
//             console.error(error);
//             return new Error(`Gagal mengambil kurikulum di backend: ${error}`);
//         }
//     };

//     /**
//      *
//      * @param {Number} jenjang_pendidikan_id
//      * @returns Array
//      */
//     const getKurikulum = async (jenjang_pendidikan_id) => {
//         try {
//             let response = store.getters['kurikulumService/getKurikulum'];
//             console.log(!response || (Array.isArray(response) && response.length == 0) || jenjang_pendidikan_id != response.jenjangPendidikanId);
//             if (!response || (Array.isArray(response) && response.length == 0) || jenjang_pendidikan_id != response.jenjangPendidikanId) {
//                 response = await fetchKurikulum(jenjang_pendidikan_id);
//                 console.log(response);
//                 if (response.status) {
//                     return {
//                         status: response.status,
//                         message: response.message,
//                         kurikulum: response.kurikulum
//                     };
//                 }
//             }
//             return {
//                 status: true,
//                 message: 'kurikulumList berhasil diambil dari store',
//                 kurikulum: response.kurikulum
//                 // jenjangPendidikanId: response?.jenjang_pendidikan_id
//             };
//         } catch (error) {
//             return new Error(`Gagal mendapatkan kurikulum di store: ${error}`);
//         }
//     };

//     return {
//         fetchKurikulum,
//         getKurikulum
//     };
// }
