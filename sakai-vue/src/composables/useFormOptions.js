// composables/useFormOptions.js

import { debounce } from 'lodash-es';
import { computed, ref } from 'vue';
import { useStore } from 'vuex';
import { useSekolah } from './sekolah_composable/useSekolah';

export function useFormOptions() {
    const schemaname = computed(() => {
        return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    });
    const { sekolah } = useSekolah();
    const store = useStore();
    const selectedJenisKelamin = ref();
    const jenisKelaminOptions = ref([
        { label: 'Laki-Laki', value: 'L' },
        { label: 'Perempuan', value: 'P' }
    ]);
    // const statusAnakOptions = ref([
    //     { label: 'Anak Kandung', value: '1' },
    //     { label: 'Anak tiri', value: '2' },
    //     { label: 'Anak Angkat', value: '3' }
    // ]);

    const selectedAgama = ref();
    const agamaOptions = ref([
        { label: 'Islam', value: 'Islam' },
        { label: 'Kristen', value: 'Kristen' },
        { label: 'Katolik', value: 'Katolik' },
        { label: 'Hindu', value: 'Hindu' },
        { label: 'Buddha', value: 'Buddha' },
        { label: 'Konghucu', value: 'Konghucu' }
    ]);
    const selectedGelarAkademikDepan = ref();
    const selectedGelarAkademikBelakang = ref();
    const gelarAkademikOptions = ref();
    const gelarAkademikDepanOptions = ref();
    const gelarAkademikBelakangOptions = ref();
    const kurikulumList = computed(() => store.getters['kurikulumService/getKurikulum'] || []);

    const kurikulumOptions = ref();
    const kurikulumLoading = ref(false);
    const jurusanOptions = ref();
    const ptkLoading = ref(false);
    const ptkOptions = ref([]);
    const fetchGelarAkademik = async () => {
        try {
            let cek = store.getters['sekolahService/getGelarAkademik'];
            if (!cek || cek.length == 0) {
                cek = await store.dispatch('sekolahService/fetchGelarAkademik');
            }
            // console.log(cek);
            gelarAkademikOptions.value = cek;
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan gelar akademik', error);
        }
    };
    // const handleKeydown = (event) => {
    //     if (event.key === ' ') {
    //         searchTerm.value += ' '; // Menambahkan spasi ke query
    //     }
    // };
    const searchGelar = (posisiGelar, searcTerm) => {
        setTimeout(() => {
            if (!searcTerm.query.trim().length) {
                if (posisiGelar == 1) {
                    gelarAkademikDepanOptions.value = gelarAkademikOptions.value.filter((item) => item.posisiGelar == posisiGelar);
                } else if (posisiGelar == 2) {
                    gelarAkademikBelakangOptions.value = gelarAkademikOptions.value.filter((item) => item.posisiGelar == posisiGelar);
                }
            } else {
                if (posisiGelar == 1) {
                    gelarAkademikDepanOptions.value = gelarAkademikOptions.value.filter((item) => item.kode.toLowerCase().includes(searcTerm.query.toLowerCase()));
                } else if (posisiGelar == 2) {
                    gelarAkademikBelakangOptions.value = gelarAkademikOptions.value.filter((item) => item.kode.toLowerCase().includes(searcTerm.query.toLowerCase()));
                }
            }
        }, 250);
    };

    const fetchKurikulum = async () => {
        try {
            // let response = kurikulumList; //store.getters['sekolahService/getKurikulum'];
            if (!kurikulumList.value || kurikulumList.value.length == 0) {
                const { kurikulum } = await store.dispatch('kurikulumService/fetchKurikulum', { jenjangPendidikanId: sekolah.value?.sekolah.jenjangPendidikanId, jenjangPendidikanStr: sekolah.value?.bentukPendidikanStr });
                // if (response.value.status) {
                //     toast.add({ severity: 'success', summary: 'Successful', detail: `${response.value.message}`, life: 3000 });
                //     return response.value.kurikulum;
                // }
                kurikulumList.value = kurikulum;
                return kurikulum;
            } else {
                return kurikulumList.value;
            }
        } catch (error) {
            console.log(error);
            throw new Error('Gagal mendapatkan kurikulum', error);

            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };
    const searchKurikulum = debounce(async (searchTerm) => {
        if (!kurikulumList.value) {
            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan kurikulum`, life: 3000 });
            return;
        }
        if (!searchTerm) {
            kurikulumOptions.value = [...kurikulumList];
        } else {
            kurikulumOptions.value = kurikulumList.value.filter((item) => item.namaKurikulum.toLowerCase().includes(searchTerm.toLowerCase()));
        }
    }, 250);

    // ==============================
    // Bidang Keahlian
    const fetchBidangKeahlian = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchBidangKeahlian', { jurusanInduk: jurusanInduk });
            if (response.status) {
                // toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.bidangKeahlian;
            }
        } catch (error) {
            throw new Error('Gagal mendapatkan bidang keahlian:', error);

            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };
    const fetchProgramKeahlian = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchProgramKeahlian', { jurusanInduk: jurusanInduk });
            if (response.status) {
                // toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.programKeahlian;
            }
        } catch (error) {
            throw new Error('Gagal mendapatkan program keahlian: ', error);

            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };

    const fetchJurusan = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchJurusan', { jurusanInduk: jurusanInduk });
            if (response.status) {
                // toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.jurusan;
            }
        } catch (error) {
            throw new Error('Gagal mendapatkan jurusan:', error);

            // toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };

    const ptkSearch = debounce(async (searchTerm) => {
        try {
            ptkLoading.value = true;
            ptkOptions.value = await store.dispatch('sekolahService/searchPTKByName', { schemaname: schemaname.value, nama: searchTerm.query.toLowerCase() });
        } catch (error) {
            alert('error');
        } finally {
            ptkLoading.value = false;
        }
    }, 250);

    // const fetchMapelKurikulum = debounce(async () => {
    //     try {
    //         const response = await store.dispatch('sekolahService/fetchMapelKurikulum', { schemaname: schemaname.value, nama: searchTerm.query.toLowerCase() });
    //     } catch (error) {
    //         alert('error');
    //     } finally {
    //         ptkLoading.value = false;
    //     }
    // }, 250);

    return {
        selectedJenisKelamin,
        jenisKelaminOptions,
        selectedAgama,
        agamaOptions,
        selectedGelarAkademikDepan,
        selectedGelarAkademikBelakang,
        fetchGelarAkademik,
        searchGelar,
        gelarAkademikDepanOptions,
        gelarAkademikBelakangOptions,
        // handleKeydown,
        kurikulumOptions,
        kurikulumList,
        kurikulumLoading,
        searchKurikulum,
        fetchKurikulum,
        jurusanOptions,
        fetchJurusan,
        ptkSearch,
        ptkOptions,
        ptkLoading,
        fetchBidangKeahlian,
        fetchProgramKeahlian
    };
}
