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
    const sekolah = computed(() => store.getters['sekolahService/getSekolah']);
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
                // response = await store.dispatch('sekolahService/fetchTabeltenant', response?.user.sekolahTenantId);
            }
            // console.log(response)
            return response;
        } catch (error) {
            console.log(error);
        }
    };
    const fetchGuruTerdaftar = async () => {
        try {
            const payload = {
                tahunAjaranId: initSelectedSemester.value?.tahunAjaranId,
                schemaname: schemaname.value
            };
            console.log(payload);
            let res = await store.getters['sekolahService/getPTKTerdaftar'];
            if (!res) {
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
            console.log(response)
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
            console.log(payload)
            const response = await store.dispatch('sekolahService/updatePTKTerdaftar', payload);
            if (!response.status) {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data guru: ${response.message}`, life: 3000 });
            }
            toast.add({ severity: 'success', summary: 'Success', detail: `Berhasil update data guru: ${response.message}`, life: 3000 });

            return response;
        } catch (error) {
            console.error('Gagal update data guru:', error);
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal update data guru: ${error}`, life: 3000 });
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

    const fetchKelas = async (kelasId = null, tingkatPendidikanId = null) => {
        try {
            let response = store.getters['sekolahService/getKelas'];
            if (!response || response.length == 0) {
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
                response = await store.dispatch('sekolahService/fetchRombel', payload);
            }

            kelasList.value = response.kelas;
            // console.log(response);
            return response.kelas;
        } catch (error) {
            console.error('Gagal mengambil data kelas:', error);
        }
    };
    // const fetchAnggotaKelas = async (anggotaRombelId = null) => {
    //     try {
    //         const payload = {
    //             schemaname: schemaname.value,
    //             semester_id: initSelectedSemester.value?.semesterId
    //         };

    //         if (anggotaRombelId) {
    //             payload.anggota_rombel_id = anggotaRombelId;
    //         }
    //         const response = await store.dispatch('sekolahService/fetchSiswaAktif', payload);
    //         return response;
    //     } catch (error) {
    //         console.error('Gagal mengambil data kelas:', error);
    //     }
    // };
    const fetchSiswaAktif = async (semesterId = null) => {
        const payload = {
            schemaname: schemaname.value
        };
        if (!semesterId) {
            payload.semesterId = initSelectedSemester.value.semesterId;
        } else {
            payload.semesterId = semesterId;
        }
        let res = await store.getters['sekolahService/getSiswaAktif'];

        // console.log(res.semester_id != payload.semesterId);
        if (!res) {
            res = await store.dispatch('sekolahService/fetchSiswaAktif', payload);
        } else {
            if (res.semester_id != payload.semesterId) {
                res = await store.dispatch('sekolahService/fetchSiswaAktif', payload);
            }
        }
        siswaAktifList.value = res.peserta_didik;
        return res.peserta_didik;
    };

    const deleteSiswaAktif = async (anggotaRombelId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                anggota_rombel_id: anggotaRombelId
            };
            const response = await store.dispatch('sekolahService/deleteAnggotaKelas', payload);
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
    const searchSiswaById = async (pesertaDidikId) => {
        try {
            const response = await store.push('sekolahService');
        } catch (error) {}
    };

    const fetchSemester = async () => {
        try {
            const results = await store.dispatch('sekolahService/fetchSemester');
            if (results) {
                semester.value = store.getters['sekolahService/getSemester'];
                // Cek apakah di vuex ada nilai
                initSelectedSemester.value = await store.getters['sekolahService/getinitSelectedSemester'];
                if (initSelectedSemester.value == null) {
                    // jika tidak ada, ambil semester terbaru berdasarkan ID terbesar
                    initSelectedSemester.value = semester.value.reduce((latest, current) => (current.semesterId > latest.semesterId ? current : latest));
                }
                // tetapkan semester yang dipilih
                store.commit('sekolahService/SET_SELECTEDSEMESTER', initSelectedSemester.value);
            }
        } catch (error) {}
    };

    const fetchTahunAjaran = async () => {
        try {
            rawlistTahunAjaran.value = store.getters['sekolahService/getTahunAjaran'];
            // console.log(rawlistTahunAjaran.value);
            if (!rawlistTahunAjaran || rawlistTahunAjaran.value.length == 0) {
                const results = await store.dispatch('sekolahService/fetchTahunAjaran');
                // console.log(results);
                if (results.status) {
                    rawlistTahunAjaran.value = results.tahunAjaran;
                    toast.add({ severity: 'success', summary: 'Successful', detail: `${results.message}`, life: 3000 });
                }
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengambil tahun ajaran: ${error}`, life: 3000 });
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
            // console.log(response);
            if (!response) {
                const payload = {
                    jenjang_pendidikan_id: await store.getters['sekolahService/getSekolah']?.jenjangPendidikanId //sekolah.value?.jenjangPendidikanId
                };
                response = await store.dispatch('sekolahService/fetchTingkatPendidikan', payload);
            }
            return response;
        } catch (error) {
            throw error;
        }

        // console.log(response);
    };

    const fetchBanyakSiswaByTingkatId = async (tingkatPendidikanId) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                semester_id: `${initSelectedTahunAjaran.value?.tahunAjaranId}2`,
                tingkat_pendidikan_id: tingkatPendidikanId
            };
            const res = await store.dispatch('sekolahService/fetchBanyakSiswaByTingkatId', payload);
            console.log(res);

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

    const createDns = async (dns) => {
        try {
            const payload = {
                schemaname: schemaname.value,
                tahun_ajaran_id: `${initSelectedSemester.value?.tahunAjaranId}`,
                data_nominasi_sementara: dns
            };
            const response = await store.dispatch('sekolahService/createDns', payload);
            if (response.status) {
                toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mengirimkan DNS: ${error}`, life: 3000 });
        }
    };
    const getDns = async (tahunAjaranId) => {
        try {
            let response = await store.getters['sekolahService/getDns'];
            if (!response || response.length == 0 || response.tahun_ajaran_id != tahunAjaranId) {
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
                tahun_ajaran_id: initSelectedTahunAjaran.value?.tahunAjaranId,
                peserta_didik_id: pesertaDidikId
            };
            const response = await store.dispatch('sekolahService/searchDns', payload);
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
                response = await store.dispatch('sekolahService/fetchDashboard', payload);
            }
            // console.log(!response);
            if (!response) {
                toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi`, life: 3000 });
            }
            // console.log(response);
            // toast.add({ severity: 'success', summary: 'Success', detail: `Berhasil mendapatkan informasi`, life: 3000 });

            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error}`, life: 3000 });
        }
    };

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
        // fetchSiswa,
        kelasList,
        fetchSemester,
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
        createDns,
        getDns,
        fetchSekolah,
        fetchDashboard,
        searchDns,
        createInfoIjazah,
        sekolah,
        deleteDns,
        fetchTahunAjaran,
        deleteBatchGuruTerdaftar
    };
}
