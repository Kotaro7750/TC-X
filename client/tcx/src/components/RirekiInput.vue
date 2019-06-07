<template>
  <div class="container-fluid">
    <h3 v-show="isInputTimeStrange">入力した時間を確認してください。このままだと終了時刻は翌日の時刻となります。</h3>
    <span class="fa fa-search form-control-feedback"></span>

    <b-form class="needs-validation" novalidate>
      <b-form inline class="offset-lg-4">
        <b-form-select v-model="rireki.syubetsu" class="col-lg-2 col-3 rounded-pill theme-color">
          <option v-for="syubetsu in syubetsuList" v-bind:key="syubetsu.syubetsu"  v-bind:value="syubetsu.syubetsu">{{syubetsu.name}}</option>
        </b-form-select>

        <b-input-group  class="col-lg-5 col-9">
          <b-form-input  name="about" v-model="rireki.about" class="rounded-pill" 
              v-validate="{ required: true ,regex: /^[ぁ-んァ-ンーa-zA-Z0-9一-龠０-９\-\r]+$/}"
              :state="validateState('about') && !isAboutBlank"
              aria-describedby="about">
          </b-form-input>

          <b-form-invalid-feedback id="about">
            {{$validator.errors.first('about')}}
          </b-form-invalid-feedback>
        </b-input-group>
      </b-form>

      <p></p>

      <b-form inline class="offset-lg-4">
        <b-form-select v-model="rireki.startDay" class="col-lg-2 col-2 rounded-pill theme-color">
          <option v-for="day in days" v-bind:key="day"  v-bind:value="day">{{day}}日</option>
        </b-form-select>

        <b-input-group  class="col-lg-5 col-10 form-group has-feedback feedback-icon ">
          <b-form-input  name="start" v-model="rireki.startTime" type="time" class="time rounded-pill"
              v-validate="{ required: true }" 
              :state="!isInputTimeStrange"
              aria-describedby="start">
          </b-form-input>

          <span class="input-group-text bg-white border-0">~</span>

          <b-form-input  name="end" v-model="rireki.endTime" type="time" class="time rounded-pill"
              v-validate="{ required: true }" 
              :state="!isInputTimeStrange"
              aria-describedby="end">
          </b-form-input>

          <b-form-invalid-feedback id="end">
            <b-alert show variant="warning">Warning Alert</b-alert>
          </b-form-invalid-feedback>

        </b-input-group>
      </b-form>




    <span class ="glyphicon glyphicon-exclamation-sign"></span>
    </b-form>
    <p></p>
    <b-button @click="onSubmit" v-show="isInputCorrect" pill variant="outline-success btn-lg">追加</b-button>

  </div>   
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit,Watch,Inject} from "vue-property-decorator";
import moment from "moment-timezone"
import VeeValidate ,{Validator}from 'vee-validate';

@Component({
})

export default class RirekiInput extends Vue{
    @Inject() $validator!:Validator;
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
      joid:this.$store.getters.userInfo.joid,
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

    isLoading = false;
    isListError = false;
    listError:{} = {};
    syubetsuList:{syubetsu:number,name:string,salary:number}[] = [];

    errMsgs: string[] = [];


    @Emit('on-submit')
    onSubmit(){
      console.log(this.$validator)
      let start:moment.Moment = moment(String(this.year) + "-" + ("0" + String(this.month)).slice(-2) + "-" + ("0" + String(this.rireki.startDay)).slice(-2) + "T" + this.rireki.startTime);
      let end:moment.Moment = moment(String(this.year) + "-" + ("0" + String(this.month)).slice(-2) + "-" + ("0" + String(this.rireki.startDay)).slice(-2) + "T" + this.rireki.endTime);

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

    created(){
      this.getSyubetsuList();
    }

    validateState(ref:string) {
      return !this.$validator.errors.has(ref);
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

    get isAboutBlank():boolean{
      if (this.rireki.about == "") {
        return true;
      }
      return false;
    }

    //validate if input time is strange 
    get isInputTimeStrange():boolean{
      if (this.rireki.startTime == "" || this.rireki.endTime == "") {
        return true;
      }
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

    getSyubetsuList():void{
        this.isListError = false;
        this.isLoading = true;
        this.$store.dispatch('getSyubetsuList').then(() => {
            this.syubetsuList = this.$store.getters.getSyubetsuList
            this.isLoading = false;
        }).catch((error) => {
            this.listError = error;
            this.isListError = true
        });
    }

    clearInput(){
      this.rireki = {
        joid:this.$store.getters.userInfo.joid,
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
  .starttime-pill{
    border-radius: 50pt 50pt
  }

  .endtime-pill{
    border-radius: 50pt 50pt
  }

  .theme-color{
    color: #FFFFFF;
    background: #2D9F91;
  }

  .time.form-control.is-invalid{
    border-color: #ffc107;
    background-image: url();
  }


</style>

