<template>
  <div class="signup">
    <Logo/>
    {{ alert }}
    <h2>Sign up</h2>
      <label>Name</label>
      <input v-model=joid type="Number" placeholder="Joid">
      <input v-model=name type="text" placeholder="Name">
    <button @click="addUser">Sign up</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop} from "vue-property-decorator";
import { router } from "vue-router";

import Logo from '@/components/Logo.vue';

@Component({
  components: {
    Logo,
  }
})

export default class Signup extends Vue{
  //this.joid とかってfetchすると400えらー
  //joidが"　"joid"　ってなるのでだめ　数字にしないと
  joid!: Number;
  name!: string;
  alert: string = "";

  created(){
    this.alert = "";
  }

  addUser(){
    if (this.validateInput()) {
      fetch('http://localhost:8888/users',{
        method: 'POST',
        body: JSON.stringify({
          joid: this.joid,
          name: this.name,
        })
      }).then(response =>{
          return response.json();
        }).then(res => {
          console.log(res);
        });

      this.joid = null;
      this.name = null;
      this.$router.push({path: 'about'})
    }else{
      this.alert = "Please input joid and name"
      this.joid = null;
      this.name = null;
    }
  }

  validateInput(){
    //void string is not allowed
    if (this.joid == null || this.name == null ){
      return false;
    }
    //joid must be a number
    return true;
  }
}
</script>