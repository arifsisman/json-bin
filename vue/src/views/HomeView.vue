<script setup>
import JsonEditor from "../components/JsonEditor.vue";

let jsonString = undefined;

function save() {
  console.log(jsonString);
}

function onTextChange(value) {
  jsonString = value;
  console.log(jsonString);
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
