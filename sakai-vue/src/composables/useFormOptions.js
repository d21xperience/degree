// composables/useFormOptions.js

import { useToast } from 'primevue';
import { ref } from 'vue';
import { useStore } from 'vuex';
import { useSekolahService } from './useSekolahService';
export function useFormOptions() {
    const { sekolah } = useSekolahService();
    const store = useStore();
    const toast = useToast();
    const tahunAjaranOptions = ref();
    const selectedJenisKelamin = ref();
    const jenisKelaminOptions = ref([
        { label: 'Laki-Laki', value: 'L' },
        { label: 'Perempuan', value: 'P' }
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
    const jurusanList = store.getters['sekolahService/getJurusan'];
    const jurusanOptions = ref();
    const jurusanLoading = ref(false);

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
            let response = await store.getters['sekolahService/getKurikulum'];
            if (!response || response.length == 0) {
                console.log(response);
                response = await store.dispatch('sekolahService/fetchKurikulum', sekolah.value?.jenjangPendidikanId);
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

    const kurikulumSearch = (searchTerm) => {
        kurikulumLoading.value = true;
        setTimeout(async () => {
            if (!searchTerm.query.trim().length) {
                kurikulumOptions.value = await kurikulumList;
            } else {
                kurikulumOptions.value = await kurikulumList.filter((item) => item.namaKurikulum.toLowerCase().includes(searchTerm.query.toLowerCase()));
            }
            kurikulumLoading.value = false;
        }, 250);
    };

    const fetchJurusan = async () => {
        try {
            let response = await store.getters['sekolahService/getJurusan'];

            if (!response || response.length == 0) {
                let param = '';
                // console.log(sekolah.value.bentukPendidikanId);
                switch (sekolah.value.bentukPendidikanId) {
                    case 15 || 17:
                        param = 'untuk_smk';
                        break;
                    case 13 || 16:
                        param = 'untuk_sma';
                        break;

                    default:
                        break;
                }
                // console.log(param);
                const payload = {
                    jenjang_pendidikan_id: sekolah.value?.jenjangPendidikanId,
                    param: param
                };
                response = await store.dispatch('sekolahService/fetchJurusan', payload);
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

    const jurusanSearch = (searchTerm) => {
        jurusanLoading.value = true;
        setTimeout(async () => {
            if (!searchTerm.query.trim().length) {
                jurusanOptions.value = await jurusanList;
            } else {
                jurusanOptions.value = await jurusanList.filter((item) => item.namaJurusan.toLowerCase().includes(searchTerm.query.toLowerCase()));
            }
            jurusanLoading.value = false;
        }, 250);
    };

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
        kurikulumLoading,
        kurikulumSearch,
        fetchKurikulum,
        jurusanOptions,
        jurusanLoading,
        jurusanSearch,
        fetchJurusan
    };
}
