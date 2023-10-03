<script setup>
import {defineProps, defineEmits, ref, computed, onMounted} from 'vue';
import {CreatePost} from '../../../wailsjs/go/services/PostService';
import {getAdminId} from '../../utils/env';
import { state, clearMetadata } from './store';
import FAB from '../FAB.vue';
import Metadata from './Metadata.vue';
import BodyView from './Body.vue';

import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import Paragraph from '@editorjs/paragraph';
import List from '@editorjs/list';
import {UploadPostEmbedImage} from '../../../wailsjs/go/services/PostService';
import {getFileByteArray, getPostEmbeddedImageUrl} from '../../utils';
import CustomImageTool from '../../plugins/image-tool';

/** @type {import('@editorjs/editorjs').default} */
let editor;

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
   
/** @type {import('vue').Ref<boolean>} */
const imageUploadErrorSnackbar = ref(false);
   
/** @type {import('vue').Ref<boolean>} */
const docComileErrorSnackbar = ref(false);

const allowToSubmit = computed(function () {
   return state.title.length && state.slug.length;
});

function close() {
   emit('close');
}

/** @param {File} file */
async function uploadByFile(file) {
   /** @type {import('../../../wailsjs/go/models').types.PostImageFile} */
   let result;

   if (!file) {
      return;
   }

   try {
      const byteArray = await getFileByteArray(file);
      result = await UploadPostEmbedImage(byteArray);

      return {
         success: 1,
         file: {
            assetId: result.assetId,
            publicId: result.publicId,
            url: getPostEmbeddedImageUrl(result.publicId),
         },
      };
   } catch (error) {
      imageUploadErrorSnackbar.value = true;
      return {
         success: 0,
      };
   }
}

function initEditor() {
   editor = new EditorJS({
      minHeight: 400,
      readOnly: false,
      holder: 'codex-editor',
      autofocus: true,
      initialBlock: 'paragraph',
      placeholder: 'Start writing here...',
      tools: {
         header: {
            class: Header,
            config: {
               placeholder: 'Enter a header',
               defaultLevel: 2,
               levels: [2, 3, 4, 5],
            },
         },
         list: {
            class: List,
         },
         paragraph: {
            class: Paragraph,
         },
         image: {
            // @ts-ignore
            class: CustomImageTool,
            config: {
               types: 'image/jpeg',
               uploader: { uploadByFile },
            },
         },
      },
   });
}

async function savePost() {
   /** @type{string} */
   let body;

   loadingSavePost.value = true;

   try {
      const outputData = await editor.save();
      body = JSON.stringify(outputData);
   } catch (error) {
      console.error(error);
      docComileErrorSnackbar.value = true;
      loadingSavePost.value = false;
      return;
   }

   try {
      await CreatePost({
         title: state.title,
         slug: state.slug,
         desc: state.desc,
         topic: state.topic || 'other',
         tags: state.tags,
         body,
         public: state.publicScope,

         // app supports block-style editor only
         format: "block",
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
                     <BodyView @mount="initEditor" />
                  </template>
               </v-stepper>
            </v-container>
         </v-form>
      </v-card>
      <v-snackbar v-model="saveErrorSnackbar" :timeout="5000">Couldn't save post</v-snackbar>
      <v-snackbar v-model="imageUploadErrorSnackbar" :timeout="3000" color="error">
         Couldn't upload Image
      </v-snackbar>
   </v-dialog>
</template>

<style scoped>
.app-custom-toolbar {
   position: fixed;
   z-index: 120;
}
</style>
