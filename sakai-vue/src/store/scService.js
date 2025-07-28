// import axios from 'axios';
// const api = axios.create({
//     baseURL: import.meta.env.VITE_API_BASE_URL, //'http://localhost:8184/api/v1',
//     withCredentials: true, // Untuk mengirim cookie atau credensial
//     headers: {
//         'Content-Type': 'application/json',
//         Authorization: 'Bearer your_token'
//     }
// });
import api from './api';
const state = {
    BCNetwork: JSON.parse(localStorage.getItem('BCNETWORK')) || null,
    BCNetworkSelected: JSON.parse(localStorage.getItem('BCNetworkSelected')) || null,
    BCPlatformSelected: JSON.parse(localStorage.getItem('BCPlatformSelected')) || null,
    BCAccountActivate: null,
    MetamasConnected: JSON.parse(localStorage.getItem('METAMASK_CONNECTED')) || null,
    contract: JSON.parse(localStorage.getItem('CONTRACT')) || null,
    SCIjazah: JSON.parse(localStorage.getItem('SCIjazah')) || null,
    walletInfo: JSON.parse(localStorage.getItem('WALLET_INFO')) || null,
    bcConnected: null
};

const mutations = {
    SET_BCNETWORK(state, value) {
        state.BCNETWORK = value;
        localStorage.setItem('BCNETWORK', JSON.stringify(value));
    },
    SET_BCNetworkSelected(state, value) {
        state.BCNetworkSelected = value;
        localStorage.setItem('BCNetworkSelected', JSON.stringify(value));
    },
    SET_BCACCOUNT(state, value) {
        state.BCACCOUNT = value;
        localStorage.setItem('BCACCOUNT', JSON.stringify(value));
    },
    setBCPlatformSelected(state, value) {
        state.BCPlatformSelected = value;
        localStorage.setItem('BCPlatformSelected', JSON.stringify(value));
    },
    setBCAccountActivate(state, value) {
        state.BCAccountActivate = value;
    },
    SET_METAMASKCONNECTED(state, value) {
        state.MetamasConnected = value;
        localStorage.setItem('METAMASK_CONNECTED', JSON.stringify(value));
    },
    SET_CONTRACT(state, value) {
        state.contract = value;
        localStorage.setItem('CONTRACT', JSON.stringify(value));
    },
    SET_SCIjazah(state, value) {
        state.SCIjazah = value;
        localStorage.setItem('SCIjazah', JSON.stringify(value));
    },
    SET_WALLETINFO(state, value) {
        state.walletInfo = value;
        localStorage.setItem('WALLETINFO', JSON.stringify(value));
    },
    SET_BcConected(state, value) {
        state.bcConnected = value;
        // localStorage.setItem('SET_BcConected', JSON.stringify(value));
    }
};

