<template>
    <div>
        <div class="">
            <div class="flex justify-between p-2">
                <div class="flex items-center space-x-2">
                    <label class="text-gray-500">Platform</label>
                    <button @click="dialogSelectplatforms = true" class="rounded-full bg-slate-300 py-2 px-4 hover:opacity-80">
                        <span v-if="platformsActivate">{{ platformsActivate?.name ?? 'Pilih Platform' }}</span>
                        <i class="ml-2 pi pi-angle-up"></i>
                    </button>
                </div>
                <!-- <div class="md:flex md:items-center md:space-x-2">
                    <label class="text-slate-500 md:text-base text-sm">Akun</label>
                    <div>
                        <Select v-model="selectedAkun" :options="AkunOptions" optionLabel="accountName" optionValue="accountID" placeholder="Pilih akun" class="w-full md:w-52 mr-2" />
                    </div>
                </div> -->
            </div>
        </div>
        <div class="">
            <RouterView />
        </div>
        <Dialog v-model:visible="dialogSelectplatforms" modal header="Pilih Platform" position="top">
            <div>
                <div v-for="platform in platforms" :key="platform.id" class="my-2 flex space-x-1">
                    <input type="radio" :name="platform.id" :id="platform.id" :value="platform.name" v-model="selectedPlatform" />
                    <label :for="platform.id">{{ platform?.name }}</label>
                </div>

                <div class="flex justify-center">
                    <button type="button" class="bg-blue-400 p-2 rounded-md text-white" @click="saveSelection">Simpan</button>
                </div>
            </div>
        </Dialog>
    </div>
</template>
<script setup>
import { ref } from 'vue';
const selectedPlatform = ref();
const selectedAkun = ref();
const dialogSelectplatforms = ref(false);
const platformsActivate = ref({
    id: 1,
    name: 'Ethereum'
});
const platforms = ref([
    {
        id: 1,
        name: 'Ethereum'
    },
    {
        id: 2,
        name: 'Quorum'
    },
    {
        id: 3,
        name: 'Hyperledger Fabric'
    }
]); 
const AkunOptions = ref([
    {
        accountID: '1',
        accountName: 'Account 1',
        accountAddress: '39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp',
        amount: 300,
        platformId: 1
    },
    {
        accountID: '2',
        accountName: 'Account 2',
        accountAddress: '9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU',
        amount: 100,
        platformId: 1
    }
]);
</script>
