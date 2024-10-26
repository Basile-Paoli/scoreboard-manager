<script setup lang="ts"></script>

<template>
    <div class="layout">
        <RouterView />
    </div>
</template>

<style scoped>
.layout {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
}
</style>
