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
              primary: '#00695f',
              secondary: '#ffea00'
            }
         },
         dark: {
            dark: true,
            colors: {
              primary: '#009688',
              secondary: '#ffea00'
            }
         },
      },
   },
});

const app = createApp(App);

app.use(router);
app.use(vuetify);
app.mount('#app');
