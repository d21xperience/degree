import { computed } from 'vue';
import { useStore } from 'vuex';
import { useTableTenant } from './useTableTenant';
export function useSekolah() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const sekolah = computed(() => {
        try {
            const tes = store.getters['sekolahService/getSekolah'];
            const response = tes;
            response.uri = response?.sekolah.nama.toLowerCase().replace(/\s+/g, '');
            return response;
        } catch (error) {
            return null;
        }
    });

    const fetchSekolah = async () => {
        try {
            let response = await store.getters['sekolahService/getSekolah'];
            if (!response) {
                const sekolahID = await store.state.authService.user?.sekolahAsal.id;
                const tTenant = await store.dispatch('sekolahService/fetchTabeltenant', sekolahID);
                response = await store.dispatch('sekolahService/fetchSekolah', { schemaname: tTenant.schemaname, namaSekolah: tTenant.namaSekolah });
                // response = await store.dispatch('sekolahService/fetchTabeltenant', response?.user.sekolahTenantId);
            }
            // console.log(response)
            return response;
        } catch (error) {
            console.log(error);
        }
    };
    const updateSekolah = async (param) => {
        try {
            const payload = {
                sekolah: param.sekolah,
                bentukPendidikanStr: param.bentukPendidikanStr,
                jenjangPendidikanStr: param.jenjangPendidikanStr
            };
            store.commit('sekolahService/SET_TABELSEKOLAH', payload);
            // await fetchSekolah();

            payload.schemaname = schemaname.value;
            const response = await store.dispatch('sekolahService/updateSekolah', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal update data sekolah: ${error.message}`);
        }
    };

    const fetchTingkat = async () => {
        try {
            let response = await store.getters['sekolahService/getTingkatPendidikan'];
            if (!response) {
                const payload = {
                    jenjang_pendidikan_id: await store.getters['sekolahService/getSekolah']?.sekolah.jenjangPendidikanId //sekolah.value?.jenjangPendidikanId
                };
                // console.log(payload);
                response = await store.dispatch('sekolahService/fetchTingkatPendidikan', payload);
            }
            return response;
        } catch (error) {
            throw new Error(`Gagal mengambil tingkat: ${error.message}`);
        }
    };

    const createInfoIjazah = async (dataInfoIjazah) => {
        const payload = {
            schemaname: schemaname.value,
            info_ijazah: dataInfoIjazah
        };
        const response = await store.dispatch('sekolahService/createInfoIjazah', payload);
        console.log(response);
    };

    const fetchBentukPendidikan = async () => {
        try {
            const { bentukPendidikan } = await store.dispatch('sekolahService/fetchBentukPendidikan');
            return bentukPendidikan;
        } catch (error) {
            throw new Error(error);
        }
    };
    /**
     * @param {Object} jenjang - Default value = jenjangLembaga = 1; jenjangOrang = 0
     * @returns Object
     */
    const fetchJenjangPendidikan = async (jenjang = { isJenjangLembaga: true, jenjangLembaga: 1, isJenjangOrang: false, jenjangOrang: 0 }) => {
        try {
            const response = await store.dispatch('sekolahService/fetchJenjangPendidikan', jenjang);
            console.log(response);
            return response.jenjang;
        } catch (error) {
            throw new Error(error);
        }
    };
    return {
        fetchTingkat,
        fetchSekolah,
        createInfoIjazah,
        sekolah,
        updateSekolah,
        fetchBentukPendidikan,
        fetchJenjangPendidikan
    };
}
