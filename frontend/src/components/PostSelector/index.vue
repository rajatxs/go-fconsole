<script setup>
import { defineProps, defineEmits, ref, watch, onBeforeUnmount, computed } from 'vue';
import { GetPostsMetadata, GetPublicPostCountByTopic } from '../../../wailsjs/go/services/PostService';
import { GetPublicTopics } from '../../../wailsjs/go/services/TopicService';
import { formatTime, getPostCoverImageURL } from '../../utils';

/** @type {Record<string, any>} */
let publicTopics = {};

const props = defineProps({
   visible: {
      type: Boolean,
      default: false,
   },
   topic: {
      type: String,
      required: false,
   },
   selected: {
      type: Array,
      default: [],
   },
});
const emit = defineEmits(['close', 'select']);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
const topicOptions = ref([]);

/** @type {import('vue').Ref<{title: string, value: string}[]>} */
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

/** @type {import('vue').Ref<boolean>} */
const fetchPostError = ref(false);

/** @type {import('vue').Ref<boolean>} */
const fetchTopicError = ref(false);

/** @type {import('vue').Ref<string>} */
const topic = ref('all');

/** @type {import('vue').Ref<string>} */
const sort = ref('updated');

/** @type {import('vue').Ref<any[]>} */
const posts = ref([]);

/** @type {import('vue').Ref<number>} */
const pageIndex = ref(1);
   
/** @type {import('vue').Ref<number>} */
const maxPageIndex = ref(0);

const postsList = computed(function() {
   let _list = [];

   posts.value.forEach((_md, _mdindex) => {
      if (_mdindex > 0 && _mdindex < posts.value.length) {
         _list.push({ type: 'divider', inset: true });
      }
      _list.push(_md);
   });

   return _list;
});

/** @param {string} id  */
function topicName(id) {
   if (publicTopics[id]) {
      return publicTopics[id].name;
   } else {
      return 'Other';
   }
}

async function fetchPostsMetadata() {
   const limit = 8;

   try {
      let _posts = await GetPostsMetadata({
         private: false,
         topic: topic.value,
         sortBy: sort.value,
         skip: (pageIndex.value - 1) * limit,
         limit,
      });

      if (Array.isArray(_posts)) {
         try {
            const _postCount = await GetPublicPostCountByTopic(topic.value);
            maxPageIndex.value = Math.ceil(_postCount / limit);
         } catch (error) {
            console.error(error);
            fetchPostError.value = true;
         }
      } else {
         maxPageIndex.value = 0;
         _posts = [];
      }

      // skip selected posts
      posts.value = _posts.filter(p => !props.selected.includes(p._id));
   } catch (error) {
      console.error(error);
      fetchPostError.value = true;
   }
}

/** @param {any} $event */
function onPostSelect($event) {
   const id = $event.id;
   const post = posts.value.find(_post => _post._id === id);

   if (!post) {
      return;
   }

   emit('select', post);
}

watch([sort, pageIndex], fetchPostsMetadata);

watch(topic, async function() {
   pageIndex.value = 1;
   await fetchPostsMetadata();
});

watch(() => props.visible, async function(newState) {
   if (!newState) {
      return;
   }

   if (props.topic) {
      topic.value = props.topic;
   }

   try {
      let options = [{
         title: "All",
         value: "all",
      }];

      publicTopics = await GetPublicTopics();

      for (const topicId in publicTopics) {
         options.push({
            title: publicTopics[topicId].name,
            value: topicId,
         });
      }
      
      topicOptions.value = options;
   } catch (error) {
      console.error(error);
      fetchTopicError.value = true;
   }
   
   await fetchPostsMetadata();
});

onBeforeUnmount(function() {
   // @ts-ignore
   publicTopics = {};
})
</script>

<template>
   <v-dialog
      v-model="props.visible"
      scrollable
      :width="550"
      :height="600">
      <v-card>
         <v-toolbar dark>
            <v-btn icon dark @click="emit('close')">
               <v-icon>mdi-close</v-icon>
            </v-btn>

            <v-toolbar-title>Select Post</v-toolbar-title>
         </v-toolbar>

         <v-card-item>
            <v-row>
               <v-col cols="12" sm="6">
                  <v-select 
                     v-model="topic" 
                     :items="topicOptions" 
                     label="Topic">
                  </v-select>
               </v-col>
               <v-col cols="12" sm="6">
                  <v-select 
                     v-model="sort" 
                     :items="sortOptions" 
                     label="Sort By">
                  </v-select>
               </v-col>
            </v-row>
         </v-card-item>

         <v-card-text v-if="postsList.length > 0">
            <v-list
               :items="postsList"
               item-title="title"
               item-value="_id"
               item-props
               lines="2"
               @click:select="onPostSelect">
               <template v-slot:prepend="{ item }">
                  <v-avatar :rounded="0">
                     <v-img
                        :src="getPostCoverImageURL(item.coverImage.path)"
                        :alt="item.title">
                     </v-img>
                  </v-avatar>
               </template>

               <template v-slot:subtitle="{ item }">
                  {{ topicName(item.topic) }} - {{ formatTime(item.createdAt) }}
               </template>
               <template v-slot:divider>
                  <v-divider inset></v-divider>
               </template>
            </v-list>
         </v-card-text>

         <v-card-text v-else class="mx-auto pt-5">No Posts</v-card-text>

         <v-card-item>
            <v-pagination 
               v-model="pageIndex"
               :length="maxPageIndex" 
               density="comfortable">
            </v-pagination>
         </v-card-item>
      </v-card>
    </v-dialog>

    <v-snackbar v-model="fetchPostError" :timeout="3000" color="error">
      Couldn't get posts
   </v-snackbar>
    <v-snackbar v-model="fetchTopicError" :timeout="3000" color="error">
      Couldn't get topics
   </v-snackbar>
</template>
