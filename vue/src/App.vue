<script setup>
import { RouterView } from "vue-router";
</script>

<template>
  <h1 class="green">JsonBin</h1>

  <RouterView />
</template>

<style>
@import "@/assets/base.css";

html,
body {
  height: 100%;
}

#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;

  font-weight: normal;
}

h1 {
  font-weight: 500;
  font-size: 2.5rem;
}

.green {
  color: #14b8a6;
}
</style>
