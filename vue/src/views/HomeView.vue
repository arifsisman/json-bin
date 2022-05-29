<script setup>
import JsonEditor from "../components/JsonEditor.vue";

let jsonString = undefined;

function save() {
  console.log(jsonString);
}

function onTextChange(value) {
  console.log(value);
  jsonString = value;
}
</script>

<template>
  <JsonEditor
    :jsonString="jsonString"
    @onTextChange="onTextChange"
  ></JsonEditor>
  <Button label="Save" class="save-button" @click="save" />
</template>

<style scoped>
.save-button {
  float: right;
}
</style>
