/* eslint-disable no-unused-vars */
import api from '../api';
const state = {
    tabelTenant: JSON.parse(localStorage.getItem('tabelTenant')) || null,
    tabelSekolah: JSON.parse(localStorage.getItem('tabelSekolah')) || null,
    tabelTingkatPendidikan: JSON.parse(localStorage.getItem('tabelTingkatPendidikan')) || null,
    tabelJurusan: JSON.parse(localStorage.getItem('tabelJurusan')) || null,
    tabelMapel: JSON.parse(localStorage.getItem('tabelMapel')) || null,
    tabelGelarAkademik: JSON.parse(localStorage.getItem('gelarAkademik')) || [],
    tabelDashboard: JSON.parse(localStorage.getItem('tabelDashboard')) || null,
    tabelJenjangPendidikan: JSON.parse(localStorage.getItem('tabelJenjangPendidikan')) || null,
    tabelBentukPendidikan: JSON.parse(localStorage.getItem('tabelBentukPendidikan')) || null,
    tabelKategoriSekolah: JSON.parse(localStorage.getItem('tabelKategoriSekolah')) || [],
    isKategoriSekolahCompleted: JSON.parse(localStorage.getItem('isKategoriSekolahCompleted')) || false
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

    SET_TABELJURUSAN(state, value) {
        state.tabelJurusan = value;
        localStorage.setItem('tabelJurusan', JSON.stringify(value));
    },

    SET_TABELMAPEL(state, value) {
        state.tabelMapel = value;
        localStorage.setItem('tabelMapel', JSON.stringify(value));
    },

    SET_GELARAKADEMIK(state, value) {
        state.tabelGelarAkademik = value;
        localStorage.setItem('gelarAkademik', JSON.stringify(value));
    },
    SET_DASHBOARD(state, value) {
        state.tabelDashboard = value;
        localStorage.setItem('tabelDashboard', JSON.stringify(value));
    },

    SET_TABELJENJANGPENDIDIKAN(state, value) {
        state.tabelJenjangPendidikan = value;
        localStorage.setItem('tabelJenjangPendidikan', JSON.stringify(value));
    },
    SET_TABELBENTUKPENDIDIKAN(state, value) {
        state.tabelBentukPendidikan = value;
        localStorage.setItem('tabelBentukPendidikan', JSON.stringify(value));
    },
    SET_TABELKATEGORISEKOLAH(state, value) {
        state.tabelKategoriSekolah = value;
        localStorage.setItem('tabelKategoriSekolah', JSON.stringify(value));
    },
    SET_ISKATEGORISEKOLAHCOMPLETED(state, value) {
        state.isKategoriSekolahCompleted = value;
        localStorage.setItem('isKategoriSekolahCompleted', JSON.stringify(value));
    },
    resetState(state) {
        localStorage.clear();
        state.tabelTenant = null;
        state.tabelSekolah = null;
        state.tabelTingkatPendidikan = null;
        state.tabelJurusan = null;
        state.tabelMapel = null;
        state.selectedTahunAjaran = [];
        state.tabelGelarAkademik = [];
        state.tabelDashboard = [];
        state.tabelJenjangPendidikan = null;
        state.tabelKategoriSekolah = [];
        state.isKategoriSekolahCompleted = false;
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
            throw new Error(`Gagal membuat template: ${error}`);
        }
    },
    async fetchSekolah({ commit }, payload) {
        try {
            const response = await api.get(`/ss/${payload.schemaname}/sekolah`);
            if (response.status) {
                commit('SET_TABELSEKOLAH', response.data);
                return response.data;
            }
        } catch (error) {
            throw new Error(`Gagal mengambil data sekolah: ${error}`);
        }
    },
    async updateSekolah({ commit }, payload) {
        try {
            const response = await api.put(`/ss/${payload.schemaname}/update`, payload);
            if (response.status) {
                commit('SET_TABELSEKOLAH', payload);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal update sekolah: ${error}`);
        }
    },

    // =======================================
    // REFERENSI TABEL
    // =======================================
    async fetchTabeltenant({ commit }, sekolahTenantId) {
        try {
            const { data } = await api.get('/ss/sekolah/sekolah-terdaftar', {
                params: {
                    sekolah_tenant_id: sekolahTenantId
                }
            });
            commit('SET_TABELTENANT', data);
            return { status: true, schemaname: data.schemaname };
        } catch (error) {
            throw new Error(`Gagal mengambil data tabel tenant: ${error}`);
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
            throw new Error(`Gagal membuat tabel tenant: ${error}`);
        }
    },
    async fetchBentukPendidikan({ commit }) {
        try {
            const { data } = await api.get('/ss/ref/bentuk-pendidikan');
            // console.log('Bentuk Pendidikan ====', data);
            commit('SET_TABELBENTUKPENDIDIKAN', data?.bentukPendidikan);
            return data?.bentukPendidikan;
        } catch (error) {
            throw new Error(`Gagal membuat bentuk pendidikan: ${error}`);
        }
    },
    async fetchJenjangPendidikan({ commit }, payload) {
        try {
            const { data } = await api.get(`/ss/ref/jenjang`, {
                params: {
                    is_jenjang_orang: payload.isJenjangOrang,
                    is_jenjang_lembaga: payload.isJenjangLembaga,
                    jenjang_orang: payload.jenjangOrang,
                    jenjang_lembaga: payload.jenjangLembaga
                }
            });
            if (data.status) {
                commit('SET_TABELJENJANGPENDIDIKAN', data.jenjang);
                return data;
            }
        } catch (error) {
            throw new Error(`Gagal mendapatkan jenjang pendidikan: ${error}`);
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
            throw new Error(`Gagal mendapatkan tingkat pendidikan: ${error}`);
        }
    },

    // ==============================
    // Bidang Keahlian

    async fetchBidangKeahlian({ commit }, payload) {
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
            throw new Error(`Gagal mendapatkan bidang keahlian: ${error}`);
        }
    },
    // ==============================
    // Program Keahlian

    async fetchProgramKeahlian({ commit }, payload) {
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
            throw new Error(`Gagal mendapatkan program keahlian: ${error}`);
        }
    },
    async fetchJurusan({ commit }, payload) {
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
            throw new Error(`Gagal mendapatkan jurusan: ${error}`);
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
            throw new Error(`Gagal mendapatkan mata pelajaran: ${error}`);
        }
    },
    async filterMapel({ commit }, payload) {
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
            throw new Error(`Gagal memfilter mata pelajaran: ${error}`);
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
            throw new Error(`Gagal medapatakan mapel kurikulum: ${error}`);
        }
    },
    async fetchGelarAkademik({ commit }) {
        try {
            const response = await api.get(`/ss/ref/gelar-akademik`);
            commit('SET_GELARAKADEMIK', response.data.gelarAkademik);
            return response.data.gelarAkademik;
        } catch (error) {
            throw new Error(`Gagal mendapatkan gelar akademik: ${error}`);
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
            console.log(payload);
            if (response) {
                commit('SET_DASHBOARD', { semester_id: payload.semester_id, data: response.data });
                return response.data;
            }
        } catch (error) {
            throw new Error(`Gagal mendapatkan data dashboard: ${error}`);
        }
    },

    async fetchKategoriSekolah({ commit }, payload) {
        try {
            const { data } = await api.get(`/ss/${payload.schemaname}/kategori-sekolah`, {
                params: {
                    tahun_ajaran_id: payload.tahun_ajaran_id
                }
            });
            if (data.status) {
                commit('SET_TABELKATEGORISEKOLAH', {
                    tahunAjaranId: payload.tahun_ajaran_id,
                    kategoriSekolah: data.kategoriSekolah
                });
                if (data.kategoriSekolah.length == 0) {
                    commit('SET_ISKATEGORISEKOLAHCOMPLETED', false);
                } else {
                    commit('SET_ISKATEGORISEKOLAHCOMPLETED', true);
                }
                return data.kategoriSekolah;
            }
        } catch (error) {
            commit('SET_ISKATEGORISEKOLAHCOMPLETED', false);
            throw new Error(error.message);
        }
    },
    async createKategoriSekolah({ commit }, payload) {
        // console.log({commit}, payload);
        try {
            const response = await api.post(`/ss/${payload.schemaname}/kategori-sekolah/create`, payload);
            if (response.status) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    async updateKategoriSekolah({ commit }, payload) {
        try {
            // console.log({commit}, payload)
            // return
            const response = await api.put(`/ss/${payload.schemaname}/kategori-sekolah/update`, payload);
            if (response.status) {
                return response.data;
            }
        } catch (error) {
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    async deleteKategoriSekolah({ commit }, payload) {
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
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    async deleteKategoriSekolahKurikulum({ commit }, payload) {
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
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },

    async createProsesKelas({ commit }, payload) {
        try {
            console.log(payload);
            // return
            const response = await api.post(`/ss/${payload.schemaname}/kategori-sekolah-kelas/proses`, payload);
            if (response.status) {
                // commit('SET_DASHBOARD', response.data);
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw new Error(`Gagal membuat Kategori Mapel: ${error}`);
        }
    },

    async fetchKategoriMapel({ commit }, payload) {
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
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },

    async deleteKategoriMapel({ commit }, payload) {
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
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    async deleteBatchKategoriMapel({ commit }, payload) {
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
            throw new Error(`Gagal menghapus Kategori Mapel: ${error}`);
        }
    },
    // ==================================
    // PEMBELAJARAN SERVICE
    // ==================================
    async createPembelajaran({ commit }, payload) {
        try {
            const response = await api.post(`ss/pembelajaran/create`, payload);
            return response.data;
        } catch (error) {
            throw new Error(`Gagal membuat pembelajaran: ${error}`);
        }
    },

    async createKenaikan({ commit }, payload) {
        try {
            const response = await api.post(`ss/kenaikan/create`, payload);
            console.log(response.data);
        } catch (error) {
            // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
            throw new Error(`Gagal membuat kenaikan siswa: ${error}`);
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

    getJurusan: (state) => state.tabelJurusan,
    getTingkatPendidikan: (state) => state.tabelTingkatPendidikan,
    getBentukPendidikan: (state) => state.tabelBentukPendidikan,
    getJenjangPendidikan: (state) => state.tabelJenjangPendidikan,
    getMapel: (state) => state.tabelMapel,
    getGelarAkademik: (state) => state.tabelGelarAkademik,
    getDashboard: (state) => state.tabelDashboard,
    getTabelKategoriSekolah: (state) => state.tabelKategoriSekolah,
    getIsKategoriSekolahCompleted: (state) => state.isKategoriSekolahCompleted
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
