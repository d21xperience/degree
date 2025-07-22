<script setup>
import { onBeforeMount } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

onBeforeMount(async () => {
    console.log('bootstrap');
    await store.dispatch('authService/bootstrap');
});
</script>

<template>
    <router-view />
</template>

<style scoped></style>
