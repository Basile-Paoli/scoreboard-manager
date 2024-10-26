<script setup lang="ts">
import NumberInput from "@/components/NumberInput.vue";
import TextInput from "@/components/TextInput.vue";
import {
    SetPlayerName,
    SetScore,
} from "../../../wailsjs/go/scoreboard/Scoreboard";
import { scoreboard } from "../../../wailsjs/go/models";
import Player = scoreboard.Player;

const props = defineProps<{ index: number }>();
const player = defineModel<Player>({ required: true });

function handleScoreUpdate() {
    SetScore(props.index, player.value.Score);
}

function handleNameUpdate() {
    SetPlayerName(props.index, player.value.Name);
}
</script>

<template>
    <div>
        <TextInput
            v-model="player.Name"
            class="input-field"
            @update:model-value="() => handleNameUpdate()"
        />
        <NumberInput
            v-model.number="player.Score"
            class="input-field"
            @update:model-value="() => handleScoreUpdate()"
        />
    </div>
</template>

<style scoped></style>
