<script setup>
import {ref, onMounted, watch} from 'vue';
import {GetPostsMetadata, GetPostCount, UpdatePostScope} from '../../wailsjs/go/services/PostService';
import {groupArray, truncateText, getPostCoverImageURL} from '../utils';
import {getTopicName, getPublicTopics} from '../utils/topic';
import ComposePostDialog from '../components/ComposePostDialog.vue';
import Loader from '../components/Loader.vue';

/** @type {import('vue').Ref<any[][]>} */
const postsGroups = ref([]);

const composeMode = ref(false);

const loading = ref(true);
const topic = ref('all');
const scope = ref('public');
const sort = ref('newest');
const limit = ref(6);
const pageIndex = ref(1);
const maxPageIndex = ref(0);
const fetchErrorSnackbar = ref(false);
const copySuccessSnackbar = ref(false);
const copyErrorSnackbar = ref(false);
const errorScopeUpdateSnackbar = ref(false);

/** @type {import('vue').Ref<any[]>} */
const topicOptions = ref([]);

const sortOptions = ref([
   {
      title: 'Title',
      value: 'title',
   },
   {
      title: 'Topic',
      value: 'topic',
   },
   {
      title: 'Newest',
      value: 'newest',
   },
   {
      title: 'Oldest',
      value: 'oldest',
   },
   {
      title: 'Last Updated',
      value: 'updated',
   },
]);

function renderTopicFilter() {
   /** @type {any} */
   const _topics = getPublicTopics();

   const options = [
      {
         title: 'All',
         value: 'all',
      },
   ];

   for (const _topic of _topics) {
      options.push({
         title: _topic.name,
         value: _topic.id,
      });
   }

   topicOptions.value = options;
}

async function fetchPosts() {
   loading.value = true;

   try {
      const _posts = await GetPostsMetadata({
         private: scope.value === 'private',
         topic: topic.value,
         sortBy: sort.value,
         limit: limit.value,
         skip: (pageIndex.value - 1) * limit.value
      });

      if (Array.isArray(_posts)) {
         const _postCount = await GetPostCount(scope.value, false);
         const _group = groupArray(_posts, 3);

         postsGroups.value = _group;
         maxPageIndex.value = Math.ceil(_postCount / limit.value);
      } else {
         postsGroups.value = [];
      }
   } catch (error) {
      fetchErrorSnackbar.value = true;
   }
   loading.value = false;
}

/**
 * Opens preview of selected post in a new tab or window
 * @param {string} slug - Post slug
 */
function openPreview(slug) {
   window.open(`https://www.fivemin.in/${slug}`, '_blank');
}

/**
 * Writes given post id to clipboard
 * @param {string} id - Post Id
 */
async function copyId(id) {
   try {
      await navigator.clipboard.writeText(id);
      copySuccessSnackbar.value = true;
   } catch (error) {
      copyErrorSnackbar.value = true;
   }
}

/**
 * Makes given post private
 * @param {string} id - Post Id
 */
async function makePrivate(id) {
   try {
      await UpdatePostScope(id, "private");
      await fetchPosts();
   } catch (error) {
      errorScopeUpdateSnackbar.value = true;
   }
}

/**
 * Makes given post public
 * @param {string} id - Post Id
 */
async function makePublic(id) {
   try {
      await UpdatePostScope(id, "public");
      await fetchPosts();
   } catch (error) {
      errorScopeUpdateSnackbar.value = true;
   }
}

watch([topic, sort, pageIndex], fetchPosts);
watch(scope, async function () {
   // reset page index to avoid conflict
   pageIndex.value = 1;
   await fetchPosts();
});

onMounted(async function () {
   renderTopicFilter();
   await fetchPosts();
});
</script>

<template>
   <v-container>
      <div class="text-h4 mb-5">Posts</div>

      <v-row>
         <v-col cols="12" sm="4">
            <v-tabs v-model="scope" color="primary">
               <v-tab value="public"><v-icon>mdi-earth</v-icon>&nbsp;Public</v-tab>
               <v-tab value="private"><v-icon>mdi-lock</v-icon>&nbsp;Private</v-tab>
            </v-tabs>
         </v-col>
         <v-col cols="12" sm="4">
            <v-select label="Topic" :items="topicOptions" v-model="topic"></v-select>
         </v-col>
         <v-col cols="12" sm="4">
            <v-select label="Sort By" :items="sortOptions" v-model="sort"></v-select>
         </v-col>
      </v-row>

      <v-container>
         <Loader v-if="loading" message="Getting posts" />
         <v-sheet
            v-else-if="!loading && postsGroups.length === 0"
            :width="600"
            :height="500"
            class="d-flex flex-column align-center justify-center flex-wrap text-center mx-auto px-4">
            <div class="text-h6">No Posts</div>
         </v-sheet>
         <v-row v-else v-for="(group, groupIndex) of postsGroups" :key="groupIndex">
            <v-col v-for="post of group" :key="post._id" cols="12" sm="4">
               <v-card class="mx-auto" :elevation="2">
                  <v-img
                     class="align-end text-white"
                     height="200"
                     :src="getPostCoverImageURL(post.coverImage.path)"
                     cover>
                  </v-img>

                  <v-card-item>
                     <v-card-title>{{ post.title }}</v-card-title>

                     <v-card-subtitle>
                        <span class="me-1">{{ getTopicName(post.topic) }}</span>
                     </v-card-subtitle>
                  </v-card-item>

                  <v-card-text>
                     {{ truncateText(post.desc, 72) }}
                  </v-card-text>

                  <v-card-actions class="flex flex-row justify-space-between pl-2 pr-2">
                     <div>
                        <v-btn 
                           v-show="scope === 'public'" 
                           color="primary" 
                           @click="openPreview(post.slug)">
                           Preview
                        </v-btn>
                        <v-btn color="primary">Edit</v-btn>
                     </div>

                     <v-menu :width="180" location="start">
                        <template v-slot:activator="{ props }">
                           <v-btn v-bind="props" icon="mdi-dots-vertical"></v-btn>
                        </template>

                        <v-list>
                           <v-list-item 
                              title="Copy ID" 
                              value="copy-id"
                              @click="copyId(post._id)">
                           </v-list-item>
                           
                           <v-list-item 
                              v-if="scope === 'public'" 
                              title="Make Private" 
                              value="make-private"
                              @click="makePrivate(post._id)">
                           </v-list-item>

                           <v-list-item 
                              v-else 
                              title="Make Public" 
                              value="make-private"
                              @click="makePublic(post._id)">
                           </v-list-item>

                           <v-list-item 
                              title="Delete" 
                              value="preview">
                           </v-list-item>
                        </v-list>
                     </v-menu>
                  </v-card-actions>
               </v-card>
            </v-col>
         </v-row>
      </v-container>

      <ComposePostDialog 
         :visible="composeMode" 
         @open="composeMode = true"
         @close="composeMode = false" />
 
      <v-pagination v-model="pageIndex" :length="maxPageIndex" class="mt-5"></v-pagination>
      <v-snackbar v-model="fetchErrorSnackbar" :timeout="5000"> Couldn't get posts </v-snackbar>
      <v-snackbar v-model="copySuccessSnackbar" :timeout="3000"> Copied to clipboard </v-snackbar>
      <v-snackbar v-model="copyErrorSnackbar" :timeout="3000"> Couldn't write to clipboard </v-snackbar>
      <v-snackbar v-model="errorScopeUpdateSnackbar" :timeout="3000"> Couldn't update scope </v-snackbar>
   </v-container>
</template>
