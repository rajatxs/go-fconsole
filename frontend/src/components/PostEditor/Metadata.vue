<script setup>
import {ref, computed, onMounted} from 'vue';
import {state} from './store.js';
import PostSelector from '../PostSelector/index.vue';
import {UploadPostCoverImage, DeletePostImage} from '../../../wailsjs/go/services/PostService';
import {GetPublicTopics} from '../../../wailsjs/go/services/TopicService';
import {getFileByteArray, getPostCoverImageURL, computeSlug} from '../../utils';

/** @type {import('vue').Ref<boolean>} */
const loadingUploadImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const loadingRemoveImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const viewPostSelector = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadSuccessSnackbar = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadErrorSnackbar = ref(false);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topics = ref([]);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const licenses = ref([
   { title: "Creative Commons Attribution (CC BY) 4.0", value: "CC-BY-4.0" },
   { title: "Creative Commons Attribution-NonCommercial (CC BY-NC) 4.0", value: "CC-BY-NC-4.0" },
   { title: "Creative Commons Attribution-ShareAlike (CC BY-SA) 4.0", value: "CC-BY-SA-4.0" },
   { title: "Creative Commons Attribution-NoDerivs (CC BY-ND) 4.0", value: "CC-BY-ND-4.0" },
]);

const coverImageUploaded = computed(function () {
   return state.coverImagePublicId.length > 0 && state.coverImageAssetId.length > 0;
});

const coverImageUrl = computed(function () {
   return getPostCoverImageURL(state.coverImagePublicId);
});

const relatedPostsIds = computed(function() {
   return state.relatedPosts.map(p => p.value);
});

onMounted(async function () {
   await renderTopics();
});

async function renderTopics() {
   const _topics = await GetPublicTopics();
   const options = [];

   for (const _topicId in _topics) {
      options.push({
         title: _topics[_topicId].name,
         value: _topicId,
      });
   }

   topics.value = options;
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

   /** @type {import('../../../wailsjs/go/models').types.UploadedImageFile} */
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

/**
 * Inserts new post tag in related post input field
 * @param {import('../../../wailsjs/go/models').models.PostMetadataDocument} post 
 */
function addRelatedPost(post) {
   state.relatedPosts.push({
      title: post.title,
      value: post._id.toString(),
   });
   viewPostSelector.value = false;
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

      <v-row>
         <v-col cols="12" sm="6">
            <!-- Tags input field -->
            <v-combobox v-model="state.tags" label="Tags" multiple chips></v-combobox>
         </v-col>
         <v-col cols="12" sm="6">
            <!-- License selection dropdown -->
            <v-select v-model="state.license" label="License" :items="licenses"></v-select>
         </v-col>
      </v-row>

      <!-- Related post input field -->
      <v-row>
         <v-col cols="12">
            <v-combobox
               v-model="state.relatedPosts"
               label="Related Posts"
               multiple
               chips
               closable-chips>
               <template v-slot:append>
                  <v-btn 
                     prepend-icon="mdi-select-search" 
                     @click="viewPostSelector = true">
                     Add Post
                  </v-btn>
               </template>
            </v-combobox>
         </v-col>
      </v-row>

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

   <PostSelector 
      :visible="viewPostSelector" 
      :topic="state.topic || undefined"
      :selected="relatedPostsIds"
      @select="addRelatedPost"
      @close="viewPostSelector = false" />
</template>
