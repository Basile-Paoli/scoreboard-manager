import pluginVue from "eslint-plugin-vue";
import prettier from "@vue/eslint-config-prettier";
import vueEslintConfigTypescript from "@vue/eslint-config-typescript";

export default [
    { ignores: ["wailsjs/**/*"] },
    ...pluginVue.configs["flat/recommended"],
    ...vueEslintConfigTypescript(),
    prettier,
];
