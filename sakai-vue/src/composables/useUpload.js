import { computed, onMounted, onUnmounted, ref } from 'vue';

/**
 * useSomething composable
 * Deskripsi singkat tentang apa yang dilakukan composable ini.
 *
 * @returns {
 *   state: object,
 *   actions: object,
 *   computed: object,
 * }
 */
export function useSomething() {
    // =============================================================
    // State
    // =============================================================
    const data = ref(null);

    // =============================================================
    // Computed
    // =============================================================
    const computedValue = computed(() => {
        return data.value;
    });

    // =============================================================
    // Actions / Methods
    // =============================================================
    function fetchData() {
        // implementasi
    }
    const downloadTemplateWithGet = async (templateUrl) => {
        // console.log(templateUrl.value);
        // return;
        // submitted.value = true;
        // if (!semesterAktif.value) {
        //     // alert('Pilih tahun pelajaran');
        //     isErr.value = true;
        //     return;
        // }
        try {
            const response = await fetch(templateUrl, {
                method: 'GET',
                headers: {
                    Accept: 'application/octet-stream'
                }
            });

            if (!response.ok) {
                throw new Error('Gagal mengunduh file');
            }

            // Coba ambil nama file dari header Content-Disposition
            const contentDisposition = response.headers.get('Content-Disposition');
            let fileName = 'downloaded_file.xlsx'; // Default jika tidak ditemukan
            if (contentDisposition) {
                const match = contentDisposition.match(/filename="([^"]+)"/);
                if (match && match[1]) {
                    fileName = match[1];
                }
            }

            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);

            const a = document.createElement('a');
            a.href = url;
            a.download = fileName;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);

            window.URL.revokeObjectURL(url);
        } catch (error) {
            console.log(error);
            // toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
        }
    };
    // BACKUP DAN SUDAH BERJALAN DENGAN BAIK. JIKA ADA MASALAH COPY PASTE KE KOMPONEN DialogImport.vue
    // const downloadTemplate = async () => {
    //     // console.log(templateUrl.value);
    //     // return;
    //     submitted.value = true;
    //     if (!semesterAktif.value) {
    //         // alert('Pilih tahun pelajaran');
    //         isErr.value = true;
    //         return;
    //     }
    //     try {
    //         const response = await fetch(templateUrl.value, {
    //             method: 'GET',
    //             headers: {
    //                 Accept: 'application/octet-stream'
    //             }
    //         });

    //         if (!response.ok) {
    //             throw new Error('Gagal mengunduh file');
    //         }

    //         // Coba ambil nama file dari header Content-Disposition
    //         const contentDisposition = response.headers.get('Content-Disposition');
    //         let fileName = 'downloaded_file.xlsx'; // Default jika tidak ditemukan
    //         if (contentDisposition) {
    //             const match = contentDisposition.match(/filename="([^"]+)"/);
    //             if (match && match[1]) {
    //                 fileName = match[1];
    //             }
    //         }

    //         const blob = await response.blob();
    //         const url = window.URL.createObjectURL(blob);

    //         const a = document.createElement('a');
    //         a.href = url;
    //         a.download = fileName;
    //         document.body.appendChild(a);
    //         a.click();
    //         document.body.removeChild(a);

    //         window.URL.revokeObjectURL(url);
    //     } catch (error) {
    //         console.log(error);
    //         toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
    //     }
    // };
    // =============================================================
    // Lifecycle Hooks
    // =============================================================
    onMounted(() => {
        fetchData();
    });

    onUnmounted(() => {
        // cleanup jika diperlukan
    });

    // =============================================================
    // Return API (Ordered, Clean)
    // =============================================================
    return {
        // state
        data,

        // computed
        computedValue,

        // methods
        fetchData,
        downloadTemplateWithGet
    };
}
