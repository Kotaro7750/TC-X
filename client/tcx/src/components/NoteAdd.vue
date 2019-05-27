<template>
  <div class="noteadd">
    <h2>ノート追加</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
    <NoteInput @on-submit="onSubmit" ref="input"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Emit} from "vue-property-decorator";
import moment from "moment-timezone"
import NoteInput from '@/components/NoteInput.vue';

@Component({
  components: {
    NoteInput,
  }
})

export default class NoteAdd extends Vue{

    selectedYear: number = moment().year();
    selectedMonth: number = moment().month()+1;

    isError:boolean = false;
    fetchError:{} = {};

    @Emit('on-add')
    onAdd(){
      return ;
    }

    addNote():void{
      this.isError = false;
      var url = "http://localhost:8888/note";
      fetch(url,{
        method: 'POST',
        headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
        body: JSON.stringify(this.note)
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

    onSubmit(note:{note:number,name:string,salary:number}):void{
      this.note = note;
      this.addNote();
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