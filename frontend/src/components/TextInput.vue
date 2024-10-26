<script setup lang="ts">
const model = defineModel<string>();
const { label } = defineProps<{ label?: string }>();
</script>

<template>
    <div class="text-input-container">
        <label v-if="label">{{ label }}</label>
        <input v-model="model" type="text" />
    </div>
</template>

<style scoped>
.text-input-container {
    display: flex;
    align-items: center;
    justify-content: center;
}

input {
    text-align: center;
}
</style>
