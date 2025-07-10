<script setup>
import ContractDeployer from '@/components/ContractDeployer.vue';

// import EthereumSetting from './Client_Ethereum.vue';
// ContractDeployer

</script>


<template>
    
    <!-- tempat mendeploy smartcontract -->
    <div class="mt-4">
        <!-- <EthereumSetting/> -->
        <!-- <h3>Deploy smartcontract</h3> -->
        <div class="grid grid-cols-1 md:grid-cols-2 md:gap-4 mb-4">
           <div>
               <label for="" class="block">Private Key</label>
               <input-text />
           </div>
           <div>
            
           </div>
            
        </div>
        <!-- <EthereumSetting/> -->
         <contract-deployer/>
         
        <div>
            
        </div>
    </div>
</template>