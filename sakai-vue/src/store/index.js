import { createStore } from 'vuex'; 
import authModule from './authService';
import scService from './scService';
import sekolahModule from './sekolahService';
const store = createStore({
    modules: { 
        authService: authModule,
        sekolahService: sekolahModule,
        scService: scService
    }
});

export default store;
