<template>
  <div class="signup">
    <h2>Sign up</h2>
    <AuthInput @on-submit="SignUp" message="登録"/>
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
  SignUp(personalInfo:{joid:number,name:string,password:string}){
    console.log(personalInfo);
  }
}
</script>