<script setup lang="ts">
import { onErrorCaptured, ref } from "vue";
import ErrorDisplay from "@/components/ErrorDisplay.vue";

const errorDisplayRef = ref();

onErrorCaptured((e: unknown) => {
    if (e instanceof Error) {
        errorDisplayRef.value?.addError(e.message);
    } else {
        errorDisplayRef.value?.addError(e);
    }
    return true;
});
</script>

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
    <ErrorDisplay ref="errorDisplayRef" />
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
