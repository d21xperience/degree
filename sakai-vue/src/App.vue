<script setup>
import LoadingOverlay from '@/components/LoadingOverlay.vue';
import { isLoading } from '@/router';
</script>

<template>
    <div>
        <LoadingOverlay :visible="isLoading"> Memuat data, harap tunggu... </LoadingOverlay>
        <router-view />
    </div>
</template>

<style scoped></style>
