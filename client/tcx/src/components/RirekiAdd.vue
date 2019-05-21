<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
    <RirekiInput @on-submit="onSubmit" v-bind:year="year" v-bind:month="month" ref="input"/>
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
      startTime:string,
      endTime:string,
    } = {
      joid:63,
      syubetsu:3,
      about:"",
      startTime:moment().format('YYYY-MM-DDTHH:mm:ssZ'),
      endTime:moment().format('YYYY-MM-DDTHH:mm:ssZ'),
    };

    isError:boolean = false;
    fetchError:{} = {};

    addRireki():void{
      this.isError = false;
      var url = "http://localhost:8888/rireki/" + String(this.month);
      fetch(url,{
        method: 'POST',
        body: JSON.stringify(this.rireki)
      }).then(response => {
        if (!response.ok) {
          this.isError = true;
        }else{
          this.clearInput();
        }
        return response.json();
      }).then(json => {
        this.fetchError = json.error;
      })
    }

    onSubmit(rireki:{joid:number,syubetsu:number,about:string,startTime:string,endTime:string}):void{
      this.rireki = rireki;
      this.addRireki();
    }

    clearInput(){
      let inputForm :any = this.$refs.input;
      inputForm.clearInput();
    }
}
</script>

<style scoped>
.fetchError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }
</style>

