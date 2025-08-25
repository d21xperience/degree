import api from '../api';
const state = {
    tabelTenant: JSON.parse(localStorage.getItem('tabelTenant')) || null,
    tabelSekolah: JSON.parse(localStorage.getItem('tabelSekolah')) || null,
    tabelTingkatPendidikan: JSON.parse(localStorage.getItem('tabelTingkatPendidikan')) || null,
    tabelKurikulum: JSON.parse(localStorage.getItem('tabelKurikulum')) || null,
    tabelJurusan: JSON.parse(localStorage.getItem('tabelJurusan')) || null,
    tabelMapel: JSON.parse(localStorage.getItem('tabelMapel')) || null,
    tabelTahunAjaran: JSON.parse(localStorage.getItem('tabelTahunAjaran')) || [],
    // selectedTahunAjaran: JSON.parse(localStorage.getItem('selectedTahunAjaran')) || [],
    tabelGelarAkademik: JSON.parse(localStorage.getItem('gelarAkademik')) || [],
    tabelDashboard: JSON.parse(localStorage.getItem('tabelDashboard')) || []
};

const mutations = {
    SET_TABELTENANT(state, tabelTenant) {
        state.tabelTenant = tabelTenant;
        localStorage.setItem('tabelTenant', JSON.stringify(tabelTenant));
    },

    SET_TABELSEKOLAH(state, tabelSekolah) {
        state.tabelSekolah = tabelSekolah;
        localStorage.setItem('tabelSekolah', JSON.stringify(tabelSekolah));
    },

    SET_TABELTINGKATPENDIDIKAN(state, value) {
        state.tabelTingkatPendidikan = value;
        localStorage.setItem('tabelTingkatPendidikan', JSON.stringify(value));
    },
    SET_TABELKURIKULUM(state, value) {
        state.tabelKurikulum = value;
        localStorage.setItem('tabelKurikulum', JSON.stringify(value));
    },
    SET_TABELJURUSAN(state, value) {
        state.tabelJurusan = value;
        localStorage.setItem('tabelJurusan', JSON.stringify(value));
    },

    SET_TABELMAPEL(state, value) {
        state.tabelMapel = value;
        localStorage.setItem('tabelMapel', JSON.stringify(value));
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

    resetState(state) {
        state.tabelTenant = null;

        state.tabelSekolah = null;

        state.tabelTingkatPendidikan = null;
        state.tabelKurikulum = null;
        state.tabelJurusan = null;

        state.tabelMapel = null;

        state.tabelTahunAjaran = [];
        state.selectedTahunAjaran = [];
        state.tabelGelarAkademik = [];
        state.tabelDashboard = [];
    }
};

const actions = {
    async resetState({ commit }) {
        try {
            commit('resetState');
        } finally {
            commit('resetState');
        }
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
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan tahun ajaran:', error);
        }
    },

    async fetchSelectedTahunAjaran({ commit }, payload) {
        commit('SET_SELECTEDTAHUNAJARAN', payload);
    },
    // ================================================

    // =============================================

    async getTemplate({ commit }) {
        commit('SET_LOADING', true);
        commit('SET_ERROR', null);
        try {
            const response = await api.get(`/ss/download/template`, {
                params: {
                    'template-type': 'siswa'
                }
            });
            return response;
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat template:', error);
            return null;
        }
    },
    async fetchSekolah({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/sekolah`);
            commit('SET_TABELSEKOLAH', response.data);
            return response.data;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async updateSekolah(payload) {
        try {
            const response = await api.put(`/ss/${payload.schemaname}/update`, payload);
            if (response.status) {
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    // =======================================
    // REFERENSI TABEL
    // =======================================
    async fetchTabeltenant({ commit }, sekolahTenantId) {
        try {
            const response = await api.get('/ss/sekolah/sekolah-terdaftar', {
                params: {
                    sekolah_tenant_id: sekolahTenantId
                }
            });
            commit('SET_TABELTENANT', response.data);
            return response.data;
        } catch (error) {
            console.error('Gagal mengambil data tabel tenant:', error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async createTabeltenant({ commit }, sekolah) {
        const payload = {
            sekolah: sekolah.sekolahData
        };
        try {
            const response = await api.post('/ss/sekolah/registrasi-sekolah', payload);
            console.log(response.data);
            commit('SET_TABELTENANT', response.data);
            return response.data;
        } catch (error) {
            // commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat tabel tenant:', error);
            return null;
        }
    },
    async fetchBentukPendidikan() {
        try {
            const response = await api.get(`/ss/ref/bentuk-pendidikan`);
            return response.data.bentukPendidikan;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat bentuk pendidikan:', error);
            return null;
        }
    },
    async fetchJenjangPendidikan() {
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
    async fetchKurikulum({ commit }, payload) {
        try {
            const response = await api.get(`/ss/ref/kurikulum`, {
                params: {
                    jenjang_pendidikan_id: payload.jenjangPendidikanId,
                    jenjang_pendidikan_str: payload.jenjangPendidikanStr
                }
            });
            if (response.data.status) {
                commit('SET_TABELKURIKULUM', response.data.kurikulum);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    // ==============================
    // Bidang Keahlian

    async fetchBidangKeahlian(payload) {
        try {
            const response = await api.get(`/ss/ref/bidang-keahlian`, {
                params: {
                    jurusan_induk: payload.jurusanInduk
                }
            });
            if (response.data.status) {
                // commit('SET_TABELKURIKULUM', response.data.kurikulum);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    // ==============================
    // Program Keahlian

    async fetchProgramKeahlian(payload) {
        try {
            const response = await api.get(`/ss/ref/program-keahlian`, {
                params: {
                    jurusan_induk: payload.jurusanInduk
                }
            });
            if (response.data.status) {
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async fetchJurusan(payload) {
        try {
            const response = await api.get(`/ss/ref/jurusan`, {
                params: {
                    jurusan_induk: payload.jurusanInduk
                }
            });
            if (response.status) {
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
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
    async filterMapel(payload) {
        try {
            // console.log("filterMapel",payload)
            const response = await api.get(`/ss/ref/mapel/filter`, {
                params: {
                    query: payload.query
                }
            });
            // console.log('response.data.mapel', response);
            // commit('SET_TABELMAPEL', response.data.mapel);
            return response.data.mapel;
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal memuat mata pelajaran:', error);
            return null;
        }
    },
    async fetchMapelKurikulum({ commit }, payload) {
        try {
            const response = await api.get(`/ss/ref/mapel-kurikulum`, {
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
    async fetchGelarAkademik({ commit }) {
        try {
            const response = await api.get(`/ss/ref/gelar-akademik`);
            commit('SET_GELARAKADEMIK', response.data.gelarAkademik);
            return response.data.gelarAkademik;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
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
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async fetchKategoriSekolah(payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/kategori-sekolah`, {
                params: {
                    tahun_ajaran_id: payload.tahun_ajaran_id
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async createKategoriSekolah(payload) {
        // console.log(payload);
        try {
            const response = await api.post(`/ss/${payload.schemaname}/kategori-sekolah/create`, payload);
            if (response.status) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async updateKategoriSekolah(payload) {
        try {
            // console.log(payload)
            // return
            const response = await api.put(`/ss/${payload.schemaname}/kategori-sekolah/update`, payload);
            if (response.status) {
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteKategoriSekolah(payload) {
        try {
            const response = await api.delete(`/ss/${payload.schemaname}/kategori-sekolah/delete`, {
                params: {
                    kategori_sekolah_id: payload.kategori_sekolah_id
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteKategoriSekolahKurikulum(payload) {
        try {
            const response = await api.delete(`/ss/${payload.schemaname}/kategori-sekolah/delete`, {
                params: {
                    kurikulum_id: payload.kurikulum_id
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async createProsesKelas(payload) {
        try {
            // console.log(payload);
            // return
            const response = await api.post(`/ss/${payload.schemaname}/kategori-sekolah-kelas/proses`, payload);
            if (response.status) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async fetchKategoriMapel(payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/kategori-sekolah/mapel`, {
                params: {
                    tahun_ajaran_id: `${payload.tahunAjaranId}`,
                    tingkat_pendidikan: payload.tingkatPendidikan,
                    kurikulum_id: payload.kurikulumId
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },

    async deleteKategoriMapel(payload) {
        try {
            const response = await api.delete(`/ss/${payload.schemaname}/kategori-sekolah/mapel/delete`, {
                params: {
                    id: payload.id
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    async deleteBatchKategoriMapel(payload) {
        try {
            console.log('sekolahService =>', payload);
            // return;
            const response = await api.delete(`/ss/${payload.schemaname}/kategori-sekolah/mapel/batch-delete`, {
                params: {
                    id: payload.id
                }
            });
            if (response) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            console.log(error);
            throw new Error('Gagal menghapus Kategori Mapel:', error);
        }
    },
    // ==================================
    // PEMBELAJARAN SERVICE
    // ==================================
    async createPembelajaran({ commit }, payload) {
        try {
            const response = await api.post(`ss/pembelajaran/create`, payload);
            console.log(response.data);
            return response.data;
        } catch (error) {
            commit('SET_ERROR', error.response?.data || 'Terjadi kesalahan');
            console.error('Gagal membuat pembelajaran:', error);
            return null;
        }
    },

    async createKenaikan(payload) {
        try {
            const response = await api.post(`ss/kenaikan/create`, payload);
            console.log(response.data);
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            console.error('Gagal membuat kenaikan siswa:', error);
            return null;
        }
    },

    async updateSettingSekolah({ commit }, payload) {
        commit('sekolahService/SET_TABELSEKOLAH', payload);
    }
};

const getters = {
    isLoading: (state) => state.loading,
    getError: (state) => state.error,
    getTabeltenant: (state) => state.tabelTenant,
    getSekolah: (state) => state.tabelSekolah,
    getKurikulum: (state) => state.tabelKurikulum,
    getJurusan: (state) => state.tabelJurusan,
    getTingkatPendidikan: (state) => state.tabelTingkatPendidikan,
    getMapel: (state) => state.tabelMapel,

    getGelarAkademik: (state) => state.tabelGelarAkademik,
    getDashboard: (state) => state.tabelDashboard
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
