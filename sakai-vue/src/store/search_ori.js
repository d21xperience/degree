// import axios from "axios";
// import SekolahService from "@/services/SekolahService";
// const state = {
//   searchCache: {},
//   results: [],
//   cacheTimestamps: [],
// };

// const mutations = {
//   // SET_DAPO(state, { query, results }) {
//   //   state.searchCache[query] = results;
//   //   state.cacheTimestamps[query] = Date.now(); // Simpan timestamp
//   // },
//   SET_RESULTS(state, { query, results }) {
//     state.searchCache[query] = results;
//     state.results = results;
//   },
// };

// const actions = {
//   // async fetchResults({ state, commit }, query) {
//   //   const now = Date.now();
//   //   const cacheTTL = 300000; // TTL: 5 menit

//   //   if (
//   //     state.searchCache[query] &&
//   //     now - state.cacheTimestamps[query] < cacheTTL
//   //   ) {
//   //     console.log("Data from cache..");
//   //     return state.searchCache[query];
//   //   }

//   //   try {
//   //     console.log("Ambil data dari server");
//   //     // const response = await axios.get("/dapo/api/getHasilPencarian", {
//   //     //   params: { keyword: query },
//   //     // });
//   //     const response = await fetch(
//   //       `https://api.allorigins.win/get?url=${encodeURIComponent(
//   //         "https://dapo.kemdikbud.go.id/api/getHasilPencarian?keyword="
//   //       )}${query}`
//   //     );

//   //     if (!response.ok) {
//   //       throw new Error("Network response was not ok.");
//   //     }

//   //     const data = await response.json();
//   //     // Parse JSON contents
//   //     const results = JSON.parse(data.contents);
//   //     commit("SET_RESULTS", { query, results });
//   //     // console.log(results);
//   //     return results;
//   //   } catch (error) {
//   //     console.error("Error fetching results:", error);
//   //     throw error; // Opsional: Lempar error agar komponen tahu ada masalah
//   //   }
//   // },
//   async fetchResults({ state, commit }, query) {},
// };

// const getters = {
//   results: (state) => state.results,
// };

// export default {
//   namespaced: true,
//   state,
//   mutations,
//   actions,
//   getters,
// };
