<script setup>
import LoadingOverlay from '@/components/LoadingOverlay.vue';
import { isLoading } from '@/router';
</script>

<template>
    <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay>
    <router-view />
</template>

<style scoped></style>
