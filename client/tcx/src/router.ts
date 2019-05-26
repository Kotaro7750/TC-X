import Vue from 'vue';
import Router from 'vue-router';

import Home from './views/Home.vue';
import Signup from './views/Signup.vue';
import Signin from './views/Signin.vue';

import Rireki from './views/Rireki.vue';
import Syubetsu from './views/Syubetsu.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup,
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin,
    },
    {
      path: '/rireki',
      name: 'Rireki',
      component: Rireki,
    },
    {
      path: '/syubetsu',
      name: 'Syubetsu',
      component: Syubetsu,
    },
    {
      path: '*',
      name: '404',
      component: Home,
    },
  ],
});
