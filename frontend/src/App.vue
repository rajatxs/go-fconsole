<script setup>
import { ref, onBeforeMount } from 'vue';
import { RouterView } from 'vue-router';
import { GetAppConfigVariables } from '../wailsjs/go/main/App';
import { setVariables } from './utils/env';

const loading = ref(true);

/** Load default configuration from main process */
async function preload() {
   if (!loading.value) {
      loading.value = true;
   }

   const env = await GetAppConfigVariables();
   setVariables(env);
   loading.value = false;
}

onBeforeMount(preload);
</script>

<template>
   <v-card>
      <v-layout>
         <!-- Sidenav -->
         <v-navigation-drawer permanent theme="light" :width="220">
            <v-list nav>
               <v-list-item prepend-icon="mdi-text-box-multiple" title="Posts" value="posts" to="/"></v-list-item>
               <v-list-item prepend-icon="mdi-list-box" title="Topics" value="topics" to="/topics"></v-list-item>
               <v-list-item prepend-icon="mdi-information" title="About" value="about" to="/about"></v-list-item>
            </v-list>
         </v-navigation-drawer>

         <v-main class="py-5">
            <!-- Preconnect progresbar -->
            <v-dialog v-model="loading" :scrim="false" persistent :width="320" :height="80">
               <v-card color="primary" :height="80">
                  <v-card-text>
                     <span>Connecting</span>
                     <v-progress-linear indeterminate color="white" class="mt-4"></v-progress-linear>
                  </v-card-text>
               </v-card>
            </v-dialog>

            <RouterView v-if="!loading" />
         </v-main>
      </v-layout>
   </v-card>
</template>
