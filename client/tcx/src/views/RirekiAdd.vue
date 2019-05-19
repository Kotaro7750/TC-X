<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>
    <ul v-show="!isInputCorrect">
      <li v-for="errMsg in errMsgs" v-bind:key="errMsg">{{errMsg}}</li>
    </ul>
    <input type="number" v-model="rireki.joid">
    <input type="number" v-model="rireki.syubetsu">
    <input type="text" v-model="rireki.about">
    <input type="time" v-model="rireki.startTime">
    <input type="time" v-model="rireki.endTime">
    {{rireki}}
    <button @click="log">送信</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
})

export default class RirekiAdd extends Vue{
    rireki:{
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    } = {
        joid:63,
        syubetsu:3,
        about:"",
        startTime:"",
        endTime:""
    };

    errMsgs: string[] = []

    addRireki():void{
      var url = "http://localhost:8888/rireki/" + "5";
      fetch(url,{
        method: 'POST',
        body: JSON.stringify(this.rireki)
      }).then(response => {
        return response.json();
      }).then(json => {
        console.log(json);
      })
    }

    //validate if input is correct
    get isInputCorrect():boolean{
      this.errMsgs = [];
      if (isNaN(this.rireki.joid) || isNaN(this.rireki.syubetsu)) {
        this.errMsgs.push("joidと業務種別は数字で入力してください");
      }

      if (this.rireki.about == "") {
        this.errMsgs.push("業務内容を入力してください。");
      }

      if (this.rireki.startTime.match("/^23:56$/") == null) {
        this.errMsgs.push("時刻は00:00〜23:59までの範囲で入力してください");
        console.log(this.rireki.startTime.match("/^23\:56$/"));
        console.log(typeof this.rireki.startTime);
      }

      if (this.errMsgs.length != 0) {
        return false;
      }else{
        return true;
      }
    }

    log():void{
      console.log(this.errMsgs);
    }

}
</script>
