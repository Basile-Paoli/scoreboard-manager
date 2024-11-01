<script setup lang="ts">
import { onMounted, ref } from "vue";
import TextInput from "@/components/TextInput.vue";
import NumberInput from "@/components/NumberInput.vue";
import PlayerPart from "@/views/scoreboard/PlayerPart.vue";
import {
    GetScoreBoard,
    SaveEntireScoreboard,
    SetBestOf,
    SetRoundName,
} from "../../../wailsjs/go/scoreboard/Scoreboard";

const scoreBoard = ref(await GetScoreBoard());
onMounted(async () => {
    await SaveEntireScoreboard();
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
                v-for="i in scoreBoard.Players.length"
                :key="i"
                v-model="scoreBoard.Players[i - 1]"
                :index="i - 1"
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

    > .players {
        display: flex;
        justify-content: space-around;
        gap: 1rem;
    }
}
</style>
