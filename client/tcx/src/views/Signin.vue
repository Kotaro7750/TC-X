<template>
  <div class="signin">
    <Logo/>
    <h2>サインイン</h2>
    <div class="alert alert-danger" role="alert" v-show="isError">{{ signInError }}</div>
    <AuthInput @on-submit="SignIn" message="ログイン" />
    <p>まだ登録していない方はこちら
      <router-link to="/signup">Sign Up</router-link>
    </p>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import AuthInput from '@/components/AuthInput.vue';

@Component({
  components: {
    AuthInput,
  }
})

export default class Signin extends Vue{
  isError: boolean = false;
  signInError: string[] = [];

  personalInfo :{
    joid: number,
    name: string,
    hashedPass: string,
  } = {
    joid: 0,
    name: "",
    hashedPass: "",
  };

  SignIn(personalInfo:{joid:number,name:string,hashedPass:string}){
    this.isError = false;
    let url = process.env.VUE_APP_SERVER_URL+  "/user/" + String(personalInfo.joid);
    console.log(url);
    fetch(url,{
      method: 'GET',
      headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + personalInfo.hashedPass,},
    }).then(response => {
      if (!response.ok) {
        this.isError = true;
      }else{
      }
      return response.json();
    }).then(json => {
      if (!this.isError) {
        this.$store.dispatch('signIn',json.result);
        this.$router.push('/');
      }
        this.signInError = json.error;
    });

  }

}
</script>

<style scoped>
/* TODO: add css */ 
</style>

