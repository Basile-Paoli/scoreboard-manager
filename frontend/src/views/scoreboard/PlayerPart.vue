<script setup lang="ts">
import NumberInput from "@/components/NumberInput.vue";
import TextInput from "@/components/TextInput.vue";
import {
    SetPlayerName,
    SetScore,
    SetTeamName,
} from "../../../wailsjs/go/scoreboard/Scoreboard";
import { scoreboard } from "../../../wailsjs/go/models";
import Player = scoreboard.Player;

const props = defineProps<{ index: number }>();
const player = defineModel<Player>({ required: true });

function scoreUpdate() {
    SetScore(props.index, player.value.Score);
}

function nameUpdate() {
    SetPlayerName(props.index, player.value.Name);
}

function teamNameUpdate() {
    SetTeamName(props.index, player.value.TeamName);
}
</script>

<template>
    <div class="player-part">
        <div class="name-group">
            <TextInput
                v-model="player.TeamName"
                class="team-input"
                @blur="teamNameUpdate"
            />
            <TextInput v-model="player.Name" @blur="nameUpdate" />
        </div>
        <NumberInput
            v-model.number="player.Score"
            @update:model-value="scoreUpdate"
        />
    </div>
</template>

<style scoped>
.player-part {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;

    > .name-group {
        display: flex;
        gap: 0.5rem;

        > .team-input {
            max-width: 100px;
        }
    }
}
</style>
