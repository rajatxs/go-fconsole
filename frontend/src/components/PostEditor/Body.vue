<script setup>
import {defineEmits, onMounted} from 'vue';

const emit = defineEmits(['mount']);

onMounted(function () {
   emit('mount');
});
</script>

<template>
   <v-card flat>
      <div id="codex-editor"></div>
   </v-card>
</template>

<style></style>
