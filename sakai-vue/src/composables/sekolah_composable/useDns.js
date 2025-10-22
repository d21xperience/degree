import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useDns() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester, initSelectedTahunAjaran } = useSemester();

    const addDns = async (dns) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`,
                data_nominasi_sementara: dns
            };
            // console.log(payload);
            // return;
            const response = await store.dispatch('dnsService/createDns', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menambahkan DNS:', error);
        }
    };
    const updateDns = async (dns) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`,
                data_nominasi_sementara: dns
            };
            const response = await store.dispatch('dnsService/updateDns', payload);
            if (response.status) {
                const dnsTabel = store.getters['dnsService/getDns'];
                const dns = dnsTabel.dataNominasiSementara.find((item) => item.pesertaDidikId == payload.data_nominasi_sementara.pesertaDidikId);
                if (dns) {
                    Object.assign(dns, payload.data_nominasi_sementara);
                }
                store.commit('dnsService/SET_TABELDNS', dnsTabel);
                return true;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal update DNS:', error);
        }
    };
    const getDns = async (tahunAjaranId) => {
        try {
            let response = await store.getters['dnsService/getDns'];
            // console.log(response)
            if (!response || !Array.isArray(response.dataNominasiSementara) || response.dataNominasiSementara.length === 0 || response.tahun_ajaran_id != tahunAjaranId) {
                const payload = {
                    schemaname: schemaname.value,
                    tahun_ajaran_id: tahunAjaranId,
                    is_complete: false
                };
                response = await store.dispatch('dnsService/fetchDns', payload);
                console.log(response);
                if (response) {
                    return response.dataNominasiSementara;
                } else {
                    return [];
                }
            }
            return response.dataNominasiSementara;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan DNS:', error);
        }
    };
    const searchDns = async (pesertaDidikId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedTahunAjaran.value?.tahunAjaranId}`,
                peserta_didik_id: pesertaDidikId
            };
            // console.log(payload);
            // return;
            const response = await store.dispatch('dnsService/searchDns', payload);
            return response.dataNominasiSementara;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian DNS:', error);
        }
    };
    const searchDnsLokal = async (pesertaDidikId) => {
        try {
            const dnsTabel = store.getters['dnsService/getDns'];
            const dns = dnsTabel.dataNominasiSementara.find((item) => item.pesertaDidikId == pesertaDidikId);
            // console.log(dns);
            // return
            return dns;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal melakukan pencarian DNS lokal:', error);
        }
    };

    const deleteDns = async (pesertaDidikId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                peserta_didik_id: pesertaDidikId
            };
            console.log(payload);
            const response = await store.dispatch('dnsService/deleteDns', payload);
            return response;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus DNS:', error);
        }
    };
    // const deleteBatchDns = async (anggotaRombelIds) => {
    //     // console.log(anggotaRombelIds);

    //     // return;
    //     try {
    //         const payload = {
    //             schemaname: schemaname.value,
    //             anggota_rombel_id: anggotaRombelIds
    //         };
    //         const response = await store.dispatch('dnsService/deleteBatchAnggotaKelas', payload);
    //         if (response.status) {
    //             toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
    //         }
    //     } catch (error) {
    //         toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus banyak siswa: ${error}`, life: 3000 });
    //     }
    // };
    return {
        addDns,
        getDns,
        searchDns,
        deleteDns,
        updateDns,
        searchDnsLokal
    };
}
