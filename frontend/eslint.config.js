import pluginVue from "eslint-plugin-vue";
import prettier from "@vue/eslint-config-prettier";
import vueEslintConfigTypescript from "@vue/eslint-config-typescript";
import js from "@eslint/js";

export default [
    { ignores: ["wailsjs/**/*"] },
    ...pluginVue.configs["flat/recommended"],
    prettier,
    ...vueEslintConfigTypescript(),
    js.configs.recommended,
];
