// @ts-nocheck
import { createRouter, createWebHashHistory } from 'vue-router';
import PostsView from './views/Posts.vue';
import TopicsView from './views/Topics.vue';
import AboutView from './views/About.vue';

const router = createRouter({
   // @ts-ignore
   history: createWebHashHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: '/',
         name: 'posts',
         component: PostsView,
      },
      {
         path: '/topics',
         name: 'topics',
         component: TopicsView,
      },
      {
         path: '/about',
         name: 'about',
         component: AboutView,
      },
   ],
});

export default router;
