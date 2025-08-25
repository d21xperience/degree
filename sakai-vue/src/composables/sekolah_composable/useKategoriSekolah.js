import { computed, ref } from 'vue';
import { useStore } from 'vuex';
import { useSemester } from './useSemester';
import { useTableTenant } from './useTableTenant';
export function useKategoriSekolah() {
    const store = useStore();

    const { schemaname } = useTableTenant();
    const { initSelectedSemester } = useSemester();
    const fetchKategoriSekolah = async () => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: initSelectedSemester.value?.tahunAjaranId
            };
            const response = await store.dispatch('sekolahService/fetchKategoriSekolah', payload);
            kategoriSekolahList.value = response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan Kategori Sekolah:', error);
        }
    };
    const kategoriSekolahList = ref([]);
    const kategoriSekolahTabel = computed(() => {
        const results = kategoriSekolahList.value.reduce((acc, tes) => {
            const existing = acc.find((item) => item.kurikulum_id === tes.kurikulum_id);

            if (existing) {
                existing.kategorikelas.push({
                    kategori_sekolah_id: tes.kategori_sekolah_id,
                    tingkat_id: tes.tingkat_id,
                    jumlah: tes.jumlah
                });
            } else {
                acc.push({
                    kategori_sekolah_id: tes.kategori_sekolah_id,
                    kurikulum_id: tes.kurikulum_id,
                    jurusan_id: tes.jurusan_id,
                    nama_kurikulum: tes.nama_kurikulum,
                    nama_bidang_keahlian: tes.nama_bidang_keahlian,
                    nama_program_keahlian: tes.nama_program_keahlian,
                    nama_jurusan: tes.nama_jurusan,
                    jenjang_pendidikan_id: tes.jenjang_pendidikan_id,
                    tahun_ajaran_id: tes.tahun_ajaran_id,
                    kategorikelas: [
                        {
                            kategori_sekolah_id: tes.kategori_sekolah_id,
                            tingkat_id: tes.tingkat_id,
                            jumlah: tes.jumlah
                        }
                    ]
                });
            }

            return acc;
        }, []);

        // Setelah grouping selesai, hitung total_kelas untuk tiap kurikulum
        const finalResults = results.map((item) => {
            const total_kelas = item.kategorikelas.reduce((sum, kls) => sum + kls.jumlah, 0);
            return {
                ...item,
                total_kelas
            };
        });

        // console.log(finalResults);
        return finalResults;
    });

    const createKategoriSekolah = async (kategoriSekolah) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kategori_sekolah: {
                    kurikulum_id: kategoriSekolah.kurikulum_id,
                    jurusan_id: kategoriSekolah.jurusan_id,
                    nama_kurikulum: kategoriSekolah.nama_kurikulum,
                    nama_bidang_keahlian: kategoriSekolah.nama_bidang_keahlian,
                    nama_program_keahlian: kategoriSekolah.nama_program_keahlian,
                    nama_jurusan: kategoriSekolah.nama_jurusan,
                    jenjang_pendidikan_id: kategoriSekolah.jenjang_pendidikan_id,
                    tingkat_id: kategoriSekolah.tingkat_id,
                    jumlah: kategoriSekolah.jumlah,
                    tahun_ajaran_id: `${kategoriSekolah.tahun_ajaran_id}`
                }
            };
            const response = await store.dispatch('sekolahService/createKategoriSekolah', payload);
            return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal membuat Kategori Sekolah:', error);
        }
    };
    const updateKategoriSekolah = async (kategoriSekolah) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kategori_sekolah: kategoriSekolah
            };
            const response = await store.dispatch('sekolahService/updateKategoriSekolah', payload);
            return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal update Kategori Sekolah:', error);
        }
    };
    const deleteKategoriSekolah = async (kategoriSekolahId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kategori_sekolah_id: kategoriSekolahId
            };
            console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriSekolah', payload);

            return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Sekolah:', error);
        }
    };
    const deleteKategoriSekolahKurikulum = async (kurikulumId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kurikulum_id: kurikulumId
            };
            // console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriSekolahKurikulum', payload);
            return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Kurikulum Sekolah:', error);
        }
    };
    const createProsesKategoriSekolahDanKelas = async () => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`
            };
            const response = await store.dispatch('sekolahService/createProsesKelas', payload);
            return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal membuat Kategori Sekolah:', error);
        }
    };

    const fetchKategoriMapel = async (mapel) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                // tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                kurikulumId: mapel.kurikulumId,
                tingkatPendidikan: mapel.tingkatPendidikan
            };
            const response = await store.dispatch('sekolahService/fetchKategoriMapel', payload);
            return response.kategoriMapel;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan Kategori Mapel:', error);
        }
    };

    const deleteKategoriMapel = async (idMapel) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                id: idMapel
            };
            // console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriMapel', payload);
            // console.log(response);
            return response;
            // return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    };
    const deleteBatchKategoriMapel = async (idMapel) => {
        try {
            // console.log(idMapel)
            const payload = {
                schemaname: schemaname.value,
                id: idMapel
            };
            // console.log(payload)
            // return
            const response = await store.dispatch('sekolahService/deleteBatchKategoriMapel', payload);
            return response;
            // return response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    };

    return {
        fetchKategoriSekolah,
        deleteKategoriSekolah,
        createKategoriSekolah,
        updateKategoriSekolah,
        createProsesKategoriSekolahDanKelas,
        fetchKategoriMapel,
        deleteKategoriMapel,
        deleteBatchKategoriMapel,
        kategoriSekolahTabel,
        kategoriSekolahList,
        deleteKategoriSekolahKurikulum
    };
}
