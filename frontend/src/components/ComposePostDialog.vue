<script setup>
import {defineProps, defineEmits, ref, onMounted, watch} from 'vue';
import { CreatePost } from '../../wailsjs/go/services/PostService';
import { computeSlug } from '../utils';
import { getPublicTopics } from '../utils/topic';
import { getAdminId } from '../utils/env';
import FAB from './FAB.vue';

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
});
const emit = defineEmits(['open', 'close', 'saved']);

const saving = ref(false);

const saveErrorSnackbar = ref(false);

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

/** @type {import('vue').Ref<string>} */
const coverImageRefName = ref('');

/** @type {import('vue').Ref<string>} */
const coverImageRefUrl = ref('');

const steppers = ref(['Metadata', 'Body']);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topics = ref([]);

function close() {
   emit('close');
}

async function savePost() {
   saving.value = true;

   try {
      console.log("saving...")
      const res = await CreatePost({
         title: title.value,
         slug: slug.value,
         desc: desc.value,
         topic: topic.value || 'other',
         tags: tags.value,
         body: '',
         public: true,
         coverImageId: '',
         coverImagePath: '',
         coverImageRefName: coverImageRefName.value,
         coverImageRefUrl: coverImageRefUrl.value,
         authorId: getAdminId(),
      });
      console.log("saved.", res);
   } catch (error) {
      console.error(error);
      saveErrorSnackbar.value = true;
   }

   saving.value = false;
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
            <v-tooltip text="Compose">
               <template v-slot:activator="{ props }">
                  <v-btn 
                     v-bind="props"
                     fixed 
                     size="large" 
                     color="secondary" 
                     icon="mdi-pencil"
                     elevation="8" 
                     @click="emit('open')">
                  </v-btn>
               </template>
            </v-tooltip>
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
               <v-btn :loading="saving" variant="text" @click="savePost">Save</v-btn>
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

                        <v-row class="mb-2">
                           <v-col cols="12" sm="4">
                              <v-file-input
                                 label="Select Cover Image"
                                 variant="filled"
                                 prepend-icon="mdi-image"
                                 accept="image/jpeg">
                              </v-file-input>
                           </v-col>
                           <v-col cols="12" sm="4">
                              <v-text-field v-model="coverImageRefName" label="Reference Name (optional)"></v-text-field>
                           </v-col>
                           <v-col cols="12" sm="4">
                              <v-text-field v-model="coverImageRefUrl" label="Reference URL (optional)"></v-text-field>
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
      <v-snackbar v-model="saveErrorSnackbar" :timeout="5000">Couldn't save post</v-snackbar>
   </v-dialog>
</template>

<style scoped>
.app-custom-toolbar {
   position: fixed;
   z-index: 120;
}
</style>
