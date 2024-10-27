<script setup lang="ts">
import { onMounted, ref } from "vue";
import {
    GetScoreBoard,
    SetBestOf,
    SetRoundName,
} from "../../../wailsjs/go/scoreboard/Scoreboard";
import TextInput from "@/components/TextInput.vue";
import NumberInput from "@/components/NumberInput.vue";
import PlayerPart from "@/views/scoreboard/PlayerPart.vue";
import { scoreboard } from "../../../wailsjs/go/models";
import Scoreboard = scoreboard.Scoreboard;

const scoreBoard = ref<Scoreboard>(Scoreboard.createFrom({}));
onMounted(async () => {
    scoreBoard.value = await GetScoreBoard();
});
</script>
<template>
    <div class="container">
        <h1>Scoreboard</h1>
        <TextInput
            v-model="scoreBoard.RoundName"
            @blur="() => SetRoundName(scoreBoard.RoundName)"
        />
        <NumberInput
            v-model="scoreBoard.BestOf"
            label="Best Of"
            @update:model-value="() => SetBestOf(scoreBoard.BestOf)"
        />
        <div class="players">
            <PlayerPart
                v-for="(_, i) in scoreBoard.Players"
                :key="i"
                v-model="scoreBoard.Players[i]"
                :index="i"
            />
        </div>
    </div>
</template>
<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
    max-width: fit-content;

    > .players {
        display: flex;
        justify-content: space-around;
        gap: 1rem;
        flex-wrap: wrap;
    }
}
</style>
