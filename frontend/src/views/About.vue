<script setup>
import {ref, onMounted} from 'vue';
import {GetVersions} from '../../wailsjs/go/main/App';
import {OpenBrowser} from '../../wailsjs/go/main/App';

let fetchErrorSnackbar = ref(false);
let sendFeedbackSnackbar = ref(false);

/** @type {import('vue').Ref<any>} */
let versions = ref({});

async function sendFeedback() {
   try {
      await OpenBrowser("mailto:rxx256+fconsole@outlook.com", false);
   } catch (error) {
      sendFeedbackSnackbar.value = true;
   }
}

async function fetchVersionInfo() {
   try {
      const _versions = await GetVersions();
      versions.value = _versions;
   } catch (error) {
      fetchErrorSnackbar.value = true;
   }
}

onMounted(fetchVersionInfo);
</script>

<template>
   <v-container>
      <div class="text-h4 mb-5">About</div>

      <v-row>
         <v-col cols="6">
            <v-list lines="two">
               <v-list-item title="Version" :subtitle="versions.app"></v-list-item>
               <v-list-item title="Date" :subtitle="versions.date"></v-list-item>
               <v-list-item title="Wails Version" :subtitle="versions.wails"></v-list-item>
               <v-list-item title="Go Version" :subtitle="versions.go"></v-list-item>
               <v-list-item title="Home Directory" :subtitle="versions.homedir"></v-list-item>
               <v-list-item
                  title="WebView2 Runtime Version"
                  :subtitle="versions.webview2"></v-list-item>
               <v-list-item title="OS" :subtitle="versions.os"></v-list-item>
               <v-list-item title="Architecture" :subtitle="versions.arch"></v-list-item>
            </v-list>
         </v-col>

         <v-col cols="6">
            <p class="text-h6">Additional Info:</p>
            <br />
            <p>
               <a
                  href="https://www.fivemin.in"
                  target="_blank"
                  class="text-decoration-none text-primary">
                  <span>Fivemin Website</span>&nbsp;<v-icon size="x-small">mdi-open-in-new</v-icon>
               </a>
            </p>
            <br />
            <div>
               <a
                  href="https://github.com/rajatxs/go-fconsole"
                  target="_blank"
                  class="text-decoration-none text-primary">
                  <span>GithHub Repository</span>&nbsp;<v-icon size="x-small">mdi-open-in-new</v-icon>
               </a>
            </div>
            <br />
            <v-btn color="primary" @click="sendFeedback">
               Send Feedback
            </v-btn>
         </v-col>
      </v-row>

      <v-snackbar v-model="fetchErrorSnackbar" :timeout="5000"> Couldn't get info </v-snackbar>
      <v-snackbar v-model="sendFeedbackSnackbar" :timeout="5000">
         Couldn't open mail app
      </v-snackbar>
   </v-container>
</template>
