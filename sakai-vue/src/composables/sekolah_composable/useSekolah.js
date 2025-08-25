import { computed } from 'vue';
import { useStore } from 'vuex';
import { useTableTenant } from './useTableTenant';
export function useSekolah() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const sekolah = computed(() => {
        const tes = store.getters['sekolahService/getSekolah'];
        const response = tes;
        response.uri = response?.sekolah.nama.toLowerCase().replace(/\s+/g, '');
        return response;
    });

    const fetchSekolah = async () => {
        try {
            let response = await store.getters['sekolahService/getSekolah'];
            if (!response) {
                const sekolahID = await store.state.authService.user?.sekolahTenantId;
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

    /**
     *
     * @param {String} rombonganBelajarId
     * @param {String} semesterId
     * @returns
     */
    const fetchAnggotaKelas = async (rombonganBelajarId = '', semesterId = '') => {
        try {
            const cachedData = await store.getters['sekolahService/getSiswaAktif'];
            if (cachedData.semester_id === semesterId) {
                const anggotaKelas = cachedData.peserta_didik.filter((val) => val.rombonganBelajarId === rombonganBelajarId);
                return anggotaKelas;
            }
            return null;
        } catch (error) {
            console.error('Gagal mengambil data kelas:', error);
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

    // const fetchMapel = async (mapel) => {
    //     try {
    //         let response = await store.dispatch('sekolahService/fetchMapel');

    //         if (response.status) {
    //             toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
    //         }

    //         return response.kategoriMapel;
    //     } catch (error) {
    //         toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
    //     }
    // };

    return {
        // fetchSiswa,

        fetchTingkat,
        fetchSekolah,
        createInfoIjazah,
        sekolah,
        updateSekolah,
        fetchAnggotaKelas
    };
}
