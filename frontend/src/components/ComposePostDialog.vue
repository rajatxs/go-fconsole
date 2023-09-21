<script setup>
import {defineProps, defineEmits, ref, onMounted, watch} from 'vue';
import { computeSlug } from '../utils';
import { getPublicTopics } from '../utils/topic';
import FAB from './FAB.vue';

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
});
const emit = defineEmits(['open', 'close', 'saved']);

/** @type {import('vue').Ref<string>} */
const title = ref('');

/** @type {import('vue').Ref<string>} */
const slug = ref('');

/** @type {import('vue').Ref<string|null>} */
const topic = ref(null);

/** @type {import('vue').Ref<string>} */
const desc = ref('');

/** @type {import('vue').Ref<string[]>} */
const tags = ref([]);

const steppers = ref(['Metadata', 'Body']);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topics = ref([]);

function close() {
   emit('close');
}

async function savePost() {
   emit('saved');
}

function renderTopics() {
   /** @type {any[]} */
   const _topics = getPublicTopics();

   topics.value = _topics.map(_topic => {
      return {
         title: _topic.name,
         value: _topic.id,
      }
   });
}

watch(title, function(newTitle) {
   if (newTitle.length > 0) {
      slug.value = computeSlug(newTitle);
   } else {
      slug.value = '';
   }
});

onMounted(function() {
   renderTopics();
});
</script>

<template>
   <v-dialog v-model="props.visible" fullscreen :scrim="false" scrollable transition="dialog-bottom-transition">
      <template v-slot:activator="{ props }">
         <FAB>
            <v-btn 
               fab 
               fixed 
               size="large" 
               color="secondary" 
               icon="mdi-pencil" 
               elevation="8" 
               @click="emit('open')">
            </v-btn>
         </FAB>
      </template>
      <v-card>
         <v-toolbar dark color="primary" class="app-custom-toolbar">
            <v-btn icon dark @click="close">
               <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Compose</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
               <v-btn variant="text" @click="savePost">Done</v-btn>
            </v-toolbar-items>
         </v-toolbar>

         <v-form class="mt-16">
            <v-container>
               <v-stepper :items="steppers">
                  <template v-slot:item.1>
                     <v-card flat>
                        <v-text-field v-model="title" label="Title"></v-text-field>
                        <v-row>
                           <v-col cols="12" sm="6">
                              <v-select
                                 v-model="topic"
                                 label="Topic"
                                 :items="topics">
                              </v-select>
                           </v-col>
                           <v-col cols="12" sm="6">
                              <v-text-field 
                                 v-model="slug" 
                                 :title="slug"
                                 label="Slug (Auto Generated)" 
                                 readonly></v-text-field>
                           </v-col>
                        </v-row>
                        <v-textarea v-model="desc" label="Description"></v-textarea>

                        <v-combobox
                           v-model="tags"
                           label="Tags"
                           multiple
                           chips>
                        </v-combobox>

                        <v-row class="mb-2" style="min-height: 240px;">
                           <v-col cols="12" sm="6">
                              <!-- <v-img
                                 aspect-ratio="16/9"
                                 src="https://cdn.vuetifyjs.com/images/parallax/material.jpg"
                                 cover>
                              </v-img> -->
                              <v-sheet class="d-flex flex-column justify-center align-center" height="100%" border="md" :rounded="true">
                                 <v-btn variant="text">Upload Cover Image</v-btn>
                              </v-sheet>
                           </v-col>
                           <v-col cols="12" sm="6">
                              <v-text-field label="Reference Name (optional)"></v-text-field>
                              <v-text-field label="Reference URL (optional)"></v-text-field>
                           </v-col>
                        </v-row>
                     </v-card>
                  </template>
                  
                  <template v-slot:item.2>
                     <v-card title="Body" flat>...</v-card>
                  </template>
               </v-stepper>
            </v-container>
         </v-form>
      </v-card>
   </v-dialog>
</template>

<style scoped>
.app-custom-toolbar {
   position: fixed;
   z-index: 120;
}
</style>
