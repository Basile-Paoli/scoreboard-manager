<script setup lang="ts">
import { computed, onMounted, ref } from "vue";

const { message } = defineProps<{ message: string }>();
const isOpen = ref(true);

function close() {
    isOpen.value = false;
    setTimeout(() => {
        emit("close");
    }, 300);
}

onMounted(() => {
    isOpen.value = true;
});

const displayMessage = computed(() => {
    const capitalized = message.charAt(0).toUpperCase() + message.slice(1);
    return `Error: ${capitalized}`;
});

const emit = defineEmits<{ close: [] }>();
</script>

<template>
    <div class="error" :class="{ isOpen }" @click="close">
        {{ displayMessage }}
    </div>
</template>

<style scoped>
.error {
    width: 15rem;
    font-size: 1rem;
    background-color: #fa3f3f;
    color: white;
    padding: 1rem;
    text-align: center;
    border-radius: 0.5rem;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    opacity: 0;
    transform: translateY(50%);
    transition:
        opacity 0.3s ease-in-out,
        transform 0.3s ease-in-out;

    &.isOpen {
        opacity: 1;
        cursor: pointer;
        transform: translateY(0);
    }
}
</style>
