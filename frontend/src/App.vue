<script setup>
import { RouterView } from 'vue-router';
</script>

<template>
   <v-card>
      <v-layout>
         <v-navigation-drawer permanent theme="light">
            <v-list nav>
               <v-list-item
                  prepend-icon="mdi-text-box-multiple"
                  title="Posts"
                  value="posts" to="/"></v-list-item>
               <v-list-item prepend-icon="mdi-list-box" title="Topics" value="topics" to="/topics"></v-list-item>
               <v-list-item
                  prepend-icon="mdi-information"
                  title="About"
                  value="about" to="/about"></v-list-item>
            </v-list>
         </v-navigation-drawer>

         <v-main class="py-5">
            <RouterView />
         </v-main>
      </v-layout>
   </v-card>
</template>

<style></style>
