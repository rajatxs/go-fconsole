<script setup>
import {ref, computed, onMounted} from 'vue';
import {state} from './store.js';
import {getPublicTopics} from '../../utils/topic';
import {UploadPostCoverImage, DeletePostImage} from '../../../wailsjs/go/services/PostService';
import {getFileByteArray, getPostCoverImageURL, computeSlug} from '../../utils';

/** @type {import('vue').Ref<boolean>} */
const loadingUploadImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const loadingRemoveImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadSuccessSnackbar = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadErrorSnackbar = ref(false);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topics = ref([]);

const coverImageUploaded = computed(function () {
   return state.coverImagePublicId.length > 0 && state.coverImageAssetId.length > 0;
});

const coverImageUrl = computed(function () {
   return getPostCoverImageURL(state.coverImagePublicId);
});

onMounted(function () {
   renderTopics();
});

function renderTopics() {
   /** @type {any[]} */
   const _topics = getPublicTopics();

   topics.value = _topics.map((_topic) => {
      return {
         title: _topic.name,
         value: _topic.id,
      };
   });
}

/** @param {any} event  */
function onTitleKeyUp(event) {
   const newTitle = event.target.value;

   if (newTitle.length > 0) {
      state.slug = computeSlug(newTitle);
   } else {
      state.slug = '';
   }
}

async function onCoverImageUpload(event) {
   const [file] = event.target.files;

   /** @type {number[]} */
   let byteArray = [];

   /** @type {import('../../../wailsjs/go/models').types.PostImageFile} */
   let result;

   if (!file) {
      return;
   }

   loadingUploadImage.value = true;

   try {
      byteArray = await getFileByteArray(file);
      result = await UploadPostCoverImage(byteArray);

      state.coverImagePublicId = result.publicId;
      state.coverImageAssetId = result.assetId;
      coverImageUploadSuccessSnackbar.value = true;
   } catch (error) {
      console.error(error);
      coverImageUploadErrorSnackbar.value = true;
   }

   loadingUploadImage.value = false;
}

async function onCoverImageRemove() {
   loadingRemoveImage.value = true;

   try {
      await DeletePostImage(state.coverImagePublicId);

      state.coverImagePublicId = '';
      state.coverImageAssetId = '';
   } catch (error) {
      console.error(error);
   }

   loadingRemoveImage.value = false;
}
</script>

<template>
   <v-card flat>
      <!-- Title text field -->
      <v-text-field 
         v-model="state.title" 
         label="Title" 
         @keyup="onTitleKeyUp">
      </v-text-field>
      <v-row>
         <v-col cols="12" sm="6">
            <!-- Topic selection dropdown -->
            <v-select v-model="state.topic" label="Topic" :items="topics"></v-select>
         </v-col>
         <v-col cols="12" sm="6">
            <!-- Auto generated slug text field -->
            <v-text-field
               v-model="state.slug"
               :title="state.slug"
               label="Slug (Auto Generated)"
               readonly>
            </v-text-field>
         </v-col>
      </v-row>

      <!-- Description multiline text field -->
      <v-textarea v-model="state.desc" label="Description"></v-textarea>

      <!-- Tags input field -->
      <v-combobox v-model="state.tags" label="Tags" multiple chips> </v-combobox>

      <v-row class="mb-2">
         <v-col cols="12" sm="6">
            <!-- Uploaded cover image preview -->
            <v-card v-if="coverImageUploaded">
               <v-img 
                  :src="coverImageUrl" 
                  alt="Cover Image" 
                  aspect-ratio="16/9" 
                  width="auto" 
                  cover>
               </v-img>

               <v-expansion-panels>
                  <v-expansion-panel>
                     <v-expansion-panel-title>
                        <v-icon icon="mdi-information"></v-icon>&nbsp;<span class="font-weight-bold">
                           Cover Image Info
                        </span>
                     </v-expansion-panel-title>
                     <v-expansion-panel-text>
                        <p>
                           <span class="font-weight-bold">Asset Id:</span>
                           <br />
                           <code>{{ state.coverImageAssetId }}</code>
                        </p>
                        <p>
                           <span class="font-weight-bold">Public Id:</span>
                           <br />
                           <code>{{ state.coverImagePublicId }}</code>
                        </p>
                     </v-expansion-panel-text>
                  </v-expansion-panel>
               </v-expansion-panels>
            </v-card>

            <!-- Cover image input -->
            <template v-else>
               <v-file-input
                  v-show="!loadingUploadImage"
                  label="Upload Cover Image"
                  variant="filled"
                  prepend-icon="mdi-image"
                  accept="image/jpeg"
                  @change="onCoverImageUpload">
               </v-file-input>
               <v-row
                  v-show="loadingUploadImage"
                  class="fill-height"
                  align-content="center"
                  justify="center">
                  <v-col class="text-subtitle-1 text-center" cols="12">
                     Uploading Cover Image
                  </v-col>
                  <v-col cols="6">
                     <v-progress-linear color="primary-darken-2" height="6" indeterminate rounded>
                     </v-progress-linear>
                  </v-col>
               </v-row>
            </template>
         </v-col>

         <v-col cols="12" sm="6">
            <!-- Cover image reference name text field -->
            <v-text-field
               v-model="state.coverImageRefName"
               label="Reference Name (optional)"></v-text-field>

            <!-- Cover image reference url text field -->
            <v-text-field
               v-model="state.coverImageRefUrl"
               label="Reference URL (optional)"></v-text-field>

            <!-- Remove cover image action button -->
            <v-btn
               v-show="coverImageUploaded"
               :loading="loadingRemoveImage"
               variant="text"
               color="error"
               prepend-icon="mdi-image-remove-outline"
               @click="onCoverImageRemove">
               Remove Cover Image
            </v-btn>
         </v-col>
      </v-row>
      <v-snackbar v-model="coverImageUploadSuccessSnackbar" :timeout="3000" color="primary">
         Cover Image Uploaded
      </v-snackbar>
      <v-snackbar v-model="coverImageUploadErrorSnackbar" :timeout="3000" color="error">
         Couldn't upload Cover Image
      </v-snackbar>
   </v-card>
</template>
