import { createStore } from 'vuex';
import authService from './auth_store/authService';
import anggotaKelas from './sekolah_store/anggotaKelas';
import dnsService from './sekolah_store/dnsService';
import guruService from './sekolah_store/guruService';
import kelasService from './sekolah_store/kelasService';
import kurikulumService from './sekolah_store/kurikulumService';
import nilaiService from './sekolah_store/nilaiService';
import sekolahService from './sekolah_store/sekolahService';
import semesterService from './sekolah_store/semesterService';
import siswaService from './sekolah_store/siswaService';
import scService from './smartcontract_store/scService';
const store = createStore({
    modules: {
        authService,
        scService,
        anggotaKelas,
        dnsService,
        guruService,
        sekolahService,
        semesterService,
        siswaService,
        kelasService,
        nilaiService,
        kurikulumService
    }
});

export default store;
