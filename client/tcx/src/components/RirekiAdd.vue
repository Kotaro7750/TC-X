<template>
  <div class="rirekiadd">
    <h2 class="left col-lg-3 ">履歴追加</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
    <RirekiInput @on-submit="onSubmit" v-bind:year="year" v-bind:month="month" ref="input"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit} from "vue-property-decorator";
import moment from "moment-timezone"
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
      joid:this.$store.getters.userInfo.joid,
      syubetsu:3,
      about:"",
      startTime:moment().format('YYYY-MM-DDTHH:mm:ss'),
      endTime:moment().format('YYYY-MM-DDTHH:mm:ss'),
    };

    isError:boolean = false;
    fetchError:{} = {};

    @Emit('on-add')
    onAdd(){
      return ;
    }

    addRireki():void{
      this.isError = false;
      var url = process.env.VUE_APP_SERVER_URL+"/rireki/" + String(this.year) + "/"+ String(this.month) + "/" + String(this.$store.getters.userInfo.joid);
      fetch(url,{
        method: 'POST',
        headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
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
        this.onAdd();
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
/* TODO: add css */
.fetchError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }

.left{
  text-align: left;
}
</style>

