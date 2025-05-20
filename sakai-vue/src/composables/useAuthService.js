// import { computed, ref } from 'vue';
// import { useStore } from 'vuex';

// export function useSekolahService() {
//     const store = useStore();
//     // const siswaList = ref([]);
//     const kelasList = ref([]);
//     const schemaname = computed(() => store.getters['sekolahService/getTabeltenant']?.schemaname);
//     const selectedSemester = computed(() => store.getters['sekolahService/getSelectedSemester']);
//     const fetchTabeltenant = async () => {};
//     const fetchUser = async () => {
//         try {
//             const payload = {
//                 tahunAjaranId: selectedSemester.value?.tahunAjaranId,
//                 schemaname: schemaname.value
//             };
//             let res = await store.getters['sekolahService/getPTKTerdaftar'];
//             if (!res) {
//                 res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
//             } else {
//                 if (res.tahun_ajaran_id != selectedSemester.value?.tahunAjaranId) {
//                     res = await store.dispatch('sekolahService/fetchPTKTerdaftar', payload);
//                 }
//             }
//             guruTerdaftarList.value = res.ptkTerdaftar;
//             return res.ptkTerdaftar;
//         } catch (error) {
//             console.error('Gagal mengambil data guru:', error);
//         }
//     };

//     return {
//         currentUser
//     };
// }
