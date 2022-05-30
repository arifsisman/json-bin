<script setup>
import { ref } from "vue";

const emit = defineEmits(["onTextChange"]);
const props = defineProps(["jsonString"]);
const json = ref(props.jsonString);

let isInvalidJson = false;
// let json = "1111";

function formatPaste(event) {
  if (event.inputType === "insertText") {
    emit("onTextChange", event.target.text);
    return;
  }

  // paste
  const clipboardData = event.target.value;
  console.log(clipboardData);

  try {
    // json = JSON.stringify(JSON.parse(clipboardData), null, 2);

    isInvalidJson = false;
    emit("onTextChange", clipboardData);
  } catch (e) {
    console.log("Invalid JSON");
    isInvalidJson = true;
  }
}
</script>

<template>
  <div class="main-area">
    <Textarea
      v-model="json"
      class="text-area"
      rows="30"
      :class="{ invalid: isInvalidJson }"
      @input="formatPaste"
    />
  </div>
</template>

<style scoped>
.text-area {
  width: 100%;
  font-family: monospace;
  background: hsl(0, 0%, 98%);
}

.invalid {
  border: 10px solid red;
  border-color: red;
}
</style>
