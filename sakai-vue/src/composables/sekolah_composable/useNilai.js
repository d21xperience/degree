import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useNilai() {
    const store = useStore();
    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();

    const fetchNilaiSiswa = async (pesertaDidikId = null) => {
        const payload = {
            // page: 1,
            semesterId: initSelectedSemester.value.semesterId,
            schemaname: schemaname.value
        };
        // console.log(payload);
        if (pesertaDidikId) {
            payload.peserta_didik_id = pesertaDidikId;
        }
        const results = await store.dispatch('nilaiService/fetchNilaiSiswa', payload);
        // console.log(results)
        return results;
        // siswaList.value = results;
        // results.forEach(item => {
        //     siswa.value.push(item)
        // });
    };

    const searchNilai = async (pesertaDidikId) => {
        try {
            // console.log(pesertaDidikId)
            // return
            const payload = {
                semesterId: initSelectedSemester.value.semesterId,
                schemaname: schemaname.value,
                pesertaDidikId: pesertaDidikId
            };
            const results = await store.dispatch('nilaiService/searchNilai', payload);
            if (results.status) {
                return results;
            }
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    return {
        fetchNilaiSiswa,
        searchNilai
    };
}
