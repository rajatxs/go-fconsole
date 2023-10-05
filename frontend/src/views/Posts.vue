<script setup>
import {ref, onMounted, watch} from 'vue';
import {
   GetPostsMetadata, 
   GetPostCount, 
   UpdatePostScope, 
   SetPostDeleteFlag
} from '../../wailsjs/go/services/PostService';
import { ClipboardSetText } from '../../wailsjs/go/main/App';
import {groupArray, truncateText, getPostCoverImageURL} from '../utils';
import {getTopicName, getPublicTopics} from '../utils/topic';
import PostEditor from '../components/PostEditor/index.vue';
import Loader from '../components/Loader.vue';

/** @type {import('vue').Ref<any[][]>} */
const postsGroups = ref([]);

const deletePostConfirmDialog = ref(false);
const composeMode = ref(false);

const loading = ref(true);
const topic = ref('all');
const scope = ref('public');
const sort = ref('newest');
const limit = ref(9);
const deletePostId = ref('');
const editPostId = ref('');
const pageIndex = ref(1);
const maxPageIndex = ref(0);
const fetchErrorSnackbar = ref(false);
const copySuccessSnackbar = ref(false);
const copyErrorSnackbar = ref(false);
const errorScopeUpdateSnackbar = ref(false);
const deleteSuccessSnackbar = ref(false);
const deleteErrorSnackbar = ref(false);

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

         postsGroups.value = groupArray(_posts, 3);
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
      await ClipboardSetText(id);
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

/** Deletes post by given post id */
async function deletePost() {
   try {
      await SetPostDeleteFlag(deletePostId.value, true);
      deleteSuccessSnackbar.value = true;
      await fetchPosts();
   } catch (error) {
      deleteErrorSnackbar.value = true;
   }

   deletePostId.value = '';
}

async function onNewPost() {
   closePostEditor();
   await fetchPosts();
}

/**
 * Opens PostEditor to perform update operation
 * @param {string} id 
 */
function editPost(id) {
   editPostId.value = id;
   composeMode.value = true;
}

function closePostEditor() {
   if (composeMode.value) {
      composeMode.value = false;
   }

   if (editPostId.value.length > 0) {
      editPostId.value = '';
   }
}

watch([topic, sort, pageIndex], fetchPosts);
watch(scope, async function () {
   // reset page index to avoid conflict
   pageIndex.value = 1;
   await fetchPosts();
});
watch(deletePostId, function(newPostId) {
   deletePostConfirmDialog.value = (newPostId.length > 0)
});

onMounted(async function () {
   renderTopicFilter();
   await fetchPosts();
});
</script>

<template>
   <v-container>
      <v-row>
         <v-col cols="12" sm="3">
            <div class="text-h4 m-0">Posts</div>
         </v-col>

         <!-- Scope selection button group -->
         <v-col cols="12" sm="3">
            <v-btn-toggle
               v-model="scope"
               color="primary-darken-1"
               rounded="lg"
               mandatory
               group>
               <v-btn value="public">
                  <v-icon icon="mdi-earth"></v-icon>&nbsp;
                  <span>Public</span>
               </v-btn>
               <v-btn value="private">
                  <v-icon icon="mdi-lock"></v-icon>&nbsp;
                  <span>Private</span>
               </v-btn>
            </v-btn-toggle>
         </v-col>

         <!-- Topic selection dropdown -->
         <v-col cols="12" sm="3">
            <v-select v-model="topic" :items="topicOptions" label="Topic"></v-select>
         </v-col>

         <!-- Sort filter -->
         <v-col cols="12" sm="3">
            <v-select v-model="sort" :items="sortOptions" label="Sort By"></v-select>
         </v-col>
      </v-row>

      <v-container class="px-0">
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
               <v-card class="mx-auto">
                  <v-img
                     :src="getPostCoverImageURL(post.coverImage.path)"
                     height="200"
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

                  <v-card-actions class="flex flex-row justify-space-between pl-3 pr-3">
                     <!-- Post primary actions -->
                     <div>
                        <v-btn 
                           v-show="scope === 'public'" 
                           color="primary-darken-4"
                           @click="openPreview(post.slug)">
                           Preview
                        </v-btn>
                        <v-btn 
                           color="primary-darken-4" 
                           @click="editPost(post._id)">
                           Edit
                        </v-btn>
                     </div>

                     <!-- Post context menu -->
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
                              value="preview"
                              class="text-error"
                              @click="deletePostId = post._id">
                           </v-list-item>
                        </v-list>
                     </v-menu>
                  </v-card-actions>
               </v-card>
            </v-col>
         </v-row>
      </v-container>

      <PostEditor
         :visible="composeMode" 
         :id="editPostId"
         @saved="onNewPost"
         @open="composeMode = true"
         @close="closePostEditor" />

      <!-- Delete Post Dialog -->
      <v-dialog v-model="deletePostConfirmDialog" :width="460" persistent>
         <v-card>
            <v-card-text>Are you sure you want to delete this post?</v-card-text>
            <v-card-actions>
               <v-spacer></v-spacer>
               <v-btn variant="text" @click="deletePostId = ''">
                  No
               </v-btn>

               <v-btn
                  color="error"
                  variant="tonal"
                  @click="deletePost()">
                  Yes
               </v-btn>
            </v-card-actions>
         </v-card>
      </v-dialog>
 
      <v-pagination
         v-model="pageIndex"
         :length="maxPageIndex"
         class="mt-5">
      </v-pagination>

      <v-snackbar
         v-model="fetchErrorSnackbar"
         :timeout="5000"
         color="error">
         Couldn't get posts
      </v-snackbar>

      <v-snackbar
         v-model="copySuccessSnackbar"
         :timeout="3000"
         color="primary">
         Copied to clipboard
      </v-snackbar>

      <v-snackbar
         v-model="copyErrorSnackbar"
         :timeout="3000"
         color="error">
         Couldn't write to clipboard
      </v-snackbar>

      <v-snackbar
         v-model="errorScopeUpdateSnackbar"
         :timeout="3000"
         color="error">
         Couldn't update scope
      </v-snackbar>

      <v-snackbar
         v-model="deleteSuccessSnackbar"
         :timeout="3000"
         color="primary">
         Post deleted successfully
      </v-snackbar>

      <v-snackbar
         v-model="deleteErrorSnackbar"
         :timeout="3000"
         color="error">
         Couldn't delete post
      </v-snackbar>
   </v-container>
</template>
