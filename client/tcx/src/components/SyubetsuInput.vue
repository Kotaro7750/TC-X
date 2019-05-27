<template>
  <div class="input">
    <ul v-show="!isInputCorrect">
      <li v-for="errMsg in errMsgs" v-bind:key="errMsg">{{errMsg}}</li>
    </ul>
    <p>
      <label for="syubetsu">業務種別</label>
      <input type="number" id="syubetsu" min="0" v-model="syubetsu.syubetsu">
    </p>

    <p>
      <label for="name">業務内容</label>
      <input type="text" id="name" v-model="syubetsu.name">
    </p>

    <p>
      <label for="salary">時給</label>
      <input type="number" id="salary" min="0" v-model="syubetsu.salary">
    </p>
    <button @click="onSubmit" v-show="isInputCorrect">送信</button>
  </div>   
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit,Watch} from "vue-property-decorator";

@Component({
})

export default class SyubetsuInput extends Vue{
    @Prop() initData!:{
      syubetsu:number,
      name:string,
      salary:number,
    };


    syubetsu:{
      syubetsu:number,
      name:string,
      salary:number,
    } = this.initData == null ? 
    //when not initialized
    {
      syubetsu:0,
      name:"",
      salary:1000,
    } :{
      syubetsu:this.initData.syubetsu,
      name:this.initData.name,
      salary:this.initData.salary,
    };


    errMsgs: string[] = [];

    @Emit('on-submit')
    onSubmit(){
      let formattedsyubetsu = {
        syubetsu:Number(this.syubetsu.syubetsu),
        name:this.syubetsu.name,
        salary:Number(this.syubetsu.salary),
      }
        return formattedsyubetsu;
    }

    //validate if input is correct
    get isInputCorrect():boolean{
      this.errMsgs = [];
      if (isNaN(this.syubetsu.syubetsu) || isNaN(this.syubetsu.salary)) {
        this.errMsgs.push("業務種別と時給は数字で入力してください");
      }

      if (this.syubetsu.syubetsu < 0 || this.syubetsu.salary < 0) {
        this.errMsgs.push("業務種別と時給は正の値を入力してください");
      }

      if (this.syubetsu.name == "") {
        this.errMsgs.push("業務内容を入力してください。");
      }

      if (this.errMsgs.length != 0) {
        return false;
      }else{
        return true;
      }
    }

    clearInput(){
      this.syubetsu = {
        syubetsu:0,
        name:"",
        salary:1000,
      };
    }
}
</script>

<style scoped>
/* add css */ 

</style>

