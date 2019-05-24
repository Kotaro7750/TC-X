<template>
  <div class="signin">
    <Logo/>
    <h2>Sign In</h2>
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
    var url = "http://localhost:8888/user";
    fetch(url,{
      method: 'GET',
      headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + personalInfo.hashedPass,},
    }).then(response => {
      if (!response.ok) {
      }else{
      }
      return response.json();
    }).then(json => {
        //vuexに登録
        console.log(json.header);
    });

  }

}
</script>

<style scoped>
/* TODO: add css */ 
</style>

