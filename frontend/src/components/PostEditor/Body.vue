<script setup>
import {ref, onMounted} from 'vue';
import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import Paragraph from '@editorjs/paragraph';
import List from '@editorjs/list';
import {UploadPostEmbedImage} from '../../../wailsjs/go/services/PostService';
import {getFileByteArray, getPostEmbeddedImageUrl} from '../../utils';
import CustomImageTool from '../../plugins/image-tool';

/** @type {import('@editorjs/editorjs').default} */
let editor;

const imageUploadErrorSnackbar = ref(false);

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

   window['editor'] = editor;
}

onMounted(initEditor);
</script>

<template>
   <v-card flat>
      <div id="codex-editor"></div>

      <v-snackbar v-model="imageUploadErrorSnackbar" :timeout="3000" color="error">
         Couldn't upload Cover Image
      </v-snackbar>
   </v-card>
</template>

<style></style>
