import { createRouter, createWebHashHistory } from "vue-router";
import ScoreBoardView from "@/views/scoreboard/ScoreBoardView.vue";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "home",
            component: ScoreBoardView,
        },
    ],
});

export default router;
