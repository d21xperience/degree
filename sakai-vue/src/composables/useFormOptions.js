// composables/useFormOptions.js

import { debounce } from 'lodash-es';
import { useToast } from 'primevue';
import { ref } from 'vue';
import { useStore } from 'vuex';
import { useSekolahService } from './useSekolahService';

export function useFormOptions() {
    const { sekolah, schemaname } = useSekolahService();
    const store = useStore();
    const toast = useToast();
    const tahunAjaranOptions = ref();
    const selectedJenisKelamin = ref();
    const jenisKelaminOptions = ref([
        { label: 'Laki-Laki', value: 'L' },
        { label: 'Perempuan', value: 'P' }
    ]);
    const statusAnakOptions = ref([
        { label: 'Anak Kandung', value: '1' },
        { label: 'Anak tiri', value: '2' },
        { label: 'Anak Angkat', value: '3' }
    ]);

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
    const kurikulumList = store.getters['sekolahService/getKurikulum'];
    const kurikulumOptions = ref();
    const kurikulumLoading = ref(false);
    const jurusanOptions = ref();
    const ptkLoading = ref(false);
    const ptkOptions = ref([]);
    const fetchGelarAkademik = () => {
        try {
            let cek = store.getters['sekolahService/getGelarAkademik'];
            if (!cek || cek.length == 0) {
                cek = store.dispatch('sekolahService/fetchGelarAkademik');
            }
            // console.log(cek);
            gelarAkademikOptions.value = cek;
        } catch (error) {
            throw error;
        }
    };
    const handleKeydown = (event) => {
        if (event.key === ' ') {
            searchTerm.value += ' '; // Menambahkan spasi ke query
        }
    };
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
            let response = kurikulumList; //await store.getters['sekolahService/getKurikulum'];
            if (!response || response.length == 0) {
                response = await store.dispatch('sekolahService/fetchKurikulum', { jenjangPendidikanId: sekolah.value?.sekolah.jenjangPendidikanId, jenjangPendidikanStr: sekolah.value?.bentukPendidikanStr });
                if (response.status) {
                    toast.add({ severity: 'success', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                    return response.kurikulum;
                }
            }
            return response;
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };
    const searchKurikulum = debounce(async (searchTerm) => {
        if (!kurikulumList) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan kurikulum`, life: 3000 });
            return;
        }
        if (!searchTerm) {
            kurikulumOptions.value = [...kurikulumList];
        } else {
            kurikulumOptions.value = kurikulumList.filter((item) => item.namaKurikulum.toLowerCase().includes(searchTerm.toLowerCase()));
        }
    }, 250);

    // ==============================
    // Bidang Keahlian
    const fetchBidangKeahlian = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchBidangKeahlian', { jurusanInduk: jurusanInduk });
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.bidangKeahlian;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };
    const fetchProgramKeahlian = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchProgramKeahlian', { jurusanInduk: jurusanInduk });
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.programKeahlian;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
        }
    };

    const fetchJurusan = async (jurusanInduk) => {
        try {
            const response = await store.dispatch('sekolahService/fetchJurusan', { jurusanInduk: jurusanInduk });
            if (response.status) {
                toast.add({ severity: 'info', summary: 'Successful', detail: `${response.message}`, life: 3000 });
                return response.jurusan;
            }
        } catch (error) {
            toast.add({ severity: 'error', summary: 'Failled', detail: `Gagal mendapatkan informasi: ${error.message}`, life: 3000 });
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

    const fetchMapelKurikulum = debounce(async () => {
        try {
            const response = await store.dispatch('sekolahService/fetchMapelKurikulum', { schemaname: schemaname.value, nama: searchTerm.query.toLowerCase() });
        } catch (error) {
            alert('error');
        } finally {
            ptkLoading.value = false;
        }
    }, 250); 

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
        handleKeydown,
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
