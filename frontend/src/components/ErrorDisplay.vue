<script setup lang="ts">
import { onMounted, ref } from "vue";
import ErrorDialog from "@/components/ErrorDialog.vue";
import { EventsOn } from "../../wailsjs/runtime";
import { events } from "../../wailsjs/go/models";
import Event = events.Event;

const errors = ref<
    {
        message: string;
        key: number;
    }[]
>([]);

const addError = (error: string) => {
    errors.value = [
        ...errors.value,
        {
            message: error,
            key: Date.now(),
        },
    ];
};

const close = (key: number) => {
    errors.value = errors.value.filter((error) => error.key !== key);
};

defineExpose({
    addError,
});

onMounted(() => {
    EventsOn(Event.error, (data) => {
        addError(data);
    });
});
</script>

<template>
    <div class="error-display">
        <div v-for="error in errors" :key="error.key">
            <ErrorDialog
                :message="error.message"
                class="error-dialog"
                @close="close(error.key)"
            />
        </div>
    </div>
</template>

<style scoped>
.error-display {
    position: fixed;
    bottom: 11px;
    right: 11px;

    .error-dialog {
        margin-top: 0.5rem;
    }
}
</style>
