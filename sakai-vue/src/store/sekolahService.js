import axios from 'axios';
const api = axios.create({
    baseURL: import.meta.env.VITE_API_SEKOLAH_BASE_URL, // 'http://localhost:8183/api/v1',
    withCredentials: true, // Untuk mengirim cookie atau credensial
    headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer your_token'
    },
    timeout: 20000 // 20 detik
});

const state = {
    // tabelKelas: JSON.parse(localStorage.getItem("tabelKelas")) || null,
    tabelKelas: (() => {
        try {
            const data = localStorage.getItem('tabelKelas');
            return data ? JSON.parse(data) : [];
        } catch (e) {
            console.error('Gagal parse tabelKelas dari localStorage', e);
            return [];
        }
    })(),
    tabelTenant: JSON.parse(localStorage.getItem('tabelTenant')) || null,
    tabelSemester: JSON.parse(localStorage.getItem('tabelSemester')) || null,
    tabelSekolah: JSON.parse(localStorage.getItem('tabelSekolah')) || null,
    tabelSiswaAktif: JSON.parse(localStorage.getItem('tabelSiswaAktif')) || null,
    tabelGuru: JSON.parse(localStorage.getItem('tabelGuru')) || null,
    tabelPTKTerdaftar: JSON.parse(localStorage.getItem('tabelPTKTerdaftar')) || null,
    tabelTingkatPendidikan: JSON.parse(localStorage.getItem('tabelTingkatPendidikan')) || null,
    tabelKurikulum: JSON.parse(localStorage.getItem('tabelKurikulum')) || null,
    selectedSemester: JSON.parse(localStorage.getItem('selectedSemester')) || null,
    tabelAnggotaKelas: JSON.parse(localStorage.getItem('tabelAnggotaKelas')) || [],
    tabelMapel: JSON.parse(localStorage.getItem('tabelMapel')) || null,

    tabelNilaiakhir: JSON.parse(localStorage.getItem('tabelNilaiakhir')) || null,
    tabelTahunAjaran: JSON.parse(localStorage.getItem('tabelTahunAjaran')) || [],
    selectedTahunAjaran: JSON.parse(localStorage.getItem('selectedTahunAjaran')) || [],
    tabelGelarAkademik: JSON.parse(localStorage.getItem('gelarAkademik')) || [],
    tabelDashboard: JSON.parse(localStorage.getItem('tabelDashboard')) || [],
    tabelDns: JSON.parse(localStorage.getItem('tabelDns')) || []
};

