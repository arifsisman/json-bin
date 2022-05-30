<script setup>
import JsonEditor from "../components/JsonEditor.vue"
import { onMounted, ref } from "vue"
import api from "../client/axios"

const json = ref('{"a":"aaas"}')
const isValidJson = ref(false)

onMounted(() => validateJson())

const onTextChange = (text) => {
  json.value = text
  validateJson()
}

const validateJson = () => {
  try {
    JSON.parse(json.value)
    isValidJson.value = true
  } catch (e) {
    isValidJson.value = false
  }
}

const newBin = () => {
  api
    .post("/", json.value)
    .then((res) => console.log(res))
    .catch((err) => console.error(err))
}
</script>

<template>
  <JsonEditor :json="json" @onTextChange="onTextChange"></JsonEditor>
  <Button label="Save" class="save-button" :disabled="!isValidJson" @click="newBin"></Button>
  <span>json: {{ json }}</span> <br />
  <span>isValid: {{ isValidJson }}</span>
</template>

<style scoped>
.save-button {
  float: right;
}
</style>
