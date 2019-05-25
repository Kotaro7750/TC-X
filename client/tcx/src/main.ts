import Vue from 'vue';
import Vuex from 'vuex';
import App from './App.vue';
import router from './router';

import BootstrapVue from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
Vue.use(BootstrapVue);
Vue.use(Vuex);

Vue.config.productionTip = false;

export const store = new Vuex.Store({
  state: {
    // when singnin, this represent joid of that user. if not, -1
    userInfo: {
      signInJoid: -1,
      name: '',
      apiToken: '',
    },
  },

  getters: {
    userInfo: (state) => {
      let userInfo = {
        joid: state.userInfo.signInJoid,
        name: state.userInfo.name,
        token: state.userInfo.apiToken,
      };
      return userInfo;
    },
    isSignIn: (state) => {
      if (state.userInfo.signInJoid === -1) {
        return false;
      } else {
        return true;
      }
    },
  },

  mutations: {
    signIn(state: any , payload: {joid: number, name: string, hashedPass:  string, token: string}) {
      state.userInfo.signInJoid = payload.joid;
      state.userInfo.name = payload.name;
      state.userInfo.apiToken = payload.token;
    },
    signOut(state) {
      state.userInfo.signInJoid = -1;
      state.userInfo.apiToken = '';
      state.userInfo.name = '';
    },
  },
  actions: {
    signIn(context, userInfo: {joid: number, name: string, hashedPass: string, token: string}) {
      context.commit('signIn', userInfo);
    },
    signOut(context) {
      context.commit('signOut');
    },
  },
});

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');