const mutations = {
    SET_TABELTENANT(state, tabelTenant) {
        state.tabelTenant = tabelTenant;
        localStorage.setItem('tabelTenant', JSON.stringify(tabelTenant));
    },
    SET_TABELSEMESTER(state, value) {
        state.tabelSemester = value;
        localStorage.setItem('tabelSemester', JSON.stringify(value));
    },
    SET_TABELSEKOLAH(state, tabelSekolah) {
        state.tabelSekolah = tabelSekolah;
        localStorage.setItem('tabelSekolah', JSON.stringify(tabelSekolah));
    },
    SET_TABELSISWAAKTIF(state, tabelSiswaAktif) {
        state.tabelSiswaAktif = tabelSiswaAktif;
        localStorage.setItem('tabelSiswaAktif', JSON.stringify(tabelSiswaAktif));
    },
    SET_TABELGURU(state, tabelGuru) {
        state.tabelGuru = tabelGuru;
        localStorage.setItem('tabelGuru', JSON.stringify(tabelGuru));
    },
    SET_TABELPTKTERDAFTAR(state, value) {
        state.tabelPTKTerdaftar = value;
        localStorage.setItem('tabelPTKTerdaftar', JSON.stringify(value));
    },
    SET_TABELTINGKATPENDIDIKAN(state, value) {
        state.tabelTingkatPendidikan = value;
        localStorage.setItem('tabelTingkatPendidikan', JSON.stringify(value));
    },
    SET_TABELKURIKULUM(state, value) {
        state.tabelKurikulum = value;
        localStorage.setItem('tabelKurikulum', JSON.stringify(value));
    },
    SET_SELECTEDSEMESTER(state, value) {
        state.selectedSemester = value;
        localStorage.setItem('selectedSemester', JSON.stringify(value));
    },
    SET_TABELKELAS(state, value) {
        state.tabelMapel = value;
        localStorage.setItem('tabelKelas', JSON.stringify(value));
    },
    SET_TABELANGGOTAKELAS(state, value) {
        state.tabelAnggotaKelas = value;
        localStorage.setItem('tabelAnggotaKelas', JSON.stringify(value));
    },
    SET_TABELMAPEL(state, value) {
        state.tabelMapel = value;
        localStorage.setItem('tabelMapel', JSON.stringify(value));
    },
    SET_TABELNILAIAKHIR(state, value) {
        state.tabelNilaiakhir = value;
        localStorage.setItem('tabelNilaiakhir', JSON.stringify(value));
    },
    SET_TABELTAHUNAJARAN(state, value) {
        state.tabelTahunAjaran = value;
        localStorage.setItem('tabelTahunAjaran', JSON.stringify(value));
    },
    SET_SELECTEDTAHUNAJARAN(state, value) {
        state.selectedTahunAjaran = value;
        localStorage.setItem('selectedTahunAjaran', JSON.stringify(value));
    },
    SET_GELARAKADEMIK(state, value) {
        state.tabelGelarAkademik = value;
        localStorage.setItem('gelarAkademik', JSON.stringify(value));
    },
    SET_DASHBOARD(state, value) {
        state.tabelDashboard = value;
        localStorage.setItem('tabelDashboard', JSON.stringify(value));
    },
    SET_TABELDNS(state, value) {
        state.tabelDns = value;
        localStorage.setItem('tabelDns', JSON.stringify(value));
    }
};