const actions = {
    async fetchBCPlatform({ commit }, payload) {
        try {
            const { data } = await api.get('/scs/platform');
            data.bcPlatform.find((platform) => {
                if (platform.active) {
                    commit('setBCPlatformSelected', platform);
                }
            });
            return data;
        } catch (error) {
            throw error;
        }
    },

    async setBCPlatform({ commit }, payload) {
        // console.log("in vuex: ", payload);
        const py = {
            bc_platform: {
                id: payload.bc_platform.id,
                name: payload.bc_platform.name,
                active: payload.bc_platform.active
            },
            schemaname: payload.schemaname
        };

        // console.log(py);
        try {
            const response = await api.put('/scs/platform', py);
            return response.data;
        } catch (error) {
            console.error('Gagal set BC platform:', error);
            return null;
        }
        // console.log(response);
    },
    async updateBCPlatformSelected({ commit }, value) {
        commit('setBCPlatformSelected', value);
    },

    // ================================
    async fetchBlockchainNetworks({ commit }, networkArchitecture) {
        try {
            const response = await api.get(`scs/bc-networks`, {
                params: {
                    network_architecture: networkArchitecture
                }
            });
            commit('SET_BCNETWORK', response.data.network);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            console.error('Gagal memuat network:', error);
            throw error;
        }
    },

    async updateBlockchainNetwork({ commit }, payload) {
        try {
            // console.log(payload)
            const { data } = await api.put('scs/bc-network', payload);
            return data;
        } catch (error) {
            throw error;
        }
    },

    async deleteBlockchainNetwork({ commit }, payload) {
        try {
            const networkIds = payload.map((item) => Number(item.Id)).filter((id) => Number.isInteger(id) && id >= 0 && id <= 4294967295); // validasi uint32

            const params = new URLSearchParams();
            networkIds.forEach((id) => {
                params.append('network_ids', id);
            });

            const { data } = await api.delete(`scs/bc-network?${params.toString()}`);
            return data;
        } catch (error) {
            console.error('deleteBlockchainNetwork error:', error.response?.data || error.message);
            throw error;
        }
    },
    async setBlockchainNetwork({ commit }, value) {
        commit('SET_BCNetworkSelected', value);
    },
    // ==========================akun=========================
    // ===================================================
    async fetchBCAccount({ commit }, username) {
        try {
            const { data } = await api.get('scs/blockchainaccount/list', {
                params: {
                    username: username
                }
            });
            // console.log('scService/fetchBCAccount', data);
            return data;
            // console.log(response.data);
            // commit('SET_BCACCOUNT', response.data.blockchainaccounts);
            // return response.data; // Mengembalikan data sekolah
        } catch (error) {
            console.error('Gagal mengambil akun blockchain:', error);
            throw error;
        }
    },
    async createBCAccount({ commit }, payload) {
        let network_name = payload.network_name || 0;
        let user_id = payload.user_id || 0;
        let password = payload.password || 0;
        let schemaname = payload.schemaname || null;

        try {
            const response = await api.post(`/blockchainaccount/create`, {
                params: {
                    schemaname: schemaname,
                    password: password,
                    user_id: user_id,
                    network_name: network_name
                }
            });
            // commit("SET_BCNETWORK", response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },
    async importBCAccount({ commit }, payload) {
        console.log('Payload yang dikirim:', JSON.stringify(payload, null, 2));

        try {
            const response = await api.post(`/blockchainaccount/import`, JSON.stringify(payload, null, 2), {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            // commit("SET_BCNETWORK", response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },
    async updateBCAccount({ commit }, payload) {
        let network_name = payload.network_name || 0;
        let user_id = payload.user_id || 0;
        let password = payload.password || 0;
        let schemaname = payload.schemaname || null;

        try {
            const response = await api.get(`/blockchainaccount/create`, {
                params: {
                    schemaname: schemaname,
                    password: password,
                    user_id: user_id,
                    network_name: network_name
                }
            });
            // commit("SET_BCNETWORK", response.data);
            return response.data; // Mengembalikan data sekolah
        } catch (error) {
            throw error;
        }
    },

    // ================================

    async createIjazahBC({ commit }, payload) {
        try {
            // console.log(JSON.stringify(payload));
            // return;
            const response = await api.post(`/scs/ijazah-bc/create`, JSON.stringify(payload));
            return response.data;
        } catch (error) {
            throw error;
        }
    },
    async fetchIjazahBC({ commit }, payload) {
        try {
            const response = await api.get(`/scs/ijazah-bc`, {
                params: {
                    sekolah_id: payload.sekolah_id,
                    tahun_ajaran_id: payload.tahun_ajaran_id
                }
            });
            return response.data;
        } catch (error) {
            throw error;
        }
    },
    async searchIjazahBC({ commit }, payload) {
        try {
            const response = await api.get(`/scs/ijazah-bc/search`, {
                params: {
                    nisn: payload.nisn
                }
            });
            return response.data.ijazahBc;
        } catch (error) {
            throw error;
        }
    },
    // ================================

    async updateBCAccountActivate({ commit }, value) {
        commit('setBCAccountActivate', value);
    },
    // ==================================
    // METAMASK
    // ==================================
    // async createMetamaskConnected ({commit}, value){
    //     try {

    //     } catch (error) {
    //         throw error
    //     }

    async fetchContract({ commit }, payload) {
        try {
            const response = await api.get(`/scs/contract-address`);
            if (response) {
                commit('SET_CONTRACT', response.data.contract);
                return response.data;
            }
        } catch (error) {
            throw error;
        }
    },

    // ============================================
    // BC IJAZAH
    // ============================================
    async fetchSCIjazah({ commit }, payload) {
        try {
            // console.log(payload);
            const response = await api.get(`/scs/ijazah-bc`, {
                params: {
                    sekolah_id: payload.sekolah_id,
                    tahun_ajaran_id: `${payload.tahun_ajaran_id}`
                }
            });
            if (response) {
                const results = { tahun_ajaran_id: `${payload.tahun_ajaran_id}`, degreeData: response.data.degreeData };
                commit('SET_SCIjazah', results);
                return response.data;
            }
        } catch (error) {
            throw error;
        }
    },
    async fetchBCTransaction({ commit }, payload) {
        try {
            const response = await api.get('/scs/bc-transaction', {
                params: {
                    schemaname: payload.schemaname
                }
            });
            if (response.status) {
                return response.data;
            }
        } catch (error) {
            console.log(error);
            throw error;
        }
    },

    // ==================================
    // EVM OR NONEVM
    // ==================================
    async getNetworkChainId({ commit }, payload) {
        try {
            const { data } = await api.get('scs/bc-networks/ethereum-network', {
                params: {
                    rpc_url: payload
                }
            });
            return data;
        } catch (error) {
            throw error;
        }
    },

    async createWalletInfo({ commit }, payload) {
        try {
            commit('SET_WALLETINFO', payload);
        } catch (error) {}
    },

    async fetchWalletInfo({ commit }, payload) {
        try {
            const { data } = await api.post('scs/blockchainaccount/wallet', payload);
            if (data.status) {
                commit('SET_WALLETINFO', data.walletData);
                return data;
            }
        } catch (error) {
            console.log(error);
            throw error;
        }
    },
    // ======== SMART CONTRACT =======

    async getsolVersion({ commit }) {
        try {
            return await api.get('/scs/contract/deploy/solc-version');
        } catch (error) {
            console.log(error);
        }
    },
    async deployContract({ commit }, formData) {
        const response = await fetch('/scs/contract/compile-contract', {
            method: 'POST',
            body: formData
        });

        const result = await response.json();
        console.log(result);
    },

    // ==========================================
    // Blockhain Service
    // ==========================================
    async setBCConfig({ commit }, payload) {
        try {
            const { data } = await api.post('scs/blockchain/config', payload);
            console.log(data);
            if (data.status) {
                return data;
            }
        } catch (error) {
            throw error;
        }
    }
};

// ==========================================
// GETTERS
// ==========================================
const getters = {
    getBCNETWORK:
        (state) =>
        (filter = {}) => {
            const list = state.BCNetwork;
            if (Object.keys(filter).length === 0) return list;

            return list.filter((item) => Object.entries(filter).every(([key, value]) => item[key] === value));
        },
    getBCPlatformSelected: (state) => state.BCPlatformSelected,
    getBCNetworkSelected: (state) => state.BCNetworkSelected,
    getBCAccount: (state) => state.BCACCOUNT,
    getBCAccountActivate: (state) => state.BCAccountActivate,
    getMetamaskConnected: (state) => state.MetamasConnected,
    getContract: (state) => state.contract,
    getSCIjazah: (state) => state.SCIjazah,
    getWalletInfo: (state) => state.walletInfo,
    getBcConnected: (state) => state.bcConnected
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
};
