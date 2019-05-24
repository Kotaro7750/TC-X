<template>
  <div class="input">
    <h3 v-show="isInputTimeStrange">入力した時間を確認してください。このままだと終了時刻は翌日の時刻となります。</h3>

    <ul v-show="!isInputCorrect">
      <li v-for="errMsg in errMsgs" v-bind:key="errMsg">{{errMsg}}</li>
    </ul>
    <input type="number" v-model="rireki.joid">
    <input type="number" v-model="rireki.syubetsu">
    <input type="text" v-model="rireki.about">

    <select v-model="rireki.startDay">
      <option v-for="day in days" v-bind:key="day"  v-bind:value="day">{{day}}日</option>
    </select>
    <input type="time" v-model="rireki.startTime">

    <input type="time" v-model="rireki.endTime">
    <button @click="onSubmit" v-show="isInputCorrect">送信</button>
  </div>   
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit,Watch} from "vue-property-decorator";
import moment from "moment-timezone"

@Component({
})

export default class RirekiInput extends Vue{
    @Prop(Number) month!:number;
    @Prop(Number) year!:number;
    @Prop() initData!:{
      joid:number,
      syubetsu:number,
      about:string,
      startDay:number,
      startTime:string,
      endTime:string,
    };

    @Watch('month') onMonthChanged(){
      this.days = this.dayOfMonth(this.month);
    }

    rireki:{
      joid:number,
      syubetsu:number,
      about:string,
      startDay:number,
      startTime:string,
      endTime:string,
    } = this.initData == null ? 
    //when not initialized
    {
      joid:63,
      syubetsu:3,
      about:"",
      startDay:moment().date(),
      startTime:moment().tz("Asia/Tokyo").format('HH:mm'),
      endTime:moment().tz("Asia/Tokyo").format('HH:mm')
    } :{
      joid:this.initData.joid,
      syubetsu:this.initData.syubetsu,
      about:this.initData.about,
      startDay:this.initData.startDay,
      startTime:moment(this.initData.startTime).tz("Asia/Tokyo").format('HH:mm'),
      endTime:moment(this.initData.endTime).tz("Asia/Tokyo").format('HH:mm')
    };

    days = this.dayOfMonth(this.month);

    errMsgs: string[] = [];

    @Emit('on-submit')
    onSubmit(){
      let start:moment.Moment = moment(String(this.year) + "-" + String(this.month) + "-" + String(this.rireki.startDay) + " " + this.rireki.startTime);
      let end:moment.Moment = moment(String(this.year) + "-" + String(this.month) + "-" + String(this.rireki.startDay) + " " + this.rireki.endTime);

      if (end.isBefore(start)) {
        end = end.add(1,'d');
      }

      let formattedrireki = {
        joid:Number(this.rireki.joid),
        syubetsu:Number(this.rireki.syubetsu),
        about:this.rireki.about,
        startTime:start.format("YYYY-MM-DDTHH:mm:ss+09:00"),
        endTime:end.format("YYYY-MM-DDTHH:mm:ss+09:00"),
      }
        return formattedrireki;
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
      let startTimeNum = moment(this.rireki.startTime,'HH:mm');
      let endTimeNum = moment(this.rireki.endTime,'HH:mm');
      if (endTimeNum <= startTimeNum) {
        return true;
      }
      return false;
    }

    //return array filled of day of month
    dayOfMonth(month:number){
      let days=  Array(moment(String(month),"MM").daysInMonth());
      for (let i:number = 0; i < days.length;i++){
        days[i] = i + 1;
      }
      return days;
    }

    clearInput(){
      this.rireki = {
        joid:63,
        syubetsu:3,
        about:"",
        startDay:moment().date(),
        startTime:moment().format('HH:mm'),
        endTime:moment().format('HH:mm')
      };
    }
}
</script>

<style scoped>
/* add css */ 

</style>

