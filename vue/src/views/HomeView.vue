<script setup>
import JsonEditor from "../components/JsonEditor.vue"
import { onMounted, ref, watch } from "vue"
import axios from "axios"

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

const newBin = async () => {
  try {
    const body = JSON.stringify(json)
    const response = await axios.post("http://localhost:3000/", body)
    if (response.status === 200) {
      console.log(response.data)
    }
  } catch (error) {
    console.error(error)
  }
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
