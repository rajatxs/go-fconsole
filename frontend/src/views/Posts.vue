<script setup>
import {ref, onMounted, watch} from 'vue';
import {GetPostsMetadata, GetPostCount} from '../../wailsjs/go/main/App';
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
      const _posts = await GetPostsMetadata(
         scope.value,
         topic.value,
         sort.value,
         limit.value,
         (pageIndex.value - 1) * limit.value
      );

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

                  <v-card-actions>
                     <v-btn color="primary">Edit</v-btn>
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
   </v-container>
</template>
