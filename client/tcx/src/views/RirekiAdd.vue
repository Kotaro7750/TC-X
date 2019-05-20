<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>

    <h3 v-show="isInputTimeStrange">入力した時間を確認してください。このままだと終了時刻は翌日の時刻となります。</h3>

    <ul v-show="!isInputCorrect">
      <li v-for="errMsg in errMsgs" v-bind:key="errMsg">{{errMsg}}</li>
    </ul>
    <input type="number" v-model="rireki.joid">
    <input type="number" v-model="rireki.syubetsu">
    <input type="text" v-model="rireki.about">
    <input type="time" v-model="rireki.startTime">
    <input type="time" v-model="rireki.endTime">
    <button @click="addRireki" v-show="isInputCorrect">送信</button>
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

    errMsgs: string[] = [];

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

      if (this.rireki.startTime.match(/([0-1][0-9]|2[0-3]):[0-5][0-9]/) == null || this.rireki.endTime.match(/([0-1][0-9]|2[0-3]):[0-5][0-9]/) == null) {
        this.errMsgs.push("時刻は00:00〜23:59までの範囲で入力してください");
      }

      if (this.errMsgs.length != 0) {
        return false;
      }else{
        return true;
      }
    }

    //validate if input time is strange 
    get isInputTimeStrange():boolean{
      let startTimeNum :number = Number(this.rireki.startTime.split(':')[0])*100 +Number(this.rireki.startTime.split(':')[1]); 
      let endTimeNum :number = Number(this.rireki.endTime.split(':')[0])*100 +Number(this.rireki.endTime.split(':')[1]); 
      if (endTimeNum <= startTimeNum) {
        return true;
      }
      return false;
    }

    log():void{
      console.log(this.errMsgs);
    }

}
</script>
