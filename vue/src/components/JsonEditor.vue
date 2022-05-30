<script setup>
const props = defineProps({
  json: {
    type: String,
    default: "",
    required: true
  },
  isValidJson: {
    type: Boolean,
    default: false
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
      :class="props.isValidJson ? 'valid' : 'invalid'"
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
/* .valid {
  border: 2px solid #14b8a6;
}

.invalid {
  border: 2px solid crimson;
} 

.p-inputtext:enabled:hover,
.p-inputtext:enabled:focus {
  border-color: initial;
  box-shadow: initial;
} */
</style>
