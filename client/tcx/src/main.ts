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
    syubetsuList: [
      {syubetsu: 0,
       name: '',
       salary:  1000,
      },
    ],
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

    getSyubetsuList: (state) => {
      return state.syubetsuList;
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
    getSyubetsuList(state, payload: Array<{syubetsu: number, name: string, salary: number}> ) {
      state.syubetsuList = payload;
    },
  },
  actions: {
    signIn(context, userInfo: {joid: number, name: string, hashedPass: string, token: string}) {
      context.commit('signIn', userInfo);
    },
    signOut(context) {
      context.commit('signOut');
    },

    getSyubetsuList(context) {
      return new Promise((resolve, reject) => {
        context.commit('syubetsuLoading');
        const url = 'http://localhost:8888/syubetsu';
        let syubetsuList: Array<{syubetsu: number, name: string, salary: number}> = [];
        fetch(url, {
            method: 'GET',
        }).then((response) => {
          return response.json();
        }).then((json) => {
          if (json.error != null) {
            reject(json.error);
          } else {
            syubetsuList = json.result;
            context.commit('getSyubetsuList', syubetsuList);
            resolve();
          }
        });
      });
    },
  },
});

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
