<script setup>
import { useSCService } from '@/composables/useSCService';
import { onMounted, ref } from 'vue';

const scService = useSCService();

const bcTransactions = ref();
onMounted(async () => {
    bcTransactions.value = await scService.getBCTransaction();
    console.log(bcTransactions.value);
});
// const fetchTransaksi = async () => {
//     try {
//         let payload = {
//             sekolah_id: await store.getters["sekolahService/getSekolah"]?.sekolah_id,
//             tahun_ajaran_id: tahunAjaranId.value
//         }
//         const results = await store.dispatch("scService/fetchIjazahBC", payload)
//         bcTransactions.value = results.degreeData
//         console.log(results)
//     } catch (error) {
//         toast.add({ severity: 'error', summary: 'Error', detail: error, life: 3000 });
//     }
// }

// watch(tahunAjaranId, async (newVal) => {
//     await fetchTransaksi()
// })
</script>
<template>
    <div class="p2 my-2">
        <div>
            <DataTable striped-rows :value="bcTransactions">
                <Column field="ijazah.nama" header="Nama" />
                <Column field="degreeHash" header="Degree Hash" />
                <Column field="txHash" header="Trx Hash" />
                <!-- <Column field="bcType" header="BC Type"></Column> -->
                <Column field="ipfsUrl" header="IPFS URL" />
                <Column field="ipfsUrl" header="BC URL" />
                <Column header="Created" />
            </DataTable>
        </div>
        <Toast />
    </div>
</template>
