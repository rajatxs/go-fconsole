<script setup>
import { defineProps } from 'vue';

const props = defineProps({
   width: {
      type: Number,
      default: 600, 
   },
   height: {
      type: Number,
      default: 600,
   },
   message: {
      type: String,
      default: 'Loading'
   }
});
</script>

<template>
   <v-sheet
      class="d-flex flex-column align-center justify-center flex-wrap text-center mx-auto px-4"
      :width="props.width"
      :height="props.height"
      :elevation="0">
      <v-progress-circular indeterminate class="mb-5"></v-progress-circular>
      <p class="font-weight-bold">{{ props.message }}</p>
   </v-sheet>
</template>
