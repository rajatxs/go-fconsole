<script setup>
import {ref, onBeforeMount} from 'vue';
import {useTheme} from 'vuetify';
import {RouterView} from 'vue-router';
import {GetAppConfigVariables} from '../wailsjs/go/main/App';
import {setVariables} from './utils/env';
import {setString} from './utils/kvstore';

const loading = ref(true);
const theme = useTheme();

/** Load default configuration from main process */
async function preload() {
   if (!loading.value) {
      loading.value = true;
   }

   const env = await GetAppConfigVariables();
   setVariables(env);
   loading.value = false;
}

/**
 * Sync app theme with system settings
 * @param {any} event
 */
function updateTheme(event) {
   const currentTheme = theme.global.current.value.dark ? 'dark' : 'light';

   if (event.matches) {
      // switch to dark theme
      if (currentTheme !== 'dark') {
         theme.global.name.value = 'dark';
         setString('theme', 'dark');
      }
   } else {
      // switch to light theme
      if (currentTheme !== 'light') {
         theme.global.name.value = 'light';
         setString('theme', 'light');
      }
   }
}

onBeforeMount(async () => {
   const themeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
   themeMediaQuery.onchange = updateTheme;
   updateTheme(themeMediaQuery);

   await preload();
});
</script>

<template>
   <v-layout>
      <!-- Sidenav -->
      <v-navigation-drawer permanent :width="240">
         <v-list nav>
            <v-list-item
               prepend-icon="mdi-text-box-multiple"
               title="Posts"
               value="posts"
               to="/">
            </v-list-item>
            <v-list-item
               prepend-icon="mdi-list-box"
               title="Topics"
               value="topics"
               to="/topics">
            </v-list-item>
            <v-list-item
               prepend-icon="mdi-information"
               title="About"
               value="about"
               to="/about">
            </v-list-item>
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
</template>

<style>
#app {
   height: 100vh;
}
</style>
