<template>
  <div class="container-fluid">
    <div class="row" v-show="isInputError">
      <b-alert show variant="warning" class="col-md-6 offset-md-3">
        <ul v-show="isInputError">
          <li v-for="inputError in inputErrors" v-bind:key="inputError">{{inputError}}</li>
        </ul>
      </b-alert>
    </div>
    <div class="row">
      <b-input-group prepend="Joid" class="col-md-4 offset-md-4">
        <b-form-input type="number" v-model="personalInfo.joid" size="20"></b-form-input>
      </b-input-group>
    </div>
    <div class="row">
      <b-input-group prepend="名前" class="col-md-4 offset-md-4">
        <b-form-input type="text" maxlength="100" v-model="personalInfo.name"></b-form-input>
      </b-input-group>
    </div>
    <div class="row">
      <b-input-group prepend="パスワード" class="col-md-4 offset-md-4">
        <b-form-input type="password" maxlength="100" v-model="personalInfo.password"></b-form-input>
      </b-input-group>
    </div>
    <div class="row">
      <b-button pill v-show="!isInputError" @click="onSubmit" class="col-md-2 offset-md-5">{{message}}</b-button>
    </div>
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
      if (this.personalInfo.joid < 0) {
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
ul {
  list-style: none
}

</style>
