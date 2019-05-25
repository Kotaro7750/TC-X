<template>
  <div class="signup">
    <h2>Sign up</h2>
    <div class="alert alert-danger" role="alert" v-show="isError">{{ signUpError }}</div>
    <AuthInput @on-submit="onSubmit" message="登録"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop} from "vue-property-decorator";
import AuthInput from '@/components/AuthInput.vue';


@Component({
  components: {
    AuthInput,
  }
})

export default class Signup extends Vue{
  isError: boolean = false;
  signUpError: string[] = [];

  personalInfo :{
    joid: number,
    name: string,
    hashedPass: string,
  } = {
    joid: 0,
    name: "",
    hashedPass: "",
  };

  SignUp(personalInfo:{joid:number,name:string,hashedPass:string}){
    this.isError = false;
    var url = "http://localhost:8888/user";
    fetch(url,{
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(this.personalInfo)
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
      this.signUpError = json.error;
    });
  }

  onSubmit(personalInfo:{joid:number,name:string,hashedPass:string}){
    this.personalInfo = personalInfo;
    this.SignUp(this.personalInfo);
  }
}
</script>

<style scoped>
/* TODO: add css */
</style>
