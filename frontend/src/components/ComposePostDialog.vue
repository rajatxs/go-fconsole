<script setup>
import {defineProps, defineEmits, ref, onMounted, watch, computed} from 'vue';
import { CreatePost, UploadPostCoverImage, DeletePostCoverImage } from '../../wailsjs/go/services/PostService';
import { computeSlug, getFileByteArray, getPostCoverImageURL } from '../utils';
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

/** @type {import('vue').Ref<boolean>} */
const publicScope = ref(true);

/** @type {import('vue').Ref<string>} */
const coverImageRefName = ref('');

/** @type {import('vue').Ref<string>} */
const coverImageRefUrl = ref('');

/** @type {import('vue').Ref<string>} */
const coverImagePublicId = ref('');

/** @type {import('vue').Ref<string>} */
const coverImageAssetId = ref('');

const steppers = ref(['Metadata', 'Body']);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topics = ref([]);

/** @type {import('vue').Ref<boolean>} */
const loadingUploadImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const loadingRemoveImage = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadSuccessSnackbar = ref(false);

/** @type {import('vue').Ref<boolean>} */
const coverImageUploadErrorSnackbar = ref(false);

const coverImageUploaded = computed(function() {
   return (
      coverImagePublicId.value.length > 0 && 
      coverImageAssetId.value.length > 0
   );
});

const coverImageUrl = computed(function() {
   return getPostCoverImageURL(coverImagePublicId.value);
});

const allowToSubmit = computed(function() {
   return (title.value.length && slug.value.length)
});

function close() {
   emit('close');
}

async function uploadCoverImage(event) {
   const [ file ] = event.target.files;

   /** @type {number[]} */
   let byteArray = new Array();

   /** @type {import('../../wailsjs/go/models').types.PostImageFile} */
   let result;

   if (!file) {
      return;
   }

   loadingUploadImage.value = true;

   try {
      byteArray = await getFileByteArray(file);
      result = await UploadPostCoverImage(byteArray);

      coverImagePublicId.value = result.publicId;
      coverImageAssetId.value = result.assetId;
      coverImageUploadSuccessSnackbar.value = true;
   } catch (error) {
      console.error(error);
      coverImageUploadErrorSnackbar.value = true;
   }

   loadingUploadImage.value = false;
}

async function removeCoverImage() {
   loadingRemoveImage.value = true;

   try {
      await DeletePostCoverImage(coverImagePublicId.value);

      coverImagePublicId.value = '';
      coverImageAssetId.value = '';
   } catch (error) {
      console.error(error);
   }

   loadingRemoveImage.value = false;
}

async function savePost() {
   saving.value = true;

   try {
      await CreatePost({
         title: title.value,
         slug: slug.value,
         desc: desc.value,
         topic: topic.value || 'other',
         tags: tags.value,
         body: 'Hello, world',
         public: publicScope.value,
         coverImageId: coverImageAssetId.value,
         coverImagePath: coverImagePublicId.value,
         coverImageRefName: coverImageRefName.value,
         coverImageRefUrl: coverImageRefUrl.value,
         authorId: getAdminId(),
      });
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
         <v-toolbar dark color="primary" class="app-custom-toolbar">
            <v-btn icon dark @click="close">
               <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Compose</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items class="d-flex flex-row align-center mr-3">
               <v-switch 
                  v-model="publicScope"
                  label="Public" 
                  color="white" 
                  style="grid-template-areas: unset;">
               </v-switch>
            </v-toolbar-items>

            <v-toolbar-items>
               <v-btn 
                  :loading="saving" 
                  :width="120" 
                  :prepend-icon="publicScope? 'mdi-upload': 'mdi-file-check-outline'"
                  :disabled="!allowToSubmit"
                  @click="savePost">
                  {{ publicScope? 'Publish': 'Save' }}
               </v-btn>
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
                           <v-col cols="12" sm="6">
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
                                          <v-icon icon="mdi-information"></v-icon>&nbsp;<span class="font-weight-bold">Cover Image Info</span>
                                       </v-expansion-panel-title>
                                       <v-expansion-panel-text>
                                          <p>
                                             <span class="font-weight-bold">Asset Id:</span>
                                             <br />
                                             <code>{{ coverImageAssetId }}</code>
                                          </p>
                                          <p>
                                             <span class="font-weight-bold">Public Id:</span>
                                             <br />
                                             <code>{{ coverImagePublicId }}</code>
                                          </p>
                                       </v-expansion-panel-text>
                                    </v-expansion-panel>
                                 </v-expansion-panels>
                              </v-card>

                              <template v-else>
                                 <v-file-input
                                    v-show="!loadingUploadImage"
                                    label="Upload Cover Image"
                                    variant="filled"
                                    prepend-icon="mdi-image"
                                    accept="image/jpeg"
                                    @change="uploadCoverImage">
                                 </v-file-input>
                                 <v-row
                                    v-show="loadingUploadImage"
                                    class="fill-height"
                                    align-content="center"
                                    justify="center">
                                    <v-col
                                       class="text-subtitle-1 text-center"
                                       cols="12">
                                       Uploading Cover Image
                                    </v-col>
                                    <v-col cols="6">
                                       <v-progress-linear
                                          color="primary-darken-2"
                                          height="6"
                                          indeterminate
                                          rounded>
                                       </v-progress-linear>
                                    </v-col>
                                 </v-row>
                              </template>
                           </v-col>

                           <v-col cols="12" sm="6">
                              <v-text-field v-model="coverImageRefName" label="Reference Name (optional)"></v-text-field>
                              <v-text-field v-model="coverImageRefUrl" label="Reference URL (optional)"></v-text-field>
                              <v-btn 
                                 v-show="coverImageUploaded"
                                 :loading="loadingRemoveImage"
                                 variant="text" 
                                 color="error" 
                                 prepend-icon="mdi-image-remove-outline"
                                 @click="removeCoverImage">
                                 Remove Cover Image
                              </v-btn>
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
      <v-snackbar v-model="coverImageUploadSuccessSnackbar" :timeout="3000" color="primary">Cover Image Uploaded</v-snackbar>
      <v-snackbar v-model="coverImageUploadErrorSnackbar" :timeout="3000" color="error">Couldn't upload Cover Image</v-snackbar>
   </v-dialog>
</template>

<style scoped>
.app-custom-toolbar {
   position: fixed;
   z-index: 120;
}
</style>
