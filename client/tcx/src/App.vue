<template>
  <div id="app">
    <b-navbar toggleable="lg" type="dark" variant="info">
        <b-navbar-brand to="/">TC-X</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse" right></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav >
            <b-navbar-nav v-if="!isSignIn">
                <b-nav-item to="/signup">登録</b-nav-item>
                <b-nav-item to="/signin">サインイン</b-nav-item>
            </b-navbar-nav>
            <b-navbar-nav v-else>
                <b-nav-item to="/rireki">履歴ノート</b-nav-item>
                <b-nav-item to="/syubetsu">業務種別管理</b-nav-item>
                <b-nav-item to="/note">ノート管理</b-nav-item>
                <b-nav-item to="/" @click.native="signOut" >サインアウト</b-nav-item>
            </b-navbar-nav>
    </b-collapse>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Watch} from "vue-property-decorator";

@Component({
  components: {
  }
})

export default class App extends Vue{
  personalInfo :{
    joid: number,
    name: string,
    token: string,
  } = this.$store.getters.userInfo;

  @Watch('userInfo') onUserInfoChanged(){
    this.personalInfo = this.userInfo;
  }

  signOut(){
    this.$router.push('/');
    this.$store.dispatch('signOut');
  }

  get userInfo(): {joid:number,name:string,token:string}{
    let userInfo = this.$store.getters.userInfo;
    return userInfo;
  }

  get isSignIn():boolean{
    return this.$store.getters.isSignIn;
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
