import { useToast } from 'primevue/usetoast';
import { computed, ref, watch } from 'vue';
import { useStore } from 'vuex';
export function useSekolahService() {
    const store = useStore();
    const toast = useToast();

    const guruList = ref([]);
    const guruTerdaftarList = ref([]);
    const siswaAktifList = ref([]);
    const kelasList = ref([]);
    const sekolah = computed(() => {
        const tes = store.getters['sekolahService/getSekolah'];
        const response = tes;
        response.uri = response?.sekolah.nama.toLowerCase().replace(/\s+/g, '');
        return response;
    });
    const listTahunAjaran = computed(() => rawlistTahunAjaran || []);
    const rawlistTahunAjaran = ref();
    const listSemester = store.getters['sekolahService/getSemester'];

    const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
    const initSelectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
    const initSelectedTahunAjaran = computed(() => store.getters['sekolahService/getSelectedTahunAjaran']);
    const selectedSemester = ref();
    const selectedTahunAjaran = ref();

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
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Info', detail: response?.message, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Info', detail: error, life: 3000 });
        }
    };
    const fetchGuruTerdaftar = async () => {
        try {
            const payload = {
                tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            let res = await store.getters['sekolahService/getPTKTerdaftar'];
            // console.log('useSekolahService/fetchGuruTerdaftar', res);
            if (!res || res.length == 0) {
                // console.log(payload);
                res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
            } else {
                if (res.tahun_ajaran_id != initSelectedSemester.value?.tahunAjaranId) {
                    res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
                }
            }
            guruTerdaftarList.value = res.ptkTerdaftar;
            return res.ptkTerdaftar;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil guru terdaftar: ${error}`, life: 3000 });
        }
    };
    const searchGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };
            const response = await store.dispatch('sekolahService/searchPTKTerdaftar', payload);
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `Sukses: ${response.message}`, life: 3000 });
            }
            return response.ptkTerdaftar;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil guru terdaftar: ${error}`, life: 3000 });
        }
    };
    const deleteGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };

            // console.log('useSekolahService', payload);
            const response = await store.dispatch('sekolahService/deletePTKTerdaftar', payload);
            return response;
        } catch (error) {
            console.error('Gagal menghapus data guru:', error);
        }
    };
    const deleteBatchGuruTerdaftar = async (ptkTerdaftarId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar_id: ptkTerdaftarId
            };

            // console.log('useSekolahService', payload);
            const response = await store.dispatch('sekolahService/deleteBatchPTKTerdaftar', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus Guru: ${error}`, life: 3000 });
        }
    };
    const updateGuruTerdaftar = async (ptkTerdaftar) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar: [ptkTerdaftar._rawValue]
            };
            const response = await store.dispatch('sekolahService/updatePTKTerdaftar', payload);
            if (!response.status) {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data guru: ${response.message}`, life: 3000 });
            }
            await store.dispatch('sekolahService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
            toast.add({ severity: 'success', summary: 'Success', detail: 'Berhasil update data guru', life: 3000 });
            return response;
        } catch (error) {
            // console.error('Gagal update data guru:', error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data guru: ${error}`, life: 3000 });
        }
    };
    const guruTerdaftarLoading = ref(false);
    const addGuruTerdaftar = async (ptkTerdaftar) => {
        guruTerdaftarLoading.value = true;
        try {
            const payload = {
                schemaname: schemaname.value,
                ptk_terdaftar: ptkTerdaftar._rawValue
            };
            // console.log(payload);
            const response = await store.dispatch('sekolahService/addPTKTerdaftar', payload);
            if (!response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: 'Berhasil menambah data guru', life: 3000 });
            }
            await store.dispatch('sekolahService/fetchPTKTerdaftar', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
            // toast.add({ severity: 'success', summary: 'Success', detail: 'Berhasil menambah data guru', life: 3000 });

            return response;
        } catch (error) {
            console.error('Gagal update data guru:', error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data guru: ${error}`, life: 3000 });
        } finally {
            guruTerdaftarLoading.value = false;
        }
    };

    const fetchGuru = async (ptkId = null) => {
        try {
            const payload = {
                schemaname: schemaname.value
            };

            if (ptkId) {
                payload.ptk_id = ptkId;
            }
            const response = await store.dispatch('sekolahService/fetchGuru', payload);
            guruList.value = response;
        } catch (error) {
            console.error('Gagal mengambil data guru:', error);
        }
    };
    // =================================================
    // KELAS
    // =================================================
    const fetchKelas = async (kelasId = null, tingkatPendidikanId = null) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: initSelectedSemester.value?.semesterId
            };
            if (kelasId) {
                payload.kelas_id = kelasId;
            }
            if (tingkatPendidikanId) {
                payload.tingkat_pendidikan_id = tingkatPendidikanId;
            }
            const response = await store.dispatch('sekolahService/fetchKelas', payload);
            return response;
        } catch (error) {
            throw error;
        }
    };
    const getKelas = async () => {
        try {
            let response = store.getters['sekolahService/getKelas'];
            if (!response || Array.isArray(response.kelas) || response.kelas.length == 0 || initSelectedSemester.value?.semesterId != response?.semesterId) {
                response = await fetchKelas();
            }

            kelasList.value = response.kelas;
            return response;
        } catch (error) {
            throw error;
        }
    };

    const searchKelas = async (kelasId = null) => {
        try {
            let response = store.getters['sekolahService/getKelas'];

            if (!response || response.kelas.length == 0 || initSelectedSemester.value?.semesterId != response?.semesterId) {
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: initSelectedSemester.value?.semesterId
                };
                response = await store.dispatch('sekolahService/fetchKelas', payload);
                if (response.status) {
                    toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
                }
            }
            const result = response.kelas.find((item) => item.rombonganBelajarId == kelasId);
            return result;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil kelas: ${error}`, life: 3000 });
        }
    };
    const addKelas = async (kelas, anggotaKelas = null) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kelas: [kelas._rawValue],
                anggota_kelas: anggotaKelas
            };
            console.log(payload);
            const response = await store.dispatch('sekolahService/createKelas', payload);
            if (!response.status) {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menambah data kelas: ${response.message}`, life: 3000 });
            }
            // await store.dispatch('sekolahService/fetchKelas', { tahunAjaranId: initSelectedSemester.value?.tahunAjaranId, schemaname: schemaname.value });
            toast.add({ severity: 'success', summary: 'Success', detail: 'Berhasil menambah data kelas', life: 3000 });

            return response;
        } catch (error) {
            console.error('Gagal update data kelas:', error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data kelas: ${error}`, life: 3000 });
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
    const fetchSiswaAktif = async (semesterId = null) => {
        try {
            const requestData = {
                schemaname: schemaname.value,
                semesterId: semesterId || initSelectedSemester.value.semesterId
            };
            const cachedData = await store.getters['sekolahService/getSiswaAktif'];
            const shouldFetchNewData = !cachedData || !cachedData?.peserta_didik?.length || cachedData.semester_id !== requestData.semesterId;

            let studentData = cachedData;
            if (shouldFetchNewData) {
                studentData = await store.dispatch('sekolahService/fetchSiswaAktif', requestData);
            }

            // Update reactive data
            siswaAktifList.value = studentData.peserta_didik;

            return studentData.peserta_didik;
        } catch (error) {
            console.error('Failed to fetch active students:', error);
            throw error;
        }
    };

    const deleteSiswaAktif = async (anggotaRombelId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelId
            };
            const response = await store.dispatch('sekolahService/createAnggotaKelas', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus siswa: ${error}`, life: 3000 });
        }
    };
    const deleteBatchSiswaAktif = async (anggotaRombelIds) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelIds
            };
            const response = await store.dispatch('sekolahService/deleteBatchAnggotaKelas', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus banyak siswa: ${error}`, life: 3000 });
        }
    };
    const addSiswaAktif = async (anggotaRombelIds) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelIds
            };
            const response = await store.dispatch('sekolahService/deleteBatchAnggotaKelas', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus banyak siswa: ${error}`, life: 3000 });
        }
    };
    const searchSiswaAktif = async (pesertaDidikId) => {
        try {
            const response = await store.getters['sekolahService/getSiswaAktif'];
            if (response) {
                const siswa = response.peserta_didik.find((item) => item.pesertaDidikId.includes(pesertaDidikId));
                toast.add({ severity: 'success', summary: 'Successful', detail: 'Berhasil mengambil data siswa', life: 3000 });
                return siswa;
            }
        } catch (error) {
            console.log(error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil data siswa: ${error}`, life: 3000 });
        }
    };

    const fetchSemester = async () => {
        try {
            const results = await store.dispatch('sekolahService/fetchSemester');
            if (results.status) {
                return results.semester;
            }
        } catch (error) {
            throw error;
        }
    };

    /**
     * Gets a contract for the given owner address
     * @param {Array} semester
     * @returns {Promise} A promise that resolves with the contract response
     * @throws {Error} If there's an error fetching the contract
     */
    /**
     * Deletes semesters by their IDs
     * @param {Array<Object>} semesters - Array of semester objects containing semesterId
     * @returns {Promise<Object>} A promise that resolves with the deletion response
     * @throws {Error} If the deletion request fails
     *
     * @example
     * await deleteSemester([{ semesterId: 1 }, { semesterId: 2 }]);
     */
    const deleteSemester = async (semesters) => {
        // Validasi input
        if (!Array.isArray(semesters)) {
            throw new Error('Parameter must be an array of semester objects');
        }

        // Ekstrak semesterId secara langsung
        const semesterIds = semesters
            .map((semester) => {
                if (semester.semesterId == null) {
                    console.warn('Semester object missing semesterId:', semester);
                }
                return semester.semesterId;
            })
            .filter((id) => id != null); // Filter null/undefined

        if (semesterIds.length === 0) {
            throw new Error('No valid semester IDs provided');
        }
        console.log(semesterIds);
        // return;
        try {
            // Tunggu hasil dispatch dengan await
            const response = await store.dispatch('sekolahService/deleteSemester', semesterIds);
            console.log('deleteSemester', response);
            // Asumsi response memiliki struktur { status: true, data: ... } atau sejenisnya
            if (response.status) {
                return response;
            } else {
                throw new Error(response?.message || 'Failed to delete semesters');
            }
        } catch (error) {
            // Tambahkan konteks error
            console.error('Error deleting semesters:', error);
            throw new Error(`Failed to delete semesters: ${error.message || 'Unknown error'}`);
        }
    };

    const updateSemester = async (semester) => {
        try {
            console.log('updateSemester', semester);
            // return
            const res = await store.dispatch('sekolahService/updateSemester', semester);
            return res;
        } catch (error) {
            throw error;
        }
    };
    const fetchTahunAjaran = async () => {
        try {
            rawlistTahunAjaran.value = store.getters['sekolahService/getTahunAjaran'];
            if (!rawlistTahunAjaran || rawlistTahunAjaran.value.length == 0) {
                const results = await store.dispatch('sekolahService/fetchTahunAjaran');
                if (results.status) {
                    rawlistTahunAjaran.value = results.tahunAjaran;
                    // toast.add({ severity: 'success', summary: 'Successful', detail: `${results.message}`, life: 3000 });
                }
            }
        } catch (error) {
            throw error;
            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil tahun ajaran: ${error}`, life: 3000 });
        }
    };

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
        const results = await store.dispatch('sekolahService/fetchNilaiSiswa', payload);
        // console.log(results)
        return results;
        // siswaList.value = results;
        // results.forEach(item => {
        //     siswa.value.push(item)
        // });
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
            throw error;
        }
    };

    const fetchBanyakSiswaByTingkatId = async (tingkatPendidikanId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: `${initSelectedTahunAjaran.value?.tahunAjaranId}2`,
                tingkat_pendidikan_id: tingkatPendidikanId
            };
            const res = await store.dispatch('sekolahService/fetchBanyakSiswaByTingkatId', payload);
            // console.log(res);

            return res;
        } catch (error) {
            throw error;
        }
    };
    const fetchBanyakSiswaByRombelId = async (rombelId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: `${initSelectedSemester.value?.semesterId}`,
                rombongan_belajar_id: rombelId
            };
            // return
            const res = await store.dispatch('sekolahService/fetchBanyakSiswaByRombelId', payload);

            return res;
        } catch (error) {
            throw error;
        }
    };

    const addDns = async (dns) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`,
                data_nominasi_sementara: dns
            };
            // console.log(payload);
            // return;
            const response = await store.dispatch('sekolahService/createDns', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
        }
    };
    const updateDns = async (dns) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`,
                data_nominasi_sementara: dns
            };
            const response = await store.dispatch('sekolahService/updateDns', payload);
            if (response.status) {
                const dnsTabel = store.getters['sekolahService/getDns'];
                const dns = dnsTabel.dataNominasiSementara.find((item) => item.pesertaDidikId == payload.data_nominasi_sementara.pesertaDidikId);
                if (dns) {
                    Object.assign(dns, payload.data_nominasi_sementara);
                }
                store.commit('sekolahService/SET_TABELDNS', dnsTabel);
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return true;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
        }
    };
    const getDns = async (tahunAjaranId) => {
        try {
            let response = await store.getters['sekolahService/getDns'];
            // console.log(response)
            if (!response || !Array.isArray(response.dataNominasiSementara) || response.dataNominasiSementara.length === 0 || response.tahun_ajaran_id != tahunAjaranId) {
                const payload = {
                    schemaname: schemaname.value,
                    tahun_ajaran_id: tahunAjaranId,
                    is_complete: false
                };
                response = await store.dispatch('sekolahService/fetchDns', payload);
                console.log(response);
                if (response) {
                    toast.add({ severity: 'success', summary: 'Successful', detail: `"${response.message}"`, life: 3000 });
                    return response.dataNominasiSementara;
                } else {
                    toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: "${response.message}"`, life: 3000 });
                    return [];
                }
            }
            return response.dataNominasiSementara;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
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
            const response = await store.dispatch('sekolahService/searchDns', payload);
            // console.log(response)
            // return
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.dataNominasiSementara;
            } else {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
        }
    };
    const searchDnsLokal = async (pesertaDidikId) => {
        try {
            const dnsTabel = store.getters['sekolahService/getDns'];
            const dns = dnsTabel.dataNominasiSementara.find((item) => item.pesertaDidikId == pesertaDidikId);
            // console.log(dns);
            // return
            if (dns) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${pesertaDidikId}`, life: 3000 });
                return dns;
            } else {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${pesertaDidikId}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
        }
    };

    const deleteDns = async (pesertaDidikId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                peserta_didik_id: pesertaDidikId
            };
            console.log(payload);
            const response = await store.dispatch('sekolahService/deleteDns', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus siswa: ${error}`, life: 3000 });
        }
    };
    const deleteBatchDns = async (anggotaRombelIds) => {
        // console.log(anggotaRombelIds);

        // return;
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelIds
            };
            const response = await store.dispatch('sekolahService/deleteBatchAnggotaKelas', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal menghapus banyak siswa: ${error}`, life: 3000 });
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

    const fetchDashboard = async () => {
        try {
            let response = await store.getters['sekolahService/getDashboard'];
            if (response) {
                const payload = {
                    schemaname: schemaname.value,
                    semester_id: initSelectedSemester.value?.semesterId
                };
                if (payload.schemaname === '') {
                    return;
                }
                response = await store.dispatch('sekolahService/fetchDashboard', payload);
            }
            if (!response) {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi`, life: 3000 });
            }
            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
        }
    };
    const fetchKategoriSekolah = async () => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: initSelectedSemester.value?.tahunAjaranId
            };
            const response = await store.dispatch('sekolahService/fetchKategoriSekolah', payload);
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            kategoriSekolahList.value = response.kategoriSekolah;
        } catch (error) {
            console.log(error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            // console.log(payload)
            // return
            const response = await store.dispatch('sekolahService/createKategoriSekolah', payload);
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
        }
    };
    const updateKategoriSekolah = async (kategoriSekolah) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                kategori_sekolah: kategoriSekolah
            };

            // console.log(payload);
            // return;
            const response = await store.dispatch('sekolahService/updateKategoriSekolah', payload);
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
        }
    };
    const createProsesKategoriSekolahDanKelas = async () => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`
            };
            const response = await store.dispatch('sekolahService/createProsesKelas', payload);
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }

            return response.kategoriMapel;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            // return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
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
            // console.log(response);
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Success', detail: `${response.message}`, life: 3000 });
            }
            // return response.kategoriSekolah;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
        }
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

    watch(selectedTahunAjaran, (e) => {
        store.commit('sekolahService/SET_SELECTEDTAHUNAJARAN', e);
    });
    watch(selectedSemester, (e) => {
        store.commit('sekolahService/SET_SELECTEDSEMESTER', e);
    });

    return {
        fetchGuru,
        guruList,
        fetchGuruTerdaftar,
        searchGuruTerdaftar,
        deleteGuruTerdaftar,
        updateGuruTerdaftar,
        guruTerdaftarList,
        fetchKelas,
        getKelas,
        // fetchSiswa,
        kelasList,
        fetchSemester,
        deleteSemester,
        updateSemester,
        fetchSiswaAktif,
        deleteSiswaAktif,
        deleteBatchSiswaAktif,
        fetchNilaiSiswa,
        fetchTingkat,
        schemaname,
        selectedSemester,
        initSelectedSemester,
        listTahunAjaran,
        listSemester,
        selectedTahunAjaran,
        initSelectedTahunAjaran,
        fetchBanyakSiswaByTingkatId,
        fetchBanyakSiswaByRombelId,
        addDns,
        getDns,
        fetchSekolah,
        fetchDashboard,
        searchDns,
        createInfoIjazah,
        sekolah,
        deleteDns,
        fetchTahunAjaran,
        deleteBatchGuruTerdaftar,
        fetchKategoriSekolah,
        deleteKategoriSekolah,
        createKategoriSekolah,
        updateKategoriSekolah,
        addGuruTerdaftar,
        addKelas,
        updateSekolah,
        createProsesKategoriSekolahDanKelas,
        fetchKategoriMapel,
        deleteKategoriMapel,
        deleteBatchKategoriMapel,
        kategoriSekolahTabel,
        kategoriSekolahList,
        deleteKategoriSekolahKurikulum,
        updateDns,
        searchDnsLokal,
        searchSiswaAktif,
        searchKelas,
        guruTerdaftarLoading,
        fetchAnggotaKelas
    };
}
