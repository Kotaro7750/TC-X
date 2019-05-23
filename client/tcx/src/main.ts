import Vue from 'vue';
import Vuex from 'vuex';
import App from './App.vue';
import router from './router';

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue);
Vue.use(Vuex);

Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');

const store = new Vuex.Store({
  state: {
    // when singnin, this represent joid of that user. if not, -1
    userInfo: {
      signInJoid: -1,
      apiToken: '',
    },
  },

  getters: {
    joid: (state) => state.userInfo.signInJoid,
    token: (state) => state.userInfo.apiToken,
  },

  mutations: {
    signIn(state: any , payload: {joid: number, token: string}) {
      state.userInfo.signInJoid = payload.joid;
      state.userInfo.apiToken = payload.token;
    },
    signOut(state) {
      state.userInfo.signInJoid = -1;
      state.userInfo.apiToken = '';
    },
  },
});
