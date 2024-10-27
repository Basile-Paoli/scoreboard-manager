<script setup lang="ts">
const model = defineModel<string>();
const { label } = defineProps<{ label?: string }>();
const emit = defineEmits<{ blur: [] }>();
</script>

<template>
    <div class="text-input-container">
        <label v-if="label">{{ label }}</label>
        <div class="input-div">
            <input v-model="model" type="text" @blur="emit('blur')" />
        </div>
    </div>
</template>

<style scoped>
.text-input-container {
    display: flex;
    align-items: center;
    justify-content: center;

    .input-div {
        input {
            text-align: center;
            width: 100%;
            margin: 0;
        }
    }
}
</style>
