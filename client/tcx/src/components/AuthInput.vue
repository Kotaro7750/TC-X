<template>
  <div class="authinput">
    <ul v-show="isInputError">
      <li v-for="inputError in inputErrors" v-bind:key="inputError">{{inputError}}</li>
    </ul>
    <input type="number"  v-model="personalInfo.joid">
    <input type="text" maxlength="100" v-model="personalInfo.name">
    <input type="password" maxlength="100" v-model="personalInfo.password">
    <button v-show="!isInputError" @click="onSubmit">{{message}}</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit} from "vue-property-decorator";
import sha256 from "js-sha256";
@Component({
})

export default class AuthInput extends Vue{
    @Prop() message!:string;

    personalInfo:{
        joid: number,
        name: string,
        password: string,
    } = {
      joid: 0,
      name: "",
      password: "",
    };

    inputErrors:string[] = [];

    @Emit('on-submit')
    onSubmit(){
        let hashedPass = sha256.sha256(this.personalInfo.password);
        console.log(hashedPass);
        let personalInfoWithHashedPass = {
            joid:Number(this.personalInfo.joid),
            name:this.personalInfo.name,
            hashedPass:hashedPass,
        };
        return  personalInfoWithHashedPass;
    }

    get isInputError(){
        this.inputErrors = [];
      if (this.personalInfo.joid == 0) {
        this.inputErrors.push("Joidを入力してください")
      }
      if (this.personalInfo.name == "") {
        this.inputErrors.push("名前を入力してください")
      }
      if (this.personalInfo.password == "") {
        this.inputErrors.push("パスワードを入力してください")
      }
      if (this.inputErrors.length) {
        return true;
      }
      return false;
    }
}
</script>

// TODO: add css
<style scoped>

</style>
