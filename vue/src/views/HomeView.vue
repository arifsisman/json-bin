<script setup>
import JsonEditor from "../components/JsonEditor.vue"
import { onMounted, ref } from "vue"
import api from "../client/axios"
import { useToast } from "primevue/usetoast"

const toast = useToast()
toast.add({ severity: "info", summary: "Info Message" })

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
    .then((res) => {
      console.log(res)
      toast.add({
        severity: "success",
        summary: "JsonBin created!",
        detail: "You can view it at: " + res.data,
        life: 3000
      })
    })
    .catch((err) => {
      console.log(err)
      toast.add({
        severity: "error",
        summary: "Whoops!",
        summary: err.response.data,
        life: 3000
      })
    })
}
</script>

<template>
  <JsonEditor :json="json" :isValidJson="isValidJson" @onTextChange="onTextChange"></JsonEditor>
  <Button label="Save" class="save-button" @click="newBin"></Button>
  <Toast position="bottom-center" />
  <span>json: {{ json }}</span> <br />
  <span>isValid: {{ isValidJson }}</span>
</template>

<style scoped>
.save-button {
  float: right;
}
</style>
