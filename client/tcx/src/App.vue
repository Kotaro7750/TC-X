<template>
  <div id="app" >
    <b-navbar toggleable="lg"  class="navbar-custom">
        <b-navbar-brand to="/" class="navbar-light navbar-brand">TC-X</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse" right></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav >
            <b-navbar-nav v-if="!isSignIn" class="ml-auto">
                <b-nav-item to="/signup">登録</b-nav-item>
                <b-nav-item to="/signin">サインイン</b-nav-item>
            </b-navbar-nav>
            <b-navbar-nav v-else class="ml-auto">
                  <b-nav-item to="/rireki" v-show="!isSuperUser">履歴ノート</b-nav-item>
                  <b-nav-item to="/syubetsu" v-show="isSuperUser">業務種別管理</b-nav-item>
                  <b-nav-item to="/note" v-show="isSuperUser">ノート管理</b-nav-item>
                  <b-nav-item to="/" @click="signOut" >サインアウト</b-nav-item>
            </b-navbar-nav>
    </b-collapse>
    </b-navbar>
    <p></p>
    <router-view />
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
    console.log("signout!!")
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

  get isSuperUser():boolean{
    return this.$store.getters.isSuperUser;
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

/* change the background color */
.navbar-custom {
    background-color: #2d9F91;
}
/* change the brand and text color */
.navbar-custom.navbar-light .navbar-brand, 
.navbar-custom.navbar-light .navbar-brand:hover,
.navbar-custom.navbar-light .navbar-brand:focus {
    color: #FFFFFF;
}
/* change the link color */
.navbar-custom .navbar-nav .nav-link {
    color: rgba(255, 255, 255, 0.8);
}
/* change the color of active or hovered links */
.navbar-custom .nav-item.active .nav-link,
.navbar-custom .nav-item:focus .nav-link,
.navbar-custom .nav-item:hover .nav-link {
    color: #ffffff;
}
</style>
