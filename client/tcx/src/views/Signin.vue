<template>
  <div class="signin">
    <Logo/>
    <h2>Sign in</h2>
    <AuthInput @on-submit="SignIn" message="ログイン" />
    <p>You don't have an account?
      <router-link to="/signup">Create account now!</router-link>
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
  SignIn(personalInfo:{joid:number,name:string,hashedPass:string}){
    console.log(personalInfo);
  }

}
</script>

<style scoped>
/* TODO: add css */ 
</style>

