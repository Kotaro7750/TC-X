<template>
  <div class="syubetsuadd">
    <h2>種別追加</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
    <SyubetsuInput @on-submit="onSubmit" ref="input"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit} from "vue-property-decorator";
import SyubetsuInput from '@/components/SyubetsuInput.vue';

@Component({
  components: {
    SyubetsuInput,
  }
})

export default class SyubetsuAdd extends Vue{

    syubetsu:{
      syubetsu:number,
      name:string,
      salary:number,
    } = {
      syubetsu:0,
      name:"",
      salary:1000,
    };

    isError:boolean = false;
    fetchError:{} = {};

    @Emit('on-add')
    onAdd(){
      return ;
    }

    addSyubetsu():void{
      this.isError = false;
      var url =  process.env.VUE_APP_SERVER_URL+"/syubetsu";
      fetch(url,{
        method: 'POST',
        headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
        body: JSON.stringify(this.syubetsu)
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

    onSubmit(syubetsu:{syubetsu:number,name:string,salary:number}):void{
      this.syubetsu = syubetsu;
      this.addSyubetsu();
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
</style>