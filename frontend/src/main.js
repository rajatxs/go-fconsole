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

const vuetify = createVuetify({
   components,
   directives,
   icons: {
      defaultSet: 'mdi',
   },
   theme: {
      defaultTheme: getString('theme', 'light'),
   },
});

const app = createApp(App);

app.use(router);
app.use(vuetify);
app.mount('#app');
