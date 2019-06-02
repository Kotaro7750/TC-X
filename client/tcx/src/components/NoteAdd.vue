<template>
  <div class="noteadd">
    <h2>ノート追加</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
    <select v-model="selectedYear">
        <option v-for="year in yearList" v-bind:key="year" v-bind:value="year">{{year}}年</option>
    </select>
    <select v-model="selectedMonth">
        <option v-for="month in monthList" v-bind:key="month" v-bind:value="month">{{month}}月</option>
    </select>
    <button @click="addNote">送信</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit} from "vue-property-decorator";
import moment from "moment-timezone"

@Component({
  components: {
  }
})

export default class NoteAdd extends Vue{

    selectedYear: number = Number(moment().year());
    selectedMonth: number = moment().month()+1;

    yearList: number[] = [Number(moment().year()),Number(moment().year())+1];
    monthList: number[] = [1,2,3,4,5,6,7,8,9,10,11,12];

    isError:boolean = false;
    fetchError:{} = {};

    @Emit('on-add')
    onAdd(){
      return ;
    }


    addNote():void{
      this.isError = false;
      let confirmDelete = confirm("作成していいですか？");
      if (confirmDelete == true) {
        var url = process.env.VUE_APP_SERVER_URL+"/note/" + this.selectedYear + "/" + this.selectedMonth;
        fetch(url,{
          method: 'POST',
          headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
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