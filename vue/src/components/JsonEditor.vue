<script setup>
import { ref } from "vue";

defineEmits(["onTextChange"]);
const props = defineProps(["jsonString"]);
const json = ref(props.jsonString);
</script>

<template>
  <div class="main-area">
    <Textarea
      v-model="json"
      class="text-area"
      rows="30"
      @keypress="$emit('onTextChange', json)"
    />
  </div>
</template>

<style scoped>
.text-area {
  width: 100%;
  font-family: monospace;
  background: hsl(0, 0%, 98%);
}
</style>
