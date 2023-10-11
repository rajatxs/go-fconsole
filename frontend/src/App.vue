<script setup>
import {ref, onBeforeMount} from 'vue';
import {useTheme} from 'vuetify';
import {RouterView} from 'vue-router';
import {GetAppConfigVariables} from '../wailsjs/go/main/App';
import {setVariables} from './utils/env';
import {setString} from './utils/kvstore';

const loading = ref(true);
const collapsed = ref(true);
const drawerMenuItems = ref([
   {
      title: "Posts",
      value: "posts",
      path: "/",
      icon: "mdi-text-box-outline",
   },
   {
      title: "Topics",
      value: "topics",
      path: "/topics",
      icon: "mdi-format-list-bulleted",
   },
   {
      title: "About",
      value: "about",
      path: "/about",
      icon: "mdi-information-outline",
   },
]);

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

function toggleNavigationDrawer() {
   collapsed.value = !collapsed.value;
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
      <v-navigation-drawer :width="240" :rail="collapsed" permanent>
         <v-list nav>
            <v-list-item
               prepend-icon="mdi-menu"
               @click="toggleNavigationDrawer">
            </v-list-item>

            <template v-for="item of drawerMenuItems" :key="item.path">
               <v-tooltip
                  :text="item.title"
                  :disabled="!collapsed"
                  :open-delay="500">
                  <template v-slot:activator="{ props }">
                     <v-list-item
                        v-bind="props"
                        :prepend-icon="item.icon"
                        :title="item.title"
                        :value="item.value"
                        :to="item.path"
                        color="primary-darken-3">
                     </v-list-item>
                  </template>
               </v-tooltip>
            </template>

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
