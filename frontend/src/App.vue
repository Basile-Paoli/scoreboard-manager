<script setup lang="ts"></script>

<template>
    <div class="layout">
        <RouterView v-slot="{ Component }">
            <Suspense timeout="0">
                <component :is="Component" />
                <template #fallback>
                    <h1>Loading...</h1>
                </template>
            </Suspense>
        </RouterView>
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
