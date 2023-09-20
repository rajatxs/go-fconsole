<script setup>
import {defineProps, defineEmits} from 'vue';
import FAB from './FAB.vue';

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
});
const emit = defineEmits(['open', 'close', 'saved'])

function close() {
   emit('close');
}

async function savePost() {
   emit('saved');
}
</script>

<template>
   <v-dialog v-model="props.visible" fullscreen :scrim="false" transition="dialog-bottom-transition">
      <template v-slot:activator="{ props }">
         <FAB>
            <v-btn 
               fab 
               fixed 
               size="large" 
               color="primary" 
               icon="mdi-pencil" 
               elevation="8" 
               @click="emit('open')">
            </v-btn>
         </FAB>
      </template>
      <v-card>
         <v-toolbar dark color="primary">
            <v-btn icon dark @click="close">
               <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Compose</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
               <v-btn variant="text" @click="savePost"> Save </v-btn>
            </v-toolbar-items>
         </v-toolbar>
         <v-divider></v-divider>
      </v-card>
   </v-dialog>
</template>
