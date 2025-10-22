import store from '@/store';

export function useKurikulum() {
    /**
     *
     * @param {Number} jenjang_pendidikan_id
     * @returns Array
     */
    const fetchKurikulum = async (jenjang_pendidikan_id) => {
        try {
            // console.log(jenjang_pendidikan_id);
            const response = await store.dispatch('kurikulumService/fetchKurikulum', { jenjangPendidikanId: jenjang_pendidikan_id });
            return response;
        } catch (error) {
            console.error(error);
            return new Error(`Gagal mengambil kurikulum di backend: ${error}`);
        }
    };

    /**
     *
     * @param {Number} jenjang_pendidikan_id
     * @returns Array
     */
    const getKurikulum = async (jenjang_pendidikan_id) => {
        try {
            let response = store.getters['kurikulumService/getKurikulum'];
            console.log(!response || (Array.isArray(response) && response.length == 0) || jenjang_pendidikan_id != response.jenjangPendidikanId);
            if (!response || (Array.isArray(response) && response.length == 0) || jenjang_pendidikan_id != response.jenjangPendidikanId) {
                response = await fetchKurikulum(jenjang_pendidikan_id);
                console.log(response);
                if (response.status) {
                    return {
                        status: response.status,
                        message: response.message,
                        kurikulum: response.kurikulum
                    };
                }
            }
            return {
                status: true,
                message: 'Data berhasil diambil dari store',
                kurikulum: response.kurikulum
                // jenjangPendidikanId: response?.jenjang_pendidikan_id
            };
        } catch (error) {
            return new Error(`Gagal mendapatkan kurikulum di store: ${error}`);
        }
    };

    return {
        fetchKurikulum,
        getKurikulum
    };
}
