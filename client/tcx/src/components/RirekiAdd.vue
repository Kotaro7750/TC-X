<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>
    <RirekiInput @on-submit="log" v-bind:year="year" v-bind:month="month" />
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop} from "vue-property-decorator";
import moment from "moment"
import RirekiInput from '@/components/RirekiInput.vue';

@Component({
  components: {
    RirekiInput,
  }
})

export default class RirekiAdd extends Vue{
    @Prop(Number) year!:number;
    @Prop(Number) month!:number;

    rireki:{
      joid:number,
      syubetsu:number,
      about:string,
      startDay:number,
      startTime:string,
      endDay:number,
      endTime:string,
    } = {
      joid:63,
      syubetsu:3,
      about:"",
      startDay:moment().date(),
      startTime:moment().format('hh:mm'),
      endDay:moment().date(),
      endTime:moment().format('hh:mm')
    };

    //hh:mm => yy-mm-ddThh:mm:ssZ
    get formatRireki(){
      let rireki = this.rireki;

      rireki.startTime = this.year + "-" + this.month + "-" + rireki.startDay + "T" + rireki.startTime.split(':')[0] + ":" + rireki.startTime.split(':')[1]; 
      rireki.endTime = this.year + "-" + this.month + "-" + rireki.startDay + "T" + rireki.endTime.split(':')[0] + ":" + rireki.endTime.split(':')[1]; 
      return rireki;
    }

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

    log(rireki:{day:number,joid:number,syubetsu:number,about:string,startDay:number,startTime:string,endDay:number,endTime:string}):void{
      this.rireki = rireki;
      console.log(this.rireki);
    }
}
</script>
