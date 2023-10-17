<script setup>
import {defineProps, defineEmits, ref, computed, watch} from 'vue';
import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import Paragraph from '@editorjs/paragraph';
import List from '@editorjs/list';
import Table from '@editorjs/table';
import CodeTool from '@editorjs/code';
import WarningTool from '@editorjs/warning';
import InlineCodeTool from '@editorjs/inline-code';
import {
   CreatePost,
   UpdatePostById,
   UploadPostEmbedImage,
   GetPostById,
} from '../../../wailsjs/go/services/PostService';
import {getFileByteArray, getPostEmbeddedImageUrl} from '../../utils';
import {getAdminId} from '../../utils/env';
import {state, setMetadata, clearMetadata} from './store';
import CustomImageTool from '../../plugins/image-tool';
import FAB from '../FAB.vue';
import Metadata from './Metadata.vue';
import BodyView from './Body.vue';

/** @type {import('@editorjs/editorjs').default} */
let editor;

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
   id: {
      type: String,
      default: '',
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

/** @type {import('vue').Ref<boolean>} */
const errorLoadData = ref(false);

/** @type {import('vue').Ref<'create'|'update'>} */
const action = computed(function () {
   return props.id.length ? 'update' : 'create';
});

const allowToSubmit = computed(function () {
   return state.title.length && state.slug.length;
});

const toolbarColor = computed(function() {
   return state.publicScope? 'primary': 'default';
});

function close() {
   if (action.value === 'update') {
      if (errorLoadData.value) {
         errorLoadData.value = false;
      }

      clearMetadata();
   }

   emit('close');
}

/** @param {File} file */
async function uploadByFile(file) {
   /** @type {import('../../../wailsjs/go/models').types.UploadedImageFile} */
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
            id: result.assetId,
            path: result.publicId,
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
      data: state.body,
      minHeight: 400,
      readOnly: false,
      holder: 'codex-editor',
      autofocus: true,
      defaultBlock: 'paragraph',
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
            inlineToolbar: true,
         },
         paragraph: {
            class: Paragraph,
         },
         image: {
            // @ts-ignore
            class: CustomImageTool,
            config: {
               types: 'image/jpeg',
               uploader: {uploadByFile},
            },
         },
         table: {
            class: Table,
            config: {
               rows: 3,
               cols: 3,
            },
         },
         code: {
            class: CodeTool,
         },
         warning: {
            class: WarningTool,
         },
         inlineCode: {
            class: InlineCodeTool,
         }
      },
   });
}

/** @param {import('@editorjs/editorjs').OutputData} body  */
async function createPost(body) {
   const relatedPosts = state.relatedPosts.map(p => p.value);

   try {
      await CreatePost({
         title: state.title,
         slug: state.slug,
         desc: state.desc,
         topic: state.topic || 'other',
         tags: state.tags,
         body,
         public: state.publicScope,
         format: 'block', // app supports block-style editor only
         coverImageId: state.coverImageAssetId,
         coverImagePath: state.coverImagePublicId,
         coverImageRefName: state.coverImageRefName,
         coverImageRefUrl: state.coverImageRefUrl,
         authorId: getAdminId(),
         relatedPosts,
      });
   } catch (error) {
      console.error(error);
      saveErrorSnackbar.value = true;
   }
}

/** @param {import('@editorjs/editorjs').OutputData} body  */
async function updatePost(body) {
   const relatedPosts = state.relatedPosts.map(p => p.value);

   try {
      await UpdatePostById(props.id, {
         title: state.title,
         slug: state.slug,
         desc: state.desc,
         topic: state.topic || 'other',
         tags: state.tags,
         body,
         public: state.publicScope,
         coverImageId: state.coverImageAssetId,
         coverImagePath: state.coverImagePublicId,
         coverImageRefName: state.coverImageRefName,
         coverImageRefUrl: state.coverImageRefUrl,
         relatedPosts,
      });
   } catch (error) {
      console.error(error);
      saveErrorSnackbar.value = true;
   }
}

async function savePost() {
   /** @type {import('@editorjs/editorjs').OutputData} */
   let body;

   loadingSavePost.value = true;

   try {
      if (editor) {
         body = await editor.save();
      } else {
         body = state.body;
      }
   } catch (error) {
      console.error(error);
      docComileErrorSnackbar.value = true;
      loadingSavePost.value = false;
      return;
   }

   if (action.value === 'create') {
      await createPost(body);
   } else if (action.value === 'update') {
      await updatePost(body);
   }

   loadingSavePost.value = false;
   clearMetadata();
   emit('saved');
}

watch(
   () => props.visible,
   async function (newState) {
      if (!newState) {
         return;
      }

      if (action.value === 'update') {
         try {
            const post = await GetPostById(props.id);
            console.log("post::", post);
            setMetadata(post);
         } catch (error) {
            console.error(error);
            errorLoadData.value = true;
         }
      }
   }
);
</script>

<template>
   <v-dialog
      v-model="props.visible"
      :scrim="false"
      transition="dialog-bottom-transition"
      fullscreen
      scrollable>
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
         <v-toolbar :color="toolbarColor" class="app-custom-toolbar" dark>
            <v-btn icon dark @click="close">
               <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title v-if="action === 'create'">Compose</v-toolbar-title>
            <v-toolbar-title v-else>Edit</v-toolbar-title>
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
      <v-snackbar v-model="saveErrorSnackbar" :timeout="5000" color="error">Couldn't save post</v-snackbar>
      <v-snackbar v-model="errorLoadData" :timeout="3000" color="error">Couldn't get post data</v-snackbar>
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
