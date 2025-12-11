import { computed, readonly, ref, watch } from 'vue';
import { useStore } from 'vuex';

/**
 *
 * @param {String} schemaname Wajib - schemaname tidak boleh kosong
 * @returns {Object} API
 */
export function useKategoriSekolah(schemaname = '') {
    if (schemaname == '') {
        return new Error('schemaname & selectedTahunAjaran tidak boleh kosong');
    }

    const store = useStore();
    const isFetching = ref(false);
    // const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname || null);
    // const selectedTahunAjaran = computed(() => store.getters['semesterService/getSelectedTahunAjaran']);
    // state
    const isKategoriSekolahCompleted = computed(() => store.getters['sekolahService/getIsKategoriSekolahCompleted']);
    const storeKategoriSekolahList = computed(() => store.getters['sekolahService/getTabelKategoriSekolah']);
    const kategoriSekolahList = ref();
    // const kategoriSekolahTabel = ref([]);
    const storeKategoriSekolahTabel = computed(() => {
        const kategoriSekolah = storeKategoriSekolahList.value.kategoriSekolah;
        if (!kategoriSekolah) return;
        const results = kategoriSekolah.reduce((acc, tes) => {
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

    const convertToSekolahTabel = (kategoriSekolah) => {
        const results = kategoriSekolah.reduce((acc, tes) => {
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
    };
    watch(storeKategoriSekolahList, (newVal) => {
        if (newVal == kategoriSekolahList.value) {
            return;
        } else {
            kategoriSekolahList.value = newVal;
        }
        // kategoriSekolahTabel.value = createKategoriSekolahTabel();
    });
    /**
     *
     * @param {String} tahunAjaran Wajib
     * @returns {Object}
     */
    const fetchKategoriSekolah = async (tahunAjaranId = '') => {
        isFetching.value = true;
        // console.log('isFething....');
        try {
            const cached = store.getters['sekolahService/getTabelKategoriSekolah'];

            if (cached.length > 0) {
                console.log('dieksesi blok 1');
                if (cached?.tahunAjaranId?.length > 0) {
                    console.log('dieksesi blok 2');
                    if (cached.tahunAjaranId == tahunAjaranId) {
                        console.log('dieksesi blok 2.1');
                        return cached;
                    }
                }
            }
            // }
            // console.log('dieksesi blok 3');
            // console.log(selectedTahunAjaran.value.tahunAjaranId);
            if (tahunAjaranId == '') {
                throw new Error('Silahkan pilih tahun ajaran terlebih dahulu');
                // return;
            }
            // if (!tahunAjaran) tahunAjaran = selectedTahunAjaran.value.tahunAjaranId;
            const payload = {
                schemaname: schemaname,
                tahun_ajaran_id: tahunAjaranId
            };
            // console.log(payload);
            // console.log(payload);
            const response = await store.dispatch('sekolahService/fetchKategoriSekolah', payload);
            // kategoriSekolahList.value = response;
            // console.log(response);
            return response;
        } catch (error) {
            throw new Error(`Gagal mendapatkan Kategori Sekolah:'${error.message || error}`);
        } finally {
            isFetching.value = false;
        }
    };

    /**
     *
     * @param {Array} kategoriSekolah Wajib diisi
     * @returns {Object}
     */
    const createKategoriSekolah = async (kategoriSekolah = []) => {
        try {
            isFetching.value = true;
            console.log(kategoriSekolah);
            return;
            const payload = {
                schemaname: schemaname,
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
            console.log(payload);
            alert('berhasil ditambahkan!!!');
            return;
            const response = await store.dispatch('sekolahService/createKategoriSekolah', payload);
            return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal membuat Kategori Sekolah:'${error.message || error}`);
        } finally {
            isFetching.value = false;
        }
    };
    const updateKategoriSekolah = async (kategoriSekolah) => {
        try {
            const payload = {
                schemaname: schemaname,
                kategori_sekolah: kategoriSekolah
            };
            const response = await store.dispatch('sekolahService/updateKategoriSekolah', payload);
            return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal update Kategori Sekolah:'${error.message || error}`);
        }
    };
    const deleteKategoriSekolah = async (kategoriSekolahId) => {
        try {
            const payload = {
                schemaname: schemaname,
                kategori_sekolah_id: kategoriSekolahId
            };
            console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriSekolah', payload);

            return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Sekolah:'${error.message || error}`);
        }
    };
    const deleteKategoriSekolahKurikulum = async (kurikulumId) => {
        try {
            const payload = {
                schemaname: schemaname,
                kurikulum_id: kurikulumId
            };
            // console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriSekolahKurikulum', payload);
            return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Kurikulum Sekolah:'${error.message || error}`);
        }
    };
    const createProsesKategoriSekolahDanKelas = async (selectedTahunAjaran) => {
        try {
            const payload = {
                schemaname: schemaname,
                tahun_ajaran_id: selectedTahunAjaran.value.label
            };
            const response = await store.dispatch('sekolahService/createProsesKelas', payload);
            console.log(response);
            return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal membuat Kategori Sekolah:'${error.message || error}`);
        }
    };

    const fetchKategoriMapel = async (mapel) => {
        try {
            const payload = {
                schemaname: schemaname,
                // tahunAjaranId: selectedSemester.value?.tahunAjaranId,
                kurikulumId: mapel.kurikulumId,
                tingkatPendidikan: mapel.tingkatPendidikan
            };
            const response = await store.dispatch('sekolahService/fetchKategoriMapel', payload);
            return response.kategoriMapel;
        } catch (error) {
            throw new Error(`Gagal mendapatkan Kategori Mapel:'${error.message || error}`);
        }
    };

    const deleteKategoriMapel = async (idMapel) => {
        try {
            const payload = {
                schemaname: schemaname,
                id: idMapel
            };
            // console.log(payload);
            const response = await store.dispatch('sekolahService/deleteKategoriMapel', payload);
            // console.log(response);
            return response;
            // return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Mapel:'${error.message || error}`);
        }
    };
    const deleteBatchKategoriMapel = async (idMapel) => {
        try {
            // console.log(idMapel)
            const payload = {
                schemaname: schemaname,
                id: idMapel
            };
            // console.log(payload)
            // return
            const response = await store.dispatch('sekolahService/deleteBatchKategoriMapel', payload);
            return response;
            // return response.kategoriSekolah;
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Mapel:'${error.message || error}`);
        }
    };

    const createKategoriSekolahId = (existingData = []) => {
        // 1. Validasi input dan handle array kosong
        if (!Array.isArray(existingData) || existingData.length === 0) {
            return 1; // Jika array kosong, mulai dari 1
        }

        // 2. Ambil semua kategori_sekolah_id yang valid
        const allIds = existingData
            .map((item) => item.kategori_sekolah_id)
            .filter((id) => {
                // Filter hanya number yang valid
                const num = Number(id);
                return !isNaN(num) && isFinite(num) && Number.isInteger(num) && num > 0;
            });

        // 3. Jika tidak ada ID yang valid, return 1
        if (allIds.length === 0) {
            return 1;
        }

        // 4. Cari nilai tertinggi
        const maxId = Math.max(...allIds);

        // 5. Return nilai tertinggi + 1
        return maxId + 1;
    };

    // Fungsi untuk validasi kurikulum_id duplikat
    const isKurikulumIdDuplicate = (existingData, newKurikulumId) => {
        if (!Array.isArray(existingData) || existingData.length === 0) {
            return false; // Tidak ada data, berarti tidak duplikat
        }

        // Cek apakah kurikulum_id sudah ada
        return existingData.some((item) => item.kurikulum_id === newKurikulumId);
    };
    // Fungsi untuk validasi kurikulum_id duplikat
    const isTingkatIdDuplicate = (existingData, newTingkatId) => {
        if (!Array.isArray(existingData) || existingData.length === 0) {
            return false; // Tidak ada data, berarti tidak duplikat
        }

        // Cek apakah kurikulum_id sudah ada
        return existingData.some((item) => item.tingkat_id === newTingkatId);
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
        storeKategoriSekolahTabel,
        kategoriSekolahList,
        deleteKategoriSekolahKurikulum,
        isFetching: readonly(isFetching),
        isKategoriSekolahCompleted: readonly(isKategoriSekolahCompleted),
        createKategoriSekolahId,
        isKurikulumIdDuplicate,
        isTingkatIdDuplicate,
        convertToSekolahTabel
    };
}
