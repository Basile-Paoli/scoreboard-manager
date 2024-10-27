<script setup lang="ts">
const { label } = defineProps<{ label?: string }>();
const model = defineModel<number>();

function increment() {
    if (!model.value) {
        model.value = 0;
    }
    model.value++;
}

function decrement() {
    if (!model.value) {
        model.value = 0;
    }
    model.value--;
}
</script>

<template>
    <div class="number-input-container">
        <label v-if="label">{{ label }}</label>
        <input v-model.number="model" type="number" />
        <div class="buttons">
            <button @click="increment">
                <span class="up-button svg" />
            </button>
            <button @click="decrement">
                <span class="down-button svg" />
            </button>
        </div>
    </div>
</template>

<style scoped>
.number-input-container {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;

    > input {
        font-size: 29pt;
        width: 2.5rem;
        text-align: center;
    }

    > .buttons {
        display: flex;
        flex-direction: column;

        > button {
            &:hover {
                background-color: var(--grey3);
            }

            &:active {
                background-color: var(--grey5);
            }

            > .svg {
                display: block;
                background-color: var(--grey8);
                width: 2rem;
                height: 1.5rem;

                &.up-button {
                    -webkit-mask: url(assets/svg/chevron-up.svg) no-repeat 50%
                        50%;
                    mask: url(assets/svg/chevron-up.svg) no-repeat 50% 50%;
                    mask-size: 2.5rem;
                }

                &.down-button {
                    -webkit-mask: url(../../assets/svg/chevron-down.svg)
                        no-repeat 50% 50%;
                    mask: url(../../assets/svg/chevron-down.svg) no-repeat 50%
                        50%;
                    mask-size: 2.5rem;
                }
            }
        }
    }
}
</style>
