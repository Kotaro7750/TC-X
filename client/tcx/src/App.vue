<template>
  <div id="app">
    <nav class="navbar navbar-expand-sm navbar-dark bg-dark mt-3 mb-3">
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav4" aria-controls="navbarNav4" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <a class="navbar-brand" href="/">TC-X</a>
        <div class="collapse navbar-collapse justify-content-end">
            <ul class="navbar-nav" v-if="!isSignIn">
                <li class="nav-item active">
                  <router-link to="/signup" class=nav-link>SignUp</router-link>
                </li>
                <li class="nav-item active">
                  <router-link to="/signin" class=nav-link>SignIn</router-link>
                </li>
            </ul>
            <ul class="navbar-nav" v-else>
                <li>{{personalInfo.joid}}</li>
                <li>{{personalInfo.name}}</li>
                <li class="nav-item active">
                  <router-link to="/" @click.native="signOut" class=nav-link>SignOut</router-link>
                </li>
            </ul>
        </div>
    </nav>
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
