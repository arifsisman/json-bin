module.exports = {
    env: {
      node: true,
    },
    extends: [
      'eslint:recommended',
      'plugin:vue/vue3-recommended',
      "prettier"
    ],
    rules: {
      "vue/no-reserved-component-names": "off",
      "vue/multi-word-component-names": "off",
      "vue/v-on-event-hyphenation": "off",
      "vue/attribute-hyphenation": "off",
      "vue/require-prop-types": "off",
      "vue/require-explicit-emits": "off",
    }
  }