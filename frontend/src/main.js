import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { getString } from './utils/kvstore';
import './style.css';

// Vuetify
import 'vuetify/styles';
import '@mdi/font/css/materialdesignicons.css';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import { VStepper } from 'vuetify/labs/VStepper';

const vuetify = createVuetify({
   components: {
      VStepper,
      ...components,
   },
   directives,
   icons: {
      defaultSet: 'mdi',
   },
   theme: {
      defaultTheme: getString('theme', 'light'),
      themes: {
         light: {
           dark: false,
           colors: {
              primary: '#6200EE',
              'primary-lighten-1': '#7e3ff2',
              'primary-lighten-2': '#9965f4',
              'primary-lighten-3': '#b794f6',
              'primary-lighten-4': '#d4bff9',
              'primary-lighten-5': '#efe5fd',

              'primary-darken-1': '#5300e8',
              'primary-darken-2': '#3d00e0',
              'primary-darken-3': '#1c00db',
              'primary-darken-4': '#0000d6',
           },
         },
         dark: {
            dark: true,
            colors: {
               primary: '#5300e8',
               'primary-lighten-1': '#5300e8',
               'primary-lighten-2': '#3d00e0',
               'primary-lighten-3': '#1c00db',
               'primary-lighten-4': '#0000d6',
               'primary-lighten-5': '#6200EE',

               'primary-darken-1': '#7e3ff2',
               'primary-darken-2': '#9965f4',
               'primary-darken-3': '#b794f6',
               'primary-darken-4': '#d4bff9',
            },
         },
      },
   },
});

const app = createApp(App);

app.use(router);
app.use(vuetify);
app.mount('#app');
