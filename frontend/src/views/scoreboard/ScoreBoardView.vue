<script setup lang="ts">
import { onMounted, ref } from "vue";
import { GetScoreBoard } from "../../../wailsjs/go/scoreboard/Scoreboard";
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
        <div class="scoreboard">
            <TextInput v-model="scoreBoard.RoundName" />
            <NumberInput
                v-model="scoreBoard.BestOf"
                label="Best Of"
                class="input-field"
            />
            <div class="players">
                <PlayerPart
                    v-for="(player, i) in scoreBoard.Players"
                    :key="i"
                    v-model="scoreBoard.Players[i]"
                    :index="i"
                />
            </div>
        </div>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: fit-content;
}

.scoreboard {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.players {
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
}
</style>
