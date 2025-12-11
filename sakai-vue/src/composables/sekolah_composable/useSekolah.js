import { computed, onMounted, readonly, ref, watch } from 'vue';
import { useStore } from 'vuex';

/**
 * Composable untuk sekolah: read + edit sederhana.
 * currentSekolah bisa diubah langsung (mutable ref), dan tetap sinkron saat store berubah.
 */
export function useSekolah({ schemaname = '', autoload = true } = {}) {
    // console.log('[COMPOSABLES] useSekolah dipanggil!!', schemaname);
    const store = useStore();
    // const schemaname = computed(() => {
    //     return store.getters['sekolahService/getTabeltenant']?.schemaname || null;
    // });

    const isFetching = ref(false);
    const isError = ref(null);
    const storeSekolah = computed(() => store.getters['sekolahService/getSekolah'] || null);

    // 🟢 Local editable copy — satu-satunya sumber untuk tampilan & edit
    const currentSekolah = ref(null);

    // 🔁 Isi awal dari store, dan update hanya jika currentSekolah masih null
    // → artinya: sekali isi, lalu biarkan komponen kelola perubahan lokal
    // watch(
    //     storeSekolah,
    //     (newVal) => {
    //         if (newVal && currentSekolah.value === null) {
    //             // Clone agar tidak shared reference
    //             currentSekolah.value = JSON.parse(JSON.stringify(newVal));
    //         }
    //     },
    //     { immediate: true } // ✅ Wajib: isi saat composable dipanggil
    // );

    watch(
        storeSekolah,
        (newVal) => {
            if (newVal) {
                currentSekolah.value = JSON.parse(JSON.stringify(newVal));
            }
        },
        { immediate: true }
    );
    const _fetchSekolah = async () => {
        if (!schemaname) throw new Error('schemaname belum tersedia');

        // return await store.dispatch('sekolahService/fetchSekolah', {
        //     schemaname: schemaname
        // });
        const tes = await store.dispatch('sekolahService/fetchSekolah', {
            schemaname: schemaname
        });
        return tes;
    };

    const initialize = async () => {
        if (storeSekolah.value) return; // sudah ada → skip
        try {
            isFetching.value = true;
            await _fetchSekolah();
            // Setelah fetch, store berubah → watch akan isi currentSekolah (karena masih null)
        } catch (err) {
            isError.value = err;
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    if (schemaname == '') {
        return new Error('Skema tidak boleh kosong');
    }
    if (autoload) {
        onMounted(() => {
            initialize().catch((err) => {
                console.warn('[useSekolah] Gagal autoload:', err.message);
            });
        });
    }

    const updateSekolah = async (param) => {
        try {
            isFetching.value = true;
            const payload = { ...param, schemaname: schemaname };
            const result = await store.dispatch('sekolahService/updateSekolah', payload);

            return result;
        } catch (err) {
            isError.value = err;
            throw err;
        } finally {
            isFetching.value = false;
        }
    };

    const fetchTingkat = async () => {
        const jenjangId = currentSekolah.value?.sekolah?.jenjangPendidikanId;
        if (!jenjangId) throw new Error('jenjangPendidikanId tidak ditemukan');

        return await store.dispatch('sekolahService/fetchTingkatPendidikan', {
            jenjang_pendidikan_id: jenjangId
        });
    };

    const createInfoIjazah = async (data) => {
        if (!schemaname) throw new Error('schemaname belum tersedia');
        return await store.dispatch('sekolahService/createInfoIjazah', {
            schemaname: schemaname,
            info_ijazah: data
        });
    };

    return {
        // 🔹 State
        currentSekolah, // ✅ mutable — langsung pakai di v-model
        isFetching: readonly(isFetching),
        isError: readonly(isError),

        // 🔹 Actions
        initialize,
        updateSekolah,
        fetchTingkat,
        createInfoIjazah
    };
}
