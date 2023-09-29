<script setup>
import {defineProps, defineEmits, ref, computed} from 'vue';
import {CreatePost} from '../../../wailsjs/go/services/PostService';
import {getAdminId} from '../../utils/env';
import { state, clearMetadata } from './store';
import FAB from '../FAB.vue';
import Metadata from './Metadata.vue';

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
});
const emit = defineEmits(['open', 'close', 'saved']);

/** @type {import('vue').Ref<string[]>} */
const steppers = ref(['Metadata', 'Body']);

/** @type {import('vue').Ref<boolean>} */
const loadingSavePost = ref(false);

/** @type {import('vue').Ref<boolean>} */
const saveErrorSnackbar = ref(false);

const allowToSubmit = computed(function () {
   return state.title.length && state.slug.length;
});

function close() {
   emit('close');
}

async function savePost() {
   loadingSavePost.value = true;

   try {
      await CreatePost({
         title: state.title,
         slug: state.slug,
         desc: state.desc,
         topic: state.topic || 'other',
         tags: state.tags,
         body: 'Hello, world',
         public: state.publicScope,
         coverImageId: state.coverImageAssetId,
         coverImagePath: state.coverImagePublicId,
         coverImageRefName: state.coverImageRefName,
         coverImageRefUrl: state.coverImageRefUrl,
         authorId: getAdminId(),
      });
   } catch (error) {
      console.error(error);
      saveErrorSnackbar.value = true;
   }

   loadingSavePost.value = false;
   clearMetadata();
   emit('saved');
}
</script>

<template>
   <v-dialog
      v-model="props.visible"
      fullscreen
      :scrim="false"
      scrollable
      transition="dialog-bottom-transition">
      <template v-slot:activator="{props}">
         <FAB>
            <v-tooltip text="Compose">
               <template v-slot:activator="{props}">
                  <v-btn
                     v-bind="props"
                     fixed
                     size="large"
                     color="primary"
                     icon="mdi-pencil"
                     elevation="8"
                     @click="emit('open')">
                  </v-btn>
               </template>
            </v-tooltip>
         </FAB>
      </template>

      <v-card>
         <v-toolbar color="primary" class="app-custom-toolbar" dark>
            <v-btn icon dark @click="close">
               <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Compose</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="d-flex flex-row align-center mr-3">
               <v-switch
                  v-model="state.publicScope"
                  label="Public"
                  color="white"
                  style="grid-template-areas: unset">
               </v-switch>
            </v-toolbar-items>

            <v-toolbar-items>
               <v-btn
                  :loading="loadingSavePost"
                  :width="120"
                  :prepend-icon="state.publicScope ? 'mdi-upload' : 'mdi-file-check-outline'"
                  :disabled="!allowToSubmit"
                  @click="savePost">
                  {{ state.publicScope ? 'Publish' : 'Save' }}
               </v-btn>
            </v-toolbar-items>
         </v-toolbar>

         <v-form class="mt-16">
            <v-container>
               <v-stepper :items="steppers">
                  <template v-slot:item.1>
                     <Metadata />
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