const actions = {
    // ================================================
    // Semester
    // ================================================
    async fetchSemester({ commit }, semester_id) {
        try {
            const response = await api.get(`/ss/semester`, {
                params: {
                    semester_id: semester_id
                }
            });
            commit('SET_TABELSEMESTER', response.data.semester);
            const tahunAjaran = state.selectedTahunAjaran?.tahunAjaranId;
            // const selectedSemester = response.data.semester.reduce((max, item) => (item.semesterId > max.semesterId ? item : max), response.data.semester[0]);
            if (!state.selectedSemester) {
                const selectedSemester = response.data.semester.filter((item) => item.tahunAjaranId == tahunAjaran);
                commit('SET_SELECTEDSEMESTER', selectedSemester[0]);
            }
            return true;
        } catch (error) {
            throw error;
        }
    },

    async fetchSelectedSemester({ commit }, payload) {
        commit('SET_SELECTEDSEMESTER', payload);
    },
    // ================================================
    // Tahun Ajaran
    // ================================================
    async fetchTahunAjaran({ commit }, tahun_ajaran_id) {
        try {
            const response = await api.get(`/ss/tahun-ajaran`, {
                params: {
                    tahun_ajaran_id: tahun_ajaran_id
                }
            });
            commit('SET_TABELTAHUNAJARAN', response.data.tahunAjaran);
            const selectedTahunAjaran = response.data.tahunAjaran.reduce((max, item) => (item.tahunAjaranId > max.tahunAjaranId ? item : max), response.data.tahunAjaran[0]);
            commit('SET_SELECTEDTAHUNAJARAN', selectedTahunAjaran);
            return true; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },

    async fetchSelectedTahunAjaran({ commit }, payload) {
        commit('SET_SELECTEDTAHUNAJARAN', payload);
    },
    // ================================================

    async fetchRombel({ commit }, payload) {
        try {
            // console.log(payload);
            const response = await api.get(`/ss/${payload.schemaname}/kelas`, {
                params: {
                    semester_id: payload.semester_id,
                    kelas_id: payload.kelas_id,
                    tingkat_pendidikan_id: payload.tingkat_pendidikan_id
                }
            });
            const data = {
                semesterId: payload.semester_id,
                kelas: response.data.kelas
            };
            commit('SET_TABELKELAS', data.kelas);
            return data; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },

    // ================================================
    // Service Guru
    // ================================================
    async fetchGuru({ commit }, payload) {
        try {
            const response = await api.get('/ss/ptk', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_id: payload.ptk_id
                }
            });
            // console.log(response.data);
            // commit("SET_TABELGURU", response.data.PTK);
            return response.data.PTK; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },
    async fetchPTKTerdaftar({ commit }, payload) {
        try {
            const response = await api.get('/ss/ptk-terdaftar', {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahunAjaranId
                }
            });
            const results = {
                tahun_ajaran_id: payload.tahunAjaranId,
                ptkTerdaftar: response.data.ptkTerdaftar
            };
            // console.log(response.data);
            commit('SET_TABELPTKTERDAFTAR', results);
            return results; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },
    async deletePTKTerdaftar({ commit }, payload) {
        try {
            const response = await api.delete('/ss/ptk-terdaftar/delete', {
                params: {
                    schemaname: payload.schemaname,
                    ptk_terdaftar_id: payload.ptk_terdaftar_id
                }
            });
            // update tabelPtkTerdaftarId
            if (response.status) {
                console.log(state.tabelPTKTerdaftar);
                const updatePTK = {
                    tahun_ajaran_id: state.selectedSemester?.semesterId,
                    ptkTerdaftar: state.tabelPTKTerdaftar.ptkTerdaftar.filter((item) => item.ptkTerdaftarId != payload.ptk_terdaftar_id)
                };
                console.log(updatePTK);
                commit('SET_TABELPTKTERDAFTAR', updatePTK);
            }
            return response;
        } catch (error) {
            throw error;
        }
    },

    async saveGuru({ commit }, payload) {
        try {
            const response = await api.post('/ss/ptk/create', payload);
            return response.data.status;
        } catch (error) {
            console.log(error);
        }
    },
    async savePTKTerdaftar({ commit }, payload) {
        try {
            console.log(payload);
            const response = await api.post(`/ss/${payload.schemaname}/ptk-terdaftar/create`, payload);
            return response.data.status;
        } catch (error) {
            throw error;
        }
    },
    async updatePTKTerdaftar({ commit }, payload) {
        try {
            console.log(payload);
            // return;
            const response = await api.put(`/ss/ptk-terdaftar/update`, payload);
            // console.log(response);
            return response.data;
        } catch (error) {
            throw error;
        }
    },

    // =============================================

    async getTemplate({ commit }, payload) {
        commit('SET_LOADING', true);
        commit('SET_ERROR', null);
        try {
            const response = await api.get(`/ss/download/template`, {
                params: {
                    'template-type': 'siswa'
                }
            });
            // console.log(response.data.semester);
            // commit("SET_TABELSEMESTER", response.data.semester);
            return response; // Mengembalikan data sekolah
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat semester:', error);
            return null;
        }
    },
    async fetchSekolah({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/sekolah`);
            commit('SET_TABELSEKOLAH', response.data.sekolah);
            return response.data.sekolah; // Mengembalikan data sekolah
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat semester:', error);
            return null;
        }
    },
    async updateSekolah({ commit }, payload) {
        try {
            console.log(payload);
            const response = await api.post(`/ss/${payload.schemaname}/update`, payload);
            // console.log(response.data);
            // commit("SET_TABELSEMESTER", response.data.semester);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat semester:', error);
            return null;
        }
    },

    // ==================================
    // SISWA
    // ==================================
    async createBanyakSiswa({ commit }, payload) {
        try {
            const response = await api.post(`/ss/${payload.schemaname}/siswa/create-banyak`, payload);
            console.log(response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            return null;
        }
    },
    async fetchSiswa({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/siswa`, {
                params: {
                    page: payload.page,
                    perpage: payload.perpage
                    // schemaname: schemaname,
                }
            });
            // console.log(response);

            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            throw error;
        }
    },
    async fetchSiswaAktif({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas`, {
                params: {
                    semester_id: payload.semesterId,
                    schemaname: payload.schemaname
                }
            });
            const results = {
                semester_id: payload.semesterId,
                peserta_didik: response.data.anggotaKelas
            };
            commit('SET_TABELSISWAAKTIF', results);
            return results;
        } catch (error) {
            throw error;
        }
    },

    // async searchSiswa({ commit }, payload) {
    //     try {
    //         const response = await api.get(`/ss/${payload.schemaname}/siswa/search`, {
    //             params: {
    //                 nm_siswa: payload.nm_siswa
    //             }
    //         });
    //         // console.log("results",response.data);
    //         return response.data;
    //     } catch (error) {
    //         commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
    //         throw error;
    //     }
    // },
    async searchSiswaAktifById({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/siswa/search`, {
                params: {
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            // console.log("results",response.data);
            return response.data;
        } catch (error) {
            throw error;
        }
    },

    async fetchBanyakSiswaByTingkatId({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/filter`, {
                params: {
                    semester_id: payload.semester_id,
                    tingkat_pendidikan_id: payload.tingkat_pendidikan_id
                }
            });
            console.log(response);
            return response.data.anggotaKelas;
        } catch (error) {
            throw error;
        }
    },
    async fetchBanyakSiswaByRombelId({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/filter`, {
                params: {
                    semester_id: payload.semester_id,
                    rombongan_belajar_id: payload.rombongan_belajar_id
                }
            });
            console.log(response);
            return response.data.anggotaKelas;
        } catch (error) {
            throw error;
        }
    },
    // =======================================
    // ANGGOTA KELAS
    // =======================================
    async searchAnggotaKelas({ commit }, payload) {
        try {
            // console.log(payload);
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas/search`, {
                params: {
                    semester_id: payload.semester_id,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            return response.data.anggotaKelas; // Mengembalikan data sekolah
        } catch (error) {
            if (error.code === 'ECONNABORTED') {
                console.error('Permintaan terlalu lama, server lambat atau tidak merespons.');
            } else {
                console.error('Terjadi kesalahan:', error.message);
            }
            return null;
        }
    },
    async fetchAnggotaKelas({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/anggota-kelas`, {
                params: {
                    semester_id: payload.semester_id,
                    rombongan_belajar_id: payload.rombongan_belajar_id
                }
            });
            return response.data.anggotaKelas;
        } catch (error) {
            console.error('Gagal mendapatkan anggota kelas:', error);
            return null;
        } finally {
        }
    },

    async deleteAnggotaKelas({ commit }, payload) {
        try {
            const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete`, {
                params: {
                    schemaname: payload.schemaname,
                    anggota_rombel_id: payload.anggota_rombel_id
                }
            });
            const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => item.anggotaRombelId != payload.anggota_rombel_id);
            const updateState = {
                semester_id: state.selectedSemester.semesterId,
                peserta_didik: updateAnggotaRombel
            };
            // console.log(updateState);
            commit('SET_TABELSISWAAKTIF', updateState);
            return response.data; // Mengembalikan data
        } catch (error) {
            throw error;
        }
    },
    async deleteBatchAnggotaKelas({ commit }, payload) {
        try {
            const filteredIds = payload.anggota_rombel_id.filter((id) => id && id.trim() !== '');
            const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete-batch`, {
                params: {
                    schemaname: payload.schemaname,
                    anggota_rombel_id: filteredIds
                }
            });
            if (response.data.status) {
                const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => !filteredIds.includes(item.anggotaRombelId));
                const updateState = {
                    semester_id: state.selectedSemester.semesterId,
                    peserta_didik: updateAnggotaRombel
                };
                // console.log(updateState);
                commit('SET_TABELSISWAAKTIF', updateState);
                return response.data; // Mengembalikan data
            }
        } catch (error) {
            throw error;
        }
    },

    // =======================================
    // REFERENSI TABEL
    // =======================================
    async fetchTabeltenant({ commit }, sekolahTenantId) {
        try {
            const response = await api.get('/sekolah/sekolah-terdaftar', {
                params: {
                    sekolah_tenant_id: sekolahTenantId
                }
            });
            commit('SET_TABELTENANT', response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            console.error('Gagal mengambil data tabel tenant:', error);
            throw error;
        }
    },

    async createTabeltenant({ commit }, sekolah) {
        const payload = {
            sekolah: sekolah.sekolahData
        };

        // return;
        try {
            const response = await api.post('/sekolah/registrasi-sekolah', payload);
            console.log(response.data);
            commit('SET_TABELTENANT', response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            // commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat tabel tenant:', error);
            return null;
        }
    },
    async fetchBentukPendidikan({ commit }) {
        try {
            const response = await api.get(`/ss/ref/bentuk-pendidikan`);
            return response.data.bentukPendidikan;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat bentuk pendidikan:', error);
            return null;
        }
    },
    async fetchJenjangPendidikan({ commit }) {
        try {
            const response = await api.get(`/ss/ref/jenjang`);
            return response.data.jenjang;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat jenjang pendidikan:', error);
            return null;
        }
    },
    async fetchTingkatPendidikan({ commit }, payload) {
        try {
            const response = await api.get(`/ss/ref/tingkat-pendidikan`, {
                params: {
                    jenjang_pendidikan_id: payload.jenjang_pendidikan_id
                }
            });
            commit('SET_TABELTINGKATPENDIDIKAN', response.data.tingkatPendidikan);
            return response.data.tingkatPendidikan;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat tingkat pendidikan:', error);
            return null;
        }
    },
    async fetchKurikulum({ commit }, jenjangPendidikanId) {
        try {
            const response = await api.get(`/ss/ref/kurikulum`, {
                params: {
                    jenjang_pendidikan_id: jenjangPendidikanId
                }
            });
            if (response.data.status) {
                commit('SET_TABELKURIKULUM', response.data.kurikulum);
                return response.data;
            }
        } catch (error) {
            throw error;
        }
    },
    async fetchJurusan({ commit }, payload) {
        try {
            const response = await api.get(`/ss/ref/jurusan`, {
                params: {
                    jenjang_pendidikan_id: payload.jenjang_pendidikan_id,
                    param: payload.param
                }
            });
            return response.data.jurusan;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat kurikulum:', error);
            return null;
        }
    },
    async fetchMapel({ commit }, payload) {
        try {
            const response = await api.get(`/ss/ref/mapel`, {
                params: {
                    mapel_id: payload.mapel_id
                }
            });
            console.log('response.data.mapel');
            commit('SET_TABELMAPEL', response.data.mapel);
            return response.data.mapel;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal memuat mata pelajaran:', error);
            return null;
        }
    },
    async fetchGelarAkademik({ commit }, posisiGelar) {
        try {
            const response = await api.get(`/ss/ref/gelar-akademik`);
            commit('SET_GELARAKADEMIK', response.data.gelarAkademik);
            return response.data.gelarAkademik;
        } catch (error) {
            throw error;
        }
    },
    // =======================================
    // KELAS
    // =======================================

    async createKelas({ commit }, payload) {
        try {
            const response = await api.post(`/ss/${payload.schemaname}/tambah-kelas`, payload);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat kelas:', error);
            return null;
        }
    },
    async editKelas({ commit }, payload) {
        try {
            console.log('Payload di edit kelas:', payload);
            const response = await api.put(`/ss/${payload.schemaname}/kelas`, payload);
            // console.log(response.data);
            // commit("SET_TABELSEMESTER", response.data.semester);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal update kelas:', error);
            return null;
        }
    },
    async deleteKelas({ commit }, payload) {
        // try {
        //     const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete`, {
        //         params: {
        //             schemaname: payload.schemaname,
        //             anggota_rombel_id: payload.anggota_rombel_id
        //         }
        //     });
        //     const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => item.anggotaRombelId != payload.anggota_rombel_id);
        //     const updateState = {
        //         semester_id: state.selectedSemester.semesterId,
        //         peserta_didik: updateAnggotaRombel
        //     };
        //     // console.log(updateState);
        //     commit('SET_TABELSISWAAKTIF', updateState);
        //     return response.data; // Mengembalikan data
        // } catch (error) {
        //     throw error;
        // }
    },
    async deleteBatchKelas({ commit }, payload) {
        // try {
        //     const filteredIds = payload.anggota_rombel_id.filter((id) => id && id.trim() !== '');
        //     const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete-batch`, {
        //         params: {
        //             schemaname: payload.schemaname,
        //             anggota_rombel_id: filteredIds
        //         }
        //     });
        //     if (response.data.status) {
        //         const updateAnggotaRombel = state.tabelSiswaAktif.peserta_didik.filter((item) => !filteredIds.includes(item.anggotaRombelId));
        //         const updateState = {
        //             semester_id: state.selectedSemester.semesterId,
        //             peserta_didik: updateAnggotaRombel
        //         };
        //         // console.log(updateState);
        //         commit('SET_TABELSISWAAKTIF', updateState);
        //         return response.data; // Mengembalikan data
        //     }
        // } catch (error) {
        //     throw error;
        // }
    },

    // =========================
    // DASHBOARD
    // =========================
    async fetchDashboard({ commit }, payload) {
        try {
            const response = await api.get('/ss/dashboard/get-dashboard', {
                params: {
                    schemaname: payload.schemaname,
                    semester_id: payload.semester_id
                }
            });
            if (response) {
                commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            throw error;
        }
    },
    // =======================
    // ANGGOTA KELAS
    // =======================

    // async deleteAnggotaKelas({ commit }, payload) {
    //     try {
    //         const response = await api.delete(`/ss/${payload.schemaname}/anggota-kelas/delete`, {
    //             params: {
    //                 anggota_rombel_id: payload.anggota_rombel_id
    //             }
    //         });
    //         // console.log(response.data);
    //         // commit("SET_TABELSEMESTER", response.data.semester);
    //         return response.data; // Mengembalikan data sekolah
    //     } catch (error) {
    //         // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
    //         console.error('Gagal menghapus anggota kelas:', error);
    //         return null;
    //     }
    // },

    // ==================================
    // PEMBELAJARAN SERVICE
    // ==================================
    async createPembelajaran({ commit }, payload) {
        try {
            const response = await api.post(`ss/pembelajaran/create`, payload);
            console.log(response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal mendaftarkan siswa baru:', error);
            return null;
        }
    },

    // ==================================
    // IJAZAH SERVICE
    // ==================================
    async createProsesIjazah({ commit }, payload) {
        try {
            const response = await api.post(`ss/ijazah/create`, payload);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            // console.error("Gagal mendaftarkan siswa baru:", error);
            throw error;
        }
    },
    async fetchProsesIjazah({ commit }, payload) {
        try {
            const response = await api.get(`ss/proses-ijazah`, {
                params: {
                    schemaname: payload.schemaname,
                    semester_id: payload.tahun_ajaran_id,
                    ijazah_id: payload.ijazah_id
                }
            });
            return response.data.anggotaKelas; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },

    // =============================================
    //  DNS
    // =============================================
    async createDns({ commit }, payload) {
        try {
            const response = await api.post(`ss/ijazah/data-nominasi_sementara`, payload);
            return response.data;
        } catch (error) {
            throw error;
        }
    },
    async fetchDns({ commit }, payload) {
        try {
            const response = await api.get(`/ss/data-nominasi_sementara`, {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    is_complete: payload.is_complete
                }
            });
            // console.log(response.data);
            if (response.status) {
                const results = {
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    dataNominasiSementara: response.data.dataNominasiSementara
                };
                commit('SET_TABELDNS', results);
                results.message = response.data.message;
                results.status = response.data.status;
                console.log(results);
                return results;
            }
            return nil;
        } catch (error) {
            // console.log(error);
            throw error;
        }
    },
    async searchDns({ commit }, payload) {
        try {
            const response = await api.get('/ss/data-nominasi_sementara/search', {
                params: {
                    schemaname: payload.schemaname,
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            return response.data;
        } catch (error) {
            throw error;
        }
    },
    async deleteDns({ commit }, payload) {
        try {
            const response = await api.delete(`ss/data-nominasi_sementara`, {
                params: {
                    schemaname: payload.schemaname,
                    peserta_didik_id: payload.peserta_didik_id
                }
            });
            if (response.status) {
                const cek = state.tabelDns.dataNominasiSementara.filter((item) => item.pesertaDidikId != payload.peserta_didik_id);
                const results = {
                    tahun_ajaran_id: payload.tahun_ajaran_id,
                    dataNominasiSementara: cek
                };
                commit('SET_TABELDNS', results);
                results.message = response.data.message;
                results.status = response.data.status;
                // console.log(results);
                return results;
            }
            return response.data;
        } catch (error) {
            throw error;
        }
    },

    async createInfoIjazah({ commit }, payload) {
        console.log(payload);

        const response = await api.post(`ss/ijazah/seting-ijazah`, payload);
        console.log(response);
    },
    async fetchNilaiSiswa({ commit }, payload) {
        try {
            const response = await api.get(`ss/${payload.schemaname}/nilai-akhir`, {
                params: {
                    semester_id: payload.semesterId,
                    peserta_didik_id: payload.semester_id
                }
            });
            commit('SET_TABELNILAIAKHIR', response.data?.nilaiSiswa);
            return response.data.nilaiSiswa; // Mengembalikan nilai siswa
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal nilai siswa:', error);
            return null;
        }
    },
    async createKenaikan({ commit }, payload) {
        try {
            const response = await api.post(`ss/kenaikan/create`, payload);
            // commit("SET_TABELNILAIAKHIR", response.data?.nilaiSiswa);

            console.log(response.data);
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal nilai siswa:', error);
            return null;
        }
    }
};

const getters = {
    getKelas: (state) => (semesterId) => {
        return state.tabelKelas.filter((kelas) => kelas.semesterId === semesterId);
    },
    isLoading: (state) => state.loading,
    getError: (state) => state.error,
    getTabeltenant: (state) => state.tabelTenant,
    getSemester: (state) => state.tabelSemester,
    getSelectedSemester: (state) => state.selectedSemester,
    getSekolah: (state) => state.tabelSekolah,
    getKurikulum: (state) => state.tabelKurikulum,
    getGuru: (state) => state.tabelGuru,
    getPTKTerdaftar: (state) => state.tabelPTKTerdaftar,
    getTingkatPendidikan: (state) => state.tabelTingkatPendidikan,
    getTabelAnggotaKelas: (state) => state.tabelAnggotaKelas,
    getSiswaAktif: (state) => state.tabelSiswaAktif,
    getMapel: (state) => state.tabelMapel,
    getNilaiSiswa: (state) => state.tabelNilaiakhir,
    getTahunAjaran: (state) => state.tabelTahunAjaran,
    getSelectedTahunAjaran: (state) => state.selectedTahunAjaran,
    getGelarAkademik: (state) => state.tabelGelarAkademik,
    getDashboard: (state) => state.tabelDashboard,
    getDns: (state) => state.tabelDns
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
