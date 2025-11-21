<script setup>
import TopbarWidget from './TopbarWidget.vue';
</script>

<template>
    <nav class="navigation-bar" role="navigation" aria-label="Main navigation">
        <div class="nav-container">
            <TopbarWidget />
        </div>
    </nav>
</template>

<style scoped>
.navigation-bar {
    /* position: absolute; */
    /* top: 0;
    left: 0;
    right: 0;
    z-index: 50; */
    /* background: transparent; */
    /* background-color: blue; */
}

.nav-container {
    padding: 1.5rem 1.5rem;
    margin: 0 auto;
    max-width: 1200px;
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>
