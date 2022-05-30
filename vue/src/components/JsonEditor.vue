<script setup>
const props = defineProps({
  json: {
    type: String,
    default: "",
    required: true
  }
})

const emit = defineEmits(["onTextChange", "onJsonValidate"])
</script>

<template>
  <div class="main-area">
    <Textarea
      :value="props.json"
      @input="emit('onTextChange', $event.target.value)"
      class="text-area"
      rows="30"
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
  border: 1px solid red;
}
</style>
