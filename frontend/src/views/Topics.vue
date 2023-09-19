<script setup>
import { onMounted, ref, watch } from 'vue';
import { groupArray, getPostTopicImageUrl } from '../utils';
import { getPublicTopics, getPrivateTopics } from '../utils/topic';
import topics from '../data/topics.json';
import Loader from '../components/Loader.vue';

/** @type {import('vue').Ref<any[][]>} */
const topicGroups = ref([]);

const scope = ref('public');
const loading = ref(true);
const fetchErrorSnackbar = ref(false);
const previewErrorSnackbar = ref(false);

const scopeOptions = ref([
   {
      title: 'Public',
      value: 'public',
   },
   {
      title: 'Private',
      value: 'private',
   },
]);

/**
 * Opens live topic page
 * @param {string} id - Topic Id
 */
async function preview(id) {
   const url = `https://www.fivemin.in/t/${id}`;

   try {
      window.open(url, '_blank');
   } catch (error) {
      previewErrorSnackbar.value = true;
   }
}

async function fetchTopics() {
   loading.value = true;

   try {
      let _topics = [];

      if (scope.value === "public") {
         _topics = getPublicTopics();
      } else {
         _topics = getPrivateTopics();
      }

      topicGroups.value = groupArray(_topics);
   } catch (error) {
      fetchErrorSnackbar.value = true;
   }
   loading.value = false;
}

watch(scope, fetchTopics);
onMounted(fetchTopics);
</script>

<template>
   <v-container>
      <div class="text-h4 mb-5">Topics</div>

      <v-row>
         <v-col cols="12" sm="4"> </v-col>
         <v-col cols="12" sm="4"> </v-col>
         <v-col cols="12" sm="4">
            <v-select label="Scope" :items="scopeOptions" v-model="scope"></v-select>
         </v-col>
      </v-row>

      <v-container>
         <Loader v-if="loading" message="Getting topics" />
         <v-row v-else v-for="(group, groupIndex) of topicGroups" :key="groupIndex">
            <v-col v-for="topic of group" :key="topic.id" cols="12" sm="4">
               <v-card class="mx-auto">
                  <v-img class="align-end text-white" height="200" :src="getPostTopicImageUrl(topic.thumbPath)" cover>
                  </v-img>

                  <v-card-item>
                     <v-card-title>{{ topic.name }}</v-card-title>
                  </v-card-item>

                  <v-card-actions v-if="scope === 'public'">
                     <v-btn color="primary" @click="preview(topic.id)">Preview</v-btn>
                  </v-card-actions>
               </v-card>
            </v-col>
         </v-row>
      </v-container>

      <v-snackbar v-model="fetchErrorSnackbar" :timeout="5000"> Couldn't get topics </v-snackbar>
      <v-snackbar v-model="previewErrorSnackbar" :timeout="5000"> Couldn't open preview </v-snackbar>
   </v-container>
</template>
